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

package balances

import (
	"context"
	"fmt"

	"github.com/coinbase-samples/core-go"
	"github.com/coinbase-samples/prime-sdk-go/client"
	"github.com/coinbase-samples/prime-sdk-go/model"
	"github.com/coinbase-samples/prime-sdk-go/utils"
)

type ListEntityBalancesRequest struct {
	EntityId        string                  `json:"entity_id"`
	Symbols         []string                `json:"symbols,omitempty"`
	Pagination      *model.PaginationParams `json:"pagination_params"`
	AggregationType model.AggregationType   `json:"aggregation_type,omitempty"`
}

type ListEntityBalancesResponse struct {
	Balances   []*model.EntityBalance     `json:"balances"`
	Pagination *model.Pagination          `json:"pagination"`
	Request    *ListEntityBalancesRequest `json:"-"`
}

func (s *balancesServiceImpl) ListEntityBalances(
	ctx context.Context,
	request *ListEntityBalancesRequest,
) (*ListEntityBalancesResponse, error) {

	path := fmt.Sprintf("/entities/%s/balances", request.EntityId)

	var queryParams string

	queryParams = utils.AppendPaginationParams(queryParams, request.Pagination)

	for _, v := range request.Symbols {
		queryParams = core.AppendHttpQueryParam(queryParams, "symbols", v)
	}

	if request.AggregationType != "" && request.AggregationType != model.AggregationTypeUnknown {
		queryParams = core.AppendHttpQueryParam(queryParams, "aggregation_type", string(request.AggregationType))
	}

	response := &ListEntityBalancesResponse{Request: request}

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
