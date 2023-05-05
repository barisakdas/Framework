package secret

import (
	interfaces "github.com/barisakdas/Framework/Secret/Interfaces"
	services "github.com/barisakdas/Framework/Secret/Services"
)

func NewFileLoggerService(url string) interfaces.IAzureKeyVaultService {
	return &services.AzureKeyVaultService{Url: url}
}
