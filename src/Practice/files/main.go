package main

import (
	"fmt"
	"practice/internal/runner"
)

func main() {
	fmt.Println("fadfa")
	runner.Run(map[string]func(){
		"number": number_demo,
		"file":   read_file,
	})
}
