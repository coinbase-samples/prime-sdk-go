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

// Network family constants
const (
	EvmNetworkFamily    = "NETWORK_FAMILY_EVM"
	SolanaNetworkFamily = "NETWORK_FAMILY_SOLANA"
)

// Asset represents a Prime asset
type Asset struct {
	Name             string     `json:"name"`
	Symbol           string     `json:"symbol"`
	DecimalPrecision string     `json:"decimal_precision"`
	TradingSupported bool       `json:"trading_supported"`
	ExplorerUrl      string     `json:"explorer_url"`
	Networks         []*Network `json:"networks"`
}

// Network represents network information for an asset
type Network struct {
	Network                *NetworkDetails `json:"network"`
	Name                   string          `json:"name"`
	MaxDecimals            string          `json:"max_decimals"`
	Default                bool            `json:"default"`
	TradingSupported       bool            `json:"trading_supported"`
	VaultSupported         bool            `json:"vault_supported"`
	PrimeCustodySupported  bool            `json:"prime_custody_supported"`
	DestinationTagRequired bool            `json:"destination_tag_required"`
	NetworkLink            string          `json:"network_link"`
	NetworkScopedSymbol    string          `json:"network_scoped_symbol"`
}

// NetworkDetails represents detailed network information
type NetworkDetails struct {
	Id   string `json:"id"`
	Type string `json:"type"`
}
