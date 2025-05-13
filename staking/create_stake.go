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

package staking

import (
	"context"
	"fmt"

	"github.com/coinbase-samples/core-go"
	"github.com/coinbase-samples/prime-sdk-go/client"
)

type CreateStakeInputs struct {
	// Optional amount to stake or unstake. If omitted, the wallet will stake or unstake the maximum amount available
	Amount string `json:"amount"`
}

type CreateStakeRequest struct {
	// required
	PortfolioId string `json:"portfolio_id"`
	// required
	WalletId string `json:"wallet_id"`
	// The client generated idempotency key for requested execution. Subsequent requests using the same key will fail
	IdempotencyKey string            `json:"idempotency_key"`
	Inputs         CreateStakeInputs `json:"inputs,omitempty"`
}

type CreateStakeResponse struct {
	// The wallet ID
	WalletId string `json:"wallet_id"`
	// ID of the newly created transaction, can be used to fetch details of the current state of execution
	TransactionId string `json:"transaction_id"`
	// The ID for the activity generated for this request
	ActivityId string              `json:"activity_id"`
	Request    *CreateStakeRequest `json:"-"`
}

func (s *stakingServiceImpl) CreateStake(
	ctx context.Context,
	request *CreateStakeRequest,
) (*CreateStakeResponse, error) {

	path := fmt.Sprintf("/portfolios/%s/wallets/%s/staking/initiate", request.PortfolioId, request.WalletId)

	var queryParams string

	response := &CreateStakeResponse{Request: request}

	if err := core.HttpPost(
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
