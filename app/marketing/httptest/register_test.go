package httptest

import (
	"github.com/gavv/httpexpect"
	"mall-go/app/marketing/internal/dao"
	"mall-go/app/marketing/internal/handler"
	"mall-go/app/marketing/internal/model"
	"mall-go/app/marketing/internal/redis"
	"mall-go/app/marketing/internal/service"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

const registerUserPath = "/service/users/"

func startServer(t *testing.T) (*httptest.Server, *httpexpect.Expect) {
	// 启动服务器
	server := httptest.NewServer(handler.SecKillEngine())

	// 通过server创建测试引擎
	return server, httpexpect.WithConfig(httpexpect.Config{
		BaseURL:  server.URL,
		Reporter: httpexpect.NewAssertReporter(t),

		// use http.Client with a cookie jar and timeout
		Client: &http.Client{
			Jar:     httpexpect.NewJar(),
			Timeout: time.Second * 30,
		},
		// use verbose logging
		Printers: []httpexpect.Printer{
			httpexpect.NewCurlPrinter(t),
			httpexpect.NewDebugPrinter(t, true),
		},
	})
}

func testFormat(e *httpexpect.Expect) {
	// 足够长的密码
	longPassword := ""
	for i := 0; i < 15; i++ {
		longPassword += "p"
	}
	// 足够长的用户名
	longUserName := ""
	for i := 0; i < 15; i++ {
		longUserName += "u"
	}

	// 用户名过短
	tooShortUserName := ""
	for i := 0; i < model.MinUserNameLen-1; i++ {
		tooShortUserName += "t"
	}
	e.POST(registerUserPath).
		WithForm(RegisterForm{tooShortUserName, longPassword, model.NormalSeller}).
		Expect().
		Status(http.StatusBadRequest).JSON().Object().
		ValueEqual(service.ErrMsgKey, "User name too short.")

	// 密码过短
	tooShortPassword := ""
	for i := 0; i < model.MinPasswordLen-1; i++ {
		tooShortPassword += "p"
	}
	e.POST(registerUserPath).
		WithForm(RegisterForm{longUserName, tooShortPassword, model.NormalSeller}).
		Expect().
		Status(http.StatusBadRequest).JSON().Object().
		ValueEqual(service.ErrMsgKey, "Password too short.")

	// 定义错误用户类型表格
	type BadKindRegisterForm struct {
		Username string `form:"username"`
		Password string `form:"password"`
		Kind     string `form:"kind"`
	}
	// 用户类型为空
	e.POST(registerUserPath).
		WithForm(BadKindRegisterForm{longUserName, longPassword, ""}).
		Expect().
		Status(http.StatusBadRequest).JSON().Object().
		ValueEqual(service.ErrMsgKey, "Empty field of kind.")

	// 用户类型不存在
	impossibleKind := "ImpossibleKind"
	e.POST(registerUserPath).
		WithForm(BadKindRegisterForm{longUserName, longPassword, impossibleKind}).
		Expect().
		Status(http.StatusBadRequest).JSON().Object().
		ValueEqual(service.ErrMsgKey, "Unexpected value of kind, "+impossibleKind)

}

func testDuplicateRegisterSeller(e *httpexpect.Expect) {
	// 注册一个用户
	e.POST("/service/users/").
		WithForm(RegisterForm{"kiana", "shen6508", model.NormalSeller}).
		Expect().
		Status(http.StatusOK).JSON().Object().
		ValueEqual(service.ErrMsgKey, "")

	// 重复注册
	e.POST("/service/users/").
		WithForm(RegisterForm{"kiana", "password2", model.NormalSeller}).
		Expect().
		Status(http.StatusBadRequest).JSON().Object().
		ValueEqual(service.ErrMsgKey, "Insert user failed. Maybe user name duplicates.")
}

func TestRegisterCases(t *testing.T) {
	_, e := startServer(t)
	defer dao.Close()
	defer redis.Close()

	// 注册失败
	testFormat(e)

	// 注册商家kiana成功
	testDuplicateRegisterSeller(e)
}
