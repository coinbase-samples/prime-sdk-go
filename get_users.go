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

type GetPortfolioUsersRequest struct {
	PortfolioId string            `json:"portfolio_id"`
	Pagination  *PaginationParams `json:"pagination_params"`
}

type GetPortfolioUsersResponse struct {
	Users      []*User                        `json:"users"`
	Request    *GetPortfolioUsersRequest `json:"request"`
	Pagination *Pagination                    `json:"pagination"`
}

type GetUsersRequest struct {
	EntityId   string            `json:"entity_id"`
	Pagination *PaginationParams `json:"pagination_params"`
}

type GetUsersResponse struct {
	Users      []*User               `json:"users"`
	Request    *GetUsersRequest `json:"request"`
	Pagination *Pagination           `json:"pagination"`
}

func (c Client) GetUsers(
	ctx context.Context,
	request *GetUsersRequest,
) (*GetUsersResponse, error) {

	path := fmt.Sprintf("/entities/%s/users", request.EntityId)

	queryParams := appendPaginationParams(emptyQueryParams, request.Pagination)

	response := &GetUsersResponse{Request: request}

	if err := get(ctx, c, path, queryParams, request, response); err != nil {
		return nil, err
	}

	return response, nil
}

func (c Client) GetPortfolioUsers(
	ctx context.Context,
	request *GetPortfolioUsersRequest,
) (*GetPortfolioUsersResponse, error) {

	path := fmt.Sprintf("/portfolios/%s/users", request.PortfolioId)

	queryParams := appendPaginationParams(emptyQueryParams, request.Pagination)

	response := &GetPortfolioUsersResponse{Request: request}

	if err := get(ctx, c, path, queryParams, request, response); err != nil {
		return nil, err
	}

	return response, nil
}