package registry

import (
	mysql "github.com/yuk1ty/easy-learn-clean-architecture/iamclean/adapter/infra/mysql/user"
	userrepository "github.com/yuk1ty/easy-learn-clean-architecture/iamclean/adapter/infra/repository/user"
	"github.com/yuk1ty/easy-learn-clean-architecture/iamclean/adapter/web/api"
	userusecase "github.com/yuk1ty/easy-learn-clean-architecture/iamclean/usecase"
)

type EndpointRegistry struct{}

func NewEndpointRegistry() EndpointRegistry {
	return EndpointRegistry{}
}

func (r EndpointRegistry) UserFilterListEndpoint() api.UserListFilterEndpoint {
	// TODO get from registries
	userDao := mysql.NewUserDao()
	userRepository := userrepository.NewUserRepository(userDao)
	userUseCase := userusecase.NewUserUsecase(userRepository)
	endpoint := api.NewUserListFilterEndpoint(userUseCase)
	return endpoint
}
