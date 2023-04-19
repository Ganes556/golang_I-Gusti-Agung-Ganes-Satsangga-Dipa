package pkg

import "golang.org/x/crypto/bcrypt"

type Password interface {
	HashPassword(password string) (string, error)
	CheckPasswordHash(password, hash string) bool
}

type password struct{}

func NewPassword() Password {
	return &password{}
}

func (p *password) HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func (p *password) CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}