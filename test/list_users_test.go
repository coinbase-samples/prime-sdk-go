package test

import (
	"context"
	"testing"
	"time"

	prime "github.com/coinbase-samples/prime-sdk-go"
)

func TestListPortfolioUsers(t *testing.T) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := newLiveTestClient()
	if err != nil {
		t.Fatal(err)
	}

	response, err := client.ListPortfolioUsers(ctx, &prime.ListPortfolioUsersRequest{
		PortfolioId: client.Credentials.PortfolioId,
		Pagination:  &prime.PaginationParams{Limit: "100"},
	})

	if err != nil {
		t.Fatal(err)
	}

	if response == nil {
		t.Fatal(err)
	}

	if len(response.Users) == 0 {
		t.Fatal("expected users in get")
	}

	if len(response.Users[0].Id) == 0 {
		t.Fatal("expected user id to be set")
	}

	if len(response.Users[0].PortfolioId) == 0 {
		t.Fatal("expected user portfolio id to be set")
	}

}
