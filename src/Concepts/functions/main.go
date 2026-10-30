package main

import "concepts/internal/runner"

func main() {
	runner.Run(map[string]func(){
		"functions":        demoFunctions,
		"multiple-returns": demoMultipleReturns,
		"closures":         demoClosure,
		"recursion":        demoRecursion,
		"variadic":         demoVariadic,
	})
}
