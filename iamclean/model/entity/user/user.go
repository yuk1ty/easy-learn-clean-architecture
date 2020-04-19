package entity

import (
	"github.com/yuk1ty/easy-learn-clean-architecture/iamclean/model/entity/role"
	vo "github.com/yuk1ty/easy-learn-clean-architecture/iamclean/model/vo/user"
)

type User struct {
	Id    *vo.UserId
	Name  *vo.UserName
	Roles []entity.Role
}

func NewUser(name string, roles []entity.Role) (*User, error) {
	userName, err := vo.NewUserName(name)
	if err != nil {
		return nil, err
	}

	return &User{
		Id:    vo.NewUserId(),
		Name:  userName,
		Roles: roles,
	}, nil
}
