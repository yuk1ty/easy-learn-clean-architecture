package user

import (
	"github.com/yuk1ty/easy-learn-clean-architecture/iamclean/model/entity/role"
	"github.com/yuk1ty/easy-learn-clean-architecture/iamclean/model/entity/user"
)

type UserDto struct {
	Id    string      `json:"id"`
	Name  string      `json:"name"`
	Roles []role.Role `json:"roles"` // TODO to DTO
}

func FromUser(user user.User) *UserDto {
	return &UserDto{
		Id:    user.Id.Value,
		Name:  user.Name.Value,
		Roles: user.Roles,
	}
}
