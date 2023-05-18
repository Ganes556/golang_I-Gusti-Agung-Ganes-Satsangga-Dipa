package category

import "github.com/Ganes556/golang_I-Gusti-Agung-Ganes-Satsangga-Dipa/entities"

func (c *categoryRepository) GetById(id uint) (entities.Category, error) {
	var category entities.Category

	err := c.db.Where("id = ?", id).First(&category).Error
	if err != nil {
		return category, err
	}

	return category, nil
}

func (c *categoryRepository) Create(category entities.Category) error {
	err := c.db.Create(&category).Error
	if err != nil {
		return err
	}

	return nil
}


func (c *categoryRepository) GetByName(category entities.Category) error {
	err := c.db.Where("name = ?", category.Name).First(&category).Error
	if err != nil {
		return err
	}

	return nil
}