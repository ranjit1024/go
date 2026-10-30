# Go Practice Exercises — #1 to #46

These 5 exercises are designed to practice the Go-by-Example topics from the beginning through **Sorting (#46)**, progressively increasing in difficulty.

Completed endpoint:
- #44 Mutexes
- #45 Stateful Goroutines
- #46 Sorting

---

## Exercise 1 — Beginner: Student Grade Analyzer

**Difficulty:** 2/10

Build a program that receives:

```go
scores := []int{72, 91, 65, 88, 54, 79, 95, 61}
```

Implement functions to:

- Find the highest score
- Find the lowest score
- Calculate the average
- Count how many students passed (`>= 60`)
- Print the scores in sorted order

Use:

- Variables and constants
- `for`
- `if/else`
- Functions
- Arrays/slices
- `range`
- Sorting

---

## Exercise 2 — Medium: Word Frequency Analyzer

**Difficulty:** 4/10

Given:

```go
text := "go is simple and go is powerful and go is fast"
```

Create a program that:

1. Splits the sentence into words.
2. Counts each word using a `map[string]int`.
3. Prints the frequency of every word.
4. Finds the most frequently occurring word.
5. Sorts the words by frequency.

Example output:

```text
go        3
is        2
and       2
simple    1
powerful  1
fast      1
```

Create functions such as:

```go
func countWords(text string) map[string]int
func mostFrequent(words map[string]int) (string, int)
```

Use:

- Strings
- Slices
- Maps
- Functions
- Multiple return values
- `range`
- Sorting
- Conditionals

---

## Exercise 3 — Medium/Hard: Bank Account System

**Difficulty:** 6/10

Create:

```go
type Account struct {
    ID      int
    Name    string
    Balance int
}
```

Implement:

```go
Deposit(amount int)
Withdraw(amount int) error
GetBalance() int
```

Then create multiple accounts and perform transactions.

Add validation:

- Cannot deposit negative money.
- Cannot withdraw negative money.
- Cannot withdraw more than the balance.
- Return meaningful errors.

Then create **multiple goroutines** performing transactions on the same account.

You must make the account **thread-safe**.

Use:

- Structs
- Methods
- Pointers
- Errors
- Custom errors
- Goroutines
- Mutex
- WaitGroup

---

## Exercise 4 — Hard: Concurrent Worker Pool

**Difficulty:** 8/10

Create **20 jobs**:

```go
type Job struct {
    ID   int
    Name string
}
```

You should have exactly **3 workers**.

Architecture:

```text
             Jobs
              │
              ▼
       ┌─────────────┐
       │ jobs channel│
       └──────┬──────┘
          ┌───┼───┐
          ▼   ▼   ▼
         W1  W2  W3
```

Each worker:

- Receives a job
- Processes it
- Sleeps for a random amount of time
- Prints the worker ID and job ID

Requirements:

- All 20 jobs must be processed.
- Maximum 3 jobs execute concurrently.
- Main must wait for all workers.
- Close the jobs channel correctly.
- Use `WaitGroup`.
- Use channel directions where appropriate.

For example:

```go
func worker(id int, jobs <-chan Job, wg *sync.WaitGroup)
```

---

## Exercise 5 — Very Hard: Concurrent In-Memory Task Manager

**Difficulty:** 10/10

Build a mini task system.

```go
type Task struct {
    ID     int
    Title  string
    Status string
}
```

Operations:

```text
Create
Get
Complete
Delete
List
```

You should have multiple goroutines performing operations concurrently.

For example:

```text
Goroutine 1 → Create tasks
Goroutine 2 → Complete tasks
Goroutine 3 → Read tasks
Goroutine 4 → Delete tasks
```

### Version 1 — Mutex

First implement the task manager using:

```go
sync.Mutex
```

Make sure concurrent access to the task map is safe.

### Version 2 — Stateful Goroutine

Then implement the same system using the **stateful goroutine pattern**:

```text
              requests
                  │
                  ▼
        ┌─────────────────┐
        │ State Goroutine  │
        │                 │
        │    map[int]Task │
        └─────────────────┘
             ▲    ▲    ▲
             │    │    │
           Get  Create Complete
```

The state-owning goroutine should be the **only goroutine that directly modifies the map**.

### Final requirement

Sort the task output by:

```text
Status
→ then Task ID
```

Use:

- Structs
- Methods
- Pointers
- Interfaces where useful
- Errors
- Goroutines
- Channels
- Channel directions
- Select
- WaitGroups
- Mutex
- Stateful goroutines
- Sorting

---

# Recommended Order

Complete them in this exact order:

```text
1. Student Grade Analyzer
          ↓
2. Word Frequency Analyzer
          ↓
3. Bank Account System
          ↓
4. Concurrent Worker Pool
          ↓
5. Concurrent Task Manager
```

The goal is not just to make the programs work. For every exercise, be able to answer:

```text
1. What state exists?
2. Who owns the state?
3. Which parts can run concurrently?
4. Where is the critical section?
5. How do goroutines communicate?
6. How is shared state protected?
7. How does the program know all work is finished?
8. What happens if two goroutines access the same data simultaneously?
```

## Rule

Try each exercise **without looking at a solution first**.

After completing an exercise, run:

```bash
go run -race .
```

especially for the concurrent exercises.

The progression is intended to take you from basic Go syntax and data structures toward practical concurrency design.
