package registry

import (
	repository "github.com/yuk1ty/easy-learn-clean-architecture/iamclean/model/repository/user"
)
import infra "github.com/yuk1ty/easy-learn-clean-architecture/iamclean/adapter/infra/repository/user"

type RepositoryRegistry struct {
	mysqlDaoRegistry MySqlDaoRegistry
}

func NewRepositoryRegistry() RepositoryRegistry {
	return RepositoryRegistry{mysqlDaoRegistry: NewMySqlDaoRegistry()}
}

func (r RepositoryRegistry) UserRepository() repository.UserRepository {
	return infra.NewUserRepository(r.mysqlDaoRegistry.UserDao())
}