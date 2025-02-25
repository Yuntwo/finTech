package httptest

import (
	"SecKill/internal/service"
	"github.com/gavv/httpexpect"
	"net/http"
)

var fetchCouponPath = "/service/users/{username}/coupons/{name}"

func fetchDemoCouponSuccess(e *httpexpect.Expect) {
	e.PATCH(fetchCouponPath, demoSellerName, demoCouponName).
		Expect().
		Status(http.StatusCreated).JSON().Object().
		ValueEqual(service.ErrMsgKey, "")
}

func fetchDemoCouponFail(e *httpexpect.Expect) {
	e.PATCH(fetchCouponPath, demoSellerName, demoCouponName).
		Expect().
		Status(http.StatusNoContent).
		Body().Empty()
}
