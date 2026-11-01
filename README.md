# Go Practice

A collection of Go concept demos and practice exercises, kept as two standalone
Go modules.

## Project structure

```text
.
├── docs/                  # Notes, exercises, and roadmap
├── src/
│   ├── Concepts/          # Topic-by-topic Go concept demos
│   │   ├── basics/
│   │   ├── functions/
│   │   ├── collections/
│   │   ├── structs/
│   │   ├── errors/
│   │   ├── generics/
│   │   ├── concurrency/
│   │   ├── time/
│   │   ├── sorting/
│   │   └── internal/runner/
│   └── Practice/          # Practice exercises and mini projects
│       ├── concurrency/
│       ├── errors/
│       ├── exercises/
│       ├── generics/
│       ├── sorting/
│       ├── strings/
│       ├── structs/
│       ├── templates/
│       ├── time/
│       └── internal/runner/
├── Makefile
└── README.md
```

Both modules have their own `go.mod` and are split into topic packages; each
package lists its demos and can run them by name.

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

List and run individual concept demos:

```bash
cd src/Concepts
go run ./concurrency              # list demos in the category
go run ./concurrency select       # run one demo
go run ./basics variables loops   # run several demos

make run-concept CATEGORY=concurrency DEMO=select
```

List and run individual practice demos:

```bash
cd src/Practice
go run ./concurrency                 # list demos in the category
go run ./concurrency channels select # run one or more demos

make run-practice-demo PRACTICE_CATEGORY=concurrency PRACTICE_DEMO=select
```

## Development tasks

```bash
make fmt      # format Go sources
make vet      # run go vet on all packages
make build    # build all packages and binaries into bin/
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
