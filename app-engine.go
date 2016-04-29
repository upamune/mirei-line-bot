// +build appengine

package main

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/fasthttp"
)

func createMux() *echo.Echo {
	e := echo.New()

	s := fasthttp.New("")
	s.SetHandler(e)
	http.Handle("/", s)

	return e
}
