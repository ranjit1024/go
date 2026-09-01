package main

import "fmt"

type argError struct {
	arg     int
	message string
}

func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.message)
}
func main() {
	var a int
	fmt.Scanf("%d", &a)
	fmt.Println(a)

	fmt.Println("Data is the king")
}
