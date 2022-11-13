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
	r.POST("/login", func(c *gin.Context) {
		openIds := c.Request.Header["X-Wx-Openid"]
		if (len(openIds) == 0) {
			c.JSON(200, gin.H{
				"userInfo": nil,
			})
			return
		}
		openId := openIds[0]
		userInfo, err := service.NewUser(openId)
		c.JSON(200, gin.H{
			"userInfo": userInfo,
			"err": err,
		})
	})
	r.Run()
}

