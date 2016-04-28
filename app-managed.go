// +build appenginevm

package main

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	"google.golang.org/appengine"
)

func createMux() *echo.Echo {
	e := echo.New()

	return e
}

func main() {
	s := standard.New(":8080")
	s.SetHandler(e)
	http.Handle("/", s)
	appengine.Main()
}
