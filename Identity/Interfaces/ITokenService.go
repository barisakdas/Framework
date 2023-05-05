package identity

type ITokenService interface {
	GenerateToken(username, password string) (string, error)
}
