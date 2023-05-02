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

package prime

import (
	"context"
	"fmt"
	"time"
)

type AddressBookEntry struct {
	Id                    string                   `json:"id"`
	Symbol                string                   `json:"currency_symbol"`
	Name                  string                   `json:"name"`
	Address               string                   `json:"address"`
	AccountIdentifier     string                   `json:"account_identifier"`
	AccountIdentifierName string                   `json:"account_identifier_name"`
	State                 string                   `json:"state"`
	ExplorerLink          string                   `json:"explorer_link"`
	LastUsedAt            time.Time                `json:"last_used_at"`
	AddedAt               time.Time                `json:"added_at"`
	AddedBy               *AddressBookEntryAddedBy `json:"added_by"`
}

type AddressBookEntryAddedBy struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	AvatarUrl string `json:"avatar_url"`
}

type DescribeAddressBookRequest struct {
	PortfolioId    string          `json:"portfolio_id"`
	Symbol         string          `json:"currency_symbol"`
	Search         string          `json:"search"`
	IteratorParams *IteratorParams `json:"iteratorParams"`
}

type DescribeAddressBookResponse struct {
	Addresses  []*AddressBookEntry         `json:"addresses"`
	Request    *DescribeAddressBookRequest `json:"request"`
	Pagination *Pagination                 `json:"pagination"`
}

func (c Client) DescribeAddressBook(
	ctx context.Context,
	request *DescribeAddressBookRequest,
) (*DescribeAddressBookResponse, error) {

	path := fmt.Sprintf("/portfolios/%s/address_book", request.PortfolioId)

	var queryParams string
	if len(request.Symbol) > 0 {
		queryParams = appendQueryParam(queryParams, "currency_symbol", request.Symbol)
	}

	if len(request.Search) > 0 {
		queryParams = appendQueryParam(queryParams, "search", request.Search)
	}

	queryParams = iteratorParams(queryParams, request.IteratorParams)

	response := &DescribeAddressBookResponse{Request: request}

	if err := get(ctx, c, path, queryParams, request, response); err != nil {
		return nil, err
	}

	return response, nil
}
