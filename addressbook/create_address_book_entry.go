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

package addressbook

import (
	"context"
	"fmt"

	"github.com/coinbase-samples/core-go"
	"github.com/coinbase-samples/prime-sdk-go/client"
)

type CreateAddressBookEntryRequest struct {
	PortfolioId       string `json:"portfolio_id"`
	Address           string `json:"address"`
	Symbol            string `json:"currency_symbol"`
	Name              string `json:"name"`
	AccountIdentifier string `json:"account_identifier,omitempty"`
}

type CreateAddressBookEntryResponse struct {
	ActivityId         string                         `json:"activity_id"`
	Type               string                         `json:"activity_type"`
	RemainingApprovals int32                          `json:"num_approvals_remaining"`
	Request            *CreateAddressBookEntryRequest `json:"-"`
}

func (s *addressBookServiceImpl) CreateAddressBookEntry(
	ctx context.Context,
	request *CreateAddressBookEntryRequest,
) (*CreateAddressBookEntryResponse, error) {

	path := fmt.Sprintf("/portfolios/%s/address_book", request.PortfolioId)

	response := &CreateAddressBookEntryResponse{Request: request}

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
