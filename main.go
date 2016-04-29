// +build !appengine

package app

import (
	"github.com/labstack/echo/engine/standard"
)

func main() {
	e.Run(standard.New(":8080"))
}
