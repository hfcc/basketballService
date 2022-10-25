package service

import (
	"basketballService/db/dao"
	"basketballService/db/model"

	"github.com/samber/lo"
)
func Getschedules() (map[string][]model.ScheduleModel, error) {
	schedules, err := dao.Getschedules()
	if err != nil {
		return nil, err
	}
	schedulesInDay := lo.GroupBy(schedules, func (schedule model.ScheduleModel) string  {
		return schedule.StartTimeUTC.Format("2006-01-02")
	})
	return schedulesInDay, nil
}
