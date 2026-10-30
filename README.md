# Go Practice

A collection of Go concept demos and practice exercises, kept as two standalone
Go modules.

## Project structure

```text
.
├── docs/                  # Notes, exercises, and roadmap
├── src/
│   ├── Concepts/          # Topic-by-topic Go concept demos
│   └── Practice/          # Practice exercises and mini projects
├── Makefile
└── README.md
```

Each module has its own `go.mod` and `main.go`. The `main.go` file selects the
demo or exercise that is currently being explored; the remaining files provide
the implementations.

## Requirements

- Go 1.25 or newer

## Run

```bash
make run-concepts
make run-practice
```

Or run a module directly:

```bash
cd src/Concepts && go run .
cd src/Practice && go run .
```

## Development tasks

```bash
make fmt      # format Go sources
make vet      # run go vet on both modules
make build    # build both modules into bin/
make clean    # remove build artifacts
```

## Topics covered

- Basics: variables, values, loops, conditionals, switch, range, strings
- Functions: closures, recursion, variadic functions
- Collections: arrays, slices, maps
- Structs, interfaces, embedding, enums
- Errors: custom errors, panic
- Generics and range-over-func iterators
- Concurrency: goroutines, channels, select, mutexes, wait groups, worker
  pools, rate limiting, stateful goroutines
- Time: timers, tickers, timeouts
- Sorting and other standard library examples

## Documentation

- [Understanding notes](docs/understanding.md)
- [Practice exercises](docs/practice-exercises.md)
- [Learning roadmap](docs/roadmap.md)
- [Select notes](docs/select.md)
- [Stateful goroutines exercise](docs/stateful-goroutines-exercise.md)

## License

[MIT](LICENSE)
