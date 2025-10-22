// Code generated from Apple documentation for Dispatch. DO NOT EDIT.

package dispatch

import (
	"unsafe"

	"github.com/ebitengine/purego"
)


// Dispatch Functions (39 total)
//
// Type-safe package-level functions with graceful error handling.
// Missing symbols are silently ignored during init; functions will panic when called if unavailable.

var (
	_dispatch_group_wait func(unsafe.Pointer) unsafe.Pointer
	_dispatch_barrier_async func(unsafe.Pointer) unsafe.Pointer
	_dispatch_after func(unsafe.Pointer) unsafe.Pointer
	_dispatch_async func(unsafe.Pointer) unsafe.Pointer
	_dispatch_semaphore_wait func(unsafe.Pointer) unsafe.Pointer
	_dispatch_retain func(unsafe.Pointer) unsafe.Pointer
	_dispatch_release func(unsafe.Pointer) unsafe.Pointer
	_dispatch_group_enter func(unsafe.Pointer) unsafe.Pointer
	_dispatch_group_leave func(unsafe.Pointer) unsafe.Pointer
	_dispatch_activate func(unsafe.Pointer) unsafe.Pointer
	_dispatch_resume func(unsafe.Pointer) unsafe.Pointer
	_dispatch_suspend func(unsafe.Pointer) unsafe.Pointer
	_dispatch_main func() unsafe.Pointer
	_dispatch_after_f func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_dispatch_allow_send_signals func(int) int
	_dispatch_async_f func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_dispatch_barrier_async_and_wait func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_dispatch_block_create_with_qos_class func(unsafe.Pointer, unsafe.Pointer, int, unsafe.Pointer) unsafe.Pointer
	_dispatch_block_perform func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_dispatch_data_create func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_dispatch_data_create_map func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_dispatch_debugv func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_dispatch_get_context func(unsafe.Pointer) unsafe.Pointer
	_dispatch_get_current_queue func() unsafe.Pointer
	_dispatch_get_global_queue func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_dispatch_group_async_f func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_dispatch_group_notify_f func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_dispatch_introspection_hook_queue_item_complete func(unsafe.Pointer) unsafe.Pointer
	_dispatch_io_close func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_dispatch_queue_attr_make_with_autorelease_frequency func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_dispatch_set_context func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_dispatch_set_finalizer_f func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_dispatch_set_qos_class_floor func(unsafe.Pointer, unsafe.Pointer, int) unsafe.Pointer
	_dispatch_source_create func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_dispatch_source_get_handle func(unsafe.Pointer) unsafe.Pointer
	_dispatch_time func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_dispatch_walltime func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_dispatch_workloop_create_inactive func(unsafe.Pointer) unsafe.Pointer
	_dispatch_workloop_set_autorelease_frequency func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
)

