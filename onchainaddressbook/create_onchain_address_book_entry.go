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

package onchainaddressbook

import (
	"context"
	"fmt"

	"github.com/coinbase-samples/core-go"
	"github.com/coinbase-samples/prime-sdk-go/client"
	"github.com/coinbase-samples/prime-sdk-go/model"
)

type CreateOnchainAddressBookEntryRequest struct {
	AddressGroup *model.OnchainAddressGroup `json:"address_group"`
	PortfolioId  string                     `json:"portfolio_id"`
}

type CreateOnchainAddressBookEntryResponse struct {
	ActivityId         string                                `json:"activity_id"`
	ActivityType       string                                `json:"activity_type"`
	RemainingApprovals int32                                 `json:"num_approvals_remaining"`
	Request            *CreateOnchainAddressBookEntryRequest `json:"request"`
}

func (s *onchainAddressBookServiceImpl) CreateOnchainAddressBookEntry(
	ctx context.Context,
	request *CreateOnchainAddressBookEntryRequest,
) (*CreateOnchainAddressBookEntryResponse, error) {

	path := fmt.Sprintf("/portfolios/%s/onchain_address_group", request.PortfolioId)

	response := &CreateOnchainAddressBookEntryResponse{Request: request}

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
