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
	"github.com/coinbase-samples/prime-sdk-go/model"
)

// GetOrderEditHistoryRequest represents the request to get order edit history
type GetOrderEditHistoryRequest struct {
	PortfolioId string `json:"portfolio_id"`
	OrderId     string `json:"order_id"`
}

// GetOrderEditHistoryResponse represents the response containing order edit history
type GetOrderEditHistoryResponse struct {
	OrderId string `json:"order_id"`
	// Deprecated: Use EditHistory instead
	OrderEditHistory []*model.OrderEditHistory   `json:"order_edit_history,omitempty"`
	EditHistory      []*model.EditHistory        `json:"edit_history"`
	Request          *GetOrderEditHistoryRequest `json:"-"`
}

// GetOrderEditHistory retrieves the edit history for a specific order
func (s *ordersServiceImpl) GetOrderEditHistory(
	ctx context.Context,
	request *GetOrderEditHistoryRequest,
) (*GetOrderEditHistoryResponse, error) {

	path := fmt.Sprintf("/portfolios/%s/orders/%s/edit_history", request.PortfolioId, request.OrderId)

	response := &GetOrderEditHistoryResponse{Request: request}

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
