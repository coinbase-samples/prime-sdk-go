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

	"github.com/coinbase-samples/core-go"
	"github.com/shopspring/decimal"
)

const (
	BalanceTypeTrading = "TRADING_BALANCES"
	BalanceTypeVault   = "VAULT_BALANCES"
	BalanceTypeTotal   = "TOTAL_BALANCES"
)

type Balance struct {
	Symbol               string `json:"symbol"`
	Amount               string `json:"amount"`
	Holds                string `json:"holds"`
	BondedAmount         string `json:"bonded_amount"`
	ReservedAmount       string `json:"reserved_amount"`
	UnbondingAmount      string `json:"unbonding_amount"`
	UnvestedAmount       string `json:"unvested_amount"`
	PendingRewardsAmount string `json:"pending_rewards_amount"`
	PastRewardsAmount    string `json:"past_rewards_amount"`
	BondableAmount       string `json:"bondable_amount"`
	WithdrawableAmount   string `json:"withdrawable_amount"`
}

func (b Balance) AmountNum() (amount decimal.Decimal, err error) {
	amount, err = core.StrToNum(b.Amount)
	if err != nil {
		err = fmt.Errorf("Invalid asset amount: %s - symbol: %s - msg: %v", b.Amount, b.Symbol, err)
	}
	return
}

func (b Balance) HoldsNum() (holds decimal.Decimal, err error) {
	holds, err = core.StrToNum(b.Holds)
	if err != nil {
		err = fmt.Errorf("Invalid asset holds: %s - symbol: %s - msg: %v", b.Holds, b.Symbol, err)
	}
	return
}

type BalanceWithHolds struct {
	Total string `json:"total"`
	Holds string `json:"holds"`
}
