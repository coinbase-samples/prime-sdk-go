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

type DescribeAssetsRequest struct {
	EntityId string `json:"entity_id"`
}

type DescribeAssetsResponse struct {
	Assets  []*Asset               `json:"assets"`
	Request *DescribeAssetsRequest `json:"request"`
}

func (c Client) DescribeAssets(
	ctx context.Context,
	request *DescribeAssetsRequest,
) (*DescribeAssetsResponse, error) {

	path := fmt.Sprintf("/entities/%s/assets", request.EntityId)

	response := &DescribeAssetsResponse{Request: request}

	if err := get(ctx, c, path, emptyQueryParams, request, response); err != nil {
		return response, fmt.Errorf("unable to DescribeAssets: %w", err)
	}

	return response, nil
}
