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
	"time"
)

type DescribePortfolioCreditRequest struct {
	PortfolioId string `json:"portfolio_id"`
}

type DescribePortfolioCreditResponse struct {
	PostTradeCredit *PortfolioPostTradeCredit       `json:"post_trade_credit"`
	Request         *DescribePortfolioCreditRequest `json:"request"`
}

type PortfolioPostTradeCreditAmountDue struct {
	Currency string    `json:"currency"`
	Amount   string    `json:"amount"`
	DueDate  time.Time `json:"due_date"`
}

type PortfolioPostTradeCredit struct {
	Id                      string                               `json:"portfolio_id"`
	Currency                string                               `json:"currency"`
	Limit                   string                               `json:"limit"`
	Utilized                string                               `json:"utilized"`
	Available               string                               `json:"available"`
	Frozen                  bool                                 `json:"frozen"`
	AmountsDue              []*PortfolioPostTradeCreditAmountDue `json:"amounts_due"`
	FrozenReason            string                               `json:"frozen_reason"`
	Enabled                 bool                                 `json:"enabled"`
	AdjustedCreditUtilized  string                               `json:"adjusted_credit_utilized"`
	AdjustedPortfolioEquity string                               `json:"adjusted_portfolio_equity"`
}

func (c Client) DescribePortfolioCredit(
	ctx context.Context,
	request *DescribePortfolioCreditRequest,
) (*DescribePortfolioCreditResponse, error) {

	path := fmt.Sprintf("/portfolios/%s/credit", request.PortfolioId)

	response := &DescribePortfolioCreditResponse{Request: request}

	if err := get(ctx, c, path, emptyQueryParams, request, response); err != nil {
		return nil, err
	}

	return response, nil
}
