package redis

import (
	"github.com/go-redis/redis/v7"
	"mall-go/app/lottery/internal/config"
)

var client *redis.Client

// 开启redis连接池
func initRedisConnection(config config.AppConfig) {
	client = redis.NewClient(&redis.Options{
		Addr:     config.App.Redis.Address,
		Password: config.App.Redis.Password, // It's ok if password is "".
		DB:       0,                         // use default DB
	})

	if _, err := FlushAll(); err != nil {
		println("Error when flushAll. " + err.Error())
	}
}

// FlushAll 用于测试
func FlushAll() (string, error) {
	return client.FlushAll().Result()
}

// PrepareScript 确保redis加载lua脚本，若未加载则加载
func PrepareScript(script string) string {
	// sha := sha1.Sum([]byte(script))
	scriptsExists, err := client.ScriptExists(script).Result()
	if err != nil {
		panic("Failed to check if script exists: " + err.Error())
	}
	if !scriptsExists[0] {
		scriptSHA, err := client.ScriptLoad(script).Result()
		if err != nil {
			panic("Failed to load script " + script + " err: " + err.Error())
		}
		return scriptSHA
	}
	print("Script Exists.")
	return ""
}

// EvalSHA 执行lua脚本
func EvalSHA(sha string, args []string) (interface{}, error) {
	val, err := client.EvalSha(sha, args).Result()
	if err != nil {
		print("Error executing evalSHA... " + err.Error())
		return nil, err
	}
	return val, nil
}

// SetForever redis operation SET
func SetForever(key string, value interface{}) (string, error) {
	val, err := client.Set(key, value, 0).Result() // expiration表示无过期时间
	return val, err
}

// SetMapForever redis operation hmset
func SetMapForever(key string, field map[string]interface{}) (string, error) {
	return client.HMSet(key, field).Result()
}

// GetMap redis operation hmget
func GetMap(key string, fields ...string) ([]interface{}, error) {
	return client.HMGet(key, fields...).Result()
}

// SetAdd redis SADD
func SetAdd(key string, field string) (int64, error) {
	return client.SAdd(key, field).Result()
}

// SetIsMember redis SISMEMBER
func SetIsMember(key string, field string) (bool, error) {
	return client.SIsMember(key, field).Result()
}

// GetSetMembers redis SMEMBERS
func GetSetMembers(key string) ([]string, error) {
	return client.SMembers(key).Result()
}
