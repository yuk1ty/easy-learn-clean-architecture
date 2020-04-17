package user

import "github.com/yuk1ty/easy-learn-clean-architecture/iamclean/usecase/dto/user"

type UserListResponse struct {
	Items []user.UserDto `json:"items"`
}
