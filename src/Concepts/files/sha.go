package main

import (
	"crypto/sha256"
	"fmt"
	"strconv"
)

func demo_sha() {
	s := "Sha256 this string"
	h := sha256.New()
	h.Write([]byte(s))
	bs := h.Sum(nil)
	fmt.Println(s)
	fmt.Printf("%x \n", bs)
	fmt.Println("Hello World....")
	// ----------- Number pasing-------------
	n, err := strconv.ParseInt("1010", 2, 64)
	if err != nil {
		panic("incorret data")
	} //

	fmt.Println(n)
}
