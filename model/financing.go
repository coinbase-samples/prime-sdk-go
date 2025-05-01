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

type Locate struct {
	// The currency symbol
	Symbol string `json:"symbol"`
	// The available quantity located
	Quantity string `json:"quantity"`
	// The interest rate for located symbol
	Rate string `json:"rate"`
}

type MarginCallRecord struct {
	// The unique ID of the margin call
	MarginCallId string `json:"margin_call_id"`

	// The initial margin call amount in notional value
	InitialNotionalAmount string `json:"initial_notional_amount"`

	// The outstanding margin call amount in notional value
	OutstandingNotionalAmount string `json:"outstanding_notional_amount"`

	// The time the margin call is created in RFC3339 format
	CreatedAt string `json:"created_at"`

	// The time the margin call is due in RFC3339 format
	DueAt string `json:"due_at"`
}

type LoanInfo struct {
	// The unique ID of the portfolio
	PortfolioId string `json:"portfolio_id"`

	// The currency symbol
	Symbol string `json:"symbol"`

	// Balance amount
	Amount string `json:"amount"`

	// Notional balance amount
	NotionalAmount string `json:"notional_amount"`

	// Settlement due date
	DueDate string `json:"due_date"`
}

type MarginAddOnType string

const (
	MarginAddOnTypeUnspecified     MarginAddOnType = "MARGIN_ADD_ON_TYPE_UNSPECIFIED"
	MarginAddOnSingleCoinStress    MarginAddOnType = "SINGLE_COIN_STRESS"
	MarginAddOnConcentrationStress MarginAddOnType = "CONCENTRATION_STRESS"
	MarginAddOnMacroStress         MarginAddOnType = "MACRO_STRESS"
	MarginAddOnShortBiasedStress   MarginAddOnType = "SHORT_BIASED_STRESS"
)

type LoanType string

const (
	LoanTypeTypeUnspecified     LoanType = "LOAN_TYPE_UNSET"
	LoanTypeBilateralLending    LoanType = "BILATERAL_LENDING"
	LoanTypeTradeFinance        LoanType = "TRADE_FINANCE"
	LoanTypePortfolioMargin     LoanType = "PORTFOLIO_MARGIN"
	LoanTypeShortCollateralLoan LoanType = "SHORT_COLLATERAL_LOAN"
	LoanTypeShortCollateral     LoanType = "SHORT_COLLATERAL"
)

type PMAssetInfo struct {
	// The currency symbol
	Symbol string `json:"symbol,omitempty"`

	// Nominal amount of the currency
	Amount string `json:"amount,omitempty"`

	// Spot price for the currency
	Price string `json:"price,omitempty"`

	// Notional amount of the currency
	NotionalAmount string `json:"notional_amount,omitempty"`

	// Asset tier of the currency
	AssetTier string `json:"asset_tier,omitempty"`

	// Whether the currency is margin eligible
	MarginEligible bool `json:"margin_eligible,omitempty"`

	// Base margin requirement of the currency
	BaseMarginRequirement string `json:"base_margin_requirement,omitempty"`

	// Notional amount of the currency's base margin requirement
	BaseMarginRequirementNotional string `json:"base_margin_requirement_notional,omitempty"`

	// The 30d adv of the currency
	Adv30d string `json:"adv_30d,omitempty"`

	// Historic 5d volatility of the currency
	Hist5dVol string `json:"hist_5d_vol,omitempty"`

	// Historic 30d volatility of the currency
	Hist30dVol string `json:"hist_30d_vol,omitempty"`

	// Historic 90d volatility of the currency
	Hist90dVol string `json:"hist_90d_vol,omitempty"`

	// Volatility margin addon of the currency position
	VolatilityAddon string `json:"volatility_addon,omitempty"`

	// Liquidity margin addon of the currency position
	LiquidityAddon string `json:"liquidity_addon,omitempty"`

	// Total position margin of the currency
	TotalPositionMargin string `json:"total_position_margin,omitempty"`

	// Nominal short position of the currency
	ShortNominal string `json:"short_nominal,omitempty"`

	// Nominal long position of the currency
	LongNominal string `json:"long_nominal,omitempty"`
}

type ShortCollateral struct {
	// Existing short collateral balance
	OldBalance string `json:"old_balance,omitempty"`

	// New short collateral balance required
	NewBalance string `json:"new_balance,omitempty"`

	// Loan interest rate
	LoanInterestRate string `json:"loan_interest_rate,omitempty"`

	// Collateral interest rate
	CollateralInterestRate string `json:"collateral_interest_rate,omitempty"`
}

type AssetBalance struct {
	// The unique ID of the portfolio
	PortfolioId string `json:"portfolio_id,omitempty"`

	// The currency symbol
	Symbol string `json:"symbol,omitempty"`

	// Balance amount
	Amount string `json:"amount,omitempty"`

	// Notional balance amount
	NotionalAmount string `json:"notional_amount,omitempty"`

	// Conversion rate
	ConversionRate string `json:"conversion_rate,omitempty"`
}

