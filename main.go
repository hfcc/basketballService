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
	r.POST("/favor/:gameId", func(c *gin.Context) {
		openIds := c.Request.Header["X-Wx-Openid"]
		if (len(openIds) == 0) {
			c.JSON(200, gin.H{
				"favorInfo": nil,
				"err": nil,
				"errCode": 100,
			})
			return
		}
		openId := openIds[0]
		userInfo, err := service.GetUserByOpenId(openId)
		if err != nil && userInfo.Id <= 0 {
			c.JSON(200, gin.H{
				"favorInfo": nil,
				"err": err,
				"errCode": 101, // 不是我们的用户
			})
			return
		}
		gameId := c.Param("gameId")
		favorInfo, err := service.InsertFavor(userInfo.Id, gameId)
		if (err != nil) {
			c.JSON(200, gin.H{
				"favorInfo": nil,
				"err": err,
			})
			return
		}
		c.JSON(200, gin.H{
			"favorInfo": favorInfo,
			"err": err,
		})
	})
	/* r.GET("/test", func(c *gin.Context) {
		backendtool.InsertSchedules()
		c.JSON(200, gin.H{
			"ok": "",
      "err": "",
		})
	}) */
	r.Run()
}

