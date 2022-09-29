package encryptor

type PasswordEncryptor interface {
	Encrypt(password string) string
	Compare(password, hash string) bool
}