func init() {
	lib, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
	tryRegister(&_dispatch_group_wait, lib, "dispatch_group_wait")
	tryRegister(&_dispatch_barrier_async, lib, "dispatch_barrier_async")
	tryRegister(&_dispatch_after, lib, "dispatch_after")
	tryRegister(&_dispatch_async, lib, "dispatch_async")
	tryRegister(&_dispatch_semaphore_wait, lib, "dispatch_semaphore_wait")
	tryRegister(&_dispatch_retain, lib, "dispatch_retain")
	tryRegister(&_dispatch_release, lib, "dispatch_release")
	tryRegister(&_dispatch_group_enter, lib, "dispatch_group_enter")
	tryRegister(&_dispatch_group_leave, lib, "dispatch_group_leave")
	tryRegister(&_dispatch_activate, lib, "dispatch_activate")
	tryRegister(&_dispatch_resume, lib, "dispatch_resume")
	tryRegister(&_dispatch_suspend, lib, "dispatch_suspend")
	tryRegister(&_dispatch_main, lib, "dispatch_main")
	tryRegister(&_dispatch_after_f, lib, "dispatch_after_f")
	tryRegister(&_dispatch_allow_send_signals, lib, "dispatch_allow_send_signals")
	tryRegister(&_dispatch_async_f, lib, "dispatch_async_f")
	tryRegister(&_dispatch_barrier_async_and_wait, lib, "dispatch_barrier_async_and_wait")
	tryRegister(&_dispatch_block_create_with_qos_class, lib, "dispatch_block_create_with_qos_class")
	tryRegister(&_dispatch_block_perform, lib, "dispatch_block_perform")
	tryRegister(&_dispatch_data_create, lib, "dispatch_data_create")
	tryRegister(&_dispatch_data_create_map, lib, "dispatch_data_create_map")
	tryRegister(&_dispatch_debugv, lib, "dispatch_debugv")
	tryRegister(&_dispatch_get_context, lib, "dispatch_get_context")
	tryRegister(&_dispatch_get_current_queue, lib, "dispatch_get_current_queue")
	tryRegister(&_dispatch_get_global_queue, lib, "dispatch_get_global_queue")
	tryRegister(&_dispatch_group_async_f, lib, "dispatch_group_async_f")
	tryRegister(&_dispatch_group_notify_f, lib, "dispatch_group_notify_f")
	tryRegister(&_dispatch_introspection_hook_queue_item_complete, lib, "dispatch_introspection_hook_queue_item_complete")
	tryRegister(&_dispatch_io_close, lib, "dispatch_io_close")
	tryRegister(&_dispatch_queue_attr_make_with_autorelease_frequency, lib, "dispatch_queue_attr_make_with_autorelease_frequency")
	tryRegister(&_dispatch_set_context, lib, "dispatch_set_context")
	tryRegister(&_dispatch_set_finalizer_f, lib, "dispatch_set_finalizer_f")
	tryRegister(&_dispatch_set_qos_class_floor, lib, "dispatch_set_qos_class_floor")
	tryRegister(&_dispatch_source_create, lib, "dispatch_source_create")
	tryRegister(&_dispatch_source_get_handle, lib, "dispatch_source_get_handle")
	tryRegister(&_dispatch_time, lib, "dispatch_time")
	tryRegister(&_dispatch_walltime, lib, "dispatch_walltime")
	tryRegister(&_dispatch_workloop_create_inactive, lib, "dispatch_workloop_create_inactive")
	tryRegister(&_dispatch_workloop_set_autorelease_frequency, lib, "dispatch_workloop_set_autorelease_frequency")
}

// tryRegister attempts to register a function, silently ignoring failures.
// This allows the library to load even if some symbols are missing.
func tryRegister(fn interface{}, lib uintptr, name string) {
	defer func() {
		if r := recover(); r != nil {
			// Symbol not found - function will remain nil and panic when called
			// This is expected for inline functions, macros, or version-specific APIs
		}
	}()
	purego.RegisterLibFunc(fn, lib, name)
}



// Waits synchronously for the previously submitted block objects to finish; returns if the blocks do not complete before the specified timeout period has elapsed. [Full Topic]
//
// Added in macOS 10.6.
//
// [Full Topic]: https://developer.apple.com/documentation/dispatch/1452794-dispatch_group_wait
func dispatch_group_wait(p0 unsafe.Pointer) unsafe.Pointer {
	return _dispatch_group_wait(p0)
	}


// Submits a barrier block for asynchronous execution and returns immediately. [Full Topic]
//
// Added in macOS 10.7.
//
// [Full Topic]: https://developer.apple.com/documentation/dispatch/1452797-dispatch_barrier_async
func dispatch_barrier_async(p0 unsafe.Pointer) {
	_dispatch_barrier_async(p0)
	}


// Enqueues a block for execution at the specified time. [Full Topic]
//
// Added in macOS 10.6.
//
// [Full Topic]: https://developer.apple.com/documentation/dispatch/1452876-dispatch_after
func dispatch_after(p0 unsafe.Pointer) {
	_dispatch_after(p0)
	}


// Submits a block for asynchronous execution on a dispatch queue and returns immediately. [Full Topic]
//
// Added in macOS 10.6.
//
// [Full Topic]: https://developer.apple.com/documentation/dispatch/1453057-dispatch_async
func dispatch_async(p0 unsafe.Pointer) {
	_dispatch_async(p0)
	}


// Waits for (decrements) a semaphore. [Full Topic]
//
// Added in macOS 10.6.
//
// [Full Topic]: https://developer.apple.com/documentation/dispatch/1453087-dispatch_semaphore_wait
func dispatch_semaphore_wait(p0 unsafe.Pointer) unsafe.Pointer {
	return _dispatch_semaphore_wait(p0)
	}


