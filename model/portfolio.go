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

import "time"

// Portfolio represents a Prime portfolio
type Portfolio struct {
	Id             string `json:"id"`
	Name           string `json:"name"`
	EntityId       string `json:"entity_id"`
	OrganizationId string `json:"organization_id"`
}

// Counterparty represents a counterparty for a portfolio
type Counterparty struct {
	CounterpartyId string `json:"counterparty_id"`
}

// PostTradeCreditAmountDue represents an amount due for post trade credit
type PostTradeCreditAmountDue struct {
	Currency string    `json:"currency"`
	Amount   string    `json:"amount"`
	DueDate  time.Time `json:"due_date"`
}

// PostTradeCredit represents post trade credit information for a portfolio
type PostTradeCredit struct {
	Id                     string                      `json:"portfolio_id"`
	Currency               string                      `json:"currency"`
	Limit                  string                      `json:"limit"`
	Utilized               string                      `json:"utilized"`
	Available              string                      `json:"available"`
	Frozen                 bool                        `json:"frozen"`
	AmountsDue             []*PostTradeCreditAmountDue `json:"amounts_due"`
	FrozenReason           string                      `json:"frozen_reason"`
	Enabled                bool                        `json:"enabled"`
	AdjustedCreditUtilized string                      `json:"adjusted_credit_utilized"`
	AdjustedEquity         string                      `json:"adjusted_portfolio_equity"`
}
