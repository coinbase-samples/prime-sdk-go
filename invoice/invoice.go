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

package invoice

import (
	"context"

	"github.com/coinbase-samples/prime-sdk-go/client"
	"github.com/coinbase-samples/prime-sdk-go/model"
)

type InvoiceService interface {
	ListInvoices(ctx context.Context, request *ListInvoicesRequest) (*ListInvoicesResponse, error)
	ServiceConfig() *model.ServiceConfig
}

func NewInvoiceService(c client.RestClient) InvoiceService {
	return &invoiceServiceImpl{
		client:        c,
		serviceConfig: model.DefaultServiceConfig(),
	}
}

func NewInvoiceServiceWithConfig(c client.RestClient, config *model.ServiceConfig) InvoiceService {
	return &invoiceServiceImpl{
		client:        c,
		serviceConfig: config,
	}
}

type invoiceServiceImpl struct {
	client        client.RestClient
	serviceConfig *model.ServiceConfig
}

func (s *invoiceServiceImpl) ServiceConfig() *model.ServiceConfig {
	return s.serviceConfig
}
