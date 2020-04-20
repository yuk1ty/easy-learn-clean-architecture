package registry

import usecase "github.com/yuk1ty/easy-learn-clean-architecture/iamclean/usecase/user"

type UseCaseRegistry struct {
	repositoryRegistry RepositoryRegistry
}

func NewUseCaseRegistry() UseCaseRegistry {
	return UseCaseRegistry{repositoryRegistry: NewRepositoryRegistry()}
}

func (r UseCaseRegistry) UserUseCase() usecase.UserUseCase {
	return usecase.NewUserUsecase(r.repositoryRegistry.UserRepository())
}