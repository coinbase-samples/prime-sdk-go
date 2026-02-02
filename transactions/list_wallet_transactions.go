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

package transactions

import (
	"context"
	"fmt"
	"time"

	"github.com/coinbase-samples/core-go"
	"github.com/coinbase-samples/prime-sdk-go/client"
	"github.com/coinbase-samples/prime-sdk-go/model"
	"github.com/coinbase-samples/prime-sdk-go/utils"
)

type ListWalletTransactionsRequest struct {
	PortfolioId string                  `json:"portfolio_id"`
	WalletId    string                  `json:"wallet_id"`
	Symbols     string                  `json:"symbols"`
	Types       []string                `json:"types"`
	Start       time.Time               `json:"start_time"`
	End         time.Time               `json:"end_time"`
	Pagination  *model.PaginationParams `json:"pagination_params"`
}

type ListWalletTransactionsResponse struct {
	model.PaginationMixin
	Transactions  []*model.Transaction           `json:"transactions"`
	Request       *ListWalletTransactionsRequest `json:"-"`
	service       TransactionsService
	serviceConfig *model.ServiceConfig
}

// Next fetches the next page of results. Returns nil, nil if no more pages.
func (r *ListWalletTransactionsResponse) Next(ctx context.Context) (*ListWalletTransactionsResponse, error) {
	if !r.HasNext() {
		return nil, nil
	}

	nextReq := *r.Request
	nextReq.Pagination = model.PrepareNextPagination(r.Request.Pagination, r.GetNextCursor())

	return r.service.ListWalletTransactions(ctx, &nextReq)
}

// Iterator returns a PageIterator for convenient iteration and FetchAll.
func (r *ListWalletTransactionsResponse) Iterator() *model.PageIterator[*ListWalletTransactionsResponse, *model.Transaction] {
	return model.NewPageIteratorWithConfig(r, func(resp *ListWalletTransactionsResponse) []*model.Transaction {
		return resp.Transactions
	}, r.serviceConfig)
}

func (s *transactionsServiceImpl) ListWalletTransactions(
	ctx context.Context,
	request *ListWalletTransactionsRequest,
) (*ListWalletTransactionsResponse, error) {

	path := fmt.Sprintf(
		"/portfolios/%s/wallets/%s/transactions",
		request.PortfolioId,
		request.WalletId,
	)

	request.Pagination = utils.ApplyDefaultLimit(request.Pagination, s.serviceConfig)

	var queryParams string

	if len(request.Symbols) > 0 {
		queryParams = core.AppendHttpQueryParam(queryParams, "symbols", request.Symbols)
	}

	for _, ty := range request.Types {
		queryParams = core.AppendHttpQueryParam(queryParams, "types", ty)
	}

	if !request.Start.IsZero() {
		queryParams = core.AppendHttpQueryParam(queryParams, "start_time", utils.TimeToStr(request.Start))
	}

	if !request.End.IsZero() {
		queryParams = core.AppendHttpQueryParam(queryParams, "end_time", utils.TimeToStr(request.End))
	}

	queryParams = utils.AppendPaginationParams(queryParams, request.Pagination)

	response := &ListWalletTransactionsResponse{
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
