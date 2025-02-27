package httptest

import (
	"SecKill/internal/service"
	"github.com/gavv/httpexpect"
	"net/http"
	"strconv"
)

/*
该文件下依赖于注册过的demo用户，需要先调用registerDemoUsers
该文件定义了添加优惠券的各种函数
*/

/* 定义了添加优惠券的表格，函数等等 */
const addLotteryPath = "/service/users/{username}/lotterys"

type AddLotteryForm struct {
	Name        string `form:"name"`
	Amount      string `form:"amount"` // 应当int
	Description string `form:"description"`
	Stock       string `form:"stock"` // 应当int
}

// 定义了demo优惠券
var demoLotteryName = "my_lottery"
var demoAmount = 10
var demoStock = 50
var demoAddLotteryForm AddLotteryForm = AddLotteryForm{
	Name:        demoLotteryName,
	Amount:      strconv.Itoa(demoAmount),
	Stock:       strconv.Itoa(demoStock),
	Description: "kiana: this is my good lottery",
}

// 测试添加优惠券时的表格格式
func testAddLotteryWrongFormat(e *httpexpect.Expect) {
	// 登录商家
	logout(e)
	demoSellerLogin(e)

	// amount值不是数字
	amountNotNumberForm := demoAddLotteryForm
	amountNotNumberForm.Amount = "blah-blah"
	e.POST(addLotteryPath, demoSellerName).
		WithForm(amountNotNumberForm).
		Expect().
		Status(http.StatusBadRequest).JSON().Object().
		ValueEqual(service.ErrMsgKey, "Amount field wrong format.")

	// stock值不是数字
	stockNotNumberForm := demoAddLotteryForm
	stockNotNumberForm.Stock = "blah-blah"
	e.POST(addLotteryPath, demoSellerName).
		WithForm(stockNotNumberForm).
		Expect().
		Status(http.StatusBadRequest).JSON().Object().
		ValueEqual(service.ErrMsgKey, "Stock field wrong format.")

	// 优惠券名为空
	emptyLotteryNameForm := demoAddLotteryForm
	emptyLotteryNameForm.Name = ""
	e.POST(addLotteryPath, demoSellerName).
		WithForm(emptyLotteryNameForm).
		Expect().
		Status(http.StatusBadRequest).JSON().Object().
		ValueEqual(service.ErrMsgKey, "Lottery name or description should not be empty.")

	// 优惠券描述为空
	emptyDescriptionForm := demoAddLotteryForm
	emptyDescriptionForm.Description = ""
	e.POST(addLotteryPath, demoSellerName).
		WithForm(emptyDescriptionForm).
		Expect().
		Status(http.StatusBadRequest).JSON().Object().
		ValueEqual(service.ErrMsgKey, "Lottery name or description should not be empty.")
}

// 测试非商家添加优惠券或为其它用户添加优惠券
func testAddLotteryWrongUser(e *httpexpect.Expect) {
	// 登录顾客
	demoCustomerLogin(e)
	// 顾客不可添加优惠券
	e.POST(addLotteryPath, demoCustomerName).
		WithForm(demoAddLotteryForm).
		Expect().
		Status(http.StatusUnauthorized).JSON().Object().
		ValueEqual(service.ErrMsgKey, "Only sellers can create lotterys.")

	// 登录商家
	demoSellerLogin(e)
	// 不可为其它用户添加优惠券
	e.POST(addLotteryPath, demoCustomerName).
		WithForm(demoAddLotteryForm).
		Expect().
		Status(http.StatusUnauthorized).JSON().Object().
		ValueEqual(service.ErrMsgKey, "Cannot create lotterys for other users.")
}

// 测试未登录添加优惠券
func testAddLotteryNotLogIn(e *httpexpect.Expect) {
	logout(e)

	e.POST(addLotteryPath, demoSellerName).
		WithForm(demoAddLotteryForm).
		Expect().
		Status(http.StatusUnauthorized).JSON().Object().
		ValueEqual(service.ErrMsgKey, "Not Logged in.")
}

func testAddLotteryDuplicate(e *httpexpect.Expect) {
	demoSellerLogin(e)

	e.POST(addLotteryPath, demoSellerName).
		WithForm(demoAddLotteryForm).
		Expect().
		Status(http.StatusCreated).JSON().Object().
		ValueEqual(service.ErrMsgKey, "")

	// 添加重复优惠券失败
	e.POST(addLotteryPath, demoSellerName).
		WithForm(demoAddLotteryForm).
		Expect().
		Status(http.StatusBadRequest).JSON().Object().
		ValueEqual(service.ErrMsgKey, "Create failed. Maybe (username,lottery name) duplicates")
}

// 添加demo优惠券(事先不得添加过)
func demoAddLottery(e *httpexpect.Expect) {
	demoSellerLogin(e)

	e.POST(addLotteryPath, demoSellerName).
		WithForm(demoAddLotteryForm).
		Expect().
		Status(http.StatusCreated).JSON().Object().
		ValueEqual(service.ErrMsgKey, "")
}
