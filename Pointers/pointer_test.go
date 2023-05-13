package pointers

import "testing"

func TestWallet(t *testing.T) {

	check := func(t testing.TB, w Wallet, expected Bitcoin) {
		t.Helper()
		got := w.get_balance()
		if got != expected {
			t.Errorf("got %s and expected %s", got, expected)
		}
	}
	t.Run("testing the deposit function", func(t *testing.T) {
		wallet := Wallet{}
		wallet.deposit(Bitcoin(10))
		expected := Bitcoin(10)
		check(t, wallet, expected)
	})

	t.Run("testing the withdraw function", func(t *testing.T) {
		wallet := Wallet{balance: 50}
		wallet.withdraw(Bitcoin(10))
		expected := Bitcoin(40)
		check(t, wallet, expected)

	})

}
