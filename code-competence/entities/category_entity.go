package entities

type Category struct {
	ID    uint   `json:"id" gorm:"primary_key"`
	Name  string `json:"name" gorm:"type:varchar(150)"`
	Items []Item `json:"items" gorm:"foreignKey:CategoryID"`
}
