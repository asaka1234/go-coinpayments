package coinpayments

import (
	"testing"
)

func TestCallGetBasicInfo(t *testing.T) {
	client, err := testClient()
	if err != nil {
		t.Fatalf("Should have instantiated a new client with valid config and http client, but it threw error: %s", err.Error())
	}

	_, err = client.CallGetBasicInfo()
	if err != nil {
		t.Fatalf("Could not call get basic info: %s", err.Error())
	}

}

func TestCallRates(t *testing.T) {
	client, err := testClient()
	if err != nil {
		t.Fatalf("Should have instantiated a new client with valid config and http client, but it threw error: %s", err.Error())
	}

	_, err = client.CallRates(&RatesRequest{Short: "1", Accepted: "0"})

	if err != nil {
		t.Fatalf("Could not call get basic info: %s", err.Error())
	}

}
