package main

import "concepts/internal/runner"

func main() {
	runner.Run(map[string]func(){
		"timers":   demoTimers,
		"tickers":  tickerDemo,
		"timeouts": demoTimeouts,
	})
}
