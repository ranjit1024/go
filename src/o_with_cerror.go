package main

import "fmt"

type numError struct {
	num     int
	message string
}

func (e *numError) Error() string {
	return fmt.Sprintf("%d - %s", e.num, e.message)
}

func Sum(a, b int) (int, error) {
	if b == 0 {
		return -1, &numError{num: b, message: "Canaot work with it"}
	}
	return a + b, nil
}
