package test

import (
	"context"
	"testing"
	"time"

	prime "github.com/coinbase-samples/prime-sdk-go"
)

func TestCreateOrder(t *testing.T) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := newLiveTestClient()
	if err != nil {
		t.Fatal(err)
	}

	response, err := client.CreateOrder(
		ctx,
		&prime.CreateOrderRequest{
			PortfolioId:   client.Credentials.PortfolioId,
			Side:          prime.OrderSideBuy,
			ClientOrderId: time.Now().String(),
			ProductId:     "ADA-USD",
			BaseQuantity:  "20",
			Type:          prime.OrderTypeLimit,
			LimitPrice:    "0.15",
			TimeInForce:   prime.TimeInForceGoodUntilCancelled,
		},
	)
	if err != nil {
		t.Fatal(err)
	}

	if response == nil {
		t.Fatal("expected response to not be nil")
	}

	if len(response.OrderId) == 0 {
		t.Fatal("expected an order id in the response")
	}

	testCancelOrder(t, client, response.OrderId)
}

func testCancelOrder(t *testing.T, client *prime.Client, orderId string) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	response, err := client.CancelOrder(
		ctx,
		&prime.CancelOrderRequest{
			PortfolioId: client.Credentials.PortfolioId,
			OrderId:     orderId,
		},
	)
	if err != nil {
		t.Fatal(err)
	}

	if len(response.OrderId) == 0 {
		t.Error("expected order id to be set in cancel order")
	}

	if response.OrderId != orderId {
		t.Error("expected order id in response to match passed order id")
	}

}
