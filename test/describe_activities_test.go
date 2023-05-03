package test

import (
	"context"
	"testing"
	"time"

	prime "github.com/coinbase-samples/prime-sdk-go"
)

func TestDescribeActivities(t *testing.T) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := newLiveTestClient()
	if err != nil {
		t.Fatal(err)
	}

	response, err := client.DescribeActivities(ctx, &prime.DescribeActivitiesRequest{
		PortfolioId: client.Credentials.PortfolioId,
		Pagination:  &prime.PaginationParams{Limit: "10"},
	})

	if err != nil {
		t.Fatal(err)
	}

	if response == nil {
		t.Fatal(err)
	}

	if len(response.Activities) == 0 {
		t.Fatal("expected activities in describe")
	}

	if len(response.Activities[0].Id) == 0 {
		t.Fatal("expected activities id to be set")
	}

	for _, a := range response.Activities {
		//fmt.Println(fmt.Sprintf("id: %s - category: %s - primary type: %s - secondary type: %s - symbols: %v - created by: %s - at: %s - update: %s", a.Id, a.Category, a.PrimaryType, a.SecondaryType, a.Symbols, a.CreatedBy, a.Created, a.Updated))

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

		if len(a.CreatedBy) == 0 {
			t.Error("expected created by to be set")
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

			if len(u.UserId) == 0 {
				t.Error("expected user id to be set")
			}

			if len(u.Timestamp) == 0 {
				t.Error("expected timestamp to be set")
			}

			//fmt.Println(fmt.Sprintf("    user action: %s - user id: %s - timestamp: %s", u.Action, u.UserId, u.Timestamp))
		}

		if a.AccountMetadata != nil && a.AccountMetadata.Consensus != nil {
			if len(a.AccountMetadata.Consensus.ApprovalDeadline) == 0 {
				t.Error("expected approval deadline to be set")
			}
		}

		testDescribeActivity(t, client, client.Credentials.PortfolioId, a.Id)
	}
}

func testDescribeActivity(t *testing.T, client *prime.Client, portfolioId, activityId string) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	response, err := client.DescribeActivity(ctx, &prime.DescribeActivityRequest{
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
