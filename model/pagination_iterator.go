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

package model

import "context"

// ServiceConfig controls pagination behavior for services
type ServiceConfig struct {
	// MaxPages is the maximum number of pages to fetch (0 = unlimited)
	MaxPages int
	// MaxItems is the maximum number of items to fetch (0 = unlimited)
	MaxItems int
	// DefaultLimit is the default page size if not specified in the request
	DefaultLimit int32
}

// DefaultServiceConfig returns a config with no limits
func DefaultServiceConfig() *ServiceConfig {
	return &ServiceConfig{
		MaxPages:     0,
		MaxItems:     0,
		DefaultLimit: 25,
	}
}

// PaginationMixin provides HasNext and GetNextCursor functionality.
// Embed this in response structs to avoid duplicating these methods.
type PaginationMixin struct {
	Pagination *Pagination `json:"pagination"`
}

// HasNext returns true if there are more pages available
func (m *PaginationMixin) HasNext() bool {
	return m.Pagination != nil && m.Pagination.HasNext
}

// GetNextCursor returns the cursor for the next page, or empty string if none
func (m *PaginationMixin) GetNextCursor() string {
	if m.Pagination == nil {
		return ""
	}
	return m.Pagination.NextCursor
}

// PrepareNextPagination creates pagination params for the next page request.
// It safely copies existing params (if any) and sets the cursor for the next page.
func PrepareNextPagination(current *PaginationParams, nextCursor string) *PaginationParams {
	if current == nil {
		return &PaginationParams{Cursor: nextCursor}
	}
	cp := *current
	cp.Cursor = nextCursor
	return &cp
}

// PaginatedResponse is implemented by any response that supports pagination
type PaginatedResponse[T any] interface {
	HasNext() bool
	GetNextCursor() string
	Next(ctx context.Context) (T, error)
}

// ItemExtractor extracts a slice of items from a response
type ItemExtractor[R any, I any] func(R) []I

// PageIterator provides iteration over paginated responses
type PageIterator[R PaginatedResponse[R], I any] struct {
	current   R
	extractor ItemExtractor[R, I]
	config    *ServiceConfig
}

// NewPageIterator creates an iterator from an initial response
func NewPageIterator[R PaginatedResponse[R], I any](
	initial R,
	extractor ItemExtractor[R, I],
) *PageIterator[R, I] {
	return &PageIterator[R, I]{
		current:   initial,
		extractor: extractor,
		config:    nil,
	}
}

// NewPageIteratorWithConfig creates an iterator with pagination config
func NewPageIteratorWithConfig[R PaginatedResponse[R], I any](
	initial R,
	extractor ItemExtractor[R, I],
	config *ServiceConfig,
) *PageIterator[R, I] {
	return &PageIterator[R, I]{
		current:   initial,
		extractor: extractor,
		config:    config,
	}
}

// WithConfig sets the pagination config and returns the iterator for chaining
func (it *PageIterator[R, I]) WithConfig(config *ServiceConfig) *PageIterator[R, I] {
	it.config = config
	return it
}

// Current returns the current page response
func (it *PageIterator[R, I]) Current() R {
	return it.current
}

// Items returns items from the current page
func (it *PageIterator[R, I]) Items() []I {
	return it.extractor(it.current)
}

// HasNext returns true if there are more pages
func (it *PageIterator[R, I]) HasNext() bool {
	return it.current.HasNext()
}

// Next advances to the next page and returns the new response
func (it *PageIterator[R, I]) Next(ctx context.Context) (R, error) {
	next, err := it.current.Next(ctx)
	if err != nil {
		return it.current, err
	}
	it.current = next
	return next, nil
}

// FetchAll retrieves all items across all pages starting from current page.
// Respects MaxPages and MaxItems from config if set.
func (it *PageIterator[R, I]) FetchAll(ctx context.Context) ([]I, error) {
	all := make([]I, 0)
	all = append(all, it.Items()...)

	pages := 1
	for it.HasNext() {
		// Check MaxPages limit
		if it.config != nil && it.config.MaxPages > 0 && pages >= it.config.MaxPages {
			break
		}
		// Check MaxItems limit
		if it.config != nil && it.config.MaxItems > 0 && len(all) >= it.config.MaxItems {
			break
		}

		_, err := it.Next(ctx)
		if err != nil {
			return all, err
		}
		all = append(all, it.Items()...)
		pages++
	}

	// Trim to MaxItems if exceeded
	if it.config != nil && it.config.MaxItems > 0 && len(all) > it.config.MaxItems {
		all = all[:it.config.MaxItems]
	}

	return all, nil
}

// ForEach iterates through all pages starting from current, calling fn for each page.
// Respects MaxPages from config if set.
func (it *PageIterator[R, I]) ForEach(ctx context.Context, fn func(R) error) error {
	if err := fn(it.current); err != nil {
		return err
	}

	pages := 1
	for it.HasNext() {
		// Check MaxPages limit
		if it.config != nil && it.config.MaxPages > 0 && pages >= it.config.MaxPages {
			break
		}

		_, err := it.Next(ctx)
		if err != nil {
			return err
		}
		if err := fn(it.current); err != nil {
			return err
		}
		pages++
	}

	return nil
}
