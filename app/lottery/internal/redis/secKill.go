package redis

import (
	"fmt"
	"github.com/prometheus/common/promslog"
)

// 下面是一大堆自定义的Error
type redisEvalError struct {
}

func (e redisEvalError) Error() string {
	return "Error when executing redisService eval."
}

type userHasLotteryError struct {
	userName    string
	lotteryName string
}

func (e userHasLotteryError) Error() string {
	return fmt.Sprintf("User %s has had lottery %s.", e.userName, e.lotteryName)
}

type noSuchLotteryError struct {
	userName    string
	lotteryName string
}

func (e noSuchLotteryError) Error() string {
	return fmt.Sprintf("Lottery %s created by %s doesn't exist.", e.lotteryName, e.userName)
}

type noLotteryLeftError struct {
	userName    string
	lotteryName string
}

func (e noLotteryLeftError) Error() string {
	return fmt.Sprintf("No Lottery %s created by %s left.", e.lotteryName, e.userName)
}

type LotteryLeftResError struct {
	lotteryLeftRes interface{}
}

func (e LotteryLeftResError) Error() string {
	switch e.lotteryLeftRes.(type) {
	case int:
		return fmt.Sprintf("Unexpected lotteryLeftRes Num: %v.", e.lotteryLeftRes)
	default:
		return fmt.Sprintf("lotteryLeftRes : %v with wrong type.", e.lotteryLeftRes)
	}
}

func IsRedisEvalError(err error) bool {
	switch err.(type) {
	case redisEvalError:
		return true
	default:
		return false
	}
}

// CacheAtomicSecKill 尝试在redis进行原子性的秒杀操作
func CacheAtomicSecKill(userName string, sellerName string, lotteryName string) (int64, error) {
	// 根据sha，执行预先加载的秒杀lua脚本
	userHasLotterysKey := getHasLotterysKeyByName(userName)
	lotteryKey := getLotteryKeyByName(lotteryName)
	res, err := EvalSHA(secKillSHA, []string{userHasLotterysKey, lotteryName, lotteryKey})
	if err != nil {
		return -1, redisEvalError{}
	}

	// 该lua脚本应当返回int值
	lotteryLeftRes, ok := res.(int64)
	if !ok {
		return -1, LotteryLeftResError{res}
	}

	// 此处的-1, -2, -3 和 1的判断依据, 与secKillSHA变量lua脚本的返回值保持一致
	// 请看secKillSHA
	switch {
	case lotteryLeftRes == -1:
		return lotteryLeftRes, userHasLotteryError{userName, lotteryName}
	case lotteryLeftRes == -2:
		return lotteryLeftRes, noSuchLotteryError{sellerName, lotteryName}
	case lotteryLeftRes == -3:
		return lotteryLeftRes, noLotteryLeftError{sellerName, lotteryName}
	case lotteryLeftRes == 1: // 抢券成功，注意这里lotteryLeftRes返回的不是剩余的优惠券数量，而只是状态码
		return lotteryLeftRes, nil
	default:
		{
			logger := promslog.New(&promslog.Config{})
			logger.Error("Unexpected return value.")
			return -1, LotteryLeftResError{lotteryLeftRes}
		}

	}
}
