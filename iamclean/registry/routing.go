package registry

import (
	"github.com/yuk1ty/easy-learn-clean-architecture/iamclean/adapter/web/api"
)

type RoutingRegistry struct {
	endpointRegistry EndpointRegistry
}

func NewRoutingRegistry() RoutingRegistry {
	return RoutingRegistry{endpointRegistry: NewEndpointRegistry()}
}

func (r RoutingRegistry) Route() api.Routing {
	return api.NewRouting(r.endpointRegistry.UserFilterListEndpoint())
}