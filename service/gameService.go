package service

import (
	"basketballService/db/dao"
	"basketballService/db/model"
)
func Getschedules() (*[]model.ScheduleModel, error) {
	schedule, err := dao.Getschedules()
	if err != nil {
		return nil, err
	}
	return schedule, nil
}
