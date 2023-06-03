package secman

import (
	"log"

	vault "github.com/hashicorp/vault/api"
)

func GetListSecrets(path string) (map[string]interface{}, error) {
	config := vault.DefaultConfig()

	config.Address = "http://45.8.99.135:8200/"

	client, err := vault.NewClient(config)
	if err != nil {
		log.Fatalf("unable to initialize Vault client: %v", err)
	}

	// Authenticate
	client.SetToken("hvs.qIxmTShSDDHOnWiWjFUL0rHF")
	// secret, err := client.KVv1("secret").Get(context.Background(), path)
	// if err != nil {
	// 	log.Fatalf("unable to read secret: %v", err)
	// }

	// value, ok := secret.Data["password"].(string)
	// if !ok {
	// 	log.Fatalf("value type assertion failed: %T %#v", secret.Data["password"], secret.Data["password"])
	// }
	secret, err := client.Logical().List(path)
	if err != nil {
		log.Fatalf("unable to write secret: %v", err)
	}
	if secret == nil {
		return nil, nil
	}
	secretData := secret.Data
	if len(secretData) == 0 {
		return secretData, nil
	}
	return secretData, nil
}
