package test

import (
	"context"
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

	testDescribePortfolio(t, client, response.Portfolios[0].Id)

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
