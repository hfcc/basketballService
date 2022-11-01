package dao

import (
	"basketballService/db"
	"basketballService/db/model"
)

// Getschedules
func Getschedules() ([]model.ScheduleModel, error) {
	var schedules []model.ScheduleModel
	cli := db.Get()
	var result = cli.Table("schedule").Find(&schedules)
	return schedules, result.Error
}	
func InsertSchedule(schedules []model.ScheduleModel) {
	cli := db.Get()
	cli.Table("schedule").CreateInBatches(schedules, 1500)
}
