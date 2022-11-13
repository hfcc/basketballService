package dao

import (
	"basketballService/db"
	"basketballService/db/model"
)
func InsertUser(userInfo model.UserModel) {
	cli := db.Get()
	cli.Table("user").Create(&userInfo)
}
func SelectUserByOpenId(openId string) (model.UserModel, error) {
	var userInfo model.UserModel
	cli := db.Get()
	var result = cli.Table("user").Where("openId = ?", openId).First(&userInfo)
	return userInfo, result.Error
}