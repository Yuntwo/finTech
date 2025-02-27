package model

import "time"

// Promotion 营销投放：管理营销资源投放、业务接入和数据建模
// 投放业务接入：如decision、商详曝光、支付咨询
// 投放决策引擎：如投放位、投放计划、资源位
// 数据模型：如投放位、素材
type Promotion struct {
	ID         int64
	Name       string
	Type       string
	Placement  string
	ResourceID int64
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

// Activity 营销活动：管理活动创建、规则引擎和奖品
// 业务接入：如收银台预咨询、直播间预咨询等
// 活动决策引擎：如规则、活动、奖品组等
// 数据模型：如规则、活动
type Activity struct {
	ID         int64
	Name       string
	Rules      string
	PrizeGroup string
	StartTime  time.Time
	EndTime    time.Time
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

// Coupon 营销资产：管理优惠券、立减等营销资产的发放、核销和管理
// 业务接入：如发放支付券、优惠核销、前置券透出等
// 资产决策引擎：如支付券、立减、累计券等
// 数据模型：如支付券、立减等
type Coupon struct {
	ID         int64
	Code       string
	Discount   float64
	ExpiryDate time.Time
	Status     string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
