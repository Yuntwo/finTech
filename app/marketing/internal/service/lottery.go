package service

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"log"
	"mall-go/app/marketing/internal/dao"
	"mall-go/app/marketing/internal/model"
	"mall-go/app/marketing/internal/redis"
	jwt "mall-go/common/middleware" // 算是很实用的命名方法，因为middleware下面没有给auth放在单独的包里面取别名
	"net/http"
	"strconv"
)

const (
	lotteryPageSize int64 = 20
	// ErrMsgKey Visible for testing
	ErrMsgKey = "errMsg"
	DataKey   = "data"
)

// FetchLottery 秒杀优惠券，seller不允许获取优惠券
// TODO 应该可以抽象为切面来处理的，但是go似乎没有切面的概念，中间件似乎不太适合做切面
func FetchLottery(ctx *gin.Context) {
	// 这里是业务逻辑的「授权(Authorization)」，用于判断合法用户是否有权访问当前资源(例如，某个用户是否可以获取优惠券，某个角色是否有权限执行某个操作)
	claims := ctx.MustGet("claims").(*jwt.CustomClaims)
	if claims == nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{ErrMsgKey: "Not Authorized."})
		return
	}

	// 后续直接由gin Web框架根据gin.Context的内容打包成HTTP响应以及发送等

	if claims.Kind == "seller" {
		ctx.JSON(http.StatusUnauthorized, gin.H{ErrMsgKey: "Sellers aren't allowed to get lotterys."})
		return
	}

	// HTTP请求参数
	paramSellerName := ctx.Param("username")
	paramLotteryName := ctx.Param("name")

	// ---用户抢优惠券。后面需要高并发处理---
	// 先在缓存执行原子性的秒杀操作。将原子性地完成"判断能否秒杀-执行秒杀"的步骤
	_, err := redis.CacheAtomicSecKill(claims.Username, paramSellerName, paramLotteryName)
	if err == nil {
		lottery := redis.GetLottery(paramLotteryName)
		// 交给[协程]完成数据库写入操作
		SecKillChannel <- secKillMessage{claims.Username, lottery}
		ctx.JSON(http.StatusCreated, gin.H{ErrMsgKey: ""})
		return
	} else {
		if redis.IsRedisEvalError(err) {
			log.Printf("Server error" + err.Error())
			ctx.JSON(http.StatusInternalServerError, gin.H{ErrMsgKey: err.Error()})
			return
		} else {
			log.Println("Fail to fetch lottery. " + err.Error())
			ctx.JSON(http.StatusNoContent, gin.H{})
			return
		}
	}
}

// GetLotterys 查询优惠券
func GetLotterys(ctx *gin.Context) {
	// 登陆检查token
	claims := ctx.MustGet("claims").(*jwt.CustomClaims)
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

	//log.Printf("Querying lottery with name %s, page %d\n", queryUserName, page)
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
		var allLotterys []model.Lottery
		var err error
		if allLotterys, err = redis.GetLotterys(claims.Username); err != nil {
			log.Println("Server error.")
			ctx.JSON(http.StatusInternalServerError, gin.H{ErrMsgKey: "Server error."})
			return
		}

		lotterys := getValidLotterySlice(allLotterys, page)

		if queryUser.IsSeller() {
			sellerLotterys := model.ParseSellerResLotterys(lotterys)
			statusCode := getDataStatusCode(len(sellerLotterys))
			ctx.JSON(statusCode, gin.H{ErrMsgKey: "", DataKey: sellerLotterys})
			return
		} else if queryUser.IsCustomer() {
			customerLotterys := model.ParseCustomerResLotterys(lotterys)
			statusCode := getDataStatusCode(len(customerLotterys))
			ctx.JSON(statusCode, gin.H{ErrMsgKey: "", DataKey: customerLotterys})
			return
		}
	} else {
		// 查询名与用户名不同
		if queryUser.IsCustomer() {
			// 不可查询其它顾客的优惠券
			log.Println("Cannot check other customer.")
			ctx.JSON(http.StatusUnauthorized, gin.H{ErrMsgKey: "Cannot check other customer.", DataKey: []model.Lottery{}})
			return
		} else if queryUser.IsSeller() {
			// 可以查询其它商家拥有的优惠券
			var allLotterys []model.Lottery
			var err error
			if allLotterys, err = redis.GetLotterys(queryUserName); err != nil {
				log.Println("Error when getting seller's lotterys.")
				ctx.JSON(http.StatusInternalServerError, gin.H{ErrMsgKey: "Error when getting seller's lotterys.", DataKey: allLotterys})
				return
			}
			lotterys := getValidLotterySlice(allLotterys, page)

			sellerLotterys := model.ParseSellerResLotterys(lotterys)
			statusCode := getDataStatusCode(len(sellerLotterys))
			ctx.JSON(statusCode, gin.H{ErrMsgKey: "", DataKey: sellerLotterys})
			return
		}
	}

}

