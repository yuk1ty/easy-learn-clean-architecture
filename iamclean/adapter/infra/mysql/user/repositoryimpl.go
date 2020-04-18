package user

import "github.com/yuk1ty/easy-learn-clean-architecture/iamclean/model/entity/user"

type UserRepositoryImpl struct {

}

func (u *UserRepositoryImpl) FindAll() (*user.Users, error) {
	return nil, nil
}