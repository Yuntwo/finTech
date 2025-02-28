package httptest

import (
	"github.com/gavv/httpexpect"
	"mall-go/app/marketing/internal/service"
	"net/http"
)

var fetchLotteryPath = "/service/users/{username}/lotterys/{name}"

// 定义了demo优惠券
// TODO Q:为什么放在其他文件无法访问？
var demoLotteryName = "my_lottery"
var demoAmount = 10
var demoStock = 50

func fetchDemoLotterySuccess(e *httpexpect.Expect) {
	e.PATCH(fetchLotteryPath, demoSellerName, demoLotteryName).
		Expect().
		Status(http.StatusCreated).JSON().Object().
		ValueEqual(service.ErrMsgKey, "")
}

func fetchDemoLotteryFail(e *httpexpect.Expect) {
	e.PATCH(fetchLotteryPath, demoSellerName, demoLotteryName).
		Expect().
		Status(http.StatusNoContent).
		Body().Empty()
}
