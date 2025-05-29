package coinpayments

import (
	"errors"
	"net/http"
	"net/url"
	"os"
	"testing"
)

func testClient() (*Client, error) {
	pubKey, ok := os.LookupEnv("COINPAYMENTS_PUBLIC_KEY")
	if !ok {
		return nil, errors.New("no public key provided in environment")
	}

	privateKey, ok := os.LookupEnv("COINPAYMENTS_PRIVATE_KEY")
	if !ok {
		return nil, errors.New("no priatekey provided in environment")
	}

	return NewClient(&Config{PublicKey: pubKey, PrivateKey: privateKey}, &http.Client{})

}
func TestNewClient(t *testing.T) {
	if _, err := NewClient(&Config{PublicKey: "", PrivateKey: ""}, &http.Client{}); err == nil {
		t.Fatalf("Should have thrown an error with emptu public and private key, but it didn't")
	}

	if _, err := NewClient(&Config{PublicKey: "publickey", PrivateKey: "privateKey"}, nil); err == nil {
		t.Fatalf("Should have thrown an error with a nil http client, but it didn't")
	}

	if _, err := testClient(); err != nil {
		t.Fatalf("Should have instantiated a new client with valid config and http client, but it threw error: %s", err.Error())
	}
}

func TestCall(t *testing.T) {
	client, err := testClient()
	if err != nil {
		t.Fatalf("Should have instantiated a new client with valid config and http client, but it threw error: %s", err.Error())
	}

	failCommand := "doesntexist"
	var txResponse TransactionResponse
	if err := client.Call(failCommand, url.Values{}, &txResponse); err == nil {
		t.Fatalf("Should have failed with non-supported command of %s", failCommand)
	}

}
