package manager

import "github.com/jutionck/go-sinar-makmur-mongodb/repository"

type RepoManager interface {
	BrandRepo() repository.BrandRepository
}

type repoManager struct {
	infra InfraManager
}

func (r *repoManager) BrandRepo() repository.BrandRepository {
	return repository.NewBrandRepository(r.infra.Conn())
}

func NewRepoManager(infra InfraManager) RepoManager {
	return &repoManager{infra: infra}
}
