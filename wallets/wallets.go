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

package wallets

import (
	"context"

	"github.com/coinbase-samples/prime-sdk-go/client"
	"github.com/coinbase-samples/prime-sdk-go/model"
)

type WalletsService interface {
	ListWallets(ctx context.Context, request *ListWalletsRequest) (*ListWalletsResponse, error)
	CreateWallet(ctx context.Context, request *CreateWalletRequest) (*CreateWalletResponse, error)
	GetWallet(ctx context.Context, request *GetWalletRequest) (*GetWalletResponse, error)
	GetWalletDepositInstructions(ctx context.Context, request *GetWalletDepositInstructionsRequest) (*GetWalletDepositInstructionsResponse, error)
	ListWalletAddresses(ctx context.Context, request *ListWalletAddressesRequest) (*ListWalletAddressesResponse, error)
	CreateWalletAddress(ctx context.Context, request *CreateWalletAddressRequest) (*CreateWalletAddressResponse, error)
	PaginationConfig() *model.PaginationConfig
}

// NewWalletsService creates a new WalletsService with default pagination config
func NewWalletsService(c client.RestClient) WalletsService {
	return &walletsServiceImpl{
		client:           c,
		paginationConfig: model.DefaultPaginationConfig(),
	}
}

// NewWalletsServiceWithConfig creates a new WalletsService with custom pagination config
func NewWalletsServiceWithConfig(c client.RestClient, config *model.PaginationConfig) WalletsService {
	if config == nil {
		config = model.DefaultPaginationConfig()
	}
	return &walletsServiceImpl{
		client:           c,
		paginationConfig: config,
	}
}

type walletsServiceImpl struct {
	client           client.RestClient
	paginationConfig *model.PaginationConfig
}

func (s *walletsServiceImpl) PaginationConfig() *model.PaginationConfig {
	return s.paginationConfig
}
