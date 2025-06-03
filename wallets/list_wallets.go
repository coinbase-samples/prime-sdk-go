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
	PortfolioId string                  `json:"portfolio_id"`
	Type        string                  `json:"type"`
	Symbols     []string                `json:"symbols"`
	Pagination  *model.PaginationParams `json:"pagination_params"`
}

type ListWalletsResponse struct {
	Wallets    []*model.Wallet     `json:"wallets"`
	Pagination *model.Pagination   `json:"pagination"`
	Request    *ListWalletsRequest `json:"-"`
}

func (r ListWalletsResponse) HasNext() bool {
	return r.Pagination != nil && r.Pagination.HasNext
}

func (s *walletsServiceImpl) ListWallets(
	ctx context.Context,
	request *ListWalletsRequest,
) (*ListWalletsResponse, error) {

	path := fmt.Sprintf("/portfolios/%s/wallets", request.PortfolioId)

	queryParams := core.EmptyQueryParams
	if request.Type != "" {
		queryParams = core.AppendHttpQueryParam(queryParams, "type", request.Type)
	}
	for _, v := range request.Symbols {
		queryParams = core.AppendHttpQueryParam(queryParams, "symbols", v)
	}

	queryParams = utils.AppendPaginationParams(queryParams, request.Pagination)

	response := &ListWalletsResponse{Request: request}

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
