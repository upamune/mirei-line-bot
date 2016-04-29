// +build !appengine

package main

import (
	"github.com/labstack/echo/engine/standard"
)

func main() {
	e.Run(standard.New(":8080"))
}
