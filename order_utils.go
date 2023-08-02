package prime

import (
	"github.com/shopspring/decimal"
)

func CalculateOrderSize(
	product *Product,
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
