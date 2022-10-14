package main

import (
	"basketballService/db"
	"basketballService/service"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	if err := db.Init(); err != nil {
		panic(fmt.Sprintf("mysql init failed with %+v", err))
	}
	r := gin.Default()
	r.GET("/schedules", func(c *gin.Context) {
		schedules, err := service.Getschedules()
		c.JSON(200, gin.H{
			"schedules": schedules,
      "err": err,
		})
	})
	r.Run()
}

