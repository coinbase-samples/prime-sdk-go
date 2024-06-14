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
	"github.com/coinbase-samples/core-go"
)

type GetAddressBookRequest struct {
	PortfolioId string            `json:"portfolio_id"`
	Symbol      string            `json:"currency_symbol"`
	Search      string            `json:"search"`
	Pagination  *PaginationParams `json:"pagination_params"`
}

type GetAddressBookResponse struct {
	Addresses  []*AddressBookEntry    `json:"addresses"`
	Request    *GetAddressBookRequest `json:"request"`
	Pagination *Pagination            `json:"pagination"`
}

func (c *Client) GetAddressBook(
	ctx context.Context,
	request *GetAddressBookRequest,
) (*GetAddressBookResponse, error) {

	path := fmt.Sprintf("/portfolios/%s/address_book", request.PortfolioId)

	var queryParams string
	if len(request.Symbol) > 0 {
		queryParams = appendQueryParam(queryParams, "currency_symbol", request.Symbol)
	}

	if len(request.Search) > 0 {
		queryParams = appendQueryParam(queryParams, "search", request.Search)
	}

	queryParams = appendPaginationParams(queryParams, request.Pagination)

	response := &GetAddressBookResponse{Request: request}

	if err := core.Get(ctx, c, path, queryParams, request, response, addPrimeHeaders); err != nil {
		return nil, err
	}

	return response, nil
}
