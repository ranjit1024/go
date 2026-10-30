# Go Skills — Project Practice Checklist

> Goal: Improve Go from a decent working level to strong real-world/backend/systems proficiency.
>
> Use C++ mainly for DSA/problem solving and use Go to build practical software, backend services, concurrent programs, and systems tools.

---

## How to Use This Roadmap

For every project:

- [ ] Define requirements before coding
- [ ] Build the first version yourself
- [ ] Read compiler/runtime errors and debug them
- [ ] Use Go documentation when stuck
- [ ] Avoid following tutorials line-by-line
- [ ] Add tests
- [ ] Refactor after the first working version
- [ ] Rebuild important parts later without looking at the original code

---

# 1. CLI Task Manager

**Difficulty:** ★★☆☆☆

Build a terminal application:

```text
todo add "learn Go channels"
todo list
todo complete 3
todo delete 3
todo search "Go"
```

### Tasks

- [ ] Design the command structure
- [ ] Create task struct
- [ ] Add tasks
- [ ] List tasks
- [ ] Complete tasks
- [ ] Delete tasks
- [ ] Search tasks
- [ ] Save tasks to JSON
- [ ] Load tasks from JSON
- [ ] Handle invalid input
- [ ] Handle errors properly
- [ ] Split code into packages
- [ ] Add unit tests
- [ ] Add help/usage output

### Go Concepts

- [ ] Structs
- [ ] Slices
- [ ] Maps
- [ ] Methods
- [ ] Interfaces
- [ ] Errors
- [ ] JSON
- [ ] File I/O
- [ ] Packages
- [ ] CLI arguments

---

# 2. Log Analyzer

Input:

```text
2026-09-14 INFO user logged in
2026-09-14 ERROR database timeout
2026-09-14 INFO request completed
```

Output:

```text
INFO:  15342
WARN:    231
ERROR:    82

Top errors:
database timeout     41
connection refused   23
```

### Tasks

- [ ] Read a log file
- [ ] Process it line by line
- [ ] Parse log levels
- [ ] Count INFO/WARN/ERROR
- [ ] Count individual error messages
- [ ] Show top N errors
- [ ] Add CLI flags
- [ ] Add filtering by log level
- [ ] Handle large files efficiently
- [ ] Process multiple files
- [ ] Add concurrent processing
- [ ] Add tests

### CLI

```bash
loganalyzer --file app.log
loganalyzer --level ERROR
loganalyzer --top 10
```

### Go Concepts

- [ ] File I/O
- [ ] Buffered I/O
- [ ] Strings
- [ ] Maps
- [ ] Structs
- [ ] CLI flags
- [ ] Concurrency
- [ ] Testing

---

# 3. Concurrent URL Checker

Input:

```text
google.com
github.com
example.com
openai.com
```

Output:

```text
google.com     200    121ms
github.com     200    182ms
example.com    404    240ms
```

### Phase 1 — Sequential

- [ ] Read URLs
- [ ] Make HTTP requests
- [ ] Record status codes
- [ ] Measure response time
- [ ] Handle errors

### Phase 2 — Concurrent

- [ ] Create goroutines
- [ ] Use channels
- [ ] Build a worker pool
- [ ] Use `sync.WaitGroup`
- [ ] Collect results
- [ ] Collect errors

### Phase 3 — Production Features

- [ ] Limit concurrency
- [ ] Add request timeout
- [ ] Add retries
- [ ] Add exponential backoff
- [ ] Add context cancellation
- [ ] Gracefully shut down workers
- [ ] Add structured output
- [ ] Measure performance

### Go Concepts

- [ ] Goroutines
- [ ] Channels
- [ ] `select`
- [ ] `sync.WaitGroup`
- [ ] `context`
- [ ] HTTP
- [ ] Timeouts
- [ ] Worker pools

---

# 4. HTTP Server From Scratch

Build:

```text
POST   /users
GET    /users
GET    /users/:id
PUT    /users/:id
DELETE /users/:id
```

Use Go's standard library first:

```go
net/http
encoding/json
```

### Tasks

