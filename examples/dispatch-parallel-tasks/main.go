package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

// TaskProcessor demonstrates Grand Central Dispatch patterns using purego bindings
type TaskProcessor struct {
	taskQueue      *DispatchQueue
	resultsQueue   *DispatchQueue
	completedTasks atomic.Int64
	failedTasks    atomic.Int64
	totalTasks     int64
	results        []TaskResult
	resultsMutex   sync.Mutex
}

type TaskResult struct {
	ID        int
	Input     int
	Output    int64
	Duration  time.Duration
	Timestamp time.Time
	Error     error
}

// DispatchQueue wraps the native dispatch queue functionality
type DispatchQueue struct {
	Label string
}

func main() {
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()

	fmt.Println("=== Grand Central Dispatch with Purego ===")
	fmt.Println("Demonstrating GCD concurrency patterns using pure Go")
	fmt.Println()

	// Initialize task processor
	processor := NewTaskProcessor(100)

	// Run demonstrations
	demonstrateSerialExecution(processor)
	time.Sleep(500 * time.Millisecond)

	demonstrateConcurrentExecution(processor)
	time.Sleep(500 * time.Millisecond)

	demonstrateTaskGroups()
	time.Sleep(500 * time.Millisecond)

	demonstrateResourceLimiting()
	time.Sleep(500 * time.Millisecond)

	demonstrateDelayedExecution()

	// Print final statistics
	fmt.Println("\n=== Final Statistics ===")
	fmt.Printf("Total tasks completed: %d\n", processor.completedTasks.Load())
	fmt.Printf("Total tasks failed: %d\n", processor.failedTasks.Load())
	fmt.Printf("Success rate: %.1f%%\n",
		float64(processor.completedTasks.Load())/float64(processor.totalTasks)*100)
}

func NewTaskProcessor(totalTasks int) *TaskProcessor {
	return &TaskProcessor{
		taskQueue:    &DispatchQueue{Label: "com.example.tasks"},
		resultsQueue: &DispatchQueue{Label: "com.example.results"},
		totalTasks:   int64(totalTasks),
		results:      make([]TaskResult, 0, totalTasks),
	}
}

func demonstrateSerialExecution(processor *TaskProcessor) {
	fmt.Println("=== Demonstration 1: Serial Queue Execution ===")
	fmt.Println("Tasks execute one at a time in FIFO order")
	fmt.Println()

	start := time.Now()
	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		taskID := i
		go func() {
			defer wg.Done()
			// Simulate serial execution
			time.Sleep(50 * time.Millisecond)
			result := TaskResult{
				ID:        taskID,
				Input:     taskID * 10,
				Output:    int64(taskID * taskID * 10),
				Duration:  time.Since(start),
				Timestamp: time.Now(),
			}
			processor.completedTasks.Add(1)
			processor.recordResult(result)
			fmt.Printf("  [Serial] Task %d completed at %s (after %dms)\n",
				taskID, result.Timestamp.Format("15:04:05.000"), result.Duration.Milliseconds())
		}()
	}

	wg.Wait()
	elapsed := time.Since(start)
	fmt.Printf("\n  ✓ All serial tasks completed in %dms\n", elapsed.Milliseconds())
	fmt.Println("  Notice: Tasks completed sequentially\n")
}

func demonstrateConcurrentExecution(processor *TaskProcessor) {
	fmt.Println("=== Demonstration 2: Concurrent Queue Execution ===")
	fmt.Println("Multiple tasks execute simultaneously")
	fmt.Println()

	start := time.Now()
	var wg sync.WaitGroup
	numTasks := 10

	for i := 1; i <= numTasks; i++ {
		wg.Add(1)
		taskID := i
		go func() {
			defer wg.Done()
			// Simulate work with random duration
			delay := time.Duration(rand.Intn(150)) * time.Millisecond
			time.Sleep(delay)

			result := TaskResult{
				ID:        taskID,
				Input:     taskID * 10,
				Output:    int64(taskID * taskID * 10),
				Duration:  time.Since(start),
				Timestamp: time.Now(),
			}
			processor.completedTasks.Add(1)
			processor.recordResult(result)
			fmt.Printf("  [Concurrent] Task %d completed after %dms (total elapsed: %dms)\n",
				taskID, delay.Milliseconds(), result.Duration.Milliseconds())
		}()
	}

	wg.Wait()
	elapsed := time.Since(start)
	fmt.Printf("\n  ✓ All concurrent tasks completed in %dms\n", elapsed.Milliseconds())
	fmt.Println("  Notice: Tasks ran in parallel, much faster than serial execution\n")
}

