package main

import "fmt"

func devide(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("Cannot devide by zero")
	}
	return a / b, nil
}
func main() {
	result, err := devide(10, 0)
	fmt.Println(result)
	fmt.Println(err)
	fmt.Println("Data is the oil")
}
