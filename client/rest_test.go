/**
 * Copyright 2026-present Coinbase Global, Inc.
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

import "testing"

func TestVersionedBaseUrl(t *testing.T) {
	cases := []struct {
		base    string
		version string
		want    string
	}{
		{"https://api.prime.coinbase.com/v1", "v2", "https://api.prime.coinbase.com/v2"},
		{"https://api.prime.coinbase.com/v1/", "v2", "https://api.prime.coinbase.com/v2"},
		{"https://api.prime.coinbase.com", "v2", "https://api.prime.coinbase.com/v2"},
		{"https://api.prime.coinbase.com/", "v2", "https://api.prime.coinbase.com/v2"},
		{"https://api.prime.coinbase.com/v1", "v1", "https://api.prime.coinbase.com/v1"},
	}

	for _, tc := range cases {
		got := VersionedBaseUrl(tc.base, tc.version)
		if got != tc.want {
			t.Errorf("VersionedBaseUrl(%q, %q) = %q; want %q", tc.base, tc.version, got, tc.want)
		}
	}
}
