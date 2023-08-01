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

type ListOpenOrdersRequest struct {
	PortfolioId string `json:"portfolio_id"`
	ProductId   string `json:"product_id"`
}

type ListOpenOrdersResponse struct {
	Orders  []*Order               `json:"orders"`
	Request *ListOpenOrdersRequest `json:"request"`
}

// ListOpenOrders enables searching for open orders by product id.
// This API endpoint is currently being adjusted by Coinbase.
// This function will change once the Prime endpoint design is finalized.
// This will not return more than 1k open orders and pagination is not supported.
// https://docs.cloud.coinbase.com/prime/reference/primerestapi_getopenorders
func (c Client) ListOpenOrders(
	ctx context.Context,
	request *ListOpenOrdersRequest,
) (*ListOpenOrdersResponse, error) {

	path := fmt.Sprintf("/portfolios/%s/open_orders", request.PortfolioId)

	queryParams := appendQueryParam(emptyQueryParams, "product_ids", request.ProductId)

	response := &ListOpenOrdersResponse{Request: request}

	if err := get(ctx, c, path, queryParams, request, response); err != nil {
		return nil, err
	}

	return response, nil
}
