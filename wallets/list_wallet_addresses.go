/**
 * Copyright 2025-present Coinbase Global, Inc.
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

type ListWalletAddressesRequest struct {
	PortfolioId string                  `json:"portfolio_id"`
	WalletId    string                  `json:"wallet_id"`
	NetworkId   string                  `json:"network_id"`
	Pagination  *model.PaginationParams `json:"pagination_params"`
}

type ListWalletAddressesResponse struct {
	Addresses  []*model.BlockchainAddress  `json:"addresses"`
	Pagination *model.Pagination           `json:"pagination"`
	Request    *ListWalletAddressesRequest `json:"-"`
}

func (r ListWalletAddressesResponse) HasNext() bool {
	return r.Pagination != nil && r.Pagination.HasNext
}

func (s *walletsServiceImpl) ListWalletAddresses(
	ctx context.Context,
	request *ListWalletAddressesRequest,
) (*ListWalletAddressesResponse, error) {

	path := fmt.Sprintf("/portfolios/%s/wallets/%s/addresses", request.PortfolioId, request.WalletId)

	queryParams := core.EmptyQueryParams
	if request.NetworkId != "" {
		queryParams = core.AppendHttpQueryParam(queryParams, "network_id", request.NetworkId)
	}

	queryParams = utils.AppendPaginationParams(queryParams, request.Pagination)

	response := &ListWalletAddressesResponse{Request: request}

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
