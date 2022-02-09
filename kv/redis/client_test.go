//go:build integration

package redis

import (
	"os"
	"reflect"
	"testing"
)

var testClient *Client

func TestNewClient(t *testing.T) {
	client, err := testCreateClient()
	if err != nil {
		t.Fatalf("unexpected error, got: %s, want: nil.", err.Error())
	}
	if reflect.TypeOf(client) != reflect.TypeOf(&Client{}) {
		t.Fatalf("unexpected client type, got: %s, want: %s", reflect.TypeOf(client), reflect.TypeOf(&Client{}))
	}
	if client == nil {
		t.Fatalf("expected client, got: nil")
	}
}

func testCreateClient() (*Client, error) {
	if testClient != nil {
		return testClient, nil
	}

	newClient, err := NewClient(os.Getenv("TEST_REDIS"), 0, os.Getenv("TEST_REDIS_PASS"))
	if err != nil {
		return nil, err
	}
	testClient = newClient

	return testClient, nil
}
