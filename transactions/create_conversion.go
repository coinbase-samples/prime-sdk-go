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

package transactions

import (
	"context"
	"fmt"

	"github.com/coinbase-samples/core-go"
	"github.com/coinbase-samples/prime-sdk-go/client"
)

type CreateConversionRequest struct {
	PortfolioId         string `json:"portfolio_id"`
	SourceWalletId      string `json:"wallet_id"`
	SourceSymbol        string `json:"source_symbol"`
	DestinationWalletId string `json:"destination"`
	DestinationSymbol   string `json:"destination_symbol"`
	IdempotencyKey      string `json:"idempotency_key"`
	Amount              string `json:"amount"`
}

type CreateConversionResponse struct {
	ActivityId          string                   `json:"activity_id"`
	SourceSymbol        string                   `json:"source_symbol"`
	DestinationSymbol   string                   `json:"destination_symbol"`
	Amount              string                   `json:"amount"`
	DestinationWalletId string                   `json:"destination"`
	SourceWalletId      string                   `json:"source"`
	Request             *CreateConversionRequest `json:"-"`
}

func (s *transactionsServiceImpl) CreateConversion(
	ctx context.Context,
	request *CreateConversionRequest,
) (*CreateConversionResponse, error) {

	path := fmt.Sprintf("/portfolios/%s/wallets/%s/conversion",
		request.PortfolioId,
		request.SourceWalletId,
	)

	response := &CreateConversionResponse{Request: request}

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
