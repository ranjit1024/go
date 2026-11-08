package main

import (
	"fmt"
	"path/filepath"
)

func read_file() {
	path := filepath.Join("tmp", "a.txt")
	fmt.Println(path)
	fmt.Println("Hello World")

}
