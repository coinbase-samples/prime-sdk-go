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

type GetEntityPaymentMethodRequest struct {
	Id              string `json:"entity_id"`
	PaymentMethodId string `json:"payment_method_id"`
}

type GetEntityPaymentMethodResponse struct {
	Id                string                         `json:"id"`
	Symbol            string                         `json:"symbol"`
	PaymentMethodType string                         `json:"payment_method_type"`
	Name              string                         `json:"name"`
	AccountNumber     string                         `json:"account_number"`
	BankCode          string                         `json:"bank_code"`
	Request           *GetEntityPaymentMethodRequest `json:"request"`
}

func (c Client) GetEntityPaymentMethod(
	ctx context.Context,
	request *GetEntityPaymentMethodRequest,
) (*GetEntityPaymentMethodResponse, error) {

	path := fmt.Sprintf(
		"/entities/%s/payment-methods/%s",
		request.Id,
		request.PaymentMethodId,
	)

	response := &GetEntityPaymentMethodResponse{Request: request}

	if err := get(ctx, c, path, emptyQueryParams, request, response); err != nil {
		return nil, err
	}

	return response, nil
}
