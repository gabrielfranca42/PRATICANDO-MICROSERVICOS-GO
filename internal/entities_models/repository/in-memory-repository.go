package repository

import entitiesmodels "github.com/gabrielfranca42/go-microservices/internal/entities_models"

type inMemoryCategoryRepository struct {
	db []*entitiesmodels.Category
}

func NewInMemoryCategoryRepository() *inMemoryCategoryRepository {
	return &inMemoryCategoryRepository{
		db: make([]*entitiesmodels.Category, 0),
	}
}

func (r *inMemoryCategoryRepository) Save(category *entitiesmodels.Category) error {
	r.db = append(r.db, category)

	return nil
}
