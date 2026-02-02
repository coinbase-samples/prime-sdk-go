/**
 * Copyright 2026-present Coinbase Global, Inc.
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

type GetProductCandlesRequest struct {
	PortfolioId string                  `json:"portfolio_id"`
	ProductId   string                  `json:"product_id"`
	StartTime   time.Time               `json:"start_time"`
	EndTime     time.Time               `json:"end_time"`
	Granularity model.CandleGranularity `json:"granularity"`
}

type GetProductCandlesResponse struct {
	Candles []*model.Candle           `json:"candles"`
	Request *GetProductCandlesRequest `json:"-"`
}

func (s *productsServiceImpl) GetProductCandles(
	ctx context.Context,
	request *GetProductCandlesRequest,
) (*GetProductCandlesResponse, error) {

	path := fmt.Sprintf("/portfolios/%s/candles", request.PortfolioId)

	queryParams := core.AppendHttpQueryParam(core.EmptyQueryParams, "product_id", request.ProductId)
	queryParams = core.AppendHttpQueryParam(queryParams, "start_time", utils.TimeToStr(request.StartTime))
	queryParams = core.AppendHttpQueryParam(queryParams, "end_time", utils.TimeToStr(request.EndTime))
	queryParams = core.AppendHttpQueryParam(queryParams, "granularity", string(request.Granularity))

	response := &GetProductCandlesResponse{Request: request}

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
