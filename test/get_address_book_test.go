package test

import (
	"context"
	"testing"
	"time"

	prime "github.com/coinbase-samples/prime-sdk-go"
)

func TestGetAddressBook(t *testing.T) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := newLiveTestClient()
	if err != nil {
		t.Fatal(err)
	}

	response, err := client.GetAddressBook(ctx, &prime.GetAddressBookRequest{
		PortfolioId: client.Credentials.PortfolioId,
		Search:      "test",
	})

	if err != nil {
		t.Fatal(err)
	}

	if response == nil {
		t.Fatal(err)
	}

	if len(response.Addresses) == 0 {
		t.Fatal("expected addresses in get")
	}

	if len(response.Addresses[0].Id) == 0 {
		t.Fatal("expected address book entry id to be set")
	}

}
