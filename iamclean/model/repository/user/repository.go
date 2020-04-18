package user

import "github.com/yuk1ty/easy-learn-clean-architecture/iamclean/model/entity/user"

type UserRepository interface {
	FindAll() (*user.Users, error)
}
