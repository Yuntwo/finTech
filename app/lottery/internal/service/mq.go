package service

import (
	"log"
	"mall-go/app/lottery/internal/dao"
	"mall-go/app/lottery/internal/model"
)

type secKillMessage struct {
	username string
	lottery  model.Lottery
}

const maxMessageNum = 20000

var SecKillChannel = make(chan secKillMessage, maxMessageNum) //有缓存的channel

func secKillConsumer() { //从channel中读取信息，更新数据库（就是当redis中优惠券数量减一时，就让数据库也减一）
	for {
		message := <-SecKillChannel
		log.Println("Got one message: " + message.username)

		username := message.username               //抢购成功的用户的用户名
		sellerName := message.lottery.Username     //优惠券的商家名
		lotteryName := message.lottery.LotteryName //优惠券名

		var err error
		err = dao.UserHasLottery(username, message.lottery) //用户优惠券数+1
		if err != nil {
			println("Error when inserting user's lottery. " + err.Error())
		}
		err = dao.DecreaseOneLotteryLeft(sellerName, lotteryName) //优惠券库存自减1
		if err != nil {
			println("Error when decreasing lottery left. " + err.Error())
		}
	}

}

var isConsumerRun = false

func RunSecKillConsumer() {
	// Only Run one consumer.
	if !isConsumerRun {
		go secKillConsumer() //开启一个消费者goroutine，作用是接收redis的改动信息，更新数据库
		isConsumerRun = true
	}
}
