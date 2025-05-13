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

	"github.com/coinbase-samples/core-go"
	"github.com/coinbase-samples/prime-sdk-go/client"
	"github.com/coinbase-samples/prime-sdk-go/model"
)

type CreatePortfolioNetAllocationsRequest struct {
	NettingId                       string                 `json:"netting_id"`
	SourcePortfolioId               string                 `json:"source_portfolio_id"`
	ProductId                       string                 `json:"product_id"`
	OrderIds                        []string               `json:"order_ids"`
	AllocationLegs                  []*model.AllocationLeg `json:"allocation_legs"`
	SizeType                        string                 `json:"size_type"`
	RemainderDestinationPortfolioId string                 `json:"remainder_destination_portfolio"`
}

type CreatePortfolioNetAllocationsResponse struct {
	Success          bool                                  `json:"success"`
	NettingId        string                                `json:"netting_id"`
	FailureReason    string                                `json:"failure_reason"`
	BuyAllocationId  string                                `json:"buy_allocation_id"`
	SellAllocationId string                                `json:"sell_allocation_id"`
	Request          *CreatePortfolioNetAllocationsRequest `json:"-"`
}

func (s *allocationsServiceImpl) CreatePortfolioNetAllocations(
	ctx context.Context,
	request *CreatePortfolioNetAllocationsRequest,
) (*CreatePortfolioNetAllocationsResponse, error) {

	path := "/allocations/net"

	response := &CreatePortfolioNetAllocationsResponse{Request: request}

	if err := core.HttpPost(
		ctx,
		s.client,
		path,
		core.EmptyQueryParams,
		client.DefaultSuccessHttpStatusCodes,
		request,
		response,
		s.client.HeadersFunc(),
	); err != nil {
		return nil, err
	}

	return response, nil
}