// AddLottery 商家添加优惠券
func AddLottery(ctx *gin.Context) {
	// 登陆检查token
	claims := ctx.MustGet("claims").(*jwt.CustomClaims)
	if claims == nil {
		log.Println("Not Authorized.")
		ctx.JSON(http.StatusUnauthorized, gin.H{ErrMsgKey: "Not Authorized."})
		return
	}

	if claims.Kind == "customer" { //!user.IsSeller() {
		log.Println("Only sellers can create lotterys.")
		ctx.JSON(http.StatusUnauthorized, gin.H{ErrMsgKey: "Only sellers can create lotterys."})
		return
	}

	// 检查参数
	paramUserName := ctx.Param("username") // 注意: 该参数是网址路径参数
	var postLottery model.ReqLottery
	if err := ctx.BindJSON(&postLottery); err != nil {
		log.Println("Only receive JSON format.")
		ctx.JSON(http.StatusBadRequest, gin.H{ErrMsgKey: "Only receive JSON format."})
		return
	}
	lotteryName := postLottery.Name
	formAmount := postLottery.Amount
	description := postLottery.Description
	formStock := postLottery.Stock
	if claims.Username != paramUserName {
		log.Println("Cannot create lotterys for other users.")
		ctx.JSON(http.StatusUnauthorized, gin.H{ErrMsgKey: "Cannot create lotterys for other users."})
		return
	}
	amount := formAmount
	stock := formStock

	// 优惠券描述可以为空的，不需要检查长度
	//if len(lotteryName) == 0 || len(description) == 0 {
	//	ctx.JSON(http.StatusBadRequest, gin.H{ErrMsgKey: "Lottery name or description should not be empty."})
	//	return
	//}

	// 在数据库添加优惠券
	lottery := model.Lottery{
		Username:    claims.Username,
		LotteryName: lotteryName,
		Amount:      amount,
		Left:        amount,
		Stock:       stock,
		Description: description,
	}
	var err error
	err = dao.Db.Create(&lottery).Error
	if err != nil {
		log.Println("Create failed. Maybe (username,lottery name) duplicates")
		ctx.JSON(http.StatusBadRequest, gin.H{ErrMsgKey: "Create failed. Maybe (username,lottery name) duplicates"})
		return
	}

	// 在Redis添加优惠券
	if err = redis.CacheLotteryAndHasLottery(lottery); err != nil {
		log.Println("Create Cache failed. ", err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{ErrMsgKey: "Create Cache failed. " + err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{ErrMsgKey: ""})
	return

}

// RegisterUser 用户注册
// TODO refactor到user包下面
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

// 根据page页数, 取得合理切片范围的lottery
// 需要保证切片索引startIndex, endIndex不越界
func getValidLotterySlice(allLotterys []model.Lottery, page int64) []model.Lottery {
	if len(allLotterys) == 0 {
		return allLotterys
	}
	lotteryLen := int64(len(allLotterys))
	startIndex := page * lotteryPageSize
	endIndex := page*lotteryPageSize + lotteryPageSize
	if startIndex < 0 {
		startIndex = 0
	} else if startIndex > lotteryLen {
		startIndex = lotteryLen
	}
	if endIndex < 1 {
		if lotteryLen < lotteryPageSize {
			endIndex = lotteryLen
		} else {
			endIndex = lotteryPageSize
		}
	} else if endIndex > lotteryLen {
		endIndex = lotteryLen
	}
	return allLotterys[startIndex:endIndex]
}

// 数据长度为空则返回204,否则返回200
func getDataStatusCode(len int) int {
	if len == 0 {
		return http.StatusNoContent
	} else {
		return http.StatusOK
	}
}
