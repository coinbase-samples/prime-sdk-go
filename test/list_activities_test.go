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

package test

import (
	"context"
	"testing"
	"time"

	"github.com/coinbase-samples/prime-sdk-go/activities"
	"github.com/coinbase-samples/prime-sdk-go/model"
)

func TestListActivities(t *testing.T) {

	c, err := newLiveTestClient()
	if err != nil {
		t.Fatal(err)
	}

	service := activities.NewActivitiesService(c)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	response, err := service.ListActivities(ctx, &activities.ListActivitiesRequest{
		PortfolioId: c.Credentials().PortfolioId,
		Pagination:  &model.PaginationParams{Limit: "10"},
	})

	if err != nil {
		t.Fatal(err)
	}

	if response == nil {
		t.Fatal(err)
	}

	if len(response.Activities) == 0 {
		t.Fatal("expected activities in get")
	}

	if len(response.Activities[0].Id) == 0 {
		t.Fatal("expected activities id to be set")
	}

	for _, a := range response.Activities {

		if len(a.Id) == 0 {
			t.Error("expected id to be set")
		}

		if len(a.Category) == 0 {
			t.Error("expected category to be set")
		}

		if len(a.PrimaryType) == 0 {
			t.Error("expected primary type to be set")
		}

		if len(a.SecondaryType) == 0 {
			t.Error("expected secondary type to be set")
		}

		if len(a.Created) == 0 {
			t.Error("expected created to be set")
		}

		if len(a.Updated) == 0 {
			t.Error("expected updated to be set")
		}

		for _, u := range a.UserActions {
			if len(u.Action) == 0 {
				t.Error("expected user action to be set")
			}
			if len(u.Timestamp) == 0 {
				t.Error("expected timestamp to be set")
			}
		}

		testGetActivity(t, service, c.Credentials().PortfolioId, a.Id)
	}
}

func testGetActivity(t *testing.T, svc activities.ActivitiesService, portfolioId, activityId string) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	response, err := svc.GetActivity(ctx, &activities.GetActivityRequest{
		PortfolioId: portfolioId,
		Id:          activityId,
	})

	if err != nil {
		t.Fatal(err)
	}

	if response == nil {
		t.Fatal("expected activity response to not be nil")
	}

	if response.Activity == nil {
		t.Fatal("expected activity to not be nil")
	}

	if response.Activity.Id != activityId {
		t.Fatalf("expected activity id: %s - received activity id: %s", activityId, response.Activity.Id)
	}

}
