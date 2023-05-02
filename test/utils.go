package test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	prime "github.com/coinbase-samples/prime-sdk-go"
)

func newLiveTestClient() (*prime.Client, error) {

	credentials, err := loadCredentialsFromEnv()
	if err != nil {
		return nil, err
	}

	client := prime.NewClient(credentials, http.Client{})
	return client, nil

}

func loadCredentialsFromEnv() (*prime.Credentials, error) {

	credentials := &prime.Credentials{}
	if err := json.Unmarshal([]byte(os.Getenv("PRIME_CREDENTIALS")), credentials); err != nil {
		return nil, fmt.Errorf("unable to deserialize prime credentials JSON: %w", err)
	}

	return credentials, nil
}
