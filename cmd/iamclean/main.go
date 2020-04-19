package main

import (
	"github.com/labstack/echo/v4"
	"github.com/yuk1ty/easy-learn-clean-architecture/iamclean/registry"
)

func main() {
	e := echo.New()
	registry.NewRoutingRegistry().Route()
	e.Logger.Fatal(e.Start(":1323"))
}
