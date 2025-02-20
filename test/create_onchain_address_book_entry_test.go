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
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/coinbase-samples/prime-sdk-go/model"
	"github.com/coinbase-samples/prime-sdk-go/onchainaddressbook"
)

func TestCreateOnchainAddressBookEntry(t *testing.T) {

	if os.Getenv("PRIME_SDK_FULL_TESTS") != "enabled" {
		t.Skip()
	}

	c, err := newLiveTestClient()
	if err != nil {
		t.Fatal(err)
	}

	service := onchainaddressbook.NewOnchainAddressBookService(c)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	response, err := service.CreateOnchainAddressBookEntry(
		ctx,
		&onchainaddressbook.CreateOnchainAddressBookEntryRequest{
			PortfolioId: c.Credentials().PortfolioId,
			AddressGroup: &model.OnchainAddressGroup{
				Name:        fmt.Sprintf("PrimeSdkTestAddressGroup-%d", time.Now().UnixMilli()),
				NetworkType: "NETWORK_TYPE_EVM",
				Addresses: []*model.OnchainAddress{&model.OnchainAddress{
					Name:     fmt.Sprintf("PrimeSdkTestAddress-%d", time.Now().UnixMilli()),
					Address:  "0x836fa72D2aF55d698e8767acBE88c042b8201036",
					ChainIds: []string{"8453"},
				},
				},
			},
		})
	if err != nil {
		t.Fatal(err)
	}

	if response == nil {
		t.Fatal("expected a not nil response")
	}

	fmt.Println(response.ActivityId)

	if len(response.ActivityId) == 0 {
		t.Fatal("expected an activity id")
	}

	if len(response.ActivityType) == 0 {
		t.Fatal("expected an activity type")
	}

	if response.RemainingApprovals == 0 {
		t.Fatal("expected consensus approvals > 0")
	}
}
