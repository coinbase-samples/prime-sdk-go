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
	"testing"

	"github.com/coinbase-samples/prime-sdk-go/model"
)

func TestPaginationParams(t *testing.T) {

	cases := []struct {
		description string
		params      *model.PaginationParams
		query       string
		expected    string
	}{
		{
			description: "TestPaginationParams0",
			params:      &model.PaginationParams{Cursor: "test"},
			query:       "",
			expected:    "?cursor=test",
		},
		{
			description: "TestPaginationParams1",
			params:      &model.PaginationParams{Cursor: "test"},
			query:       "?test=new",
			expected:    "?test=new&cursor=test",
		},
		{
			description: "TestPaginationParams1",
			params:      &model.PaginationParams{Cursor: "test", Limit: "10", SortDirection: "ASC"},
			query:       "?test=new",
			expected:    "?test=new&cursor=test&limit=10&sort_direction=ASC",
		},
	}

	for _, tt := range cases {
		t.Run(tt.description, func(t *testing.T) {
			result := AppendPaginationParams(tt.query, tt.params)
			if result != tt.expected {
				t.Errorf("test: %s - expected: %s - received: %s", tt.description, tt.expected, result)
			}
		})
	}
}
