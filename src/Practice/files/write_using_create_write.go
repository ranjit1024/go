package main

import (
	"fmt"
	"os"
)

func write_using_crate_and_write() {
	fmt.Println("Hello World..")

	data, err := os.ReadFile("./tmp/a.txt")
	check(err)
	fmt.Println(string(data))
}
