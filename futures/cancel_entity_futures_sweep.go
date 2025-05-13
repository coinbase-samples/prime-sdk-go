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
)

type CancelEntityFuturesSweepRequest struct {
	EntityId string `json:"entity_id"`
}

type CancelEntityFuturesSweepResponse struct {
	Success   bool                             `json:"success"`
	RequestId string                           `json:"request_id"`
	Request   *CancelEntityFuturesSweepRequest `json:"-"`
}

func (s *futuresServiceImpl) CancelEntityFuturesSweep(
	ctx context.Context,
	request *CancelEntityFuturesSweepRequest,
) (*CancelEntityFuturesSweepResponse, error) {

	path := fmt.Sprintf("/entities/%s/futures/sweeps", request.EntityId)

	response := &CancelEntityFuturesSweepResponse{Request: request}

	if err := core.HttpDelete(
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
