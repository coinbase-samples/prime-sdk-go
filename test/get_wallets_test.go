package test

import (
	"context"
	"testing"
	"time"

	prime "github.com/coinbase-samples/prime-sdk-go"
)

func TestGetWallets(t *testing.T) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := newLiveTestClient()
	if err != nil {
		t.Fatal(err)
	}

	walletsResponse, err := client.GetWallets(ctx, &prime.GetWalletsRequest{
		PortfolioId: client.Credentials.PortfolioId,
		Type:        prime.WalletTypeTrading,
	})

	if err != nil {
		t.Fatal(err)
	}

	if walletsResponse == nil {
		t.Fatal(err)
	}

	if len(walletsResponse.Wallets) == 0 {
		t.Fatal("expected trading wallets in get")
	}

	testGetWallet(t, client, walletsResponse.Wallets[0].Id)
}

func testGetWallet(t *testing.T, client *prime.Client, walletId string) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	response, err := client.GetWallet(ctx, &prime.GetWalletRequest{
		PortfolioId: client.Credentials.PortfolioId,
		Id:          walletId,
	})

	if err != nil {
		t.Fatal(err)
	}

	if response == nil {
		t.Fatal("expected wallet response to not be nil")
	}

	if response.Wallet == nil {
		t.Fatal("expected wallet to not be nil")
	}

	if response.Wallet.Id != walletId {
		t.Fatalf("expected wallet id: %s - received wallet id: %s", walletId, response.Wallet.Id)
	}

}
