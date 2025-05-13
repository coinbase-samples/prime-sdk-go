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
	PortfolioId         string                  `json:"portfolio_id"`
	WalletId            string                  `json:"wallet_id"`
	VisiblilityStatuses []string                `json:"visibility_statuses"`
	Pagination          *model.PaginationParams `json:"pagination_params"`
}

type ListOnchainWalletBalancesResponse struct {
	Balances              []*model.Balance                  `json:"balances"`
	Type                  string                            `json:"type"`
	TradingWalletBalances *model.BalanceWithHolds           `json:"trading_balances"`
	VaultWalletBalances   *model.BalanceWithHolds           `json:"vault_balances"`
	Request               *ListOnchainWalletBalancesRequest `json:"-"`
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

	var queryParams string

	queryParams = utils.AppendPaginationParams(queryParams, request.Pagination)

	for _, v := range request.VisiblilityStatuses {
		queryParams = core.AppendHttpQueryParam(queryParams, "visibility_statuses", v)
	}

	response := &ListOnchainWalletBalancesResponse{Request: request}

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
