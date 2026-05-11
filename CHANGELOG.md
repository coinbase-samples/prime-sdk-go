# Changelog


## [0.7.0] - 2026-MAY-11

### Added

- New Beta Financing endpoints
  - `GetCrossMarginRiskParameters` — retrieves XM 2.0 tier risk parameters and offset credit matrices for an entity
  - `GetCrossMarginPrimeOverview` — returns full live Prime cross-margin information (served from `/v2`)
  - `SetFundingSettings` — sets FCM funding configuration for an entity (creates a PCS proposal)
  - `GetMarketData` — retrieves paginated volatility and ADV market data for an entity
- `client.VersionedBaseUrl` and `client.WithBaseUrl` helpers for per-call API version overrides (used internally by `GetCrossMarginPrimeOverview` for the `/v2` path, without affecting other calls)
- `NewFinancingServiceWithConfig` constructor for pagination control on `GetMarketData`
- Financing models: `XMLiquidationStatus`, `ActiveLiquidationSummary`; `ActiveLiquidation` field on `CrossMarginOverview`
- Financing models: `MarginAddOn`, `XMPosition`, `XMRiskNettingInfo`; `XMMarginLimit`, `SpotEquity`, `FuturesEquity`, `RiskNettingInfo` fields on `XMSummary`
- Beta financing models: `PrimeXMControlStatus`, `PrimeXMMarginLevel`, `PrimeXMHealthStatus`, `PrimeXMMarginRequirementType`, `PrimeXMMarginThresholdType`, `CrossMarginRiskParameters`, `TierPairRateEntry`, `CrossMarginPrimeMarginSummary`, `CrossMarginPrimeSpotEquityBreakdown`, `CrossMarginPrimeDerivativesEquityBreakdown`, `CrossMarginPrimeRiskNettingInfo`, `PrimeXMMarginRequirementBreakdown`, `PrimeXMOffsetCreditBreakdown`, `CrossMarginPrimeXMPosition`, `PrimeXMMarginCallThresholds`, `PrimeXMMarginThreshold`, `MarketData`
- Staking model: `ValidatorUnstakePreview`; `WalletId`, `WalletAddress`, `CurrentTimestamp`, `Validators` fields on `PreviewUnstakeResponse`
- User model: `BUSINESS_MANAGER` user role
- RFQ: `QuoteDurationMs` optional field on `CreateQuoteRequest` and `CreateQuoteResponse`


## [0.6.3] - 2026-APR-30

### Added
- Add RFQ information to products


## [0.6.2] - 2026-APR-21

### Added
- Add entity name to portfolio struct


## [0.6.1] - 2026-APR-20

### Added
- New attributes on List Assets

## [0.6.0] - 2026-MAR-30

### Added

- New `advancedtransfers` package with four endpoints
  - ListAdvancedTransfers
  - CreateAdvancedTransfer
  - CancelAdvancedTransfer
  - ListAdvancedTransferTransactions
- New Transaction endpoint: GetTransactionTravelRuleData
- New examples: listAdvancedTransfers, createAdvancedTransfer, cancelAdvancedTransfer, getTransactionTravelRuleData
- New models: `AdvancedTransfer`, `AdvancedTransferState`, `AdvancedTransferType`, `BlindMatchMetadata`, `FundMovement`, `TransferLocation`
- New product types: `ProductType`, `ContractExpiryType`, `ExpiringContractStatus`, `FutureProductDetails`, `PerpetualProductDetails`, `FcmTradingSessionDetails`
- New model types: `CommissionDetailTotal`, `UserRole`, `SecondaryPermission`, `FcmMarginHealthState`, `StakingRewardType`, `ValidatorAllocation`

### Updated

- ListProducts supports three new optional query params: `product_type`, `contract_expiry_type`, `expiring_contract_status`
- `Product` model has new fields: `product_type`, `fcm_trading_session_details`, `future_product_details`
- `Order` and `OrderFill` models have new fields: `product_type`, `commission_detail_total`
- `GetFcmRiskLimitsResponse` has new fields: `cfm_unsettled_accrued_funding_pnl`, `margin_utilization_percent`, `margin_health_state`
- `User` model has new fields: `roles`, `secondary_permissions`
- `CreateUnstakeInputs` has new field: `validator_allocations` (Alpha — ETH V2 validator-level unstaking)
- `StakingRewardType` has new enum value: `BUIDL_DIVIDEND`

## [0.5.2] - 2025-JUN-17

### Added

- Wallet Service has two new endpoints
  - listWalletAddresses
  - createWalletAddress

## [0.5.0] - 2025-JUN-11

### Fix

- Align financing models with rest of the SDK

## [0.4.3] - 2025-JUN-02

### Added

- Add Network info to Get and List Wallets
- Added TransactionMetadata to activities


## [0.4.1] - 2025-MAY-13

### Added

- Add disable dynamic nonce to EVM params for Onchain Txs
- Add settle currency to quote requests

## [0.4.0] - 2025-MAY-13

### Fix

- Request structs are now excluded from JSON serialization when marshaling a response

### Added

- Add pagination support for ListOpenOrders and slice support for product IDs

## [0.3.8] - 2025-MAY-07

### Added

- Adding support for new Financing endpoints
  - ListExistingLocations
  - ListInterestAccruals
  - ListPortfolioInterestAccruals
  - ListMarginCallSummaries
  - ListMarginConversions
  - GetEntityLocateAvailabilities
  - GetMarginInformation
  - GetPortfolioBuyingPower
  - GetPortfolioCreditInformation
  - GetPortfolioWithdrawalPower
  - GetTieredPricingFees
  - CreateNewLocates
- Adding support for new Positions endpoints
  - ListAggregateEntityPositions
  - ListEntityPositions
- Adding support for new Balance endpoint
  - ListEntityBalances
- Added support for new staking endpoints
  - CreateStake
  - CreateUnstake

## [0.3.7] - 2025-MAY-01

### Fix

- Fix missing allocations on list request

## [0.3.6] - 2025-APR-30

### Added

- Added NetworkFamily to CreateWallet endpoints
- Now supports evm and solana network families

## [0.3.5] - 2025-APR-08

### Added

- Added OnchainEvmParams option to transation
