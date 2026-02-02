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

type ListEntityActivitiesRequest struct {
	EntityId                    string                  `json:"entity_id"`
	ActivityLevel               string                  `json:"activity_level"`
	Symbols                     []string                `json:"symbols"`
	Categories                  []string                `json:"categories"`
	Statuses                    []string                `json:"statuses"`
	StartTime                   time.Time               `json:"start_time"`
	EndTime                     time.Time               `json:"end_time"`
	GetNetworkUnifiedActivities bool                    `json:"get_network_unified_activities,omitempty"`
	Pagination                  *model.PaginationParams `json:"pagination_params"`
}

type ListEntityActivitiesResponse struct {
	model.PaginationMixin
	Activities    []*model.Activity            `json:"activities"`
	Request       *ListEntityActivitiesRequest `json:"-"`
	service       ActivitiesService
	serviceConfig *model.ServiceConfig
}

// Next fetches the next page of entity activities using the pagination cursor.
// Returns nil if there are no more pages.
func (r *ListEntityActivitiesResponse) Next(ctx context.Context) (*ListEntityActivitiesResponse, error) {
	if !r.HasNext() {
		return nil, nil
	}

	nextRequest := *r.Request
	nextRequest.Pagination = model.PrepareNextPagination(r.Request.Pagination, r.GetNextCursor())

	return r.service.ListEntityActivities(ctx, &nextRequest)
}

// Iterator returns a PageIterator for iterating through all pages of entity activities.
func (r *ListEntityActivitiesResponse) Iterator() *model.PageIterator[*ListEntityActivitiesResponse, *model.Activity] {
	return model.NewPageIteratorWithConfig(
		r,
		func(resp *ListEntityActivitiesResponse) []*model.Activity {
			return resp.Activities
		},
		r.serviceConfig,
	)
}

func (s *activitiesServiceImpl) ListEntityActivities(
	ctx context.Context,
	request *ListEntityActivitiesRequest,
) (*ListEntityActivitiesResponse, error) {

	path := fmt.Sprintf("/entities/%s/activities", request.EntityId)

	request.Pagination = utils.ApplyDefaultLimit(request.Pagination, s.serviceConfig)

	queryParams := utils.AppendPaginationParams(core.EmptyQueryParams, request.Pagination)

	if request.ActivityLevel != "" {
		queryParams = core.AppendHttpQueryParam(queryParams, "activity_level", request.ActivityLevel)
	}

	if !request.StartTime.IsZero() {
		queryParams = core.AppendHttpQueryParam(queryParams, "start_time", utils.TimeToStr(request.StartTime))
	}

	if !request.EndTime.IsZero() {
		queryParams = core.AppendHttpQueryParam(queryParams, "end_time", utils.TimeToStr(request.EndTime))
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

	response := &ListEntityActivitiesResponse{
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
