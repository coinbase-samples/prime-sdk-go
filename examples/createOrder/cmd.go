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
	"github.com/coinbase-samples/prime-sdk-go/model"
	"github.com/coinbase-samples/prime-sdk-go/orders"
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

	if len(os.Args) < 6 {
		log.Fatalf("product ID, side, type, base quantity, and limit price are required as command-line arguments")
	}
	productId := os.Args[1]
	side := os.Args[2]
	orderType := os.Args[3]
	baseQuantity := os.Args[4]
	limitPrice := os.Args[5]

	ordersSvc := orders.NewOrdersService(client)

	request := &orders.CreateOrderRequest{
		Order: &model.Order{
			PortfolioId:  credentials.PortfolioId,
			ProductId:    productId,
			Side:         side,
			Type:         orderType,
			BaseQuantity: baseQuantity,
			LimitPrice:   limitPrice,
		},
	}

	response, err := ordersSvc.CreateOrder(context.Background(), request)
	if err != nil {
		log.Fatalf("unable to create order: %v", err)
	}

	output, err := json.MarshalIndent(response, "", "  ")
	if err != nil {
		log.Fatalf("error marshaling response to JSON: %v", err)
	}
	fmt.Println(string(output))
}
