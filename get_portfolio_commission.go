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
	"github.com/coinbase-samples/core-go"
)

type GetPortfolioCommissionRequest struct {
	PortfolioId string `json:"portfolio_id"`
}

type GetPortfolioCommissionResponse struct {
	Commission *Commission                    `json:"commission"`
	Request    *GetPortfolioCommissionRequest `json:"request"`
}

func (c *Client) GetPortfolioCommission(
	ctx context.Context,
	request *GetPortfolioCommissionRequest,
) (*GetPortfolioCommissionResponse, error) {

	path := fmt.Sprintf("/portfolios/%s/commission", request.PortfolioId)

	response := &GetPortfolioCommissionResponse{Request: request}

	if err := core.Get(ctx, c, path, core.EmptyQueryParams, request, response, addPrimeHeaders); err != nil {
		return nil, err
	}

	return response, nil
}
