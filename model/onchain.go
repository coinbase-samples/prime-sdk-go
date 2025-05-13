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

type OnchainNetworkType string

const (
	OnchainNetworkTypeUnspecified OnchainNetworkType = "NETWORK_TYPE_UNSPECIFIED"
	OnchainNetworkTypeEvm         OnchainNetworkType = "NETWORK_TYPE_EVM"
	OnchainNetworkTypeSolana      OnchainNetworkType = "NETWORK_TYPE_SOLANA"
)

type OnchainActivityType string

const (
	OnchainActivityTypeUnknown                 OnchainActivityType = "ACTIVITY_TYPE_UNKNOWN"
	OnchainActivityTypeGovernanceVote          OnchainActivityType = "ACTIVITY_TYPE_GOVERNANCE_VOTE"
	OnchainActivityTypeInvitiation             OnchainActivityType = "ACTIVITY_TYPE_INVITATION"
	OnchainActivityTypeWalletChange            OnchainActivityType = "ACTIVITY_TYPE_WALLET_CHANGE"
	OnchainActivityTypeApiKeyChange            OnchainActivityType = "ACTIVITY_TYPE_API_KEY_CHANGE"
	OnchainActivityTypeSettingsChange          OnchainActivityType = "ACTIVITY_TYPE_SETTINGS_CHANGE"
	OnchainActivityTypeBillingPreferenceChange OnchainActivityType = "ACTIVITY_TYPE_BILLING_PREFERENCE_CHANGE"
	OnchainActivityTypePaymentMethodChange     OnchainActivityType = "ACTIVITY_TYPE_PAYMENT_METHOD_CHANGE"
	OnchainActivityTypeWithdrawal              OnchainActivityType = "ACTIVITY_TYPE_WITHDRAWAL"
	OnchainActivityTypeDeposit                 OnchainActivityType = "ACTIVITY_TYPE_DEPOSIT"
	OnchainActivityTypeCreateWallet            OnchainActivityType = "ACTIVITY_TYPE_CREATE_WALLET"
	OnchainActivityTypeRemoveWallet            OnchainActivityType = "ACTIVITY_TYPE_REMOVE_WALLET"
	OnchainActivityTypeUpdateWallet            OnchainActivityType = "ACTIVITY_TYPE_UPDATE_WALLET"
	OnchainActivityTypeCastVote                OnchainActivityType = "ACTIVITY_TYPE_CAST_VOTE"
	OnchainActivityTypeEnableVoting            OnchainActivityType = "ACTIVITY_TYPE_ENABLE_VOTING"
	OnchainActivityTypeStake                   OnchainActivityType = "ACTIVITY_TYPE_STAKE"
	OnchainActivityTypeUnstake                 OnchainActivityType = "ACTIVITY_TYPE_UNSTAKE"
	OnchainActivityTypeChangeValidator         OnchainActivityType = "ACTIVITY_TYPE_CHANGE_VALIDATOR"
	OnchainActivityTypeRestake                 OnchainActivityType = "ACTIVITY_TYPE_RESTAKE"
	OnchainActivityTypeAddressBook             OnchainActivityType = "ACTIVITY_TYPE_ADDRESS_BOOK"
	OnchainActivityTypeTeamMembes              OnchainActivityType = "ACTIVITY_TYPE_TEAM_MEMBERS"
	OnchainActivityTypeBilling                 OnchainActivityType = "ACTIVITY_TYPE_BILLING"
	OnchainActivityTypeSecurity                OnchainActivityType = "ACTIVITY_TYPE_SECURITY"
	OnchainActivityTypeApi                     OnchainActivityType = "ACTIVITY_TYPE_API"
	OnchainActivityTypeSettings                OnchainActivityType = "ACTIVITY_TYPE_SETTINGS"
	OnchainActivityTypeSmartContract           OnchainActivityType = "ACTIVITY_TYPE_SMART_CONTRACT"
	OnchainActivityTypeUserChangeRequestNoPas  OnchainActivityType = "ACTIVITY_TYPE_USER_CHANGE_REQUEST_NO_PAS"
	OnchainActivityTypeWeb3Transaction         OnchainActivityType = "ACTIVITY_TYPE_WEB3_TRANSACTION"
	OnchainActivityTypeWeb3Message             OnchainActivityType = "ACTIVITY_TYPE_WEB3_MESSAGE"
	OnchainActivityTypeClaimRewards            OnchainActivityType = "ACTIVITY_TYPE_CLAIM_REWARDS"
)

type OnchainAddressGroup struct {
	Id          string             `json:"id"`
	Name        string             `json:"name"`
	NetworkType OnchainNetworkType `json:"network_type"`
	Addresses   []*OnchainAddress  `json:"addresses"`
}

type OnchainAddress struct {
	Name     string   `json:"name"`
	Address  string   `json:"address"`
	ChainIds []string `json:"chain_ids"`
}

type OnchainTransaction struct {
	RawUnsignedTransaction string            `json:"raw_unsigned_txn"`
	Rpc                    *OnchainRpc       `json:"rpc,omitempty"`
	EvmParams              *OnchainEvmParams `json:"evm_params,omitempty"`
}

type OnchainRpc struct {
	Url           string `json:"url,omitempty"`
	SkipBroadcast bool   `json:"skip_broadcast"`
}

type OnchainEvmParams struct {
	DisableDynamicGas     bool   `json:"disable_dynamic_gas"`
	DisableDynamicNonce   bool   `json:"disable_dynamic_nonce"`
	ReplacedTransactionId string `json:"replaced_transaction_id,omitempty"`
	ChainId               string `json:"chain_id"`
}
