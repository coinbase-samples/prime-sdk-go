/**
 * Copyright 2024-present Coinbase Global, Inc.
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

package balances

import (
	"context"
	"fmt"

	"github.com/coinbase-samples/core-go"
	"github.com/coinbase-samples/prime-sdk-go/client"
	"github.com/coinbase-samples/prime-sdk-go/model"
	"github.com/coinbase-samples/prime-sdk-go/utils"
)

type ListOnchainWalletBalancesRequest struct {
	PortfolioId        string                  `json:"portfolio_id"`
	WalletId           string                  `json:"wallet_id"`
	VisibilityStatuses []string                `json:"visibility_statuses"`
	Pagination         *model.PaginationParams `json:"pagination_params"`
}

type ListOnchainWalletBalancesResponse struct {
	model.PaginationMixin                                   // provides Pagination, HasNext(), GetNextCursor()
	Balances              []*model.Web3Balance              `json:"balances"`
	DefiBalances          []*model.DefiBalance              `json:"defi_balances"`
	Request               *ListOnchainWalletBalancesRequest `json:"-"`
	service               BalancesService                   // unexported, injected by service
	serviceConfig         *model.ServiceConfig              // unexported, injected by service
}

// Next fetches the next page of results. Returns nil, nil if no more pages.
func (r *ListOnchainWalletBalancesResponse) Next(ctx context.Context) (*ListOnchainWalletBalancesResponse, error) {
	if !r.HasNext() {
		return nil, nil
	}

	nextReq := *r.Request
	nextReq.Pagination = model.PrepareNextPagination(r.Request.Pagination, r.Pagination.NextCursor)

	return r.service.ListOnchainWalletBalances(ctx, &nextReq)
}

// Iterator returns a PageIterator for convenient iteration and FetchAll.
// The iterator respects the service's ServiceConfig for MaxPages and MaxItems.
func (r *ListOnchainWalletBalancesResponse) Iterator() *model.PageIterator[*ListOnchainWalletBalancesResponse, *model.Web3Balance] {
	return model.NewPageIteratorWithConfig(r, func(resp *ListOnchainWalletBalancesResponse) []*model.Web3Balance {
		return resp.Balances
	}, r.serviceConfig)
}

func (s *balancesServiceImpl) ListOnchainWalletBalances(
	ctx context.Context,
	request *ListOnchainWalletBalancesRequest,
) (*ListOnchainWalletBalancesResponse, error) {

	path := fmt.Sprintf(
		"/portfolios/%s/wallets/%s/web3_balances",
		request.PortfolioId,
		request.WalletId,
	)

	request.Pagination = utils.ApplyDefaultLimit(request.Pagination, s.serviceConfig)

	var queryParams string

	queryParams = utils.AppendPaginationParams(queryParams, request.Pagination)

	for _, v := range request.VisibilityStatuses {
		queryParams = core.AppendHttpQueryParam(queryParams, "visibility_statuses", v)
	}

	response := &ListOnchainWalletBalancesResponse{
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
