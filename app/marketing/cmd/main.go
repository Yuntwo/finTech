package main

import (
	"flag"
	"fmt"
	"mall-go/app/marketing/internal/dao"
	"mall-go/app/marketing/internal/handler"
	"mall-go/app/marketing/internal/redis"
	"net/http"
	_ "net/http/pprof" // 加载这个包自动初始化执行它的init函数，启用性能分析路由，但不会把包中的内容暴露给当前文件；可以不使用这个包

	"mall-go/app/marketing/internal/config"

	"github.com/mix-plus/core/conf"
)

const port = 20080

var configFile = flag.String("f", "etc/config.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config

	if err := conf.MustLoad(*configFile, &c); err != nil {
		panic(err)
	}

	app, err := initApp(&c)
	if err != nil {
		panic(err)
	}

	app.Run()

	// Separate HTTP server from Gin
	router := handler.SecKillEngine() //路由跳转都写在这里
	defer dao.Close()
	defer redis.Close()

	go func() { //可视化性能测试
		fmt.Println("pprof start...")
		fmt.Println(http.ListenAndServe(":9876", nil))
	}()

	// 这里相当于if条件句可以先声明变量再判断，揉了一下
	if err := router.Run(fmt.Sprintf(":%d", port)); err != nil {
		println("Error when running server. " + err.Error())
	}
}
