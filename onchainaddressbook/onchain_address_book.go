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

package onchainaddressbook

import (
	"context"

	"github.com/coinbase-samples/prime-sdk-go/client"
)

type OnchainAddressBookService interface {
	CreateOnchainAddressBookEntry(ctx context.Context, request *CreateOnchainAddressBookEntryRequest) (*CreateOnchainAddressBookEntryResponse, error)
	UpdateOnchainAddressBookEntry(ctx context.Context, request *UpdateOnchainAddressBookEntryRequest) (*UpdateOnchainAddressBookEntryResponse, error)
	DeleteOnchainAddressBookEntry(ctx context.Context, request *DeleteOnchainAddressBookEntryRequest) (*DeleteOnchainAddressBookEntryResponse, error)
	ListOnchainAddressBookGroups(ctx context.Context, request *ListOnchainAddressBookGroupsRequest) (*ListOnchainAddressBookGroupsResponse, error)
}

func NewOnchainAddressBookService(c client.RestClient) OnchainAddressBookService {
	return &onchainAddressBookServiceImpl{client: c}
}

type onchainAddressBookServiceImpl struct {
	client client.RestClient
}
