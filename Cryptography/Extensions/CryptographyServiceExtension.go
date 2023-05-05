package cryptography

import (
	interfaces "github.com/barisakdas/Framework/Cryptography/Interfaces"
	services "github.com/barisakdas/Framework/Cryptography/Services"
)

func NewFileLoggerService(key []byte) interfaces.ICryptographyService {
	return &services.CryptographyService{Key: key}
}
