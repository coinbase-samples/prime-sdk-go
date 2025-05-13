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

package orders

import (
	"context"
	"errors"
	"fmt"

	"github.com/coinbase-samples/core-go"
	"github.com/coinbase-samples/prime-sdk-go/client"
	"github.com/coinbase-samples/prime-sdk-go/model"
)

type CreateOrderPreviewResponse struct {
	Order   *model.Order        `json:"order"`
	Request *CreateOrderRequest `json:"-"`
}

func (s *ordersServiceImpl) CreateOrderPreview(
	ctx context.Context,
	request *CreateOrderRequest,
) (*CreateOrderPreviewResponse, error) {

	if request.Order == nil {
		return nil, errors.New("order not set on request")
	}

	path := fmt.Sprintf("/portfolios/%s/order_preview", request.Order.PortfolioId)

	response := &CreateOrderPreviewResponse{Request: request}

	responseOrder := &model.Order{}

	if err := core.HttpPost(
		ctx,
		s.client,
		path,
		core.EmptyQueryParams,
		client.DefaultSuccessHttpStatusCodes,
		request.Order,
		responseOrder,
		s.client.HeadersFunc(),
	); err != nil {
		return nil, err
	}

	response.Order = responseOrder

	return response, nil
}
