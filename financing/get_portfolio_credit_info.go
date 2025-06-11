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

type GetPortfolioCreditInfoRequest struct {
	PortfolioId   string `json:"portfolio_id"`   // required
	BaseCurrency  string `json:"base_currency"`  // required
	QuoteCurrency string `json:"quote_currency"` // required
}

type GetPortfolioCreditInfoResponse struct {
	PortfolioCreditInfo *model.PostTradeCreditInfo     `json:"post_trade_credit"`
	Request             *GetPortfolioCreditInfoRequest `json:"-"`
}

func (s *financingServiceImpl) GetPortfolioCreditInfo(
	ctx context.Context,
	request *GetPortfolioCreditInfoRequest,
) (*GetPortfolioCreditInfoResponse, error) {

	path := fmt.Sprintf("/portfolios/%s/credit", request.PortfolioId)

	var queryParams string

	if request.BaseCurrency != "" {
		queryParams = core.AppendHttpQueryParam(queryParams, "base_currency", request.BaseCurrency)
	} else {
		return nil, fmt.Errorf("base_currency is required")
	}

	if request.QuoteCurrency != "" {
		queryParams = core.AppendHttpQueryParam(queryParams, "quote_currency", request.QuoteCurrency)
	} else {
		return nil, fmt.Errorf("quote_currency is required")
	}

	response := &GetPortfolioCreditInfoResponse{Request: request}

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
