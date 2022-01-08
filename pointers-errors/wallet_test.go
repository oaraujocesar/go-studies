package wallet

import (
	"testing"
)

func TestWallet(t *testing.T) {

	assertBalance := func(t testing.TB, wallet Wallet, expected Bitcoin) {
		t.Helper()

		received := wallet.Balance()

		if received != expected {
			t.Errorf("Received %v, but expected %v", received, expected)
		}
	}

	assertError := func(t *testing.T, received error, expected string) {
		t.Helper()

		if received == nil {
			t.Fatal("Expected an error, but didn't get one")
		}

		if received.Error() != expected {
			t.Errorf("Received %q, but expected %q", received, expected)
		}
	}

	assertInexistentError := func(t *testing.T, err error) {
		t.Helper()

		if err != nil {
			t.Errorf("unexpected error")
		}
	}

	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10.3))

		assertBalance(t, wallet, Bitcoin(10.3))
		// fmt.Printf("address of balance in test is %v \n", &wallet.balance)
	})

	t.Run("Withdraw with enough balance", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20.0)}
		err := wallet.Withdraw(Bitcoin(10.0))

		assertBalance(t, wallet, Bitcoin(10.0))
		assertInexistentError(t, err)
	})

	t.Run("Withdraw without enough balance", func(t *testing.T) {
		initialBalance := Bitcoin(20.0)
		wallet := Wallet{balance: initialBalance}
		err := wallet.Withdraw(Bitcoin(100.0))

		assertBalance(t, wallet, initialBalance)
		assertError(t, err, ErrInsufficientFunds.Error())
	})
}
