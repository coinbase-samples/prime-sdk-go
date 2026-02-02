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

	"github.com/coinbase-samples/prime-sdk-go/client"
	"github.com/coinbase-samples/prime-sdk-go/model"
)

type StakingService interface {
	// Portfolio-level staking
	PortfolioStakeInitiate(ctx context.Context, request *PortfolioStakeInitiateRequest) (*PortfolioStakeInitiateResponse, error)
	PortfolioUnstake(ctx context.Context, request *PortfolioUnstakeRequest) (*PortfolioUnstakeResponse, error)
	QueryTransactionValidators(ctx context.Context, request *QueryTransactionValidatorsRequest) (*QueryTransactionValidatorsResponse, error)

	// Wallet-level staking
	CreateStake(ctx context.Context, request *CreateStakeRequest) (*CreateStakeResponse, error)
	CreateUnstake(ctx context.Context, request *CreateUnstakeRequest) (*CreateUnstakeResponse, error)
	ClaimStakingRewards(ctx context.Context, request *ClaimStakingRewardsRequest) (*ClaimStakingRewardsResponse, error)
	GetStakingStatus(ctx context.Context, request *GetStakingStatusRequest) (*GetStakingStatusResponse, error)
	PreviewUnstake(ctx context.Context, request *PreviewUnstakeRequest) (*PreviewUnstakeResponse, error)
	GetUnstakingStatus(ctx context.Context, request *GetUnstakingStatusRequest) (*GetUnstakingStatusResponse, error)

	ServiceConfig() *model.ServiceConfig
}

func NewStakingService(c client.RestClient) StakingService {
	return &stakingServiceImpl{
		client:        c,
		serviceConfig: model.DefaultServiceConfig(),
	}
}

func NewStakingServiceWithConfig(c client.RestClient, config *model.ServiceConfig) StakingService {
	return &stakingServiceImpl{
		client:        c,
		serviceConfig: config,
	}
}

type stakingServiceImpl struct {
	client        client.RestClient
	serviceConfig *model.ServiceConfig
}

func (s *stakingServiceImpl) ServiceConfig() *model.ServiceConfig {
	return s.serviceConfig
}
