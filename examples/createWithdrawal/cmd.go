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

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/coinbase-samples/prime-sdk-go/client"
	"github.com/coinbase-samples/prime-sdk-go/credentials"
	"github.com/coinbase-samples/prime-sdk-go/model"
	"github.com/coinbase-samples/prime-sdk-go/transactions"
	"github.com/google/uuid"
)

func main() {

	credentials, err := credentials.ReadEnvCredentials("PRIME_CREDENTIALS")
	if err != nil {
		log.Fatalf("unable to read credentials from environment: %v", err)
	}

	httpClient, err := client.DefaultHttpClient()
	if err != nil {
		log.Fatalf("unable to load default http client: %v", err)
	}

	client := client.NewRestClient(credentials, httpClient)

	if len(os.Args) < 4 {
		log.Fatalf("wallet ID, amount, symbol, and address are required as command-line arguments")
	}
	walletId := os.Args[1]
	amount := os.Args[2]
	symbol := os.Args[3]
	address := os.Args[4]

	transactionsSvc := transactions.NewTransactionsService(client)

	request := &transactions.CreateWalletWithdrawalRequest{
		PortfolioId:     credentials.PortfolioId,
		SourceWalletId:  walletId,
		DestinationType: "DESTINATION_BLOCKCHAIN",
		Amount:          amount,
		Symbol:          symbol,
		BlockchainAddress: &model.BlockchainAddress{
			Address: address,
		},
		IdempotencyKey: uuid.New().String(),
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	response, err := transactionsSvc.CreateWalletWithdrawal(ctx, request)
	if err != nil {
		log.Fatalf("unable to create withdrawal: %v", err)
	}

	output, err := json.MarshalIndent(response, "", "  ")
	if err != nil {
		panic(fmt.Sprintf("error marshaling response to JSON: %v", err))
	}
	fmt.Println(string(output))
}
