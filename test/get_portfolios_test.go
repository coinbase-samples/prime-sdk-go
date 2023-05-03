package test

import (
	"context"
	"log"
	"testing"
	"time"

	prime "github.com/coinbase-samples/prime-sdk-go"
)

func TestGetPortfolios(t *testing.T) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := newLiveTestClient()
	if err != nil {
		t.Fatal(err)
	}

	response, err := client.GetPortfolios(ctx, &prime.GetPortfoliosRequest{})

	if err != nil {
		t.Fatal(err)
	}

	if response == nil {
		t.Fatal(err)
	}

	if len(response.Portfolios) == 0 {
		t.Fatal("expected portfolios in describe")
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
		t.Fatal("expected describe to include credentials portfolio")
	}

	testGetPortfolio(t, client, portfolio.Id)

	//testGetPortfolioCredit(t, client, response.Portfolios[0].Id)
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

func testGetPortfolioCredit(t *testing.T, client *prime.Client, portfolioId string) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	response, err := client.GetPortfolioCredit(ctx, &prime.GetPortfolioCreditRequest{
		PortfolioId: portfolioId,
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