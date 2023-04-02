package models

import (
	"time"

	"gorm.io/gorm"
)

type Base struct {
	ID        uint           `gorm:"primarykey" json:"id" form:"id"`
	CreatedAt time.Time      `json:"-" form:"-"`
	UpdatedAt time.Time      `json:"-" form:"-"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-" form:"-"`
}
