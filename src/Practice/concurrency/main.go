package main

import "practice/internal/runner"

func main() {
	runner.Run(map[string]func(){
		"goroutines":          demoGoroutines,
		"channels":            demoChannels,
		"buffered-channels":   demoBufferedChannels,
		"channel-sync":        demoChannelSync,
		"channel-directions":  demoChannelDirections,
		"closing-channels":    demoClosingChannels,
		"non-blocking":        demoNonBlockingChannels,
		"select":              demoSelect,
		"wait-groups":         demoWaitGroups,
		"mutexes":             demoMutexes,
		"atomic-counters":     demoAtomicCounters,
		"stateful-goroutines": demoStatefulGoroutines,
		"worker-pools":        demoWorkerPools,
		"rate-limiting":       demoRateLimiting,
		"rate-limiting-burst": demoRateLimitingBurst,
		"key-value-store":     demoKeyValueStore,
	})
}
