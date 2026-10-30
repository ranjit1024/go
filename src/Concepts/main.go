package main

import "fmt"

func main() {
	fmt.Println("Go concept demos")
	fmt.Println()
	fmt.Println("List the demos in a category:")
	for _, category := range []string{
		"basics",
		"functions",
		"collections",
		"structs",
		"errors",
		"generics",
		"concurrency",
		"time",
		"sorting",
	} {
		fmt.Printf("  go run ./%s\n", category)
	}
	fmt.Println()
	fmt.Println("Run a specific demo:")
	fmt.Println("  go run ./concurrency select")
}
