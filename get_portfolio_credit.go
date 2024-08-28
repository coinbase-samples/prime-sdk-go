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

type GetPortfolioCreditRequest struct {
	Id string `json:"portfolio_id"`
}

type GetPortfolioCreditResponse struct {
	PostTradeCredit *PostTradeCredit           `json:"post_trade_credit"`
	Request         *GetPortfolioCreditRequest `json:"request"`
}

func (c *Client) GetPortfolioCredit(
	ctx context.Context,
	request *GetPortfolioCreditRequest,
) (*GetPortfolioCreditResponse, error) {

	path := fmt.Sprintf("/portfolios/%s/credit", request.Id)

	response := &GetPortfolioCreditResponse{Request: request}

	if err := core.Get(ctx, c, path, core.EmptyQueryParams, request, response, c.headersFunc); err != nil {
		return nil, err
	}

	return response, nil
}
