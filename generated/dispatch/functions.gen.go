// Code generated from Apple documentation for Dispatch. DO NOT EDIT.

package dispatch

import (
	"unsafe"

	"github.com/ebitengine/purego"
)


// Dispatch Functions (112 total)
//
// Type-safe package-level functions with graceful error handling.
// Missing symbols are silently ignored during init; functions will panic when called if unavailable.

var (
	_dispatch_after func(unsafe.Pointer)
	_dispatch_async func(unsafe.Pointer)
	_dispatch_group_enter func(unsafe.Pointer)
	_dispatch_group_create func() unsafe.Pointer
	_dispatch_group_leave func(unsafe.Pointer)
	_dispatch_io_barrier func(unsafe.Pointer, unsafe.Pointer)
	_dispatch_io_get_descriptor func(unsafe.Pointer) unsafe.Pointer
	_dispatch_io_set_high_water func(unsafe.Pointer, uintptr)
	_dispatch_io_set_low_water func(unsafe.Pointer, uintptr)
	_dispatch_activate func(unsafe.Pointer)
	_dispatch_resume func(unsafe.Pointer)
	_dispatch_set_target_queue func(unsafe.Pointer, unsafe.Pointer)
	_dispatch_suspend func(unsafe.Pointer)
	_dispatch_async_and_wait func(unsafe.Pointer, unsafe.Pointer)
	_dispatch_sync func(unsafe.Pointer, unsafe.Pointer)
	_dispatch_semaphore_create func(unsafe.Pointer) unsafe.Pointer
	_dispatch_main func()
	_dispatch_after_f func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer)
	_dispatch_allow_send_signals func(int) int
	_dispatch_apply func(uintptr, unsafe.Pointer)
	_dispatch_apply_f func(uintptr, unsafe.Pointer, unsafe.Pointer)
	_dispatch_assert_queue func(unsafe.Pointer)
	_dispatch_assert_queue_barrier func(unsafe.Pointer)
	_dispatch_assert_queue_not func(unsafe.Pointer)
	_dispatch_async_and_wait_f func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer)
	_dispatch_async_f func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer)
	_dispatch_barrier_async func(unsafe.Pointer, unsafe.Pointer)
	_dispatch_barrier_async_and_wait func(unsafe.Pointer, unsafe.Pointer)
	_dispatch_barrier_async_and_wait_f func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer)
	_dispatch_barrier_async_f func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer)
	_dispatch_barrier_sync func(unsafe.Pointer, unsafe.Pointer)
	_dispatch_barrier_sync_f func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer)
	_dispatch_block_cancel func(unsafe.Pointer)
	_dispatch_block_create func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_dispatch_block_create_with_qos_class func(unsafe.Pointer, unsafe.Pointer, int, unsafe.Pointer) unsafe.Pointer
	_dispatch_block_notify func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer)
	_dispatch_block_perform func(unsafe.Pointer, unsafe.Pointer)
	_dispatch_block_testcancel func(unsafe.Pointer) unsafe.Pointer
	_dispatch_block_wait func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_dispatch_data_apply func(unsafe.Pointer, unsafe.Pointer) bool
	_dispatch_data_copy_region func(unsafe.Pointer, uintptr, unsafe.Pointer) unsafe.Pointer
	_dispatch_data_create func(unsafe.Pointer, uintptr, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_dispatch_data_create_concat func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_dispatch_data_create_map func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_dispatch_data_create_subrange func(unsafe.Pointer, uintptr, uintptr) unsafe.Pointer
	_dispatch_data_get_size func(unsafe.Pointer) uintptr
	_dispatch_debug func(unsafe.Pointer, unsafe.Pointer)
	_dispatch_debugv func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer)
	_dispatch_get_context func(unsafe.Pointer) unsafe.Pointer
	_dispatch_get_current_queue func() unsafe.Pointer
	_dispatch_get_global_queue func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_dispatch_get_specific func(unsafe.Pointer) unsafe.Pointer
	_dispatch_group_async func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer)
	_dispatch_group_async_f func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer)
	_dispatch_group_notify func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer)
	_dispatch_group_notify_f func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer)
	_dispatch_group_wait func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_dispatch_introspection_hook_queue_callout_begin func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer)
	_dispatch_introspection_hook_queue_callout_end func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer)
	_dispatch_introspection_hook_queue_create func(unsafe.Pointer)
	_dispatch_introspection_hook_queue_destroy func(unsafe.Pointer)
	_dispatch_introspection_hook_queue_item_complete func(unsafe.Pointer)
	_dispatch_introspection_hook_queue_item_dequeue func(unsafe.Pointer, unsafe.Pointer)
	_dispatch_introspection_hook_queue_item_enqueue func(unsafe.Pointer, unsafe.Pointer)
	_dispatch_io_close func(unsafe.Pointer, unsafe.Pointer)
	_dispatch_io_create func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_dispatch_io_create_with_io func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_dispatch_io_create_with_path func(unsafe.Pointer, unsafe.Pointer, int, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_dispatch_io_read func(unsafe.Pointer, unsafe.Pointer, uintptr, unsafe.Pointer, unsafe.Pointer)
	_dispatch_io_set_interval func(unsafe.Pointer, uint64, unsafe.Pointer)
	_dispatch_io_write func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer)
	_dispatch_once func(unsafe.Pointer, unsafe.Pointer)
	_dispatch_once_f func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer)
	_dispatch_queue_attr_make_initially_inactive func(unsafe.Pointer) unsafe.Pointer
	_dispatch_queue_attr_make_with_autorelease_frequency func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_dispatch_queue_attr_make_with_qos_class func(unsafe.Pointer, unsafe.Pointer, int) unsafe.Pointer
	_dispatch_queue_create func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_dispatch_queue_create_with_target func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_dispatch_queue_get_label func(unsafe.Pointer) unsafe.Pointer
	_dispatch_queue_get_qos_class func(unsafe.Pointer, []int) unsafe.Pointer
	_dispatch_queue_get_specific func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_dispatch_queue_set_specific func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer)
	_dispatch_read func(unsafe.Pointer, uintptr, unsafe.Pointer)
	_dispatch_release func(unsafe.Pointer)
	_dispatch_retain func(unsafe.Pointer)
	_dispatch_semaphore_signal func(unsafe.Pointer) unsafe.Pointer
	_dispatch_semaphore_wait func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_dispatch_set_context func(unsafe.Pointer, unsafe.Pointer)
	_dispatch_set_finalizer_f func(unsafe.Pointer, unsafe.Pointer)
	_dispatch_set_qos_class_floor func(unsafe.Pointer, unsafe.Pointer, int)
	_dispatch_source_cancel func(unsafe.Pointer)
	_dispatch_source_create func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_dispatch_source_get_data func(unsafe.Pointer) unsafe.Pointer
	_dispatch_source_get_handle func(unsafe.Pointer) unsafe.Pointer
	_dispatch_source_get_mask func(unsafe.Pointer) unsafe.Pointer
	_dispatch_source_merge_data func(unsafe.Pointer, unsafe.Pointer)
	_dispatch_source_set_cancel_handler func(unsafe.Pointer, unsafe.Pointer)
	_dispatch_source_set_cancel_handler_f func(unsafe.Pointer, unsafe.Pointer)
	_dispatch_source_set_event_handler func(unsafe.Pointer, unsafe.Pointer)
	_dispatch_source_set_event_handler_f func(unsafe.Pointer, unsafe.Pointer)
	_dispatch_source_set_registration_handler func(unsafe.Pointer, unsafe.Pointer)
	_dispatch_source_set_registration_handler_f func(unsafe.Pointer, unsafe.Pointer)
	_dispatch_source_set_timer func(unsafe.Pointer, unsafe.Pointer, uint64, uint64)
	_dispatch_source_testcancel func(unsafe.Pointer) unsafe.Pointer
	_dispatch_sync_f func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer)
	_dispatch_time func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_dispatch_walltime func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_dispatch_workloop_create func(unsafe.Pointer) unsafe.Pointer
	_dispatch_workloop_create_inactive func(unsafe.Pointer) unsafe.Pointer
	_dispatch_workloop_set_autorelease_frequency func(unsafe.Pointer, unsafe.Pointer)
	_dispatch_workloop_set_os_workgroup func(unsafe.Pointer, unsafe.Pointer)
	_dispatch_write func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer)
)

