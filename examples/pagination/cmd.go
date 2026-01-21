/**
 * Copyright 2026-present Coinbase Global, Inc.
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
	"fmt"
	"log"

	"github.com/coinbase-samples/prime-sdk-go/client"
	"github.com/coinbase-samples/prime-sdk-go/credentials"
	"github.com/coinbase-samples/prime-sdk-go/model"
	"github.com/coinbase-samples/prime-sdk-go/transactions"
	"github.com/coinbase-samples/prime-sdk-go/wallets"
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

	restClient := client.NewRestClient(credentials, httpClient)
	ctx := context.Background()

	// Example 1: Fetch the second page manually
	fetchSecondPage(ctx, restClient, credentials.PortfolioId)

	// Example 2: FetchAll with MaxItems of 500
	fetchAllWithMaxItems(ctx, restClient, credentials.PortfolioId)

	// Example 3: FetchAll with MaxPages of 5
	fetchAllWithMaxPages(ctx, restClient, credentials.PortfolioId)

	// Example 4: Using DefaultLimit of 100
	fetchWithDefaultLimit(ctx, restClient, credentials.PortfolioId)
}

// fetchSecondPage demonstrates manually fetching the second page of results
func fetchSecondPage(ctx context.Context, restClient client.RestClient, portfolioId string) {
	fmt.Println("\n=== Example 1: Fetch Second Page Manually ===")

	walletsSvc := wallets.NewWalletsService(restClient)

	// Fetch the first page
	request := &wallets.ListWalletsRequest{
		PortfolioId: portfolioId,
		Pagination: &model.PaginationParams{
			Limit: 10, // 10 items per page
		},
	}

	firstPage, err := walletsSvc.ListWallets(ctx, request)
	if err != nil {
		log.Fatalf("error fetching first page: %v", err)
	}

	fmt.Printf("First page: %d wallets\n", len(firstPage.Wallets))
	fmt.Printf("Has next page: %v\n", firstPage.HasNext())
	fmt.Printf("Next cursor: %s\n", firstPage.GetNextCursor())

	// Fetch the second page if available
	if firstPage.HasNext() {
		secondPage, err := firstPage.Next(ctx)
		if err != nil {
			log.Fatalf("error fetching second page: %v", err)
		}
		if secondPage != nil {
			fmt.Printf("Second page: %d wallets\n", len(secondPage.Wallets))
			fmt.Printf("Has next page: %v\n", secondPage.HasNext())
		}
	} else {
		fmt.Println("No second page available")
	}
}

// fetchAllWithMaxItems demonstrates using FetchAll with a MaxItems limit
func fetchAllWithMaxItems(ctx context.Context, restClient client.RestClient, portfolioId string) {
	fmt.Println("\n=== Example 2: FetchAll with MaxItems of 500 ===")

	// Create service with MaxItems config
	config := &model.PaginationConfig{
		MaxItems:     500, // Stop after collecting 500 items
		DefaultLimit: 100, // Request 100 items per page
	}

	txnSvc := transactions.NewTransactionsServiceWithConfig(restClient, config)

	request := &transactions.ListPortfolioTransactionsRequest{
		PortfolioId: portfolioId,
	}

	resp, err := txnSvc.ListPortfolioTransactions(ctx, request)
	if err != nil {
		log.Fatalf("error fetching transactions: %v", err)
	}

	fmt.Printf("First page: %d transactions\n", len(resp.Transactions))

	// FetchAll will respect MaxItems from the service config
	allTransactions, err := resp.Iterator().FetchAll(ctx)
	if err != nil {
		log.Fatalf("error fetching all transactions: %v", err)
	}

	fmt.Printf("Total transactions fetched (max 500): %d\n", len(allTransactions))

	// Alternative: Override config per-iterator
	fmt.Println("\n--- Alternative: Override config for specific iterator ---")

	resp2, err := txnSvc.ListPortfolioTransactions(ctx, request)
	if err != nil {
		log.Fatalf("error fetching transactions: %v", err)
	}

	customConfig := &model.PaginationConfig{MaxItems: 200}
	limitedTransactions, err := resp2.Iterator().WithConfig(customConfig).FetchAll(ctx)
	if err != nil {
		log.Fatalf("error fetching limited transactions: %v", err)
	}

	fmt.Printf("Total transactions with custom limit (max 200): %d\n", len(limitedTransactions))
}

// fetchAllWithMaxPages demonstrates using FetchAll with a MaxPages limit
func fetchAllWithMaxPages(ctx context.Context, restClient client.RestClient, portfolioId string) {
	fmt.Println("\n=== Example 3: FetchAll with MaxPages of 5 ===")

	// Create service with MaxPages config
	config := &model.PaginationConfig{
		MaxPages:     5,  // Stop after 5 pages
		DefaultLimit: 25, // Request 25 items per page
	}

	walletsSvc := wallets.NewWalletsServiceWithConfig(restClient, config)

	request := &wallets.ListWalletsRequest{
		PortfolioId: portfolioId,
	}

	resp, err := walletsSvc.ListWallets(ctx, request)
	if err != nil {
		log.Fatalf("error fetching wallets: %v", err)
	}

	fmt.Printf("First page: %d wallets\n", len(resp.Wallets))

	// FetchAll will respect MaxPages from the service config
	allWallets, err := resp.Iterator().FetchAll(ctx)
	if err != nil {
		log.Fatalf("error fetching all wallets: %v", err)
	}

	fmt.Printf("Total wallets fetched (max 5 pages): %d\n", len(allWallets))

	// Alternative: Using ForEach to process page by page
	fmt.Println("\n--- Alternative: Process pages with ForEach ---")

	resp2, err := walletsSvc.ListWallets(ctx, request)
	if err != nil {
		log.Fatalf("error fetching wallets: %v", err)
	}

	pageCount := 0
	err = resp2.Iterator().ForEach(ctx, func(page *wallets.ListWalletsResponse) error {
		pageCount++
		fmt.Printf("Processing page %d: %d wallets\n", pageCount, len(page.Wallets))
		return nil
	})
	if err != nil {
		log.Fatalf("error in ForEach: %v", err)
	}

	fmt.Printf("Processed %d pages total\n", pageCount)
}

// fetchWithDefaultLimit demonstrates using DefaultLimit to set page size at service level
func fetchWithDefaultLimit(ctx context.Context, restClient client.RestClient, portfolioId string) {
	fmt.Println("\n=== Example 4: Using DefaultLimit of 100 ===")

	// Create service with DefaultLimit config
	// This means all requests will use 100 items per page unless overridden
	config := &model.PaginationConfig{
		DefaultLimit: 100, // All requests will default to 100 items per page
	}

	txnSvc := transactions.NewTransactionsServiceWithConfig(restClient, config)

	// Request without specifying pagination - DefaultLimit will be applied
	request := &transactions.ListPortfolioTransactionsRequest{
		PortfolioId: portfolioId,
		// No Pagination specified - DefaultLimit of 100 will be used
	}

	resp, err := txnSvc.ListPortfolioTransactions(ctx, request)
	if err != nil {
		log.Fatalf("error fetching transactions: %v", err)
	}

	fmt.Printf("First page (using DefaultLimit 100): %d transactions\n", len(resp.Transactions))
	fmt.Printf("Has next page: %v\n", resp.HasNext())

	// You can still override the limit per-request if needed
	fmt.Println("\n--- Override DefaultLimit per-request ---")

	requestWithCustomLimit := &transactions.ListPortfolioTransactionsRequest{
		PortfolioId: portfolioId,
		Pagination: &model.PaginationParams{
			Limit: 50, // Override to 50 items per page
		},
	}

	resp2, err := txnSvc.ListPortfolioTransactions(ctx, requestWithCustomLimit)
	if err != nil {
		log.Fatalf("error fetching transactions: %v", err)
	}

	fmt.Printf("First page (override to 50): %d transactions\n", len(resp2.Transactions))
}
