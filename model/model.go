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

	BalanceTypeTrading = "TRADING_BALANCES"
	BalanceTypeVault   = "VAULT_BALANCES"
	BalanceTypeTotal   = "TOTAL_BALANCES"

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
}

type Portfolio struct {
	Id             string `json:"id"`
	Name           string `json:"name"`
	EntityId       string `json:"entity_id"`
	OrganizationId string `json:"organization_id"`
}

type Balance struct {
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

func (b Balance) AmountNum() (amount decimal.Decimal, err error) {
	amount, err = core.StrToNum(b.Amount)
	if err != nil {
		err = fmt.Errorf("Invalid asset amount: %s - symbol: %s - msg: %v", b.Amount, b.Symbol, err)
	}
	return
}

func (b Balance) HoldsNum() (holds decimal.Decimal, err error) {
	holds, err = core.StrToNum(b.Holds)
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

type RfqProductDetails struct {
	Tradable        bool   `json:"tradable"`
	MinNotionalSize string `json:"min_notional_size"`
	MaxNotionalSize string `json:"max_notional_size"`
}

type Product struct {
	Id                string             `json:"id"`
	BaseIncrement     string             `json:"base_increment"`
	QuoteIncrement    string             `json:"quote_increment"`
	BaseMinSize       string             `json:"base_min_size"`
	BaseMaxSize       string             `json:"base_max_size"`
	QuoteMinSize      string             `json:"quote_min_size"`
	QuoteMaxSize      string             `json:"quote_max_size"`
	Permissions       []string           `json:"permissions"`
	PriceIncrement    string             `json:"price_increment"`
	RfqProductDetails *RfqProductDetails `json:"rfq_product_details"`
}

func (p Product) BaseMinSizeNum() (amount decimal.Decimal, err error) {
	amount, err = core.StrToNum(p.BaseMinSize)
	if err != nil {
		err = fmt.Errorf("invalid base min: %s - id: %s - err: %w", p.BaseMinSize, p.Id, err)
	}
	return
}

func (p Product) BaseMaxSizeNum() (amount decimal.Decimal, err error) {
	amount, err = core.StrToNum(p.BaseMaxSize)
	if err != nil {
		err = fmt.Errorf("invalid base max: %s - id: %s - err: %v", p.BaseMaxSize, p.Id, err)
	}
	return
}

func (p Product) BaseIncrementNum() (amount decimal.Decimal, err error) {
	amount, err = core.StrToNum(p.BaseIncrement)
	if err != nil {
		err = fmt.Errorf("invalid base increment: %s - id: %s - msg: %w", p.BaseIncrement, p.Id, err)
	}
	return
}

func (p Product) QuoteMinSizeNum() (amount decimal.Decimal, err error) {
	amount, err = core.StrToNum(p.QuoteMinSize)
	if err != nil {
		err = fmt.Errorf("invalid quote min: %s - id: %s - err: %w", p.QuoteMinSize, p.Id, err)
	}
	return
}

func (p Product) QuoteMaxSizeNum() (amount decimal.Decimal, err error) {
	amount, err = core.StrToNum(p.QuoteMaxSize)
	if err != nil {
		err = fmt.Errorf("invalid quote max: %s - id: %s - err: %v", p.QuoteMaxSize, p.Id, err)
	}
	return
}

func (p Product) QuoteIncrementNum() (amount decimal.Decimal, err error) {
	amount, err = core.StrToNum(p.QuoteIncrement)
	if err != nil {
		err = fmt.Errorf("invalid quote increment: %s - id: %s - msg: %w", p.QuoteIncrement, p.Id, err)
	}
	return
}

type OrderFill struct {
	Id             string    `json:"id"`
	OrderId        string    `json:"order_id"`
	Side           string    `json:"side"`
	ProductId      string    `json:"product_id"`
	FilledQuantity string    `json:"filled_quantity"`
	FilledValue    string    `json:"filled_value"`
	Price          string    `json:"price"`
	Time           time.Time `json:"time"`
	Commission     string    `json:"commission"`
	Venue          string    `json:"venue"`
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

type EstimatedNetworkFees struct {
	LowerBound string `json:"lower_bound,omitempty"`
	UpperBound string `json:"upper_bound,omitempty"`
}

type MatchMetadata struct {
	ReferenceId    string `json:"reference_id,omitempty"`
	SettlementDate string `json:"settlement_date,omitempty"`
}

type TransactionMetadata struct {
	MatchMetadata *MatchMetadata `json:"match_metadata,omitempty"`
}

type AssetChange struct {
	Symbol string `json:"symbol,omitempty"`
}

type OnchainDetail struct {
	SignedTransaction     string          `json:"signed_transaction"`
	RiskAssessment        *RiskAssessment `json:"risk_assessment"`
	ChainId               string          `json:"chain_id"`
	Nonce                 string          `json:"nonce"`
	ReplacedTransactionId string          `json:"replaced_transaction_id"`
	DestinationAddress    string          `json:"destination_address"`
	SkipBroadcast         bool            `json:"skip_broadcast"`
	FailureReason         string          `json:"failure_reason"`
	SigningStatus         string          `json:"signing_status"`
}

type Transaction struct {
	Id                    string                `json:"id"`
	WalletId              string                `json:"wallet_id"`
	PortfolioId           string                `json:"portfolio_id"`
	Type                  string                `json:"type"`
	Status                string                `json:"status"`
	Symbol                string                `json:"symbol"`
	Created               time.Time             `json:"created_at"`
	Completed             time.Time             `json:"completed_at"`
	Amount                string                `json:"amount"`
	TransferFrom          *Transfer             `json:"transfer_from,omitempty"`
	TransferTo            *Transfer             `json:"transfer_to,omitempty"`
	NetworkFees           string                `json:"network_fees"`
	Fees                  string                `json:"fees"`
	FeeSymbol             string                `json:"fee_symbol"`
	BlockchainIds         []string              `json:"blockchain_ids"`
	TransactionId         string                `json:"transaction_id"`
	DestinationSymbol     string                `json:"destination_symbol"`
	EstimatedNetworkFees  *EstimatedNetworkFees `json:"estimated_network_fees,omitempty"`
	Network               string                `json:"network"`
	EstimatedAssetChanges []AssetChange         `json:"estimated_asset_changes"`
	Metadata              *TransactionMetadata  `json:"metadata,omitempty"`
	IdempotencyKey        string                `json:"idempotency_key"`
	OnchainDetails        *OnchainDetail        `json:"onchain_details,omitempty"`
}

type RiskAssessment struct {
	ComplianceRiskDetected bool `json:"compliance_risk_detected"`
	SecurityRiskDetected   bool `json:"security_risk_detected"`
}

type Transfer struct {
	Type              string `json:"type"`
	Value             string `json:"value"`
	Address           string `json:"address"`
	AccountIdentifier string `json:"account_identifier"`
}

func (tr Transfer) ValueNum() (amount decimal.Decimal, err error) {
	amount, err = core.StrToNum(tr.Value)
	if err != nil {
		err = fmt.Errorf("invalid transfer value: %s - type: %s - msg: %v", tr.Value, tr.Type, err)
	}
	return
}

type Activity struct {
	Id                  string                `json:"id"`
	ReferenceId         string                `json:"reference_id"`
	Category            string                `json:"category"`
	PrimaryType         string                `json:"type"`
	SecondaryType       string                `json:"secondary_type"`
	Status              string                `json:"status"`
	CreatedBy           string                `json:"created_by"`
	Title               string                `json:"title"`
	Description         string                `json:"description"`
	UserActions         []*UserAction         `json:"user_actions,omitempty"`
	AccountMetadata     *AccountMetadata      `json:"account_metadata,omitempty"`
	OrdersMetadata      *OrdersMetadata       `json:"orders_metadata,omitempty"`
	TransactionMetadata *TransactionsMetadata `json:"transaction_metadata,omitempty"`
	Symbols             []string              `json:"symbols,omitempty"`
	Created             string                `json:"created_at"`
	Updated             string                `json:"updated_at"`
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
	TransactionsMetadata *TransactionsMetadata `json:"transactions_metadata,omitempty"`
}

type AddressBookEntry struct {
	Id                    string                   `json:"id"`
	Symbol                string                   `json:"currency_symbol"`
	Name                  string                   `json:"name"`
	Address               string                   `json:"address"`
	AccountIdentifier     string                   `json:"account_identifier"`
	AccountIdentifierName string                   `json:"account_identifier_name"`
	State                 string                   `json:"state"`
	ExplorerLink          string                   `json:"explorer_link"`
	LastUsed              time.Time                `json:"last_used_at"`
	Added                 time.Time                `json:"added_at"`
	AddedBy               *AddressBookEntryAddedBy `json:"added_by"`
}

type AddressBookEntryAddedBy struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	AvatarUrl string `json:"avatar_url"`
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

type Invoice struct {
	Id            string         `json:"id"`
	BillingYear   int32          `json:"billing_year"`
	BillingMonth  int32          `json:"billing_month"`
	DueDate       string         `json:"due_date"`
	InvoiceNumber string         `json:"invoice_number"`
	State         string         `json:"state"`
	UsdAmountPaid float64        `json:"usd_amount_paid"`
	UsdAmountOwed float64        `json:"usd_amount_owed"`
	Items         []*InvoiceItem `json:"invoice_items"`
}

type InvoiceItem struct {
	Description    string  `json:"description"`
	CurrencySymbol string  `json:"currency_symbol"`
	InvoiceType    string  `json:"invoice_type"`
	Rate           float64 `json:"rate"`
	Quantity       float64 `json:"quantity"`
	Price          float64 `json:"price"`
	AverageAuc     float64 `json:"average_auc"`
	Total          float64 `json:"total"`
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

type FcmBalance struct {
	PortfolioId        string `json:"portfolio_id"`
	CfmUsdBalance      string `json:"cfm_usd_balance"`
	UnrealizedPnl      string `json:"unrealized_pnl"`
	DailyRealizedPnl   string `json:"daily_realized_pnl"`
	ExcessLiquidity    string `json:"excess_liquidity"`
	FuturesBuyingPower string `json:"futures_buying_power"`
	InitialMargin      string `json:"initial_margin"`
	MaintenanceMargin  string `json:"maintenance_margin"`
	ClearingAccountId  string `json:"clearing_account_id"`
}

type Position struct {
	ProductId         string    `json:"product_id"`
	Side              string    `json:"side"`
	NumberOfContracts string    `json:"number_of_contracts"`
	DailyRealizedPnl  string    `json:"daily_realized_pnl"`
	UnrealizedPnl     string    `json:"unrealized_pnl"`
	CurrentPrice      string    `json:"current_price"`
	AvgEntryPrice     string    `json:"avg_entry_price"`
	ExpirationTime    time.Time `json:"expiration_time"`
}

type Sweep struct {
	Id              string           `json:"id"`
	RequestedAmount *RequestedAmount `json:"requested_amount"`
	ShouldSweepAll  bool             `json:"should_sweep_all"`
	Status          string           `json:"status"`
	ScheduledTime   string           `json:"scheduled_time"`
}

type RequestedAmount struct {
	Currency string `json:"currency"`
	Amount   string `json:"amount"`
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

type CandleGranularity string

const (
	CandleGranularityOneMinute      CandleGranularity = "ONE_MINUTE"
	CandleGranularityFiveMinutes    CandleGranularity = "FIVE_MINUTES"
	CandleGranularityFifteenMinutes CandleGranularity = "FIFTEEN_MINUTES"
	CandleGranularityThirtyMinutes  CandleGranularity = "THIRTY_MINUTES"
	CandleGranularityOneHour        CandleGranularity = "ONE_HOUR"
	CandleGranularityTwoHours       CandleGranularity = "TWO_HOURS"
	CandleGranularityFourHours      CandleGranularity = "FOUR_HOURS"
	CandleGranularitySixHours       CandleGranularity = "SIX_HOURS"
	CandleGranularityOneDay         CandleGranularity = "ONE_DAY"
)

type Candle struct {
	Timestamp string `json:"timestamp"`
	Open      string `json:"open"`
	High      string `json:"high"`
	Low       string `json:"low"`
	Close     string `json:"close"`
	Volume    string `json:"volume"`
}

// FCM Futures types
type FcmMarginCallType string

const (
	FcmMarginCallTypeUnspecified FcmMarginCallType = "FCM_MARGIN_CALL_TYPE_UNSPECIFIED"
)

type FcmMarginCallState string

const (
	FcmMarginCallStateUnspecified FcmMarginCallState = "FCM_MARGIN_CALL_STATE_UNSPECIFIED"
)

type FcmMarginCall struct {
	Type            FcmMarginCallType  `json:"type"`
	State           FcmMarginCallState `json:"state"`
	InitialAmount   string             `json:"initial_amount"`
	RemainingAmount string             `json:"remaining_amount"`
	BusinessDate    string             `json:"business_date"`
	CureDeadline    string             `json:"cure_deadline"`
}

type FcmRiskLimits struct {
	CfmRiskLimit                  string `json:"cfm_risk_limit"`
	CfmRiskLimitUtilization       string `json:"cfm_risk_limit_utilization"`
	CfmTotalMargin                string `json:"cfm_total_margin"`
	CfmDeltaOte                   string `json:"cfm_delta_ote"`
	CfmUnsettledRealizedPnl       string `json:"cfm_unsettled_realized_pnl"`
	CfmUnsettledAccruedFundingPnl string `json:"cfm_unsettled_accrued_funding_pnl"`
}

type UnstakeEstimateType string

const (
	UnstakeEstimateTypeUnspecified UnstakeEstimateType = "UNSPECIFIED"
)

type UnstakeType string

const (
	UnstakeTypeUnspecified UnstakeType = "UNSTAKE_TYPE_UNSPECIFIED"
)

type UnstakeStatus struct {
	Amount              string              `json:"amount"`
	EstimateType        UnstakeEstimateType `json:"estimate_type"`
	EstimateDescription string              `json:"estimate_description"`
	UnstakeType         UnstakeType         `json:"unstake_type"`
	FinishingAt         string              `json:"finishing_at"`
	RemainingHours      int                 `json:"remaining_hours"`
	RequestedAt         string              `json:"requested_at"`
}

type UnstakeValidator struct {
	ValidatorAddress string           `json:"validator_address"`
	Statuses         []*UnstakeStatus `json:"statuses"`
}

type ValidatorStatus string

const (
	ValidatorStatusUnspecified ValidatorStatus = "VALIDATOR_STATUS_UNSPECIFIED"
)

type TransactionValidator struct {
	TransactionId    string          `json:"transaction_id"`
	ValidatorAddress string          `json:"validator_address"`
	ValidatorStatus  ValidatorStatus `json:"validator_status"`
}
