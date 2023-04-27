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
	"strings"
)

func sign(path, body, method, signingKey string, t int64) string {
	h := hmac.New(sha256.New, []byte(signingKey))
	h.Write([]byte(fmt.Sprintf("%d%s%s%s", t, method, path, body)))
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

func urlIteratorParams(url string, p *IteratorParams) string {

	appended := strings.Contains(url, "?")

	if len(p.Cursor) > 0 {
		url += fmt.Sprintf("%scursor=%s", urlParamSep(appended), p.Cursor)
		appended = true
	}

	if len(p.Limit) > 0 {
		url += fmt.Sprintf("%slimit=%s", urlParamSep(appended), p.Limit)
		appended = true
	}

	if len(p.SortDirection) > 0 {
		url += fmt.Sprintf("%sort_direction=%s", urlParamSep(appended), p.SortDirection)
		appended = true
	}

	return url
}

func urlParamSep(appended bool) string {
	if appended {
		return "&"
	}
	return "?"
}
