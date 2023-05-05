package secret

import (
	"context"
	"log"
	"os"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/keyvault/azsecrets"
)

type AzureKeyVaultService struct {
	Url string
}

func (s *AzureKeyVaultService) GetSecret(secret string) (string, error) {
	vaultURI := os.Getenv("AZURE_KEY_VAULT_URI")

	// Create a credential using the NewDefaultAzureCredential type.
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return "", err
	}

	// Establish a connection to the Key Vault client
	client, err := azsecrets.NewClient(vaultURI, cred, nil)

	// Get a secret. An empty string version gets the latest version of the secret.
	version := ""
	resp, err := client.GetSecret(context.TODO(), secret, version, nil)
	if err != nil {
		return "", err
	}

	return *resp.Value, nil
}

func (s *AzureKeyVaultService) CreateSecret(secret, value string) error {
	vaultURI := os.Getenv("AZURE_KEY_VAULT_URI")

	// Create a credential using the NewDefaultAzureCredential type.
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return err
	}

	// Establish a connection to the Key Vault client
	client, err := azsecrets.NewClient(vaultURI, cred, nil)

	// Create a secret
	params := azsecrets.SetSecretParameters{Value: &value}
	_, err = client.SetSecret(context.TODO(), secret, params, nil)
	if err != nil {
		log.Fatalf("failed to create a secret: %v", err)
		return err
	}

	return nil
}

func (s *AzureKeyVaultService) DeleteSecret(secret string) error {
	vaultURI := os.Getenv("AZURE_KEY_VAULT_URI")

	// Create a credential using the NewDefaultAzureCredential type.
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return err
	}

	// Establish a connection to the Key Vault client
	client, err := azsecrets.NewClient(vaultURI, cred, nil)

	delResp, err := client.DeleteSecret(context.TODO(), secret, nil)
	if err != nil {
		log.Fatalf("failed to delete secret %s: %v", delResp.ID.Name(), err)
		return err
	}

	return nil
}
