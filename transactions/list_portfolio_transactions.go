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
	Transactions []*model.Transaction              `json:"transactions"`
	Pagination   *model.Pagination                 `json:"pagination"`
	Request      *ListPortfolioTransactionsRequest `json:"request"`
}

func (s *transactionsServiceImpl) ListPortfolioTransactions(
	ctx context.Context,
	request *ListPortfolioTransactionsRequest,
) (*ListPortfolioTransactionsResponse, error) {

	path := fmt.Sprintf("/portfolios/%s/transactions", request.PortfolioId)

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

	response := &ListPortfolioTransactionsResponse{Request: request}

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