func init() {
	lib, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
	tryRegister(&_dispatch_after, lib, "dispatch_after")
	tryRegister(&_dispatch_async, lib, "dispatch_async")
	tryRegister(&_dispatch_group_enter, lib, "dispatch_group_enter")
	tryRegister(&_dispatch_group_create, lib, "dispatch_group_create")
	tryRegister(&_dispatch_group_leave, lib, "dispatch_group_leave")
	tryRegister(&_dispatch_io_barrier, lib, "dispatch_io_barrier")
	tryRegister(&_dispatch_io_get_descriptor, lib, "dispatch_io_get_descriptor")
	tryRegister(&_dispatch_io_set_high_water, lib, "dispatch_io_set_high_water")
	tryRegister(&_dispatch_io_set_low_water, lib, "dispatch_io_set_low_water")
	tryRegister(&_dispatch_activate, lib, "dispatch_activate")
	tryRegister(&_dispatch_resume, lib, "dispatch_resume")
	tryRegister(&_dispatch_set_target_queue, lib, "dispatch_set_target_queue")
	tryRegister(&_dispatch_suspend, lib, "dispatch_suspend")
	tryRegister(&_dispatch_async_and_wait, lib, "dispatch_async_and_wait")
	tryRegister(&_dispatch_sync, lib, "dispatch_sync")
	tryRegister(&_dispatch_semaphore_create, lib, "dispatch_semaphore_create")
	tryRegister(&_dispatch_main, lib, "dispatch_main")
	tryRegister(&_dispatch_after_f, lib, "dispatch_after_f")
	tryRegister(&_dispatch_allow_send_signals, lib, "dispatch_allow_send_signals")
	tryRegister(&_dispatch_apply, lib, "dispatch_apply")
	tryRegister(&_dispatch_apply_f, lib, "dispatch_apply_f")
	tryRegister(&_dispatch_assert_queue, lib, "dispatch_assert_queue")
	tryRegister(&_dispatch_assert_queue_barrier, lib, "dispatch_assert_queue_barrier")
	tryRegister(&_dispatch_assert_queue_not, lib, "dispatch_assert_queue_not")
	tryRegister(&_dispatch_async_and_wait_f, lib, "dispatch_async_and_wait_f")
	tryRegister(&_dispatch_async_f, lib, "dispatch_async_f")
	tryRegister(&_dispatch_barrier_async, lib, "dispatch_barrier_async")
	tryRegister(&_dispatch_barrier_async_and_wait, lib, "dispatch_barrier_async_and_wait")
	tryRegister(&_dispatch_barrier_async_and_wait_f, lib, "dispatch_barrier_async_and_wait_f")
	tryRegister(&_dispatch_barrier_async_f, lib, "dispatch_barrier_async_f")
	tryRegister(&_dispatch_barrier_sync, lib, "dispatch_barrier_sync")
	tryRegister(&_dispatch_barrier_sync_f, lib, "dispatch_barrier_sync_f")
	tryRegister(&_dispatch_block_cancel, lib, "dispatch_block_cancel")
	tryRegister(&_dispatch_block_create, lib, "dispatch_block_create")
	tryRegister(&_dispatch_block_create_with_qos_class, lib, "dispatch_block_create_with_qos_class")
	tryRegister(&_dispatch_block_notify, lib, "dispatch_block_notify")
	tryRegister(&_dispatch_block_perform, lib, "dispatch_block_perform")
	tryRegister(&_dispatch_block_testcancel, lib, "dispatch_block_testcancel")
	tryRegister(&_dispatch_block_wait, lib, "dispatch_block_wait")
	tryRegister(&_dispatch_data_apply, lib, "dispatch_data_apply")
	tryRegister(&_dispatch_data_copy_region, lib, "dispatch_data_copy_region")
	tryRegister(&_dispatch_data_create, lib, "dispatch_data_create")
	tryRegister(&_dispatch_data_create_concat, lib, "dispatch_data_create_concat")
	tryRegister(&_dispatch_data_create_map, lib, "dispatch_data_create_map")
	tryRegister(&_dispatch_data_create_subrange, lib, "dispatch_data_create_subrange")
	tryRegister(&_dispatch_data_get_size, lib, "dispatch_data_get_size")
	tryRegister(&_dispatch_debug, lib, "dispatch_debug")
	tryRegister(&_dispatch_debugv, lib, "dispatch_debugv")
	tryRegister(&_dispatch_get_context, lib, "dispatch_get_context")
	tryRegister(&_dispatch_get_current_queue, lib, "dispatch_get_current_queue")
	tryRegister(&_dispatch_get_global_queue, lib, "dispatch_get_global_queue")
	tryRegister(&_dispatch_get_specific, lib, "dispatch_get_specific")
	tryRegister(&_dispatch_group_async, lib, "dispatch_group_async")
	tryRegister(&_dispatch_group_async_f, lib, "dispatch_group_async_f")
	tryRegister(&_dispatch_group_notify, lib, "dispatch_group_notify")
	tryRegister(&_dispatch_group_notify_f, lib, "dispatch_group_notify_f")
	tryRegister(&_dispatch_group_wait, lib, "dispatch_group_wait")
	tryRegister(&_dispatch_introspection_hook_queue_callout_begin, lib, "dispatch_introspection_hook_queue_callout_begin")
	tryRegister(&_dispatch_introspection_hook_queue_callout_end, lib, "dispatch_introspection_hook_queue_callout_end")
	tryRegister(&_dispatch_introspection_hook_queue_create, lib, "dispatch_introspection_hook_queue_create")
	tryRegister(&_dispatch_introspection_hook_queue_destroy, lib, "dispatch_introspection_hook_queue_destroy")
	tryRegister(&_dispatch_introspection_hook_queue_item_complete, lib, "dispatch_introspection_hook_queue_item_complete")
	tryRegister(&_dispatch_introspection_hook_queue_item_dequeue, lib, "dispatch_introspection_hook_queue_item_dequeue")
	tryRegister(&_dispatch_introspection_hook_queue_item_enqueue, lib, "dispatch_introspection_hook_queue_item_enqueue")
	tryRegister(&_dispatch_io_close, lib, "dispatch_io_close")
	tryRegister(&_dispatch_io_create, lib, "dispatch_io_create")
	tryRegister(&_dispatch_io_create_with_io, lib, "dispatch_io_create_with_io")
	tryRegister(&_dispatch_io_create_with_path, lib, "dispatch_io_create_with_path")
	tryRegister(&_dispatch_io_read, lib, "dispatch_io_read")
	tryRegister(&_dispatch_io_set_interval, lib, "dispatch_io_set_interval")
	tryRegister(&_dispatch_io_write, lib, "dispatch_io_write")
	tryRegister(&_dispatch_once, lib, "dispatch_once")
	tryRegister(&_dispatch_once_f, lib, "dispatch_once_f")
	tryRegister(&_dispatch_queue_attr_make_initially_inactive, lib, "dispatch_queue_attr_make_initially_inactive")
	tryRegister(&_dispatch_queue_attr_make_with_autorelease_frequency, lib, "dispatch_queue_attr_make_with_autorelease_frequency")
	tryRegister(&_dispatch_queue_attr_make_with_qos_class, lib, "dispatch_queue_attr_make_with_qos_class")
	tryRegister(&_dispatch_queue_create, lib, "dispatch_queue_create")
	tryRegister(&_dispatch_queue_create_with_target, lib, "dispatch_queue_create_with_target")
	tryRegister(&_dispatch_queue_get_label, lib, "dispatch_queue_get_label")
	tryRegister(&_dispatch_queue_get_qos_class, lib, "dispatch_queue_get_qos_class")
	tryRegister(&_dispatch_queue_get_specific, lib, "dispatch_queue_get_specific")
	tryRegister(&_dispatch_queue_set_specific, lib, "dispatch_queue_set_specific")
	tryRegister(&_dispatch_read, lib, "dispatch_read")
	tryRegister(&_dispatch_release, lib, "dispatch_release")
	tryRegister(&_dispatch_retain, lib, "dispatch_retain")
	tryRegister(&_dispatch_semaphore_signal, lib, "dispatch_semaphore_signal")
	tryRegister(&_dispatch_semaphore_wait, lib, "dispatch_semaphore_wait")
	tryRegister(&_dispatch_set_context, lib, "dispatch_set_context")
	tryRegister(&_dispatch_set_finalizer_f, lib, "dispatch_set_finalizer_f")
	tryRegister(&_dispatch_set_qos_class_floor, lib, "dispatch_set_qos_class_floor")
	tryRegister(&_dispatch_source_cancel, lib, "dispatch_source_cancel")
	tryRegister(&_dispatch_source_create, lib, "dispatch_source_create")
	tryRegister(&_dispatch_source_get_data, lib, "dispatch_source_get_data")
	tryRegister(&_dispatch_source_get_handle, lib, "dispatch_source_get_handle")
	tryRegister(&_dispatch_source_get_mask, lib, "dispatch_source_get_mask")
	tryRegister(&_dispatch_source_merge_data, lib, "dispatch_source_merge_data")
	tryRegister(&_dispatch_source_set_cancel_handler, lib, "dispatch_source_set_cancel_handler")
	tryRegister(&_dispatch_source_set_cancel_handler_f, lib, "dispatch_source_set_cancel_handler_f")
	tryRegister(&_dispatch_source_set_event_handler, lib, "dispatch_source_set_event_handler")
	tryRegister(&_dispatch_source_set_event_handler_f, lib, "dispatch_source_set_event_handler_f")
	tryRegister(&_dispatch_source_set_registration_handler, lib, "dispatch_source_set_registration_handler")
	tryRegister(&_dispatch_source_set_registration_handler_f, lib, "dispatch_source_set_registration_handler_f")
	tryRegister(&_dispatch_source_set_timer, lib, "dispatch_source_set_timer")
	tryRegister(&_dispatch_source_testcancel, lib, "dispatch_source_testcancel")
	tryRegister(&_dispatch_sync_f, lib, "dispatch_sync_f")
	tryRegister(&_dispatch_time, lib, "dispatch_time")
	tryRegister(&_dispatch_walltime, lib, "dispatch_walltime")
	tryRegister(&_dispatch_workloop_create, lib, "dispatch_workloop_create")
	tryRegister(&_dispatch_workloop_create_inactive, lib, "dispatch_workloop_create_inactive")
	tryRegister(&_dispatch_workloop_set_autorelease_frequency, lib, "dispatch_workloop_set_autorelease_frequency")
	tryRegister(&_dispatch_workloop_set_os_workgroup, lib, "dispatch_workloop_set_os_workgroup")
	tryRegister(&_dispatch_write, lib, "dispatch_write")
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



// Enqueues a block for execution at the specified time.
//
// Added in macOS 10.6.
// Enqueues a block for execution at the specified time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/dispatch/1452876-dispatch_after
func dispatch_after(p0 unsafe.Pointer) {
	_dispatch_after(p0)
}

// Submits a block for asynchronous execution on a dispatch queue and returns immediately.
//
// Added in macOS 10.6.
// Submits a block for asynchronous execution on a dispatch queue and returns immediately.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/dispatch/1453057-dispatch_async
func dispatch_async(p0 unsafe.Pointer) {
	_dispatch_async(p0)
}

// Explicitly indicates that a block has entered the group.
//
// Added in macOS 10.6.
// Explicitly indicates that a block has entered the group.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Dispatch/DispatchGroup/enter()
func dispatch_group_enter(group unsafe.Pointer) {
	_dispatch_group_enter(group)
}

// Creates a new group to which you can assign block objects.
//
// Added in macOS 10.6.
// Creates a new group to which you can assign block objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Dispatch/DispatchGroup/init()
func dispatch_group_create() unsafe.Pointer {
	return _dispatch_group_create()
}

// Explicitly indicates that a block in the group finished executing.
//
// Added in macOS 10.6.
// Explicitly indicates that a block in the group finished executing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Dispatch/DispatchGroup/leave()
func dispatch_group_leave(group unsafe.Pointer) {
	_dispatch_group_leave(group)
}

// Schedules a barrier operation on the specified channel.
//
// Added in macOS 10.7.
// Schedules a barrier operation on the specified channel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Dispatch/DispatchIO/barrier(execute:)
func dispatch_io_barrier(channel unsafe.Pointer, barrier unsafe.Pointer) {
	_dispatch_io_barrier(channel, barrier)
}

// Returns the file descriptor associated with the specified channel.
//
// Added in macOS 10.7.
// Returns the file descriptor associated with the specified channel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Dispatch/DispatchIO/fileDescriptor
func dispatch_io_get_descriptor(channel unsafe.Pointer) unsafe.Pointer {
	return _dispatch_io_get_descriptor(channel)
}

// Sets the maximum number of bytes to process before enqueueing a handler block.
//
// Added in macOS 10.7.
// Sets the maximum number of bytes to process before enqueueing a handler block.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Dispatch/DispatchIO/setLimit(highWater:)
func dispatch_io_set_high_water(channel unsafe.Pointer, high_water uintptr) {
	_dispatch_io_set_high_water(channel, high_water)
}

// Sets the minimum number of bytes to process before enqueueing a handler block.
//
// Added in macOS 10.7.
// Sets the minimum number of bytes to process before enqueueing a handler block.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Dispatch/DispatchIO/setLimit(lowWater:)
func dispatch_io_set_low_water(channel unsafe.Pointer, low_water uintptr) {
	_dispatch_io_set_low_water(channel, low_water)
}

// Activates the dispatch object.
//
// Added in macOS 10.12.
// Activates the dispatch object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Dispatch/DispatchObject/activate()
func dispatch_activate(object unsafe.Pointer) {
	_dispatch_activate(object)
}

// Resumes the invocation of block objects on a dispatch object.
//
// Added in macOS 10.6.
// Resumes the invocation of block objects on a dispatch object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Dispatch/DispatchObject/resume()
func dispatch_resume(object unsafe.Pointer) {
	_dispatch_resume(object)
}

// Specifies the dispatch queue on which to perform work associated with the current object.
//
// Added in macOS 10.6.
// Specifies the dispatch queue on which to perform work associated with the current object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Dispatch/DispatchObject/setTarget(queue:)
func dispatch_set_target_queue(object unsafe.Pointer, queue unsafe.Pointer) {
	_dispatch_set_target_queue(object, queue)
}

// Suspends the invocation of block objects on a dispatch object.
//
// Added in macOS 10.6.
// Suspends the invocation of block objects on a dispatch object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Dispatch/DispatchObject/suspend()
func dispatch_suspend(object unsafe.Pointer) {
	_dispatch_suspend(object)
}

// Submits a work item for execution and returns only after it finishes executing.
//
// Added in macOS 10.14.
// Submits a work item for execution and returns only after it finishes executing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Dispatch/DispatchQueue/asyncAndWait(execute:)-1udeu
func dispatch_async_and_wait(queue unsafe.Pointer, block unsafe.Pointer) {
	_dispatch_async_and_wait(queue, block)
}

// Submits a block object for execution and returns after that block finishes executing.
//
// Added in macOS 10.6.
// Submits a block object for execution and returns after that block finishes executing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Dispatch/DispatchQueue/sync(execute:)-3segw
func dispatch_sync(queue unsafe.Pointer, block unsafe.Pointer) {
	_dispatch_sync(queue, block)
}

// Creates new counting semaphore with an initial value.
//
// Added in macOS 10.6.
// Creates new counting semaphore with an initial value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Dispatch/DispatchSemaphore/init(value:)
func dispatch_semaphore_create(value unsafe.Pointer) unsafe.Pointer {
	return _dispatch_semaphore_create(value)
}

// Executes blocks submitted to the main queue.
//
// Added in macOS 10.6.
// Executes blocks submitted to the main queue.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Dispatch/dispatchMain()
func dispatch_main() {
	_dispatch_main()
}

// Enqueues an app-defined function for execution at the specified time.
//
// Added in macOS 10.6.
// Enqueues an app-defined function for execution at the specified time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Dispatch/dispatch_after_f
func dispatch_after_f(when unsafe.Pointer, queue unsafe.Pointer, context unsafe.Pointer, work unsafe.Pointer) {
	_dispatch_after_f(when, queue, context, work)
}

// dispatch_allow_send_signals is a Dispatch function.
//
// Added in macOS 14.4.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Dispatch/dispatch_allow_send_signals(_:)
func dispatch_allow_send_signals(preserve_signum int) int {
	return _dispatch_allow_send_signals(preserve_signum)
}

// Submits a single block to the dispatch queue and causes the block to be executed the specified number of times.
//
// Added in macOS 10.6.
// Submits a single block to the dispatch queue and causes the block to be executed the specified number of times.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Dispatch/dispatch_apply
func dispatch_apply(iterations uintptr, queue unsafe.Pointer) {
	_dispatch_apply(iterations, queue)
}

// Submits a single function to the dispatch queue and causes the function to be executed the specified number of times.
//
// Added in macOS 10.6.
// Submits a single function to the dispatch queue and causes the function to be executed the specified number of times.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Dispatch/dispatch_apply_f
func dispatch_apply_f(iterations uintptr, queue unsafe.Pointer, context unsafe.Pointer) {
	_dispatch_apply_f(iterations, queue, context)
}

// Generates an assertion if the current block is not running on the specified dispatch queue.
//
// Added in macOS 10.12.
// Generates an assertion if the current block is not running on the specified dispatch queue.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Dispatch/dispatch_assert_queue
func dispatch_assert_queue(queue unsafe.Pointer) {
	_dispatch_assert_queue(queue)
}

// Generates an assertion if the current block is not running as a barrier on the specified dispatch queue.
//
// Added in macOS 10.12.
// Generates an assertion if the current block is not running as a barrier on the specified dispatch queue.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Dispatch/dispatch_assert_queue_barrier
func dispatch_assert_queue_barrier(queue unsafe.Pointer) {
	_dispatch_assert_queue_barrier(queue)
}

// Generates an assertion if the current block is executing on the specified dispatch queue.
//
// Added in macOS 10.12.
// Generates an assertion if the current block is executing on the specified dispatch queue.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Dispatch/dispatch_assert_queue_not
func dispatch_assert_queue_not(queue unsafe.Pointer) {
	_dispatch_assert_queue_not(queue)
}

// Submits a function-based work item for execution and returns only after it finishes executing.
//
// Added in macOS 10.14.
// Submits a function-based work item for execution and returns only after it finishes executing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Dispatch/dispatch_async_and_wait_f
func dispatch_async_and_wait_f(queue unsafe.Pointer, context unsafe.Pointer, work unsafe.Pointer) {
	_dispatch_async_and_wait_f(queue, context, work)
}

// Submits an app-defined function for asynchronous execution on a dispatch queue and returns immediately.
//
// Added in macOS 10.6.
// Submits an app-defined function for asynchronous execution on a dispatch queue and returns immediately.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Dispatch/dispatch_async_f
func dispatch_async_f(queue unsafe.Pointer, context unsafe.Pointer, work unsafe.Pointer) {
	_dispatch_async_f(queue, context, work)
}

// Submits a barrier block for asynchronous execution and returns immediately.
//
// Added in macOS 10.7.
// Submits a barrier block for asynchronous execution and returns immediately.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Dispatch/dispatch_barrier_async
func dispatch_barrier_async(queue unsafe.Pointer, block unsafe.Pointer) {
	_dispatch_barrier_async(queue, block)
}

// Submits a work item for synchronous execution and marks the work as a barrier for subsequent concurrent tasks.
//
// Added in macOS 10.14.
// Submits a work item for synchronous execution and marks the work as a barrier for subsequent concurrent tasks.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Dispatch/dispatch_barrier_async_and_wait
func dispatch_barrier_async_and_wait(queue unsafe.Pointer, block unsafe.Pointer) {
	_dispatch_barrier_async_and_wait(queue, block)
}

// Submits a function-based work item for synchronous execution and marks the work as a barrier for subsequent concurrent tasks.
//
// Added in macOS 10.14.
// Submits a function-based work item for synchronous execution and marks the work as a barrier for subsequent concurrent tasks.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Dispatch/dispatch_barrier_async_and_wait_f
func dispatch_barrier_async_and_wait_f(queue unsafe.Pointer, context unsafe.Pointer, work unsafe.Pointer) {
	_dispatch_barrier_async_and_wait_f(queue, context, work)
}

// Submits a barrier function for asynchronous execution and returns immediately.
//
// Added in macOS 10.7.
// Submits a barrier function for asynchronous execution and returns immediately.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Dispatch/dispatch_barrier_async_f
func dispatch_barrier_async_f(queue unsafe.Pointer, context unsafe.Pointer, work unsafe.Pointer) {
	_dispatch_barrier_async_f(queue, context, work)
}

// Submits a barrier block object for execution and waits until that block completes.
//
// Added in macOS 10.7.
// Submits a barrier block object for execution and waits until that block completes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Dispatch/dispatch_barrier_sync
func dispatch_barrier_sync(queue unsafe.Pointer, block unsafe.Pointer) {
	_dispatch_barrier_sync(queue, block)
}

// Submits a barrier function for execution and waits until that function completes.
//
// Added in macOS 10.7.
// Submits a barrier function for execution and waits until that function completes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Dispatch/dispatch_barrier_sync_f
func dispatch_barrier_sync_f(queue unsafe.Pointer, context unsafe.Pointer, work unsafe.Pointer) {
	_dispatch_barrier_sync_f(queue, context, work)
}

// Cancels the specified dispatch block asynchronously.
//
// Added in macOS 10.10.
// Cancels the specified dispatch block asynchronously.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Dispatch/dispatch_block_cancel
func dispatch_block_cancel(block unsafe.Pointer) {
	_dispatch_block_cancel(block)
}

// Creates a new dispatch block on the heap using an existing block and the given flags.
//
// Added in macOS 10.10.
// Creates a new dispatch block on the heap using an existing block and the given flags.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Dispatch/dispatch_block_create
func dispatch_block_create(flags unsafe.Pointer, block unsafe.Pointer) unsafe.Pointer {
	return _dispatch_block_create(flags, block)
}

// Creates a new dispatch block from an existing block and the given flags, and assigns it the specified quality-of-service class and relative priority.
//
// Added in macOS 10.10.
// Creates a new dispatch block from an existing block and the given flags, and assigns it the specified quality-of-service class and relative priority.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Dispatch/dispatch_block_create_with_qos_class
func dispatch_block_create_with_qos_class(flags unsafe.Pointer, qos_class unsafe.Pointer, relative_priority int, block unsafe.Pointer) unsafe.Pointer {
	return _dispatch_block_create_with_qos_class(flags, qos_class, relative_priority, block)
}

// Schedules a notification block to be submitted to a queue when the execution of a specified dispatch block has completed.
//
// Added in macOS 10.10.
// Schedules a notification block to be submitted to a queue when the execution of a specified dispatch block has completed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Dispatch/dispatch_block_notify
func dispatch_block_notify(block unsafe.Pointer, queue unsafe.Pointer, notification_block unsafe.Pointer) {
	_dispatch_block_notify(block, queue, notification_block)
}

// Creates, synchronously executes, and releases a dispatch block from the specified block and flags.
//
// Added in macOS 10.10.
// Creates, synchronously executes, and releases a dispatch block from the specified block and flags.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Dispatch/dispatch_block_perform
func dispatch_block_perform(flags unsafe.Pointer, block unsafe.Pointer) {
	_dispatch_block_perform(flags, block)
}

// Tests whether the given dispatch block has been canceled.
//
// Added in macOS 10.10.
// Tests whether the given dispatch block has been canceled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Dispatch/dispatch_block_testcancel
func dispatch_block_testcancel(block unsafe.Pointer) unsafe.Pointer {
	return _dispatch_block_testcancel(block)
}

// Waits synchronously until execution of the specified dispatch block has completed or until the specified timeout has elapsed.
//
// Added in macOS 10.10.
// Waits synchronously until execution of the specified dispatch block has completed or until the specified timeout has elapsed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Dispatch/dispatch_block_wait
func dispatch_block_wait(block unsafe.Pointer, timeout unsafe.Pointer) unsafe.Pointer {
	return _dispatch_block_wait(block, timeout)
}

// Traverses the memory of a dispatch data object and executes custom code on each region.
//
// Added in macOS 10.7.
// Traverses the memory of a dispatch data object and executes custom code on each region.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Dispatch/dispatch_data_apply
func dispatch_data_apply(data unsafe.Pointer, applier unsafe.Pointer) bool {
	return _dispatch_data_apply(data, applier)
}

// Returns a data object containing a portion of the data in another data object.
//
// Added in macOS 10.7.
// Returns a data object containing a portion of the data in another data object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Dispatch/dispatch_data_copy_region
func dispatch_data_copy_region(data unsafe.Pointer, location uintptr, offset_ptr unsafe.Pointer) unsafe.Pointer {
	return _dispatch_data_copy_region(data, location, offset_ptr)
}

// Creates a new dispatch data object with the specified memory buffer.
//
// Added in macOS 10.7.
// Creates a new dispatch data object with the specified memory buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Dispatch/dispatch_data_create
func dispatch_data_create(buffer unsafe.Pointer, size uintptr, queue unsafe.Pointer, destructor unsafe.Pointer) unsafe.Pointer {
	return _dispatch_data_create(buffer, size, queue, destructor)
}

// Returns a new dispatch data object consisting of the concatenated data from two other data objects.
//
// Added in macOS 10.7.
// Returns a new dispatch data object consisting of the concatenated data from two other data objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Dispatch/dispatch_data_create_concat
func dispatch_data_create_concat(data1 unsafe.Pointer, data2 unsafe.Pointer) unsafe.Pointer {
	return _dispatch_data_create_concat(data1, data2)
}

// Returns a new dispatch data object containing a contiguous representation of the specified object’s memory.
//
// Added in macOS 10.7.
// Returns a new dispatch data object containing a contiguous representation of the specified object’s memory.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Dispatch/dispatch_data_create_map
func dispatch_data_create_map(data unsafe.Pointer, buffer_ptr unsafe.Pointer, size_ptr unsafe.Pointer) unsafe.Pointer {
	return _dispatch_data_create_map(data, buffer_ptr, size_ptr)
}

// Returns a new dispatch data object whose contents consist of a portion of another object’s memory region.
//
// Added in macOS 10.7.
// Returns a new dispatch data object whose contents consist of a portion of another object’s memory region.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Dispatch/dispatch_data_create_subrange
func dispatch_data_create_subrange(data unsafe.Pointer, offset uintptr, length uintptr) unsafe.Pointer {
	return _dispatch_data_create_subrange(data, offset, length)
}

// Returns the logical size of the memory managed by a dispatch data object
//
// Added in macOS 10.7.
// Returns the logical size of the memory managed by a dispatch data object
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Dispatch/dispatch_data_get_size
func dispatch_data_get_size(data unsafe.Pointer) uintptr {
	return _dispatch_data_get_size(data)
}

// Programmatically logs debug information about a dispatch object.
//
// Deprecated: This function was deprecated in macOS 10.9.
//
// Added in macOS 10.6.
// Programmatically logs debug information about a dispatch object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Dispatch/dispatch_debug
func dispatch_debug(object unsafe.Pointer, message unsafe.Pointer) {
	_dispatch_debug(object, message)
}

// dispatch_debugv is a Dispatch function.

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Dispatch/dispatch_debugv(_:_:_:)
func dispatch_debugv(object unsafe.Pointer, message unsafe.Pointer, ap unsafe.Pointer) {
	_dispatch_debugv(object, message, ap)
}

// Returns the application-defined context of an object.
//
// Added in macOS 10.6.
// Returns the application-defined context of an object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Dispatch/dispatch_get_context
func dispatch_get_context(object unsafe.Pointer) unsafe.Pointer {
	return _dispatch_get_context(object)
}

// Returns the queue on which the currently executing block is running.

// Returns the queue on which the currently executing block is running.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Dispatch/dispatch_get_current_queue()
func dispatch_get_current_queue() unsafe.Pointer {
	return _dispatch_get_current_queue()
}

// Returns a system-defined global concurrent queue with the specified quality-of-service class.
//
// Added in macOS 10.6.
// Returns a system-defined global concurrent queue with the specified quality-of-service class.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Dispatch/dispatch_get_global_queue
func dispatch_get_global_queue(identifier unsafe.Pointer, flags unsafe.Pointer) unsafe.Pointer {
	return _dispatch_get_global_queue(identifier, flags)
}

// Returns the value for the key associated with the current dispatch queue.
//
// Added in macOS 10.7.
// Returns the value for the key associated with the current dispatch queue.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Dispatch/dispatch_get_specific
func dispatch_get_specific(key unsafe.Pointer) unsafe.Pointer {
	return _dispatch_get_specific(key)
}

// Schedules a block asynchronously for execution and simultaneously associates it with the specified dispatch group.
//
// Added in macOS 10.6.
// Schedules a block asynchronously for execution and simultaneously associates it with the specified dispatch group.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Dispatch/dispatch_group_async
func dispatch_group_async(group unsafe.Pointer, queue unsafe.Pointer, block unsafe.Pointer) {
	_dispatch_group_async(group, queue, block)
}

// Submits an application-defined function to a dispatch queue and associates it with the specified dispatch group.
//
// Added in macOS 10.6.
// Submits an application-defined function to a dispatch queue and associates it with the specified dispatch group.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Dispatch/dispatch_group_async_f
func dispatch_group_async_f(group unsafe.Pointer, queue unsafe.Pointer, context unsafe.Pointer, work unsafe.Pointer) {
	_dispatch_group_async_f(group, queue, context, work)
}

// Schedules a block object to be submitted to a queue when a group of previously submitted block objects have completed.
//
// Added in macOS 10.6.
// Schedules a block object to be submitted to a queue when a group of previously submitted block objects have completed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Dispatch/dispatch_group_notify
func dispatch_group_notify(group unsafe.Pointer, queue unsafe.Pointer, block unsafe.Pointer) {
	_dispatch_group_notify(group, queue, block)
}

// Schedules an application-defined function to be submitted to a queue when a group of previously submitted block objects have completed.
//
// Added in macOS 10.6.
// Schedules an application-defined function to be submitted to a queue when a group of previously submitted block objects have completed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Dispatch/dispatch_group_notify_f
func dispatch_group_notify_f(group unsafe.Pointer, queue unsafe.Pointer, context unsafe.Pointer, work unsafe.Pointer) {
	_dispatch_group_notify_f(group, queue, context, work)
}

// Waits synchronously for the previously submitted block objects to finish; returns if the blocks do not complete before the specified timeout period has elapsed.
//
// Added in macOS 10.6.
// Waits synchronously for the previously submitted block objects to finish; returns if the blocks do not complete before the specified timeout period has elapsed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Dispatch/dispatch_group_wait
func dispatch_group_wait(group unsafe.Pointer, timeout unsafe.Pointer) unsafe.Pointer {
	return _dispatch_group_wait(group, timeout)
}

// dispatch_introspection_hook_queue_callout_begin is a Dispatch function.
//
// Added in macOS 10.9.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Dispatch/dispatch_introspection_hook_queue_callout_begin
func dispatch_introspection_hook_queue_callout_begin(queue unsafe.Pointer, context unsafe.Pointer, function unsafe.Pointer) {
	_dispatch_introspection_hook_queue_callout_begin(queue, context, function)
}

// dispatch_introspection_hook_queue_callout_end is a Dispatch function.
//
// Added in macOS 10.9.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Dispatch/dispatch_introspection_hook_queue_callout_end
func dispatch_introspection_hook_queue_callout_end(queue unsafe.Pointer, context unsafe.Pointer, function unsafe.Pointer) {
	_dispatch_introspection_hook_queue_callout_end(queue, context, function)
}

// dispatch_introspection_hook_queue_create is a Dispatch function.
//
// Added in macOS 10.9.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Dispatch/dispatch_introspection_hook_queue_create
func dispatch_introspection_hook_queue_create(queue unsafe.Pointer) {
	_dispatch_introspection_hook_queue_create(queue)
}

// dispatch_introspection_hook_queue_destroy is a Dispatch function.
//
// Added in macOS 10.9.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Dispatch/dispatch_introspection_hook_queue_destroy
func dispatch_introspection_hook_queue_destroy(queue unsafe.Pointer) {
	_dispatch_introspection_hook_queue_destroy(queue)
}

// dispatch_introspection_hook_queue_item_complete is a Dispatch function.
//
// Added in macOS 10.10.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Dispatch/dispatch_introspection_hook_queue_item_complete
func dispatch_introspection_hook_queue_item_complete(item unsafe.Pointer) {
	_dispatch_introspection_hook_queue_item_complete(item)
}

// dispatch_introspection_hook_queue_item_dequeue is a Dispatch function.
//
// Added in macOS 10.9.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Dispatch/dispatch_introspection_hook_queue_item_dequeue
func dispatch_introspection_hook_queue_item_dequeue(queue unsafe.Pointer, item unsafe.Pointer) {
	_dispatch_introspection_hook_queue_item_dequeue(queue, item)
}

// dispatch_introspection_hook_queue_item_enqueue is a Dispatch function.
//
// Added in macOS 10.9.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Dispatch/dispatch_introspection_hook_queue_item_enqueue
func dispatch_introspection_hook_queue_item_enqueue(queue unsafe.Pointer, item unsafe.Pointer) {
	_dispatch_introspection_hook_queue_item_enqueue(queue, item)
}

// Closes the specified channel to new read and write operations.
//
// Added in macOS 10.7.
// Closes the specified channel to new read and write operations.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Dispatch/dispatch_io_close
func dispatch_io_close(channel unsafe.Pointer, flags unsafe.Pointer) {
	_dispatch_io_close(channel, flags)
}

// Creates a dispatch I/O channel and associates it with the specified file descriptor.
//
// Added in macOS 10.7.
// Creates a dispatch I/O channel and associates it with the specified file descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Dispatch/dispatch_io_create
func dispatch_io_create(type_ unsafe.Pointer, fd unsafe.Pointer, queue unsafe.Pointer) unsafe.Pointer {
	return _dispatch_io_create(type_, fd, queue)
}

// Creates a new dispatch I/O channel from an existing channel.
//
// Added in macOS 10.7.
// Creates a new dispatch I/O channel from an existing channel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Dispatch/dispatch_io_create_with_io
func dispatch_io_create_with_io(type_ unsafe.Pointer, io unsafe.Pointer, queue unsafe.Pointer) unsafe.Pointer {
	return _dispatch_io_create_with_io(type_, io, queue)
}

// Creates a dispatch I/O channel with the associated path name.
//
// Added in macOS 10.7.
// Creates a dispatch I/O channel with the associated path name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Dispatch/dispatch_io_create_with_path
func dispatch_io_create_with_path(type_ unsafe.Pointer, path unsafe.Pointer, oflag int, mode unsafe.Pointer, queue unsafe.Pointer) unsafe.Pointer {
	return _dispatch_io_create_with_path(type_, path, oflag, mode, queue)
}

// Schedules an asynchronous read operation on the specified channel.
//
// Added in macOS 10.7.
// Schedules an asynchronous read operation on the specified channel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Dispatch/dispatch_io_read
func dispatch_io_read(channel unsafe.Pointer, offset unsafe.Pointer, length uintptr, queue unsafe.Pointer, io_handler unsafe.Pointer) {
	_dispatch_io_read(channel, offset, length, queue, io_handler)
}

// Sets the interval (in nanoseconds) at which to invoke the I/O handlers for the channel.
//
// Added in macOS 10.7.
// Sets the interval (in nanoseconds) at which to invoke the I/O handlers for the channel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Dispatch/dispatch_io_set_interval
func dispatch_io_set_interval(channel unsafe.Pointer, interval uint64, flags unsafe.Pointer) {
	_dispatch_io_set_interval(channel, interval, flags)
}

// Schedules an asynchronous write operation for the specified channel.
//
// Added in macOS 10.7.
// Schedules an asynchronous write operation for the specified channel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Dispatch/dispatch_io_write
func dispatch_io_write(channel unsafe.Pointer, offset unsafe.Pointer, data unsafe.Pointer, queue unsafe.Pointer, io_handler unsafe.Pointer) {
	_dispatch_io_write(channel, offset, data, queue, io_handler)
}

// Executes a block object only once for the lifetime of an application.
//
// Added in macOS 10.6.
// Executes a block object only once for the lifetime of an application.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Dispatch/dispatch_once-c.func
func dispatch_once(predicate unsafe.Pointer, block unsafe.Pointer) {
	_dispatch_once(predicate, block)
}

// Executes an application-defined function only once for the lifetime of an application.
//
// Added in macOS 10.6.
// Executes an application-defined function only once for the lifetime of an application.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Dispatch/dispatch_once_f-c.func
func dispatch_once_f(predicate unsafe.Pointer, context unsafe.Pointer, function unsafe.Pointer) {
	_dispatch_once_f(predicate, context, function)
}

// Returns an attribute that configures a dispatch queue as initially inactive.
//
// Added in macOS 10.12.
// Returns an attribute that configures a dispatch queue as initially inactive.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Dispatch/dispatch_queue_attr_make_initially_inactive
func dispatch_queue_attr_make_initially_inactive(attr unsafe.Pointer) unsafe.Pointer {
	return _dispatch_queue_attr_make_initially_inactive(attr)
}

// Returns an attribute that specifies how the dispatch queue manages autorelease pools for the blocks it executes.
//
// Added in macOS 10.12.
// Returns an attribute that specifies how the dispatch queue manages autorelease pools for the blocks it executes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Dispatch/dispatch_queue_attr_make_with_autorelease_frequency
func dispatch_queue_attr_make_with_autorelease_frequency(attr unsafe.Pointer, frequency unsafe.Pointer) unsafe.Pointer {
	return _dispatch_queue_attr_make_with_autorelease_frequency(attr, frequency)
}

// Returns attributes suitable for creating a dispatch queue with the desired quality-of-service information.
//
// Added in macOS 10.10.
// Returns attributes suitable for creating a dispatch queue with the desired quality-of-service information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Dispatch/dispatch_queue_attr_make_with_qos_class
func dispatch_queue_attr_make_with_qos_class(attr unsafe.Pointer, qos_class unsafe.Pointer, relative_priority int) unsafe.Pointer {
	return _dispatch_queue_attr_make_with_qos_class(attr, qos_class, relative_priority)
}

// Creates a new dispatch queue to which you can submit blocks.
//
// Added in macOS 10.6.
// Creates a new dispatch queue to which you can submit blocks.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Dispatch/dispatch_queue_create
func dispatch_queue_create(label unsafe.Pointer, attr unsafe.Pointer) unsafe.Pointer {
	return _dispatch_queue_create(label, attr)
}

// Creates a new dispatch queue to which you can submit blocks.
//
// Added in macOS 10.12.
// Creates a new dispatch queue to which you can submit blocks.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Dispatch/dispatch_queue_create_with_target
func dispatch_queue_create_with_target(label unsafe.Pointer, attr unsafe.Pointer, target unsafe.Pointer) unsafe.Pointer {
	return _dispatch_queue_create_with_target(label, attr, target)
}

// Returns the label you assigned to the dispatch queue at creation time.
//
// Added in macOS 10.6.
// Returns the label you assigned to the dispatch queue at creation time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Dispatch/dispatch_queue_get_label
func dispatch_queue_get_label(queue unsafe.Pointer) unsafe.Pointer {
	return _dispatch_queue_get_label(queue)
}

// Returns the quality-of-service class for the specified queue.
//
// Added in macOS 10.10.
// Returns the quality-of-service class for the specified queue.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Dispatch/dispatch_queue_get_qos_class
func dispatch_queue_get_qos_class(queue unsafe.Pointer, relative_priority_ptr []int) unsafe.Pointer {
	return _dispatch_queue_get_qos_class(queue, relative_priority_ptr)
}

// Gets the value for the key associated with the specified dispatch queue.
//
// Added in macOS 10.7.
// Gets the value for the key associated with the specified dispatch queue.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Dispatch/dispatch_queue_get_specific
func dispatch_queue_get_specific(queue unsafe.Pointer, key unsafe.Pointer) unsafe.Pointer {
	return _dispatch_queue_get_specific(queue, key)
}

// Sets the key/value data for the specified dispatch queue.
//
// Added in macOS 10.7.
// Sets the key/value data for the specified dispatch queue.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Dispatch/dispatch_queue_set_specific
func dispatch_queue_set_specific(queue unsafe.Pointer, key unsafe.Pointer, context unsafe.Pointer, destructor unsafe.Pointer) {
	_dispatch_queue_set_specific(queue, key, context, destructor)
}

// Schedules an asynchronous read operation using the specified file descriptor.
//
// Added in macOS 10.7.
// Schedules an asynchronous read operation using the specified file descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Dispatch/dispatch_read
func dispatch_read(fd unsafe.Pointer, length uintptr, queue unsafe.Pointer) {
	_dispatch_read(fd, length, queue)
}

// Decrements the reference count (the retain count) of a dispatch object.
//
// Added in macOS 10.6.
// Decrements the reference count (the retain count) of a dispatch object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Dispatch/dispatch_release
func dispatch_release(object unsafe.Pointer) {
	_dispatch_release(object)
}

// Increments the reference count (the retain count) of a dispatch object.
//
// Added in macOS 10.6.
// Increments the reference count (the retain count) of a dispatch object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Dispatch/dispatch_retain
func dispatch_retain(object unsafe.Pointer) {
	_dispatch_retain(object)
}

// Signals (increments) a semaphore.
//
// Added in macOS 10.6.
// Signals (increments) a semaphore.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Dispatch/dispatch_semaphore_signal
func dispatch_semaphore_signal(dsema unsafe.Pointer) unsafe.Pointer {
	return _dispatch_semaphore_signal(dsema)
}

// Waits for (decrements) a semaphore.
//
// Added in macOS 10.6.
// Waits for (decrements) a semaphore.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Dispatch/dispatch_semaphore_wait
func dispatch_semaphore_wait(dsema unsafe.Pointer, timeout unsafe.Pointer) unsafe.Pointer {
	return _dispatch_semaphore_wait(dsema, timeout)
}

// Associates an application-defined context with the object.
//
// Added in macOS 10.6.
// Associates an application-defined context with the object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Dispatch/dispatch_set_context
func dispatch_set_context(object unsafe.Pointer, context unsafe.Pointer) {
	_dispatch_set_context(object, context)
}

// Sets the finalizer function for a dispatch object.
//
// Added in macOS 10.6.
// Sets the finalizer function for a dispatch object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Dispatch/dispatch_set_finalizer_f
func dispatch_set_finalizer_f(object unsafe.Pointer, finalizer unsafe.Pointer) {
	_dispatch_set_finalizer_f(object, finalizer)
}

// Specifies the minimum quality-of-service level for a dispatch queue, source, or workloop.
//
// Added in macOS 10.14.
// Specifies the minimum quality-of-service level for a dispatch queue, source, or workloop.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Dispatch/dispatch_set_qos_class_floor
func dispatch_set_qos_class_floor(object unsafe.Pointer, qos_class unsafe.Pointer, relative_priority int) {
	_dispatch_set_qos_class_floor(object, qos_class, relative_priority)
}

// Asynchronously cancels the dispatch source, preventing any further invocation of its event handler block.
//
// Added in macOS 10.6.
// Asynchronously cancels the dispatch source, preventing any further invocation of its event handler block.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Dispatch/dispatch_source_cancel
func dispatch_source_cancel(source unsafe.Pointer) {
	_dispatch_source_cancel(source)
}

// Creates a new dispatch source to monitor low-level system events.
//
// Added in macOS 10.6.
// Creates a new dispatch source to monitor low-level system events.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Dispatch/dispatch_source_create
func dispatch_source_create(type_ unsafe.Pointer, handle unsafe.Pointer, mask unsafe.Pointer, queue unsafe.Pointer) unsafe.Pointer {
	return _dispatch_source_create(type_, handle, mask, queue)
}

// Returns pending data for the dispatch source.
//
// Added in macOS 10.6.
// Returns pending data for the dispatch source.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Dispatch/dispatch_source_get_data
func dispatch_source_get_data(source unsafe.Pointer) unsafe.Pointer {
	return _dispatch_source_get_data(source)
}

// Returns the underlying system handle associated with the specified dispatch source.
//
// Added in macOS 10.6.
// Returns the underlying system handle associated with the specified dispatch source.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Dispatch/dispatch_source_get_handle
func dispatch_source_get_handle(source unsafe.Pointer) unsafe.Pointer {
	return _dispatch_source_get_handle(source)
}

// Returns the mask of events monitored by the dispatch source.
//
// Added in macOS 10.6.
// Returns the mask of events monitored by the dispatch source.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Dispatch/dispatch_source_get_mask
func dispatch_source_get_mask(source unsafe.Pointer) unsafe.Pointer {
	return _dispatch_source_get_mask(source)
}

// Merges data into a dispatch source and submits its event handler block to its target queue.
//
// Added in macOS 10.6.
// Merges data into a dispatch source and submits its event handler block to its target queue.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Dispatch/dispatch_source_merge_data
func dispatch_source_merge_data(source unsafe.Pointer, value unsafe.Pointer) {
	_dispatch_source_merge_data(source, value)
}

// Sets the cancellation handler block for the given dispatch source.
//
// Added in macOS 10.6.
// Sets the cancellation handler block for the given dispatch source.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Dispatch/dispatch_source_set_cancel_handler
func dispatch_source_set_cancel_handler(source unsafe.Pointer, handler unsafe.Pointer) {
	_dispatch_source_set_cancel_handler(source, handler)
}

// Sets the cancellation handler function for the given dispatch source.
//
// Added in macOS 10.6.
// Sets the cancellation handler function for the given dispatch source.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Dispatch/dispatch_source_set_cancel_handler_f
func dispatch_source_set_cancel_handler_f(source unsafe.Pointer, handler unsafe.Pointer) {
	_dispatch_source_set_cancel_handler_f(source, handler)
}

// Sets the event handler block for the given dispatch source.
//
// Added in macOS 10.6.
// Sets the event handler block for the given dispatch source.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Dispatch/dispatch_source_set_event_handler
func dispatch_source_set_event_handler(source unsafe.Pointer, handler unsafe.Pointer) {
	_dispatch_source_set_event_handler(source, handler)
}

// Sets the event handler function for the given dispatch source.
//
// Added in macOS 10.6.
// Sets the event handler function for the given dispatch source.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Dispatch/dispatch_source_set_event_handler_f
func dispatch_source_set_event_handler_f(source unsafe.Pointer, handler unsafe.Pointer) {
	_dispatch_source_set_event_handler_f(source, handler)
}

// Sets the registration handler block for the given dispatch source.
//
// Added in macOS 10.7.
// Sets the registration handler block for the given dispatch source.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Dispatch/dispatch_source_set_registration_handler
func dispatch_source_set_registration_handler(source unsafe.Pointer, handler unsafe.Pointer) {
	_dispatch_source_set_registration_handler(source, handler)
}

// Sets the registration handler function for the given dispatch source.
//
// Added in macOS 10.7.
// Sets the registration handler function for the given dispatch source.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Dispatch/dispatch_source_set_registration_handler_f
func dispatch_source_set_registration_handler_f(source unsafe.Pointer, handler unsafe.Pointer) {
	_dispatch_source_set_registration_handler_f(source, handler)
}

// Sets a start time, interval, and leeway value for a timer source.
//
// Added in macOS 10.6.
// Sets a start time, interval, and leeway value for a timer source.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Dispatch/dispatch_source_set_timer
func dispatch_source_set_timer(source unsafe.Pointer, start unsafe.Pointer, interval uint64, leeway uint64) {
	_dispatch_source_set_timer(source, start, interval, leeway)
}

// Tests whether the given dispatch source has been canceled.
//
// Added in macOS 10.6.
// Tests whether the given dispatch source has been canceled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Dispatch/dispatch_source_testcancel
func dispatch_source_testcancel(source unsafe.Pointer) unsafe.Pointer {
	return _dispatch_source_testcancel(source)
}

// Submits an app-defined function for synchronous execution on a dispatch queue.
//
// Added in macOS 10.6.
// Submits an app-defined function for synchronous execution on a dispatch queue.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Dispatch/dispatch_sync_f
func dispatch_sync_f(queue unsafe.Pointer, context unsafe.Pointer, work unsafe.Pointer) {
	_dispatch_sync_f(queue, context, work)
}

// Creates a relative to the default clock or modifies an existing .
//
// Added in macOS 10.6.
// Creates a relative to the default clock or modifies an existing .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Dispatch/dispatch_time
func dispatch_time(when unsafe.Pointer, delta unsafe.Pointer) unsafe.Pointer {
	return _dispatch_time(when, delta)
}

// Creates a using an absolute time according to the wall clock.
//
// Added in macOS 10.6.
// Creates a using an absolute time according to the wall clock.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Dispatch/dispatch_walltime
func dispatch_walltime(when unsafe.Pointer, delta unsafe.Pointer) unsafe.Pointer {
	return _dispatch_walltime(when, delta)
}

// Creates a new workloop with the specified label.
//
// Added in macOS 10.14.
// Creates a new workloop with the specified label.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Dispatch/dispatch_workloop_create
func dispatch_workloop_create(label unsafe.Pointer) unsafe.Pointer {
	return _dispatch_workloop_create(label)
}

// Creates a new inactive workloop with the specified label.
//
// Added in macOS 10.14.
// Creates a new inactive workloop with the specified label.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Dispatch/dispatch_workloop_create_inactive
func dispatch_workloop_create_inactive(label unsafe.Pointer) unsafe.Pointer {
	return _dispatch_workloop_create_inactive(label)
}

// Configures how the workloop manages the autorelease pools for the blocks it executes.
//
// Added in macOS 10.14.
// Configures how the workloop manages the autorelease pools for the blocks it executes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Dispatch/dispatch_workloop_set_autorelease_frequency
func dispatch_workloop_set_autorelease_frequency(workloop unsafe.Pointer, frequency unsafe.Pointer) {
	_dispatch_workloop_set_autorelease_frequency(workloop, frequency)
}

// dispatch_workloop_set_os_workgroup is a Dispatch function.
//
// Added in macOS 11.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Dispatch/dispatch_workloop_set_os_workgroup
func dispatch_workloop_set_os_workgroup(workloop unsafe.Pointer, workgroup unsafe.Pointer) {
	_dispatch_workloop_set_os_workgroup(workloop, workgroup)
}

// Schedules an asynchronous write operation using the specified file descriptor.
//
// Added in macOS 10.7.
// Schedules an asynchronous write operation using the specified file descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Dispatch/dispatch_write
func dispatch_write(fd unsafe.Pointer, data unsafe.Pointer, queue unsafe.Pointer) {
	_dispatch_write(fd, data, queue)
}



