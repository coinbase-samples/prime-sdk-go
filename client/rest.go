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

package client

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/coinbase-samples/core-go"
	"github.com/coinbase-samples/prime-sdk-go/credentials"
)

const defaultV1ApiBaseUrl = "https://api.prime.coinbase.com/v1"

var defaultHeadersFunc = AddPrimeHeaders

var DefaultSuccessHttpStatusCodes = []int{http.StatusOK}

type RestClient interface {
	SetBaseUrl(u string) RestClient
	HttpBaseUrl() string

	HttpClient() *http.Client
	SetHeadersFunc(hf core.HttpHeaderFunc) RestClient
	HeadersFunc() core.HttpHeaderFunc

	Credentials() *credentials.Credentials
}

func DefaultHttpClient() (http.Client, error) {
	return core.DefaultHttpClient()
}

type restClientImpl struct {
	httpClient http.Client
	baseUrl    string

	headersFunc core.HttpHeaderFunc
	credentials *credentials.Credentials
}

func (c *restClientImpl) HttpBaseUrl() string {
	return c.baseUrl
}

func (c *restClientImpl) SetBaseUrl(u string) RestClient {
	c.baseUrl = u
	return c
}

func (c *restClientImpl) HttpClient() *http.Client {
	return &c.httpClient
}

func (c *restClientImpl) Credentials() *credentials.Credentials {
	return c.credentials
}

func (c *restClientImpl) SetHeadersFunc(hf core.HttpHeaderFunc) RestClient {
	c.headersFunc = hf
	return c
}

func (c *restClientImpl) HeadersFunc() core.HttpHeaderFunc {
	return c.headersFunc
}

func NewRestClient(credentials *credentials.Credentials, httpClient http.Client) RestClient {
	return &restClientImpl{
		baseUrl:     defaultV1ApiBaseUrl,
		credentials: credentials,
		httpClient:  httpClient,
		headersFunc: defaultHeadersFunc,
	}
}

func AddPrimeHeaders(req *http.Request, path string, body []byte, cl core.RestClient, t time.Time) {
	c := cl.(*restClientImpl)
	timestamp := strconv.FormatInt(t.Unix(), 10)
	signature := sign(req.Method, path, timestamp, c.Credentials().SigningKey, string(body))
	req.Header.Add("Accept", "application/json")
	req.Header.Add("X-CB-ACCESS-KEY", c.Credentials().AccessKey)
	req.Header.Add("X-CB-ACCESS-PASSPHRASE", c.Credentials().Passphrase)
	req.Header.Add("X-CB-ACCESS-SIGNATURE", signature)
	req.Header.Add("X-CB-ACCESS-TIMESTAMP", timestamp)
	req.Header.Set("User-Agent", fmt.Sprintf("prime-sdk-go/%s", sdkVersion))
}

func sign(method, path, timestamp, signingKey, body string) string {
	h := hmac.New(sha256.New, []byte(signingKey))
	h.Write([]byte(fmt.Sprintf("%s%s%s%s", timestamp, method, path, body)))
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}
