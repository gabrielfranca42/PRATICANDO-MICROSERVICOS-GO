package services

import (
	entitiesmodels "github.com/gabrielfranca42/go-microservices/internal/entities_models"
	"github.com/gabrielfranca42/go-microservices/internal/entities_models/repository"
)

type ListCategoryUseCase struct {
	repository repository.ICategoryRepository
}

func NewListCategoryUseCase(repository repository.ICategoryRepository) *ListCategoryUseCase {
	return &ListCategoryUseCase{
		repository: repository,
	}
}

func (u *ListCategoryUseCase) Execute() ([]*entitiesmodels.Category, error)  {
	category, err := u.repository.List()

	if err != nil {
		return nil, err
	}

	

	return category, nil
}
