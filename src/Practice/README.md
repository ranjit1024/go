# Practice

Practice exercises and small projects built while learning Go. Demos are
grouped by topic, and each topic folder is a standalone `package main` built on
top of a shared demo runner, so examples can be listed and run individually.

## Layout

```text
Practice/
├── concurrency/        # goroutines, channels, sync, worker pools, rate limiting
├── errors/             # error handling, custom errors, panic recovery
├── exercises/          # graded exercises and mini projects
├── generics/           # generic helpers, range-over-func iterators
├── sorting/            # sorting with comparison functions
├── strings/            # string helpers, regular expressions
├── structs/            # structs, embedding, interfaces, enums
├── templates/          # text templates
├── time/               # timers, tickers, timeouts
├── internal/runner/    # shared demo runner
└── main.go             # category menu
```

## Run

```bash
cd src/Practice

go run .                             # show the category menu
go run ./concurrency                 # list the demos in a category
go run ./concurrency channels select # run one or more demos
```

## Index

### Concurrency

| Demo | File | Description |
| --- | --- | --- |
| `goroutines` | [goroutines.go](concurrency/goroutines.go) | Starting goroutines |
| `channels` | [channels.go](concurrency/channels.go) | Channel basics |
| `buffered-channels` | [buffered_channels.go](concurrency/buffered_channels.go) | Buffered channels |
| `channel-sync` | [channel_sync.go](concurrency/channel_sync.go) | Using channels to synchronize |
| `channel-directions` | [channel_directions.go](concurrency/channel_directions.go) | Send-only and receive-only channels |
| `closing-channels` | [closing_channels.go](concurrency/closing_channels.go) | Closing channels and ranging over them |
| `non-blocking` | [non_blocking_channels.go](concurrency/non_blocking_channels.go) | Non-blocking channel operations |
| `select` | [select.go](concurrency/select.go) | Waiting on multiple channels |
| `wait-groups` | [wait_groups.go](concurrency/wait_groups.go) | Waiting for goroutines |
| `mutexes` | [mutexes.go](concurrency/mutexes.go) | Mutual exclusion |
| `atomic-counters` | [atomic_counters.go](concurrency/atomic_counters.go) | Atomic counters |
| `stateful-goroutines` | [stateful_goroutines.go](concurrency/stateful_goroutines.go) | State owned by a goroutine |
| `worker-pools` | [worker_pools.go](concurrency/worker_pools.go) | Worker pools |
| `rate-limiting` | [rate_limiting.go](concurrency/rate_limiting.go) | Rate limiting |
| `rate-limiting-burst` | [rate_limiting_burst.go](concurrency/rate_limiting_burst.go) | Burst rate limiting |
| `key-value-store` | [key_value_store.go](concurrency/key_value_store.go) | Channel-based key-value store |

### Errors

| Demo | File | Description |
| --- | --- | --- |
| `errors` | [errors.go](errors/errors.go) | Returning and handling errors |
| `custom-errors` | [custom_errors.go](errors/custom_errors.go) | Custom error types with `errors.As` |
| `recover` | [recover.go](errors/recover.go) | Recovering from panics |

### Exercises

| Demo | File | Description |
| --- | --- | --- |
| `grades` | [grades.go](exercises/grades.go) | Student grade analyzer |
| `word-count` | [word_count.go](exercises/word_count.go) | Word frequency counter |
| `bank-account` | [bank_account.go](exercises/bank_account.go) | Concurrent bank account |

### Generics

| Demo | File | Description |
| --- | --- | --- |
| `generics` | [generics.go](generics/generics.go) | Generic helpers |
| `iterators` | [iterators.go](generics/iterators.go) | Range-over-func iterators |

### Sorting

| Demo | File | Description |
| --- | --- | --- |
| `sorting-by-function` | [sorting_by_function.go](sorting/sorting_by_function.go) | Sorting with comparison functions |
