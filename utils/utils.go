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

package utils

import (
	"fmt"
	"time"

	"github.com/coinbase-samples/core-go"
	"github.com/coinbase-samples/prime-sdk-go/model"
)

func TimeToStr(t time.Time) string {
	return t.Format("2006-01-02T15:04:05Z")
}

func NewUuid() string {
	return core.NewUuid()
}

func AppendPaginationParams(v string, p *model.PaginationParams) string {

	if p == nil {
		return v
	}

	if len(p.Cursor) > 0 {
		v = core.AppendHttpQueryParam(v, "cursor", p.Cursor)
	}

	if p.Limit > 0 {
		v = core.AppendHttpQueryParam(v, "limit", fmt.Sprintf("%d", p.Limit))
	}

	if len(p.SortDirection) > 0 {
		v = core.AppendHttpQueryParam(v, "sort_direction", p.SortDirection)
	}

	return v
}
