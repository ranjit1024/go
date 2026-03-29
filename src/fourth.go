package main

import (
	"fmt"
)

func main() {
	fmt.Println("Data")
	m := make(map[string]int)
	_, prs := m["k2"]
	fmt.Println(prs)
	m["ranjit"] = 12
	m["pi"] = 31415

	fmt.Println(m, len(m))
}
