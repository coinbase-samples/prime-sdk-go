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

package advancedtransfers

import (
	"context"
	"fmt"
	"time"

	"github.com/coinbase-samples/core-go"
	"github.com/coinbase-samples/prime-sdk-go/client"
	"github.com/coinbase-samples/prime-sdk-go/model"
	"github.com/coinbase-samples/prime-sdk-go/utils"
)

type ListAdvancedTransfersRequest struct {
	PortfolioId           string                        `json:"-"`
	State                 model.AdvancedTransferState   `json:"state,omitempty"`
	Type                  model.AdvancedTransferType    `json:"type,omitempty"`
	Start                 time.Time                     `json:"start_time,omitempty"`
	End                   time.Time                     `json:"end_time,omitempty"`
	ReferenceId           string                        `json:"reference_id,omitempty"`
	SortDirection         string                        `json:"sort_direction,omitempty"`
	Pagination            *model.PaginationParams       `json:"pagination_params,omitempty"`
}

type ListAdvancedTransfersResponse struct {
	model.PaginationMixin
	AdvancedTransfers []*model.AdvancedTransfer      `json:"advanced_transfers"`
	Request           *ListAdvancedTransfersRequest  `json:"-"`
	service           AdvancedTransfersService
	serviceConfig     *model.ServiceConfig
}

// Next fetches the next page of results. Returns nil, nil if no more pages.
func (r *ListAdvancedTransfersResponse) Next(ctx context.Context) (*ListAdvancedTransfersResponse, error) {
	if !r.HasNext() {
		return nil, nil
	}

	nextReq := *r.Request
	nextReq.Pagination = model.PrepareNextPagination(r.Request.Pagination, r.GetNextCursor())

	return r.service.ListAdvancedTransfers(ctx, &nextReq)
}

// Iterator returns a PageIterator for convenient iteration and FetchAll.
func (r *ListAdvancedTransfersResponse) Iterator() *model.PageIterator[*ListAdvancedTransfersResponse, *model.AdvancedTransfer] {
	return model.NewPageIteratorWithConfig(r, func(resp *ListAdvancedTransfersResponse) []*model.AdvancedTransfer {
		return resp.AdvancedTransfers
	}, r.serviceConfig)
}

func (s *advancedTransfersServiceImpl) ListAdvancedTransfers(
	ctx context.Context,
	request *ListAdvancedTransfersRequest,
) (*ListAdvancedTransfersResponse, error) {

	path := fmt.Sprintf("/portfolios/%s/advanced_transfers", request.PortfolioId)

	request.Pagination = utils.ApplyDefaultLimit(request.Pagination, s.serviceConfig)

	queryParams := core.EmptyQueryParams

	if request.State != "" {
		queryParams = core.AppendHttpQueryParam(queryParams, "state", string(request.State))
	}
	if request.Type != "" {
		queryParams = core.AppendHttpQueryParam(queryParams, "type", string(request.Type))
	}
	if !request.Start.IsZero() {
		queryParams = core.AppendHttpQueryParam(queryParams, "start_time", utils.TimeToStr(request.Start))
	}
	if !request.End.IsZero() {
		queryParams = core.AppendHttpQueryParam(queryParams, "end_time", utils.TimeToStr(request.End))
	}
	if request.ReferenceId != "" {
		queryParams = core.AppendHttpQueryParam(queryParams, "reference_id", request.ReferenceId)
	}
	if request.SortDirection != "" {
		queryParams = core.AppendHttpQueryParam(queryParams, "sort_direction", request.SortDirection)
	}

	queryParams = utils.AppendPaginationParams(queryParams, request.Pagination)

	response := &ListAdvancedTransfersResponse{
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
