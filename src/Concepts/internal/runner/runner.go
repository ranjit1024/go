package runner

import (
	"fmt"
	"os"
	"sort"
)

func Run(demos map[string]func()) {
	if len(os.Args) < 2 {
		names := make([]string, 0, len(demos))
		for name := range demos {
			names = append(names, name)
		}
		sort.Strings(names)

		fmt.Println("Available demos:")
		for _, name := range names {
			fmt.Printf("  %s\n", name)
		}
		fmt.Println()
		fmt.Println("Run one or more demos by name:")
		fmt.Println("  go run <category> <demo> [demo...]")
		return
	}

	for _, name := range os.Args[1:] {
		demo, ok := demos[name]
		if !ok {
			fmt.Fprintf(os.Stderr, "unknown demo %q\n", name)
			os.Exit(1)
		}
		fmt.Printf("\n=== %s ===\n", name)
		demo()
	}
}
