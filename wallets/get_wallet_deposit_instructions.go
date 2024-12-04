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
)

type GetWalletDepositInstructionsRequest struct {
	PortfolioId string `json:"portfolio_id"`
	Id          string `json:"wallet_id"`
	Type        string `json:"deposit_type"`
}

type GetWalletDepositInstructionsResponse struct {
	Crypto  *model.CryptoDepositInstructions `json:"crypto_instructions"`
	Fiat    *model.FiatDepositInstructions   `json:"fiat_instructions"`
	Request *GetWalletDepositInstructionsRequest
}

func (s *walletsServiceImpl) GetWalletDepositInstructions(
	ctx context.Context,
	request *GetWalletDepositInstructionsRequest,
) (*GetWalletDepositInstructionsResponse, error) {

	path := fmt.Sprintf("/portfolios/%s/wallets/%s/deposit_instructions", request.PortfolioId, request.Id)

	queryParams := core.AppendHttpQueryParam(core.EmptyQueryParams, "deposit_type", request.Type)

	response := &GetWalletDepositInstructionsResponse{Request: request}

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
