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

package orders

import (
	"context"
	"fmt"

	"github.com/coinbase-samples/core-go"
	"github.com/coinbase-samples/prime-sdk-go/client"
)

type AcceptQuoteRequest struct {
	PortfolioId   string `json:"portfolio_id"`
	ProductId     string `json:"product_id"`
	Side          string `json:"side"`
	ClientOrderId string `json:"client_order_id"`
	QuoteId       string `json:"quote_id"`
}

type AcceptQuoteResponse struct {
	OrderId string              `json:"order_id"`
	Request *AcceptQuoteRequest `json:"-"`
}

func (s *ordersServiceImpl) AcceptQuote(ctx context.Context, request *AcceptQuoteRequest) (*AcceptQuoteResponse, error) {

	path := fmt.Sprintf("/portfolios/%s/accept_quote", request.PortfolioId)

	response := &AcceptQuoteResponse{Request: request}

	if err := core.HttpPost(
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
