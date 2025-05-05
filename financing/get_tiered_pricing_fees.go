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

type GetTieredPricingFeesRequest struct {
	EntityId string `json:"entity_id"` // required
	// The fees on a specific effective date in RFC3339 format
	EffectiveAt string `json:"effective_at"`
}

type GetTieredPricingFeesResponse struct {
	Fees    *[]model.TieredPricingFee    `json:"fees,omitempty"`
	Request *GetTieredPricingFeesRequest `json:"request"`
}

func (s *financingServiceImpl) GetTieredPricingFees(
	ctx context.Context,
	request *GetTieredPricingFeesRequest,
) (*GetTieredPricingFeesResponse, error) {

	path := fmt.Sprintf("/entities/%s/tf_tiered_fees", request.EntityId)

	var queryParams string

	if request.EffectiveAt != "" {
		queryParams = core.AppendHttpQueryParam(queryParams, "effective_at", request.EffectiveAt)
	}

	response := &GetTieredPricingFeesResponse{Request: request}

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
