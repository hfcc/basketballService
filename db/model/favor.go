package model

import "time"

type FavorModel struct {
  UserId    int32   `gorm:"column:userId" json:"userId"`
	GameId  string   `gorm:"column:gameId" json:"gameId"`
	FavorLevel  int   `gorm:"column:favorLevel" json:"favorLevel"`
	CreatedAt time.Time `gorm:"column:createdAt" json:"createdAt"`
}