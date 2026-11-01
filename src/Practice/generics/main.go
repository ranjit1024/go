package main

import "practice/internal/runner"

func main() {
	runner.Run(map[string]func(){
		"generics":  demoGenerics,
		"iterators": demoIterators,
	})
}
