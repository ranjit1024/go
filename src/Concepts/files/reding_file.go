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
	// opeining a file
	f, err := os.Open(path)
	check(err)
	b1 := make([]byte, 5)
	n1, err := f.Read(b1)
	check(err)
	fmt.Printf("%d bytes : %s\n", n1, string(b1[:n1]))
}
