package account

import (
	"sync"
)

type Account struct {
	Bal    int64
	Closed bool
	mux    sync.Mutex
}

func Open(initialDeposit int64) *Account {
	if initialDeposit < 0 {
		return nil
	}
	return &Account{Bal: initialDeposit}
}

func (acct *Account) Close() (payout int64, ok bool) {
	acct.mux.Lock()
	defer acct.mux.Unlock()

	if acct.Closed {
		payout = 0
		ok = false
	} else {
		acct.Closed = true
		payout = acct.Bal
		ok = true
	}

	return
}

func (acct *Account) Balance() (balance int64, ok bool) {
	acct.mux.Lock()
	defer acct.mux.Unlock()

	if acct.Closed {
		balance = 0
		ok = false
	} else {
		balance = acct.Bal
		ok = true
	}

	return
}

func (acct *Account) Deposit(amount int64) (newBalance int64, ok bool) {
	acct.mux.Lock()
	defer acct.mux.Unlock()

	if acct.Closed {
		newBalance = 0
		ok = false
	} else {
		if newBalance = acct.Bal + amount; newBalance < 0 {
			newBalance = acct.Bal
			ok = false
		} else {
			acct.Bal = newBalance
			ok = true
		}
	}

	return
}
