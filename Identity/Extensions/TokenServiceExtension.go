package identity

import (
	interfaces "github.com/barisakdas/Framework/Identity/Interfaces"
	services "github.com/barisakdas/Framework/Identity/Services"
)

func NewFileLoggerService(secretKey string) interfaces.ITokenService {
	return &services.TokenService{SecretKey: secretKey}
}
