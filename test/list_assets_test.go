package test

import (
	"context"
	"testing"
	"time"

	prime "github.com/coinbase-samples/prime-sdk-go"
)

func TestListAssets(t *testing.T) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := newLiveTestClient()
	if err != nil {
		t.Fatal(err)
	}

	entityId, err := loadEntityId(client)
	if err != nil {
		t.Fatal(err)
	}

	response, err := client.ListAssets(ctx, &prime.ListAssetsRequest{
		EntityId: entityId,
	})

	if err != nil {
		t.Fatal(err)
	}

	if response == nil {
		t.Fatal(err)
	}

	if len(response.Assets) == 0 {
		t.Fatal("expected assets in get")
	}
}
