package dao

import (
	"mall-go/app/marketing/internal/config"
)

func init() {
	config, err := config.GetAppConfig()
	if err != nil {
		panic("failed to load data config: " + err.Error())
	}

	initMysql(config)
}

func Close() {
	err := Db.Close()
	if err != nil {
		print("Error on closing dbService client.")
	}
}
