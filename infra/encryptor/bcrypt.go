package encryptor

import "golang.org/x/crypto/bcrypt"

type BCrypt struct {
	cost int
}

func NewBCrypt(cost int) *BCrypt {
	return &BCrypt{cost: cost}
}

func (b *BCrypt) Encrypt(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), b.cost)

	return string(hash), err
}

func (b *BCrypt) Compare(password, hash string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err == nil {
		return true, nil
	}

	if err == bcrypt.ErrMismatchedHashAndPassword {
		return false, nil
	}

	panic(err)
}
