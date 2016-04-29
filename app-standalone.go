// +build !appengine,!appenginevm

package main

import (
	"fmt"

	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/fasthttp"
	"github.com/labstack/echo/middleware"
)

func createMux() *echo.Echo {
	e := echo.New()
	e.Use(middleware.Logger())

	return e
}

func main() {
	fmt.Println("Running... :8080")

	e.Run(fasthttp.New(":8080"))
}
