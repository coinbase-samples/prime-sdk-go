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
)

type ListPortfoliosRequest struct{}

type ListPortfoliosResponse struct {
	Portfolios []*Portfolio           `json:"portfolios"`
	Request    *ListPortfoliosRequest `json:"request"`
}

func (c Client) ListPortfolios(
	ctx context.Context,
	request *ListPortfoliosRequest,
) (*ListPortfoliosResponse, error) {

	path := "/portfolios"

	response := &ListPortfoliosResponse{Request: request}

	if err := get(ctx, c, path, emptyQueryParams, request, response); err != nil {
		return nil, err
	}

	return response, nil
}
