// Grand Central Dispatch (GCD) Example using generated bindings
//
// This example demonstrates using libdispatch (GCD) for concurrent programming
// on macOS using only purego bindings without cgo.
//
// Features demonstrated:
// - Dispatch queues (serial, concurrent, main)
// - Async dispatch
// - Dispatch groups for coordination
// - Quality of Service (QoS) classes
package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
	"unsafe"

	"github.com/ebitengine/purego"
)

// Dispatch types
type (
	DispatchQueue      unsafe.Pointer
	DispatchGroup      unsafe.Pointer
	DispatchQueueAttr  uintptr
	DispatchTime       uint64
	DispatchQoSClass   uint32
)

// Dispatch constants
const (
	// QoS classes
	QOS_CLASS_USER_INTERACTIVE DispatchQoSClass = 0x21
	QOS_CLASS_USER_INITIATED   DispatchQoSClass = 0x19
	QOS_CLASS_DEFAULT          DispatchQoSClass = 0x15
	QOS_CLASS_UTILITY          DispatchQoSClass = 0x11
	QOS_CLASS_BACKGROUND       DispatchQoSClass = 0x09

	// Special times
	DISPATCH_TIME_NOW     DispatchTime = 0
	DISPATCH_TIME_FOREVER DispatchTime = ^DispatchTime(0)
)

// Dispatch function bindings
var (
	dispatchLib uintptr

	// Queue creation and management
	dispatch_queue_create            func(label *byte, attr DispatchQueueAttr) DispatchQueue
	dispatch_get_global_queue        func(qos DispatchQoSClass, flags uintptr) DispatchQueue

	// Async execution
	dispatch_async_f                 func(queue DispatchQueue, context unsafe.Pointer, work func(unsafe.Pointer))
	dispatch_sync_f                  func(queue DispatchQueue, context unsafe.Pointer, work func(unsafe.Pointer))

	// Dispatch groups
	dispatch_group_create            func() DispatchGroup
	dispatch_group_async_f           func(group DispatchGroup, queue DispatchQueue, context unsafe.Pointer, work func(unsafe.Pointer))
	dispatch_group_wait              func(group DispatchGroup, timeout DispatchTime) int
	dispatch_group_notify_f          func(group DispatchGroup, queue DispatchQueue, context unsafe.Pointer, work func(unsafe.Pointer))

	// Time functions
	dispatch_time                    func(when DispatchTime, delta int64) DispatchTime
	dispatch_after_f                 func(when DispatchTime, queue DispatchQueue, context unsafe.Pointer, work func(unsafe.Pointer))

	// Utility
	dispatch_release                 func(object unsafe.Pointer)
)

// Queue attribute constants (these are actually pointers to global vars)
var (
	DISPATCH_QUEUE_SERIAL     DispatchQueueAttr
	DISPATCH_QUEUE_CONCURRENT DispatchQueueAttr
)

func init() {
	runtime.LockOSThread()

	var err error
	dispatchLib, err = purego.Dlopen("/usr/lib/system/libdispatch.dylib", purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(fmt.Sprintf("Failed to load libdispatch: %v", err))
	}

	// Register dispatch functions
	purego.RegisterLibFunc(&dispatch_queue_create, dispatchLib, "dispatch_queue_create")
	purego.RegisterLibFunc(&dispatch_get_global_queue, dispatchLib, "dispatch_get_global_queue")
	purego.RegisterLibFunc(&dispatch_async_f, dispatchLib, "dispatch_async_f")
	purego.RegisterLibFunc(&dispatch_sync_f, dispatchLib, "dispatch_sync_f")
	purego.RegisterLibFunc(&dispatch_group_create, dispatchLib, "dispatch_group_create")
	purego.RegisterLibFunc(&dispatch_group_async_f, dispatchLib, "dispatch_group_async_f")
	purego.RegisterLibFunc(&dispatch_group_wait, dispatchLib, "dispatch_group_wait")
	purego.RegisterLibFunc(&dispatch_group_notify_f, dispatchLib, "dispatch_group_notify_f")
	purego.RegisterLibFunc(&dispatch_time, dispatchLib, "dispatch_time")
	purego.RegisterLibFunc(&dispatch_after_f, dispatchLib, "dispatch_after_f")
	purego.RegisterLibFunc(&dispatch_release, dispatchLib, "dispatch_release")

	// Queue attribute constants
	// Note: DISPATCH_QUEUE_CONCURRENT would require loading _dispatch_queue_attr_concurrent symbol
	// For now we just use 0 (serial) - concurrent queues can be created via global queue
	DISPATCH_QUEUE_SERIAL = 0
	DISPATCH_QUEUE_CONCURRENT = 0 // Simplified for now
}

