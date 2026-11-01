package main

import "practice/internal/runner"

func main() {
	runner.Run(map[string]func(){
		"timers":   demoTimers,
		"tickers":  demoTickers,
		"timeouts": demoTimeouts,
	})
}
