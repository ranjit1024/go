package main

import "practice/internal/runner"

func main() {
	runner.Run(map[string]func(){
		"errors":        demoErrors,
		"custom-errors": demoCustomErrors,
		"recover":       demoRecover,
	})
}
