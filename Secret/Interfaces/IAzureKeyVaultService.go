package secret

type IAzureKeyVaultService interface {
	GetSecret(secret string) (string, error)
	CreateSecret(secret, value string) error
	DeleteSecret(secret string) error
}
