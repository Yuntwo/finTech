package dao

import (
	"fmt"
	"mall-go/app/lottery/internal/model"
)

func GetAllLotterys() ([]model.Lottery, error) {
	var lotterys []model.Lottery
	result := Db.Find(&lotterys)
	return lotterys, result.Error
}

// UserHasLottery 下面原本应该使用gorm封装好的函数来操作数据库的，但是由于当时出了bug，没时间处理，所以直接写数据库命令来操作。
// 插入用户拥有优惠券的数据
func UserHasLottery(userName string, lottery model.Lottery) error {
	return Db.Exec(fmt.Sprintf("INSERT IGNORE INTO lotterys "+
		"(`username`,`lottery_name`,`amount`,`left`,`stock`,`description`) "+
		"values('%s', '%s', %d, %d, %f, '%s')",
		userName, lottery.LotteryName, 1, 1, lottery.Stock, lottery.Description)).Error
}

// 优惠券库存自减1
func DecreaseOneLotteryLeft(sellerName string, lotteryName string) error {
	return Db.Exec(fmt.Sprintf("UPDATE lotterys c SET c.left=c.left-1 WHERE "+
		"c.username='%s' AND c.lottery_name='%s' AND c.left>0", sellerName, lotteryName)).Error
}
