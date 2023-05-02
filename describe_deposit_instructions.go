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

type DescribeWalletDepositInstructionsRequest struct {
	PortfolioId string `json:"portfolio_id"`
	Id          string `json:"wallet_id"`
	Type        string `json:"deposit_type"`
}

type DescribeWalletDepositInstructionsResponse struct {
	Crypto  *CryptoDepositInstructions `json:"crypto_instructions"`
	Fiat    *FiatDepositInstructions   `json:"fiat_instructions"`
	Request *DescribeWalletDepositInstructionsRequest
}

type CryptoDepositInstructions struct {
	Id                string `json:"id"`
	Name              string `json:"name"`
	Type              string `json:"type"`
	Address           string `json:"address"`
	AccountIdentifier string `json:"account_identifier"`
}

type FiatDepositInstructions struct {
	Id            string `json:"id"`
	Name          string `json:"name"`
	Type          string `json:"type"`
	AccountNumber string `json:"account_number"`
	RoutingNumber string `json:"routing_number"`
	ReferenceCode string `json:"reference_code"`
}

func (c Client) DescribeWalletDepositInstructions(
	ctx context.Context,
	request *DescribeWalletDepositInstructionsRequest,
) (*DescribeWalletDepositInstructionsResponse, error) {

	path := fmt.Sprintf("/portfolios/%s/wallets/%s/deposit_instructions", request.PortfolioId, request.Id)

	queryParams := appendQueryParam(emptyQueryParams, "deposit_type", request.Type)

	response := &DescribeWalletDepositInstructionsResponse{Request: request}

	if err := get(ctx, c, path, queryParams, request, response); err != nil {
		return nil, err
	}

	return response, nil
}
