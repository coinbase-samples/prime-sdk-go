/**
 * Copyright 2025-present Coinbase Global, Inc.
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
	"os"
	"testing"
	"time"

	"github.com/coinbase-samples/prime-sdk-go/orders"
	"github.com/coinbase-samples/prime-sdk-go/utils"
)

func TestQuote(t *testing.T) {

	if os.Getenv("PRIME_SDK_FULL_TESTS") != "enabled" {
		t.Skip()
	}

	c, err := newLiveTestClient()
	if err != nil {
		t.Fatal(err)
	}

	service := orders.NewOrdersService(c)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	response, err := service.CreateQuoteRequest(ctx, &orders.CreateQuoteRequest{
		PortfolioId:   c.Credentials().PortfolioId,
		ProductId:     "ADA-USD",
		Side:          "BUY",
		ClientQuoteId: utils.NewUuid(),
		BaseQuantity:  "15",
		LimitPrice:    "0.73",
	})

	if err != nil {
		t.Fatalf("cannot create quote request: %v", err)
	}

	if response == nil {
		t.Fatal("expected a not nil response")
	}

	//fmt.Println(response.QuoteId)

	acceptResponse, err := service.AcceptQuote(ctx, &orders.AcceptQuoteRequest{
		PortfolioId:   c.Credentials().PortfolioId,
		ProductId:     "ADA-USD",
		Side:          "BUY",
		ClientOrderId: utils.NewUuid(),
		QuoteId:       response.QuoteId,
	})

	if err != nil {
		t.Fatal(err)
	}

	//fmt.Println(acceptResponse.OrderId)
}
