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

import (
	"fmt"

	"github.com/coinbase-samples/core-go"
	"github.com/shopspring/decimal"
)

// ProductType represents the general type of product.
type ProductType string

const (
	ProductTypeSpot   ProductType = "SPOT"
	ProductTypeFuture ProductType = "FUTURE"
)

// ContractExpiryType represents the expiry type of a futures contract.
type ContractExpiryType string

const (
	ContractExpiryTypeUnspecified ContractExpiryType = "CONTRACT_EXPIRY_TYPE_UNSPECIFIED"
	ContractExpiryTypeExpiring    ContractExpiryType = "CONTRACT_EXPIRY_TYPE_EXPIRING"
	ContractExpiryTypePerpetual   ContractExpiryType = "CONTRACT_EXPIRY_TYPE_PERPETUAL"
)

// ExpiringContractStatus filters expiring futures by their expiry status.
type ExpiringContractStatus string

const (
	ExpiringContractStatusUnexpired ExpiringContractStatus = "EXPIRING_CONTRACT_STATUS_UNEXPIRED"
	ExpiringContractStatusExpired   ExpiringContractStatus = "EXPIRING_CONTRACT_STATUS_EXPIRED"
	ExpiringContractStatusAll       ExpiringContractStatus = "EXPIRING_CONTRACT_STATUS_ALL"
)

// RiskManagementType represents how risk is managed for a product.
type RiskManagementType string

const (
	RiskManagementTypeUnspecified    RiskManagementType = "RISK_MANAGEMENT_TYPE_UNSPECIFIED"
	RiskManagementTypeManagedByFcm   RiskManagementType = "RISK_MANAGEMENT_TYPE_MANAGED_BY_FCM"
	RiskManagementTypeManagedByVenue RiskManagementType = "RISK_MANAGEMENT_TYPE_MANAGED_BY_VENUE"
)

// FcmTradingSessionClosedReason represents the reason for FCM trading session closure.
type FcmTradingSessionClosedReason string

const (
	FcmTradingSessionClosedReasonUndefined           FcmTradingSessionClosedReason = "FCM_TRADING_SESSION_CLOSED_REASON_UNDEFINED"
	FcmTradingSessionClosedReasonRegularMarketClose  FcmTradingSessionClosedReason = "FCM_TRADING_SESSION_CLOSED_REASON_REGULAR_MARKET_CLOSE"
	FcmTradingSessionClosedReasonExchangeMaintenance FcmTradingSessionClosedReason = "FCM_TRADING_SESSION_CLOSED_REASON_EXCHANGE_MAINTENANCE"
	FcmTradingSessionClosedReasonVendorMaintenance   FcmTradingSessionClosedReason = "FCM_TRADING_SESSION_CLOSED_REASON_VENDOR_MAINTENANCE"
)

// FcmTradingSessionState represents the current state of an FCM trading session.
type FcmTradingSessionState string

const (
	FcmTradingSessionStateUndefined       FcmTradingSessionState = "FCM_TRADING_SESSION_STATE_UNDEFINED"
	FcmTradingSessionStatePreOpen         FcmTradingSessionState = "FCM_TRADING_SESSION_STATE_PRE_OPEN"
	FcmTradingSessionStatePreOpenNoCancel FcmTradingSessionState = "FCM_TRADING_SESSION_STATE_PRE_OPEN_NO_CANCEL"
	FcmTradingSessionStateOpen            FcmTradingSessionState = "FCM_TRADING_SESSION_STATE_OPEN"
	FcmTradingSessionStateClose           FcmTradingSessionState = "FCM_TRADING_SESSION_STATE_CLOSE"
	FcmTradingSessionStateHalted          FcmTradingSessionState = "FCM_TRADING_SESSION_STATE_HALTED"
)

// FcmScheduledMaintenance contains scheduled maintenance window information.
type FcmScheduledMaintenance struct {
	StartTime string `json:"start_time,omitempty"`
	EndTime   string `json:"end_time,omitempty"`
}

