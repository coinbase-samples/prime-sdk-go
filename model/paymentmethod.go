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

// EntityPaymentMethod represents a payment method for an entity
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
