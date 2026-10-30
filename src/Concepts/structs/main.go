package main

import "concepts/internal/runner"

func main() {
	runner.Run(map[string]func(){
		"structs":    demoStructs,
		"embedding":  demoStructEmbedding,
		"interfaces": demoInterfaces,
		"enums":      demoEnums,
	})
}
