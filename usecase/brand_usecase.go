package usecase

import (
	"github.com/jutionck/go-sinar-makmur-mongodb/model"
	"github.com/jutionck/go-sinar-makmur-mongodb/repository"
)

type BrandUseCase interface {
	RegisterNewBrand(payload *model.Brand) error
	FindAllBrand() ([]model.Brand, error)
	FindById(id string) (model.Brand, error)
}

type brandUseCase struct {
	repo repository.BrandRepository
}

func (b *brandUseCase) RegisterNewBrand(payload *model.Brand) error {
	return b.repo.Create(payload)
}

func (b *brandUseCase) FindAllBrand() ([]model.Brand, error) {
	return b.repo.List()
}

func (b *brandUseCase) FindById(id string) (model.Brand, error) {
	return b.repo.Get(id)
}

func NewBrandUseCase(repo repository.BrandRepository) BrandUseCase {
	return &brandUseCase{repo: repo}
}
