package manager

import "github.com/jutionck/go-sinar-makmur-mongodb/usecase"

type UseCaseManager interface {
	BrandUseCase() usecase.BrandUseCase
}

type useCaseManager struct {
	repoManager RepoManager
}

func (u *useCaseManager) BrandUseCase() usecase.BrandUseCase {
	return usecase.NewBrandUseCase(u.repoManager.BrandRepo())
}

func NewUseCaseManager(repoManager RepoManager) UseCaseManager {
	return &useCaseManager{repoManager: repoManager}
}
