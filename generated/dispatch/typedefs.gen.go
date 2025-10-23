// Code generated from Apple documentation for Dispatch. DO NOT EDIT.

package dispatch

// Type aliases and typedefs
// dispatch_data_t - An immutable object representing a contiguous or sparse region of memory.
//
// [Full Topic]: https://developer.apple.com/documentation/Dispatch/dispatch_data_t
// dispatch_data_t has base type: NSObject<OS_dispatch_data> *
type dispatch_data_t uintptr
// dispatch_group_t - A group of block objects submitted to a queue for asynchronous invocation.
//
// [Full Topic]: https://developer.apple.com/documentation/Dispatch/dispatch_group_t
// dispatch_group_t has base type: NSObject<OS_dispatch_group> *
type dispatch_group_t uintptr
// dispatch_io_t - A dispatch I/O channel.
//
// [Full Topic]: https://developer.apple.com/documentation/Dispatch/dispatch_io_t
// dispatch_io_t has base type: NSObject<OS_dispatch_io> *
type dispatch_io_t uintptr
// dispatch_object_t - A dispatch object.
//
// [Full Topic]: https://developer.apple.com/documentation/Dispatch/dispatch_object_t
// dispatch_object_t has base type: NSObject<OS_dispatch_object> *
type dispatch_object_t uintptr
// dispatch_queue_attr_t - Attributes describing the behaviors of a dispatch queue.
//
// [Full Topic]: https://developer.apple.com/documentation/Dispatch/dispatch_queue_attr_t
// dispatch_queue_attr_t has base type: NSObject<OS_dispatch_queue_attr> *
type dispatch_queue_attr_t uintptr
// dispatch_queue_concurrent_t - A dispatch queue that executes tasks concurrently and in any order, respecting any barriers that may be in place.
//
// [Full Topic]: https://developer.apple.com/documentation/Dispatch/dispatch_queue_concurrent_t
// dispatch_queue_concurrent_t has base type: NSObject<OS_dispatch_queue_concurrent> *
type dispatch_queue_concurrent_t uintptr
// dispatch_queue_global_t - A dispatch queue that executes tasks concurrently using threads from the global thread pool.
//
// [Full Topic]: https://developer.apple.com/documentation/Dispatch/dispatch_queue_global_t
// dispatch_queue_global_t has base type: NSObject<OS_dispatch_queue_global> *
type dispatch_queue_global_t uintptr
// dispatch_queue_main_t - A dispatch queue that is bound to the app’s main thread and executes tasks serially on that thread.
//
// [Full Topic]: https://developer.apple.com/documentation/Dispatch/dispatch_queue_main_t
// dispatch_queue_main_t has base type: NSObject<OS_dispatch_queue_main> *
type dispatch_queue_main_t uintptr
// dispatch_queue_serial_executor_t type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Dispatch/dispatch_queue_serial_executor_t
// dispatch_queue_serial_executor_t has base type: NSObject<OS_dispatch_queue_serial_executor> *
type dispatch_queue_serial_executor_t uintptr
// dispatch_queue_serial_t - A dispatch queue that executes tasks serially in first-in, first-out (FIFO) order.
//
// [Full Topic]: https://developer.apple.com/documentation/Dispatch/dispatch_queue_serial_t
// dispatch_queue_serial_t has base type: NSObject<OS_dispatch_queue_serial> *
type dispatch_queue_serial_t uintptr
// dispatch_queue_t - A lightweight object to which your application submits blocks for subsequent execution.
//
// [Full Topic]: https://developer.apple.com/documentation/Dispatch/dispatch_queue_t
// dispatch_queue_t has base type: NSObject<OS_dispatch_queue> *
type dispatch_queue_t uintptr
// dispatch_semaphore_t - A dispatch semaphore object.
//
// [Full Topic]: https://developer.apple.com/documentation/Dispatch/dispatch_semaphore_t
// dispatch_semaphore_t has base type: NSObject<OS_dispatch_semaphore> *
type dispatch_semaphore_t uintptr
// dispatch_source_mach_recv_flags_t - Mach receive-right flags.
//
// [Full Topic]: https://developer.apple.com/documentation/Dispatch/dispatch_source_mach_recv_flags_t
// dispatch_source_mach_recv_flags_t has base type: unsigned long
type dispatch_source_mach_recv_flags_t uintptr
// dispatch_source_t - An object that coordinates the processing of specific low-level system events, such as file-system events, timers, and UNIX signals.
//
// [Full Topic]: https://developer.apple.com/documentation/Dispatch/dispatch_source_t
// dispatch_source_t has base type: NSObject<OS_dispatch_source> *
type dispatch_source_t uintptr
// dispatch_time_t - An abstract representation of time.
//
// [Full Topic]: https://developer.apple.com/documentation/Dispatch/dispatch_time_t
// dispatch_time_t has base type: uint64_t
type dispatch_time_t uintptr
// dispatch_workloop_t - A dispatch queue that prioritizes the execution of tasks based on their quality-of-service level.
//
// [Full Topic]: https://developer.apple.com/documentation/Dispatch/dispatch_workloop_t
// dispatch_workloop_t has base type: NSObject<OS_dispatch_workloop> *
type dispatch_workloop_t uintptr

