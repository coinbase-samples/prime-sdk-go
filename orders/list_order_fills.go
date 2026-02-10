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

	"github.com/coinbase-samples/core-go"
	"github.com/coinbase-samples/prime-sdk-go/client"
	"github.com/coinbase-samples/prime-sdk-go/model"
	"github.com/coinbase-samples/prime-sdk-go/utils"
)

type ListOrderFillsRequest struct {
	PortfolioId string                  `json:"portfolio_id"` // required
	OrderId     string                  `json:"order_id"`     // required
	Pagination  *model.PaginationParams `json:"pagination_params"`
}

type ListOrderFillsResponse struct {
	model.PaginationMixin
	Fills         []*model.OrderFill     `json:"fills"`
	Request       *ListOrderFillsRequest `json:"-"`
	service       OrdersService
	serviceConfig *model.ServiceConfig
}

// Next fetches the next page of results. Returns nil, nil if no more pages.
func (r *ListOrderFillsResponse) Next(ctx context.Context) (*ListOrderFillsResponse, error) {
	if !r.HasNext() {
		return nil, nil
	}

	nextReq := *r.Request
	nextReq.Pagination = model.PrepareNextPagination(r.Request.Pagination, r.GetNextCursor())

	return r.service.ListOrderFills(ctx, &nextReq)
}

// Iterator returns a PageIterator for convenient iteration and FetchAll.
func (r *ListOrderFillsResponse) Iterator() *model.PageIterator[*ListOrderFillsResponse, *model.OrderFill] {
	return model.NewPageIteratorWithConfig(r, func(resp *ListOrderFillsResponse) []*model.OrderFill {
		return resp.Fills
	}, r.serviceConfig)
}

func (s *ordersServiceImpl) ListOrderFills(
	ctx context.Context,
	request *ListOrderFillsRequest,
) (*ListOrderFillsResponse, error) {

	path := fmt.Sprintf("/portfolios/%s/orders/%s/fills", request.PortfolioId, request.OrderId)

	request.Pagination = utils.ApplyDefaultLimit(request.Pagination, s.serviceConfig)

	queryParams := utils.AppendPaginationParams(core.EmptyQueryParams, request.Pagination)

	response := &ListOrderFillsResponse{
		Request:       request,
		service:       s,
		serviceConfig: s.serviceConfig,
	}

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
