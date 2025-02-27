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

type userHasCouponError struct {
	userName   string
	couponName string
}

func (e userHasCouponError) Error() string {
	return fmt.Sprintf("User %s has had coupon %s.", e.userName, e.couponName)
}

type noSuchCouponError struct {
	userName   string
	couponName string
}

func (e noSuchCouponError) Error() string {
	return fmt.Sprintf("Coupon %s created by %s doesn't exist.", e.couponName, e.userName)
}

type noCouponLeftError struct {
	userName   string
	couponName string
}

func (e noCouponLeftError) Error() string {
	return fmt.Sprintf("No Coupon %s created by %s left.", e.couponName, e.userName)
}

type CouponLeftResError struct {
	couponLeftRes interface{}
}

func (e CouponLeftResError) Error() string {
	switch e.couponLeftRes.(type) {
	case int:
		return fmt.Sprintf("Unexpected couponLeftRes Num: %v.", e.couponLeftRes)
	default:
		return fmt.Sprintf("couponLeftRes : %v with wrong type.", e.couponLeftRes)
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
func CacheAtomicSecKill(userName string, sellerName string, couponName string) (int64, error) {
	// 根据sha，执行预先加载的秒杀lua脚本
	userHasCouponsKey := getHasCouponsKeyByName(userName)
	couponKey := getCouponKeyByName(couponName)
	res, err := EvalSHA(secKillSHA, []string{userHasCouponsKey, couponName, couponKey})
	if err != nil {
		return -1, redisEvalError{}
	}

	// 该lua脚本应当返回int值
	couponLeftRes, ok := res.(int64)
	if !ok {
		return -1, CouponLeftResError{res}
	}

	// 此处的-1, -2, -3 和 1的判断依据, 与secKillSHA变量lua脚本的返回值保持一致
	// 请看secKillSHA
	switch {
	case couponLeftRes == -1:
		return couponLeftRes, userHasCouponError{userName, couponName}
	case couponLeftRes == -2:
		return couponLeftRes, noSuchCouponError{sellerName, couponName}
	case couponLeftRes == -3:
		return couponLeftRes, noCouponLeftError{sellerName, couponName}
	case couponLeftRes == 1: // 抢券成功，注意这里couponLeftRes返回的不是剩余的优惠券数量，而只是状态码
		return couponLeftRes, nil
	default:
		{
			logger := promslog.New(&promslog.Config{})
			logger.Error("Unexpected return value.")
			return -1, CouponLeftResError{couponLeftRes}
		}

	}
}
