package test

import (
	"context"
	"log"
	"testing"
	"time"

	prime "github.com/coinbase-samples/prime-sdk-go"
)

func TestListPortfolios(t *testing.T) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := newLiveTestClient()
	if err != nil {
		t.Fatal(err)
	}

	response, err := client.ListPortfolios(ctx, &prime.ListPortfoliosRequest{})

	if err != nil {
		t.Fatal(err)
	}

	if response == nil {
		t.Fatal(err)
	}

	if len(response.Portfolios) == 0 {
		t.Fatal("expected portfolios in get")
	}

	if len(response.Portfolios[0].Id) == 0 {
		t.Fatal("expected portfoliio id to be set")
	}

	var portfolio *prime.Portfolio
	for _, v := range response.Portfolios {
		if v.Id == client.Credentials.PortfolioId {
			portfolio = v
			break
		}
	}

	if portfolio == nil {
		t.Fatal("expected get to include credentials portfolio")
	}

	testGetPortfolio(t, client, portfolio.Id)

	//testGetCredit(t, client, response.Portfolios[0].Id)
}

func testGetPortfolio(t *testing.T, client *prime.Client, portfolioId string) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	response, err := client.GetPortfolio(ctx, &prime.GetPortfolioRequest{
		PortfolioId: portfolioId,
	})

	if err != nil {
		t.Fatal(err)
	}

	if response == nil {
		t.Fatal("expected portfolio response to not be nil")
	}

	if response.Portfolio == nil {
		t.Fatal("expected portfolio to not be nil")
	}

	if response.Portfolio.Id != portfolioId {
		t.Fatalf("expected portfolio id: %s - received portfolio id: %s", portfolioId, response.Portfolio.Id)
	}

	if len(response.Portfolio.EntityId) == 0 {
		log.Fatal("expected entity id to be set")
	}

}

func testGetCredit(t *testing.T, client *prime.Client, portfolioId string) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	response, err := client.GetCredit(ctx, &prime.GetCreditRequest{
		Id: portfolioId,
	})

	if err != nil {
		t.Fatal(err)
	}

	if response == nil {
		t.Fatal("expected portfolio credit response to not be nil")
	}

	if response.PostTradeCredit == nil {
		t.Fatal("expected portfolio post trade credit to not be nil")
	}

	if response.PostTradeCredit.Id != portfolioId {
		t.Fatalf("expected portfolio id: %s - received portfolio id: %s", portfolioId, response.PostTradeCredit.Id)
	}

}
