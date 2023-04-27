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

type CreateOrderRequest struct {
	PortfolioId      string `json:"portfolio_id"`
	Side             string `json:"side"`
	ClientOrderId    string `json:"client_order_id"`
	ProductId        string `json:"product_id"`
	Type             string `json:"type"`
	BaseQuantity     string `json:"base_quantity"`
	QuoteValue       string `json:"quote_value,omitempty"`
	LimitPrice       string `json:"limit_price,omitempty"`
	StartTime        string `json:"start_time,omitempty"`
	ExpiryTime       string `json:"expiry_time,omitempty"`
	TimeInForce      string `json:"time_in_force,omitempty"`
	StpId            string `json:"stp_id,omitempty"`
	DisplayQuoteSize string `json:"display_quote_size,omitempty"`
	DisplayBaseSize  string `json:"display_base_size,omitempty"`
	IsRaiseExact     string `json:"is_raise_exact,omitempty"`
}

type CreateOrderResponse struct {
	OrderId string              `json:"order_id"`
	Request *CreateOrderRequest `json:"request"`
}

func CreateOrder(ctx context.Context, request *CreateOrderRequest) (*CreateOrderResponse, error) {
	return DefaultClient.CreateOrder(ctx, request)
}

func (c Client) CreateOrder(ctx context.Context, request *CreateOrderRequest) (*CreateOrderResponse, error) {

	url := fmt.Sprintf("%s/portfolios/%s/order", primeV1ApiBaseUrl, request.PortfolioId)

	response := &CreateOrderResponse{Request: request}

	if err := post(ctx, c, url, request, response); err != nil {
		return nil, err
	}

	return response, nil
}
