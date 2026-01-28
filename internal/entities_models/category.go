package entitiesmodels

import (
	"fmt"
	"time"
)

type Category struct {
	ID       int       `json:"id"`
	Name     string    `json:"name"`
	CreateAt time.Time `json:"create_at"`
	UpdateAt time.Time `json:"update_at"`
}

func NewCategory(Name string) (*Category, error) {
	category := &Category{
		Name:     Name,
		CreateAt: time.Now(),
		UpdateAt: time.Now(),
	}

	//isso e um regra de negocio bem basica e mal feita em um mvc desente ele deveria ter um pacote de service

	err := category.IsValid()

	if err != nil {
		return nil, err
	}

	return category, nil
}

func (c *Category) IsValid() error {
	if len(c.Name) < 5 {
		return fmt.Errorf("nome muito pequeno. %d, len(c.Name))")
	}

	return nil
}
