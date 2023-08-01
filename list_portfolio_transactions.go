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

type ListPortfolioTransactionsRequest struct {
	PortfolioId string            `json:"portfolio_id"`
	Symbols     string            `json:"symbols"`
	Types       []string          `json:"types"`
	Start       time.Time         `json:"start_time"`
	End         time.Time         `json:"end_time"`
	Pagination  *PaginationParams `json:"pagination_params"`
}

type ListPortfolioTransactionsResponse struct {
	Transactions []*Transaction                    `json:"transactions"`
	Pagination   *Pagination                       `json:"pagination"`
	Request      *ListPortfolioTransactionsRequest `json:"request"`
}

func (c Client) ListPortfolioTransactions(
	ctx context.Context,
	request *ListPortfolioTransactionsRequest,
) (*ListPortfolioTransactionsResponse, error) {

	path := fmt.Sprintf("/portfolios/%s/transactions", request.PortfolioId)

	var queryParams string

	if len(request.Symbols) > 0 {
		queryParams = appendQueryParam(queryParams, "symbols", request.Symbols)
	}

	for _, ty := range request.Types {
		queryParams = appendQueryParam(queryParams, "types", ty)
	}

	if !request.Start.IsZero() {
		queryParams = appendQueryParam(queryParams, "start_time", TimeToStr(request.Start))
	}

	if !request.End.IsZero() {
		queryParams = appendQueryParam(queryParams, "end_time", TimeToStr(request.End))
	}

	queryParams = appendPaginationParams(queryParams, request.Pagination)

	response := &ListPortfolioTransactionsResponse{Request: request}

	if err := get(ctx, c, path, emptyQueryParams, request, response); err != nil {
		return nil, err
	}

	return response, nil
}
