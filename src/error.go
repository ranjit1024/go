package main

import "fmt"

func devide(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("Cannot devide by zero")
	}
	return a / b, nil
}
func main() {
	var x = 12
	fmt.Println(x)
}
