/**
 * Copyright 2023-present Coinbase Global, Inc.
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

package prime

import (
	"context"
)

type CreatePortfolioAllocationsRequest struct {
	AllocationId                    string           `json:"allocation_id"`
	SourcePortfolioId               string           `json:"source_portfolio_id"`
	ProductId                       string           `json:"product_id"`
	OrderIds                        []string         `json:"order_ids"`
	AllocationLegs                  []*AllocationLeg `json:"allocation_legs"`
	SizeType                        string           `json:"size_type"`
	RemainderDestinationPortfolioId string           `json:"remainder_destination_portfolio"`
}

type AllocationLeg struct {
	LegId                  string `json:"allocation_leg_id"`
	DestinationPortfolioId string `json:"destination_portfolio_id"`
	Amount                 string `json:"amount"`
}

type CreatePortfolioAllocationsResponse struct {
	Success       bool                               `json:"success"`
	AllocationId  string                             `json:"allocation_id"`
	FailureReason string                             `json:"failure_reason"`
	Request       *CreatePortfolioAllocationsRequest `json:"request"`
}

func (c Client) CreatePortfolioAllocations(
	ctx context.Context,
	request *CreatePortfolioAllocationsRequest,
) (*CreatePortfolioAllocationsResponse, error) {

	path := "/allocations/%s/order"

	response := &CreatePortfolioAllocationsResponse{Request: request}

	if err := post(ctx, c, path, emptyQueryParams, request, response); err != nil {
		return nil, err
	}

	return response, nil
}
