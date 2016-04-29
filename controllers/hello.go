package controllers

import (
	"net/http"

	"github.com/labstack/echo"
)

func HelloHandler(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World")
}
