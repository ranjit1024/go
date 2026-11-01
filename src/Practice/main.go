package main

import "fmt"

func main() {
	fmt.Println("Go practice demos")
	fmt.Println()
	fmt.Println("List the demos in a category:")
	for _, category := range []string{
		"concurrency",
		"errors",
		"exercises",
		"generics",
		"sorting",
		"strings",
		"structs",
		"templates",
		"time",
	} {
		fmt.Printf("  go run ./%s\n", category)
	}
	fmt.Println()
	fmt.Println("Run a specific demo:")
	fmt.Println("  go run ./concurrency select")
}
