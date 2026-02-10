/**
 * Copyright 2026-present Coinbase Global, Inc.
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

// EditOrderRequest represents the request to edit an open order (Beta feature)
type EditOrderRequest struct {
	PortfolioId       string `json:"-"`
	OrderId           string `json:"-"`
	ProductId         string `json:"product_id,omitempty"`         // Deprecated
	OrigClientOrderId string `json:"orig_client_order_id"`         // Required: client order ID of order being edited
	ClientOrderId     string `json:"client_order_id"`              // Required: updated client order ID
	BaseQuantity      string `json:"base_quantity,omitempty"`      // Order size in base asset units
	QuoteValue        string `json:"quote_value,omitempty"`        // Order size in quote asset units
	LimitPrice        string `json:"limit_price,omitempty"`        // Limit price (required for TWAP, VWAP, LIMIT, STOP_LIMIT)
	ExpiryTime        string `json:"expiry_time,omitempty"`        // Expiry time in UTC (TWAP, VWAP, LIMIT, STOP_LIMIT GTD only)
	DisplayQuoteSize  string `json:"display_quote_size,omitempty"` // Display quote size for iceberg orders
}

// EditOrderResponse represents the response from editing an order
type EditOrderResponse struct {
	OrderId string            `json:"order_id"`
	Request *EditOrderRequest `json:"-"`
}

// EditOrder edits an open order. This feature is in beta.
func (s *ordersServiceImpl) EditOrder(
	ctx context.Context,
	request *EditOrderRequest,
) (*EditOrderResponse, error) {

	path := fmt.Sprintf("/portfolios/%s/orders/%s/edit", request.PortfolioId, request.OrderId)

	response := &EditOrderResponse{Request: request}

	if err := core.HttpPut(
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
