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

package transactions

import (
	"context"
	"fmt"

	"github.com/coinbase-samples/core-go"
	"github.com/coinbase-samples/prime-sdk-go/client"
	"github.com/coinbase-samples/prime-sdk-go/model"
)

type SubmitDepositTravelRuleDataRequest struct {
	PortfolioId                   string                 `json:"-"`
	TransactionId                 string                 `json:"-"`
	Originator                    *model.TravelRuleParty `json:"originator,omitempty"`
	Beneficiary                   *model.TravelRuleParty `json:"beneficiary,omitempty"`
	IsSelf                        bool                   `json:"is_self,omitempty"`
	OptOutOfOwnershipVerification bool                   `json:"opt_out_of_ownership_verification,omitempty"`
}

type SubmitDepositTravelRuleDataResponse struct {
	OwnershipVerificationRequired bool                                `json:"ownership_verification_required"`
	Request                       *SubmitDepositTravelRuleDataRequest `json:"-"`
}

func (s *transactionsServiceImpl) SubmitDepositTravelRuleData(
	ctx context.Context,
	request *SubmitDepositTravelRuleDataRequest,
) (*SubmitDepositTravelRuleDataResponse, error) {

	path := fmt.Sprintf("/portfolios/%s/transactions/%s/travel_rule/deposit", request.PortfolioId, request.TransactionId)

	response := &SubmitDepositTravelRuleDataResponse{Request: request}

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
