package service

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"log"
	"mall-go/app/lottery/internal/dao"
	"mall-go/app/lottery/internal/model"
	"mall-go/app/lottery/internal/redis"
	myjwt "mall-go/common/middleware"
	"net/http"
	"strconv"
)

// ErrMsgKey Visible for testing
const ErrMsgKey = "errMsg"
const DataKey = "data"

// FetchCoupon 秒杀优惠券，seller不允许获取优惠券
// TODO 应该可以抽象为切面来处理的，但是go似乎没有切面的概念，中间件似乎不太适合做切面
func FetchCoupon(ctx *gin.Context) {
	// 这里是业务逻辑的「授权(Authorization)」，用于判断合法用户是否有权访问当前资源(例如，某个用户是否可以获取优惠券，某个角色是否有权限执行某个操作)
	claims := ctx.MustGet("claims").(*myjwt.CustomClaims)
	if claims == nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{ErrMsgKey: "Not Authorized."})
		return
	}

	// 后续直接由gin Web框架根据gin.Context的内容打包成HTTP响应以及发送等

	if claims.Kind == "seller" {
		ctx.JSON(http.StatusUnauthorized, gin.H{ErrMsgKey: "Sellers aren't allowed to get coupons."})
		return
	}

	// HTTP请求参数
	paramSellerName := ctx.Param("username")
	paramCouponName := ctx.Param("name")

	// ---用户抢优惠券。后面需要高并发处理---
	// 先在缓存执行原子性的秒杀操作。将原子性地完成"判断能否秒杀-执行秒杀"的步骤
	_, err := redis.CacheAtomicSecKill(claims.Username, paramSellerName, paramCouponName)
	if err == nil {
		coupon := redis.GetCoupon(paramCouponName)
		// 交给[协程]完成数据库写入操作
		SecKillChannel <- secKillMessage{claims.Username, coupon}
		ctx.JSON(http.StatusCreated, gin.H{ErrMsgKey: ""})
		return
	} else {
		if redis.IsRedisEvalError(err) {
			log.Printf("Server error" + err.Error())
			ctx.JSON(http.StatusInternalServerError, gin.H{ErrMsgKey: err.Error()})
			return
		} else {
			log.Println("Fail to fetch coupon. " + err.Error())
			ctx.JSON(http.StatusNoContent, gin.H{})
			return
		}
	}
}

const (
	couponPageSize int64 = 20
)

// 工具函数 输出查询错误
func outputQueryError(ctx *gin.Context, err error) {
	if gorm.IsRecordNotFoundError(err) {
		log.Println("Record not found.")
		ctx.JSON(http.StatusBadRequest, gin.H{ErrMsgKey: "Record not found."})
	} else {
		log.Println("Query error.")
		ctx.JSON(http.StatusBadRequest, gin.H{ErrMsgKey: "Query error."})
	}
}

// 根据page页数, 取得合理切片范围的coupons
// 需要保证切片索引startIndex, endIndex不越界
func getValidCouponSlice(allCoupons []model.Coupon, page int64) []model.Coupon {
	if len(allCoupons) == 0 {
		return allCoupons
	}
	couponLen := int64(len(allCoupons))
	startIndex := page * couponPageSize
	endIndex := page*couponPageSize + couponPageSize
	if startIndex < 0 {
		startIndex = 0
	} else if startIndex > couponLen {
		startIndex = couponLen
	}
	if endIndex < 1 {
		if couponLen < couponPageSize {
			endIndex = couponLen
		} else {
			endIndex = couponPageSize
		}
	} else if endIndex > couponLen {
		endIndex = couponLen
	}
	return allCoupons[startIndex:endIndex]
}

// 数据长度为空则返回204,否则返回200
func getDataStatusCode(len int) int {
	if len == 0 {
		return http.StatusNoContent
	} else {
		return http.StatusOK
	}
}

