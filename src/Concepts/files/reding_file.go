package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func read_file() {
	fmt.Println("Hello World ..")
	path := filepath.Join("tmp", "a.txt")
	data, err := os.ReadFile(path)
	check(err)
	fmt.Println(string(data))
}
