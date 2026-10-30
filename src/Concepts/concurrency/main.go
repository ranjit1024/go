package main

import "concepts/internal/runner"

func main() {
	runner.Run(map[string]func(){
		"goroutines":          demoGoroutines,
		"channels":            demoChannelPipelines,
		"buffered-channels":   demoBasicChannel,
		"channel-sync":        demoChannelSync,
		"channel-directions":  demoChannelDirections,
		"closing-channels":    demoClosingChannels,
		"range-over-channel":  demoRangeOverChannel,
		"select":              demoSelect,
		"non-blocking":        demoNonBlocking,
		"mutex":               demo_mutex,
		"atomic-counter":      demo_automic_counter,
		"wait-group":          demo_waitgroups,
		"worker-pool":         worker_pull,
		"rate-limiting":       demo_rateLimiter,
		"stateful-goroutines": state_goroutine,
	})
}
