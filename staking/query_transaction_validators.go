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

package staking

import (
	"context"
	"fmt"

	"github.com/coinbase-samples/core-go"
	"github.com/coinbase-samples/prime-sdk-go/client"
	"github.com/coinbase-samples/prime-sdk-go/model"
)

type QueryTransactionValidatorsRequest struct {
	PortfolioId    string   `json:"-"`
	TransactionIds []string `json:"transaction_ids"`
	Cursor         string   `json:"cursor,omitempty"`
	Limit          int32    `json:"limit,omitempty"`
	SortDirection  string   `json:"sort_direction,omitempty"`
}

type QueryTransactionValidatorsResponse struct {
	model.PaginationMixin
	TransactionValidators []*model.TransactionValidator      `json:"transaction_validators"`
	Request               *QueryTransactionValidatorsRequest `json:"-"`
	service               StakingService
	serviceConfig         *model.ServiceConfig
}

func (r *QueryTransactionValidatorsResponse) Next(ctx context.Context) (*QueryTransactionValidatorsResponse, error) {
	if !r.HasNext() {
		return nil, nil
	}

	nextReq := *r.Request
	nextReq.Cursor = r.GetNextCursor()

	return r.service.QueryTransactionValidators(ctx, &nextReq)
}

func (r *QueryTransactionValidatorsResponse) Iterator() *model.PageIterator[*QueryTransactionValidatorsResponse, *model.TransactionValidator] {
	return model.NewPageIteratorWithConfig(
		r,
		func(resp *QueryTransactionValidatorsResponse) []*model.TransactionValidator {
			return resp.TransactionValidators
		},
		r.serviceConfig,
	)
}

func (s *stakingServiceImpl) QueryTransactionValidators(
	ctx context.Context,
	request *QueryTransactionValidatorsRequest,
) (*QueryTransactionValidatorsResponse, error) {

	path := fmt.Sprintf("/portfolios/%s/staking/transaction-validators/query", request.PortfolioId)

	if request.Limit == 0 && s.serviceConfig != nil && s.serviceConfig.DefaultLimit > 0 {
		request.Limit = s.serviceConfig.DefaultLimit
	}

	response := &QueryTransactionValidatorsResponse{
		Request:       request,
		service:       s,
		serviceConfig: s.serviceConfig,
	}

	if err := core.HttpPost(
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
