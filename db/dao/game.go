package dao

import (
	"basketballService/db"
	"basketballService/db/model"
	"fmt"
)

// Getschedules
func Getschedules() (*[]model.ScheduleModel, error) {
	var schedules *[]model.ScheduleModel
	cli := db.Get()
	fmt.Print(cli)
	var result = cli.Table("schedule").Find(&schedules)
	return schedules, result.Error
}	
