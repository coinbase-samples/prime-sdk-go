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
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/coinbase-samples/core-go"
	"github.com/coinbase-samples/prime-sdk-go/client"
	"github.com/coinbase-samples/prime-sdk-go/credentials"
	"github.com/coinbase-samples/prime-sdk-go/portfolios"
)

func loadEntityId(c client.RestClient) (string, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	service := portfolios.NewPortfoliosService(c)

	response, err := service.GetPortfolio(
		ctx,
		&portfolios.GetPortfolioRequest{
			PortfolioId: c.Credentials().PortfolioId,
		},
	)

	if err != nil {
		return "", err
	}

	return response.Portfolio.EntityId, nil
}

func newLiveTestClient() (client.RestClient, error) {

	credentials, err := loadCredentialsFromEnv()
	if err != nil {
		return nil, err
	}

	httpClient, err := core.DefaultHttpClient()
	if err != nil {
		return nil, err
	}

	client := client.NewRestClient(credentials, httpClient)
	return client, nil

}

func loadCredentialsFromEnv() (*credentials.Credentials, error) {

	credentials := &credentials.Credentials{}
	if err := json.Unmarshal([]byte(os.Getenv("PRIME_CREDENTIALS")), credentials); err != nil {
		return nil, fmt.Errorf("unable to deserialize prime credentials JSON: %w", err)
	}

	return credentials, nil
}
