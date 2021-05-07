// Package account provides a bank account abstraction.
package account

import "sync"

// type Account interface {
// 	Close() (payout int64, ok bool)
// 	Balance() (balance int64, ok bool)
// 	Deposit(amount int64) (newBalance int64, ok bool)
// }

// Account describes a bank account. It only handled full dollars - no change.
type Account struct {
	lock    sync.RWMutex
	closed  bool
	balance int64
}

// Open is used to create a new Bank account. The opening balance cannot be
// negative, it can be 0.
func Open(initalDeposit int64) *Account {
	if initalDeposit < 0 {
		return nil
	}

	newAccount := Account{balance: initalDeposit}
	return &newAccount
}

// Close returns all funds in the account and closes the account.
func (a *Account) Close() (payout int64, ok bool) {
	a.lock.Lock()
	defer a.lock.Unlock()

	if a.closed {
		return 0, false
	}

	a.closed = true
	return a.balance, true
}

// Balance returns the current account balance. It will be positive.
func (a *Account) Balance() (balance int64, ok bool) {
	a.lock.RLock()
	defer a.lock.RUnlock()
	if a.closed {
		return 0, false
	}

	return a.balance, true
}

// Deposit is used to add money to the account.
// It is also used to remove money from the account - with a negative deposit.
// It will not allow the account balance to go negative.
func (a *Account) Deposit(amount int64) (newBalance int64, ok bool) {
	a.lock.Lock()
	defer a.lock.Unlock()
	if a.closed {
		return 0, false
	}

	if a.balance+amount < 0 {
		return a.balance, false
	}

	a.balance += amount
	return a.balance, true
}