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

type CreateWalletTransferRequest struct {
	PortfolioId         string `json:"portfolio_id"`
	SourceWalletId      string `json:"wallet_id"`
	Symbol              string `json:"currency_symbol"`
	DestinationWalletId string `json:"destination"`
	IdempotencyKey      string `json:"idempotency_key"`
	Amount              string `json:"amount"`
}

type CreateWalletTransferResponse struct {
	ActivityId         string                       `json:"activity_id"`
	ApprovalUrl        string                       `json:"approval_url"`
	Symbol             string                       `json:"symbol"`
	Amount             string                       `json:"amount"`
	Fee                string                       `json:"fee"`
	DestinationAddress string                       `json:"destination_address"`
	DestinationType    string                       `json:"destination_type"`
	SourceAddress      string                       `json:"source_address"`
	SourceType         string                       `json:"source_type"`
	TransactionId      string                       `json:"transaction_id"`
	Request            *CreateWalletTransferRequest `json:"-"`
}

func (s *transactionsServiceImpl) CreateWalletTransfer(
	ctx context.Context,
	request *CreateWalletTransferRequest,
) (*CreateWalletTransferResponse, error) {

	path := fmt.Sprintf("/portfolios/%s/wallets/%s/transfers",
		request.PortfolioId,
		request.SourceWalletId,
	)

	response := &CreateWalletTransferResponse{Request: request}

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