- [ ] Create HTTP server
- [ ] Create routes
- [ ] Create handlers
- [ ] Parse JSON requests
- [ ] Return JSON responses
- [ ] Use correct HTTP status codes
- [ ] Validate input
- [ ] Handle errors
- [ ] Create middleware
- [ ] Add request logging
- [ ] Add request IDs
- [ ] Add authentication middleware
- [ ] Add rate limiting
- [ ] Add graceful shutdown
- [ ] Add configuration
- [ ] Add tests

### Go Concepts

- [ ] `net/http`
- [ ] HTTP handlers
- [ ] Middleware
- [ ] JSON
- [ ] Interfaces
- [ ] Context
- [ ] Error handling
- [ ] Testing

---

# 5. URL Shortener

Build:

```text
POST /shorten
GET  /:shortCode
GET  /stats/:shortCode
```

Example request:

```json
{
  "url": "https://example.com/very/long/url"
}
```

Example response:

```json
{
  "short_url": "http://localhost:8080/a8F2x"
}
```

### Tasks

- [ ] Create URL shortening endpoint
- [ ] Generate short IDs
- [ ] Store URLs
- [ ] Redirect users
- [ ] Track click count
- [ ] Track creation time
- [ ] Track last accessed time
- [ ] Add database
- [ ] Use SQL
- [ ] Handle concurrent requests
- [ ] Validate URLs
- [ ] Add tests
- [ ] Add graceful shutdown
- [ ] Add configuration

### Go Concepts

- [ ] HTTP
- [ ] Database
- [ ] SQL
- [ ] Transactions
- [ ] Concurrency
- [ ] Validation
- [ ] Configuration
- [ ] Testing

---

# 6. Job Queue

Build a background job processing system.

```text
             ┌── Worker 1
Jobs ────────┼── Worker 2
             ├── Worker 3
             └── Worker 4
```

Example jobs:

```text
send email
resize image
process file
generate report
```

API:

```text
POST /jobs
GET  /jobs/:id
```

### Tasks

- [ ] Create job struct
- [ ] Submit jobs
- [ ] Create workers
- [ ] Build job queue
- [ ] Process jobs concurrently
- [ ] Track job status
- [ ] Return job ID
- [ ] Add retries
- [ ] Add exponential backoff
- [ ] Add job priorities
- [ ] Add cancellation
- [ ] Add dead-letter queue
- [ ] Persist jobs
- [ ] Add worker health
- [ ] Gracefully shut down workers
- [ ] Add tests

### Go Concepts

- [ ] Channels
- [ ] Goroutines
- [ ] Worker pools
- [ ] Mutexes
- [ ] `sync`
- [ ] Context cancellation
- [ ] Graceful shutdown
- [ ] Persistence

---

# 7. TCP Chat Server

Build a TCP server where multiple clients can communicate.

```text
Client A ─┐
Client B ─┼──> Go Server
Client C ─┘
```

Example:

```text
Alice: hello
Bob: hi
```

### Tasks

- [ ] Create TCP listener
- [ ] Accept connections
- [ ] Handle clients concurrently
- [ ] Read messages
- [ ] Broadcast messages
- [ ] Track connected users
- [ ] Handle disconnects
- [ ] Add usernames
- [ ] Add chat rooms
- [ ] Add private messages
- [ ] Add authentication
- [ ] Add heartbeats
- [ ] Add connection timeout
- [ ] Add graceful shutdown

### Go Concepts

- [ ] `net`
- [ ] TCP
- [ ] `net.Conn`
- [ ] Goroutines
- [ ] Channels
- [ ] Mutexes
- [ ] Connection lifecycle

---

# 8. Concurrent Web Crawler

Start with:

```text
https://example.com
```

Crawl:

```text
example.com
├── /about
├── /blog
│   ├── /post1
│   └── /post2
└── /contact
```

### Requirements

- [ ] Crawl pages
- [ ] Extract links
- [ ] Avoid duplicate URLs
- [ ] Limit crawl depth
- [ ] Limit concurrency
- [ ] Handle failed requests
- [ ] Support cancellation
- [ ] Collect results
- [ ] Track visited URLs
- [ ] Add timeouts
- [ ] Add retries
- [ ] Save crawl results
- [ ] Add statistics

