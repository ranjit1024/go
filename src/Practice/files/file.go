package main

import (
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func read_file() {
	data, err := os.ReadFile("tmp/a.txt")
	check(err)
	content := string(data)
	fmt.Println(content)

}
