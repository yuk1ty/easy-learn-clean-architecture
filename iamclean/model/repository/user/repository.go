package repository

import "github.com/yuk1ty/easy-learn-clean-architecture/iamclean/model/entity/user"

type UserRepository interface {
	FindAll() (*entity.Users, error)
}
