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
)

type DescribeBalancesRequest struct {
	PortfolioId string   `json:"portfolio_id"`
	Type        string   `json:"balance_type"`
	Symbols     []string `json:"symbols"`
}

type DescribeBalancesResponse struct {
	Balances        []*AssetBalances         `json:"balances"`
	Type            string                   `json:"type"`
	TradingBalances *BalanceWithHolds        `json:"trading_balances"`
	VaultBalances   *BalanceWithHolds        `json:"vault_balances"`
	Request         *DescribeBalancesRequest `json:"request"`
}

func (c Client) DescribeBalances(
	ctx context.Context,
	request *DescribeBalancesRequest,
) (*DescribeBalancesResponse, error) {

	path := fmt.Sprintf("/portfolios/%s/balances", request.PortfolioId)

	var appended bool
	var queryParams string
	if len(request.Type) > 0 {
		queryParams += fmt.Sprintf("?balance_type=%s", request.Type)
		appended = true
	}

	for _, v := range request.Symbols {
		queryParams += fmt.Sprintf("%ssymbols=%s", queryParamSep(appended), v)
		appended = true
	}

	response := &DescribeBalancesResponse{Request: request}

	if err := get(ctx, c, path, queryParams, request, response); err != nil {
		return nil, err
	}

	return response, nil

}
