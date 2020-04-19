package api

import (
	"github.com/labstack/echo/v4"
	user "github.com/yuk1ty/easy-learn-clean-architecture/iamclean/usecase"
	"net/http"
)

type UserListFilterEndpoint struct {
	userUseCase user.UserUseCase
}

func NewUserListFilterEndpoint(userUseCase user.UserUseCase) UserListFilterEndpoint {
	return UserListFilterEndpoint{userUseCase: userUseCase}
}

func (u UserListFilterEndpoint) FilterUserList(c echo.Context) error {
	return c.String(http.StatusOK, "UserList")
}
