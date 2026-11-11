package main

import (
	"practice/internal/runner"
)

func main() {
	runner.Run(map[string]func(){
		"number":     number_demo,
		"file":       read_file,
		"5byte":      read_file_5_bytes,
		"file_write": write_demo,
		"write":      write_using_crate_and_write,
	})
}
