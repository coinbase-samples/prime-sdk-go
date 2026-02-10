/**
 * Copyright 2024-present Coinbase Global, Inc.
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

package portfolios

import (
	"context"

	"github.com/coinbase-samples/prime-sdk-go/client"
)

type PortfoliosService interface {
	ListPortfolios(ctx context.Context, request *ListPortfoliosRequest) (*ListPortfoliosResponse, error)
	GetPortfolio(ctx context.Context, request *GetPortfolioRequest) (*GetPortfolioResponse, error)
	GetPortfolioCredit(ctx context.Context, request *GetPortfolioCreditRequest) (*GetPortfolioCreditResponse, error)
	GetPortfolioCounterparty(ctx context.Context, request *GetPortfolioCounterpartyRequest) (*GetPortfolioCounterpartyResponse, error)
}

func NewPortfoliosService(c client.RestClient) PortfoliosService {
	return &portfoliosServiceImpl{client: c}
}

type portfoliosServiceImpl struct {
	client client.RestClient
}
