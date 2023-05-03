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

type GetActivitiesRequest struct {
	PortfolioId string            `json:"portfolio_id"`
	Symbols     []string          `json:"symbols"`
	Categories  []string          `json:"categories"`
	Statuses    []string          `json:"statuses"`
	Start       time.Time         `json:"start_time"`
	End         time.Time         `json:"end_time"`
	Pagination  *PaginationParams `json:"pagination_params"`
}

type GetActivitiesResponse struct {
	Activities []*Activity           `json:"activities"`
	Request    *GetActivitiesRequest `json:"request"`
	Pagination *Pagination           `json:"pagination"`
}

func (c Client) GetActivities(
	ctx context.Context,
	request *GetActivitiesRequest,
) (*GetActivitiesResponse, error) {

	path := fmt.Sprintf("/portfolios/%s/activities", request.PortfolioId)

	var queryParams string
	if !request.Start.IsZero() {
		queryParams = appendQueryParam(queryParams, "start_time", TimeToStr(request.Start))
	}

	if !request.End.IsZero() {
		queryParams = appendQueryParam(queryParams, "end_time", TimeToStr(request.End))
	}

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

	response := &GetActivitiesResponse{Request: request}

	if err := get(ctx, c, path, queryParams, request, response); err != nil {
		return nil, err
	}

	return response, nil
}
