package test

/*
import (
	"context"
	"testing"
	"time"

	prime "github.com/coinbase-samples/prime-sdk-go"
)

func TestGetInvoices(t *testing.T) {

	client, err := newLiveTestClient()
	if err != nil {
		t.Fatal(err)
	}

	entityId, err := loadEntityId(client)
	if err != nil {
		t.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	response, err := client.GetInvoices(ctx, &prime.GetInvoicesRequest{
		EntityId: entityId,
	})

	if err != nil {
		t.Fatal(err)
	}

	if response == nil {
		t.Fatal(err)
	}

	if len(response.Invoices) == 0 {
		t.Fatal("expected invoices in describe")
	}
}
*/