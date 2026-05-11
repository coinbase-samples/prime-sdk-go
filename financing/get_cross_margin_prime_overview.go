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

package financing

import (
	"context"
	"fmt"

	"github.com/coinbase-samples/core-go"
	"github.com/coinbase-samples/prime-sdk-go/client"
	"github.com/coinbase-samples/prime-sdk-go/model"
)

type GetCrossMarginPrimeOverviewRequest struct {
	EntityId string `json:"entity_id"`
}

type GetCrossMarginPrimeOverviewResponse struct {
	ControlStatus model.PrimeXMControlStatus           `json:"control_status,omitempty"`
	MarginLevel   model.PrimeXMMarginLevel              `json:"margin_level,omitempty"`
	EvaluatedAt   string                                `json:"evaluated_at,omitempty"`
	MarginSummary *model.CrossMarginPrimeMarginSummary  `json:"margin_summary,omitempty"`
	Request       *GetCrossMarginPrimeOverviewRequest   `json:"-"`
}

func (s *financingServiceImpl) GetCrossMarginPrimeOverview(
	ctx context.Context,
	request *GetCrossMarginPrimeOverviewRequest,
) (*GetCrossMarginPrimeOverviewResponse, error) {

	// This endpoint is served from /v2, so we swap the version on the base URL
	// for this call only without mutating the shared client.
	v2 := client.WithBaseUrl(s.client, client.VersionedBaseUrl(s.client.HttpBaseUrl(), "v2"))

	path := fmt.Sprintf("/entities/%s/cross_margin/prime", request.EntityId)

	response := &GetCrossMarginPrimeOverviewResponse{Request: request}

	if err := core.HttpGet(
		ctx,
		v2,
		path,
		core.EmptyQueryParams,
		client.DefaultSuccessHttpStatusCodes,
		request,
		response,
		s.client.HeadersFunc(),
	); err != nil {
		return nil, err
	}

	return response, nil
}
