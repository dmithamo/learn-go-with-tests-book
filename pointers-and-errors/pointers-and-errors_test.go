package pointersanderrors

import (
	"testing"
)

func assertError(t *testing.T, gotErr, wantErr error) {
	t.Helper()
	if gotErr != wantErr {
		t.Errorf("got error %q want error %q", gotErr, wantErr)
	}
}

func TestWallet(t *testing.T) {
	assertBalance := func(t *testing.T, w Wallet, want string) {
		t.Helper()
		got := w.Balance()
		if got != want {
			t.Errorf("got balance %q want %q", got, want)
		}
	}

	var w Wallet

	resetWallet := func() {
		w.balance = 0.00
		w.currency = "KES"
	}

	t.Run("Balance: correct balance", func(t *testing.T) {
		resetWallet()
		assertBalance(t, w, "KES 0.00")
	})

	t.Run("Deposit", func(t *testing.T) {
		t.Run("updates balance ok after deposit", func(t *testing.T) {
			resetWallet()
			err := w.Deposit("KES 10.00")
			assertBalance(t, w, "KES 10.00")
			assertError(t, err, nil)
		})

		t.Run("handles unsupported currency error", func(t *testing.T) {
			resetWallet()
			err := w.Deposit("USD 75.00")
			assertError(t, err, &UNSUPPORTED_CURRENCY_ERR)
			assertBalance(t, w, "KES 0.00")
		})
	})

	t.Run("Withdraw", func(t *testing.T) {
		t.Run("throws insufficient funds error if balance < withdrawal amount", func(t *testing.T) {
			resetWallet()
			err := w.Withdraw("KES 10.00")
			assertError(t, err, &INSUFFICIENT_FUNDS_ERR)
			assertBalance(t, w, "KES 0.00")
		})

		t.Run("handles unsupported currency error", func(t *testing.T) {
			resetWallet()
			err := w.Deposit("KES 75.00")
			assertError(t, err, nil)

			err = w.Withdraw("USD 50.00")

			assertError(t, err, &UNSUPPORTED_CURRENCY_ERR)
			assertBalance(t, w, "KES 75.00")
		})

		t.Run("updates balance ok after withdrawal", func(t *testing.T) {
			resetWallet()
			err := w.Deposit("KES 20.00")
			assertError(t, err, nil)
			err = w.Withdraw("KES 5.00")
			assertBalance(t, w, "KES 15.00")
			assertError(t, err, nil)
		})
	})
}

func TestGetCurrencyHelper(t *testing.T) {
	t.Run("extracts currency and value", func(t *testing.T) {
		value, curr, _ := GetCurrencyHelper("KES 20.00")
		if value != 20.00 {
			t.Errorf("got value %.2f want value 20.00", value)
		}

		if curr != "KES" {
			t.Errorf("got currency %q want currency 'KES'", curr)
		}
	})

	t.Run("throws appropriate err if unable to extract value and curr", func(t *testing.T) {
		// stupid test, I think. Which caller would do this?
		_, _, err := GetCurrencyHelper("KES Dennis")
		assertError(t, err, &UNHANDLED_ERR)
	})
}
