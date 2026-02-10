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

type GetCrossMarginOverviewRequest struct {
	EntityId string `json:"entity_id"`
}

type GetCrossMarginOverviewResponse struct {
	Overview *model.CrossMarginOverview     `json:"overview"`
	Request  *GetCrossMarginOverviewRequest `json:"-"`
}

func (s *financingServiceImpl) GetCrossMarginOverview(
	ctx context.Context,
	request *GetCrossMarginOverviewRequest,
) (*GetCrossMarginOverviewResponse, error) {

	path := fmt.Sprintf("/entities/%s/cross_margin", request.EntityId)

	response := &GetCrossMarginOverviewResponse{Request: request}

	if err := core.HttpGet(
		ctx,
		s.client,
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
