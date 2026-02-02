/**
 * Copyright 2023-present Coinbase Global, Inc.
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

type WalletVisibility string

const (
	WalletVisibilityUnspecified WalletVisibility = "WALLET_VISIBILITY_UNSPECIFIED"
	WalletVisibilityVisible     WalletVisibility = "WALLET_VISIBILITY_VISIBLE"
	WalletVisibilityHidden      WalletVisibility = "WALLET_VISIBILITY_HIDDEN"
)

const (
	WalletTypeVault   = "VAULT"
	WalletTypeTrading = "TRADING"
	WalletTypeOnchain = "ONCHAIN"
	WalletTypeOther   = "WALLET_TYPE_OTHER"

	WalletDepositTypeWire   = "WIRE"
	WalletDepositTypeSwift  = "SWIFT"
	WalletDepositTypeCrypto = "CRYPTO"
)

const (
	EvmNetworkFamily    = "NETWORK_FAMILY_EVM"
	SolanaNetworkFamily = "NETWORK_FAMILY_SOLANA"
)

type ErrorMessage struct {
	Value string `json:"message"`
}

type EntityPaymentMethod struct {
	Id                string `json:"id"`
	Symbol            string `json:"symbol"`
	PaymentMethodType string `json:"payment_method_type"`
	Name              string `json:"name"`
	AccountNumber     string `json:"account_number"`
	BankCode          string `json:"bank_code"`
	BankName          string `json:"bank_name,omitempty"`
	BankName2         string `json:"bank_name_2,omitempty"`
}

type Portfolio struct {
	Id             string `json:"id"`
	Name           string `json:"name"`
	EntityId       string `json:"entity_id"`
	OrganizationId string `json:"organization_id"`
}

type Counterparty struct {
	CounterpartyId string `json:"counterparty_id"`
}

type PaginationParams struct {
	Cursor        string `json:"cursor"`
	Limit         int32  `json:"limit"`
	SortDirection string `json:"sort_direction"`
}

type User struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	EntityId    string `json:"entity_id"`
	PortfolioId string `json:"portfolio_id,omitempty"`
	Role        string `json:"role"`
}

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

type BlockchainAddress struct {
	Address           string          `json:"address"`
	AccountIdentifier string          `json:"account_identifier"`
	Network           *NetworkDetails `json:"network"`
}

type Pagination struct {
	NextCursor    string `json:"next_cursor"`
	SortDirection string `json:"sort_direction"`
	HasNext       bool   `json:"has_next"`
}

type Commission struct {
	Type          string `json:"type"`
	Rate          string `json:"rate"`
	TradingVolume string `json:"trading_volume"`
}

func (p Commission) RateNum() (rate decimal.Decimal, err error) {
	rate, err = core.StrToNum(p.Rate)
	if err != nil {
		err = fmt.Errorf("Invalid commission rate: %s - err: %w", p.Rate, err)
	}
	return
}

type Asset struct {
	Name             string     `json:"name"`
	Symbol           string     `json:"symbol"`
	DecimalPrecision string     `json:"decimal_precision"`
	TradingSupported bool       `json:"trading_supported"`
	ExplorerUrl      string     `json:"explorer_url"`
	Networks         []*Network `json:"networks"`
}

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

type NetworkDetails struct {
	Id   string `json:"id"`
	Type string `json:"type"`
}

type CryptoDepositInstructions struct {
	Id                string `json:"id"`
	Name              string `json:"name"`
	Type              string `json:"type"`
	Address           string `json:"address"`
	AccountIdentifier string `json:"account_identifier"`
}

type FiatDepositInstructions struct {
	Id            string `json:"id"`
	Name          string `json:"name"`
	Type          string `json:"type"`
	AccountNumber string `json:"account_number"`
	RoutingNumber string `json:"routing_number"`
	ReferenceCode string `json:"reference_code"`
}

type PostTradeCreditAmountDue struct {
	Currency string    `json:"currency"`
	Amount   string    `json:"amount"`
	DueDate  time.Time `json:"due_date"`
}

type PostTradeCredit struct {
	Id                     string                      `json:"portfolio_id"`
	Currency               string                      `json:"currency"`
	Limit                  string                      `json:"limit"`
	Utilized               string                      `json:"utilized"`
	Available              string                      `json:"available"`
	Frozen                 bool                        `json:"frozen"`
	AmountsDue             []*PostTradeCreditAmountDue `json:"amounts_due"`
	FrozenReason           string                      `json:"frozen_reason"`
	Enabled                bool                        `json:"enabled"`
	AdjustedCreditUtilized string                      `json:"adjusted_credit_utilized"`
	AdjustedEquity         string                      `json:"adjusted_portfolio_equity"`
}
