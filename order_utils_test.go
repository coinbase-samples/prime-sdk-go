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

package prime

import (
	"testing"

	"github.com/shopspring/decimal"
)

func TestAdjustOrderSize(t *testing.T) {

	cases := []struct {
		description   string
		amount        decimal.Decimal
		baseMin       decimal.Decimal
		baseMax       decimal.Decimal
		baseIncrement decimal.Decimal
		expected      decimal.Decimal
	}{
		{
			description:   "AdjustOrder0",
			amount:        decimal.NewFromFloat(0.500937729),
			baseMin:       decimal.NewFromFloat(0.008),
			baseMax:       decimal.NewFromFloat(780000),
			baseIncrement: decimal.NewFromFloat(0.001),
			expected:      decimal.NewFromFloat(0.5),
		},
		{
			description:   "AdjustOrder1",
			amount:        decimal.NewFromFloat(0.501937729),
			baseMin:       decimal.NewFromFloat(0.008),
			baseMax:       decimal.NewFromFloat(780000),
			baseIncrement: decimal.NewFromFloat(0.001),
			expected:      decimal.NewFromFloat(0.501),
		},
		{
			description:   "AdjustOrder2",
			amount:        decimal.NewFromFloat(0.007),
			baseMin:       decimal.NewFromFloat(0.008),
			baseMax:       decimal.NewFromFloat(780000),
			baseIncrement: decimal.NewFromFloat(0.001),
			expected:      decimal.NewFromFloat(0),
		},
		{
			description:   "AdjustOrder3",
			amount:        decimal.NewFromFloat(500.09798),
			baseMin:       decimal.NewFromFloat(0.008),
			baseMax:       decimal.NewFromFloat(780000),
			baseIncrement: decimal.NewFromFloat(0.001),
			expected:      decimal.NewFromFloat(500.097),
		},
		{
			description:   "AdjustOrder4",
			amount:        decimal.NewFromFloat(20.019136794133264),
			baseMin:       decimal.NewFromFloat(0.1),
			baseMax:       decimal.NewFromFloat(770000),
			baseIncrement: decimal.NewFromFloat(0.01),
			expected:      decimal.NewFromFloat(20.01),
		},
		{
			description:   "AdjustOrder5",
			amount:        decimal.NewFromFloat(20.019136794133264),
			baseMin:       decimal.NewFromFloat(0.1),
			baseMax:       decimal.NewFromFloat(20),
			baseIncrement: decimal.NewFromFloat(0.01),
			expected:      decimal.NewFromFloat(20),
		},
		{
			description:   "AdjustOrder6",
			amount:        decimal.NewFromFloat(12),
			baseMin:       decimal.NewFromFloat(11),
			baseMax:       decimal.NewFromFloat(960000),
			baseIncrement: decimal.NewFromFloat(0.180),
			expected:      decimal.NewFromFloat(11.88),
		},
		{
			description:   "AdjustOrder7",
			amount:        decimal.NewFromFloat(12),
			baseMin:       decimal.NewFromFloat(11),
			baseMax:       decimal.NewFromFloat(960000),
			baseIncrement: decimal.NewFromFloat(0.1),
			expected:      decimal.NewFromFloat(12),
		},
	}

	for _, tt := range cases {
		t.Run(tt.description, func(t *testing.T) {
			result := AdjustOrderSize(tt.amount, tt.baseMin, tt.baseMax, tt.baseIncrement)
			if result.Cmp(tt.expected) != 0 {
				t.Errorf("test: %s - expected: %v - received: %v", tt.description, tt.expected, result)
			}
		})
	}
}

/*
func TestAdjustTwapLimitPrice(t *testing.T) {

	cases := []struct {
		description    string
		price          decimal.Decimal
		quoteIncrement decimal.Decimal
		expected       decimal.Decimal
	}{
		{
			description:    "AdjustTwapLimitPrice0",
			price:          decimal.NewFromFloat(2000.2020222),
			quoteIncrement: decimal.NewFromFloat(0.01),
			expected:       decimal.NewFromFloat(2000.20),
		},
	}

	for _, tt := range cases {
		t.Run(tt.description, func(t *testing.T) {
			result := adjustTwapLimitPrice(tt.price, tt.quoteIncrement)
			if result.Cmp(tt.expected) != 0 {
				t.Errorf("test: %s - expected: %v - received: %v", tt.description, tt.expected, result)
			}
		})
	}
}
*/

func TestGenerateUniqueId(t *testing.T) {

	cases := []struct {
		description string
		params      []string
		expected    string
	}{
		{
			description: "AdjustTwapLimitPrice0",
			params:      []string{"one", "two", "three"},
			expected:    "9af04d80fd921d0b9265ab2f3b516edc",
		},
	}

	for _, tt := range cases {
		t.Run(tt.description, func(t *testing.T) {
			result := GenerateUniqueId(tt.params...)
			if result != tt.expected {
				t.Errorf("test: %s - expected: %s - received: %s", tt.description, tt.expected, result)
			}
		})
	}
}
