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

package model

type AllocationLeg struct {
	LegId                  string `json:"allocation_leg_id"`
	DestinationPortfolioId string `json:"destination_portfolio_id"`
	Amount                 string `json:"amount"`
}

type AllocationDestination struct {
	LegId             string `json:"leg_id"`
	SourcePortfolioId string `json:"portfolio_id"`
	AllocationBase    string `json:"allocation_base"`
	AllocationQuote   string `json:"allocation_quote"`
	FeesAllocatedLeg  string `json:"fees_allocated_leg"`
}

type Allocation struct {
	RootId        string                   `json:"root_id"`
	ReversalId    string                   `json:"reversal_id"`
	Completed     string                   `json:"allocation_completed_at"`
	UserId        string                   `json:"user_id"`
	ProductId     string                   `json:"product_id"`
	Side          string                   `json:"side"`
	AvgPrice      string                   `json:"avg_price"`
	BaseQuantity  string                   `json:"base_quantity"`
	QuoteValue    string                   `json:"quote_value"`
	FeesAllocated string                   `json:"fees_allocated"`
	Status        string                   `json:"status"`
	Source        string                   `json:"source"`
	OrderIds      []string                 `json:"order_ids"`
	Destinations  []*AllocationDestination `json:"destinations"`
}