// GetCoupons 查询优惠券
func GetCoupons(ctx *gin.Context) {
	// 登陆检查token
	claims := ctx.MustGet("claims").(*myjwt.CustomClaims)
	if claims == nil {
		log.Println("Not Authorized.")
		ctx.JSON(http.StatusUnauthorized, gin.H{ErrMsgKey: "Not Authorized."})
		return
	}

	queryUserName, queryPage := ctx.Param("username"), ctx.Query("page")

	// 检查page参数, TODO：全部下标改为从1开始
	// TODO: 要不要改成BindJSON
	var page int64
	var tmpPage int64
	if queryPage == "" {
		tmpPage = 1
	} else {
		var err error
		tmpPage, err = strconv.ParseInt(ctx.Query("page"), 10, 64)
		if err != nil {
			log.Println("Wrong format of page.")
			ctx.JSON(http.StatusBadRequest, gin.H{ErrMsgKey: "Wrong format of page."})
			return
		}
	}

	// 数据库从0开始，但是查找从1开始
	page = tmpPage - 1

	//log.Printf("Querying coupon with name %s, page %d\n", queryUserName, page)
	// TODO: 查询用户需要从缓存查，这里需要改，是有错的。在主goroutine里查询数据库，会极大的拖慢处理速度
	// 查找对应用户
	queryUser := model.User{Username: queryUserName}
	queryErr := dao.Db.Where(&queryUser).
		First(&queryUser).Error
	if queryErr != nil {
		outputQueryError(ctx, queryErr)
		return
	}

	// 根据用户名查找其拥有/创建的优惠券
	if queryUserName == claims.Username {
		// 查询名与用户名相同，返回查询名用户拥有的优惠券
		var allCoupons []model.Coupon
		var err error
		if allCoupons, err = redis.GetCoupons(claims.Username); err != nil {
			log.Println("Server error.")
			ctx.JSON(http.StatusInternalServerError, gin.H{ErrMsgKey: "Server error."})
			return
		}

		coupons := getValidCouponSlice(allCoupons, page)

		if queryUser.IsSeller() {
			sellerCoupons := model.ParseSellerResCoupons(coupons)
			statusCode := getDataStatusCode(len(sellerCoupons))
			ctx.JSON(statusCode, gin.H{ErrMsgKey: "", DataKey: sellerCoupons})
			return
		} else if queryUser.IsCustomer() {
			customerCoupons := model.ParseCustomerResCoupons(coupons)
			statusCode := getDataStatusCode(len(customerCoupons))
			ctx.JSON(statusCode, gin.H{ErrMsgKey: "", DataKey: customerCoupons})
			return
		}
	} else {
		// 查询名与用户名不同
		if queryUser.IsCustomer() {
			// 不可查询其它顾客的优惠券
			log.Println("Cannot check other customer.")
			ctx.JSON(http.StatusUnauthorized, gin.H{ErrMsgKey: "Cannot check other customer.", DataKey: []model.Coupon{}})
			return
		} else if queryUser.IsSeller() {
			// 可以查询其它商家拥有的优惠券
			var allCoupons []model.Coupon
			var err error
			if allCoupons, err = redis.GetCoupons(queryUserName); err != nil {
				log.Println("Error when getting seller's coupons.")
				ctx.JSON(http.StatusInternalServerError, gin.H{ErrMsgKey: "Error when getting seller's coupons.", DataKey: allCoupons})
				return
			}
			coupons := getValidCouponSlice(allCoupons, page)

			sellerCoupons := model.ParseSellerResCoupons(coupons)
			statusCode := getDataStatusCode(len(sellerCoupons))
			ctx.JSON(statusCode, gin.H{ErrMsgKey: "", DataKey: sellerCoupons})
			return
		}
	}

}

