/**
 * Copyright 2025-present Coinbase Global, Inc.
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

package test

import (
	"context"
	"testing"
	"time"

	"github.com/coinbase-samples/prime-sdk-go/model"
	"github.com/coinbase-samples/prime-sdk-go/products"
)

func TestListProducts(t *testing.T) {

	c, err := newLiveTestClient()
	if err != nil {
		t.Fatal(err)
	}

	svc := products.NewProductsService(c)

	var cursor string
	var products []*model.Product

	for {

		products, cursor, err = listProducts(svc, c.Credentials().PortfolioId, cursor)

		if err != nil {
			t.Fatal(err)
		}

		if len(products) == 0 {
			t.Fatal("expected products in get")
		}

		for _, p := range products {

			if p == nil {
				t.Fatal("received a nil product ")
			}
			//fmt.Println(p.Id)
			//fmt.Println(fmt.Sprintf("%s-%t-%s-%s-%s-%s-%s-%s", p.Id, p.RfqProductDetails.Tradable, p.RfqProductDetails.MinNotionalSize, p.RfqProductDetails.MaxNotionalSize, p.RfqProductDetails.MinBaseSize, p.RfqProductDetails.MaxBaseSize, p.RfqProductDetails.MinQuoteSize, p.RfqProductDetails.MaxQuoteSize))

		}

		if len(cursor) == 0 {
			break
		}

	}
}

func listProducts(svc products.ProductsService, portfolioId, cursor string) ([]*model.Product, string, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var next string

	response, err := svc.ListProducts(ctx, &products.ListProductsRequest{
		PortfolioId: portfolioId,
		Pagination: &model.PaginationParams{
			Cursor: cursor,
		},
	})

	if err != nil {
		return nil, next, err
	}

	if response.Pagination != nil {
		next = response.Pagination.NextCursor
	}

	return response.Products, next, err
}
