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

package utils

import (
	"github.com/coinbase-samples/prime-sdk-go/model"
	"github.com/shopspring/decimal"
)

func CalculateOrderSize(
	product *model.Product,
	amount decimal.Decimal,
	holds decimal.Decimal,
) (orderSize decimal.Decimal, err error) {

	var (
		baseMin       decimal.Decimal
		baseMax       decimal.Decimal
		baseIncrement decimal.Decimal
	)

	if baseMin, err = product.BaseMinSizeNum(); err != nil {
		return
	}

	if baseMax, err = product.BaseMaxSizeNum(); err != nil {
		return
	}

	if baseIncrement, err = product.BaseIncrementNum(); err != nil {
		return
	}

	availableAmount := amount.Sub(holds)

	if availableAmount.IsZero() {
		orderSize = availableAmount
		return
	}

	if availableAmount.IsNegative() {

		orderSize = decimal.NewFromInt(0)

	} else {

		orderSize = AdjustOrderSize(availableAmount, baseMin, baseMax, baseIncrement)

	}

	return
}

func AdjustOrderSize(amount, baseMin, baseMax, baseIncrement decimal.Decimal) decimal.Decimal {

	if amount.Cmp(baseMax) > 0 {
		return baseMax
	}

	if amount.Cmp(baseMin) < 0 {
		return decimal.NewFromFloat(0)
	}

	quo, rem := amount.QuoRem(baseIncrement, 0)

	if rem.IsZero() {
		return amount
	}

	return quo.Floor().Mul(baseIncrement)
}
