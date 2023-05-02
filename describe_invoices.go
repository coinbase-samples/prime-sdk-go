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
)

type DescribeInvoicesRequest struct {
	EntityId       string          `json:"entity_id"`
	States         []string        `json:"states"`
	BillingYear    int32           `json:"billing_year"`
	BillingMonth   int32           `json:"billing_month"`
	PaginationParams *PaginationParams `json:"pagination_params"`
}

type Invoice struct {
	Id            string         `json:"id"`
	BillingYear   int32          `json:"billing_year"`
	BillingMonth  int32          `json:"billing_month"`
	DueDate       string         `json:"due_date"`
	InvoiceNumber string         `json:"invoice_number"`
	State         string         `json:"state"`
	UsdAmountPaid float64        `json:"usd_amount_paid"`
	UsdAmountOwed float64        `json:"usd_amount_owed"`
	Items         []*InvoiceItem `json:"invoice_items"`
}

type InvoiceItem struct {
	Description    string  `json:"description"`
	CurrencySymbol string  `json:"currency_symbol"`
	InvoiceType    string  `json:"invoice_type"`
	Type           string  `json:"invoice_type"`
	Rate           float64 `json:"rate"`
	Quantity       float64 `json:"quantity"`
	Price          float64 `json:"price"`
	AverageAuc     float64 `json:"average_auc"`
	Total          float64 `json:"total"`
}

type DescribeInvoicesResponse struct {
	Invoices   []*Invoice               `json:"invoices"`
	Request    *DescribeInvoicesRequest `json:"request"`
	Pagination *Pagination              `json:"pagination"`
}

func (c Client) DescribeInvoices(
	ctx context.Context,
	request *DescribeInvoicesRequest,
) (*DescribeInvoicesResponse, error) {

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

	queryParams = appendPaginationParams(queryParams, request.PaginationParams)

	response := &DescribeInvoicesResponse{Request: request}

	if err := get(ctx, c, path, queryParams, request, response); err != nil {
		return nil, err
	}

	return response, nil
}
