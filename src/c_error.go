package main

import "fmt"

type argError struct {
	arg     int
	message string
}

func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.message)
}
func f(arg int) (int, error) {
	if arg == 42 {
		return -1, &argError{arg: arg, message: "Can't work with it"}
	}
	return arg + 3, nil
}
func main() {
	var a int
	_, err := f(42)
	fmt.Println(err)
	fmt.Scanf("%d", &a)
	fmt.Println(a)

	fmt.Println("Data is the king")
}
