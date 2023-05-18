package dto

type ItemCreateDTO struct {
	CategoryID  uint    `json:"category_id" validate:"required,number"`
	Name        string  `json:"name" validate:"required"`
	Price       float64 `json:"price" validate:"required,number"`
	Stock       uint    `json:"stock" validate:"required,number"`
	Description string  `json:"description" validate:"required"`
}

type ItemUpdateDTO struct {
	CategoryID  uint    `json:"category_id" validate:"number"`
	Name        string  `json:"name"`
	Price       float64 `json:"price" validate:"number"`
	Stock       uint    `json:"stock" validate:"number"`
	Description string  `json:"description"`
}