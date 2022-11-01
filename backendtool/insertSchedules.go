package backendtool

import (
	"basketballService/db/dao"
	"basketballService/db/model"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"
	"time"
)
func InsertSchedules() {
	  bfRead, err := ioutil.ReadFile("./backendtool/gameList.json")
		if err != nil {
			fmt.Println("读取json文件失败", err)
			return
	  }
		res := &NbaScheduleStruct{}
	  json.Unmarshal([]byte(bfRead), &res)
		games := res.Standard
		recordLength := len(games)
		var schedules []model.ScheduleModel
		for i := 0; i < recordLength; i++ {
			game := games[i]
			var schedule model.ScheduleModel
			schedule.CreatedAt = time.Now().UTC()
			schedule.GameId = game.GameID
			hScore, _ := strconv.ParseInt(game.HTeam.Score, 10, 16)
			schedule.HScore = int16(hScore)
			schedule.HTeamId = game.HTeam.TeamID
			schedule.SeasonName = "2022-2023"
			schedule.SeasonStageId = game.SeasonStageID
			schedule.StartTimeUTC = game.StartTimeUTC
			vScore, _ := strconv.ParseInt(game.VTeam.Score, 10, 16)
			schedule.VScore = int16(vScore)
			schedule.VTeamId = game.VTeam.TeamID
			schedules = append(schedules, schedule)
		}
		fmt.Print("schedules length: ")
		fmt.Println(len(schedules))
		dao.InsertSchedule(schedules)
}