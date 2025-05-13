# Changelog

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