### Go Concepts

- [ ] Goroutines
- [ ] Channels
- [ ] Mutexes
- [ ] Maps
- [ ] HTTP
- [ ] Context
- [ ] Concurrency limits
- [ ] Error handling

---

# 9. Reverse Proxy / API Gateway

Architecture:

```text
                 ┌── Service A
Client → Go Proxy
                 ├── Service B
                 └── Service C
```

### Tasks

- [ ] Forward HTTP requests
- [ ] Configure backend services
- [ ] Implement routing
- [ ] Add load balancing
- [ ] Add health checks
- [ ] Add timeouts
- [ ] Add retries
- [ ] Add rate limiting
- [ ] Add request logging
- [ ] Add request IDs
- [ ] Add metrics
- [ ] Add graceful shutdown
- [ ] Handle unhealthy backends

### Go Concepts

- [ ] HTTP
- [ ] Networking
- [ ] Concurrency
- [ ] Context
- [ ] Middleware
- [ ] Error handling
- [ ] Systems design

---

# 10. Git-Like Tool

Build a simplified version of Git.

Commands:

```bash
mygit init
mygit add file.txt
mygit commit -m "first commit"
mygit log
```

### Tasks

- [ ] Create repository
- [ ] Track files
- [ ] Hash file contents
- [ ] Store objects
- [ ] Create commits
- [ ] Store commit metadata
- [ ] Display history
- [ ] Compare changes
- [ ] Add branches
- [ ] Add checkout
- [ ] Add basic merge

### Go Concepts

- [ ] Filesystem
- [ ] Hashing
- [ ] Serialization
- [ ] Binary data
- [ ] Directories
- [ ] CLI design
- [ ] Data structures

---

# 11. Distributed Job Processing System

## Capstone

Architecture:

```text
                  ┌──────────────┐
                  │   API Server │
                  └──────┬───────┘
                         │
                         ▼
                  ┌──────────────┐
                  │    Queue     │
                  └──────┬───────┘
                         │
             ┌───────────┼───────────┐
             ▼           ▼           ▼
          Worker 1    Worker 2    Worker 3
             │           │           │
             └───────────┼───────────┘
                         ▼
                     Database
```

### Tasks

- [ ] Build API server
- [ ] Submit jobs
- [ ] Store jobs
- [ ] Create workers
- [ ] Process jobs
- [ ] Track job status
- [ ] Add multiple workers
- [ ] Add retries
- [ ] Add job priority
- [ ] Add cancellation
- [ ] Add worker heartbeat
- [ ] Add worker health checks
- [ ] Add persistent queue
- [ ] Add graceful shutdown
- [ ] Add rate limiting
- [ ] Add structured logging
- [ ] Add metrics
- [ ] Add integration tests
- [ ] Simulate worker failure
- [ ] Handle jobs that crash halfway through
- [ ] Document architecture

---

# Supporting Go Skills

Don't learn these only when a project requires them. Gradually incorporate them.

## Standard Library

- [ ] `fmt`
- [ ] `strings`
- [ ] `strconv`
- [ ] `bytes`
- [ ] `io`
- [ ] `os`
- [ ] `path/filepath`
- [ ] `bufio`
- [ ] `encoding/json`
- [ ] `time`
- [ ] `context`
- [ ] `sync`
- [ ] `net`
- [ ] `net/http`
- [ ] `errors`
- [ ] `log`

## Error Handling

- [ ] Return errors instead of hiding them
- [ ] Wrap errors
- [ ] `errors.Is`
- [ ] `errors.As`
- [ ] Custom errors
- [ ] Sentinel errors
- [ ] Decide where errors should be handled

## Interfaces

- [ ] Understand implicit interfaces
- [ ] Small interfaces
- [ ] Interface composition
- [ ] Dependency injection
- [ ] Mocking through interfaces
- [ ] Know when NOT to use an interface

