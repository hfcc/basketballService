package dao

import (
	"basketballService/db"
	"basketballService/db/model"
)
func InsertFavor(favor model.FavorModel) (model.FavorModel, error) {
	cli := db.Get()
	err := cli.Table("favor").Create(&favor).Error
	return favor, err
}
func SelectFavorByUserId(userId int32) ([]model.FavorModel, error) {
	cli := db.Get()
	var favors []model.FavorModel
	var result = cli.Table("favor").Where("userId = ?", userId).Find(&favors)
	return favors, result.Error
}