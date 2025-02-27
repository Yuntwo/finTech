package redis

import (
	"mall-go/app/marketing/internal/config"
	"mall-go/app/marketing/internal/dao"
)

const secKillScript = `
    -- Check if User has lottery
    -- KEYS[1]: hasLotteryKey "{username}-has"
    -- KEYS[2]: lotteryName   "{lotteryName}"
    -- KEYS[3]: lotteryKey    "{lotteryName}-info"
    -- 返回值有-1, -2, -3, 都代表抢购失败
    -- 返回值为1代表抢购成功

    -- **第一大步: Check if lottery exists and has left** --
	local lotteryLeft = redis.call("hget", KEYS[3], "left"); -- left是优惠券信息Hash中的一个字段
    -- 上述脚本执行结构可能返回false或者成功实际结果
    -- 返回false有以下3种情况:
    -- 1. KEYS[3]这个Key不存在
    -- 2. KEYS[3]这个Key存在但不是哈希类型，因为hget只能用于哈希类型
    -- 3. KEYS[3]这个Key存在且是哈希类型，但是没有left字段，因为hget只能返回指定字段的值，否则返回false而不是空值或者其它类型。
    -- 后两种情况从业务逻辑的正确性保证来说应该不存在
	if (lotteryLeft == false)
	then
		return -2;  -- No such lottery
	end
    -- 执行成功返回实际KEYS[3]中left字段的值
	if (tonumber(lotteryLeft) == 0)  -- left的值为0代表没有库存了
    then
		return -3;  --  No Lottery Left
	end

    -- **第二大步: Check if the user has got the lottery** --
    -- Set命令，SISMEMBER key member，判断member元素是否是集合key的成员
    -- redis命令本身大小写不敏感，但是命令参数大小写敏感。这里SISMEMBER也可以写作sismember等
	local userHasLottery = redis.call("SISMEMBER", KEYS[1], KEYS[2]);
	if (userHasLottery == 1)
	then
		return -1;
	end

    -- **第三大步: 优惠券存在且有库存、用户也没有获得过这个优惠券时，User gets the lottery，涉及到两个key --
	redis.call("hset", KEYS[3], "left", lotteryLeft - 1);
	redis.call("SADD", KEYS[1], KEYS[2]);
	return 1;
`

var secKillSHA string // SHA expression of secKillScript

// 将数据加载到缓存预热，防止缓存穿透
// 预热加载了商品库存key
func preHeatKeys() {
	lotterys, err := dao.GetAllLotteries()
	if err != nil {
		panic("Error when getting all lotterys." + err.Error())
	}

	for _, lottery := range lotterys {
		err := CacheLotteryAndHasLottery(lottery)
		if err != nil {
			panic("Error while setting redis keys of lotterys. " + err.Error())
		}
	}
	println("---Set redis keys of lotterys success.---")
}

func init() {

	config, err := config.GetAppConfig()
	if err != nil {
		panic("failed to load data config: " + err.Error())
	}

	initRedisConnection(config)

	// 让redis加载秒杀的lua脚本

	// 将Lua脚本secKillScript预先加载到Redis，返回其SHA1值(有则直接返回没有则计算)。
	// 通过SHA1值查找对应缓存的脚本，避免每次执行时都需要重新传输脚本，没有的话会报错找不到脚本。
	secKillSHA = PrepareScript(secKillScript)

	// 预热
	preHeatKeys()
}

func Close() {
	err := client.Close()
	if err != nil {
		print("Error on closing redisService client.")
	}
}
