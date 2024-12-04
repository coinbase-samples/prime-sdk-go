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

package allocations

import (
	"context"

	"github.com/coinbase-samples/prime-sdk-go/client"
)

type AllocationsService interface {
	CreatePortfolioAllocations(ctx context.Context, request *CreatePortfolioAllocationsRequest) (*CreatePortfolioAllocationsResponse, error)
	CreatePortfolioNetAllocations(ctx context.Context, request *CreatePortfolioNetAllocationsRequest) (*CreatePortfolioNetAllocationsResponse, error)
	ListPortfolioAllocations(ctx context.Context, request *ListPortfolioAllocationsRequest) (*ListPortfolioAllocationsResponse, error)
	GetPortfolioAllocation(ctx context.Context, request *GetPortfolioAllocationRequest) (*GetPortfolioAllocationResponse, error)
	GetPortfolioNetAllocation(ctx context.Context, request *GetPortfolioNetAllocationRequest) (*GetPortfolioNetAllocationResponse, error)
}

func NewAllocationsService(c client.RestClient) AllocationsService {
	return &allocationsServiceImpl{client: c}
}

type allocationsServiceImpl struct {
	client client.RestClient
}
