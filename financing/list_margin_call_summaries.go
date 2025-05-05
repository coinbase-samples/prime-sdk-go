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

package financing

import (
	"context"
	"fmt"

	"github.com/coinbase-samples/core-go"
	"github.com/coinbase-samples/prime-sdk-go/client"
	"github.com/coinbase-samples/prime-sdk-go/model"
)

type ListMarginCallSummariesRequest struct {
	// required
	EntityId string `json:"entity_id"`
	// The start date of the range to query for in RFC3339 format
	StartDate string `json:"start_date"`
	// The end date of the range to query for in RFC3339 format
	EndDate string `json:"end_date"`
}

type ListMarginCallSummariesResponse struct {
	MarginSummaries []model.MarginSummaryHistorical `json:"margin_summaries"`
	Request         *ListMarginCallSummariesRequest `json:"request"`
}

func (s *financingServiceImpl) ListMarginCallSummaries(
	ctx context.Context,
	request *ListMarginCallSummariesRequest,
) (*ListMarginCallSummariesResponse, error) {

	path := fmt.Sprintf("/entities/%s/margin_summaries", request.EntityId)

	var queryParams string

	if request.StartDate != "" {
		queryParams = core.AppendHttpQueryParam(queryParams, "start_date", request.StartDate)
	}

	if request.EndDate != "" {
		queryParams = core.AppendHttpQueryParam(queryParams, "end_date", request.EndDate)
	}

	response := &ListMarginCallSummariesResponse{Request: request}

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
