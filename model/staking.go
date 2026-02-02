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

// StakeType represents the type of staking operation
type StakeType string

const (
	StakeTypeUnspecified    StakeType = "STAKE_TYPE_UNSPECIFIED"
	StakeTypeInitialDeposit StakeType = "STAKE_TYPE_INITIAL_DEPOSIT"
	StakeTypeTopUp          StakeType = "STAKE_TYPE_TOP_UP"
)

// UnstakeType represents the type of unstaking operation
type UnstakeType string

const (
	UnstakeTypeUnspecified UnstakeType = "UNSTAKE_TYPE_UNSPECIFIED"
	UnstakeTypePartial     UnstakeType = "UNSTAKE_TYPE_PARTIAL"
	UnstakeTypeFull        UnstakeType = "UNSTAKE_TYPE_FULL"
)

// EstimateType represents the type of estimate for unstaking
type EstimateType string

const (
	EstimateTypeUnspecified EstimateType = "UNSPECIFIED"
	EstimateTypeLive        EstimateType = "LIVE"
	EstimateTypeInterim     EstimateType = "INTERIM"
)

// ValidatorStatus represents the status of a validator
type ValidatorStatus string

const (
	ValidatorStatusUnspecified ValidatorStatus = "VALIDATOR_STATUS_UNSPECIFIED"
	ValidatorStatusPending     ValidatorStatus = "VALIDATOR_STATUS_PENDING"
	ValidatorStatusActive      ValidatorStatus = "VALIDATOR_STATUS_ACTIVE"
	ValidatorStatusExiting     ValidatorStatus = "VALIDATOR_STATUS_EXITING"
	ValidatorStatusExited      ValidatorStatus = "VALIDATOR_STATUS_EXITED"
	ValidatorStatusWithdrawn   ValidatorStatus = "VALIDATOR_STATUS_WITHDRAWN"
)

// StakingStatus represents the status of a staking operation
type StakingStatus struct {
	Amount                string    `json:"amount"`
	StakeType             StakeType `json:"stake_type"`
	EstimatedStakeDate    string    `json:"estimated_stake_date"`
	EstimatedHoursToStake int64     `json:"estimated_hours_to_stake"`
	RequestedAt           string    `json:"requested_at"`
}

// UnstakingStatus represents the status of an unstaking operation (from API spec)
type UnstakingStatus struct {
	Amount              string       `json:"amount"`
	UnstakeType         UnstakeType  `json:"unstake_type"`
	FinishingAt         string       `json:"finishing_at"`
	RemainingHours      int64        `json:"remaining_hours"`
	RequestedAt         string       `json:"requested_at"`
	EstimateType        EstimateType `json:"estimate_type"`
	EstimateDescription string       `json:"estimate_description"`
}

// UnstakeStatus represents the status of an unstake operation (legacy)
type UnstakeStatus struct {
	Amount              string       `json:"amount"`
	EstimateType        EstimateType `json:"estimate_type"`
	EstimateDescription string       `json:"estimate_description"`
	UnstakeType         UnstakeType  `json:"unstake_type"`
	FinishingAt         string       `json:"finishing_at"`
	RemainingHours      int          `json:"remaining_hours"`
	RequestedAt         string       `json:"requested_at"`
}

// ValidatorStakingInfo represents staking information for a validator
type ValidatorStakingInfo struct {
	ValidatorAddress string           `json:"validator_address"`
	Statuses         []*StakingStatus `json:"statuses"`
}

// ValidatorUnstakingInfo represents unstaking information for a validator
type ValidatorUnstakingInfo struct {
	ValidatorAddress string             `json:"validator_address"`
	Statuses         []*UnstakingStatus `json:"statuses"`
}

// UnstakeValidator represents a validator with unstake statuses (legacy)
type UnstakeValidator struct {
	ValidatorAddress string           `json:"validator_address"`
	Statuses         []*UnstakeStatus `json:"statuses"`
}

// TransactionValidator represents a transaction-to-validator association
type TransactionValidator struct {
	TransactionId    string          `json:"transaction_id"`
	ValidatorAddress string          `json:"validator_address"`
	ValidatorStatus  ValidatorStatus `json:"validator_status"`
}

// PortfolioStakingMetadata contains optional metadata for portfolio staking operations
type PortfolioStakingMetadata struct {
	ExternalId string `json:"external_id,omitempty"`
}
