package main

import "concepts/internal/runner"

func main() {
	runner.Run(map[string]func(){
		"env": demo_env,
	})
}
