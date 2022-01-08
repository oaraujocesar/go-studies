package wallet

import (
	"fmt"
	"testing"
)

func TestWallet(t *testing.T) {
	wallet := Wallet{}

	wallet.Deposit(10.0)

	received := wallet.Balance()
	expected := 10.0

	fmt.Printf("address of balance in test is %v \n", &wallet.balance)

	if received != expected {
		t.Errorf("Received %.2f but expected %.2f", received, expected)
	}
}
