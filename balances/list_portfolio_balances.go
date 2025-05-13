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

package balances

import (
	"context"
	"fmt"

	"github.com/coinbase-samples/core-go"
	"github.com/coinbase-samples/prime-sdk-go/client"
	"github.com/coinbase-samples/prime-sdk-go/model"
)

type ListPortfolioBalancesRequest struct {
	PortfolioId string   `json:"portfolio_id"`
	Type        string   `json:"balance_type"`
	Symbols     []string `json:"symbols"`
}

type ListPortfolioBalancesResponse struct {
	Balances              []*model.Balance              `json:"balances"`
	Type                  string                        `json:"type"`
	TradingWalletBalances *model.BalanceWithHolds       `json:"trading_balances"`
	VaultWalletBalances   *model.BalanceWithHolds       `json:"vault_balances"`
	Request               *ListPortfolioBalancesRequest `json:"-"`
}

func (s *balancesServiceImpl) ListPortfolioBalances(
	ctx context.Context,
	request *ListPortfolioBalancesRequest,
) (*ListPortfolioBalancesResponse, error) {

	path := fmt.Sprintf("/portfolios/%s/balances", request.PortfolioId)

	var queryParams string
	if len(request.Type) > 0 {
		queryParams = core.AppendHttpQueryParam(queryParams, "balance_type", request.Type)
	}

	for _, v := range request.Symbols {
		queryParams = core.AppendHttpQueryParam(queryParams, "symbols", v)
	}

	response := &ListPortfolioBalancesResponse{Request: request}

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
