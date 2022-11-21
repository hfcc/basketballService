package service

import (
	"basketballService/db/dao"
	"basketballService/db/model"
	"time"
)

func NewUser(openId string) (model.UserModel, error) {
	userInfo, err := dao.SelectUserByOpenId(openId)
	if err == nil && userInfo.OpenId != "" {
		return userInfo, err
	}
	userInfo.OpenId = openId
	userInfo.CreatedAt = time.Now().UTC()
	userInfo.UpdatedAt = time.Now().UTC()
	dao.InsertUser(userInfo)
	return userInfo, nil
}
func GetUserByOpenId(openId string) (model.UserModel, error) {
	userInfo, err := dao.SelectUserByOpenId(openId)
	return userInfo, err
}