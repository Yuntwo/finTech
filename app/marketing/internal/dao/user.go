package dao

import (
	"mall-go/app/marketing/internal/model"
)

func GetUser(userName string) (model.User, error) {
	user := model.User{}
	operation := Db.Where("username = ?", userName).First(&user)
	return user, operation.Error
}
