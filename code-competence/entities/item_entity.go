package entities

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Item struct {
	Base
	ID string `json:"id" gorm:"type:char(36);primary_key;uniqueIndex"`
	CategoryID uint `json:"category_id"`
	Name string `json:"name" gorm:"type:varchar(150)"`
	Price float64 `json:"price" gorm:"type:decimal(10,2)"`
	Stock uint `json:"stock" gorm:"type:uint"`
	Description string `json:"description"`
}

func (i *Item) BeforeCreate(tx *gorm.DB) error {
	i.ID = uuid.NewString()
	return nil
}