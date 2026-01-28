package services

import (
	"log"

	entitiesmodels "github.com/gabrielfranca42/go-microservices/internal/entities_models"
)

type CreateCategoryUseCase struct {
}

func NewCreateCategoryUseCase() *CreateCategoryUseCase {
	return &CreateCategoryUseCase{}
}

func (u *CreateCategoryUseCase) Execute(name string) error {
	category, err := entitiesmodels.NewCategory(name)

	if err != nil {
		return err
	}

	log.Println(category)

	return nil
}
