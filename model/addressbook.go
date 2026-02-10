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

type AddressBookType string

const (
	AddressBookTypeUnspecified    AddressBookType = "ADDRESS_BOOK_TYPE_UNSPECIFIED"
	AddressBookTypeAddress        AddressBookType = "ADDRESS_BOOK_TYPE_ADDRESS"
	AddressBookTypeCounterpartyId AddressBookType = "ADDRESS_BOOK_TYPE_COUNTERPARTY_ID"
)

type AddressBookEntry struct {
	Id                    string                   `json:"id"`
	Symbol                string                   `json:"currency_symbol"`
	Name                  string                   `json:"name"`
	Address               string                   `json:"address"`
	AccountIdentifier     string                   `json:"account_identifier"`
	AccountIdentifierName string                   `json:"account_identifier_name"`
	State                 string                   `json:"state"`
	ExplorerLink          string                   `json:"explorer_link"`
	LastUsed              time.Time                `json:"last_used_at"`
	Added                 time.Time                `json:"added_at"`
	AddedBy               *AddressBookEntryAddedBy `json:"added_by"`
	Type                  AddressBookType          `json:"type,omitempty"`
	CounterpartyId        string                   `json:"counterparty_id,omitempty"`
}

type AddressBookEntryAddedBy struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	AvatarUrl string `json:"avatar_url"`
}
