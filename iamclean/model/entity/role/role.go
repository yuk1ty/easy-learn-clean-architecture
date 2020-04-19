package entity

import vo "github.com/yuk1ty/easy-learn-clean-architecture/iamclean/model/vo/role"

type Role struct {
	Id      *vo.RoleId
	Name    *vo.RoleName
	IsAdmin bool
}

func NewRole(name string, isAdmin bool) (*Role, error) {
	roleName, err := vo.NewRoleName(name)
	if err != nil {
		return nil, err
	}

	return &Role{
		Id:      vo.NewRoleId(),
		Name:    roleName,
		IsAdmin: isAdmin,
	}, nil
}
