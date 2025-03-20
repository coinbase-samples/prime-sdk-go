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

package futures

import (
	"context"
	"fmt"

	"github.com/coinbase-samples/core-go"
	"github.com/coinbase-samples/prime-sdk-go/client"
	"github.com/coinbase-samples/prime-sdk-go/model"
)

type GetEntityPositionsRequest struct {
	EntityId  string `json:"entity_id"`
	ProductId string `json:"product_id,omitempty"`
}

type GetEntityPositionsResponse struct {
	Positions         []*model.Position          `json:"positions"`
	ClearingAccountId string                     `json:"clearing_account_id"`
	Request           *GetEntityPositionsRequest `json:"request"`
}

func (s *futuresServiceImpl) GetEntityPositions(
	ctx context.Context,
	request *GetEntityPositionsRequest,
) (*GetEntityPositionsResponse, error) {

	path := fmt.Sprintf("/entities/%s/futures/positions", request.EntityId)

	queryParams := core.AppendHttpQueryParam(core.EmptyQueryParams, "product_id", request.ProductId)

	response := &GetEntityPositionsResponse{Request: request}

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
