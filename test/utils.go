package test

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	prime "github.com/coinbase-samples/prime-sdk-go"
)

func loadEntityId(client *prime.Client) (string, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	response, err := client.DescribePortfolio(
		ctx,
		&prime.DescribePortfolioRequest{
			PortfolioId: client.Credentials.PortfolioId,
		},
	)

	if err != nil {
		return "", err
	}

	return response.Portfolio.EntityId, nil
}

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
