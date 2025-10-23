// Code generated from Apple documentation for DiskArbitration. DO NOT EDIT.

package diskarbitration

import (
	"unsafe"

	"github.com/ebitengine/purego"
)


// DiskArbitration Functions (6 total)
//
// Type-safe package-level functions with graceful error handling.
// Missing symbols are silently ignored during init; functions will panic when called if unavailable.

var (
	_DAApprovalSessionGetTypeID func() unsafe.Pointer
	_DADiskCopyWholeDisk func(unsafe.Pointer) unsafe.Pointer
	_DADiskCreateFromIOMedia func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_DADiskCreateFromVolumePath func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_DADissenterCreate func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_DASessionScheduleWithRunLoop func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer)
)

func init() {
	lib, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
	tryRegister(&_DAApprovalSessionGetTypeID, lib, "DAApprovalSessionGetTypeID")
	tryRegister(&_DADiskCopyWholeDisk, lib, "DADiskCopyWholeDisk")
	tryRegister(&_DADiskCreateFromIOMedia, lib, "DADiskCreateFromIOMedia")
	tryRegister(&_DADiskCreateFromVolumePath, lib, "DADiskCreateFromVolumePath")
	tryRegister(&_DADissenterCreate, lib, "DADissenterCreate")
	tryRegister(&_DASessionScheduleWithRunLoop, lib, "DASessionScheduleWithRunLoop")
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



// DAApprovalSessionGetTypeID is a DiskArbitration function.
//
// Added in macOS 10.4.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DiskArbitration/DAApprovalSessionGetTypeID
func DAApprovalSessionGetTypeID() unsafe.Pointer {
	return _DAApprovalSessionGetTypeID()
}

// Obtain the associated whole disk object for the specified disk.
//
// Added in macOS 10.4.
// Obtain the associated whole disk object for the specified disk.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DiskArbitration/DADiskCopyWholeDisk(_:)
func DADiskCopyWholeDisk(disk unsafe.Pointer) unsafe.Pointer {
	return _DADiskCopyWholeDisk(disk)
}

// Creates a new disk object.
//
// Added in macOS 10.4.
// Creates a new disk object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DiskArbitration/DADiskCreateFromIOMedia(_:_:_:)
func DADiskCreateFromIOMedia(allocator unsafe.Pointer, session unsafe.Pointer, media unsafe.Pointer) unsafe.Pointer {
	return _DADiskCreateFromIOMedia(allocator, session, media)
}

// Creates a new disk object.
//
// Added in macOS 10.7.
// Creates a new disk object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DiskArbitration/DADiskCreateFromVolumePath(_:_:_:)
func DADiskCreateFromVolumePath(allocator unsafe.Pointer, session unsafe.Pointer, path unsafe.Pointer) unsafe.Pointer {
	return _DADiskCreateFromVolumePath(allocator, session, path)
}

// Creates a new dissenter object.
//
// Added in macOS 10.4.
// Creates a new dissenter object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DiskArbitration/DADissenterCreate(_:_:_:)
func DADissenterCreate(allocator unsafe.Pointer, status unsafe.Pointer, string_ unsafe.Pointer) unsafe.Pointer {
	return _DADissenterCreate(allocator, status, string_)
}

// Schedules the session on a run loop.
//
// Added in macOS 10.4.
// Schedules the session on a run loop.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DiskArbitration/DASessionScheduleWithRunLoop(_:_:_:)
func DASessionScheduleWithRunLoop(session unsafe.Pointer, runLoop unsafe.Pointer, runLoopMode unsafe.Pointer) {
	_DASessionScheduleWithRunLoop(session, runLoop, runLoopMode)
}



