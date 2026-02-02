/**
 * Copyright 2024-present Coinbase Global, Inc.
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

package addressbook

import (
	"context"

	"github.com/coinbase-samples/prime-sdk-go/client"
	"github.com/coinbase-samples/prime-sdk-go/model"
)

type AddressBookService interface {
	GetAddressBook(ctx context.Context, request *GetAddressBookRequest) (*GetAddressBookResponse, error)
	CreateAddressBookEntry(ctx context.Context, request *CreateAddressBookEntryRequest) (*CreateAddressBookEntryResponse, error)
	ServiceConfig() *model.ServiceConfig
}

// NewAddressBookService creates a new AddressBookService with default pagination config
func NewAddressBookService(c client.RestClient) AddressBookService {
	return &addressBookServiceImpl{
		client:        c,
		serviceConfig: model.DefaultServiceConfig(),
	}
}

// NewAddressBookServiceWithConfig creates a new AddressBookService with custom pagination config
func NewAddressBookServiceWithConfig(c client.RestClient, config *model.ServiceConfig) AddressBookService {
	if config == nil {
		config = model.DefaultServiceConfig()
	}
	return &addressBookServiceImpl{
		client:        c,
		serviceConfig: config,
	}
}

type addressBookServiceImpl struct {
	client        client.RestClient
	serviceConfig *model.ServiceConfig
}

func (s *addressBookServiceImpl) ServiceConfig() *model.ServiceConfig {
	return s.serviceConfig
}
