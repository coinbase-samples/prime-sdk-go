package test

import (
	"context"
	"fmt"
	"testing"
	"time"
	"os"

	prime "github.com/coinbase-samples/prime-sdk-go"
)

func TestCreateWallet(t *testing.T) {

	if os.Getenv("PRIME_SDK_FULL_TESTS") != "enabled" {
		t.Skip()
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := newLiveTestClient()
	if err != nil {
		t.Fatal(err)
	}

	response, err := client.CreateWallet(
		ctx,
		&prime.CreateWalletRequest{
			PortfolioId: client.Credentials.PortfolioId,
			Name:        fmt.Sprintf("PrimeSdkTest-%d", time.Now().UnixMilli()),
			Type:        prime.WalletTypeVault,
			Symbol:      "BTC",
		})
	if err != nil {
		t.Fatal(err)
	}

	if response == nil {
		t.Fatal("expected a not nil response")
	}

	if len(response.ActivityId) == 0 {
		t.Fatal("expected an activity id")
	}

}


