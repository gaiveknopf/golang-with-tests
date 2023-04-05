package main

import "fmt"

type Bitcoin float64

type Stringer interface {
	String() string
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%f BTC", b)
}

type Wallet struct {
	amount Bitcoin
}

func (w *Wallet) Deposit(amount Bitcoin) {
	w.amount += amount
}

func (w *Wallet) Balance() Bitcoin {
	return w.amount
}

var ErrInsufficientFunds = fmt.Errorf("cannot withdraw, insufficient funds")

func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.amount {
		return ErrInsufficientFunds
	}
	w.amount -= amount
	return nil
}
