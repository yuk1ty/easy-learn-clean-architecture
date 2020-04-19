package registry

import (
	"github.com/yuk1ty/easy-learn-clean-architecture/iamclean/adapter/web/api"
)

type EndpointRegistry struct {
	useCaseRegistry UseCaseRegistry
}

func NewEndpointRegistry() EndpointRegistry {
	return EndpointRegistry{useCaseRegistry: NewUseCaseRegistry()}
}

func (r EndpointRegistry) UserFilterListEndpoint() api.UserListFilterEndpoint {
	// TODO get from registries
	userUseCase := r.useCaseRegistry.UserUseCase()
	endpoint := api.NewUserListFilterEndpoint(userUseCase)
	return endpoint
}
