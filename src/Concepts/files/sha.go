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
	strconv.ParseInt("12", 2, 64)
}
