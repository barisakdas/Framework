package identity

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

type TokenService struct {
	SecretKey string
}

func (ts *TokenService) GenerateToken(username, password string) (string, error) {
	// Token'ın ömrü 1 saat olarak ayarlanıyor.
	expirationTime := time.Now().Add(time.Hour * 1).Unix()

	claims := jwt.MapClaims{}
	claims["username"] = username
	claims["password"] = password
	claims["exp"] = expirationTime

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(ts.SecretKey))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}
