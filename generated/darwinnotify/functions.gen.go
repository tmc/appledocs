// Code generated from Apple documentation for DarwinNotify. DO NOT EDIT.

package darwinnotify

import (
	"unsafe"

	"github.com/ebitengine/purego"
)


// DarwinNotify Functions (8 total)
//
// Type-safe package-level functions with graceful error handling.
// Missing symbols are silently ignored during init; functions will panic when called if unavailable.

var (
	_notify_check func(int, unsafe.Pointer) unsafe.Pointer
	_notify_get_state func(int, unsafe.Pointer) unsafe.Pointer
	_notify_is_valid_token func(int) bool
	_notify_register_dispatch func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_notify_register_file_descriptor func(unsafe.Pointer, unsafe.Pointer, int, unsafe.Pointer) unsafe.Pointer
	_notify_resume func(int) unsafe.Pointer
	_notify_set_state func(int, unsafe.Pointer) unsafe.Pointer
	_notify_suspend func(int) unsafe.Pointer
)

func init() {
	lib, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
	tryRegister(&_notify_check, lib, "notify_check")
	tryRegister(&_notify_get_state, lib, "notify_get_state")
	tryRegister(&_notify_is_valid_token, lib, "notify_is_valid_token")
	tryRegister(&_notify_register_dispatch, lib, "notify_register_dispatch")
	tryRegister(&_notify_register_file_descriptor, lib, "notify_register_file_descriptor")
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



// notify_check is a DarwinNotify function.

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DarwinNotify/notify_check

func notify_check(token int, check unsafe.Pointer) unsafe.Pointer {
	return _notify_check(token, check)
	}


// notify_get_state is a DarwinNotify function.
//
// Added in macOS 10.5.

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DarwinNotify/notify_get_state

func notify_get_state(token int, state64 unsafe.Pointer) unsafe.Pointer {
	return _notify_get_state(token, state64)
	}


// notify_is_valid_token is a DarwinNotify function.
//
// Added in macOS 10.10.

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DarwinNotify/notify_is_valid_token

func notify_is_valid_token(val int) bool {
	return _notify_is_valid_token(val)
	}


// Request notification delivery to a dispatch queue.
//
// Added in macOS 10.6.

// Request notification delivery to a dispatch queue.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DarwinNotify/notify_register_dispatch

func notify_register_dispatch(name unsafe.Pointer, out_token unsafe.Pointer, queue unsafe.Pointer, handler unsafe.Pointer) unsafe.Pointer {
	return _notify_register_dispatch(name, out_token, queue, handler)
	}


// notify_register_file_descriptor is a DarwinNotify function.

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DarwinNotify/notify_register_file_descriptor

func notify_register_file_descriptor(name unsafe.Pointer, notify_fd unsafe.Pointer, flags int, out_token unsafe.Pointer) unsafe.Pointer {
	return _notify_register_file_descriptor(name, notify_fd, flags, out_token)
	}


// notify_resume is a DarwinNotify function.
//
// Added in macOS 10.6.

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DarwinNotify/notify_resume

func notify_resume(token int) unsafe.Pointer {
	return _notify_resume(token)
	}


// notify_set_state is a DarwinNotify function.
//
// Added in macOS 10.5.

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DarwinNotify/notify_set_state

func notify_set_state(token int, state64 unsafe.Pointer) unsafe.Pointer {
	return _notify_set_state(token, state64)
	}


// notify_suspend is a DarwinNotify function.
//
// Added in macOS 10.6.

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DarwinNotify/notify_suspend

func notify_suspend(token int) unsafe.Pointer {
	return _notify_suspend(token)
	}