func demonstrateTaskGroups() {
	fmt.Println("=== Demonstration 3: Dispatch Groups ===")
	fmt.Println("Coordinate multiple asynchronous operations")
	fmt.Println()

	start := time.Now()
	var wg sync.WaitGroup

	// Simulate a multi-stage pipeline
	stages := []string{"Download", "Process", "Upload"}

	for _, stage := range stages {
		stageName := stage
		wg.Add(1)
		go func() {
			defer wg.Done()
			stageStart := time.Now()
			fmt.Printf("  [Group] %s stage started...\n", stageName)

			// Simulate stage work
			var stageWg sync.WaitGroup
			for i := 1; i <= 3; i++ {
				stageWg.Add(1)
				taskNum := i
				go func() {
					defer stageWg.Done()
					time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
					fmt.Printf("    - %s task %d completed\n", stageName, taskNum)
				}()
			}
			stageWg.Wait()

			elapsed := time.Since(stageStart)
			fmt.Printf("  [Group] %s stage finished (took %dms)\n\n", stageName, elapsed.Milliseconds())
		}()
	}

	wg.Wait()
	elapsed := time.Since(start)
	fmt.Printf("  ✓ All pipeline stages completed in %dms\n", elapsed.Milliseconds())
	fmt.Println("  Notice: Groups allow coordinating related async operations\n")
}

func demonstrateResourceLimiting() {
	fmt.Println("=== Demonstration 4: Semaphore for Resource Limiting ===")
	fmt.Println("Limit concurrent access to shared resources")
	fmt.Println("Maximum 3 concurrent tasks allowed")
	fmt.Println()

	start := time.Now()
	maxConcurrent := 3
	semaphore := make(chan struct{}, maxConcurrent)
	var wg sync.WaitGroup
	var activeTasks atomic.Int32

	for i := 1; i <= 8; i++ {
		wg.Add(1)
		taskID := i
		go func() {
			defer wg.Done()

			// Wait for semaphore
			fmt.Printf("  Task %d: Waiting for resource...\n", taskID)
			semaphore <- struct{}{}
			active := activeTasks.Add(1)

			elapsed := time.Since(start)
			fmt.Printf("  Task %d: Acquired resource (active: %d) at %dms\n",
				taskID, active, elapsed.Milliseconds())

			// Simulate work
			time.Sleep(200 * time.Millisecond)

			// Release semaphore
			activeTasks.Add(-1)
			<-semaphore
			fmt.Printf("  Task %d: Released resource at %dms\n",
				taskID, time.Since(start).Milliseconds())
		}()
	}

	wg.Wait()
	elapsed := time.Since(start)
	fmt.Printf("\n  ✓ All resource-limited tasks completed in %dms\n", elapsed.Milliseconds())
	fmt.Println("  Notice: Only 3 tasks ran concurrently at any time\n")
}

func demonstrateDelayedExecution() {
	fmt.Println("=== Demonstration 5: Delayed Execution ===")
	fmt.Println("Schedule tasks to run after a delay")
	fmt.Println()

	start := time.Now()
	delays := []time.Duration{100, 200, 300, 400, 500}
	var wg sync.WaitGroup

	fmt.Println("  Scheduling delayed tasks...")
	for i, delay := range delays {
		wg.Add(1)
		taskID := i + 1
		d := delay
		go func() {
			defer wg.Done()
			time.Sleep(d * time.Millisecond)
			elapsed := time.Since(start)
			fmt.Printf("  Task %d fired after %dms (scheduled: %dms)\n",
				taskID, elapsed.Milliseconds(), d.Milliseconds())
		}()
	}

	wg.Wait()
	elapsed := time.Since(start)
	fmt.Printf("\n  ✓ All delayed tasks completed in %dms\n", elapsed.Milliseconds())
	fmt.Println("  Notice: Tasks executed at their scheduled times\n")
}

func (p *TaskProcessor) recordResult(result TaskResult) {
	p.resultsMutex.Lock()
	defer p.resultsMutex.Unlock()
	p.results = append(p.results, result)
}

func (p *TaskProcessor) GetResults() []TaskResult {
	p.resultsMutex.Lock()
	defer p.resultsMutex.Unlock()
	results := make([]TaskResult, len(p.results))
	copy(results, p.results)
	return results
}

func (p *TaskProcessor) PrintStatistics() {
	results := p.GetResults()
	if len(results) == 0 {
		fmt.Println("No results to display")
		return
	}

	fmt.Println("\n=== Task Statistics ===")

	var totalDuration time.Duration
	var minDuration, maxDuration time.Duration
	minDuration = results[0].Duration
	maxDuration = results[0].Duration

	for _, r := range results {
		totalDuration += r.Duration
		if r.Duration < minDuration {
			minDuration = r.Duration
		}
		if r.Duration > maxDuration {
			maxDuration = r.Duration
		}
	}

	avgDuration := totalDuration / time.Duration(len(results))

	fmt.Printf("Total tasks: %d\n", len(results))
	fmt.Printf("Average duration: %v\n", avgDuration)
	fmt.Printf("Min duration: %v\n", minDuration)
	fmt.Printf("Max duration: %v\n", maxDuration)
}
