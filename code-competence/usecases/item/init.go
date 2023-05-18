package item

import (
	"github.com/Ganes556/golang_I-Gusti-Agung-Ganes-Satsangga-Dipa/dto"
	"github.com/Ganes556/golang_I-Gusti-Agung-Ganes-Satsangga-Dipa/entities"
	"github.com/Ganes556/golang_I-Gusti-Agung-Ganes-Satsangga-Dipa/repositories/category"
	"github.com/Ganes556/golang_I-Gusti-Agung-Ganes-Satsangga-Dipa/repositories/item"
)

type ItemUsecase interface {
	GetAll() ([]entities.Item, error)
	GetById(id string) (entities.Item, error)
	Create(input dto.ItemCreateDTO) error
	Update(id string, input dto.ItemUpdateDTO) error
	Delete(id string) error
	GetByCategory(categoryId int) ([]entities.Item, error)
	GetByKeyword(keyword string) ([]entities.Item, error)
}


type itemUsecase struct {
	itemRepo item.ItemRepository
	categoryRepo category.CategoryRepository
}

func NewItemUsecase(itemRepo item.ItemRepository, categoryRepo category.CategoryRepository) ItemUsecase {
	return &itemUsecase{itemRepo: itemRepo, categoryRepo: categoryRepo}
}