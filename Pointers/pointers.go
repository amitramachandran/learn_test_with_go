package pointers

import "fmt"

type Bitcoin int

type Wallet struct {
	balance Bitcoin
}

func (w *Wallet) deposit(amount Bitcoin) {
	w.balance += amount

}

func (w *Wallet) withdraw(amount Bitcoin) {
	w.balance -= amount
}

func (w *Wallet) get_balance() Bitcoin {
	return w.balance
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}
