package main

import (
	"fmt"
	"math/rand/v2"
)

func r_number() {
	p := fmt.Println
	fmt.Println("Hello World...")
	p(rand.IntN(100), ",")
	s2 := rand.NewPCG(42, 1024)
	p(rand.New(s2))
}
