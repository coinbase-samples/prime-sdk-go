/**
 * Copyright 2024-present Coinbase Global, Inc.
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

package transactions

import (
	"context"

	"github.com/coinbase-samples/prime-sdk-go/client"
	"github.com/coinbase-samples/prime-sdk-go/model"
)

type TransactionsService interface {
	ListPortfolioTransactions(ctx context.Context, request *ListPortfolioTransactionsRequest) (*ListPortfolioTransactionsResponse, error)
	GetTransaction(ctx context.Context, request *GetTransactionRequest) (*GetTransactionResponse, error)
	CreateConversion(ctx context.Context, request *CreateConversionRequest) (*CreateConversionResponse, error)
	ListWalletTransactions(ctx context.Context, request *ListWalletTransactionsRequest) (*ListWalletTransactionsResponse, error)
	CreateWalletTransfer(ctx context.Context, request *CreateWalletTransferRequest) (*CreateWalletTransferResponse, error)
	CreateWalletWithdrawal(ctx context.Context, request *CreateWalletWithdrawalRequest) (*CreateWalletWithdrawalResponse, error)
	CreateOnchainTransaction(ctx context.Context, request *CreateOnchainTransactionRequest) (*CreateOnchainTransactionResposne, error)
	SubmitDepositTravelRuleData(ctx context.Context, request *SubmitDepositTravelRuleDataRequest) (*SubmitDepositTravelRuleDataResponse, error)
	ServiceConfig() *model.ServiceConfig
}

// NewTransactionsService creates a new TransactionsService with default pagination config
func NewTransactionsService(c client.RestClient) TransactionsService {
	return &transactionsServiceImpl{
		client:        c,
		serviceConfig: model.DefaultServiceConfig(),
	}
}

// NewTransactionsServiceWithConfig creates a new TransactionsService with custom pagination config
func NewTransactionsServiceWithConfig(c client.RestClient, config *model.ServiceConfig) TransactionsService {
	if config == nil {
		config = model.DefaultServiceConfig()
	}
	return &transactionsServiceImpl{
		client:        c,
		serviceConfig: config,
	}
}

type transactionsServiceImpl struct {
	client        client.RestClient
	serviceConfig *model.ServiceConfig
}

func (s *transactionsServiceImpl) ServiceConfig() *model.ServiceConfig {
	return s.serviceConfig
}
