package cryptography

type ICryptographyService interface {
	Encrypt(text string) (string, error)
	Decrypt(text string) (string, error)
}
