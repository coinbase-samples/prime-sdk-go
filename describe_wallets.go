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

type DescribeWalletsRequest struct {
	PortfolioId    string          `json:"string"`
	Type           string          `json:"type"`
	Symbols        []string        `json:"symbols"`
	IteratorParams *IteratorParams `json:"iteratorParams"`
}

type DescribeWalletsResponse struct {
	Wallets    []*Wallet               `json:"wallets"`
	Request    *DescribeWalletsRequest `json:"request"`
	Pagination *Pagination             `json:"pagination"`
}

func DescribeWallets(
	ctx context.Context,
	request *DescribeWalletsRequest,
) (*DescribeWalletsResponse, error) {
	return DefaultClient.DescribeWallets(ctx, request)
}

func (c Client) DescribeWallets(
	ctx context.Context,
	request *DescribeWalletsRequest,
) (*DescribeWalletsResponse, error) {

	url := fmt.Sprintf("%s/portfolios/%s/wallets?type=%s",
		primeV1ApiBaseUrl,
		request.PortfolioId,
		request.Type,
	)

	url = urlIteratorParams(url, request.IteratorParams)

	for _, v := range request.Symbols {
		url += fmt.Sprintf("&symbols=%s", v)
	}

	response := &DescribeWalletsResponse{Request: request}

	if err := get(ctx, c, url, request, response); err != nil {
		return nil, err
	}

	return response, nil
}
