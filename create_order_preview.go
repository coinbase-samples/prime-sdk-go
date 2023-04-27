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
)

type CreateOrderPreviewResponse struct {
	PortfolioId      string              `json:"portfolio_id"`
	Side             string              `json:"side"`
	ClientOrderId    string              `json:"client_order_id"`
	ProductId        string              `json:"product_id"`
	Type             string              `json:"type"`
	BaseQuantity     string              `json:"base_quantity"`
	QuoteValue       string              `json:"quote_value,omitempty"`
	LimitPrice       string              `json:"limit_price,omitempty"`
	StartTime        string              `json:"start_time,omitempty"`
	ExpiryTime       string              `json:"expiry_time,omitempty"`
	TimeInForce      string              `json:"time_in_force,omitempty"`
	StpId            string              `json:"stp_id,omitempty"`
	DisplayQuoteSize string              `json:"display_quote_size,omitempty"`
	DisplayBaseSize  string              `json:"display_base_size,omitempty"`
	IsRaiseExact     string              `json:"is_raise_exact,omitempty"`
	Commission       string              `json:"commission"`
	Slippage         string              `json:"slippage"`
	BestBid          string              `json:"best_bid"`
	BestAsk          string              `json:"best_ask"`
	AvgFillPrice     string              `json:"average_filled_price"`
	OrderTotal       string              `json:"order_total"`
	Request          *CreateOrderRequest `json:"request"`
}

func (c Client) CreateOrderPreview(
	ctx context.Context,
	request *CreateOrderRequest,
) (*CreateOrderPreviewResponse, error) {

	path := fmt.Sprintf("/portfolios/%s/order_preview", request.PortfolioId)

	response := &CreateOrderPreviewResponse{Request: request}

	if err := post(ctx, c, path, emptyQueryParams, request, response); err != nil {
		return nil, fmt.Errorf("unable to CreateOrderPreview: %w", err)
	}

	return response, nil
}
