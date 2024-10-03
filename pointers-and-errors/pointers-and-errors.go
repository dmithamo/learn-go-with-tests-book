package pointersanderrors

import (
	"fmt"
	"strconv"
	"strings"
)

type Wallet struct {
	balance  float64
	currency string
}

type InsufficientFundsError string

func (e InsufficientFundsError) Error() string {
	return "insufficient funds"
}

type UnsupportedCurrencyError string

func (e UnsupportedCurrencyError) Error() string {
	return "unsupported currency"
}

type UnhandledError string

func (e UnhandledError) Error() string {
	return "Something went wrong"
}

var INSUFFICIENT_FUNDS_ERR InsufficientFundsError
var UNSUPPORTED_CURRENCY_ERR UnsupportedCurrencyError
var UNHANDLED_ERR UnhandledError

func (w *Wallet) Balance() string {
	return fmt.Sprintf("%s %.2f", w.currency, w.balance)
}

// Deposit credit a wallet with the specified amount
func (w *Wallet) Deposit(amount string) error {
	amt, curr, _ := GetCurrencyHelper(amount)

	if curr != w.currency {
		return &UNSUPPORTED_CURRENCY_ERR
	}

	w.balance += amt
	return nil
}

// WIthdraw debits the wallet with the specified amount
//
//	given the balance >= amount
func (w *Wallet) Withdraw(amount string) error {
	value, currency, _ := GetCurrencyHelper(amount)

	if currency != w.currency {
		return &UNSUPPORTED_CURRENCY_ERR
	}

	if w.balance < value {
		return &INSUFFICIENT_FUNDS_ERR
	}

	w.balance -= value
	return nil
}

func GetCurrencyHelper(value string) (float64, string, error) {
	amountWithCurrency := strings.Fields(value)
	amount, err := strconv.ParseFloat(amountWithCurrency[1], 64)

	if err != nil {
		return 0.00, "", &UNHANDLED_ERR
	}

	return amount, amountWithCurrency[0], nil
}
