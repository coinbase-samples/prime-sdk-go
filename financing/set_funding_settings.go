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

package financing

import (
	"context"
	"fmt"

	"github.com/coinbase-samples/core-go"
	"github.com/coinbase-samples/prime-sdk-go/client"
)

type SetFundingSettingsRequest struct {
	EntityId string `json:"entity_id"`
	// DesignatedFundingPortfolioId is the Derivatives Funding Portfolio used to fund
	// FCM margin calls and receive excess margin sweeps.
	DesignatedFundingPortfolioId string `json:"designated_funding_portfolio_id"`
	// AutomaticConversionEnabled, when true, converts USDC to USD automatically to
	// meet FCM margin calls.
	AutomaticConversionEnabled bool `json:"automatic_conversion_enabled"`
	// AutomaticLoanEnabled, when true, allows Coinbase affiliates to initiate loans
	// on behalf of the entity to meet FCM margin calls.
	AutomaticLoanEnabled bool `json:"automatic_loan_enabled"`
	// AutomaticExcessReturnEnabled, when true, sweeps FCM account balance above
	// margin requirements back to the Derivatives Funding Portfolio.
	AutomaticExcessReturnEnabled bool `json:"automatic_excess_return_enabled"`
	// ExcessFundsTargetAmount is the target amount to maintain in the Futures account
	// above margin requirements (Weekend Buying Power).
	ExcessFundsTargetAmount string `json:"excess_funds_target_amount"`
}

type SetFundingSettingsResponse struct {
	ActivityId            string                     `json:"activity_id"`
	ActivityType          string                     `json:"activity_type"`
	NumApprovalsRemaining int32                      `json:"num_approvals_remaining"`
	Request               *SetFundingSettingsRequest `json:"-"`
}

func (s *financingServiceImpl) SetFundingSettings(
	ctx context.Context,
	request *SetFundingSettingsRequest,
) (*SetFundingSettingsResponse, error) {

	path := fmt.Sprintf("/entities/%s/funding/settings", request.EntityId)

	response := &SetFundingSettingsResponse{Request: request}

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
