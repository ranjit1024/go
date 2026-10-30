package main

import "concepts/internal/runner"

func main() {
	runner.Run(map[string]func(){
		"generics":    demo_gnerices,
		"linked-list": demoRangeIterator,
		"iterators":   demoRangeFunc,
	})
}
