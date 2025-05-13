/**
 * Copyright 2025-present Coinbase Global, Inc.
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

package financing

import (
	"context"
	"fmt"

	"github.com/coinbase-samples/core-go"
	"github.com/coinbase-samples/prime-sdk-go/client"
	"github.com/coinbase-samples/prime-sdk-go/model"
)

type GetBuyingPowerRequest struct {
	// required
	PortfolioId string `json:"portfolio_id"`
	// The symbol for the base currency, required
	BaseCurrency string `json:"base_currency"`
	// The symbol for the quote currency, required
	QuoteCurrency string `json:"quote_currency"`
}

type GetBuyingPowerResponse struct {
	BuyingPower model.BuyingPower      `json:"buying_power"`
	Request     *GetBuyingPowerRequest `json:"-"`
}

func (s *financingServiceImpl) GetBuyingPower(
	ctx context.Context,
	request *GetBuyingPowerRequest,
) (*GetBuyingPowerResponse, error) {

	path := fmt.Sprintf("/portfolios/%s/buying_power", request.PortfolioId)

	var queryParams string

	if request.BaseCurrency != "" {
		queryParams = core.AppendHttpQueryParam(queryParams, "base_currency", request.BaseCurrency)
	}

	if request.QuoteCurrency != "" {
		queryParams = core.AppendHttpQueryParam(queryParams, "quote_currency", request.QuoteCurrency)
	}

	response := &GetBuyingPowerResponse{Request: request}

	if err := core.HttpGet(
		ctx,
		s.client,
		path,
		queryParams,
		client.DefaultSuccessHttpStatusCodes,
		request,
		response,
		s.client.HeadersFunc(),
	); err != nil {
		return nil, err
	}

	return response, nil
}
