package main

import (
	"os"
)

func write_demo() {
	data := []byte("Hello wold\n")
	err := os.WriteFile("./tmp/a.txt", data, 0644)
	check(err)
}
