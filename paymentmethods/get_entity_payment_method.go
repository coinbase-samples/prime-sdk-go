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

package paymentmethods

import (
	"context"
	"fmt"

	"github.com/coinbase-samples/core-go"
	"github.com/coinbase-samples/prime-sdk-go/client"
	"github.com/coinbase-samples/prime-sdk-go/model"
)

type GetEntityPaymentMethodRequest struct {
	Id              string `json:"entity_id"`
	PaymentMethodId string `json:"payment_method_id"`
}

type GetEntityPaymentMethodResponse struct {
	Details *model.EntityPaymentMethod     `json:"details"`
	Request *GetEntityPaymentMethodRequest `json:"request"`
}

func (s *paymentMethodsServiceImpl) GetEntityPaymentMethod(
	ctx context.Context,
	request *GetEntityPaymentMethodRequest,
) (*GetEntityPaymentMethodResponse, error) {

	path := fmt.Sprintf(
		"/entities/%s/payment-methods/%s",
		request.Id,
		request.PaymentMethodId,
	)

	response := &GetEntityPaymentMethodResponse{Request: request}

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
