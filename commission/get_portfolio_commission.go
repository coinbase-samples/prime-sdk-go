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

package commission

import (
	"context"
	"fmt"

	"github.com/coinbase-samples/core-go"
	"github.com/coinbase-samples/prime-sdk-go/client"
	"github.com/coinbase-samples/prime-sdk-go/model"
)

type GetPortfolioCommissionRequest struct {
	PortfolioId string `json:"portfolio_id"`
	ProductId   string `json:"product_id,omitempty"`
}

type GetPortfolioCommissionResponse struct {
	Commission *model.Commission              `json:"commission"`
	Request    *GetPortfolioCommissionRequest `json:"-"`
}

func (s *commissionServiceImpl) GetPortfolioCommission(
	ctx context.Context,
	request *GetPortfolioCommissionRequest,
) (*GetPortfolioCommissionResponse, error) {

	path := fmt.Sprintf("/portfolios/%s/commission", request.PortfolioId)

	queryParams := core.EmptyQueryParams
	if request.ProductId != "" {
		queryParams = core.AppendHttpQueryParam(queryParams, "product_id", request.ProductId)
	}

	response := &GetPortfolioCommissionResponse{Request: request}

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
