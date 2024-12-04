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

package test

import (
	"context"
	"testing"
	"time"

	"github.com/coinbase-samples/prime-sdk-go/commission"
)

func TestPortfolioCommission(t *testing.T) {

	c, err := newLiveTestClient()
	if err != nil {
		t.Fatal(err)
	}

	service := commission.NewCommissionService(c)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	response, err := service.GetPortfolioCommission(ctx, &commission.GetPortfolioCommissionRequest{
		PortfolioId: c.Credentials().PortfolioId,
	})

	if err != nil {
		t.Fatal(err)
	}

	if response == nil {
		t.Fatal(err)
	}

	if response.Commission == nil {
		t.Fatal("expected commission in get")
	}
}
