package user

import (
	"github.com/Ganes556/golang_I-Gusti-Agung-Ganes-Satsangga-Dipa/entities"
	"gorm.io/gorm"
)

type UserRepository interface {
	GetByID(id uint) (entities.User, error)
	GetByEmail(email string) (entities.User, error)
	Create(ut entities.User) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}