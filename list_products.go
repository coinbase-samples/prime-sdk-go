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

type ListProductsRequest struct {
	PortfolioId string            `json:"portfolio_id"`
	Pagination  *PaginationParams `json:"pagination_params"`
}

type ListProductsResponse struct {
	Products   []*Product           `json:"products"`
	Pagination *Pagination          `json:"pagination"`
	Request    *ListProductsRequest `json:"request"`
}

func (c Client) ListProducts(
	ctx context.Context,
	request *ListProductsRequest,
) (*ListProductsResponse, error) {

	path := fmt.Sprintf("/portfolios/%s/products", request.PortfolioId)

	queryParams := appendPaginationParams(emptyQueryParams, request.Pagination)

	response := &ListProductsResponse{Request: request}

	if err := get(ctx, c, path, queryParams, request, response); err != nil {
		return nil, err
	}

	return response, nil
}
