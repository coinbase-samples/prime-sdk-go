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
	"fmt"
)

type GetPortfolioAllocationRequest struct {
	PortfolioId  string `json:"portfolio_id"`
	AllocationId string `json:"allocation_id"`
}

type GetPortfolioAllocationResponse struct {
	Allocation *Allocation                    `json:"allocation"`
	Request    *GetPortfolioAllocationRequest `json:"request"`
}

func (c Client) GetPortfolioAllocation(
	ctx context.Context,
	request *GetPortfolioAllocationRequest,
) (*GetPortfolioAllocationResponse, error) {

	path := fmt.Sprintf(
		"/portfolios/%s/allocations/%s",
		request.PortfolioId,
		request.AllocationId,
	)

	response := &GetPortfolioAllocationResponse{Request: request}

	if err := get(ctx, c, path, emptyQueryParams, request, response); err != nil {
		return nil, err
	}

	return response, nil
}
