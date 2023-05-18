package item

import (
	"errors"

	"github.com/Ganes556/golang_I-Gusti-Agung-Ganes-Satsangga-Dipa/dto"
	"github.com/Ganes556/golang_I-Gusti-Agung-Ganes-Satsangga-Dipa/entities"
)

func (i *itemUsecase) GetAll() ([]entities.Item, error) {

	items, err := i.itemRepo.GetAll()

	if err != nil {
		if err.Error() == "record not found" {
			return items, errors.New("item not found")
		}
		return items, errors.New("internal server error")
	}

	return items, nil

}

func (i *itemUsecase) Create(input dto.ItemCreateDTO) error {
	
	_, err := i.categoryRepo.GetById(input.CategoryID)

	if err != nil {
		return errors.New("category not found")
	}

	item := entities.Item{
		Name: input.Name,
		Price: input.Price,
		CategoryID: input.CategoryID,
		Stock: input.Stock,
		Description: input.Description,
	}

	err = i.itemRepo.Create(item)

	if err != nil {
		return errors.New("internal server error")
	}

	return nil
}


func (i *itemUsecase) GetById(id string) (entities.Item, error) {

	item, err := i.itemRepo.GetById(id)

	if err != nil {
		if err.Error() == "record not found" {
			return item, errors.New("item not found")
		}
		return item, errors.New("internal server error")
	}

	return item, nil
}

func (i *itemUsecase) Delete(id string) error {
	
	_, err := i.itemRepo.GetById(id)

	if err != nil {
		if err.Error() == "record not found" {
			return errors.New("item not found")
		}
		return errors.New("internal server error")
	}

	err = i.itemRepo.Delete(id)

	if err != nil {
		return err
	}

	return nil
}

func (i *itemUsecase) Update(id string, input dto.ItemUpdateDTO) error {
	
	_, err := i.itemRepo.GetById(id)

	if err != nil {
		if err.Error() == "record not found" {
			return errors.New("item not found")
		}
		return errors.New("internal server error")
	}

	item := entities.Item{
		Name: input.Name,
		Price: input.Price,
		CategoryID: input.CategoryID,
		Stock: input.Stock,
		Description: input.Description,
	}

	err = i.itemRepo.Update(id, item)

	if err != nil {
		return errors.New("internal server error")
	}

	return nil

}

func (i *itemUsecase) GetByCategory(categoryId int) ([]entities.Item, error) {
	
	items, err := i.itemRepo.GetByCategory(categoryId)

	if err != nil {
		return items, errors.New("internal server error")
	}

	return items, nil
}

func (i *itemUsecase) GetByKeyword(keyword string) ([]entities.Item, error) {
	
	items, err := i.itemRepo.GetByKeyword(keyword)

	if err != nil {
		return items, errors.New("internal server error")
	}

	if len(items) == 0 {
		return items, errors.New("item not found")
	}

	return items, nil
}