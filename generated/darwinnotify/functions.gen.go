// Code generated from Apple documentation for DarwinNotify. DO NOT EDIT.

package darwinnotify

/* debug [functions.gen.go]: Generating 13 functions for DarwinNotify */
import (
	"unsafe"

	"github.com/ebitengine/purego"
)


// DarwinNotify Functions (13 total)
//
// Type-safe package-level functions with graceful error handling.
// Missing symbols are silently ignored during init; functions will panic when called if unavailable.

var (
	_notify_cancel func(int) uint32
	_notify_check func(int, []int) uint32
	_notify_get_state func(int, []uint64) uint32
	_notify_is_valid_token func(int) bool
	_notify_post func(unsafe.Pointer) uint32
	_notify_register_check func(unsafe.Pointer, []int) uint32
	_notify_register_dispatch func(unsafe.Pointer, []int, unsafe.Pointer, unsafe.Pointer) uint32
	_notify_register_file_descriptor func(unsafe.Pointer, []int, int, []int) uint32
	_notify_register_mach_port func(unsafe.Pointer, unsafe.Pointer, int, []int) uint32
	_notify_register_signal func(unsafe.Pointer, int, []int) uint32
	_notify_resume func(int) uint32
	_notify_set_state func(int, uint64) uint32
	_notify_suspend func(int) uint32
)

func init() {
	lib, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
	tryRegister(&_notify_cancel, lib, "notify_cancel")
	tryRegister(&_notify_check, lib, "notify_check")
	tryRegister(&_notify_get_state, lib, "notify_get_state")
	tryRegister(&_notify_is_valid_token, lib, "notify_is_valid_token")
	tryRegister(&_notify_post, lib, "notify_post")
	tryRegister(&_notify_register_check, lib, "notify_register_check")
	tryRegister(&_notify_register_dispatch, lib, "notify_register_dispatch")
	tryRegister(&_notify_register_file_descriptor, lib, "notify_register_file_descriptor")
	tryRegister(&_notify_register_mach_port, lib, "notify_register_mach_port")
	tryRegister(&_notify_register_signal, lib, "notify_register_signal")
	tryRegister(&_notify_resume, lib, "notify_resume")
	tryRegister(&_notify_set_state, lib, "notify_set_state")
	tryRegister(&_notify_suspend, lib, "notify_suspend")
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



// notify_cancel is a DarwinNotify function.
//
// Added in macOS .
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DarwinNotify/notify_cancel
func notify_cancel(token int) uint32 {
	return _notify_cancel(token)
}/* debug [functions.gen.go/function]: notify_cancel */

// notify_check is a DarwinNotify function.
//
// Added in macOS .
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DarwinNotify/notify_check
func notify_check(token int, check []int) uint32 {
	return _notify_check(token, check)
}/* debug [functions.gen.go/function]: notify_check */

// notify_get_state is a DarwinNotify function.
//
// Added in macOS 10.5.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DarwinNotify/notify_get_state
func notify_get_state(token int, state64 []uint64) uint32 {
	return _notify_get_state(token, state64)
}/* debug [functions.gen.go/function]: notify_get_state */

// notify_is_valid_token is a DarwinNotify function.
//
// Added in macOS 10.10.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DarwinNotify/notify_is_valid_token
func notify_is_valid_token(val int) bool {
	return _notify_is_valid_token(val)
}/* debug [functions.gen.go/function]: notify_is_valid_token */

// notify_post is a DarwinNotify function.
//
// Added in macOS .
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DarwinNotify/notify_post
func notify_post(name unsafe.Pointer) uint32 {
	return _notify_post(name)
}/* debug [functions.gen.go/function]: notify_post */

// notify_register_check is a DarwinNotify function.
//
// Added in macOS .
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DarwinNotify/notify_register_check
func notify_register_check(name unsafe.Pointer, out_token []int) uint32 {
	return _notify_register_check(name, out_token)
}/* debug [functions.gen.go/function]: notify_register_check */

// Request notification delivery to a dispatch queue.
//
// Added in macOS 10.6.
// Request notification delivery to a dispatch queue.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DarwinNotify/notify_register_dispatch
func notify_register_dispatch(name unsafe.Pointer, out_token []int, queue unsafe.Pointer, handler unsafe.Pointer) uint32 {
	return _notify_register_dispatch(name, out_token, queue, handler)
}/* debug [functions.gen.go/function]: notify_register_dispatch */

// notify_register_file_descriptor is a DarwinNotify function.
//
// Added in macOS .
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DarwinNotify/notify_register_file_descriptor
func notify_register_file_descriptor(name unsafe.Pointer, notify_fd []int, flags int, out_token []int) uint32 {
	return _notify_register_file_descriptor(name, notify_fd, flags, out_token)
}/* debug [functions.gen.go/function]: notify_register_file_descriptor */

// notify_register_mach_port is a DarwinNotify function.
//
// Added in macOS .
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DarwinNotify/notify_register_mach_port
func notify_register_mach_port(name unsafe.Pointer, notify_port unsafe.Pointer, flags int, out_token []int) uint32 {
	return _notify_register_mach_port(name, notify_port, flags, out_token)
}/* debug [functions.gen.go/function]: notify_register_mach_port */

// notify_register_signal is a DarwinNotify function.
//
// Added in macOS .
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DarwinNotify/notify_register_signal
func notify_register_signal(name unsafe.Pointer, sig int, out_token []int) uint32 {
	return _notify_register_signal(name, sig, out_token)
}/* debug [functions.gen.go/function]: notify_register_signal */

// notify_resume is a DarwinNotify function.
//
// Added in macOS 10.6.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DarwinNotify/notify_resume
func notify_resume(token int) uint32 {
	return _notify_resume(token)
}/* debug [functions.gen.go/function]: notify_resume */

// notify_set_state is a DarwinNotify function.
//
// Added in macOS 10.5.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DarwinNotify/notify_set_state
func notify_set_state(token int, state64 uint64) uint32 {
	return _notify_set_state(token, state64)
}/* debug [functions.gen.go/function]: notify_set_state */

// notify_suspend is a DarwinNotify function.
//
// Added in macOS 10.6.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DarwinNotify/notify_suspend
func notify_suspend(token int) uint32 {
	return _notify_suspend(token)
}/* debug [functions.gen.go/function]: notify_suspend */




