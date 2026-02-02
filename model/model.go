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

package model

import (
	"fmt"
	"time"

	"github.com/coinbase-samples/core-go"
	"github.com/shopspring/decimal"
)

type WalletVisibility string

const (
	WalletVisibilityUnspecified WalletVisibility = "WALLET_VISIBILITY_UNSPECIFIED"
	WalletVisibilityVisible     WalletVisibility = "WALLET_VISIBILITY_VISIBLE"
	WalletVisibilityHidden      WalletVisibility = "WALLET_VISIBILITY_HIDDEN"
)

const (
	WalletTypeVault   = "VAULT"
	WalletTypeTrading = "TRADING"
	WalletTypeOnchain = "ONCHAIN"
	WalletTypeOther   = "WALLET_TYPE_OTHER"

	WalletDepositTypeWire   = "WIRE"
	WalletDepositTypeSwift  = "SWIFT"
	WalletDepositTypeCrypto = "CRYPTO"

	OrderTypeMarket = "MARKET"
	OrderTypeLimit  = "LIMIT"
	OrderTypeTwap   = "TWAP"
	OrderTypeBlock  = "BLOCK"

	TimeInForceGoodUntilTime      = "GOOD_UNTIL_DATE_TIME"
	TimeInForceGoodUntilCancelled = "GOOD_UNTIL_CANCELLED"
	TimeInForceImmediateOrCancel  = "IMMEDIATE_OR_CANCEL"
)

const (
	EvmNetworkFamily    = "NETWORK_FAMILY_EVM"
	SolanaNetworkFamily = "NETWORK_FAMILY_SOLANA"
)

type OrderSide string

const (
	OrderSideBuy     OrderSide = "BUY"
	OrderSideSell    OrderSide = "SELL"
	OrderSideUnknown OrderSide = "UNKNOWN_ORDER_SIDE"
)

type ErrorMessage struct {
	Value string `json:"message"`
}

type EntityPaymentMethod struct {
	Id                string `json:"id"`
	Symbol            string `json:"symbol"`
	PaymentMethodType string `json:"payment_method_type"`
	Name              string `json:"name"`
	AccountNumber     string `json:"account_number"`
	BankCode          string `json:"bank_code"`
	BankName          string `json:"bank_name,omitempty"`
	BankName2         string `json:"bank_name_2,omitempty"`
}

type Portfolio struct {
	Id             string `json:"id"`
	Name           string `json:"name"`
	EntityId       string `json:"entity_id"`
	OrganizationId string `json:"organization_id"`
}

type Counterparty struct {
	CounterpartyId string `json:"counterparty_id"`
}

type PaginationParams struct {
	Cursor        string `json:"cursor"`
	Limit         int32  `json:"limit"`
	SortDirection string `json:"sort_direction"`
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
	Id         string           `json:"id"`
	Type       string           `json:"type"`
	Name       string           `json:"name"`
	Address    string           `json:"address"`
	Visibility WalletVisibility `json:"visibility"`
	Symbol     string           `json:"symbol"`
	Created    time.Time        `json:"created_at"`
	Network    *NetworkDetails  `json:"network"`
}

type BlockchainAddress struct {
	Address           string          `json:"address"`
	AccountIdentifier string          `json:"account_identifier"`
	Network           *NetworkDetails `json:"network"`
}

type Pagination struct {
	NextCursor    string `json:"next_cursor"`
	SortDirection string `json:"sort_direction"`
	HasNext       bool   `json:"has_next"`
}

type Commission struct {
	Type          string `json:"type"`
	Rate          string `json:"rate"`
	TradingVolume string `json:"trading_volume"`
}

func (p Commission) RateNum() (rate decimal.Decimal, err error) {
	rate, err = core.StrToNum(p.Rate)
	if err != nil {
		err = fmt.Errorf("Invalid commission rate: %s - err: %w", p.Rate, err)
	}
	return
}

