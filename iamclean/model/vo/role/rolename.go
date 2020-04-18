package vo

import "github.com/yuk1ty/easy-learn-clean-architecture/iamclean/context/apperr"

type RoleName struct {
	Value string
}

func NewRoleName(name string) (*RoleName, error) {
	if len(name) > 15 {
		return nil, apperr.NewValidationErr("Role Name は 15 文字以下である必要があります。")
	}
	return &RoleName{Value: name}, nil
}
