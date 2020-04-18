package usecase

import (
	"github.com/yuk1ty/easy-learn-clean-architecture/iamclean/model/entity/user"
	user2 "github.com/yuk1ty/easy-learn-clean-architecture/iamclean/model/repository/user"
)

type UserUseCase struct {
	userRepository user2.UserRepository // FIXME ← why the namespace to be user2?
}

func (u *UserUseCase) AllUsers() (*user.Users, error) {
	return u.userRepository.FindAll()
}