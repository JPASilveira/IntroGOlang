package main

import (
	"errors"
	"fmt"
)

type Real float64

type AccountOperations interface {
	Deposit(amount Real)
	Withdraw(amount Real)
}

type Stringer interface {
	String() string
}

type Account struct {
	amount Real
}

func (r Real) String() string {
	return fmt.Sprintf("R$ %.2f", r)
}

func (a *Account) Deposit(value Real) error {
	if value <= 0 {
		return errors.New("invalid value")
	}
	a.amount += value
	return nil
}

func (a *Account) Withdraw(value Real) error {
	if value == 0 {
		return errors.New("cannot withdraw from zero amount")
	}
	if a.amount < value {
		return errors.New("insufficient funds")
	}
	a.amount -= value
	return nil
}

func (a *Account) Balance() Real {
	return a.amount
}
