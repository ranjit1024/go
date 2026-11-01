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
