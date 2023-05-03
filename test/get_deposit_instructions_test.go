package test

import (
	"context"
	"testing"
	"time"

	prime "github.com/coinbase-samples/prime-sdk-go"
)

func TestGetWalletDepositInstructions(t *testing.T) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := newLiveTestClient()
	if err != nil {
		t.Fatal(err)
	}

	walletsResponse, err := client.GetWallets(ctx, &prime.GetWalletsRequest{
		PortfolioId: client.Credentials.PortfolioId,
		Type:        prime.WalletTypeVault,
	})

	if err != nil {
		t.Fatal(err)
	}

	if walletsResponse == nil {
		t.Fatal(err)
	}

	if len(walletsResponse.Wallets) == 0 {
		t.Fatal("expected crypto wallets in describe")
	}

	testGetDepositInstructions(t, client, walletsResponse.Wallets[0].Id)
}

func testGetDepositInstructions(t *testing.T, client *prime.Client, walletId string) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	response, err := client.GetWalletDepositInstructions(ctx, &prime.GetWalletDepositInstructionsRequest{
		PortfolioId: client.Credentials.PortfolioId,
		Id:          walletId,
		Type:        prime.WalletDepositTypeCrypto,
	})

	if err != nil {
		t.Fatal(err)
	}

	if response == nil {
		t.Fatal("expected wallet deposit response to not be nil")
	}

	if response.Crypto == nil {
		t.Fatal("expected crypto deposit instructions to not be nil")
	}

	if len(response.Crypto.Address) == 0 {
		t.Fatal("expected crypto deposit address to be set")
	}

}
