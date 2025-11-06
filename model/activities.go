/**
 * Copyright 2025-present Coinbase Global, Inc.
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

type Activity struct {
	Id                  string                `json:"id"`
	ReferenceId         string                `json:"reference_id"`
	Category            string                `json:"category"`
	PrimaryType         string                `json:"type"`
	SecondaryType       string                `json:"secondary_type"`
	Status              string                `json:"status"`
	CreatedBy           string                `json:"created_by"`
	Title               string                `json:"title"`
	Description         string                `json:"description,omitempty"`
	UserActions         []*UserAction         `json:"user_actions,omitempty"`
	AccountMetadata     *AccountMetadata      `json:"account_metadata,omitempty"`
	OrdersMetadata      *OrdersMetadata       `json:"orders_metadata,omitempty"`
	TransactionMetadata *TransactionsMetadata `json:"transaction_metadata,omitempty"`
	Symbols             []string              `json:"symbols,omitempty"`
	Created             string                `json:"created_at"`
	Updated             string                `json:"updated_at"`
	HierarchyType       string                `json:"hierarchy_type,omitempty"`
}

type UserAction struct {
	Action    string `json:"action"`
	UserId    string `json:"user_id"`
	Timestamp string `json:"timestamp"`
}

type AccountMetadata struct {
	Consensus *Consensus `json:"consensus"`
}

type TransactionsMetadata struct {
	Consensus *Consensus `json:"consensus"`
}

type OrdersMetadata struct{}

type Consensus struct {
	ApprovalDeadline string `json:"approval_deadline"`
	PassedConsensus  bool   `json:"has_passed_consensus"`
}
