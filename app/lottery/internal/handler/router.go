package handler

import (
	"encoding/gob"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
	"mall-go/app/lottery/internal/config"
	"mall-go/app/lottery/internal/model"
	redis2 "mall-go/app/lottery/internal/redis"
	"mall-go/app/lottery/internal/service"
	"mall-go/common/middleware"
)

// 相当于controller层，只不过这里没有SpringBoot这种框架通过注解的形式自动处理，而是通过引擎注册的方式

// SessionHeaderKey Visible for test
const SessionHeaderKey = "Authorization"

func SecKillEngine() *gin.Engine {
	router := gin.New()

	// 设置session为Redis存储（但是后来没有用到session，而是用jwt来做用户授权）
	config, err := config.GetAppConfig()
	if err != nil {
		panic("failed to load redisService config" + err.Error())
	}
	store, _ := redis.NewStore(config.App.Redis.MaxIdle, config.App.Redis.Network,
		config.App.Redis.Address, config.App.Redis.Password, []byte("secKill"))
	router.Use(sessions.Sessions(SessionHeaderKey, store))
	gob.Register(&model.User{})

	// 设置路由（路由只需要严格按照接口文档来写就ok了）
	userRouter := router.Group("/service/users")
	userRouter.POST("", service.RegisterUser) //注册
	// Use方法就是添加中间件的，在实际路由方法处理之前，由Gin框架调用，也算SpringBoot的AOP思想
	// Q:多个中间件的执行顺序是怎样的？如何保证链式调用？
	userRouter.Use(middleware.JWTAuth()) //这些请求都需要通过jwt做用户授权
	{
		userRouter.PATCH("/:username/lotterys/:name", service.FetchLottery)
		userRouter.GET("/:username/lotterys", service.GetLotterys)
		userRouter.POST("/:username/lotterys", service.AddLottery)
	}

	authRouter := router.Group("/service/auth") //登录和注销
	{
		authRouter.POST("", service.LoginAuth)
		authRouter.POST("/logout", service.Logout)
	}

	testRouter := router.Group("/test")
	{
		testRouter.GET("/", service.Welcome)
		testRouter.GET("/flush", func(context *gin.Context) {
			if _, err := redis2.FlushAll(); err != nil {
				println("Error when flushAll. " + err.Error())
			} else {
				println("Flushall succeed.")
			}
		})
	}

	// 启动秒杀功能的消费者（用来异步更新数据库）
	service.RunSecKillConsumer()

	return router
}
