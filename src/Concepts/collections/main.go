package main

import "concepts/internal/runner"

func main() {
	runner.Run(map[string]func(){
		"arrays":      demoArrays,
		"slices":      demoSlices,
		"more-slices": demoSlice2,
		"maps":        demoMaps,
	})
}
