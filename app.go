package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	"github.com/upamune/mirei-line-bot/controllers"
	"net/http"
	"github.com/labstack/echo/middleware"
)

var e *echo.Echo

func init() {
	e = echo.New()

	e.Use(middleware.Logger())

	e.GET("/", controllers.HelloHandler)

	line := e.Group("/line")
	line.POST("/callback", controllers.LineCallBackHandler)

	s := standard.New("")
	s.SetHandler(e)

	http.Handle("/", s)
}
