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
	"time"
)

type ListPortfolioAllocationsRequest struct {
	PortfolioId string            `json:"portfolio_id"`
	ProductIds  []string          `json:"product_ids"`
	Side        string            `json:"order_side"`
	Start       time.Time         `json:"start_date"`
	End         time.Time         `json:"end_date"`
	Pagination  *PaginationParams `json:"pagination_params"`
}

type ListPortfolioAllocationsResponse struct {
	Request    *ListPortfolioAllocationsRequest `json:"request"`
	Pagination *Pagination                      `json:"pagination"`
}

func (c Client) ListPortfolioAllocations(
	ctx context.Context,
	request *ListPortfolioAllocationsRequest,
) (*ListPortfolioAllocationsResponse, error) {

	path := fmt.Sprintf("/portfolios/%s/allocations", request.PortfolioId)

	queryParams := appendPaginationParams(emptyQueryParams, request.Pagination)

	if !request.Start.IsZero() {
		queryParams = appendQueryParam(queryParams, "start_date", TimeToStr(request.Start))
	}

	if !request.End.IsZero() {
		queryParams = appendQueryParam(queryParams, "end_datae", TimeToStr(request.End))
	}

	if len(request.Side) > 0 {
		queryParams = appendQueryParam(queryParams, "side", request.Side)
	}

	for _, v := range request.ProductIds {
		queryParams = appendQueryParam(queryParams, "product_ids", v)
	}

	response := &ListPortfolioAllocationsResponse{Request: request}

	if err := get(ctx, c, path, queryParams, request, response); err != nil {
		return nil, err
	}

	return response, nil
}
