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
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"github.com/coinbase-samples/core-go"
	"net/http"
	"strconv"
	"time"
)

var defaultV1ApiBaseUrl = "https://api.prime.coinbase.com/v1"

type Client struct {
	httpClient  http.Client
	httpBaseUrl string
	Credentials *Credentials
}

func (c *Client) HttpBaseUrl() string {
	return c.httpBaseUrl
}

func (c *Client) HttpClient() *http.Client {
	return &c.httpClient
}

func (c *Client) SetBaseUrl(u string) *Client {
	c.httpBaseUrl = u
	return c
}

func NewClient(credentials *Credentials, httpClient http.Client) *Client {
	return &Client{
		httpBaseUrl: defaultV1ApiBaseUrl,
		Credentials: credentials,
		httpClient:  httpClient,
	}
}

func addPrimeHeaders(req *http.Request, path string, body []byte, client core.Client, t time.Time) {
	c := client.(*Client)
	timestamp := strconv.FormatInt(t.Unix(), 10)
	signature := sign(req.Method, path, timestamp, c.Credentials.SigningKey, string(body))
	req.Header.Add("Accept", "application/json")
	req.Header.Add("X-CB-ACCESS-KEY", c.Credentials.AccessKey)
	req.Header.Add("X-CB-ACCESS-PASSPHRASE", c.Credentials.Passphrase)
	req.Header.Add("X-CB-ACCESS-SIGNATURE", signature)
	req.Header.Add("X-CB-ACCESS-TIMESTAMP", timestamp)
}

func sign(method, path, timestamp, signingKey, body string) string {
	h := hmac.New(sha256.New, []byte(signingKey))
	h.Write([]byte(fmt.Sprintf("%s%s%s%s", timestamp, method, path, body)))
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}
