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

	WalletDepositTypeWire   = "WIRE"
	WalletDepositTypeSwift  = "SWIFT"
	WalletDepositTypeCrypto = "CRYPTO"

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

type Portfolio struct {
	Id             string `json:"id"`
	Name           string `json:"name"`
	EntityId       string `json:"entity_id"`
	OrganizationId string `json:"organization_id"`
}

type WalletBalance struct {
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

func (b WalletBalance) AmountNum() (amount decimal.Decimal, err error) {
	amount, err = strToNum(b.Amount)
	if err != nil {
		err = fmt.Errorf("Invalid asset amount: %s - symbol: %s - msg: %v", b.Amount, b.Symbol, err)
	}
	return
}

func (b WalletBalance) HoldsNum() (holds decimal.Decimal, err error) {
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

type PaginationParams struct {
	Cursor        string `json:"cursor"`
	Limit         string `json:"limit"`
	SortDirection string `json:"sort_direction"`
}

func (r DescribeWalletsResponse) HasNext() bool {
	return r.Pagination != nil && r.Pagination.HasNext
}

type User struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	EntityId    string `json:"entity_id"`
	PortfolioId string `json:"portfolio_id,omitempty"`
	Role        string `json:"role"`
}

type Wallet struct {
	Id      string    `json:"id"`
	Type    string    `json:"type"`
	Name    string    `json:"name"`
	Symbol  string    `json:"symbol"`
	Created time.Time `json:"created_at"`
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

type Order struct {
	PortfolioId string `json:"portfolio_id"`
	Side        string `json:"side"`

	// A client-generated order ID used for reference purposes (note: order will be rejected if this ID
	// is not unique among all currently active orders)
	ClientOrderId string `json:"client_order_id"`
	ProductId     string `json:"product_id"`
	Type          string `json:"type"`

	// Order size in base asset units (either `base_quantity` or `quote_value` is required)
	BaseQuantity string `json:"base_quantity"`

	// Order size in quote asset units, i.e. the amount the user wants to spend (when buying) or receive (when selling);
	// the quantity in base units will be determined based on the market liquidity and indicated `quote_value` (either
	// `base_quantity` or `quote_value` is required)
	QuoteValue string `json:"quote_value,omitempty"`

	LimitPrice string `json:"limit_price,omitempty"`

	// The start time of the order in UTC (TWAP only)
	StartTime string `json:"start_time,omitempty"`

	// The expiry time of the order in UTC (TWAP and limit GTD only)
	ExpiryTime  string `json:"expiry_time,omitempty"`
	TimeInForce string `json:"time_in_force,omitempty"`

	// An optional self trade prevention id (in the form of a UUID). The value is only honored for certain
	// clients who are permitted to specify a custom self trade prevention id
	StpId string `json:"stp_id,omitempty"`

	// Optionally specify a display size. This is the maximum order size that will show up on venue order books.
	// Specifying a value here effectively makes a LIMIT order into an "iceberg" style order.
	// This property only applies to LIMIT orders and will be ignored for other order types.
	DisplayQuoteSize string `json:"display_quote_size,omitempty"`
	DisplayBaseSize  string `json:"display_base_size,omitempty"`

	// If you pass is_raise_exact = TRUE, you must use quote_value = n where n is the amount you want,
	// so $2000 will then cost you 1 ETH + fee, requiring > 1 ETH
	IsRaiseExact string `json:"is_raise_exact,omitempty"`

	// Used for describe order and create order preview
	Id                 string `json:"id,omitempty"`
	UserId             string `json:"user_id,omitempty"`
	Created            string `json:"created_at,omitempty"`
	FilledQuantity     string `json:"filled_quantity,omitempty"`
	FilledValue        string `json:"filled_value,omitempty"`
	AverageFilledPrice string `json:"average_filled_price,omitempty"`
	Commission         string `json:"commission,omitempty"`
	ExchangeFee        string `json:"exchange_fee,omitempty"`
	Total              string `json:"order_total,omitempty"`
	BestBid            string `json:"best_bid,omitempty"`
	BestAsk            string `json:"best_ask,omitempty"`
	Slippage           string `json:"slippage,omitempty"`
}

type Activity struct {
	Id              string           `json:"id"`
	ReferenceId     string           `json:"reference_id"`
	Category        string           `json:"category"`
	PrimaryType     string           `json:"type"`
	SecondaryType   string           `json:"secondary_type"`
	Status          string           `json:"status"`
	CreatedBy       string           `json:"created_by"`
	Title           string           `json:"title"`
	Description     string           `json:"description"`
	UserActions     []*UserAction    `json:"user_actions"`
	AccountMetadata *AccountMetadata `json:"account_metadata"`
	OrdersMetadata  *OrdersMetadata  `json:"orders_metadata"`
	Symbols         []string         `json:"symbols"`
	Created         string           `json:"created_at"`
	Updated         string           `json:"updated_at"`
}

type TransactionsMetadata struct {
	Consensus *Consensus `json:"consensus"`
}

type AccountMetadata struct {
	Consensus *Consensus `json:"consensus"`
}

// An empty/unimplemented/placeholder object in Prime
type OrdersMetadata struct{}

type Consensus struct {
	ApprovalDeadline string `json:"approval_deadline"`
	PassedConsensus  bool   `json:"has_passed_consensus"`
}

type UserAction struct {
	Action               string                `json:"action"`
	UserId               string                `json:"user_id"`
	Timestamp            string                `json:"timestamp"`
	TransactionsMetadata *TransactionsMetadata `json:"transactions_metadata"`
}
