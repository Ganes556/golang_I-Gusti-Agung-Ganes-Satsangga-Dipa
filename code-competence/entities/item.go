package entities

import (
	"strings"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Item struct {
	Base
	ID string `gorm:"type:char(32);primary_key;index"`
	CategoryID uint `json:"category_id"`
	Name string `json:"name" gorm:"type:varchar(150)"`
	Price float64 `json:"price" gorm:"type:decimal(10,2)"`
	Stock int `json:"stock" gorm:"type:int"`
	Description string `json:"description"`
}

func (i *Item) BeforeCreate(tx *gorm.DB) error {
	uuidWithHyphen := uuid.New()
	uuid := strings.Replace(uuidWithHyphen.String(), "-", "", -1)
	i.ID = uuid
	return nil
}

