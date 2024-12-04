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

	"github.com/coinbase-samples/prime-sdk-go/balances"
	"github.com/coinbase-samples/prime-sdk-go/model"
	"github.com/coinbase-samples/prime-sdk-go/wallets"
)

func TestGetWalletBalance(t *testing.T) {

	c, err := newLiveTestClient()
	if err != nil {
		t.Fatal(err)
	}

	service := wallets.NewWalletsService(c)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	response, err := service.ListWallets(ctx, &wallets.ListWalletsRequest{
		PortfolioId: c.Credentials().PortfolioId,
		Type:        model.WalletTypeTrading,
	})

	if err != nil {
		t.Fatal(err)
	}

	if response == nil {
		t.Fatal(err)
	}

	if len(response.Wallets) == 0 {
		t.Fatal("expected wallets in get")
	}

	balancesService := balances.NewBalancesService(c)

	for _, w := range response.Wallets {
		testGetWalletBalance(t, balancesService, c.Credentials().PortfolioId, w.Id)
	}
}

func testGetWalletBalance(t *testing.T, svc balances.BalancesService, portfolioId, walletId string) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	response, err := svc.GetWalletBalance(ctx, &balances.GetWalletBalanceRequest{
		PortfolioId: portfolioId,
		Id:          walletId,
	})

	if err != nil {
		t.Fatal(err)
	}

	if response == nil {
		t.Fatal("expected wallet response to not be nil")
	}

	if response.Balance == nil {
		t.Fatal("expected wallet to not be nil")
	}

}
