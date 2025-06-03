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

package orders

import (
	"context"
	"fmt"
	"time"

	"github.com/coinbase-samples/core-go"
	"github.com/coinbase-samples/prime-sdk-go/client"
	"github.com/coinbase-samples/prime-sdk-go/model"
	"github.com/coinbase-samples/prime-sdk-go/utils"
)

type ListOrdersRequest struct {
	PortfolioId string                  `json:"portfolio_id"` // required
	Statuses    []string                `json:"order_statuses"`
	ProductIds  []string                `json:"product_ids"`
	Type        string                  `json:"order_type"`
	OrderSide   string                  `json:"order_side"`
	Start       time.Time               `json:"start_date"` // required
	End         time.Time               `json:"end_date"`
	Pagination  *model.PaginationParams `json:"pagination_params"`
}

type ListOrdersResponse struct {
	Orders     []*model.Order     `json:"orders"`
	Pagination *model.Pagination  `json:"pagination"`
	Request    *ListOrdersRequest `json:"-"`
}

// ListOrders returns orders based on query params. Start time is required.
// This API endpoint cannot list open orders, so do not add an OPEN status
// to the status param.
// https://docs.cloud.coinbase.com/prime/reference/primerestapi_getorders
func (s *ordersServiceImpl) ListOrders(
	ctx context.Context,
	request *ListOrdersRequest,
) (*ListOrdersResponse, error) {

	path := fmt.Sprintf("/portfolios/%s/orders", request.PortfolioId)

	var queryParams string

	if !request.Start.IsZero() {
		queryParams = core.AppendHttpQueryParam(queryParams, "start_date", utils.TimeToStr(request.Start))
	}

	if !request.End.IsZero() {
		queryParams = core.AppendHttpQueryParam(queryParams, "end_date", utils.TimeToStr(request.End))
	}

	if len(request.Type) > 0 {
		queryParams = core.AppendHttpQueryParam(queryParams, "order_type", request.Type)
	}

	if len(request.OrderSide) > 0 {
		queryParams = core.AppendHttpQueryParam(queryParams, "order_side", request.OrderSide)
	}

	for _, s := range request.Statuses {
		queryParams = core.AppendHttpQueryParam(queryParams, "order_statuses", s)
	}

	for _, p := range request.ProductIds {
		queryParams = core.AppendHttpQueryParam(queryParams, "product_ids", p)
	}

	queryParams = utils.AppendPaginationParams(queryParams, request.Pagination)

	response := &ListOrdersResponse{Request: request}

	if err := core.HttpGet(
		ctx,
		s.client,
		path,
		queryParams,
		client.DefaultSuccessHttpStatusCodes,
		request,
		response,
		s.client.HeadersFunc(),
	); err != nil {
		return nil, err
	}

	return response, nil
}
