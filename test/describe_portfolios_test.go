package test

import (
	"context"
	"log"
	"testing"
	"time"

	prime "github.com/coinbase-samples/prime-sdk-go"
)

func TestDescribePortfolios(t *testing.T) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := newLiveTestClient()
	if err != nil {
		t.Fatal(err)
	}

	response, err := client.DescribePortfolios(ctx, &prime.DescribePortfoliosRequest{})

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

	testDescribePortfolio(t, client, portfolio.Id)

	//testDescribePortfolioCredit(t, client, response.Portfolios[0].Id)
}

func testDescribePortfolio(t *testing.T, client *prime.Client, portfolioId string) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	response, err := client.DescribePortfolio(ctx, &prime.DescribePortfolioRequest{
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

func testDescribePortfolioCredit(t *testing.T, client *prime.Client, portfolioId string) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	response, err := client.DescribePortfolioCredit(ctx, &prime.DescribePortfolioCreditRequest{
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
