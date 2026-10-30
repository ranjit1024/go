# Concepts

A collection of small Go demos grouped by topic. Each topic folder is a
standalone `package main` built on top of a shared demo runner, so examples can
be listed and run individually.

## Layout

```text
Concepts/
├── basics/             # variables, loops, strings, pointers, defer
├── functions/          # functions, closures, recursion, variadic
├── collections/        # arrays, slices, maps
├── structs/            # structs, embedding, interfaces, enums
├── errors/             # error handling, custom errors, panic
├── generics/           # generics, linked list, iterators
├── concurrency/        # goroutines, channels, select, sync, worker pools
├── time/               # timers, tickers, timeouts
├── sorting/            # sorting and comparison functions
├── internal/runner/    # shared demo runner
└── main.go             # category menu
```

## Run

```bash
cd src/Concepts

go run .                      # show the category menu
go run ./concurrency          # list the demos in a category
go run ./concurrency select   # run one demo

go run ./basics variables loops   # run several demos in order
```

## Index

### Basics

| Demo | File | Description |
| --- | --- | --- |
| `variables` | [variables.go](basics/variables.go) | Variable declarations and zero values |
| `values` | [values.go](basics/values.go) | Literals and basic value types |
| `loops` | [loops.go](basics/loops.go) | `for` loops |
| `conditionals` | [conditionals.go](basics/conditionals.go) | `if`, `else`, and boolean logic |
| `switch` | [switch.go](basics/switch.go) | `switch` statements |
| `range` | [range.go](basics/range.go) | Range over collections |
| `strings` | [strings.go](basics/strings.go) | String operations and template strings |
| `pointers` | [pointers.go](basics/pointers.go) | Pointer semantics |
| `defer` | [defer.go](basics/defer.go) | Deferred calls |
| `sandbox` | [sandbox.go](basics/sandbox.go) | Scratch space for small experiments |

### Functions

| Demo | File | Description |
| --- | --- | --- |
| `functions` | [functions.go](functions/functions.go) | Function basics and multiple return values |
| `multiple-returns` | [multiple_returns.go](functions/multiple_returns.go) | More function signatures |
| `closures` | [closure.go](functions/closure.go) | Closures |
| `recursion` | [recursion.go](functions/recursion.go) | Recursion |
| `variadic` | [variadic.go](functions/variadic.go) | Variadic functions |

### Collections

| Demo | File | Description |
| --- | --- | --- |
| `arrays` | [arrays.go](collections/arrays.go) | Arrays |
| `slices` | [slices.go](collections/slices.go) | Slices |
| `more-slices` | [more_slices.go](collections/more_slices.go) | Slice operations |
| `maps` | [maps.go](collections/maps.go) | Maps |

### Structs and interfaces

| Demo | File | Description |
| --- | --- | --- |
| `structs` | [structs.go](structs/structs.go) | Struct definitions and methods |
| `embedding` | [embedding.go](structs/embedding.go) | Embedded structs |
| `interfaces` | [interfaces.go](structs/interfaces.go) | Interfaces and dynamic dispatch |
| `enums` | [enums.go](structs/enums.go) | Enum-style constants |

### Errors

| Demo | File | Description |
| --- | --- | --- |
| `errors` | [errors.go](errors/errors.go) | Basic error returns |
| `custom` | [custom.go](errors/custom.go) | Custom error types |
| `custom-sum` | [custom_sum.go](errors/custom_sum.go) | Custom errors from a `Sum` helper |
| `custom-addition` | [custom_addition.go](errors/custom_addition.go) | Custom errors from an `add` helper |
| `panic` | [panic.go](errors/panic.go) | Panics |

### Generics and iterators

| Demo | File | Description |
| --- | --- | --- |
| `generics` | [generics.go](generics/generics.go) | Generic functions |
| `linked-list` | [linked_list.go](generics/linked_list.go) | Generic linked list |
| `iterators` | [iterators.go](generics/iterators.go) | Range-over-func iterators |

### Concurrency

| Demo | File | Description |
| --- | --- | --- |
| `goroutines` | [goroutines.go](concurrency/goroutines.go) | Goroutines |
| `channels` | [channels.go](concurrency/channels.go) | Channel pipelines |
| `buffered-channels` | [buffered_channels.go](concurrency/buffered_channels.go) | Buffered channels |
| `channel-sync` | [channel_sync.go](concurrency/channel_sync.go) | Synchronizing with channels |
| `channel-directions` | [channel_directions.go](concurrency/channel_directions.go) | Channel directions |
| `closing-channels` | [closing_channels.go](concurrency/closing_channels.go) | Closing channels |
| `range-over-channel` | [range_over_channels.go](concurrency/range_over_channels.go) | Ranging over channels |
| `select` | [select.go](concurrency/select.go) | `select` across channels |
| `non-blocking` | [non_blocking_channels.go](concurrency/non_blocking_channels.go) | Non-blocking channel operations |
| `mutex` | [mutexes.go](concurrency/mutexes.go) | Mutexes |
| `atomic-counter` | [atomic_counters.go](concurrency/atomic_counters.go) | Atomic counters |
| `wait-group` | [wait_groups.go](concurrency/wait_groups.go) | Wait groups |
| `worker-pool` | [worker_pools.go](concurrency/worker_pools.go) | Worker pools |
| `rate-limiting` | [rate_limiting.go](concurrency/rate_limiting.go) | Rate limiting |
| `stateful-goroutines` | [stateful_goroutines.go](concurrency/stateful_goroutines.go) | Stateful goroutines |

### Time

| Demo | File | Description |
| --- | --- | --- |
| `timers` | [timers.go](time/timers.go) | Timers |
| `tickers` | [tickers.go](time/tickers.go) | Tickers |
| `timeouts` | [timeouts.go](time/timeouts.go) | Timeouts with `select` |

### Sorting

| Demo | File | Description |
| --- | --- | --- |
| `sorting` | [sorting.go](sorting/sorting.go) | Sorting slices |
| `sorting-by-function` | [sorting_by_function.go](sorting/sorting_by_function.go) | Sorting with comparison functions |
