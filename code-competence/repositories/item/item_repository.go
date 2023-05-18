package item

import "github.com/Ganes556/golang_I-Gusti-Agung-Ganes-Satsangga-Dipa/entities"

func (i *itemRepository) GetAll() ([]entities.Item, error){
	items := []entities.Item{}
	err := i.db.Find(&items).Error

	if err != nil {
		return items, err
	}

	return items, nil
	
}

func (i *itemRepository) GetById(id string) (entities.Item, error){
	item := entities.Item{}
	
	err := i.db.First(&item, "id = ?", id).Error

	if err != nil {
		return item, err
	}

	return item, nil
}

func (i *itemRepository) Create(item entities.Item) error{
	err := i.db.Create(&item).Error

	if err != nil {
		return err
	}

	return nil
}

func (i *itemRepository) Update(id string, item entities.Item) error{
	err := i.db.Model(&item).Where("id = ?", id).Updates(item).Error

	if err != nil {
		return err
	}

	return nil
}

func (i *itemRepository) Delete(id string) error{
	item := entities.Item{}
	err := i.db.Delete(&item, "id = ?", id).Error

	if err != nil {
		return err
	}

	return nil
}

func (i *itemRepository) GetByCategory(categoryId int) ([]entities.Item, error){
	items := []entities.Item{}
	err := i.db.Where("category_id = ?", categoryId).Find(&items).Error

	if err != nil {
		return items, err
	}

	return items, nil
}

func (i *itemRepository) GetByKeyword(keyword string) ([]entities.Item, error){
	items := []entities.Item{}
	err := i.db.Where("name LIKE ?", "%"+ keyword + "%").Find(&items).Error

	if err != nil {
		return items, err
	}

	return items, nil
}