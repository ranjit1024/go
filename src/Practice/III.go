package main

import (
	"errors"
	"fmt"
	"sync"
)

var notSufficentBalnce = errors.New("No sufficent balcnace")

type Account struct {
	Id      int
	Name    string
	Balance int
	mu      sync.Mutex
}

func (a *Account) deposite(amount int) int {
	a.mu.Lock()
	defer a.mu.Unlock()
	fmt.Println("Update", a.Balance, "+", amount, "--------->", a.Balance+amount)
	a.Balance = a.Balance + amount
	return a.Balance
}
func (a *Account) WithDraw(amount int) error {
	a.mu.Lock()
	defer a.mu.Unlock()
	if a.Balance < amount {
		return notSufficentBalnce
	} else {
		a.Balance = a.Balance - amount
		return nil
	}
}

func (a *Account) GetBalance() int {
	a.mu.Lock()
	defer a.mu.Unlock()
	return a.Balance
}

func III() {
	acc1 := Account{
		Id:      1,
		Name:    "Ranjit",
		Balance: 0,
	}
	var wg sync.WaitGroup

	for i := range 100 {
		fmt.Println(i)
		wg.Add(1)
		go func() {
			defer wg.Done()
			acc1.deposite(10)
		}()
	}
	wg.Wait()

	fmt.Println("Currunet Balance ", acc1.Balance)
}
