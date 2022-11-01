package model

import "time"

type ScheduleModel struct {
	Id        int32     `gorm:"column:id" json:"id"`
	SeasonName    string   `gorm:"column:seasonName" json:"seasonName"`
	GameId  string   `gorm:"column:gameId" json:"gameId"`
	SeasonStageId int `gorm:"column:seasonStageId" json:"seasonStageId"`
	StartTimeUTC time.Time `gorm:"column:startTimeUTC" json:"startTimeUTC"`
	HTeamId string   `gorm:"column:hTeamId" json:"hTeamId"`
	HScore        int16     `gorm:"column:hScore" json:"hScore"`
	VTeamId string   `gorm:"column:vTeamId" json:"vTeamId"`
	VScore        int16     `gorm:"column:vScore" json:"vScore"`
	CreatedAt time.Time `gorm:"column:createdAt" json:"createdAt"`
	UpdatedAt time.Time `gorm:"column:updatedAt" json:"updatedAt"`
}