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

type ListOrdersRequest struct {
	PortfolioId string            `json:"portfolio_id"` // required
	Statuses    []string          `json:"order_statuses"`
	ProductIds  []string          `json:"product_ids"`
	Type        string            `json:"order_type"`
	OtherSide   string            `json:"order_side"`
	Start       time.Time         `json:"start_date"` // required
	End         time.Time         `json:"end_date"`
	Pagination  *PaginationParams `json:"pagination_params"`
}

type ListOrdersResponse struct {
	Orders     []*Order           `json:"orders"`
	Pagination *Pagination        `json:"pagination"`
	Request    *ListOrdersRequest `json:"request"`
}

// ListOrders returns orders based on query params. Start time is requred.
// This API endpoint cannot list open orders, so do not add an OPEN status
// to the status param.
// https://docs.cloud.coinbase.com/prime/reference/primerestapi_getorders
func (c Client) ListOrders(
	ctx context.Context,
	request *ListOrdersRequest,
) (*ListOrdersResponse, error) {

	path := fmt.Sprintf("/portfolios/%s/orders", request.PortfolioId)

	var queryParams string

	queryParams = appendQueryParam(queryParams, "start_date", TimeToStr(request.Start))

	if !request.End.IsZero() {
		queryParams = appendQueryParam(queryParams, "end_date", TimeToStr(request.End))
	}

	if len(request.Type) > 0 {
		queryParams = appendQueryParam(queryParams, "order_type", request.Type)
	}

	if len(request.OtherSide) > 0 {
		queryParams = appendQueryParam(queryParams, "order_side", request.OtherSide)
	}

	for _, s := range request.Statuses {
		queryParams = appendQueryParam(queryParams, "order_statuses", s)
	}

	for _, p := range request.ProductIds {
		queryParams = appendQueryParam(queryParams, "product_ids", p)
	}

	queryParams = appendPaginationParams(queryParams, request.Pagination)

	response := &ListOrdersResponse{Request: request}
	if err := get(ctx, c, path, queryParams, request, response); err != nil {
		return response, err
	}

	return response, nil
}
