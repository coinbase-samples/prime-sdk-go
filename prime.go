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
	"net/url"
	"os"

	"log"
)

var primeV1ApiBaseUrl = "https://api.prime.coinbase.com/v1"

func init() {
	baseUrl := os.Getenv("PRIME_SDK_BASE_URL")
	if len(baseUrl) > 0 {
		_, err := url.Parse(baseUrl)
		if err != nil {
			log.Fatalf(
				"cannot parse PRIME_SDK_BASE_URL - received: %s - err: %v",
				baseUrl,
				err,
			)
		}
		primeV1ApiBaseUrl = baseUrl
	}
}
