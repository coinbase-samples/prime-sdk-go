/**
 * Copyright 2023-present Coinbase Global, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *  http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package test

import (
	"context"
	"testing"
	"time"

	"github.com/coinbase-samples/prime-sdk-go/model"
	"github.com/coinbase-samples/prime-sdk-go/orders"
)

func TestOrders(t *testing.T) {

	c, err := newLiveTestClient()
	if err != nil {
		t.Fatal(err)
	}

	testProductId := "ADA-USD"

	service := orders.NewOrdersService(c)

	order := &model.Order{
		PortfolioId:   c.Credentials().PortfolioId,
		Side:          "BUY",
		ClientOrderId: time.Now().String(),
		ProductId:     "ADA-USD",
		BaseQuantity:  "10",
		Type:          model.OrderTypeLimit,
		LimitPrice:    "0.15",
		TimeInForce:   model.TimeInForceGoodUntilCancelled,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	response, err := service.CreateOrderPreview(
		ctx,
		&orders.CreateOrderRequest{Order: order},
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

	orderId := testCreateOrder(t, service, order)

	testGetOrder(t, service, c.Credentials().PortfolioId, orderId)

	testListOpenOrders(t, service, c.Credentials().PortfolioId, testProductId, orderId)

	testCancelOrder(t, service, c.Credentials().PortfolioId, orderId)
}

func testListOpenOrders(t *testing.T, svc orders.OrdersService, portfolioId, productId, orderId string) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	response, err := svc.ListOpenOrders(
		ctx,
		&orders.ListOpenOrdersRequest{
			PortfolioId: portfolioId,
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

func testCreateOrder(t *testing.T, svc orders.OrdersService, order *model.Order) string {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	response, err := svc.CreateOrder(
		ctx,
		&orders.CreateOrderRequest{
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

func testGetOrder(t *testing.T, svc orders.OrdersService, portfolioId, orderId string) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var err error
	for idx := 0; idx < 3; idx++ {

		var response *orders.GetOrderResponse

		response, err = svc.GetOrder(
			ctx,
			&orders.GetOrderRequest{
				PortfolioId: portfolioId,
				OrderId:     orderId,
			},
		)

		if err != nil {

			time.Sleep(time.Duration(idx) * time.Second)
			continue
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
		break
	}

	if err != nil {
		t.Fatal(err)
	}
}

func testCancelOrder(t *testing.T, svc orders.OrdersService, portfolioId, orderId string) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	response, err := svc.CancelOrder(
		ctx,
		&orders.CancelOrderRequest{
			PortfolioId: portfolioId,
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
