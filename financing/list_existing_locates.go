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

type ListExistingLocatesRequest struct {
	PortfolioId string   `json:"portfolio_id"` // required
	LocateIds   []string `json:"locate_ids"`
	LocateDate  string   `json:"locate_date"`
}

type ListExistingLocatesResponse struct {
	Locates []model.ExistingLocate      `json:"locates"`
	Request *ListExistingLocatesRequest `json:"-"`
}

func (s *financingServiceImpl) ListExistingLocates(
	ctx context.Context,
	request *ListExistingLocatesRequest,
) (*ListExistingLocatesResponse, error) {

	path := fmt.Sprintf("/portfolios/%s/locates", request.PortfolioId)

	var queryParams string

	if request.LocateDate != "" {
		queryParams = core.AppendHttpQueryParam(queryParams, "locate_date", request.LocateDate)
	}

	for _, v := range request.LocateIds {
		queryParams = core.AppendHttpQueryParam(queryParams, "locate_ids", v)
	}

	response := &ListExistingLocatesResponse{Request: request}

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