## Concurrency

- [ ] Goroutines
- [ ] Channels
- [ ] Buffered channels
- [ ] Unbuffered channels
- [ ] Closing channels
- [ ] `range` over channels
- [ ] `select`
- [ ] `sync.WaitGroup`
- [ ] `sync.Mutex`
- [ ] `sync.RWMutex`
- [ ] `sync.Once`
- [ ] `sync.Pool`
- [ ] Context cancellation
- [ ] Worker pools
- [ ] Fan-out
- [ ] Fan-in
- [ ] Race conditions
- [ ] Deadlocks

## Testing

- [ ] Unit tests
- [ ] Table-driven tests
- [ ] Subtests
- [ ] Benchmarks
- [ ] Example tests
- [ ] HTTP handler tests
- [ ] Integration tests
- [ ] Test concurrent code
- [ ] Race detector

## Tooling

- [ ] `go run`
- [ ] `go build`
- [ ] `go test`
- [ ] `go test -race`
- [ ] `go test -bench`
- [ ] `go fmt`
- [ ] `go vet`
- [ ] `go mod`
- [ ] `go doc`
- [ ] `go env`
- [ ] Understand `GOPATH`
- [ ] Understand Go modules

---

# Recommended Order

Follow this order instead of randomly choosing projects:

```text
1. CLI Task Manager
        ↓
2. Log Analyzer
        ↓
3. Concurrent URL Checker
        ↓
4. HTTP Server
        ↓
5. URL Shortener
        ↓
6. Job Queue
        ↓
7. TCP Chat Server
        ↓
8. Web Crawler
        ↓
9. Reverse Proxy
        ↓
10. Git-Like Tool
        ↓
11. Distributed Job System
```

---

# Skill Progression

## Stage 1 — Core Go

- [ ] Structs
- [ ] Methods
- [ ] Slices
- [ ] Maps
- [ ] Interfaces
- [ ] Errors
- [ ] Packages
- [ ] File I/O
- [ ] JSON

## Stage 2 — Practical Go

- [ ] CLI applications
- [ ] HTTP APIs
- [ ] Database applications
- [ ] Testing
- [ ] Configuration
- [ ] Logging

## Stage 3 — Go Concurrency

- [ ] Goroutines
- [ ] Channels
- [ ] Select
- [ ] Mutexes
- [ ] Worker pools
- [ ] Context
- [ ] Cancellation
- [ ] Race conditions
- [ ] Graceful shutdown

## Stage 4 — Go Networking

- [ ] HTTP
- [ ] TCP
- [ ] Connections
- [ ] Timeouts
- [ ] Proxies
- [ ] Load balancing
- [ ] Health checks

## Stage 5 — Production Go

- [ ] Observability
- [ ] Metrics
- [ ] Structured logging
- [ ] Configuration
- [ ] Testing
- [ ] Performance
- [ ] Error handling
- [ ] Deployment
- [ ] Graceful shutdown
- [ ] Failure handling

---

# Project Completion Checklist

Before marking a project complete:

- [ ] It works
- [ ] I can explain every important part
- [ ] I wrote it without copying a tutorial
- [ ] I handled errors
- [ ] I wrote tests
- [ ] I tested edge cases
- [ ] I checked for race conditions where relevant
- [ ] I refactored messy code
- [ ] I added documentation
- [ ] I wrote a README
- [ ] I can explain the architecture
- [ ] I can explain the concurrency model
- [ ] I can explain the major design decisions

---

# Final Goal

By the end, I should be able to look at a Go problem and independently decide:

- [ ] What data structures to use
- [ ] How to structure the packages
- [ ] Where interfaces are useful
- [ ] How errors should flow
- [ ] Whether concurrency is needed
- [ ] Which concurrency primitive to use
- [ ] How to cancel work
- [ ] How to handle failures
- [ ] How to test the system
- [ ] How to make the program observable
- [ ] How to make the program production-ready

> The objective is not to "finish projects."
>
> The objective is to reach the point where **Go becomes a tool you can use to build systems without needing a tutorial to tell you what to do next.**