// Increments the reference count (the retain count) of a dispatch object. [Full Topic]
//
// Added in macOS 10.6.
//
// [Full Topic]: https://developer.apple.com/documentation/dispatch/1496306-dispatch_retain
func dispatch_retain(p0 unsafe.Pointer) {
	_dispatch_retain(p0)
	}


// Decrements the reference count (the retain count) of a dispatch object. [Full Topic]
//
// Added in macOS 10.6.
//
// [Full Topic]: https://developer.apple.com/documentation/dispatch/1496328-dispatch_release
func dispatch_release(p0 unsafe.Pointer) {
	_dispatch_release(p0)
	}


// Explicitly indicates that a block has entered the group. [Full Topic]
//
// Added in macOS 10.6.
//
// [Full Topic]: https://developer.apple.com/documentation/Dispatch/DispatchGroup/enter()
func dispatch_group_enter(group unsafe.Pointer) {
	_dispatch_group_enter(group)
	}


// Explicitly indicates that a block in the group finished executing. [Full Topic]
//
// Added in macOS 10.6.
//
// [Full Topic]: https://developer.apple.com/documentation/Dispatch/DispatchGroup/leave()
func dispatch_group_leave(group unsafe.Pointer) {
	_dispatch_group_leave(group)
	}


// Activates the dispatch object. [Full Topic]
//
// Added in macOS 10.12.
//
// [Full Topic]: https://developer.apple.com/documentation/Dispatch/DispatchObject/activate()
func dispatch_activate(object unsafe.Pointer) {
	_dispatch_activate(object)
	}


// Resumes the invocation of block objects on a dispatch object. [Full Topic]
//
// Added in macOS 10.6.
//
// [Full Topic]: https://developer.apple.com/documentation/Dispatch/DispatchObject/resume()
func dispatch_resume(object unsafe.Pointer) {
	_dispatch_resume(object)
	}


// Suspends the invocation of block objects on a dispatch object. [Full Topic]
//
// Added in macOS 10.6.
//
// [Full Topic]: https://developer.apple.com/documentation/Dispatch/DispatchObject/suspend()
func dispatch_suspend(object unsafe.Pointer) {
	_dispatch_suspend(object)
	}


// Executes blocks submitted to the main queue. [Full Topic]
//
// Added in macOS 10.6.
//
// [Full Topic]: https://developer.apple.com/documentation/Dispatch/dispatchMain()
func dispatch_main() {
	_dispatch_main()
	}


// Enqueues an app-defined function for execution at the specified time. [Full Topic]
//
// Added in macOS 10.6.
//
// [Full Topic]: https://developer.apple.com/documentation/Dispatch/dispatch_after_f
func dispatch_after_f(when unsafe.Pointer, queue unsafe.Pointer, context unsafe.Pointer, work unsafe.Pointer) {
	_dispatch_after_f(when, queue, context, work)
	}


// dispatch_allow_send_signals is a Dispatch function. [Full Topic]
//
// Added in macOS 14.4.
//
// [Full Topic]: https://developer.apple.com/documentation/Dispatch/dispatch_allow_send_signals(_:)
func dispatch_allow_send_signals(preserve_signum int) int {
	return _dispatch_allow_send_signals(preserve_signum)
	}


// Submits an app-defined function for asynchronous execution on a dispatch queue and returns immediately. [Full Topic]
//
// Added in macOS 10.6.
//
// [Full Topic]: https://developer.apple.com/documentation/Dispatch/dispatch_async_f
func dispatch_async_f(queue unsafe.Pointer, context unsafe.Pointer, work unsafe.Pointer) {
	_dispatch_async_f(queue, context, work)
	}


// Submits a work item for synchronous execution and marks the work as a barrier for subsequent concurrent tasks. [Full Topic]
//
// Added in macOS 10.14.
//
// [Full Topic]: https://developer.apple.com/documentation/Dispatch/dispatch_barrier_async_and_wait
func dispatch_barrier_async_and_wait(queue unsafe.Pointer, block unsafe.Pointer) {
	_dispatch_barrier_async_and_wait(queue, block)
	}


// Creates a new dispatch block from an existing block and the given flags, and assigns it the specified quality-of-service class and relative priority. [Full Topic]
//
// Added in macOS 10.10.
//
// [Full Topic]: https://developer.apple.com/documentation/Dispatch/dispatch_block_create_with_qos_class
func dispatch_block_create_with_qos_class(flags unsafe.Pointer, qos_class unsafe.Pointer, relative_priority int, block unsafe.Pointer) unsafe.Pointer {
	return _dispatch_block_create_with_qos_class(flags, qos_class, relative_priority, block)
	}


