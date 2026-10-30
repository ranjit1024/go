package main

import "concepts/internal/runner"

func main() {
	runner.Run(map[string]func(){
		"variables":    demoVariables,
		"values":       demoValues,
		"loops":        loop,
		"conditionals": demoConditional,
		"switch":       demoSwitch,
		"range":        demoRange,
		"strings":      demoStrings,
		"pointers":     demoPointers,
		"defer":        demo_defer,
		"sandbox":      demoTest,
	})
}
