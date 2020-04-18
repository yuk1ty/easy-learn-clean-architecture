package response

import (
	role "github.com/yuk1ty/easy-learn-clean-architecture/iamclean/model/entity/role"
	user "github.com/yuk1ty/easy-learn-clean-architecture/iamclean/model/entity/user"
)

type UserJson struct {
	Id    string      `json:"id"`
	Name  string      `json:"name"`
	Roles []role.Role `json:"roles"` // TODO to DTO
}

func FromUser(user user.User) *UserJson {
	return &UserJson{
		Id:    user.Id.Value,
		Name:  user.Name.Value,
		Roles: user.Roles,
	}
}
