package repository

import entitiesmodels "github.com/gabrielfranca42/go-microservices/internal/entities_models"

type ICategoryRepository interface {
	Save(category *entitiesmodels.Category) error
	List () ([ ] *entitiesmodels.Category,error)
}
