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
	"github.com/coinbase-samples/prime-sdk-go/model"
)

type GetStakingStatusRequest struct {
	PortfolioId string `json:"portfolio_id"`
	WalletId    string `json:"wallet_id"`
}

type GetStakingStatusResponse struct {
	PortfolioId      string                        `json:"portfolio_id"`
	WalletId         string                        `json:"wallet_id"`
	WalletAddress    string                        `json:"wallet_address"`
	CurrentTimestamp string                        `json:"current_timestamp"`
	Validators       []*model.ValidatorStakingInfo `json:"validators"`
	Request          *GetStakingStatusRequest      `json:"-"`
}

func (s *stakingServiceImpl) GetStakingStatus(
	ctx context.Context,
	request *GetStakingStatusRequest,
) (*GetStakingStatusResponse, error) {

	path := fmt.Sprintf("/portfolios/%s/wallets/%s/staking/status", request.PortfolioId, request.WalletId)

	response := &GetStakingStatusResponse{Request: request}

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
