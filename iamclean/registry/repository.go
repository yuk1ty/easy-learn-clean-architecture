package registry

import (
	mysql "github.com/yuk1ty/easy-learn-clean-architecture/iamclean/adapter/infra/mysql/user"
	repository "github.com/yuk1ty/easy-learn-clean-architecture/iamclean/model/repository/user"
)
import infra "github.com/yuk1ty/easy-learn-clean-architecture/iamclean/adapter/infra/repository/user"

type RepositoryRegistry struct {}

func NewRepositoryRegistry() RepositoryRegistry {
	return RepositoryRegistry{}
}

func (r RepositoryRegistry) UserRepository() repository.UserRepository {
	return infra.NewUserRepository(mysql.NewUserDao()) // TODO
}