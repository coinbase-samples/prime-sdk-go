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

import (
	"fmt"
	"time"

	"github.com/coinbase-samples/core-go"
	"github.com/shopspring/decimal"
)

type Product struct {
	Id                string             `json:"id"`
	BaseIncrement     string             `json:"base_increment"`
	QuoteIncrement    string             `json:"quote_increment"`
	BaseMinSize       string             `json:"base_min_size"`
	BaseMaxSize       string             `json:"base_max_size"`
	QuoteMinSize      string             `json:"quote_min_size"`
	QuoteMaxSize      string             `json:"quote_max_size"`
	Permissions       []string           `json:"permissions"`
	PriceIncrement    string             `json:"price_increment"`
	RfqProductDetails *RfqProductDetails `json:"rfq_product_details"`
}

type RfqProductDetails struct {
	Tradable        bool   `json:"tradable"`
	MinNotionalSize string `json:"min_notional_size"`
	MaxNotionalSize string `json:"max_notional_size"`
	MinBaseSize     string `json:"min_base_size"`
	MaxBaseSize     string `json:"max_base_size"`
	MinQuoteSize    string `json:"min_quote_size"`
	MaxQuoteSize    string `json:"max_quote_size"`
}

func (p Product) BaseMinSizeNum() (amount decimal.Decimal, err error) {
	amount, err = core.StrToNum(p.BaseMinSize)
	if err != nil {
		err = fmt.Errorf("invalid base min: %s - id: %s - err: %w", p.BaseMinSize, p.Id, err)
	}
	return
}

func (p Product) BaseMaxSizeNum() (amount decimal.Decimal, err error) {
	amount, err = core.StrToNum(p.BaseMaxSize)
	if err != nil {
		err = fmt.Errorf("invalid base max: %s - id: %s - err: %v", p.BaseMaxSize, p.Id, err)
	}
	return
}

func (p Product) BaseIncrementNum() (amount decimal.Decimal, err error) {
	amount, err = core.StrToNum(p.BaseIncrement)
	if err != nil {
		err = fmt.Errorf("invalid base increment: %s - id: %s - msg: %w", p.BaseIncrement, p.Id, err)
	}
	return
}

func (p Product) QuoteMinSizeNum() (amount decimal.Decimal, err error) {
	amount, err = core.StrToNum(p.QuoteMinSize)
	if err != nil {
		err = fmt.Errorf("invalid quote min: %s - id: %s - err: %w", p.QuoteMinSize, p.Id, err)
	}
	return
}

func (p Product) QuoteMaxSizeNum() (amount decimal.Decimal, err error) {
	amount, err = core.StrToNum(p.QuoteMaxSize)
	if err != nil {
		err = fmt.Errorf("invalid quote max: %s - id: %s - err: %v", p.QuoteMaxSize, p.Id, err)
	}
	return
}

func (p Product) QuoteIncrementNum() (amount decimal.Decimal, err error) {
	amount, err = core.StrToNum(p.QuoteIncrement)
	if err != nil {
		err = fmt.Errorf("invalid quote increment: %s - id: %s - msg: %w", p.QuoteIncrement, p.Id, err)
	}
	return
}

type CandlesGranularity string

const (
	OneMinuteCandlesGranularity      CandlesGranularity = "ONE_MINUTE"
	FiveMinutesCandlesGranularity    CandlesGranularity = "FIVE_MINUTES"
	FifteenMinutesCandlesGranularity CandlesGranularity = "FIFTEEN_MINUTES"
	ThirtyMinutesCandlesGranularity  CandlesGranularity = "THIRTY_MINUTE"
	OneHourCandlesGranularity        CandlesGranularity = "ONE_HOUR"
	TwoHoursCandlesGranularity       CandlesGranularity = "TWO_HOUR"
	FourHoursCandlesGranularity      CandlesGranularity = "FOUR_HOUR"
	SixHoursCandlesGranularity       CandlesGranularity = "SIX_HOURS"
	OneDayCandlesGranularity         CandlesGranularity = "ONE_DAY"
)

type Candle struct {
	Timestamp time.Time `json:"timestamp"`
	Open      string    `json:"open"`
	High      string    `json:"high"`
	Low       string    `json:"low"`
	Close     string    `json:"close"`
	Volume    string    `json:"volume"`
}
