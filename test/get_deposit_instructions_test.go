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
	"github.com/coinbase-samples/prime-sdk-go/wallets"
)

func TestGetWalletDepositInstructions(t *testing.T) {

	c, err := newLiveTestClient()
	if err != nil {
		t.Fatal(err)
	}

	service := wallets.NewWalletsService(c)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	walletsResponse, err := service.ListWallets(ctx, &wallets.ListWalletsRequest{
		PortfolioId: c.Credentials().PortfolioId,
		Type:        model.WalletTypeVault,
	})

	if err != nil {
		t.Fatal(err)
	}

	if walletsResponse == nil {
		t.Fatal(err)
	}

	if len(walletsResponse.Wallets) == 0 {
		t.Fatal("expected crypto wallets in get")
	}

	testGetDepositInstructions(t, service, c.Credentials().PortfolioId, walletsResponse.Wallets[0].Id)
}

func testGetDepositInstructions(t *testing.T, svc wallets.WalletsService, portfolioId, walletId string) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	response, err := svc.GetWalletDepositInstructions(ctx, &wallets.GetWalletDepositInstructionsRequest{
		PortfolioId: portfolioId,
		Id:          walletId,
		Type:        model.WalletDepositTypeCrypto,
	})

	if err != nil {
		t.Fatal(err)
	}

	if response == nil {
		t.Fatal("expected wallet deposit response to not be nil")
	}

	if response.Crypto == nil {
		t.Fatal("expected crypto deposit instructions to not be nil")
	}

	if len(response.Crypto.Address) == 0 {
		t.Fatal("expected crypto deposit address to be set")
	}

}
