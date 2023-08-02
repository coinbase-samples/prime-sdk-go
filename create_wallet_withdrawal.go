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

type CreateWalletWithdrawalRequest struct {
	PortfolioId       string                               `json:"portfolio_id"`
	SourceWalletId    string                               `json:"wallet_id"`
	Amount            string                               `json:"amount"`
	DestinationType   string                               `json:"destination_type"`
	IdempotencyKey    string                               `json:"idempotency_key"`
	Symbol            string                               `json:"currency_symbol"`
	PaymentMethod     *CreateWalletWithdrawalPaymentMethod `json:""payment_method`
	BlockchainAddress *BlockchainAddress                   `json:"blockchain_address"`
}

type CreateWalletWithdrawalPaymentMethod struct {
	Id string `json:"payment_method_id"`
}

type CreateWalletWithdrawalResponse struct {
	ActivityId      string                         `json:"activity_id"`
	ApprovalUrl     string                         `json:"approval_url"`
	Symbol          string                         `json:"symbol"`
	Amount          string                         `json:"amount"`
	Fee             string                         `json:"fee"`
	DestinationType string                         `json:"destination_type"`
	SourceType      string                         `json:"source_type"`
	Destination     *BlockchainAddress             `json:"blockchain_destination"`
	Source          *BlockchainAddress             `json:"blockchain_source"`
	TransactionId   string                         `json:"transaction_id"`
	Request         *CreateWalletWithdrawalRequest `json:"request"`
}

func (c Client) CreateWalletWithdrawal(
	ctx context.Context,
	request *CreateWalletWithdrawalRequest,
) (*CreateWalletWithdrawalResponse, error) {

	path := fmt.Sprintf("/portfolios/%s/wallets/%s/withdrawals",
		request.PortfolioId,
		request.SourceWalletId,
	)

	response := &CreateWalletWithdrawalResponse{Request: request}

	if err := post(ctx, c, path, emptyQueryParams, request, response); err != nil {
		return nil, err
	}

	return response, nil
}
