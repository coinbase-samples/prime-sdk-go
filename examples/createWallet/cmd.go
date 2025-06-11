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
	"github.com/coinbase-samples/prime-sdk-go/wallets"
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

	if len(os.Args) < 3 {
		log.Fatalf("name, symbol, and wallet type are required as command-line arguments")
	}
	name := os.Args[1]
	symbol := os.Args[2]
	walletType := os.Args[3]

	walletsSvc := wallets.NewWalletsService(client)

	request := &wallets.CreateWalletRequest{
		PortfolioId:    credentials.PortfolioId,
		Name:           name,
		Symbol:         symbol,
		Type:           walletType,
		IdempotencyKey: uuid.New().String(),
	}

	response, err := walletsSvc.CreateWallet(context.Background(), request)
	if err != nil {
		log.Fatalf("unable to create vault wallet: %v", err)
	}

	output, err := json.MarshalIndent(response, "", "  ")
	if err != nil {
		panic(fmt.Sprintf("error marshaling response to JSON: %v", err))
	}
	fmt.Println(string(output))
}
