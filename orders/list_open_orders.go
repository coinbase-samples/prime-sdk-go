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

type ListOpenOrdersRequest struct {
	PortfolioId string                  `json:"portfolio_id"`
	ProductIds  []string                `json:"product_ids,omitempty"`
	OrderType   string                  `json:"order_type,omitempty"`
	OrderSide   string                  `json:"order_side,omitempty"`
	Start       time.Time               `json:"start_date,omitempty"`
	End         time.Time               `json:"end_date,omitempty"`
	Pagination  *model.PaginationParams `json:"pagination_params,omitempty"`
}

type ListOpenOrdersResponse struct {
	Orders     []*model.Order         `json:"orders"`
	Pagination *model.Pagination      `json:"pagination"`
	Request    *ListOpenOrdersRequest `json:"-"`
}

// ListOpenOrders enables searching for open orders by product id.
// This API endpoint is currently being adjusted by Coinbase.
// This function will change once the Prime endpoint design is finalized.
// This will not return more than 1k open orders and pagination is not supported.
// https://docs.cloud.coinbase.com/prime/reference/primerestapi_getopenorders
func (s *ordersServiceImpl) ListOpenOrders(
	ctx context.Context,
	request *ListOpenOrdersRequest,
) (*ListOpenOrdersResponse, error) {

	path := fmt.Sprintf("/portfolios/%s/open_orders", request.PortfolioId)

	var queryParams string

	if !request.Start.IsZero() {
		queryParams = core.AppendHttpQueryParam(queryParams, "start_date", utils.TimeToStr(request.Start))
	}

	if !request.End.IsZero() {
		queryParams = core.AppendHttpQueryParam(queryParams, "end_date", utils.TimeToStr(request.End))
	}

	for _, p := range request.ProductIds {
		queryParams = core.AppendHttpQueryParam(queryParams, "product_ids", p)
	}

	response := &ListOpenOrdersResponse{Request: request}

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
