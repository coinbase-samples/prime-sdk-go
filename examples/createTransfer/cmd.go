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

	"github.com/coinbase-samples/prime-sdk-go/client"
	"github.com/coinbase-samples/prime-sdk-go/credentials"
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
		log.Fatalf("source wallet ID, destination wallet ID, amount, and symbol are required as command-line arguments")
	}
	sourceWalletId := os.Args[1]
	destinationWalletId := os.Args[2]
	amount := os.Args[3]
	symbol := os.Args[4]

	transactionsSvc := transactions.NewTransactionsService(client)

	request := &transactions.CreateWalletTransferRequest{
		PortfolioId:         credentials.PortfolioId,
		SourceWalletId:      sourceWalletId,
		DestinationWalletId: destinationWalletId,
		Amount:              amount,
		IdempotencyKey:      uuid.New().String(),
		Symbol:              symbol,
	}

	response, err := transactionsSvc.CreateWalletTransfer(context.Background(), request)
	if err != nil {
		log.Fatalf("unable to create transfer: %v", err)
	}

	output, err := json.MarshalIndent(response, "", "  ")
	if err != nil {
		log.Fatalf("error marshaling response to JSON: %v", err)
	}
	fmt.Println(string(output))
}
