/**
 * Copyright 2026-present Coinbase Global, Inc.
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

package financing

import (
	"context"
	"fmt"

	"github.com/coinbase-samples/core-go"
	"github.com/coinbase-samples/prime-sdk-go/client"
	"github.com/coinbase-samples/prime-sdk-go/model"
	"github.com/coinbase-samples/prime-sdk-go/utils"
)

type GetMarketDataRequest struct {
	EntityId   string                  `json:"entity_id"`
	Pagination *model.PaginationParams `json:"pagination_params"`
}

type GetMarketDataResponse struct {
	model.PaginationMixin
	MarketData    []*model.MarketData    `json:"market_data"`
	Request       *GetMarketDataRequest  `json:"-"`
	service       FinancingService
	serviceConfig *model.ServiceConfig
}

// Next fetches the next page of market data using the pagination cursor.
// Returns nil if there are no more pages.
func (r *GetMarketDataResponse) Next(ctx context.Context) (*GetMarketDataResponse, error) {
	if !r.HasNext() {
		return nil, nil
	}

	nextRequest := *r.Request
	nextRequest.Pagination = model.PrepareNextPagination(r.Request.Pagination, r.GetNextCursor())

	return r.service.GetMarketData(ctx, &nextRequest)
}

// Iterator returns a PageIterator for iterating through all pages of market data.
func (r *GetMarketDataResponse) Iterator() *model.PageIterator[*GetMarketDataResponse, *model.MarketData] {
	return model.NewPageIteratorWithConfig(
		r,
		func(resp *GetMarketDataResponse) []*model.MarketData {
			return resp.MarketData
		},
		r.serviceConfig,
	)
}

func (s *financingServiceImpl) GetMarketData(
	ctx context.Context,
	request *GetMarketDataRequest,
) (*GetMarketDataResponse, error) {

	path := fmt.Sprintf("/entities/%s/market_data", request.EntityId)

	request.Pagination = utils.ApplyDefaultLimit(request.Pagination, s.serviceConfig)

	queryParams := utils.AppendPaginationParams(core.EmptyQueryParams, request.Pagination)

	response := &GetMarketDataResponse{
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
