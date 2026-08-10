package main

import (
	"fmt"
)

func main() {
	m := make(map[string]int)
	fmt.Println(m)
	m["k1"] = 7
	fmt.Println(m)
	fmt.Println(m["k1"])
	fmt.Println("Data is good")
}
