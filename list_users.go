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

type ListPortfolioUsersRequest struct {
	PortfolioId string            `json:"portfolio_id"`
	Pagination  *PaginationParams `json:"pagination_params"`
}

type ListPortfolioUsersResponse struct {
	Users      []*User                    `json:"users"`
	Request    *ListPortfolioUsersRequest `json:"request"`
	Pagination *Pagination                `json:"pagination"`
}

type ListUsersRequest struct {
	EntityId   string            `json:"entity_id"`
	Pagination *PaginationParams `json:"pagination_params"`
}

type ListUsersResponse struct {
	Users      []*User           `json:"users"`
	Request    *ListUsersRequest `json:"request"`
	Pagination *Pagination       `json:"pagination"`
}

func (c Client) ListUsers(
	ctx context.Context,
	request *ListUsersRequest,
) (*ListUsersResponse, error) {

	path := fmt.Sprintf("/entities/%s/users", request.EntityId)

	queryParams := appendPaginationParams(emptyQueryParams, request.Pagination)

	response := &ListUsersResponse{Request: request}

	if err := get(ctx, c, path, queryParams, request, response); err != nil {
		return nil, err
	}

	return response, nil
}

func (c Client) ListPortfolioUsers(
	ctx context.Context,
	request *ListPortfolioUsersRequest,
) (*ListPortfolioUsersResponse, error) {

	path := fmt.Sprintf("/portfolios/%s/users", request.PortfolioId)

	queryParams := appendPaginationParams(emptyQueryParams, request.Pagination)

	response := &ListPortfolioUsersResponse{Request: request}

	if err := get(ctx, c, path, queryParams, request, response); err != nil {
		return nil, err
	}

	return response, nil
}
