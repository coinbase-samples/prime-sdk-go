package test

import (
	"context"
	"testing"
	"time"

	prime "github.com/coinbase-samples/prime-sdk-go"
)

func TestListWalletBalances(t *testing.T) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := newLiveTestClient()
	if err != nil {
		t.Fatal(err)
	}

	response, err := client.ListWalletBalances(ctx, &prime.ListWalletBalancesRequest{
		PortfolioId: client.Credentials.PortfolioId,
		Type:        prime.BalanceTypeTrading,
	})

	if err != nil {
		t.Fatal(err)
	}

	if response == nil {
		t.Fatal(err)
	}

	if len(response.Balances) == 0 {
		t.Fatal("expected balances in get")
	}

}
