package pointer

import "errors"

type Bitcoin int

type Wallet struct {
	balance Bitcoin
}

func (w *Wallet) deposit(amount Bitcoin) {
	w.balance += amount
}

func (w *Wallet) getBalance() Bitcoin {
	return w.balance
}

func (w *Wallet) withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return errors.New("exceed balance")
	}
	w.balance -= amount
	return nil
}


