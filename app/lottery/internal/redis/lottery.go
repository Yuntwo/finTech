package redis

import (
	"fmt"
	"log"
	"mall-go/app/lottery/internal/dao"
	"mall-go/app/lottery/internal/model"
	"strconv"
)

// 获取"用户持有优惠券"的key
func getHasLotterysKeyByName(userName string) string {
	return fmt.Sprintf("%s-has", userName)
}

// 获取"优惠券"的key
func getLotteryKeyByLottery(lottery model.Lottery) string {
	return getLotteryKeyByName(lottery.LotteryName)
}
func getLotteryKeyByName(lotteryName string) string {
	return fmt.Sprintf("%s-info", lotteryName)
}

// CacheHasLottery 缓存用户拥有优惠券/商家创建优惠券的信息
func CacheHasLottery(lottery model.Lottery) (int64, error) {
	key := getHasLotterysKeyByName(lottery.Username) //得到的key其实就是 lottery.Username-has
	val, err := SetAdd(key, lottery.LotteryName)
	return val, err
}

// CacheLottery 缓存优惠券的完整信息
func CacheLottery(lottery model.Lottery) (int64, error) {
	key := getLotteryKeyByLottery(lottery)
	fields := map[string]interface{}{
		"id":          lottery.Id,
		"username":    lottery.Username,
		"lotteryName": lottery.LotteryName,
		"amount":      lottery.Amount,
		"left":        lottery.Left,
		"stock":       lottery.Stock,
		"description": lottery.Description,
	}
	val, err := SetMapForever(key, fields)
	return val, err
}

// CacheLotteryAndHasLottery 缓存优惠券
func CacheLotteryAndHasLottery(lottery model.Lottery) error {
	if _, err := CacheHasLottery(lottery); err != nil {
		return err
	}

	// user = 根据优惠券的username查user
	if user, err := dao.GetUser(lottery.Username); err != nil {
		log.Println("Database service error: ", err)
		return err
	} else {
		if user.IsSeller() {
			_, err = CacheLottery(lottery)
		}
		return err
	}
}

// GetLottery 从缓存获取优惠券
func GetLottery(lotteryName string) model.Lottery {
	key := getLotteryKeyByName(lotteryName)
	values, err := GetMap(key, "id", "username", "lotteryName", "amount", "left", "stock", "description")
	if err != nil {
		println("Error on getting lottery. " + err.Error())
	}
	// log.Println(values) TODO
	// values[0]类型是nil，说明key是不存在的？
	id, err := strconv.ParseInt(values[0].(string), 10, 64)
	if err != nil {
		println("Wrong type of id. " + err.Error())
	}
	amount, err := strconv.ParseInt(values[3].(string), 10, 64)
	if err != nil {
		println("Wrong type of amount. " + err.Error())
	}
	left, err := strconv.ParseInt(values[4].(string), 10, 64)
	if err != nil {
		println("Wrong type of left. " + err.Error())
	}
	stock, err := strconv.ParseInt(values[5].(string), 10, 64)
	if err != nil {
		println("Wrong type of stock. " + err.Error())
	}
	return model.Lottery{
		Id:          id,
		Username:    values[1].(string),
		LotteryName: values[2].(string),
		Amount:      amount,
		Left:        left,
		Stock:       stock,
		Description: values[6].(string),
	}

}

// GetLotterys 从缓存获取某个用户的所有优惠券
func GetLotterys(userName string) ([]model.Lottery, error) {
	var lotterys []model.Lottery
	hasLotterysKey := getHasLotterysKeyByName(userName)
	lotteryNames, err := GetSetMembers(hasLotterysKey)
	if err != nil {
		println("Error when getting lottery members. " + err.Error())
		return nil, err
	}
	// TODO: 使用数组, 不使用slice append
	for _, lotteryName := range lotteryNames {
		lottery := GetLottery(lotteryName)
		lotterys = append(lotterys, lottery)
	}
	return lotterys, nil
}
