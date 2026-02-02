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

type LocateAvailability struct {
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

type PmAssetInfo struct {
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
	MarketRates []*MarketRate `json:"market_rates,omitempty"`

	// Asset Balances across portfolios
	AssetBalances []*AssetBalance `json:"asset_balances,omitempty"`

	// Trade finance debit loan amounts. This field is deprecated and will be removed in the future.
	TfLoans []*LoanInfo `json:"tf_loans,omitempty"`

	// Portfolio Margin debit loan amounts
	PmLoans []*LoanInfo `json:"pm_loans,omitempty"`

	// Short collateral amounts
	ShortCollateral []*LoanInfo `json:"short_collateral,omitempty"`

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
	PmAssetInfo []*PmAssetInfo `json:"pm_asset_info,omitempty"`

	// PM limit that monitors gross notional borrowings (crypto + fiat)
	PmCreditLimit string `json:"pm_credit_limit,omitempty"`

	// PM limit that monitors excess deficit
	PmMarginLimit string `json:"pm_margin_limit,omitempty"`

	// The amount of the margin limit that is consumed by the excess deficit
	PmMarginConsumed string `json:"pm_margin_consumed,omitempty"`
}

type MarginInfo struct {
	MarginCallRecords []*MarginCallRecord `json:"margin_call_records,omitempty"`
	MarginSummary     *MarginSummary      `json:"margin_summary,omitempty"`
}

type MarginSummaryHistorical struct {
	// The UTC date time used for conversion
	ConversionDatetime string `json:"conversion_datetime,omitempty"`

	// The date used for conversion
	ConversionDate string `json:"conversion_date,omitempty"`

	// The margin summary
	MarginSummary *MarginSummary `json:"margin_summary,omitempty"`
}

type BuyingPower struct {
	// The unique ID of the portfolio
	PortfolioId string `json:"portfolio_id,omitempty"`

	// The symbol for the base currency
	BaseCurrency string `json:"base_currency,omitempty"`

	// The symbol for the quote currency
	QuoteCurrency string `json:"quote_currency,omitempty"`

	// The buying power for the base currency
	BaseBuyingPower string `json:"base_buying_power,omitempty"`

	// The buying power for the quote currency
	QuoteBuyingPower string `json:"quote_buying_power,omitempty"`
}

type AmountDue struct {
	// The currency this loan is due in
	Currency string `json:"currency,omitempty"`

	// The amount due
	Amount string `json:"amount,omitempty"`

	// The date this settlement is due, expressed in UTC
	DueDate string `json:"due_date,omitempty"`
}

type PostTradeCreditInfo struct {
	// The unique ID of the portfolio
	PortfolioId string `json:"portfolio_id,omitempty"`

	// The currency symbol credit is denoted in
	Currency string `json:"currency,omitempty"`

	// The maximum credit limit
	Limit string `json:"limit,omitempty"`

	// The amount of credit used
	Utilized string `json:"utilized,omitempty"`

	// The amount of credit available
	Available string `json:"available,omitempty"`

	// Whether or not a portfolio is frozen due to balance outstanding or other reason
	Frozen bool `json:"frozen,omitempty"`

	// The reason why the portfolio is frozen
	FrozenReason string `json:"frozen_reason,omitempty"`

	// Amounts due
	AmountsDue []*AmountDue `json:"amounts_due,omitempty"`

	// Whether the portfolio has credit enabled
	Enabled bool `json:"enabled,omitempty"`

	// The amount of adjusted credit used
	AdjustedCreditUtilized string `json:"adjusted_credit_utilized,omitempty"`

	// The amount of adjusted portfolio equity
	AdjustedPortfolioEquity string `json:"adjusted_portfolio_equity,omitempty"`
}

type WithdrawalPower struct {
	// The currency symbol
	Symbol string `json:"symbol,omitempty"`

	// Withdrawal power
	Amount string `json:"amount,omitempty"`
}

type TieredPricingFee struct {
	// Asset symbol
	Symbol string `json:"symbol,omitempty"`

	// The fee in bps
	Fee string `json:"fee,omitempty"`
}

type Locate struct {
	// The locate ID
	LocateId string `json:"locate_id,omitempty"`

	// The unique ID of the entity
	EntityId string `json:"entity_id,omitempty"`

	// The unique ID of the portfolio
	PortfolioId string `json:"portfolio_id,omitempty"`

	// The currency symbol
	Symbol string `json:"symbol,omitempty"`

	// The requested locate amount
	RequestedAmount string `json:"requested_amount,omitempty"`

	// The interest rate of PM loan
	InterestRate string `json:"interest_rate,omitempty"`

	// The locate status
	Status string `json:"status,omitempty"`

	// The approved locate amount
	ApprovedAmount string `json:"approved_amount,omitempty"`

	// Deprecated: Use locate_date instead
	ConversionDate string `json:"conversion_date,omitempty"`

	// The date when the locate was submitted in RFC3339 format
	CreatedAt string `json:"created_at,omitempty"`

	// The locate date from the CreateNewLocatesRequest in RFC3339 format
	LocateDate string `json:"locate_date,omitempty"`
}

type Benchmark string

const (
	BenchmarkUnset     Benchmark = "BENCHMARK_UNSET"
	BenchmarkZero      Benchmark = "ZERO"
	BenchmarkSofr360   Benchmark = "SOFR_360"
	BenchmarkSofr365   Benchmark = "SOFR_365"
	BenchmarkCryptoRfr Benchmark = "CRYPTO_RFR"
)

type RateType string

const (
	RateTypeUnset  RateType = "RATE_TYPE_UNSET"
	RateTypeBps    RateType = "BPS"
	RateTypeApr360 RateType = "APR_360"
	RateTypeApr365 RateType = "APR_365"
	RateTypeApr    RateType = "APR"
)

type Accrual struct {
	// The accrual ID
	AccrualId string `json:"accrual_id,omitempty"`

	// The date of accrual in UTC
	Date string `json:"date,omitempty"`

	// The unique ID of the portfolio
	PortfolioId string `json:"portfolio_id,omitempty"`

	// The currency symbol
	Symbol string `json:"symbol,omitempty"`

	// The loan type
	LoanType *LoanType `json:"loan_type,omitempty"`

	// The daily or annualized interest rate for the loan, see rate_type
	InterestRate string `json:"interest_rate,omitempty"`

	// Daily accrual amount in the principal currency
	NominalAccrual string `json:"nominal_accrual,omitempty"`

	// Daily USD accrued interest
	NotionalAccrual string `json:"notional_accrual,omitempty"`

	// Accrual rate used to convert from principal to USD accrual
	ConversionRate string `json:"conversion_rate,omitempty"`

	// Outstanding principal of the loan
	LoanAmount string `json:"loan_amount,omitempty"`

	// Benchmark information
	Benchmark Benchmark `json:"benchmark,omitempty"`

	// Daily interest rate fetched from the benchmark source
	BenchmarkRate string `json:"benchmark_rate,omitempty"`

	// Daily spread offset from the benchmark rate
	Spread string `json:"spread,omitempty"`

	// The rate type
	RateType RateType `json:"rate_type,omitempty"`

	// Outstanding principal of the loan in USD
	LoanAmountNotional string `json:"loan_amount_notional,omitempty"`

	// Settled open borrow as of start-of-day in the principal currency
	NominalOpenBorrowSod string `json:"nominal_open_borrow_sod,omitempty"`

	// Settled open borrow as of start-of-day in USD
	NotionalOpenBorrowSod string `json:"notional_open_borrow_sod,omitempty"`
}

type ConversionDetail struct {
	// The currency symbol
	Symbol string `json:"symbol,omitempty"`

	// Trade finance balance after the conversion
	TfBalance string `json:"tf_balance,omitempty"`

	// Notional trade finance balance after the conversion
	NotionalTfBalance string `json:"notional_tf_balance,omitempty"`

	// Converted balance
	ConvertedBalance string `json:"converted_balance,omitempty"`

	// Notional converted balance
	NotionalConvertedBalance string `json:"notional_converted_balance,omitempty"`

	// Interest rate
	InterestRate string `json:"interest_rate,omitempty"`

	// Conversion rate
	ConversionRate string `json:"conversion_rate,omitempty"`
}

type Conversion struct {
	// Conversion details
	ConversionDetails []*ConversionDetail `json:"conversion_details,omitempty"`

	// Short collateral
	ShortCollateral *ShortCollateral `json:"short_collateral,omitempty"`

	// The UTC date time used for conversion
	ConversionDatetime string `json:"conversion_datetime,omitempty"`

	// Portfolio ID
	PortfolioId string `json:"portfolio_id,omitempty"`
}

// TFAsset represents an asset eligible for Trade Finance
type TFAsset struct {
	Symbol              string `json:"symbol"`
	AssetAdjustment     string `json:"asset_adjustment"`
	LiabilityAdjustment string `json:"liability_adjustment"`
}

// XMControlStatus represents the control status for Cross Margin trades and withdrawals
type XMControlStatus string

const (
	XMControlStatusUnspecified       XMControlStatus = "XM_CONTROL_STATUS_UNSPECIFIED"
	XMControlStatusTradesWithdrawals XMControlStatus = "TRADES_AND_WITHDRAWALS"
	XMControlStatusTradesOnly        XMControlStatus = "TRADES_ONLY"
	XMControlStatusSessionLocked     XMControlStatus = "SESSION_LOCKED"
)

// XMEntityCallStatus represents the entity call status for Cross Margin
type XMEntityCallStatus string

const (
	XMEntityCallStatusUnspecified  XMEntityCallStatus = "XM_ENTITY_CALL_STATUS_UNSPECIFIED"
	XMEntityCallStatusNoCall       XMEntityCallStatus = "ENTITY_NO_CALL"
	XMEntityCallStatusStandardCall XMEntityCallStatus = "ENTITY_OPEN_STANDARD_CALL"
	XMEntityCallStatusUrgentCall   XMEntityCallStatus = "ENTITY_OPEN_URGENT_CALL"
	XMEntityCallStatusAgedCall     XMEntityCallStatus = "ENTITY_AGED_CALL"
	XMEntityCallStatusDebitCall    XMEntityCallStatus = "ENTITY_OPEN_DEBIT_CALL"
)

// XMMarginLevel represents the margin level for Cross Margin
type XMMarginLevel string

const (
	XMMarginLevelUnspecified XMMarginLevel = "XM_MARGIN_LEVEL_UNSPECIFIED"
	XMMarginLevelHealthy     XMMarginLevel = "HEALTHY_THRESHOLD"
	XMMarginLevelDeficit     XMMarginLevel = "DEFICIT_THRESHOLD"
	XMMarginLevelWarning     XMMarginLevel = "WARNING_THRESHOLD"
	XMMarginLevelUrgent      XMMarginLevel = "URGENT_MARGIN_CALL_THRESHOLD"
	XMMarginLevelLiquidation XMMarginLevel = "LIQUIDATION_THRESHOLD"
)

// XMParty represents a Cross Margin trading venue
type XMParty string

const (
	XMPartyUnspecified XMParty = "XM_PARTY_UNSPECIFIED"
	XMPartyCBE         XMParty = "CBE"
	XMPartyFCM         XMParty = "FCM"
)

// XMCallType represents the type of Cross Margin call
type XMCallType string

const (
	XMCallTypeUnspecified XMCallType = "XM_CALL_TYPE_UNSPECIFIED"
	XMCallTypeStandard    XMCallType = "STANDARD"
	XMCallTypeUrgent      XMCallType = "URGENT"
)

// XMCallStatus represents the status of a Cross Margin call
type XMCallStatus string

const (
	XMCallStatusUnspecified XMCallStatus = "XM_CALL_STATUS_UNSPECIFIED"
	XMCallStatusOpen        XMCallStatus = "OPEN"
	XMCallStatusClosed      XMCallStatus = "CLOSED"
	XMCallStatusAged        XMCallStatus = "AGED"
)

// XMSummary represents the Cross Margin margin model summary
type XMSummary struct {
	MarginRequirement     string `json:"margin_requirement"`
	AccountEquity         string `json:"account_equity"`
	MarginExcessShortfall string `json:"margin_excess_shortfall"`
	ConsumedCredit        string `json:"consumed_credit"`
	XMCreditLimit         string `json:"xm_credit_limit"`
}

// XMMarginCall represents a Cross Margin margin call
type XMMarginCall struct {
	MarginCallId              string        `json:"margin_call_id"`
	Currency                  string        `json:"currency"`
	InitialNotionalAmount     string        `json:"initial_notional_amount"`
	OutstandingNotionalAmount string        `json:"outstanding_notional_amount"`
	MarginCallType            XMCallType    `json:"margin_call_type"`
	MarginCallStatus          XMCallStatus  `json:"margin_call_status"`
	CalledWithMarginLevel     XMMarginLevel `json:"called_with_margin_level"`
}

// XMLoan represents a Cross Margin loan
type XMLoan struct {
	LoanId                       string  `json:"loan_id"`
	LoanParty                    XMParty `json:"loan_party"`
	PrincipalCurrency            string  `json:"principal_currency"`
	PrincipalCurrencyMarketPrice string  `json:"principal_currency_market_price"`
	InitialPrincipalAmount       string  `json:"initial_principal_amount"`
	OutstandingPrincipalAmount   string  `json:"outstanding_principal_amount"`
}

// CrossMarginOverview represents the Cross Margin overview for an entity
type CrossMarginOverview struct {
	ControlStatus     XMControlStatus    `json:"control_status"`
	CallStatus        XMEntityCallStatus `json:"call_status"`
	MarginLevel       XMMarginLevel      `json:"margin_level"`
	MarginSummary     *XMSummary         `json:"margin_summary"`
	ActiveMarginCalls []*XMMarginCall    `json:"active_margin_calls"`
	ActiveLoans       []*XMLoan          `json:"active_loans"`
}
