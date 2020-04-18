package vo

import "github.com/yuk1ty/easy-learn-clean-architecture/iamclean/context/apperr"

type UserName struct {
	Value string
}

func NewUserName(name string) (*UserName, error) {
	if len(name) > 15 {
		return nil, apperr.NewValidationErr("User Name は 15 文字以下である必要があります。")
	}
	return &UserName{Value: name}, nil
}
