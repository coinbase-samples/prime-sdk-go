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

package advancedtransfers

import (
	"context"

	"github.com/coinbase-samples/prime-sdk-go/client"
	"github.com/coinbase-samples/prime-sdk-go/model"
)

type AdvancedTransfersService interface {
	ListAdvancedTransfers(ctx context.Context, request *ListAdvancedTransfersRequest) (*ListAdvancedTransfersResponse, error)
	CreateAdvancedTransfer(ctx context.Context, request *CreateAdvancedTransferRequest) (*CreateAdvancedTransferResponse, error)
	CancelAdvancedTransfer(ctx context.Context, request *CancelAdvancedTransferRequest) (*CancelAdvancedTransferResponse, error)
	ListAdvancedTransferTransactions(ctx context.Context, request *ListAdvancedTransferTransactionsRequest) (*ListAdvancedTransferTransactionsResponse, error)
	ServiceConfig() *model.ServiceConfig
}

// NewAdvancedTransfersService creates a new AdvancedTransfersService with default pagination config.
func NewAdvancedTransfersService(c client.RestClient) AdvancedTransfersService {
	return &advancedTransfersServiceImpl{
		client:        c,
		serviceConfig: model.DefaultServiceConfig(),
	}
}

// NewAdvancedTransfersServiceWithConfig creates a new AdvancedTransfersService with custom pagination config.
func NewAdvancedTransfersServiceWithConfig(c client.RestClient, config *model.ServiceConfig) AdvancedTransfersService {
	if config == nil {
		config = model.DefaultServiceConfig()
	}
	return &advancedTransfersServiceImpl{
		client:        c,
		serviceConfig: config,
	}
}

type advancedTransfersServiceImpl struct {
	client        client.RestClient
	serviceConfig *model.ServiceConfig
}

func (s *advancedTransfersServiceImpl) ServiceConfig() *model.ServiceConfig {
	return s.serviceConfig
}
