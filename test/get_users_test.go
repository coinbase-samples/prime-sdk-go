package test

import (
	"context"
	"fmt"
	"testing"
	"time"

	prime "github.com/coinbase-samples/prime-sdk-go"
)

func TestGetPortfolioUsers(t *testing.T) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := newLiveTestClient()
	if err != nil {
		t.Fatal(err)
	}

	response, err := client.GetPortfolioUsers(ctx, &prime.GetPortfolioUsersRequest{
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
		t.Fatal("expected users in describe")
	}

	if len(response.Users[0].Id) == 0 {
		t.Fatal("expected user id to be set")
	}

	if len(response.Users[0].PortfolioId) == 0 {
		t.Fatal("expected user portfolio id to be set")
	}

	for _, u := range response.Users {

		if u.Id == "367345e1-b77c-5441-8fcc-b68f0bb1cd6e" {

			fmt.Println(fmt.Sprintf("name: %s - email: %s", u.Name, u.Email))
		}

	}

}
