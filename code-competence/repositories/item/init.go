package item

import (
	"github.com/Ganes556/golang_I-Gusti-Agung-Ganes-Satsangga-Dipa/entities"
	"gorm.io/gorm"
)

type ItemRepository interface {
	GetAll() ([]entities.Item, error)
	GetById(id string) (entities.Item, error)
	Create(item entities.Item) error
	Update(id string, item entities.Item) error
	Delete(id string) error
	GetByCategory(categoryId int) ([]entities.Item, error)
	GetByKeyword(keyword string) ([]entities.Item, error)
}


type itemRepository struct {
	db *gorm.DB	
}

func NewItemRepository(db *gorm.DB) ItemRepository {
	return &itemRepository{db: db}
}