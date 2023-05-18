package category

import (
	"github.com/Ganes556/golang_I-Gusti-Agung-Ganes-Satsangga-Dipa/entities"
	"gorm.io/gorm"
)

type CategoryRepository interface {
	GetById(id uint) (entities.Category, error)
	Create(category entities.Category) error 
	GetByName(category entities.Category) error
}

type categoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) CategoryRepository {
	return &categoryRepository{db: db}
}
