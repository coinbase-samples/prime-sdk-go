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

type CreateWalletRequest struct {
	PortfolioId    string                `json:"portfolio_id"`
	Name           string                `json:"name"`
	Symbol         string                `json:"symbol,omitempty"`
	Type           string                `json:"wallet_type"`
	IdempotencyKey string                `json:"idempotency_key,omitempty"`
	Network        *model.NetworkDetails `json:"network,omitempty"`
	NetworkFamily  string                `json:"network_family,omitempty"`
}

type CreateWalletResponse struct {
	ActivityId string               `json:"activity_id"`
	Name       string               `json:"name"`
	Symbol     string               `json:"symbol"`
	Type       string               `json:"wallet_type"`
	Request    *CreateWalletRequest `json:"request"`
}

func (s *walletsServiceImpl) CreateWallet(ctx context.Context, request *CreateWalletRequest) (*CreateWalletResponse, error) {

	path := fmt.Sprintf("/portfolios/%s/wallets", request.PortfolioId)

	if len(request.IdempotencyKey) == 0 {
		request.IdempotencyKey = utils.NewUuid()
	}

	response := &CreateWalletResponse{Request: request}

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
