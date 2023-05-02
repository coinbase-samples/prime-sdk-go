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
	"fmt"
	"time"

	"github.com/shopspring/decimal"
)

const (
	WalletTypeVault   = "VAULT"
	WalletTypeTrading = "TRADING"
	WalletTypeOther   = "WALLET_TYPE_OTHER"

	BalanceTypeTrading = "TRADING_BALANCES"
	BalanceTypeVault   = "VAULT_BALANCES"
	BalanceTypeTotal   = "TOTAL_BALANCES"

	OrderSideBuy  = "BUY"
	OrderSideSell = "SELL"

	OrderTypeMarket = "MARKET"
	OrderTypeLimit  = "LIMIT"
	OrderTypeTwap   = "TWAP"
	OrderTypeBlock  = "BLOCK"

	TimeInForceGoodUntilTime      = "GOOD_UNTIL_DATE_TIME"
	TimeInForceGoodUntilCancelled = "GOOD_UNTIL_CANCELLED"
	TimeInForceImmediateOrCancel  = "IMMEDIATE_OR_CANCEL"
)

type ErrorMessage struct {
	Value string `json:"message"`
}

type AssetBalances struct {
	Symbol               string `json:"symbol"`
	Amount               string `json:"amount"`
	Holds                string `json:"holds"`
	BondedAmount         string `json:"bonded_amount"`
	ReservedAmount       string `json:"reserved_amount"`
	UnbondingAmount      string `json:"unbonding_amount"`
	UnvestedAmount       string `json:"unvested_amount"`
	PendingRewardsAmount string `json:"pending_rewards_amount"`
	PastRewardsAmount    string `json:"past_rewards_amount"`
	BondableAmount       string `json:"bondable_amount"`
	WithdrawableAmount   string `json:"withdrawable_amount"`
}

func (b AssetBalances) AmountNum() (amount decimal.Decimal, err error) {
	amount, err = strToNum(b.Amount)
	if err != nil {
		err = fmt.Errorf("Invalid asset amount: %s - symbol: %s - msg: %v", b.Amount, b.Symbol, err)
	}
	return
}

func (b AssetBalances) HoldsNum() (holds decimal.Decimal, err error) {
	holds, err = strToNum(b.Holds)
	if err != nil {
		err = fmt.Errorf("Invalid asset holds: %s - symbol: %s - msg: %v", b.Holds, b.Symbol, err)
	}
	return
}

type BalanceWithHolds struct {
	Total string `json:"total"`
	Holds string `json:"holds"`
}

type IteratorParams struct {
	Cursor        string `json:"cursor"`
	Limit         string `json:"limit"`
	SortDirection string `json:"sort_direction"`
}

func (r DescribeWalletsResponse) HasNext() bool {
	return r.Pagination != nil && r.Pagination.HasNext
}

type Wallet struct {
	Id        string    `json:"id"`
	Type      string    `json:"type"`
	Name      string    `json:"name"`
	Symbol    string    `json:"symbol"`
	CreatedAt time.Time `json:"created_at"`
}

type Pagination struct {
	NextCursor    string `json:"next_cursor"`
	SortDirection string `json:"sort_direction"`
	HasNext       bool   `json:"has_next"`
}

type PortfolioCommission struct {
	Type          string `json:"type"`
	Rate          string `json:"rate"`
	TradingVolume string `json:"trading_volume"`
}

func (p PortfolioCommission) RateNum() (rate decimal.Decimal, err error) {
	rate, err = strToNum(p.Rate)
	if err != nil {
		err = fmt.Errorf("Invalid commission rate: %s - err: %w", p.Rate, err)
	}
	return
}

type Product struct {
	Id             string   `json:"id"`
	BaseIncrement  string   `json:"base_increment"`
	QuoteIncrement string   `json:"quote_increment"`
	BaseMinSize    string   `json:"base_min_size"`
	BaseMaxSize    string   `json:"base_max_size"`
	QuoteMinSize   string   `json:"quote_min_size"`
	QuoteMaxSize   string   `json:"quote_max_size"`
	Permissions    []string `json:"permissions"`
}

func (p Product) BaseMinSizeNum() (amount decimal.Decimal, err error) {
	amount, err = strToNum(p.BaseMinSize)
	if err != nil {
		err = fmt.Errorf("invalid base min: %s - id: %s - err: %w", p.BaseMinSize, p.Id, err)
	}
	return
}

func (p Product) BaseMaxSizeNum() (amount decimal.Decimal, err error) {
	amount, err = strToNum(p.BaseMaxSize)
	if err != nil {
		err = fmt.Errorf("invalid base max: %s - id: %s - err: %v", p.BaseMaxSize, p.Id, err)
	}
	return
}

func (p Product) BaseIncrementNum() (amount decimal.Decimal, err error) {
	amount, err = strToNum(p.BaseIncrement)
	if err != nil {
		err = fmt.Errorf("invalid base increment: %s - id: %s - msg: %w", p.BaseIncrement, p.Id, err)
	}
	return
}

func (p Product) QuoteMinSizeNum() (amount decimal.Decimal, err error) {
	amount, err = strToNum(p.QuoteMinSize)
	if err != nil {
		err = fmt.Errorf("invalid quote min: %s - id: %s - err: %w", p.QuoteMinSize, p.Id, err)
	}
	return
}

func (p Product) QuoteMaxSizeNum() (amount decimal.Decimal, err error) {
	amount, err = strToNum(p.QuoteMaxSize)
	if err != nil {
		err = fmt.Errorf("invalid quote max: %s - id: %s - err: %v", p.QuoteMaxSize, p.Id, err)
	}
	return
}

func (p Product) QuoteIncrementNum() (amount decimal.Decimal, err error) {
	amount, err = strToNum(p.QuoteIncrement)
	if err != nil {
		err = fmt.Errorf("invalid quote increment: %s - id: %s - msg: %w", p.QuoteIncrement, p.Id, err)
	}
	return
}

func strToNum(v string) (amount decimal.Decimal, err error) {
	amount, err = decimal.NewFromString(v)
	return
}
