package main

import (
	"github.com/labstack/echo/v4"
	"github.com/yuk1ty/easy-learn-clean-architecture/iamclean/adapter/web/api"
	"github.com/yuk1ty/easy-learn-clean-architecture/iamclean/registry"
)

func main() {
	e := echo.New()
	reg := registry.NewEndpointRegistry()
	api.Routing(e, reg)
	e.Logger.Fatal(e.Start(":1323"))
}
