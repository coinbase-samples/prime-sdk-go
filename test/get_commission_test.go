package test

import (
	"context"
	"testing"
	"time"

	prime "github.com/coinbase-samples/prime-sdk-go"
)

func TestPortfolioCommission(t *testing.T) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := newLiveTestClient()
	if err != nil {
		t.Fatal(err)
	}

	response, err := client.GetPortfolioCommission(ctx, &prime.GetPortfolioCommissionRequest{
		PortfolioId: client.Credentials.PortfolioId,
	})

	if err != nil {
		t.Fatal(err)
	}

	if response == nil {
		t.Fatal(err)
	}

	if response.Commission == nil {
		t.Fatal("expected commision in get")
	}

}