func cString(s string) *byte {
	b := append([]byte(s), 0)
	return &b[0]
}

func main() {
	fmt.Println("=== Grand Central Dispatch (GCD) Example ===\n")

	// Example 1: Serial queue
	fmt.Println("1. Serial Queue Example:")
	serialQueue := dispatch_queue_create(cString("com.example.serial"), DISPATCH_QUEUE_SERIAL)
	var wg1 sync.WaitGroup
	wg1.Add(3)

	for i := 0; i < 3; i++ {
		num := i
		work := func(_ unsafe.Pointer) {
			fmt.Printf("   Serial task %d executing\n", num)
			time.Sleep(100 * time.Millisecond)
			wg1.Done()
		}
		dispatch_async_f(serialQueue, nil, work)
	}
	wg1.Wait()
	fmt.Println()

	// Example 2: Concurrent queue (using global queue)
	fmt.Println("2. Concurrent Queue Example (using global queue):")
	concurrentQueue := dispatch_get_global_queue(QOS_CLASS_DEFAULT, 0)
	var wg2 sync.WaitGroup
	wg2.Add(3)

	for i := 0; i < 3; i++ {
		num := i
		work := func(_ unsafe.Pointer) {
			fmt.Printf("   Concurrent task %d starting\n", num)
			time.Sleep(100 * time.Millisecond)
			fmt.Printf("   Concurrent task %d finished\n", num)
			wg2.Done()
		}
		dispatch_async_f(concurrentQueue, nil, work)
	}
	wg2.Wait()
	fmt.Println()

	// Example 3: Global queue with QoS
	fmt.Println("3. Global Queue (QoS) Example:")
	globalQueue := dispatch_get_global_queue(QOS_CLASS_USER_INITIATED, 0)
	var wg3 sync.WaitGroup
	wg3.Add(1)

	work := func(_ unsafe.Pointer) {
		fmt.Println("   Task on global user-initiated queue")
		wg3.Done()
	}
	dispatch_async_f(globalQueue, nil, work)
	wg3.Wait()
	fmt.Println()

	// Example 4: Dispatch groups
	fmt.Println("4. Dispatch Group Example:")
	group := dispatch_group_create()
	var wg4 sync.WaitGroup
	wg4.Add(1)

	for i := 0; i < 3; i++ {
		num := i
		groupWork := func(_ unsafe.Pointer) {
			fmt.Printf("   Group task %d executing\n", num)
			time.Sleep(50 * time.Millisecond)
		}
		dispatch_group_async_f(group, concurrentQueue, nil, groupWork)
	}

	// Wait for group to complete
	notifyWork := func(_ unsafe.Pointer) {
		fmt.Println("   All group tasks completed!")
		wg4.Done()
	}
	dispatch_group_notify_f(group, globalQueue, nil, notifyWork)
	wg4.Wait()
	fmt.Println()

	// Example 5: Delayed execution
	fmt.Println("5. Delayed Execution Example:")
	var wg5 sync.WaitGroup
	wg5.Add(1)

	delayedWork := func(_ unsafe.Pointer) {
		fmt.Println("   Delayed task executed after 500ms")
		wg5.Done()
	}

	// Schedule task 500ms from now
	when := dispatch_time(DISPATCH_TIME_NOW, 500*1000*1000) // 500ms in nanoseconds
	dispatch_after_f(when, globalQueue, nil, delayedWork)
	wg5.Wait()
	fmt.Println()

	fmt.Println("✅ GCD example completed successfully!")
	fmt.Println("   All dispatch operations demonstrated:")
	fmt.Println("   - Serial queues")
	fmt.Println("   - Concurrent queues")
	fmt.Println("   - Global QoS queues")
	fmt.Println("   - Dispatch groups")
	fmt.Println("   - Delayed execution")
}
