package usecase

import (
	entity "github.com/yuk1ty/easy-learn-clean-architecture/iamclean/model/entity/user"
	repository "github.com/yuk1ty/easy-learn-clean-architecture/iamclean/model/repository/user"
)

type UserUseCase struct {
	userRepository repository.UserRepository // FIXME ← why the namespace to be user2?
}

func (u *UserUseCase) AllUsers() (*entity.Users, error) {
	return u.userRepository.FindAll()
}