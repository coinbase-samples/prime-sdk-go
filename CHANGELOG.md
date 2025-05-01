# Changelog

## [0.3.7] - 2025-MAY-02

### Added

- Adding support for new Financing endpoints
  - list_existing_locations
  - list_interest_accruals
  - list_portfolio_interest_accruals
  - list_margin_call_summaries
  - list_margin_conversions
  - get_entity_locate_availabilities
  - get_margin_information
  - get_portfolio_buying_power
  - get_portfolio_credit_information
  - get_portfolio_withdrawal_power
  - get_tiered_pricing_fees
  - create_new_locates
- Adding support for new Positions endpoints
  - ListAggregateEntityPositions
  - ListEntityPositions
- Adding support for various other new endpoints
  - BalancesService
    - ListEntityBalances
- Added support for new staking endpoints
  - create_stake
  - create_unstake

## [0.3.6] - 2025-APR-30

### Added

- Added NetworkFamily to CreateWallet endpoints
- Now supports evm and solana network families

## [0.3.5] - 2025-APR-08

### Added

- Added OnchainEvmParams option to transation
