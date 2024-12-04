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

package allocations

import (
	"context"
	"fmt"

	"github.com/coinbase-samples/core-go"
	"github.com/coinbase-samples/prime-sdk-go/client"
	"github.com/coinbase-samples/prime-sdk-go/model"
)

type GetPortfolioNetAllocationRequest struct {
	PortfolioId string `json:"portfolio_id"`
	NettingId   string `json:"netting_id"`
}

type GetPortfolioNetAllocationResponse struct {
	Allocations []*model.Allocation               `json:"allocations"`
	Request     *GetPortfolioNetAllocationRequest `json:"request"`
}

func (s *allocationsServiceImpl) GetPortfolioNetAllocation(
	ctx context.Context,
	request *GetPortfolioNetAllocationRequest,
) (*GetPortfolioNetAllocationResponse, error) {

	path := fmt.Sprintf(
		"/portfolios/%s/allocations/net/%s",
		request.PortfolioId,
		request.NettingId,
	)

	response := &GetPortfolioNetAllocationResponse{Request: request}

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
