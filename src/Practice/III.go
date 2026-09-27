package main

import (
	"errors"
	"fmt"
)

var notSufficentBalnce = errors.New("No sufficent balcnace")

type Account struct {
	Id      int
	Name    string
	Balance int
}

func (a *Account) deposite(amount int) int {
	a.Balance = a.Balance + amount
	return a.Balance
}

func (a *Account) WithDraw(amount int) error {
	if a.Balance < amount {
		return notSufficentBalnce
	} else {
		a.Balance = a.Balance - amount
		return nil
	}
}

func (a *Account) GetBalance() int {
	return a.Balance
}

func III() {
	acc1 := Account{
		Id:      1,
		Name:    "Ranjit",
		Balance: 40000,
	}
	acc1.deposite(1000)
	fmt.Println(acc1.GetBalance())
	with_draw := acc1.WithDraw(100000)
	if with_draw != nil {
		fmt.Println("withDraw failed ")
	} else {
		fmt.Println("Withdrawl successful")
	}
}
