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
	"time"

	"github.com/coinbase-samples/core-go"
	"github.com/shopspring/decimal"
)

// TravelRuleWalletType represents the type of wallet for travel rule compliance
type TravelRuleWalletType string

const (
	TravelRuleWalletTypeUnspecified   TravelRuleWalletType = "TRAVEL_RULE_WALLET_TYPE_UNSPECIFIED"
	TravelRuleWalletTypeVASP          TravelRuleWalletType = "TRAVEL_RULE_WALLET_TYPE_VASP"
	TravelRuleWalletTypeSelfCustodied TravelRuleWalletType = "TRAVEL_RULE_WALLET_TYPE_SELF_CUSTODIED"
)

// NaturalPersonName represents natural person name components
type NaturalPersonName struct {
	FirstName  string `json:"first_name,omitempty"`
	MiddleName string `json:"middle_name,omitempty"`
	LastName   string `json:"last_name,omitempty"`
}

// DetailedAddress represents detailed address information
type DetailedAddress struct {
	Address1    string `json:"address_1,omitempty"`
	Address2    string `json:"address_2,omitempty"`
	Address3    string `json:"address_3,omitempty"`
	City        string `json:"city,omitempty"`
	State       string `json:"state,omitempty"`
	CountryCode string `json:"country_code,omitempty"`
	PostalCode  string `json:"postal_code,omitempty"`
}

// TravelRuleDate represents a date for travel rule (year, month, day)
type TravelRuleDate struct {
	Year  int32 `json:"year,omitempty"`
	Month int32 `json:"month,omitempty"`
	Day   int32 `json:"day,omitempty"`
}

// TravelRuleParty represents a party in a travel rule transaction
type TravelRuleParty struct {
	Name              string               `json:"name,omitempty"`
	NaturalPersonName *NaturalPersonName   `json:"natural_person_name,omitempty"`
	Address           *DetailedAddress     `json:"address,omitempty"`
	WalletType        TravelRuleWalletType `json:"wallet_type,omitempty"`
	VaspId            string               `json:"vasp_id,omitempty"`
	VaspName          string               `json:"vasp_name,omitempty"`
	PersonalId        string               `json:"personal_id,omitempty"`
	DateOfBirth       *TravelRuleDate      `json:"date_of_birth,omitempty"`
}

// EstimatedNetworkFees represents estimated network fees for a transaction
type EstimatedNetworkFees struct {
	LowerBound string `json:"lower_bound,omitempty"`
	UpperBound string `json:"upper_bound,omitempty"`
}

// MatchMetadata represents metadata for matched transactions
type MatchMetadata struct {
	ReferenceId    string `json:"reference_id,omitempty"`
	SettlementDate string `json:"settlement_date,omitempty"`
}

// TransactionMetadata represents additional metadata for a transaction
type TransactionMetadata struct {
	MatchMetadata *MatchMetadata `json:"match_metadata,omitempty"`
}

// AssetChange represents a change in asset for a transaction
type AssetChange struct {
	Symbol string `json:"symbol,omitempty"`
}

// RiskAssessment represents risk assessment results for a transaction
type RiskAssessment struct {
	ComplianceRiskDetected bool `json:"compliance_risk_detected"`
	SecurityRiskDetected   bool `json:"security_risk_detected"`
}

// OnchainDetail represents on-chain details for a transaction
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

// Transfer represents a transfer from or to in a transaction
type Transfer struct {
	Type              string `json:"type"`
	Value             string `json:"value"`
	Address           string `json:"address"`
	AccountIdentifier string `json:"account_identifier"`
}

// ValueNum converts the transfer value string to a decimal
func (tr Transfer) ValueNum() (amount decimal.Decimal, err error) {
	amount, err = core.StrToNum(tr.Value)
	if err != nil {
		err = fmt.Errorf("invalid transfer value: %s - type: %s - msg: %v", tr.Value, tr.Type, err)
	}
	return
}

// Transaction represents a Prime transaction
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
