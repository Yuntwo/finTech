package model

// Lottery 数据库实体
type Lottery struct {
	Id          int64  `gorm:"primary_key;auto_increment"`
	Username    string `gorm:"type:varchar(20); not null"` // 用户名
	LotteryName string `gorm:"type:varchar(60); not null"` // 优惠券名称
	Amount      int64  // 最大优惠券数
	Left        int64  // 剩余优惠券数
	Stock       int64  // 面额
	Description string `gorm:"type:varchar(60)"` // 优惠券描述信息
}

type ReqLottery struct {
	Name        string
	Amount      int64
	Description string
	Stock       int64
}

type ResLottery struct {
	Name        string `json:"name"`
	Stock       int64  `json:"stock"`
	Description string `json:"description"`
}

// SellerResLottery 商家查询优惠券时，返回的数据结构
type SellerResLottery struct {
	ResLottery
	Amount int64 `json:"amount"`
	Left   int64 `json:"left"`
}

// CustomerResLottery 顾客查询优惠券时，返回的数据结构
type CustomerResLottery struct {
	ResLottery
}

func ParseSellerResLotterys(lotterys []Lottery) []SellerResLottery {
	var sellerLotterys []SellerResLottery
	for _, lottery := range lotterys {
		sellerLotterys = append(sellerLotterys,
			SellerResLottery{ResLottery{lottery.LotteryName, lottery.Stock, lottery.Description},
				lottery.Amount, lottery.Left})
	}
	return sellerLotterys
}

func ParseCustomerResLotterys(lotterys []Lottery) []CustomerResLottery {
	var sellerLotterys []CustomerResLottery
	for _, lottery := range lotterys {
		sellerLotterys = append(sellerLotterys,
			CustomerResLottery{ResLottery{lottery.LotteryName, lottery.Stock, lottery.Description}})
	}
	return sellerLotterys
}
