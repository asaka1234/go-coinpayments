package coinpayments

import (
	"testing"
)

func TestCallBalances(t *testing.T) {
	client, err := testClient()
	if err != nil {
		t.Fatalf("Should have instantiated a new client with valid config and http client, but it threw error: %s", err.Error())
	}

	_, err = client.CallBalances(&BalancesRequest{All: "1"})
	if err != nil {
		t.Fatalf("Could not call balances %s", err.Error())
	}

}
