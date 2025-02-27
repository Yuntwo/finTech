package httptest

import (
	"SecKill/internal/service"
	"github.com/gavv/httpexpect"
	"net/http"
)

var fetchLotteryPath = "/service/users/{username}/lotterys/{name}"

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
