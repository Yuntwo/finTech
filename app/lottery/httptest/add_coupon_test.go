package httptest

import (
	"SecKill/internal/data"
	"testing"
)

// 接连测试多个函数
func TestAddLotteryWrongCases(t *testing.T) {
	_, e := startServer(t)
	defer data.Close()

	registerDemoUsers(e)

	testAddLotteryWrongFormat(e)
	testAddLotteryWrongUser(e)
	testAddLotteryNotLogIn(e)
	testAddLotteryDuplicate(e)
}
