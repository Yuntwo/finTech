package dao

import (
	"fmt"
	"mall-go/app/lottery/internal/model"
)

func GetAllLotteries() ([]model.Lottery, error) {
	var lotteries []model.Lottery
	result := Db.Find(&lotteries)
	return lotteries, result.Error
}

// UserHasLottery 下面原本应该使用gorm封装好的函数来操作数据库的，但是由于当时出了bug，没时间处理，所以直接写数据库命令来操作。
// 插入用户拥有优惠券的数据，这里和商家的优惠券信息共用一张表
// 这个表设计得不行，记录表和优惠券信息表没有拆分，用户和商家的数据也是放在一起的
// 这里的stock是指拥有的优惠券数，left是指剩余的优惠券数；用户默认stock只能是1，left取决于用了没有
func UserHasLottery(userName string, lottery model.Lottery) error {
	return Db.Exec(fmt.Sprintf("INSERT IGNORE INTO lottery "+
		"(`username`,`lottery_name`,`amount`,`left`,`stock`,`description`) "+
		`values('%s', '%s', %d, %d, %f, '%s')`,
		userName, lottery.LotteryName, 1, 1, lottery.Stock, lottery.Description)).Error
}

// DecreaseOneLotteryLeft 优惠券库存自减1
func DecreaseOneLotteryLeft(sellerName string, lotteryName string) error {
	return Db.Exec(fmt.Sprintf("UPDATE lottery c SET c.left=c.left-1 WHERE "+
		"c.username='%s' AND c.lottery_name='%s' AND c.left>0", sellerName, lotteryName)).Error
}
