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

type ListOrderFillsRequest struct {
	PortfolioId string            `json:"portfolio_id"` // required
	OrderId     string            `json:"order_id"`     // required
	Pagination  *PaginationParams `json:"pagination_params"`
}

type ListOrderFillsResponse struct {
	Fills      []*OrderFill           `json:"fills"`
	Pagination *Pagination            `json:"pagination"`
	Request    *ListOrderFillsRequest `json:"request"`
}

func (c Client) ListOrderFills(
	ctx context.Context,
	request *ListOrderFillsRequest,
) (*ListOrderFillsResponse, error) {

	path := fmt.Sprintf("/portfolios/%s/orders/%s/fills", request.PortfolioId, request.OrderId)

	queryParams := appendPaginationParams(emptyQueryParams, request.Pagination)

	response := &ListOrderFillsResponse{Request: request}
	if err := get(ctx, c, path, queryParams, request, response); err != nil {
		return response, err
	}

	return response, nil
}
