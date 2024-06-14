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
	"github.com/coinbase-samples/core-go"
)

type ListPortfolioBalancesRequest struct {
	PortfolioId string   `json:"portfolio_id"`
	Type        string   `json:"balance_type"`
	Symbols     []string `json:"symbols"`
}

type ListPortfolioBalancesResponse struct {
	Balances              []*Balance                    `json:"balances"`
	Type                  string                        `json:"type"`
	TradingWalletBalances *BalanceWithHolds             `json:"trading_balances"`
	VaultWalletBalances   *BalanceWithHolds             `json:"vault_balances"`
	Request               *ListPortfolioBalancesRequest `json:"request"`
}

func (c *Client) ListPortfolioBalances(
	ctx context.Context,
	request *ListPortfolioBalancesRequest,
) (*ListPortfolioBalancesResponse, error) {

	path := fmt.Sprintf("/portfolios/%s/balances", request.PortfolioId)

	var queryParams string
	if len(request.Type) > 0 {
		queryParams = appendQueryParam(queryParams, "balance_type", request.Type)
	}

	for _, v := range request.Symbols {
		queryParams = appendQueryParam(queryParams, "symbols", v)
	}

	response := &ListPortfolioBalancesResponse{Request: request}

	if err := core.Get(ctx, c, path, queryParams, request, response, addPrimeHeaders); err != nil {
		return nil, err
	}

	return response, nil
}
