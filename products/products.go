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

package products

import (
	"context"

	"github.com/coinbase-samples/prime-sdk-go/client"
	"github.com/coinbase-samples/prime-sdk-go/model"
)

type ProductsService interface {
	ListProducts(ctx context.Context, request *ListProductsRequest) (*ListProductsResponse, error)
	GetProductCandles(ctx context.Context, request *GetProductCandlesRequest) (*GetProductCandlesResponse, error)
	ServiceConfig() *model.ServiceConfig
}

func NewProductsService(c client.RestClient) ProductsService {
	return &productsServiceImpl{
		client:        c,
		serviceConfig: model.DefaultServiceConfig(),
	}
}

func NewProductsServiceWithConfig(c client.RestClient, config *model.ServiceConfig) ProductsService {
	return &productsServiceImpl{
		client:        c,
		serviceConfig: config,
	}
}

type productsServiceImpl struct {
	client        client.RestClient
	serviceConfig *model.ServiceConfig
}

func (s *productsServiceImpl) ServiceConfig() *model.ServiceConfig {
	return s.serviceConfig
}
