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

import "time"

// WalletVisibility represents the visibility state of a wallet
type WalletVisibility string

const (
	WalletVisibilityUnspecified WalletVisibility = "WALLET_VISIBILITY_UNSPECIFIED"
	WalletVisibilityVisible     WalletVisibility = "WALLET_VISIBILITY_VISIBLE"
	WalletVisibilityHidden      WalletVisibility = "WALLET_VISIBILITY_HIDDEN"
)

// Wallet type constants
const (
	WalletTypeVault   = "VAULT"
	WalletTypeTrading = "TRADING"
	WalletTypeOnchain = "ONCHAIN"
	WalletTypeOther   = "WALLET_TYPE_OTHER"
)

// Wallet deposit type constants
const (
	WalletDepositTypeWire   = "WIRE"
	WalletDepositTypeSwift  = "SWIFT"
	WalletDepositTypeCrypto = "CRYPTO"
)

// Wallet represents a Prime wallet
type Wallet struct {
	Id         string           `json:"id"`
	Type       string           `json:"type"`
	Name       string           `json:"name"`
	Address    string           `json:"address"`
	Visibility WalletVisibility `json:"visibility"`
	Symbol     string           `json:"symbol"`
	Created    time.Time        `json:"created_at"`
	Network    *NetworkDetails  `json:"network"`
}

// BlockchainAddress represents a blockchain address
type BlockchainAddress struct {
	Address           string          `json:"address"`
	AccountIdentifier string          `json:"account_identifier"`
	Network           *NetworkDetails `json:"network"`
}

// CryptoDepositInstructions represents instructions for crypto deposits
type CryptoDepositInstructions struct {
	Id                string `json:"id"`
	Name              string `json:"name"`
	Type              string `json:"type"`
	Address           string `json:"address"`
	AccountIdentifier string `json:"account_identifier"`
}

// FiatDepositInstructions represents instructions for fiat deposits
type FiatDepositInstructions struct {
	Id            string `json:"id"`
	Name          string `json:"name"`
	Type          string `json:"type"`
	AccountNumber string `json:"account_number"`
	RoutingNumber string `json:"routing_number"`
	ReferenceCode string `json:"reference_code"`
}
