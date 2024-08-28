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
	"strconv"

	"github.com/coinbase-samples/core-go"
)

type ListInvoicesRequest struct {
	EntityId     string            `json:"entity_id"`
	States       []string          `json:"states"`
	BillingYear  int32             `json:"billing_year"`
	BillingMonth int32             `json:"billing_month"`
	Pagination   *PaginationParams `json:"pagination_params"`
}

type ListInvoicesResponse struct {
	Invoices   []*Invoice           `json:"invoices"`
	Request    *ListInvoicesRequest `json:"request"`
	Pagination *Pagination          `json:"pagination"`
}

func (c *Client) ListInvoices(
	ctx context.Context,
	request *ListInvoicesRequest,
) (*ListInvoicesResponse, error) {

	path := fmt.Sprintf("/entities/%s/invoices", request.EntityId)

	var queryParams string
	if request.BillingYear > 0 {
		queryParams = appendQueryParam(queryParams, "billing_year", strconv.Itoa(int(request.BillingYear)))
	}

	if request.BillingMonth > 0 {
		queryParams = appendQueryParam(queryParams, "billing_month", strconv.Itoa(int(request.BillingMonth)))
	}

	for _, v := range request.States {
		queryParams = appendQueryParam(queryParams, "states", v)
	}

	queryParams = appendPaginationParams(queryParams, request.Pagination)

	response := &ListInvoicesResponse{Request: request}

	if err := core.Get(ctx, c, path, queryParams, request, response, c.headersFunc); err != nil {
		return nil, err
	}

	return response, nil
}
