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
	"github.com/coinbase-samples/prime-sdk-go/model"
)

type CreateQuoteRequest struct {
	PortfolioId   string          `json:"portfolio_id"`
	ProductId     string          `json:"product_id"`
	Side          model.OrderSide `json:"side"`
	ClientQuoteId string          `json:"client_quote_id"`
	BaseQuantity  string          `json:"base_quantity,omitempty"`
	QuoteValue    string          `json:"quote_value,omitempty"`
	LimitPrice    string          `json:"limit_price"`
}

type CreateQuoteResponse struct {
	QuoteId              string              `json:"quote_id"`
	ExpirationTime       string              `json:"expiration_time"`
	BestPrice            string              `json:"best_price"`
	OrderTotal           string              `json:"order_total"`
	PriceInclusiveOfFees string              `json:"price_inclusive_of_fees"`
	Request              *CreateQuoteRequest `json:"request"`
}

func (s *ordersServiceImpl) CreateQuoteRequest(ctx context.Context, request *CreateQuoteRequest) (*CreateQuoteResponse, error) {

	path := fmt.Sprintf("/portfolios/%s/rfq", request.PortfolioId)

	response := &CreateQuoteResponse{Request: request}

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
