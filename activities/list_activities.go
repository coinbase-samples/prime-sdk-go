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

package activities

import (
	"context"
	"fmt"
	"time"

	"github.com/coinbase-samples/core-go"
	"github.com/coinbase-samples/prime-sdk-go/client"
	"github.com/coinbase-samples/prime-sdk-go/model"
	"github.com/coinbase-samples/prime-sdk-go/utils"
)

type ListActivitiesRequest struct {
	PortfolioId                 string                  `json:"portfolio_id"`
	Symbols                     []string                `json:"symbols"`
	Categories                  []string                `json:"categories"`
	Statuses                    []string                `json:"statuses"`
	Start                       time.Time               `json:"start_time"`
	End                         time.Time               `json:"end_time"`
	GetNetworkUnifiedActivities bool                    `json:"get_network_unified_activities,omitempty"`
	Pagination                  *model.PaginationParams `json:"pagination_params"`
}

type ListActivitiesResponse struct {
	model.PaginationMixin
	Activities    []*model.Activity      `json:"activities"`
	Request       *ListActivitiesRequest `json:"-"`
	service       ActivitiesService
	serviceConfig *model.ServiceConfig
}

// Next fetches the next page of activities using the pagination cursor.
// Returns nil if there are no more pages.
func (r *ListActivitiesResponse) Next(ctx context.Context) (*ListActivitiesResponse, error) {
	if !r.HasNext() {
		return nil, nil
	}

	nextRequest := *r.Request
	nextRequest.Pagination = model.PrepareNextPagination(r.Request.Pagination, r.GetNextCursor())

	return r.service.ListActivities(ctx, &nextRequest)
}

// Iterator returns a PageIterator for iterating through all pages of activities.
func (r *ListActivitiesResponse) Iterator() *model.PageIterator[*ListActivitiesResponse, *model.Activity] {
	return model.NewPageIteratorWithConfig(
		r,
		func(resp *ListActivitiesResponse) []*model.Activity {
			return resp.Activities
		},
		r.serviceConfig,
	)
}

func (s *activitiesServiceImpl) ListActivities(
	ctx context.Context,
	request *ListActivitiesRequest,
) (*ListActivitiesResponse, error) {

	path := fmt.Sprintf("/portfolios/%s/activities", request.PortfolioId)

	request.Pagination = utils.ApplyDefaultLimit(request.Pagination, s.serviceConfig)

	queryParams := utils.AppendPaginationParams(core.EmptyQueryParams, request.Pagination)

	if !request.Start.IsZero() {
		queryParams = core.AppendHttpQueryParam(queryParams, "start_time", utils.TimeToStr(request.Start))
	}

	if !request.End.IsZero() {
		queryParams = core.AppendHttpQueryParam(queryParams, "end_time", utils.TimeToStr(request.End))
	}

	for _, v := range request.Symbols {
		queryParams = core.AppendHttpQueryParam(queryParams, "symbols", v)
	}

	for _, v := range request.Categories {
		queryParams = core.AppendHttpQueryParam(queryParams, "categories", v)
	}

	for _, v := range request.Statuses {
		queryParams = core.AppendHttpQueryParam(queryParams, "statuses", v)
	}

	if request.GetNetworkUnifiedActivities {
		queryParams = core.AppendHttpQueryParam(queryParams, "get_network_unified_activities", "true")
	}

	response := &ListActivitiesResponse{
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
