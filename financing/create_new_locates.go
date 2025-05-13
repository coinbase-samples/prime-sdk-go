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
)

type CreateNewLocatesRequest struct {
	// required
	PortfolioId string `json:"portfolio_id"`
	// Currency symbol
	Symbol string `json:"symbol"`
	// Locate Amount
	Amount string `json:"amount"`
	// The target date of the locate (YYYY-MM-DD)
	LocateDate string `json:"locate_date"`
}

type CreateNewLocatesResponse struct {
	LocateId string                   `json:"locate_id"`
	Request  *CreateNewLocatesRequest `json:"-"`
}

func (s *financingServiceImpl) CreateNewLocates(
	ctx context.Context,
	request *CreateNewLocatesRequest,
) (*CreateNewLocatesResponse, error) {

	path := fmt.Sprintf("/portfolios/%s/locates", request.PortfolioId)

	var queryParams string

	response := &CreateNewLocatesResponse{Request: request}

	if err := core.HttpPost(
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
