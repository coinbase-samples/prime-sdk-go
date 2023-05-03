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

type DescribeWalletsRequest struct {
	PortfolioId string            `json:"portfolio_id"`
	Type        string            `json:"type"`
	Symbols     []string          `json:"symbols"`
	Pagination  *PaginationParams `json:"pagination_params"`
}

type DescribeWalletsResponse struct {
	Wallets    []*Wallet               `json:"wallets"`
	Request    *DescribeWalletsRequest `json:"request"`
	Pagination *Pagination             `json:"pagination"`
}

func (c Client) DescribeWallets(
	ctx context.Context,
	request *DescribeWalletsRequest,
) (*DescribeWalletsResponse, error) {

	path := fmt.Sprintf("/portfolios/%s/wallets", request.PortfolioId)

	queryParams := appendQueryParam(emptyQueryParams, "type", request.Type)

	for _, v := range request.Symbols {
		queryParams = appendQueryParam(queryParams, "symbols", v)
	}

	queryParams = appendPaginationParams(queryParams, request.Pagination)

	response := &DescribeWalletsResponse{Request: request}

	if err := get(ctx, c, path, queryParams, request, response); err != nil {
		return nil, err
	}

	return response, nil
}