// Creates, synchronously executes, and releases a dispatch block from the specified block and flags. [Full Topic]
//
// Added in macOS 10.10.
//
// [Full Topic]: https://developer.apple.com/documentation/Dispatch/dispatch_block_perform
func dispatch_block_perform(flags unsafe.Pointer, block unsafe.Pointer) {
	_dispatch_block_perform(flags, block)
	}


// Creates a new dispatch data object with the specified memory buffer. [Full Topic]
//
// Added in macOS 10.7.
//
// [Full Topic]: https://developer.apple.com/documentation/Dispatch/dispatch_data_create
func dispatch_data_create(buffer unsafe.Pointer, size unsafe.Pointer, queue unsafe.Pointer, destructor unsafe.Pointer) unsafe.Pointer {
	return _dispatch_data_create(buffer, size, queue, destructor)
	}


// Returns a new dispatch data object containing a contiguous representation of the specified object’s memory. [Full Topic]
//
// Added in macOS 10.7.
//
// [Full Topic]: https://developer.apple.com/documentation/Dispatch/dispatch_data_create_map
func dispatch_data_create_map(data unsafe.Pointer, buffer_ptr unsafe.Pointer, size_ptr unsafe.Pointer) unsafe.Pointer {
	return _dispatch_data_create_map(data, buffer_ptr, size_ptr)
	}


// dispatch_debugv is a Dispatch function. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/Dispatch/dispatch_debugv(_:_:_:)
func dispatch_debugv(object unsafe.Pointer, message unsafe.Pointer, ap unsafe.Pointer) {
	_dispatch_debugv(object, message, ap)
	}


// Returns the application-defined context of an object. [Full Topic]
//
// Added in macOS 10.6.
//
// [Full Topic]: https://developer.apple.com/documentation/Dispatch/dispatch_get_context
func dispatch_get_context(object unsafe.Pointer) unsafe.Pointer {
	return _dispatch_get_context(object)
	}


// Returns the queue on which the currently executing block is running. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/Dispatch/dispatch_get_current_queue()
func dispatch_get_current_queue() unsafe.Pointer {
	return _dispatch_get_current_queue()
	}


// Returns a system-defined global concurrent queue with the specified quality-of-service class. [Full Topic]
//
// Added in macOS 10.6.
//
// [Full Topic]: https://developer.apple.com/documentation/Dispatch/dispatch_get_global_queue
func dispatch_get_global_queue(identifier unsafe.Pointer, flags unsafe.Pointer) unsafe.Pointer {
	return _dispatch_get_global_queue(identifier, flags)
	}


// Submits an application-defined function to a dispatch queue and associates it with the specified dispatch group. [Full Topic]
//
// Added in macOS 10.6.
//
// [Full Topic]: https://developer.apple.com/documentation/Dispatch/dispatch_group_async_f
func dispatch_group_async_f(group unsafe.Pointer, queue unsafe.Pointer, context unsafe.Pointer, work unsafe.Pointer) {
	_dispatch_group_async_f(group, queue, context, work)
	}


// Schedules an application-defined function to be submitted to a queue when a group of previously submitted block objects have completed. [Full Topic]
//
// Added in macOS 10.6.
//
// [Full Topic]: https://developer.apple.com/documentation/Dispatch/dispatch_group_notify_f
func dispatch_group_notify_f(group unsafe.Pointer, queue unsafe.Pointer, context unsafe.Pointer, work unsafe.Pointer) {
	_dispatch_group_notify_f(group, queue, context, work)
	}


// dispatch_introspection_hook_queue_item_complete is a Dispatch function. [Full Topic]
//
// Added in macOS 10.10.
//
// [Full Topic]: https://developer.apple.com/documentation/Dispatch/dispatch_introspection_hook_queue_item_complete
func dispatch_introspection_hook_queue_item_complete(item unsafe.Pointer) {
	_dispatch_introspection_hook_queue_item_complete(item)
	}


