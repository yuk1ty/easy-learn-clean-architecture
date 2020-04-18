package entity

import "github.com/yuk1ty/easy-learn-clean-architecture/iamclean/model/vo/role"

type Role struct {
	Id *role.Id
	Name *role.Name
	IsAdmin bool
}

func NewRole(name string, isAdmin bool) (*Role, error) {
	roleName, err := role.ValidateWith(name)
	if err != nil {
		return nil, err
	}

	return &Role{
		Id:      role.NewId(),
		Name:    roleName,
		IsAdmin: isAdmin,
	}, nil
}