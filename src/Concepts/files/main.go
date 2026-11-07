package main

import (
	"concepts/internal/runner"
	"fmt"
)

func main() {
	fmt.Println("---------------------XXX------------")
	runner.Run(map[string]func(){
		"sha":     demo_sha,
		"base_64": base_demo,
		"read":    read_file,
	})
}
