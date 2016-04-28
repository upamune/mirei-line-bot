package main

import "github.com/upamune/mirei-line-bot/controllers"

func init() {
	e.GET("/", controllers.HelloHandler)

	line := e.Group("/line")
	line.POST("/callback", controllers.LineCallBackHandler)
}