// FcmTradingSessionDetails contains trading session details for FCM products.
type FcmTradingSessionDetails struct {
	SessionOpen                  bool                          `json:"session_open"`
	OpenTime                     string                        `json:"open_time,omitempty"`
	CloseTime                    string                        `json:"close_time,omitempty"`
	SessionState                 FcmTradingSessionState        `json:"session_state,omitempty"`
	AfterHoursOrderEntryDisabled bool                          `json:"after_hours_order_entry_disabled"`
	ClosedReason                 FcmTradingSessionClosedReason `json:"closed_reason,omitempty"`
	Maintenance                  *FcmScheduledMaintenance      `json:"maintenance,omitempty"`
	SettlementTimestamp          string                        `json:"settlement_timestamp,omitempty"`
	SettlementPrice              string                        `json:"settlement_price,omitempty"`
}

// PerpetualProductDetails contains details specific to perpetual futures products.
type PerpetualProductDetails struct {
	OpenInterest   string `json:"open_interest,omitempty"`
	FundingRate    string `json:"funding_rate,omitempty"`
	FundingTime    string `json:"funding_time,omitempty"`
	MaxLeverage    string `json:"max_leverage,omitempty"`
	UnderlyingType string `json:"underlying_type,omitempty"`
}

// FutureProductDetails contains details specific to futures products.
type FutureProductDetails struct {
	ContractCode           string                   `json:"contract_code,omitempty"`
	ContractSize           string                   `json:"contract_size,omitempty"`
	ContractExpiry         string                   `json:"contract_expiry,omitempty"`
	ContractRootUnit       string                   `json:"contract_root_unit,omitempty"`
	ContractExpiryType     ContractExpiryType       `json:"contract_expiry_type,omitempty"`
	RiskManagedBy          RiskManagementType       `json:"risk_managed_by,omitempty"`
	Venue                  string                   `json:"venue,omitempty"`
	GroupDescription       string                   `json:"group_description,omitempty"`
	ContractExpiryTimezone string                   `json:"contract_expiry_timezone,omitempty"`
	GroupShortDescription  string                   `json:"group_short_description,omitempty"`
	PerpetualDetails       *PerpetualProductDetails `json:"perpetual_details,omitempty"`
}

type RfqProductDetails struct {
	Tradable     bool   `json:"tradable"`
	MinBaseSize  string `json:"min_base_size"`
	MaxBaseSize  string `json:"max_base_size"`
	MinQuoteSize string `json:"min_quote_size"`
	MaxQuoteSize string `json:"max_quote_size"`
	// Deprecated: Value will be an empty string. Use Min/Max Base/Quote Size instead.
	MinNotionalSize string `json:"min_notional_size"`
	// Deprecated: Value will be an empty string. Use Min/Max Base/Quote Size instead.
	MaxNotionalSize string `json:"max_notional_size"`
}

type Product struct {
	Id                       string                    `json:"id"`
	BaseIncrement            string                    `json:"base_increment"`
	QuoteIncrement           string                    `json:"quote_increment"`
	BaseMinSize              string                    `json:"base_min_size"`
	BaseMaxSize              string                    `json:"base_max_size"`
	QuoteMinSize             string                    `json:"quote_min_size"`
	QuoteMaxSize             string                    `json:"quote_max_size"`
	Permissions              []string                  `json:"permissions"`
	PriceIncrement           string                    `json:"price_increment"`
	RfqProductDetails        *RfqProductDetails        `json:"rfq_product_details"`
	ProductType              ProductType               `json:"product_type,omitempty"`
	FcmTradingSessionDetails *FcmTradingSessionDetails `json:"fcm_trading_session_details,omitempty"`
	FutureProductDetails     *FutureProductDetails     `json:"future_product_details,omitempty"`
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
