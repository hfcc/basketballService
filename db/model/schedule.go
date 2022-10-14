package model

import "time"

type ScheduleModel struct {
	Id        int     `gorm:"column:id" json:"id"`
	SeasonName    string   `gorm:"column:seasonName" json:"seasonName"`
	GameId  string   `gorm:"column:gameId" json:"gameId"`
	SeasonStageId int `gorm:"column:seasonStageId" json:"seasonStageId"`
	StartTimeUTC time.Time `gorm:"column:startTimeUTC" json:"startTimeUTC"`
	HdoTeamId string   `gorm:"column:hTeamId" json:"hTeamId"`
	HScore        int     `gorm:"column:hScore" json:"hScore"`
	VTeamId string   `gorm:"column:vTeamId" json:"vTeamId"`
	VScore        int     `gorm:"column:vScore" json:"vScore"`
	CreatedAt time.Time `gorm:"column:createdAt" json:"createdAt"`
	UpdatedAt time.Time `gorm:"column:updatedAt" json:"updatedAt"`
}