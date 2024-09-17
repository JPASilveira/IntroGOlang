package main

import (
	"testing"
)

func TestAccount(t *testing.T) {
	verify := func(t *testing.T, result, expected Real) {
		if result != expected {
			t.Errorf("result: %s, expected: %s", result, expected)
		}
	}

	confirmMistake := func(t *testing.T, mistake error) {
		t.Helper()
		if mistake != nil {
			t.Fatal(mistake)
		}
	}

	t.Run("deposit", func(t *testing.T) {
		t.Helper()
		account := Account{}
		mistake := account.Deposit(10.0)
		verify(t, account.Balance(), Real(10.0))
		confirmMistake(t, mistake)
	})

	t.Run("withdraw", func(t *testing.T) {
		t.Helper()
		account := Account{10}
		mistake := account.Withdraw(5)
		verify(t, account.Balance(), Real(5))
		confirmMistake(t, mistake)
	})
}
