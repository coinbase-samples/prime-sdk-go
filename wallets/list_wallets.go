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

package wallets

import (
	"context"
	"fmt"

	"github.com/coinbase-samples/core-go"
	"github.com/coinbase-samples/prime-sdk-go/client"
	"github.com/coinbase-samples/prime-sdk-go/model"
	"github.com/coinbase-samples/prime-sdk-go/utils"
)

type ListWalletsRequest struct {
	PortfolioId              string                  `json:"portfolio_id"`
	Type                     string                  `json:"type"`
	Symbols                  []string                `json:"symbols"`
	GetNetworkUnifiedWallets bool                    `json:"get_network_unified_wallets,omitempty"`
	Pagination               *model.PaginationParams `json:"pagination_params"`
}

type ListWalletsResponse struct {
	model.PaginationMixin                      // provides Pagination, HasNext(), GetNextCursor()
	Wallets               []*model.Wallet      `json:"wallets"`
	Request               *ListWalletsRequest  `json:"-"`
	service               WalletsService       // unexported, injected by service
	serviceConfig         *model.ServiceConfig // unexported, injected by service
}

// Next fetches the next page of results. Returns nil, nil if no more pages.
func (r *ListWalletsResponse) Next(ctx context.Context) (*ListWalletsResponse, error) {
	if !r.HasNext() {
		return nil, nil
	}

	nextReq := *r.Request
	nextReq.Pagination = model.PrepareNextPagination(r.Request.Pagination, r.Pagination.NextCursor)

	return r.service.ListWallets(ctx, &nextReq)
}

// Iterator returns a PageIterator for convenient iteration and FetchAll.
// The iterator respects the service's ServiceConfig for MaxPages and MaxItems.
func (r *ListWalletsResponse) Iterator() *model.PageIterator[*ListWalletsResponse, *model.Wallet] {
	return model.NewPageIteratorWithConfig(r, func(resp *ListWalletsResponse) []*model.Wallet {
		return resp.Wallets
	}, r.serviceConfig)
}

func (s *walletsServiceImpl) ListWallets(
	ctx context.Context,
	request *ListWalletsRequest,
) (*ListWalletsResponse, error) {

	path := fmt.Sprintf("/portfolios/%s/wallets", request.PortfolioId)

	request.Pagination = utils.ApplyDefaultLimit(request.Pagination, s.serviceConfig)

	queryParams := core.EmptyQueryParams
	if request.Type != "" {
		queryParams = core.AppendHttpQueryParam(queryParams, "type", request.Type)
	}
	for _, v := range request.Symbols {
		queryParams = core.AppendHttpQueryParam(queryParams, "symbols", v)
	}
	if request.GetNetworkUnifiedWallets {
		queryParams = core.AppendHttpQueryParam(queryParams, "get_network_unified_wallets", "true")
	}

	queryParams = utils.AppendPaginationParams(queryParams, request.Pagination)

	response := &ListWalletsResponse{
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
