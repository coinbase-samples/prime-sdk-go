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

	"github.com/coinbase-samples/prime-sdk-go/model"
	"github.com/coinbase-samples/prime-sdk-go/transactions"
	"github.com/coinbase-samples/prime-sdk-go/utils"
)

func TestCreateWithdrawal(t *testing.T) {

	if os.Getenv("PRIME_SDK_FULL_TESTS") != "enabled" {
		t.Skip()
	}

	c, err := newLiveTestClient()
	if err != nil {
		t.Fatal(err)
	}

	service := transactions.NewTransactionsService(c)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	response, err := service.CreateWalletWithdrawal(
		ctx,
		&transactions.CreateWalletWithdrawalRequest{
			PortfolioId:     c.Credentials().PortfolioId,
			SourceWalletId:  "12c89466-71a6-5e64-b849-d0b94fc8d40a",
			DestinationType: "DESTINATION_BLOCKCHAIN",
			Symbol:          "USDC",
			Amount:          "1.00",
			IdempotencyKey:  utils.NewUuid(),
			BlockchainAddress: &model.BlockchainAddress{
				Address: "0x836fa72D2aF55d698e8767acBE88c042b8201036",
				Network: &model.NetworkDetails{
					Id:   "base",
					Type: "mainnet",
				},
			},
		})

	if err != nil {
		t.Fatal(err)
	}

	if response == nil {
		t.Fatal("expected a not nil response")
	}

	if len(response.ActivityId) == 0 {
		t.Fatal("expected an activity id")
	}

}
