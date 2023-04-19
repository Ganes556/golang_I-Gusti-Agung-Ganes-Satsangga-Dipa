package utils

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string, hashPtr *string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	*hashPtr = string(hash)
	return nil
}

func ComparePassword(hashPassword string, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(password))
}