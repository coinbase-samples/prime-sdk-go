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
)

type DescribeActivitiesRequest struct {
	PortfolioId string            `json:"portfolio_id"`
	Symbols     []string          `json:"symbols"`
	Categories  []string          `json:"categories"`
	Statuses    []string          `json:"statuses"`
	StartTime   string            `json:"start_time"`
	EndTime     string            `json:"end_time"`
	Pagination  *PaginationParams `json:"pagination_params"`
}

type DescribeActivitiesResponse struct {
	Activities []*Activity                `json:"activities"`
	Request    *DescribeActivitiesRequest `json:"request"`
	Pagination *Pagination                `json:"pagination"`
}

func (c Client) DescribeActivities(
	ctx context.Context,
	request *DescribeActivitiesRequest,
) (*DescribeActivitiesResponse, error) {

	path := fmt.Sprintf("/portfolios/%s/activities", request.PortfolioId)

	queryParams := appendQueryParam(emptyQueryParams, "start_time", request.StartTime)

	queryParams = appendQueryParam(emptyQueryParams, "end_time", request.EndTime)

	for _, v := range request.Symbols {
		queryParams = appendQueryParam(queryParams, "symbols", v)
	}

	for _, v := range request.Categories {
		queryParams = appendQueryParam(queryParams, "categories", v)
	}

	for _, v := range request.Statuses {
		queryParams = appendQueryParam(queryParams, "statuses", v)
	}

	queryParams = appendPaginationParams(queryParams, request.Pagination)

	response := &DescribeActivitiesResponse{Request: request}

	if err := get(ctx, c, path, queryParams, request, response); err != nil {
		return nil, err
	}

	return response, nil
}