package main

import (
	"fmt"
	"strconv"
)

func parse_number() {
	f, _ := strconv.ParseFloat("1.12121", 64)
	fmt.Println(f + 10)
}
