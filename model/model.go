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

package model

// ErrorMessage represents a generic error response
type ErrorMessage struct {
	Value string `json:"message"`
}

// PaginationParams represents pagination parameters for list requests
type PaginationParams struct {
	Cursor        string `json:"cursor"`
	Limit         int32  `json:"limit"`
	SortDirection string `json:"sort_direction"`
}

// Pagination represents pagination information in responses
type Pagination struct {
	NextCursor    string `json:"next_cursor"`
	SortDirection string `json:"sort_direction"`
	HasNext       bool   `json:"has_next"`
}
