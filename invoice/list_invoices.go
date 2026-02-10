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

package invoice

import (
	"context"
	"fmt"
	"strconv"

	"github.com/coinbase-samples/core-go"
	"github.com/coinbase-samples/prime-sdk-go/client"
	"github.com/coinbase-samples/prime-sdk-go/model"
	"github.com/coinbase-samples/prime-sdk-go/utils"
)

type ListInvoicesRequest struct {
	EntityId     string                  `json:"entity_id"`
	States       []model.InvoiceState    `json:"states"`
	BillingYear  int32                   `json:"billing_year"`
	BillingMonth int32                   `json:"billing_month"`
	Pagination   *model.PaginationParams `json:"pagination_params"`
}

type ListInvoicesResponse struct {
	model.PaginationMixin
	Invoices      []*model.Invoice     `json:"invoices"`
	Request       *ListInvoicesRequest `json:"-"`
	service       InvoiceService
	serviceConfig *model.ServiceConfig
}

func (r *ListInvoicesResponse) Next(ctx context.Context) (*ListInvoicesResponse, error) {
	if !r.HasNext() {
		return nil, nil
	}

	nextReq := *r.Request
	nextReq.Pagination = model.PrepareNextPagination(r.Request.Pagination, r.GetNextCursor())

	return r.service.ListInvoices(ctx, &nextReq)
}

func (r *ListInvoicesResponse) Iterator() *model.PageIterator[*ListInvoicesResponse, *model.Invoice] {
	return model.NewPageIteratorWithConfig(
		r,
		func(resp *ListInvoicesResponse) []*model.Invoice {
			return resp.Invoices
		},
		r.serviceConfig,
	)
}

func (s *invoiceServiceImpl) ListInvoices(
	ctx context.Context,
	request *ListInvoicesRequest,
) (*ListInvoicesResponse, error) {

	path := fmt.Sprintf("/entities/%s/invoices", request.EntityId)

	request.Pagination = utils.ApplyDefaultLimit(request.Pagination, s.serviceConfig)

	var queryParams string
	if request.BillingYear > 0 {
		queryParams = core.AppendHttpQueryParam(queryParams, "billing_year", strconv.Itoa(int(request.BillingYear)))
	}

	if request.BillingMonth > 0 {
		queryParams = core.AppendHttpQueryParam(queryParams, "billing_month", strconv.Itoa(int(request.BillingMonth)))
	}

	for _, v := range request.States {
		queryParams = core.AppendHttpQueryParam(queryParams, "states", string(v))
	}

	queryParams = utils.AppendPaginationParams(queryParams, request.Pagination)

	response := &ListInvoicesResponse{
		Request:       request,
		service:       s,
		serviceConfig: s.serviceConfig,
	}

	if err := core.HttpGet(
		ctx,
		s.client,
		path,
		queryParams,
		client.DefaultSuccessHttpStatusCodes,
		request,
		response,
		s.client.HeadersFunc(),
	); err != nil {
		return nil, err
	}

	return response, nil
}
