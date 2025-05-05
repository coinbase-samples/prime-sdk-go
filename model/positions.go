/**
 * Copyright 2025-present Coinbase Global, Inc.
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

type AggregationType string

const (
	AggregationTypeUnknown              AggregationType = "UNKNOWN_BALANCE_TYPE"
	AggregationTypeTradingBalances      AggregationType = "TRADING_BALANCES"
	AggregationTypeVaultBalances        AggregationType = "VAULT_BALANCES"
	AggregationTypeTotalBalances        AggregationType = "TOTAL_BALANCES"
	AggregationTypePrimeCustodyBalances AggregationType = "PRIME_CUSTODY_BALANCES"
	AggregationTypeUnifiedTotalBalances AggregationType = "UNIFIED_TOTAL_BALANCES"
)

type EntityBalance struct {
	Symbol        string `json:"symbol"`
	LongAmount    string `json:"long_amount"`
	LongNotional  string `json:"long_notional"`
	ShortAmount   string `json:"short_amount"`
	ShortNotional string `json:"short_notional"`
}

type EntityPositionReferenceType string

const (
	EntityPositionReferenceTypeUnspecified EntityPositionReferenceType = "POSITION_REFERENCE_TYPE_UNSPECIFIED"
	EntityPositionReferenceTypeEntity      EntityPositionReferenceType = "ENTITY"
	EntityPositionReferenceTypePortfolio   EntityPositionReferenceType = "PORTFOLIO"
)

type EntityPositionReference struct {
	Id   string                      `json:"id"`
	Type EntityPositionReferenceType `json:"type"`
}

type EntityPosition struct {
	Symbol            string                  `json:"symbol"`
	Long              string                  `json:"long"`
	Short             string                  `json:"short"`
	PositionReference EntityPositionReference `json:"position_reference"`
}
