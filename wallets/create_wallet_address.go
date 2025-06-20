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
)

type CreateWalletAddressRequest struct {
	PortfolioId string `json:"portfolio_id"`
	WalletId    string `json:"wallet_id"`
	NetworkId   string `json:"network_id"`
}

type CreateWalletAddressResponse struct {
	Address           string                      `json:"address"`
	AccountIdentifier string                      `json:"account_identifier"`
	Network           *model.NetworkDetails       `json:"network"`
	Request           *CreateWalletAddressRequest `json:"-"`
}

func (s *walletsServiceImpl) CreateWalletAddress(ctx context.Context, request *CreateWalletAddressRequest) (*CreateWalletAddressResponse, error) {

	path := fmt.Sprintf("/portfolios/%s/wallets/%s/addresses", request.PortfolioId, request.WalletId)

	response := &CreateWalletAddressResponse{Request: request}

	if err := core.HttpPost(
		ctx,
		s.client,
		path,
		core.EmptyQueryParams,
		client.DefaultSuccessHttpStatusCodes,
		request,
		response,
		s.client.HeadersFunc(),
	); err != nil {
		return nil, err
	}

	return response, nil
}
