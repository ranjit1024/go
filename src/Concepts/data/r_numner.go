package main

import (
	"fmt"
	"math/rand/v2"
)

func r_number() {
	p := fmt.Println
	fmt.Println("Hello World...")
	p(rand.IntN(100), ",")
}
