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
	"github.com/coinbase-samples/prime-sdk-go/model"
)

type GetTransactionRequest struct {
	PortfolioId   string `json:"portfolio_id"`
	TransactionId string `json:"transaction_id"`
}

type GetTransactionResponse struct {
	Transaction *model.Transaction     `json:"transaction"`
	Request     *GetTransactionRequest `json:"-"`
}

func (s *transactionsServiceImpl) GetTransaction(
	ctx context.Context,
	request *GetTransactionRequest,
) (*GetTransactionResponse, error) {

	path := fmt.Sprintf("/portfolios/%s/transactions/%s", request.PortfolioId, request.TransactionId)

	response := &GetTransactionResponse{Request: request}

	if err := core.HttpGet(
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
