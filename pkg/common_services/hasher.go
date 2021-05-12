package common_services

import "golang.org/x/crypto/bcrypt"

type Hasher interface {
	HashPassword(password string) (string, error)
	CheckPassword(hashedPassword, password string) error
}

type AppHasher struct {}

func (h *AppHasher) HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}

func (h *AppHasher) CheckPassword(hashedPassword, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return err
	}

	return nil
}
