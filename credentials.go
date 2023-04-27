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
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

var DefaultCredentials *Credentials

type Credentials struct {
	AccessKey    string `json:"accessKey"`
	Passphrase   string `json:"passphrase"`
	SigningKey   string `json:"signingKey"`
	PortfolioId  string `json:"portfolioId"`
	SvcAccountId string `json:"svcAccountId"`
}

func ReadEnvCredentials() (*Credentials, error) {

	v := os.Getenv("PRIME_CREDENTIALS")

	if len(v) == 0 {
		return nil, errors.New("PRIME_CREDENTIALS not set as environment variable")
	}

	credentials := &Credentials{}
	if err := json.Unmarshal([]byte(v), &credentials); err != nil {
		return nil, fmt.Errorf("unable unarmshal PRIME_CREDENTIALS to JSON: %w", err)
	}

	return credentials, nil
}
