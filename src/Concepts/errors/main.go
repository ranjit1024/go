package main

import "concepts/internal/runner"

func main() {
	runner.Run(map[string]func(){
		"errors":          demoErrors,
		"custom":          demoCustomErrors,
		"custom-sum":      demoSum,
		"custom-addition": demoOperations,
		"panic":           panic_demo,
	})
}
