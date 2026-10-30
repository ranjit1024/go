# Concepts

Small, topic-focused Go programs. Every file in this package defines helpers
for one concept; `main.go` calls the demo that is currently in focus.

## Run

```bash
go run .
```

## Index

### Basics

| File | Demo |
| --- | --- |
| [variables.go](variables.go) | Variable declarations and zero values |
| [values.go](values.go) | Literals and basic value types |
| [loops.go](loops.go) | `for` loops |
| [conditionals.go](conditionals.go) | `if`, `else`, and boolean logic |
| [switch.go](switch.go) | `switch` statements |
| [range.go](range.go) | Range over collections |
| [strings.go](strings.go) | String operations and template strings |
| [pointers.go](pointers.go) | Pointer semantics |
| [defer_demo.go](defer_demo.go) | Deferred calls |
| [sandbox.go](sandbox.go) | Scratch space for small experiments |

### Functions

| File | Demo |
| --- | --- |
| [functions.go](functions.go) | Function basics and multiple return values |
| [functions2.go](functions2.go) | More function signatures |
| [closure.go](closure.go) | Closures |
| [recursion.go](recursion.go) | Recursion |
| [variadic.go](variadic.go) | Variadic functions |

### Collections

| File | Demo |
| --- | --- |
| [arrays.go](arrays.go) | Arrays |
| [slices.go](slices.go) | Slices |
| [slices2.go](slices2.go) | Slice operations |
| [maps.go](maps.go) | Maps |

### Structs and interfaces

| File | Demo |
| --- | --- |
| [structs.go](structs.go) | Struct definitions and methods |
| [struct_embedding.go](struct_embedding.go) | Embedded structs |
| [interface.go](interface.go) | Interfaces and dynamic dispatch |
| [enums.go](enums.go) | Enum-style constants |

### Errors

| File | Demo |
| --- | --- |
| [errors.go](errors.go) | Basic error returns |
| [custom_error.go](custom_error.go) | Custom error types |
| [custom_error_sum.go](custom_error_sum.go) | Custom errors from a `Sum` helper |
| [custom_error_addition.go](custom_error_addition.go) | Custom errors from an `add` helper |
| [panic.go](panic.go) | Panics |

### Generics and iterators

| File | Demo |
| --- | --- |
| [generics.go](generics.go) | Generic functions |
| [generic_list.go](generic_list.go) | Generic linked list |
| [range_over_iterator.go](range_over_iterator.go) | Range-over-func iterators |

### Sorting

| File | Demo |
| --- | --- |
| [sorting.go](sorting.go) | Sorting slices |
| [sorting_by_function.go](sorting_by_function.go) | Sorting with comparison functions |

### Concurrency

| File | Demo |
| --- | --- |
| [goroutines.go](goroutines.go) | Goroutines |
| [channels.go](channels.go) | Channel communication |
| [buffered_channel.go](buffered_channel.go) | Buffered channels |
| [channel_sync.go](channel_sync.go) | Synchronizing with channels |
| [channel_direction.go](channel_direction.go) | Channel directions |
| [closing_channel.go](closing_channel.go) | Closing channels |
| [range_over_channel.go](range_over_channel.go) | Ranging over channels |
| [select.go](select.go) | `select` across channels |
| [non_blocking_channel.go](non_blocking_channel.go) | Non-blocking channel operations |
| [mutex.go](mutex.go) | Mutexes |
| [atomic_counter.go](atomic_counter.go) | Atomic counters |
| [wait_group.go](wait_group.go) | Wait groups |
| [worker_pool.go](worker_pool.go) | Worker pools |
| [rate_limiting.go](rate_limiting.go) | Rate limiting |
| [stateful_goroutines.go](stateful_goroutines.go) | Stateful goroutines |

### Time

| File | Demo |
| --- | --- |
| [timers.go](timers.go) | Timers |
| [tickers.go](tickers.go) | Tickers |
| [timeouts.go](timeouts.go) | Timeouts with `select` |
