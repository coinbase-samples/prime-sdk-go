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

package paymentmethods

import (
	"context"
	"fmt"

	"github.com/coinbase-samples/core-go"
	"github.com/coinbase-samples/prime-sdk-go/client"
	"github.com/coinbase-samples/prime-sdk-go/model"
)

type ListEntityPaymentMethodsRequest struct {
	EntityId   string                  `json:"entity_id"`
	Pagination *model.PaginationParams `json:"pagination_params"`
}

type ListEntityPaymentMethodsResponse struct {
	PaymentMethods []*model.EntityPaymentMethod     `json:"payment_methods"`
	Request        *ListEntityPaymentMethodsRequest `json:"request"`
}

func (s *paymentMethodsServiceImpl) ListEntityPaymentMethods(
	ctx context.Context,
	request *ListEntityPaymentMethodsRequest,
) (*ListEntityPaymentMethodsResponse, error) {

	path := fmt.Sprintf("/entities/%s/payment-methods", request.EntityId)

	response := &ListEntityPaymentMethodsResponse{Request: request}

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
