package main

import "concepts/internal/runner"

func main() {
	runner.Run(map[string]func(){
		"sorting":             demo_sort,
		"sorting-by-function": sorting_by_funcion,
	})
}
