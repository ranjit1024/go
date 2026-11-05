package main

import (
	"concepts/internal/runner"
	"fmt"
)

func main() {
	fmt.Println("---------------------XXX------------")
	runner.Run(map[string]func(){
		"sha": demo_sha,
	})
}
