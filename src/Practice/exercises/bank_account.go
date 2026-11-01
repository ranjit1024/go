package main

import (
	"errors"
	"fmt"
	"sync"
)

var errInsufficientBalance = errors.New("insufficient balance")

type Account struct {
	ID      int
	Name    string
	Balance int
	mu      sync.Mutex
}

func (a *Account) Deposit(amount int) int {
	a.mu.Lock()
	defer a.mu.Unlock()
	a.Balance += amount
	return a.Balance
}

func (a *Account) Withdraw(amount int) error {
	a.mu.Lock()
	defer a.mu.Unlock()
	if a.Balance < amount {
		return errInsufficientBalance
	}
	a.Balance -= amount
	return nil
}

func (a *Account) GetBalance() int {
	a.mu.Lock()
	defer a.mu.Unlock()
	return a.Balance
}

func demoBankAccount() {
	account := Account{ID: 1, Name: "Ranjit"}

	var deposits sync.WaitGroup
	for range 100 {
		deposits.Go(func() {
			account.Deposit(100)
		})
	}
	deposits.Wait()

	var withdrawals sync.WaitGroup
	for range 50 {
		withdrawals.Go(func() {
			if err := account.Withdraw(100); err != nil {
				fmt.Println("withdraw failed:", err)
			}
		})
	}
	withdrawals.Wait()

	fmt.Println("balance:", account.GetBalance())
	if err := account.Withdraw(1_000_000); err != nil {
		fmt.Println("withdraw failed:", err)
	}
}
