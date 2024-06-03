package search

import "github.com/OctavianoRyan25/be-agriculture/modules/plant"

type Usecase interface {
	Search(params PlantSearchParams) ([]plant.Plant, error)
}

type searchUsecase struct {
	repo Repository
}

func NewUsecase(repo Repository) *searchUsecase {
	return &searchUsecase{
		repo: repo,
	}
}

func (uc *searchUsecase) Search(params PlantSearchParams) ([]plant.Plant, error) {
	return uc.repo.Search(params)
}