type MarketRate struct {
	// The currency symbol
	Symbol string `json:"symbol,omitempty"`

	// The current market rate of currency
	Rate string `json:"rate,omitempty"`
}

type MarginSummary struct {
	// The unique ID of the entity
	EntityId string `json:"entity_id,omitempty"`

	// The margin equity at the entity level. Margin Equity = LMV + SMV + Trading Cash Balance + Short Collateral - Pending Withdrawals
	MarginEquity string `json:"margin_equity,omitempty"`

	// USD notional value of required equity in entity portfolios
	MarginRequirement string `json:"margin_requirement,omitempty"`

	// margin_equity - margin_requirement
	ExcessDeficit string `json:"excess_deficit,omitempty"`

	// The raw amount of portfolio margin credit used
	PmCreditConsumed string `json:"pm_credit_consumed,omitempty"`

	// The maximum trade finance credit limit. This field is deprecated and will be removed in the future.
	TfCreditLimit string `json:"tf_credit_limit,omitempty"`

	// The amount of trade finance credit used (USD). This field is deprecated and will be removed in the future.
	TfCreditConsumed string `json:"tf_credit_consumed,omitempty"`

	// TF Asset Adjusted Value (USD). This field is deprecated and will be removed in the future.
	TfAdjustedAssetValue string `json:"tf_adjusted_asset_value,omitempty"`

	// TF Adjusted Liability Value (USD). This field is deprecated and will be removed in the future.
	TfAdjustedLiabilityValue string `json:"tf_adjusted_liability_value,omitempty"`

	// The amount of adjusted credit used. This field is deprecated and will be removed in the future.
	TfAdjustedCreditConsumed string `json:"tf_adjusted_credit_consumed,omitempty"`

	// The amount of adjusted equity. This field is deprecated and will be removed in the future.
	TfAdjustedEquity string `json:"tf_adjusted_equity,omitempty"`

	// Whether or not an entity is frozen due to balance outstanding or other reason
	Frozen bool `json:"frozen,omitempty"`

	// The reason why an entity is frozen
	FrozenReason string `json:"frozen_reason,omitempty"`

	// Whether TF is enabled for the entity. This field is deprecated and will be removed in the future.
	TfEnabled bool `json:"tf_enabled,omitempty"`

	// Whether PM is enabled for the entity
	PmEnabled bool `json:"pm_enabled,omitempty"`

	// Market rates for the list of assets
	MarketRates []MarketRate `json:"market_rates,omitempty"`

	// Asset Balances across portfolios
	AssetBalances []AssetBalance `json:"asset_balances,omitempty"`

	// Trade finance debit loan amounts. This field is deprecated and will be removed in the future.
	TfLoans []LoanInfo `json:"tf_loans,omitempty"`

	// Portfolio Margin debit loan amounts
	PmLoans []LoanInfo `json:"pm_loans,omitempty"`

	// Short collateral amounts
	ShortCollateral []LoanInfo `json:"short_collateral,omitempty"`

	// Gross market value (GMV) = LMV + Abs (SMV)
	GrossMarketValue string `json:"gross_market_value,omitempty"`

	// Net Market Value (NMV) = LMV + SMV
	NetMarketValue string `json:"net_market_value,omitempty"`

	// Long Market Value (LMV) = Sum of positive notional for all assets
	LongMarketValue string `json:"long_market_value,omitempty"`

	// Non_Marginable LMV: Sum of positive notional for each non-margin eligible coin
	NonMarginableLongMarketValue string `json:"non_marginable_long_market_value,omitempty"`

	// Short Market Value (SMV) = Sum of negative notional for each margin eligible coin
	ShortMarketValue string `json:"short_market_value,omitempty"`

	// Gross Leverage = GMV / Margin Requirement
	GrossLeverage string `json:"gross_leverage,omitempty"`

	// Net Exposure = (LMV + SMV) / GMV
	NetExposure string `json:"net_exposure,omitempty"`

	// Portfolio stress triggered
	PortfolioStressTriggered MarginAddOnType `json:"portfolio_stress_triggered,omitempty"`

	// PM asset info netted across the entity
	PmAssetInfo []PMAssetInfo `json:"pm_asset_info,omitempty"`

	// PM limit that monitors gross notional borrowings (crypto + fiat)
	PmCreditLimit string `json:"pm_credit_limit,omitempty"`

	// PM limit that monitors excess deficit
	PmMarginLimit string `json:"pm_margin_limit,omitempty"`

	// The amount of the margin limit that is consumed by the excess deficit
	PmMarginConsumed string `json:"pm_margin_consumed,omitempty"`
}

type MarginInformation struct {
	MarginCallRecords []MarginCallRecord `json:"margin_call_records,omitempty"`
	MarginSummary     MarginSummary      `json:"margin_summary,omitempty"`
}
