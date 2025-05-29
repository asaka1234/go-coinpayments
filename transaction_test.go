package coinpayments

import (
	"fmt"
	"testing"
)

func TestCallGetTxInfo(t *testing.T) {
	client, err := testClient()
	if err != nil {
		t.Fatal("should have instantiated a new client with valid configuration")
	}

	resp, err := client.CallGetTxInfo(&TxInfoRequest{TxID: "CPCH3A2HBY7R92IOCQNYPL1G8X", Full: "0"})
	if err != nil {
		t.Fatal("error getting tx info ", err)
	}

	fmt.Printf("%+v\n", resp)
}
func TestCallCreateTransaction(t *testing.T) {
	client, err := testClient()
	if err != nil {
		t.Fatalf("Should have instantiated a new client with valid config and http client, but it threw error: %s", err.Error())
	}

	resp, err := client.CallCreateTransaction(&TransactionRequest{Amount: "100", Currency1: "USD", Currency2: "BTC", BuyerEmail: "jeff@internet.com"})
	if err != nil {
		t.Fatalf("Could not call create transaction: %s", err.Error())
	}

	_, err = client.CallGetTxInfo(&TxInfoRequest{TxID: resp.TxnID})

	if err != nil {
		t.Fatalf("Could not call get tx info: %s", err.Error())
	}

}

func TestCallGetCallbackAddress(t *testing.T) {
	client, err := testClient()
	if err != nil {
		t.Fatalf("Should have instantiated a new client with valid config and http client, but it threw error: %s", err.Error())
	}

	_, err = client.CallGetCallbackAddress(&CallbackAddressRequest{Currency: "BTC"})

	if err != nil {
		t.Fatalf("Could not call get callback address: %s", err.Error())
	}

}

func TestCallGetDepositAddress(t *testing.T) {
	client, err := testClient()
	if err != nil {
		t.Fatalf("Should have instantiated a new client with valid config and http client, but it threw error: %s", err.Error())
	}

	_, err = client.CallGetDepositAddress(&DepositAddressRequest{Currency: "BTC"})

	if err != nil {
		t.Fatalf("Could not call get callback address: %s", err.Error())
	}

}

func TestCallGetTxList(t *testing.T) {
	client, err := testClient()
	if err != nil {
		t.Fatalf("Should have instantiated a new client with valid config and http client, but it threw error: %s", err.Error())
	}

	_, err = client.CallGetTxList(&TxListRequest{})

	if err != nil {
		t.Fatalf("Could not call get tx list: %s", err.Error())
	}

}
