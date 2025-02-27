package httptest

import (
	"SecKill/internal/model"
	"SecKill/internal/service"
	"fmt"
	"github.com/gavv/httpexpect"
	"net/http"
)

/*
该文件下依赖于注册过的demo用户，需要先调用registerDemoUsers
该文件定义了查看优惠券的各种函数
*/

/* 定义了查看优惠券的路径，函数等等 */
const getLotteryPath = "/service/users/{username}/lotterys"
const pageQueryKey = "page"

var customerSchema = fmt.Sprintf(`{
	"type": "object",
	"properties": {
		"%s": {
				"type": "string"
			},
        "%s": {
				"type": "array",
				"items": {
					"type":        "object",
					"name":        "string",
					"amount":      "integer",
					"left":        "integer",
					"stock":       "integer",
					"description": "string"
				}
			}
	}
}`, service.ErrMsgKey, service.DataKey)

var sellerSchema = fmt.Sprintf(`{
	"type": "object",
	"properties": {
		"%s": {
				"type": "string"
			},
        "%s": {
				"type": "array",
				"items": {
					"type":        "object",
					"name":        "string",
					"stock":       "integer",
					"description": "string"
				}
			}
	}
}`, service.ErrMsgKey, service.DataKey)

// 当返回的数据为空，则状态码应为204，且没有返回内容
func isEmptyBody(e *httpexpect.Expect, username string, page int) {
	e.GET(getLotteryPath, username).
		WithQuery(pageQueryKey, page).
		Expect().
		Status(http.StatusNoContent).Body().Empty()
}

// 当返回的"data"不为空，则状态码应为200
func isNonEmptyLotterys(e *httpexpect.Expect, username string, page int) {
	e.GET(getLotteryPath, username).
		WithQuery(pageQueryKey, page).
		Expect().
		Status(http.StatusOK).JSON().Object().
		ValueEqual(service.ErrMsgKey, "").
		Value("data").Array().Length().Gt(0)
}

func isGetLotteryUnauthorized(e *httpexpect.Expect, username string, page int) {
	jsonObject := e.GET(getLotteryPath, username).
		WithQuery(pageQueryKey, page).
		Expect().
		Status(http.StatusUnauthorized).JSON().Object()
	jsonObject.Value(service.ErrMsgKey).Equal("Cannot check other customer.")
	jsonObject.Value(service.DataKey).Equal([]model.Lottery{})

}

// 验证符合顾客的格式
func isCustomerSchema(e *httpexpect.Expect, username string, page int) {
	e.GET(getLotteryPath, username).WithQuery(pageQueryKey, page).
		Expect().JSON().Schema(customerSchema)

}

// 验证符合商家的格式
func isSellerSchema(e *httpexpect.Expect, username string, page int) {
	e.GET(getLotteryPath, username).WithQuery(pageQueryKey, page).
		Expect().JSON().Schema(sellerSchema)
}

// 验证优惠券的剩余量与预期一致
func isLotteryExpectedLeft(e *httpexpect.Expect, username string, page int, index int, expectedLeft int) {
	e.GET(getLotteryPath, username).WithQuery(pageQueryKey, page).
		Expect().JSON().Object().Value(service.DataKey).Array().
		Element(index).Object().Value("left").Equal(expectedLeft)
}
