package main

import (
	"fmt"
	"os"
)

func write_using_crate_and_write() {
	message := "Hello Goloang..."
	// data := []byte(message)
	// err := os.WriteFile("./tmp/a.txt", data, 0644)
	// check(err)
	// fmt.Println(message)
	fmt.Println(message)
	file, err := os.Create("./tmp/data.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	fmt.Println("file Created...")
}
