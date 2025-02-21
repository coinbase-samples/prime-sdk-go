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

type DeleteOnchainAddressBookEntryRequest struct {
	AddressGroupId string `json:"address_group_id"`
	PortfolioId    string `json:"portfolio_id"`
}

type DeleteOnchainAddressBookEntryResponse struct {
	ActivityId         string                                `json:"activity_id"`
	ActivityType       model.OnchainActivityType             `json:"activity_type"`
	RemainingApprovals int32                                 `json:"num_approvals_remaining"`
	Request            *DeleteOnchainAddressBookEntryRequest `json:"request"`
}

func (s *onchainAddressBookServiceImpl) DeleteOnchainAddressBookEntry(
	ctx context.Context,
	request *DeleteOnchainAddressBookEntryRequest,
) (*DeleteOnchainAddressBookEntryResponse, error) {

	path := fmt.Sprintf(
		"/portfolios/%s/onchain_address_group/%s",
		request.PortfolioId,
		request.AddressGroupId,
	)

	response := &DeleteOnchainAddressBookEntryResponse{Request: request}

	if err := core.HttpDelete(
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
