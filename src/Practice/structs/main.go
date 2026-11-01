package main

import "practice/internal/runner"

func main() {
	runner.Run(map[string]func(){
		"enums":      demoEnums,
		"interfaces": demoInterfaces,
		"embedding":  demoEmbedding,
	})
}
