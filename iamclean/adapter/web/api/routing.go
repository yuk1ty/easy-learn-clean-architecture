package api

import (
	"github.com/labstack/echo/v4"
	"github.com/yuk1ty/easy-learn-clean-architecture/iamclean/registry"
)

func Routing(e *echo.Echo, registry registry.EndpointRegistry) {
	e.GET("/users", registry.UserFilterListEndpoint().FilterUserList)
}