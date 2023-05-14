package user

import (
	"github.com/Ganes556/golang_I-Gusti-Agung-Ganes-Satsangga-Dipa/entities"
	"github.com/google/uuid"
)

func (ur *userRepository) GetByID(id uuid.UUID) (entities.User, error) {
	user := entities.User{}
	err := ur.db.First(&user, "id = ?", id)
	if err != nil {
		return user, err.Error
	}
	return user, nil
}
func (ur *userRepository) GetByEmail(email string) (entities.User, error) {
	user := entities.User{}
	err := ur.db.First(&user, "email = ?", email).Error
	if err != nil {
		return user, err
	}
	return user, nil
}
func (ur *userRepository) Create(ut entities.User) error {
	err := ur.db.Create(&ut).Error
	if err != nil {
		return err
	}
	return nil
}