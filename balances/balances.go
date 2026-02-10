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

package balances

import (
	"context"

	"github.com/coinbase-samples/prime-sdk-go/client"
	"github.com/coinbase-samples/prime-sdk-go/model"
)

type BalancesService interface {
	ListPortfolioBalances(ctx context.Context, request *ListPortfolioBalancesRequest) (*ListPortfolioBalancesResponse, error)
	GetWalletBalance(ctx context.Context, request *GetWalletBalanceRequest) (*GetWalletBalanceResponse, error)
	ListOnchainWalletBalances(ctx context.Context, request *ListOnchainWalletBalancesRequest) (*ListOnchainWalletBalancesResponse, error)
	ListEntityBalances(ctx context.Context, request *ListEntityBalancesRequest) (*ListEntityBalancesResponse, error)
	ServiceConfig() *model.ServiceConfig
}

// NewBalancesService creates a new BalancesService with default pagination config
func NewBalancesService(c client.RestClient) BalancesService {
	return &balancesServiceImpl{
		client:        c,
		serviceConfig: model.DefaultServiceConfig(),
	}
}

// NewBalancesServiceWithConfig creates a new BalancesService with custom pagination config
func NewBalancesServiceWithConfig(c client.RestClient, config *model.ServiceConfig) BalancesService {
	if config == nil {
		config = model.DefaultServiceConfig()
	}
	return &balancesServiceImpl{
		client:        c,
		serviceConfig: config,
	}
}

type balancesServiceImpl struct {
	client        client.RestClient
	serviceConfig *model.ServiceConfig
}

func (s *balancesServiceImpl) ServiceConfig() *model.ServiceConfig {
	return s.serviceConfig
}
