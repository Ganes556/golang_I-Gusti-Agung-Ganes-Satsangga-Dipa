package entities

type Category struct {
	Base
	ID          uint   `json:"id" gorm:"primary_key"`
	Name        string `json:"name" gorm:"type:varchar(150)"`
	Description string `json:"description"`
	Items       []Item `json:"items" gorm:"foreignKey:CategoryID"`
}
