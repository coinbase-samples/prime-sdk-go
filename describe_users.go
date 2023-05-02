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

type DescribePortfolioUsersRequest struct {
	PortfolioId    string          `json:"portfolio_id"`
	IteratorParams *IteratorParams `json:"iteratorParams"`
}

type DescribePortfolioUsersResponse struct {
	Users      []*User                        `json:"users"`
	Request    *DescribePortfolioUsersRequest `json:"request"`
	Pagination *Pagination                    `json:"pagination"`
}

type DescribeUsersRequest struct {
	EntityId       string          `json:"entity_id"`
	IteratorParams *IteratorParams `json:"iteratorParams"`
}

type DescribeUsersResponse struct {
	Users      []*User               `json:"users"`
	Request    *DescribeUsersRequest `json:"request"`
	Pagination *Pagination           `json:"pagination"`
}

func (c Client) DescribeUsers(
	ctx context.Context,
	request *DescribeUsersRequest,
) (*DescribeUsersResponse, error) {

	path := fmt.Sprintf("/entities/%s/users", request.EntityId)

	queryParams := iteratorParams(emptyQueryParams, request.IteratorParams)

	response := &DescribeUsersResponse{Request: request}

	if err := get(ctx, c, path, queryParams, request, response); err != nil {
		return nil, err
	}

	return response, nil
}

func (c Client) DescribePortfolioUsers(
	ctx context.Context,
	request *DescribePortfolioUsersRequest,
) (*DescribePortfolioUsersResponse, error) {

	path := fmt.Sprintf("/portfolios/%s/users", request.PortfolioId)

	queryParams := iteratorParams(emptyQueryParams, request.IteratorParams)

	response := &DescribePortfolioUsersResponse{Request: request}

	if err := get(ctx, c, path, queryParams, request, response); err != nil {
		return nil, err
	}

	return response, nil
}
