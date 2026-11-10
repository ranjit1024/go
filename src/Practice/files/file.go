package main

import (
	"bufio"
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func read_file() {
	data, err := os.Open("tmp/a.txt")
	check(err)
	scanner := bufio.NewScanner(data)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
	}
}
