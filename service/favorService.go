package service

import (
	"basketballService/db/dao"
	"basketballService/db/model"
	"time"
)
func GetFavors(userId int32) ([]model.FavorModel, error) {
	favors, err := dao.SelectFavorByUserId(userId)
	if err != nil {
		return nil, err
	}
	return favors, nil
}
func InsertFavor(userId int32, gameId string) (model.FavorModel, error) {
	var favor model.FavorModel
	favor.UserId = userId
	favor.GameId = gameId
	favor.FavorLevel = 1
	favor.CreatedAt = time.Now().UTC()
	favor, err := dao.InsertFavor(favor)
	return favor, err
}