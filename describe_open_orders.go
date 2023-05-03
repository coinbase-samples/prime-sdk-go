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

// This API endpoint is currently being adjusted by Coinbase
// https://docs.cloud.coinbase.com/prime/reference/primerestapi_getopenorders
// This function will change once the Prime endpoint desing is finalized
// This will not return more than 1k open orders and pagination is not supported

type DescribeOpenOrdersRequest struct {
	PortfolioId string `json:"portfolio_id"`
	ProductId   string `json:"product_id"`
}

type DescribeOpenOrdersResponse struct {
	Orders  []*Order                   `json:"orders"`
	Request *DescribeOpenOrdersRequest `json:"request"`
}

func (c Client) DescribeOpenOrders(
	ctx context.Context,
	request *DescribeOpenOrdersRequest,
) (*DescribeOpenOrdersResponse, error) {

	path := fmt.Sprintf("/portfolios/%s/open_orders", request.PortfolioId)

	queryParams := fmt.Sprintf("?product_ids=%s", request.ProductId)

	response := &DescribeOpenOrdersResponse{Request: request}

	if err := get(ctx, c, path, queryParams, request, response); err != nil {
		return response, fmt.Errorf("unable to DescribeOpenOrders: %w", err)
	}

	return response, nil
}