// Closes the specified channel to new read and write operations. [Full Topic]
//
// Added in macOS 10.7.
//
// [Full Topic]: https://developer.apple.com/documentation/Dispatch/dispatch_io_close
func dispatch_io_close(channel unsafe.Pointer, flags unsafe.Pointer) {
	_dispatch_io_close(channel, flags)
	}


// Returns an attribute that specifies how the dispatch queue manages autorelease pools for the blocks it executes. [Full Topic]
//
// Added in macOS 10.12.
//
// [Full Topic]: https://developer.apple.com/documentation/Dispatch/dispatch_queue_attr_make_with_autorelease_frequency
func dispatch_queue_attr_make_with_autorelease_frequency(attr unsafe.Pointer, frequency unsafe.Pointer) unsafe.Pointer {
	return _dispatch_queue_attr_make_with_autorelease_frequency(attr, frequency)
	}


// Associates an application-defined context with the object. [Full Topic]
//
// Added in macOS 10.6.
//
// [Full Topic]: https://developer.apple.com/documentation/Dispatch/dispatch_set_context
func dispatch_set_context(object unsafe.Pointer, context unsafe.Pointer) {
	_dispatch_set_context(object, context)
	}


// Sets the finalizer function for a dispatch object. [Full Topic]
//
// Added in macOS 10.6.
//
// [Full Topic]: https://developer.apple.com/documentation/Dispatch/dispatch_set_finalizer_f
func dispatch_set_finalizer_f(object unsafe.Pointer, finalizer unsafe.Pointer) {
	_dispatch_set_finalizer_f(object, finalizer)
	}


// Specifies the minimum quality-of-service level for a dispatch queue, source, or workloop. [Full Topic]
//
// Added in macOS 10.14.
//
// [Full Topic]: https://developer.apple.com/documentation/Dispatch/dispatch_set_qos_class_floor
func dispatch_set_qos_class_floor(object unsafe.Pointer, qos_class unsafe.Pointer, relative_priority int) {
	_dispatch_set_qos_class_floor(object, qos_class, relative_priority)
	}


// Creates a new dispatch source to monitor low-level system events. [Full Topic]
//
// Added in macOS 10.6.
//
// [Full Topic]: https://developer.apple.com/documentation/Dispatch/dispatch_source_create
func dispatch_source_create(type_ unsafe.Pointer, handle unsafe.Pointer, mask unsafe.Pointer, queue unsafe.Pointer) unsafe.Pointer {
	return _dispatch_source_create(type_, handle, mask, queue)
	}


// Returns the underlying system handle associated with the specified dispatch source. [Full Topic]
//
// Added in macOS 10.6.
//
// [Full Topic]: https://developer.apple.com/documentation/Dispatch/dispatch_source_get_handle
func dispatch_source_get_handle(source unsafe.Pointer) unsafe.Pointer {
	return _dispatch_source_get_handle(source)
	}


// Creates a relative to the default clock or modifies an existing . [Full Topic]
//
// Added in macOS 10.6.
//
// [Full Topic]: https://developer.apple.com/documentation/Dispatch/dispatch_time
func dispatch_time(when unsafe.Pointer, delta unsafe.Pointer) unsafe.Pointer {
	return _dispatch_time(when, delta)
	}


// Creates a using an absolute time according to the wall clock. [Full Topic]
//
// Added in macOS 10.6.
//
// [Full Topic]: https://developer.apple.com/documentation/Dispatch/dispatch_walltime
func dispatch_walltime(when unsafe.Pointer, delta unsafe.Pointer) unsafe.Pointer {
	return _dispatch_walltime(when, delta)
	}


// Creates a new inactive workloop with the specified label. [Full Topic]
//
// Added in macOS 10.14.
//
// [Full Topic]: https://developer.apple.com/documentation/Dispatch/dispatch_workloop_create_inactive
func dispatch_workloop_create_inactive(label unsafe.Pointer) unsafe.Pointer {
	return _dispatch_workloop_create_inactive(label)
	}


// Configures how the workloop manages the autorelease pools for the blocks it executes. [Full Topic]
//
// Added in macOS 10.14.
//
// [Full Topic]: https://developer.apple.com/documentation/Dispatch/dispatch_workloop_set_autorelease_frequency
func dispatch_workloop_set_autorelease_frequency(workloop unsafe.Pointer, frequency unsafe.Pointer) {
	_dispatch_workloop_set_autorelease_frequency(workloop, frequency)
	}




