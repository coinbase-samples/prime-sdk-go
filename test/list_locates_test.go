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

package test

import (
	"context"
	"testing"
	"time"

	"github.com/coinbase-samples/prime-sdk-go/financing"
)

func TestListLocates(t *testing.T) {

	c, err := newLiveTestClient()
	if err != nil {
		t.Fatal(err)
	}

	service := financing.NewFinancingService(c)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	response, err := service.ListLocates(ctx, &financing.ListLocatesRequest{
		PortfolioId: c.Credentials().PortfolioId,
		LocateDate:  time.Now().Format("2006-01-02"),
	})

	if err != nil {
		t.Fatal(err)
	}

	if response == nil {
		t.Fatal(err)
	}

	if len(response.Locates) == 0 {
		return
	}

	for _, locate := range response.Locates {

		if len(locate.LocateId) == 0 {
			t.Error("expected locate id in locate")
		}

		if len(locate.Symbol) == 0 {
			t.Error("expected symbol in locate")
		}
	}
}
