package main

import (
	"errors"
	"fmt"
)

func divide(a, b int) (int, error) {
	if b == 0 {
		return -1, errors.New("cannot divide by zero")
	}
	return a / b, nil
}

func demoErrors() {
	res, err := divide(12, 12)
	fmt.Println(res, err)

	res, err = divide(12, 0)
	fmt.Println(res, err)
}
