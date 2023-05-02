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

func appendQueryParam(queryParams, key, value string) string {
	return fmt.Sprintf("%s%s%s=%s", queryParams, queryParamSep(strings.Contains(queryParams, "?")), key, value)
}

func appendPaginationParams(v string, p *PaginationParams) string {

	if p == nil {
		return v
	}

	if len(p.Cursor) > 0 {
		v = appendQueryParam(v, "cursor", p.Cursor)
	}

	if len(p.Limit) > 0 {
		v = appendQueryParam(v, "limit", p.Limit)
	}

	if len(p.SortDirection) > 0 {
		v = appendQueryParam(v, "sort_direction", p.SortDirection)
	}

	return v
}

func queryParamSep(appended bool) string {
	if appended {
		return "&"
	}
	return "?"
}
