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

// AdvancedTransferState represents the lifecycle state of an advanced transfer.
type AdvancedTransferState string

const (
	AdvancedTransferStateCreated    AdvancedTransferState = "ADVANCED_TRANSFER_STATE_CREATED"
	AdvancedTransferStateProcessing AdvancedTransferState = "ADVANCED_TRANSFER_STATE_PROCESSING"
	AdvancedTransferStateDone       AdvancedTransferState = "ADVANCED_TRANSFER_STATE_DONE"
	AdvancedTransferStateCancelled  AdvancedTransferState = "ADVANCED_TRANSFER_STATE_CANCELLED"
	AdvancedTransferStateFailed     AdvancedTransferState = "ADVANCED_TRANSFER_STATE_FAILED"
	AdvancedTransferStateExpired    AdvancedTransferState = "ADVANCED_TRANSFER_STATE_EXPIRED"
)

// AdvancedTransferType specifies the type of advanced transfer.
type AdvancedTransferType string

const (
	AdvancedTransferTypeBlindMatch AdvancedTransferType = "ADVANCED_TRANSFER_TYPE_BLIND_MATCH"
)

// TransferLocationType identifies the kind of transfer location.
type TransferLocationType string

const (
	TransferLocationTypeUnknown           TransferLocationType = "TRANSFER_LOCATION_TYPE_UNKNOWN"
	TransferLocationTypePaymentMethod     TransferLocationType = "PAYMENT_METHOD"
	TransferLocationTypeWallet            TransferLocationType = "WALLET"
	TransferLocationTypeAddress           TransferLocationType = "ADDRESS"
	TransferLocationTypeOther             TransferLocationType = "OTHER"
	TransferLocationTypeMultipleAddresses TransferLocationType = "MULTIPLE_ADDRESSES"
	TransferLocationTypeCounterpartyId    TransferLocationType = "COUNTERPARTY_ID"
)

// TransferLocation represents a source or target location in a fund movement.
type TransferLocation struct {
	Type              TransferLocationType `json:"type,omitempty"`
	Value             string               `json:"value,omitempty"`
	Address           string               `json:"address,omitempty"`
	AccountIdentifier string               `json:"account_identifier,omitempty"`
}

// FundMovement represents a single movement of funds between two counterparties.
type FundMovement struct {
	Id       string            `json:"id,omitempty"`
	Source   *TransferLocation `json:"source,omitempty"`
	Target   *TransferLocation `json:"target,omitempty"`
	Currency string            `json:"currency,omitempty"`
	Amount   string            `json:"amount,omitempty"`
}

// BlindMatchMetadata contains metadata specific to blind match advanced transfers.
type BlindMatchMetadata struct {
	ReferenceId    string `json:"reference_id,omitempty"`
	SettlementDate string `json:"settlement_date,omitempty"`
	TradeDate      string `json:"trade_date,omitempty"`
	SettlementTime string `json:"settlement_time,omitempty"`
}

// AdvancedTransfer represents a complex transfer operation such as a blind match settlement.
type AdvancedTransfer struct {
	Id                 string                `json:"id,omitempty"`
	Type               AdvancedTransferType  `json:"type,omitempty"`
	State              AdvancedTransferState `json:"state,omitempty"`
	FundMovements      []*FundMovement       `json:"fund_movements,omitempty"`
	BlindMatchMetadata *BlindMatchMetadata   `json:"blind_match_metadata,omitempty"`
}
