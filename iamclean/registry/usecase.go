package registry

import userusecase "github.com/yuk1ty/easy-learn-clean-architecture/iamclean/usecase"

type UseCaseRegistry struct {
	repositoryRegistry RepositoryRegistry
}

func NewUseCaseRegistry() UseCaseRegistry {
	return UseCaseRegistry{repositoryRegistry: NewRepositoryRegistry()}
}

func (r UseCaseRegistry) UserUseCase() userusecase.UserUseCase {
	return userusecase.NewUserUsecase(r.repositoryRegistry.UserRepository())
}