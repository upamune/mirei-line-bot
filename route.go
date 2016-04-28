package main

import "github.com/upamune/mirei-line-bot/controllers"

func init() {
	e.GET("/", controllers.HelloHandler)
}
