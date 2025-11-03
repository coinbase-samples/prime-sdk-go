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

package products

import (
	"context"
	"fmt"
	"time"

	"github.com/coinbase-samples/core-go"
	"github.com/coinbase-samples/prime-sdk-go/client"
	"github.com/coinbase-samples/prime-sdk-go/model"
	"github.com/coinbase-samples/prime-sdk-go/utils"
)

type GetCandlesRequest struct {
	PortfolioId string                   `json:"portfolio_id"`
	ProductId   string                   `json:"product_id"`
	Granularity model.CandlesGranularity `json:"granularity"`
	Start       time.Time                `json:"start_time"`
	End         time.Time                `json:"end_time"`
}

type GetCandlesResponse struct {
	Candles []*model.Candle    `json:"candles"`
	Request *GetCandlesRequest `json:"-"`
}

func (s *productsServiceImpl) GetCandles(
	ctx context.Context,
	request *GetCandlesRequest,
) (*GetCandlesResponse, error) {

	path := fmt.Sprintf("/portfolios/%s/candles", request.PortfolioId)

	var queryParams string

	queryParams = core.AppendHttpQueryParam(queryParams, "product_id", request.ProductId)

	queryParams = core.AppendHttpQueryParam(queryParams, "granularity", string(request.Granularity))

	if !request.Start.IsZero() {
		queryParams = core.AppendHttpQueryParam(queryParams, "start_time", utils.TimeToStr(request.Start))
	}

	if !request.End.IsZero() {
		queryParams = core.AppendHttpQueryParam(queryParams, "end_time", utils.TimeToStr(request.End))
	}

	response := &GetCandlesResponse{Request: request}

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
