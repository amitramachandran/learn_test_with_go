package pointers

import "testing"

func TestWallet(t *testing.T) {

	assertNoError := func(t testing.TB, err error) {
		t.Helper()
		if err != nil {
			t.Fatal("Should not get an error but an error has occured")
		}
	}

	assertError := func(t testing.TB, got, want error) {
		t.Helper()
		if got == nil {
			t.Fatal("didnt got error but need one")
		}

		if got != want {
			t.Errorf("got %q, but wanted %q", got, want)
		}

	}

	assertBalance := func(t testing.TB, w Wallet, expected Bitcoin) {
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
		assertBalance(t, wallet, expected)
	})

	t.Run("testing the withdraw function", func(t *testing.T) {
		wallet := Wallet{balance: 50}
		err := wallet.withdraw(Bitcoin(10))
		expected := Bitcoin(40)
		assertNoError(t, err)
		assertBalance(t, wallet, expected)

	})

	t.Run("insufficient funds while withdrawing", func(t *testing.T) {
		wallet := Wallet{balance: 10}
		err := wallet.withdraw(Bitcoin(20))
		assertError(t, err, ErrInsufficientFunds)
		assertBalance(t, wallet, Bitcoin(10))

	})

}
