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
	"log"
	"testing"
	"time"

	"github.com/coinbase-samples/prime-sdk-go/model"
	"github.com/coinbase-samples/prime-sdk-go/portfolios"
)

func TestListPortfolios(t *testing.T) {

	c, err := newLiveTestClient()
	if err != nil {
		t.Fatal(err)
	}

	service := portfolios.NewPortfoliosService(c)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	response, err := service.ListPortfolios(ctx, &portfolios.ListPortfoliosRequest{})

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

	var portfolio *model.Portfolio
	for _, v := range response.Portfolios {
		if v.Id == c.Credentials().PortfolioId {
			portfolio = v
			break
		}
	}

	if portfolio == nil {
		t.Fatal("expected get to include credentials portfolio")
	}

	testGetPortfolio(t, service, portfolio.Id)

}

func testGetPortfolio(t *testing.T, svc portfolios.PortfoliosService, portfolioId string) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	response, err := svc.GetPortfolio(ctx, &portfolios.GetPortfolioRequest{
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
