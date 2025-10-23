// Code generated from Apple documentation for DarwinNotify. DO NOT EDIT.

package darwinnotify

import (
	"unsafe"

	"github.com/ebitengine/purego"
)


// DarwinNotify Functions (4 total)
//
// Type-safe package-level functions with graceful error handling.
// Missing symbols are silently ignored during init; functions will panic when called if unavailable.

var (
	_notify_cancel func(int) uint32
	_notify_register_file_descriptor func(unsafe.Pointer, []int, int, []int) uint32
	_notify_resume func(int) uint32
	_notify_set_state func(int, uint64) uint32
)

func init() {
	lib, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
	tryRegister(&_notify_cancel, lib, "notify_cancel")
	tryRegister(&_notify_register_file_descriptor, lib, "notify_register_file_descriptor")
	tryRegister(&_notify_resume, lib, "notify_resume")
	tryRegister(&_notify_set_state, lib, "notify_set_state")
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

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DarwinNotify/notify_cancel
func notify_cancel(token int) uint32 {
	return _notify_cancel(token)
}

// notify_register_file_descriptor is a DarwinNotify function.

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DarwinNotify/notify_register_file_descriptor
func notify_register_file_descriptor(name unsafe.Pointer, notify_fd []int, flags int, out_token []int) uint32 {
	return _notify_register_file_descriptor(name, notify_fd, flags, out_token)
}

// notify_resume is a DarwinNotify function.
//
// Added in macOS 10.6.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DarwinNotify/notify_resume
func notify_resume(token int) uint32 {
	return _notify_resume(token)
}

// notify_set_state is a DarwinNotify function.
//
// Added in macOS 10.5.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DarwinNotify/notify_set_state
func notify_set_state(token int, state64 uint64) uint32 {
	return _notify_set_state(token, state64)
}



