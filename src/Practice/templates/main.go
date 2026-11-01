package main

import "practice/internal/runner"

func main() {
	runner.Run(map[string]func(){
		"text-template": demoTextTemplate,
	})
}