// AddCoupon 商家添加优惠券
func AddCoupon(ctx *gin.Context) {
	// 登陆检查token
	claims := ctx.MustGet("claims").(*myjwt.CustomClaims)
	if claims == nil {
		log.Println("Not Authorized.")
		ctx.JSON(http.StatusUnauthorized, gin.H{ErrMsgKey: "Not Authorized."})
		return
	}

	if claims.Kind == "customer" { //!user.IsSeller() {
		log.Println("Only sellers can create coupons.")
		ctx.JSON(http.StatusUnauthorized, gin.H{ErrMsgKey: "Only sellers can create coupons."})
		return
	}

	// 检查参数
	paramUserName := ctx.Param("username") // 注意: 该参数是网址路径参数
	var postCoupon model.ReqCoupon
	if err := ctx.BindJSON(&postCoupon); err != nil {
		log.Println("Only receive JSON format.")
		ctx.JSON(http.StatusBadRequest, gin.H{ErrMsgKey: "Only receive JSON format."})
		return
	}
	couponName := postCoupon.Name
	formAmount := postCoupon.Amount
	description := postCoupon.Description
	formStock := postCoupon.Stock
	if claims.Username != paramUserName {
		log.Println("Cannot create coupons for other users.")
		ctx.JSON(http.StatusUnauthorized, gin.H{ErrMsgKey: "Cannot create coupons for other users."})
		return
	}
	amount := formAmount
	stock := formStock

	// 优惠券描述可以为空的，不需要检查长度
	//if len(couponName) == 0 || len(description) == 0 {
	//	ctx.JSON(http.StatusBadRequest, gin.H{ErrMsgKey: "Coupon name or description should not be empty."})
	//	return
	//}

	// 在数据库添加优惠券
	coupon := model.Coupon{
		Username:    claims.Username,
		CouponName:  couponName,
		Amount:      amount,
		Left:        amount,
		Stock:       stock,
		Description: description,
	}
	var err error
	err = dao.Db.Create(&coupon).Error
	if err != nil {
		log.Println("Create failed. Maybe (username,coupon name) duplicates")
		ctx.JSON(http.StatusBadRequest, gin.H{ErrMsgKey: "Create failed. Maybe (username,coupon name) duplicates"})
		return
	}

	// 在Redis添加优惠券
	if err = redis.CacheCouponAndHasCoupon(coupon); err != nil {
		log.Println("Create Cache failed. ", err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{ErrMsgKey: "Create Cache failed. " + err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{ErrMsgKey: ""})
	return

}

// RegisterUser 用户注册
func RegisterUser(ctx *gin.Context) {
	var postUser model.RegisterUser
	if err := ctx.BindJSON(&postUser); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{ErrMsgKey: "Only receive JSON format."})
		return
	}
	// 查看参数长度、是否为空、格式
	if len(postUser.Username) < model.MinUserNameLen {
		log.Println("User name too short.")
		ctx.JSON(http.StatusBadRequest, gin.H{ErrMsgKey: "User name too short."})
		return
	} else if len(postUser.Password) < model.MinPasswordLen {
		log.Println("Password too short.")
		ctx.JSON(http.StatusBadRequest, gin.H{ErrMsgKey: "Password too short."})
		return
	} else if postUser.Kind == "" {
		log.Println("Empty field of kind.")
		ctx.JSON(http.StatusBadRequest, gin.H{ErrMsgKey: "Empty field of kind."})
		return
	} else if !model.IsValidKind(postUser.Kind) {
		log.Println("Unexpected value of kind, ", postUser.Kind)
		ctx.JSON(http.StatusBadRequest, gin.H{ErrMsgKey: "Unexpected value of kind, " + postUser.Kind})
		return
	}

	// 插入用户
	user := model.User{Username: postUser.Username, Kind: postUser.Kind, Password: model.GetMD5(postUser.Password)} //密码用MD5加密再存储
	err := dao.Db.Create(&user).Error
	if err != nil {
		log.Println("Insert user failed. Maybe user name duplicates.")
		ctx.JSON(http.StatusBadRequest, gin.H{ErrMsgKey: "Insert user failed. Maybe user name duplicates."})
		return
	} else {
		ctx.JSON(http.StatusCreated, gin.H{ErrMsgKey: ""})
		return
	}
}
