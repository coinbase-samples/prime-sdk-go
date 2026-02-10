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

// InvoiceState represents the state of an invoice
type InvoiceState string

const (
	InvoiceStateUnspecified   InvoiceState = "INVOICE_STATE_UNSPECIFIED"
	InvoiceStateImported      InvoiceState = "INVOICE_STATE_IMPORTED"
	InvoiceStateBilled        InvoiceState = "INVOICE_STATE_BILLED"
	InvoiceStatePartiallyPaid InvoiceState = "INVOICE_STATE_PARTIALLY_PAID"
	InvoiceStatePaid          InvoiceState = "INVOICE_STATE_PAID"
)

// InvoiceType represents the type of an invoice item
type InvoiceType string

const (
	InvoiceTypeUnspecified   InvoiceType = "INVOICE_TYPE_UNSPECIFIED"
	InvoiceTypeAucFee        InvoiceType = "INVOICE_TYPE_AUC_FEE"
	InvoiceTypeMinimumFee    InvoiceType = "INVOICE_TYPE_MINIMUM_FEE"
	InvoiceTypeWithdrawalFee InvoiceType = "INVOICE_TYPE_WITHDRAWAL_FEE"
	InvoiceTypeNewWalletFee  InvoiceType = "INVOICE_TYPE_NEW_WALLET_FEE"
	InvoiceTypeStakingFee    InvoiceType = "INVOICE_TYPE_STAKING_FEE"
)

type Invoice struct {
	Id            string         `json:"id"`
	BillingYear   int32          `json:"billing_year"`
	BillingMonth  int32          `json:"billing_month"`
	DueDate       string         `json:"due_date"`
	InvoiceNumber string         `json:"invoice_number"`
	State         InvoiceState   `json:"state"`
	UsdAmountPaid float64        `json:"usd_amount_paid"`
	UsdAmountOwed float64        `json:"usd_amount_owed"`
	Items         []*InvoiceItem `json:"invoice_items"`
}

type InvoiceItem struct {
	Description    string      `json:"description"`
	CurrencySymbol string      `json:"currency_symbol"`
	InvoiceType    InvoiceType `json:"invoice_type"`
	Rate           float64     `json:"rate"`
	Quantity       float64     `json:"quantity"`
	Price          float64     `json:"price"`
	AverageAuc     float64     `json:"average_auc"`
	Total          float64     `json:"total"`
}
