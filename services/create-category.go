package services

import (
	entitiesmodels "github.com/gabrielfranca42/go-microservices/internal/entities_models"
	"github.com/gabrielfranca42/go-microservices/internal/entities_models/repository"
)

type CreateCategoryUseCase struct {
	repository repository.ICategoryRepository
}

func NewCreateCategoryUseCase(repository repository.ICategoryRepository) *CreateCategoryUseCase {
	return &CreateCategoryUseCase{repository}
}

func (u *CreateCategoryUseCase) Execute(name string) error {
	category, err := entitiesmodels.NewCategory(name)

	if err != nil {
		return err
	}

	err = u.repository.Save(category)

	if err != nil {
		return err
	}

	return nil
}
