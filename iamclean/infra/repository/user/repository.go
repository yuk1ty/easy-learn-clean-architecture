package repository

import (
	mysql "github.com/yuk1ty/easy-learn-clean-architecture/iamclean/infra/mysql/user"
	entity "github.com/yuk1ty/easy-learn-clean-architecture/iamclean/model/entity/user"
	repository "github.com/yuk1ty/easy-learn-clean-architecture/iamclean/model/repository/user"
)

type UserRepositoryImpl struct {
	dao mysql.UserDao
}

func NewUserRepository(dao mysql.UserDao) repository.UserRepository {
	return &UserRepositoryImpl{dao: dao}
}

func (u *UserRepositoryImpl) FindAll() (*entity.Users, error) {
	var gormUsers []mysql.GormUser = u.dao.FindAllUsers()
	// Gorm User から User に変換する処理を噛ませる
	return nil, nil
}