type OrderFill struct {
	Id              string    `json:"id"`
	OrderId         string    `json:"order_id"`
	Side            string    `json:"side"`
	ProductId       string    `json:"product_id"`
	ClientProductId string    `json:"client_product_id"`
	FilledQuantity  string    `json:"filled_quantity"`
	FilledValue     string    `json:"filled_value"`
	Price           string    `json:"price"`
	Time            time.Time `json:"time"`
	Commission      string    `json:"commission"`
	Venue           string    `json:"venue"`
	VenueFees       string    `json:"venue_fees"`
	CesCommission   string    `json:"ces_commission"`
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
	IsRaiseExact bool `json:"is_raise_exact,omitempty"`

	// Used for describe order, create order preview, and list portfolio orders
	Id                    string `json:"id,omitempty"`
	UserId                string `json:"user_id,omitempty"`
	Created               string `json:"created_at,omitempty"`
	FilledQuantity        string `json:"filled_quantity,omitempty"`
	FilledValue           string `json:"filled_value,omitempty"`
	AverageFilledPrice    string `json:"average_filled_price,omitempty"`
	Commission            string `json:"commission,omitempty"`
	ExchangeFee           string `json:"exchange_fee,omitempty"`
	Total                 string `json:"order_total,omitempty"`
	BestBid               string `json:"best_bid,omitempty"`
	BestAsk               string `json:"best_ask,omitempty"`
	Slippage              string `json:"slippage,omitempty"`
	Status                string `json:"status,omitempty"`
	HistoricalPov         string `json:"historical_pov,omitempty"`
	StopPrice             string `json:"stop_price,omitempty"`
	NetAverageFilledPrice string `json:"net_average_filled_price,omitempty"`
	UserContext           string `json:"user_context,omitempty"`
	ClientProductId       string `json:"client_product_id,omitempty"`
	PostOnly              bool   `json:"post_only,omitempty"`
	// Deprecated: Use EditHistory instead
	OrderEditHistory []*OrderEditHistory `json:"order_edit_history,omitempty"`
	DisplaySize      string              `json:"display_size,omitempty"`
	EditHistory      []*EditHistory      `json:"edit_history,omitempty"`
	PegOffsetType    string              `json:"peg_offset_type,omitempty"`
	Offset           string              `json:"offset,omitempty"`
	WigLevel         string              `json:"wig_level,omitempty"`
}

type Asset struct {
	Name             string     `json:"name"`
	Symbol           string     `json:"symbol"`
	DecimalPrecision string     `json:"decimal_precision"`
	TradingSupported bool       `json:"trading_supported"`
	ExplorerUrl      string     `json:"explorer_url"`
	Networks         []*Network `json:"networks"`
}

type Network struct {
	Network                *NetworkDetails `json:"network"`
	Name                   string          `json:"name"`
	MaxDecimals            string          `json:"max_decimals"`
	Default                bool            `json:"default"`
	TradingSupported       bool            `json:"trading_supported"`
	VaultSupported         bool            `json:"vault_supported"`
	PrimeCustodySupported  bool            `json:"prime_custody_supported"`
	DestinationTagRequired bool            `json:"destination_tag_required"`
	NetworkLink            string          `json:"network_link"`
	NetworkScopedSymbol    string          `json:"network_scoped_symbol"`
}

type NetworkDetails struct {
	Id   string `json:"id"`
	Type string `json:"type"`
}

type CryptoDepositInstructions struct {
	Id                string `json:"id"`
	Name              string `json:"name"`
	Type              string `json:"type"`
	Address           string `json:"address"`
	AccountIdentifier string `json:"account_identifier"`
}

type FiatDepositInstructions struct {
	Id            string `json:"id"`
	Name          string `json:"name"`
	Type          string `json:"type"`
	AccountNumber string `json:"account_number"`
	RoutingNumber string `json:"routing_number"`
	ReferenceCode string `json:"reference_code"`
}

type PostTradeCreditAmountDue struct {
	Currency string    `json:"currency"`
	Amount   string    `json:"amount"`
	DueDate  time.Time `json:"due_date"`
}

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

type EditHistory struct {
	Price            string `json:"price"`
	BaseQuantity     string `json:"base_quantity"`
	QuoteValue       string `json:"quote_value"`
	DisplayBaseSize  string `json:"display_base_size"`
	DisplayQuoteSize string `json:"display_quote_size"`
	StopPrice        string `json:"stop_price"`
	ExpiryTime       string `json:"expiry_time"`
	AcceptTime       string `json:"accept_time"`
	ClientOrderId    string `json:"client_order_id"`
}

type OrderEditHistory struct {
	Price          string `json:"price"`
	Size           string `json:"size"`
	DisplaySize    string `json:"display_size"`
	StopPrice      string `json:"stop_price"`
	StopLimitPrice string `json:"stop_limit_price"`
	EndTime        string `json:"end_time"`
	AcceptTime     string `json:"accept_time"`
	ClientOrderId  string `json:"client_order_id"`
}
