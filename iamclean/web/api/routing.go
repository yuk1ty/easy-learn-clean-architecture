package api

import (
	"github.com/labstack/echo/v4"
)

type Routing struct {
	userListFilterEndpoint UserListFilterEndpoint
}

func NewRouting(userListFilterEndpoint UserListFilterEndpoint) Routing {
	return Routing{
		userListFilterEndpoint: userListFilterEndpoint,
	}
}

func (r Routing) Routes(e *echo.Echo) {
	e.GET("/users", r.userListFilterEndpoint.FilterUserList)
}