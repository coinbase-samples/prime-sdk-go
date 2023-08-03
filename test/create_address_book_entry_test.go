package test

import (
	"context"
	"fmt"
	"os"
	"testing"
	"time"

	prime "github.com/coinbase-samples/prime-sdk-go"
)

func TestCreateAddressBookEntry(t *testing.T) {

	if os.Getenv("PRIME_SDK_FULL_TESTS") != "enabled" {
		t.Skip()
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := newLiveTestClient()
	if err != nil {
		t.Fatal(err)
	}

	response, err := client.CreateAddressBookEntry(
		ctx,
		&prime.CreateAddressBookEntryRequest{
			PortfolioId: client.Credentials.PortfolioId,
			Name:        fmt.Sprintf("PrimeSdkTest-%d", time.Now().UnixMilli()),
			Symbol:      "ETH",
			Address:     "0x836fa72D2aF55d698e8767acBE88c042b8201036",
		},
	)
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
