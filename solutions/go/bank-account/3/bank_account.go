package account

import "sync"

type Account struct {
	active  bool
	balance int64
	lock    *sync.Mutex
}

func Open(amount int64) *Account {
	if amount < 0 {
		return nil
	}

	return &Account{balance: amount, active: true, lock: &sync.Mutex{}}
}

func (a *Account) Balance() (int64, bool) {
	defer a.lock.Unlock()
	a.lock.Lock()

	if !a.active {
		return 0, false
	}
	return a.balance, true
}

func (a *Account) Deposit(amount int64) (int64, bool) {
	defer a.lock.Unlock()
	a.lock.Lock()

	if !a.active {
		return a.balance, false
	}

	if amount >= 0 {
		a.balance += amount
		return a.balance, true
	}

	if a.balance+amount < 0 {
		return a.balance, false
	}

	a.balance += amount

	return a.balance, true
}

func (a *Account) Close() (int64, bool) {
	defer a.lock.Unlock()
	a.lock.Lock()

	if !a.active {
		return 0, false
	}

	res := a.balance

	a.active = false
	a.balance = 0

	return res, true
}
