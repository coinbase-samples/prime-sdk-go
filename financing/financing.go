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

package financing

import (
	"context"

	"github.com/coinbase-samples/prime-sdk-go/client"
)

type FinancingService interface {
	CreateLocate(ctx context.Context, request *CreateLocateRequest) (*CreateLocateResponse, error)
	GetEntityLocateAvailabilities(ctx context.Context, request *GetEntityLocateAvailabilitiesRequest) (*GetEntityLocateAvailabilitiesResponse, error)
	GetBuyingPower(ctx context.Context, request *GetBuyingPowerRequest) (*GetBuyingPowerResponse, error)
	GetWithdrawalPower(ctx context.Context, request *GetWithdrawalPowerRequest) (*GetWithdrawalPowerResponse, error)
	GetMarginInfo(ctx context.Context, request *GetMarginInfoRequest) (*GetMarginInfoResponse, error)
	GetPortfolioCreditInfo(ctx context.Context, request *GetPortfolioCreditInfoRequest) (*GetPortfolioCreditInfoResponse, error)
	GetTieredPricingFees(ctx context.Context, request *GetTieredPricingFeesRequest) (*GetTieredPricingFeesResponse, error)
	ListLocates(ctx context.Context, request *ListLocatesRequest) (*ListLocatesResponse, error)
	ListInterestAccruals(ctx context.Context, request *ListInterestAccrualsRequest) (*ListInterestAccrualsResponse, error)
	ListPortfolioInterestAccruals(ctx context.Context, request *ListPortfolioInterestAccrualsRequest) (*ListPortfolioInterestAccrualsResponse, error)
	ListMarginCallSummaries(ctx context.Context, request *ListMarginCallSummariesRequest) (*ListMarginCallSummariesResponse, error)
	ListMarginConversions(ctx context.Context, request *ListMarginConversionsRequest) (*ListMarginConversionsResponse, error)
}

func NewFinancingService(c client.RestClient) FinancingService {
	return &financingServiceImpl{client: c}
}

type financingServiceImpl struct {
	client client.RestClient
}
