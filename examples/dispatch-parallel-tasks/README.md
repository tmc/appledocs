# Grand Central Dispatch - Parallel Task Processing

A sophisticated demonstration of Grand Central Dispatch (GCD) concurrency patterns using pure Go, implementing patterns commonly found in the darwinkit dispatch example.

## Overview

This example demonstrates key GCD concepts and patterns:

- **Serial Queue Execution** - Tasks execute one at a time in FIFO order
- **Concurrent Queue Execution** - Multiple tasks run simultaneously
- **Dispatch Groups** - Coordinate multiple asynchronous operations
- **Semaphores** - Limit concurrent access to shared resources
- **Delayed Execution** - Schedule tasks to run after a specific delay

## Reference

Inspired by darwinkit's dispatch example at:
`/Users/tmc/go/src/github.com/progrium/darwinkit/macos/_examples/dispatch/main.go`

## What it demonstrates

- Serial vs concurrent execution patterns
- Task coordination using wait groups
- Resource limiting with semaphores (max 3 concurrent tasks)
- Pipeline processing with staged operations
- Delayed/scheduled task execution
- Performance comparison between execution strategies
- Atomic counters for statistics tracking

## Key Features

### 1. Serial Execution
```go
// Tasks execute sequentially, one after another
// Completion order is predictable and deterministic
```

### 2. Concurrent Execution
```go
// Tasks run in parallel on multiple goroutines
// Completion order depends on task duration
// Significantly faster for independent tasks
```

### 3. Dispatch Groups
```go
// Coordinate related async operations
// Wait for all tasks in a group to complete
// Useful for pipeline stages (download → process → upload)
```

### 4. Resource Limiting
```go
// Use semaphores to limit concurrent access
// Only N tasks can run simultaneously
// Prevents resource exhaustion
```

### 5. Delayed Execution
```go
// Schedule tasks to run after a specific delay
// Useful for rate limiting, retries, timeouts
```

## Usage

```bash
go run .
```

## Output

The example produces detailed output showing:
- Task completion order and timing
- Active task counts for resource limiting
- Duration comparisons between serial and concurrent execution
- Pipeline stage coordination
- Final statistics (completed tasks, success rate)

## Statistics

Example run results:
- **Serial execution**: ~50ms for 5 tasks (sequential)
- **Concurrent execution**: ~132ms for 10 tasks (parallel)
- **Resource limiting**: Max 3 tasks run concurrently
- **Pipeline stages**: Download, Process, Upload coordinated with groups

## Implementation Notes

This example uses pure Go concurrency primitives to demonstrate GCD patterns:

- `sync.WaitGroup` - Equivalent to dispatch groups
- `chan struct{}` - Semaphores for resource limiting
- `sync.atomic` - Atomic counters for statistics
- `goroutines` - Concurrent task execution
- `time.Sleep` - Delayed execution

While this implementation uses Go's native concurrency, the patterns and concepts directly translate to Apple's Grand Central Dispatch API.

## Complexity

**MEDIUM** - Demonstrates multiple concurrency patterns, resource management, and performance optimization strategies.
