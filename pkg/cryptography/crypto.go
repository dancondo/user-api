package cryptography

import "golang.org/x/crypto/bcrypt"

type Crypto interface {
	EncryptPassword(password string) (string, error)
	ValidatePassword(value string, comparation string) bool
}

type crypto struct{}

func New() Crypto {
	return &crypto{}
}

func (c *crypto) EncryptPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)

	if err != nil {
		return "", err
	}

	return string(bytes), nil
}

func (c *crypto) ValidatePassword(value string, comparation string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(value), []byte(comparation))

	if err != nil {
		return false
	}

	return true
}
