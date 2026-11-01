package main

import "practice/internal/runner"

func main() {
	runner.Run(map[string]func(){
		"strings": demoStrings,
		"regex":   demoRegex,
	})
}
