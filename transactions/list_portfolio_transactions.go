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

type ListPortfolioTransactionsRequest struct {
	PortfolioId string                  `json:"portfolio_id"`
	Symbols     string                  `json:"symbols"`
	Types       []string                `json:"types"`
	Start       time.Time               `json:"start_time"`
	End         time.Time               `json:"end_time"`
	Pagination  *model.PaginationParams `json:"pagination_params"`
}

type ListPortfolioTransactionsResponse struct {
	model.PaginationMixin                                   // provides Pagination, HasNext(), GetNextCursor()
	Transactions          []*model.Transaction              `json:"transactions"`
	Request               *ListPortfolioTransactionsRequest `json:"-"`
	service               TransactionsService               // unexported, injected by service
	paginationConfig      *model.PaginationConfig           // unexported, injected by service
}

// Next fetches the next page of results. Returns nil, nil if no more pages.
func (r *ListPortfolioTransactionsResponse) Next(ctx context.Context) (*ListPortfolioTransactionsResponse, error) {
	if !r.HasNext() {
		return nil, nil
	}

	nextReq := *r.Request
	if nextReq.Pagination == nil {
		nextReq.Pagination = &model.PaginationParams{}
	} else {
		paginationCopy := *nextReq.Pagination
		nextReq.Pagination = &paginationCopy
	}
	nextReq.Pagination.Cursor = r.Pagination.NextCursor

	return r.service.ListPortfolioTransactions(ctx, &nextReq)
}

// Iterator returns a PageIterator for convenient iteration and FetchAll.
// The iterator respects the service's PaginationConfig for MaxPages and MaxItems.
func (r *ListPortfolioTransactionsResponse) Iterator() *model.PageIterator[*ListPortfolioTransactionsResponse, *model.Transaction] {
	return model.NewPageIteratorWithConfig(r, func(resp *ListPortfolioTransactionsResponse) []*model.Transaction {
		return resp.Transactions
	}, r.paginationConfig)
}

func (s *transactionsServiceImpl) ListPortfolioTransactions(
	ctx context.Context,
	request *ListPortfolioTransactionsRequest,
) (*ListPortfolioTransactionsResponse, error) {

	path := fmt.Sprintf("/portfolios/%s/transactions", request.PortfolioId)

	// Apply default limit from config if not specified in request
	if s.paginationConfig != nil && s.paginationConfig.DefaultLimit > 0 {
		if request.Pagination == nil {
			request.Pagination = &model.PaginationParams{Limit: s.paginationConfig.DefaultLimit}
		} else if request.Pagination.Limit == 0 {
			request.Pagination.Limit = s.paginationConfig.DefaultLimit
		}
	}

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

	response := &ListPortfolioTransactionsResponse{
		Request:          request,
		service:          s,
		paginationConfig: s.paginationConfig,
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
