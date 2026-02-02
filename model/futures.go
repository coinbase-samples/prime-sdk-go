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

// FcmMarginCallType represents the type of margin call
type FcmMarginCallType string

const (
	FcmMarginCallTypeUnspecified FcmMarginCallType = "FCM_MARGIN_CALL_TYPE_UNSPECIFIED"
	FcmMarginCallTypeUrgent      FcmMarginCallType = "FCM_MARGIN_CALL_TYPE_URGENT"
	FcmMarginCallTypeRegular     FcmMarginCallType = "FCM_MARGIN_CALL_TYPE_REGULAR"
)

// FcmMarginCallState represents the state of a margin call
type FcmMarginCallState string

const (
	FcmMarginCallStateUnspecified FcmMarginCallState = "FCM_MARGIN_CALL_STATE_UNSPECIFIED"
	FcmMarginCallStateClosed      FcmMarginCallState = "FCM_MARGIN_CALL_STATE_CLOSED"
	FcmMarginCallStateRolledOver  FcmMarginCallState = "FCM_MARGIN_CALL_STATE_ROLLED_OVER"
	FcmMarginCallStateDefault     FcmMarginCallState = "FCM_MARGIN_CALL_STATE_DEFAULT"
	FcmMarginCallStateOfficial    FcmMarginCallState = "FCM_MARGIN_CALL_STATE_OFFICIAL"
)

// FcmMarginCall represents an FCM margin call
type FcmMarginCall struct {
	Type            FcmMarginCallType  `json:"type"`
	State           FcmMarginCallState `json:"state"`
	InitialAmount   string             `json:"initial_amount"`
	RemainingAmount string             `json:"remaining_amount"`
	BusinessDate    string             `json:"business_date"`
	CureDeadline    string             `json:"cure_deadline"`
}

// FcmRiskLimits represents FCM risk limits for an entity
type FcmRiskLimits struct {
	CfmRiskLimit                  string `json:"cfm_risk_limit"`
	CfmRiskLimitUtilization       string `json:"cfm_risk_limit_utilization"`
	CfmTotalMargin                string `json:"cfm_total_margin"`
	CfmDeltaOte                   string `json:"cfm_delta_ote"`
	CfmUnsettledRealizedPnl       string `json:"cfm_unsettled_realized_pnl"`
	CfmUnsettledAccruedFundingPnl string `json:"cfm_unsettled_accrued_funding_pnl"`
}

// FcmSettings represents FCM settings for an entity
type FcmSettings struct {
	TargetDerivativesExcess string `json:"target_derivatives_excess"`
}

// FcmBalance represents FCM balance information for a portfolio
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

// FcmPosition represents a futures position
type FcmPosition struct {
	ProductId         string `json:"product_id"`
	Side              string `json:"side"`
	NumberOfContracts string `json:"number_of_contracts"`
	DailyRealizedPnl  string `json:"daily_realized_pnl"`
	UnrealizedPnl     string `json:"unrealized_pnl"`
	CurrentPrice      string `json:"current_price"`
	AvgEntryPrice     string `json:"avg_entry_price"`
	ExpirationTime    string `json:"expiration_time"`
}

// FcmSweep represents a futures sweep
type FcmSweep struct {
	Id              string           `json:"id"`
	RequestedAmount *RequestedAmount `json:"requested_amount"`
	ShouldSweepAll  bool             `json:"should_sweep_all"`
	Status          string           `json:"status"`
	ScheduledTime   string           `json:"scheduled_time"`
}

// RequestedAmount represents a requested amount with currency
type RequestedAmount struct {
	Currency string `json:"currency"`
	Amount   string `json:"amount"`
}
