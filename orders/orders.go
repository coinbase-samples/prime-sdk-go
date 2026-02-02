/**
 * Copyright 2024-present Coinbase Global, Inc.
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

	"github.com/coinbase-samples/prime-sdk-go/client"
	"github.com/coinbase-samples/prime-sdk-go/model"
)

type OrdersService interface {
	ListOpenOrders(ctx context.Context, request *ListOpenOrdersRequest) (*ListOpenOrdersResponse, error)
	CreateOrder(ctx context.Context, request *CreateOrderRequest) (*CreateOrderResponse, error)
	CreateOrderPreview(ctx context.Context, request *CreateOrderRequest) (*CreateOrderPreviewResponse, error)
	ListOrders(ctx context.Context, request *ListOrdersRequest) (*ListOrdersResponse, error)
	GetOrder(ctx context.Context, request *GetOrderRequest) (*GetOrderResponse, error)
	CancelOrder(ctx context.Context, request *CancelOrderRequest) (*CancelOrderResponse, error)
	EditOrder(ctx context.Context, request *EditOrderRequest) (*EditOrderResponse, error)
	GetOrderEditHistory(ctx context.Context, request *GetOrderEditHistoryRequest) (*GetOrderEditHistoryResponse, error)
	ListOrderFills(ctx context.Context, request *ListOrderFillsRequest) (*ListOrderFillsResponse, error)
	ListPortfolioFills(ctx context.Context, request *ListPortfolioFillsRequest) (*ListPortfolioFillsResponse, error)
	CreateQuoteRequest(ctx context.Context, request *CreateQuoteRequest) (*CreateQuoteResponse, error)
	AcceptQuote(ctx context.Context, request *AcceptQuoteRequest) (*AcceptQuoteResponse, error)
	ServiceConfig() *model.ServiceConfig
}

func NewOrdersService(c client.RestClient) OrdersService {
	return &ordersServiceImpl{
		client:        c,
		serviceConfig: model.DefaultServiceConfig(),
	}
}

func NewOrdersServiceWithConfig(c client.RestClient, config *model.ServiceConfig) OrdersService {
	return &ordersServiceImpl{
		client:        c,
		serviceConfig: config,
	}
}

type ordersServiceImpl struct {
	client        client.RestClient
	serviceConfig *model.ServiceConfig
}

func (s *ordersServiceImpl) ServiceConfig() *model.ServiceConfig {
	return s.serviceConfig
}
