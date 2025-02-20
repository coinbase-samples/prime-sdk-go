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
	"fmt"
	"time"

	"github.com/coinbase-samples/core-go"
	"github.com/coinbase-samples/prime-sdk-go/client"
	"github.com/coinbase-samples/prime-sdk-go/model"
	"github.com/coinbase-samples/prime-sdk-go/utils"
)

type ListPortfolioFillsRequest struct {
	PortfolioId string                  `json:"portfolio_id"` // required
	Start       time.Time               `json:"start_date"`   // required
	End         time.Time               `json:"end_date"`     // required
	Pagination  *model.PaginationParams `json:"pagination_params"`
}

type ListPortfolioFillsResponse struct {
	Fills      []*model.OrderFill         `json:"fills"`
	Pagination *model.Pagination          `json:"pagination"`
	Request    *ListPortfolioFillsRequest `json:"request"`
}

func (s *ordersServiceImpl) ListPortfolioFills(
	ctx context.Context,
	request *ListPortfolioFillsRequest,
) (*ListPortfolioFillsResponse, error) {

	path := fmt.Sprintf("/portfolios/%s/fills", request.PortfolioId)

	queryParams := utils.AppendPaginationParams(core.EmptyQueryParams, request.Pagination)

	queryParams = core.AppendHttpQueryParam(queryParams, "start_date", utils.TimeToStr(request.Start))

	if !request.End.IsZero() {
		queryParams = core.AppendHttpQueryParam(queryParams, "end_date", utils.TimeToStr(request.End))
	}

	response := &ListPortfolioFillsResponse{Request: request}

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
