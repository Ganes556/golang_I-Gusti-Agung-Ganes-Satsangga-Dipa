package dto

type CategoryCreateDTO struct {
	Name string `json:"name" validate:"required"`
}