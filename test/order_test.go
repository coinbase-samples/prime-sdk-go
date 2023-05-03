package test

import (
	"context"
	"testing"
	"time"

	prime "github.com/coinbase-samples/prime-sdk-go"
)

func TestOrders(t *testing.T) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := newLiveTestClient()
	if err != nil {
		t.Fatal(err)
	}

	testProductId := "ADA-USD"

	order := &prime.Order{
		PortfolioId:   client.Credentials.PortfolioId,
		Side:          prime.OrderSideBuy,
		ClientOrderId: time.Now().String(),
		ProductId:     "ADA-USD",
		BaseQuantity:  "20",
		Type:          prime.OrderTypeLimit,
		LimitPrice:    "0.15",
		TimeInForce:   prime.TimeInForceGoodUntilCancelled,
	}

	response, err := client.CreateOrderPreview(
		ctx,
		&prime.CreateOrderRequest{Order: order},
	)
	if err != nil {
		t.Fatal(err)
	}

	if response == nil {
		t.Fatal("expected create order preview response to not be nil")
	}

	if response.Order == nil {
		t.Fatal("expected create order preview to not be nil")
	}

	if len(response.Order.Total) == 0 {
		t.Error("expected an order total in the response")
	}

	orderId := testCreateOrder(t, client, order)

	testDescribeOrder(t, client, orderId)

	testDescribeOpenOrders(t, client, testProductId, orderId)

	testCancelOrder(t, client, orderId)
}

func testDescribeOpenOrders(t *testing.T, client *prime.Client, productId, orderId string) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	response, err := client.DescribeOpenOrders(
		ctx,
		&prime.DescribeOpenOrdersRequest{
			PortfolioId: client.Credentials.PortfolioId,
			ProductId:   productId,
		},
	)
	if err != nil {
		t.Fatal(err)
	}

	if len(response.Orders) == 0 {
		t.Error("expected open orders to have > 0")
	}

	var found bool

	for _, o := range response.Orders {

		if o.Id == orderId {
			found = true
			break
		}
	}

	if !found {
		t.Error("expected to find an existing open order")
	}

}

func testCreateOrder(t *testing.T, client *prime.Client, order *prime.Order) string {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	response, err := client.CreateOrder(
		ctx,
		&prime.CreateOrderRequest{
			Order: order,
		},
	)
	if err != nil {
		t.Fatal(err)
	}

	if response == nil {
		t.Fatal("expected create order response to not be nil")
	}

	if len(response.OrderId) == 0 {
		t.Fatal("expected an order id in the response")
	}

	return response.OrderId
}

func testDescribeOrder(t *testing.T, client *prime.Client, orderId string) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	response, err := client.DescribeOrder(
		ctx,
		&prime.DescribeOrderRequest{
			PortfolioId: client.Credentials.PortfolioId,
			OrderId:     orderId,
		},
	)
	if err != nil {
		t.Fatal(err)
	}

	if response.Order == nil {
		t.Fatal("expected order to be set in response")
	}

	if len(response.Order.Id) == 0 {
		t.Error("expected order id to be set in response")
	}

	if response.Order.Id != orderId {
		t.Error("expected order id in response to match passed order id")
	}
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
