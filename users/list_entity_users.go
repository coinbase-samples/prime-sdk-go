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

package users

import (
	"context"
	"fmt"

	"github.com/coinbase-samples/core-go"
	"github.com/coinbase-samples/prime-sdk-go/client"
	"github.com/coinbase-samples/prime-sdk-go/model"
	"github.com/coinbase-samples/prime-sdk-go/utils"
)

type ListEntityUsersRequest struct {
	EntityId   string                  `json:"entity_id"`
	Pagination *model.PaginationParams `json:"pagination_params"`
}

type ListEntityUsersResponse struct {
	Users      []*model.User           `json:"users"`
	Pagination *model.Pagination       `json:"pagination"`
	Request    *ListEntityUsersRequest `json:"-"`
}

func (s *usersServiceImpl) ListEntityUsers(
	ctx context.Context,
	request *ListEntityUsersRequest,
) (*ListEntityUsersResponse, error) {

	path := fmt.Sprintf("/entities/%s/users", request.EntityId)

	queryParams := utils.AppendPaginationParams(core.EmptyQueryParams, request.Pagination)

	response := &ListEntityUsersResponse{Request: request}

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
