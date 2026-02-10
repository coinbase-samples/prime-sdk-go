/**
 * Copyright 2026-present Coinbase Global, Inc.
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

type ClaimStakingRewardsInputs struct {
	Amount string `json:"amount,omitempty"`
}

type ClaimStakingRewardsRequest struct {
	PortfolioId    string                     `json:"portfolio_id"`
	WalletId       string                     `json:"wallet_id"`
	IdempotencyKey string                     `json:"idempotency_key"`
	Inputs         *ClaimStakingRewardsInputs `json:"inputs,omitempty"`
}

type ClaimStakingRewardsResponse struct {
	WalletId      string                      `json:"wallet_id"`
	TransactionId string                      `json:"transaction_id"`
	ActivityId    string                      `json:"activity_id"`
	Request       *ClaimStakingRewardsRequest `json:"-"`
}

func (s *stakingServiceImpl) ClaimStakingRewards(
	ctx context.Context,
	request *ClaimStakingRewardsRequest,
) (*ClaimStakingRewardsResponse, error) {

	path := fmt.Sprintf("/portfolios/%s/wallets/%s/staking/claim_rewards", request.PortfolioId, request.WalletId)

	response := &ClaimStakingRewardsResponse{Request: request}

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
