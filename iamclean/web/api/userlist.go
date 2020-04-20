package api

import (
	"github.com/labstack/echo/v4"
	usecase "github.com/yuk1ty/easy-learn-clean-architecture/iamclean/usecase/user"
	"net/http"
)

type UserListFilterEndpoint struct {
	userUseCase usecase.UserUseCase
}

func NewUserListFilterEndpoint(userUseCase usecase.UserUseCase) UserListFilterEndpoint {
	return UserListFilterEndpoint{userUseCase: userUseCase}
}

func (u UserListFilterEndpoint) FilterUserList(c echo.Context) error {
	return c.String(http.StatusOK, "UserList")
}
