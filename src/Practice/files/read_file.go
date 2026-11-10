package main

import (
	"fmt"
	"io"
	"os"
)

func check_new(e error) {
	if e != nil {
		panic(e)
	}
}
func read_file_5_bytes() {
	file, err := os.Open("tmp/a.txt")
	check_new(err)
	defer file.Close()

	buffer := make([]byte, 5)
	for {
		n, err := file.Read(buffer)
		if n > 0 {
			fmt.Printf("Read %q \n", string(buffer[:n]))
		}
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}

	}
	fmt.Println("Hello World...")
}
