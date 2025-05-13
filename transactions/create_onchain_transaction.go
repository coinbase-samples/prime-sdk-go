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

package transactions

import (
	"context"
	"fmt"

	"github.com/coinbase-samples/core-go"
	"github.com/coinbase-samples/prime-sdk-go/client"
	"github.com/coinbase-samples/prime-sdk-go/model"
)

type CreateOnchainTransactionRequest struct {
	PortfolioId        string                    `json:"portfolio_id"`
	WalletId           string                    `json:"wallet_id"`
	OnchainTransaction *model.OnchainTransaction `json:"onchain_tx"`
}

type CreateOnchainTransactionResposne struct {
	TransactionId string                           `json:"transaction_id"`
	Request       *CreateOnchainTransactionRequest `json:"-"`
}

func (s *transactionsServiceImpl) CreateOnchainTransaction(
	ctx context.Context,
	request *CreateOnchainTransactionRequest,
) (*CreateOnchainTransactionResposne, error) {

	path := fmt.Sprintf("/portfolios/%s/wallets/%s/onchain_transaction",
		request.PortfolioId,
		request.WalletId,
	)

	if request.OnchainTransaction == nil {
		return nil, fmt.Errorf("OnchainTransaction not set in request")
	}

	response := &CreateOnchainTransactionResposne{Request: request}

	if err := core.HttpPost(
		ctx,
		s.client,
		path,
		core.EmptyQueryParams,
		client.DefaultSuccessHttpStatusCodes,
		request.OnchainTransaction,
		response,
		s.client.HeadersFunc(),
	); err != nil {
		return nil, err
	}

	return response, nil
}
