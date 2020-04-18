package entity

import (
	"github.com/yuk1ty/easy-learn-clean-architecture/iamclean/model/entity/role"
	"github.com/yuk1ty/easy-learn-clean-architecture/iamclean/model/vo/user"
)

type User struct {
	Id    *user.Id
	Name  *user.Name
	Roles []entity.Role
}

func NewUser(name string, roles []entity.Role) (*User, error) {
	userName, err := user.ValidateWith(name)
	if err != nil {
		return nil, err
	}

	return &User{
		Id:    user.NewId(),
		Name:  userName,
		Roles: roles,
	}, nil
}
