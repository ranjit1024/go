package main

import "fmt"

func zero_val(ival int) {
	ival = 0
}
func zero_valPointer(inter *int) {
	*inter = 0
}
func main() {
	fmt.Println("data")
	i := 1
	zero_val(i)
	zero_valPointer(&i)
}
