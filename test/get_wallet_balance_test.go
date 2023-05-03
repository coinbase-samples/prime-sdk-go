package test

import (
	"context"
	"testing"
	"time"

	prime "github.com/coinbase-samples/prime-sdk-go"
)

func TestGetWalletBalance(t *testing.T) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := newLiveTestClient()
	if err != nil {
		t.Fatal(err)
	}

	response, err := client.GetWallets(ctx, &prime.GetWalletsRequest{
		PortfolioId: client.Credentials.PortfolioId,
		Type:        prime.WalletTypeTrading,
	})

	if err != nil {
		t.Fatal(err)
	}

	if response == nil {
		t.Fatal(err)
	}

	if len(response.Wallets) == 0 {
		t.Fatal("expected wallets in describe")
	}

	for _, w := range response.Wallets {

		testGetWalletBalance(t, client, w.Id)

	}
}

func testGetWalletBalance(t *testing.T, client *prime.Client, walletId string) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	response, err := client.GetWalletBalance(ctx, &prime.GetWalletBalanceRequest{
		PortfolioId: client.Credentials.PortfolioId,
		Id:          walletId,
	})

	if err != nil {
		t.Fatal(err)
	}

	if response == nil {
		t.Fatal("expected wallet response to not be nil")
	}

	if response.Balance == nil {
		t.Fatal("expected wallet to not be nil")
	}

}
