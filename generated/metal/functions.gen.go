// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"unsafe"

	"github.com/ebitengine/purego"
)


// Metal Functions (6 total)
//
// Type-safe package-level functions with graceful error handling.
// Missing symbols are silently ignored during init; functions will panic when called if unavailable.

var (
	_MTLCopyAllDevices func() unsafe.Pointer
	_MTLCreateSystemDefaultDevice func() unsafe.Pointer
	_MTLIOCompressionContextAppendData func(unsafe.Pointer, unsafe.Pointer, uintptr)
	_MTLIOCompressionContextDefaultChunkSize func() uintptr
	_MTLIOCreateCompressionContext func(unsafe.Pointer, unsafe.Pointer, uintptr) unsafe.Pointer
	_MTLIOFlushAndDestroyCompressionContext func(unsafe.Pointer) unsafe.Pointer
)

func init() {
	lib, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
	tryRegister(&_MTLCopyAllDevices, lib, "MTLCopyAllDevices")
	tryRegister(&_MTLCreateSystemDefaultDevice, lib, "MTLCreateSystemDefaultDevice")
	tryRegister(&_MTLIOCompressionContextAppendData, lib, "MTLIOCompressionContextAppendData")
	tryRegister(&_MTLIOCompressionContextDefaultChunkSize, lib, "MTLIOCompressionContextDefaultChunkSize")
	tryRegister(&_MTLIOCreateCompressionContext, lib, "MTLIOCreateCompressionContext")
	tryRegister(&_MTLIOFlushAndDestroyCompressionContext, lib, "MTLIOFlushAndDestroyCompressionContext")
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



// Returns an array of all the Metal device instances in the system. [Full Topic]
//
// Added in macOS 10.11.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLCopyAllDevices()
func MTLCopyAllDevices() unsafe.Pointer {
	return _MTLCopyAllDevices()
	}


// Returns the device instance Metal selects as the default. [Full Topic]
//
// Added in macOS 10.11.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLCreateSystemDefaultDevice()
func MTLCreateSystemDefaultDevice() unsafe.Pointer {
	return _MTLCreateSystemDefaultDevice()
	}


// Adds data to a compression context. [Full Topic]
//
// Added in macOS 13.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLIOCompressionContextAppendData(_:_:_:)
func MTLIOCompressionContextAppendData(context unsafe.Pointer, data unsafe.Pointer, size uintptr) {
	_MTLIOCompressionContextAppendData(context, data, size)
	}


// Returns a compression chunk size you can use as a default for creating a compression context. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLIOCompressionContextDefaultChunkSize()
func MTLIOCompressionContextDefaultChunkSize() uintptr {
	return _MTLIOCompressionContextDefaultChunkSize()
	}


// Creates a compression context that you use to compress data into a single file. [Full Topic]
//
// Added in macOS 13.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLIOCreateCompressionContext
func MTLIOCreateCompressionContext(path unsafe.Pointer, type_ unsafe.Pointer, chunkSize uintptr) unsafe.Pointer {
	return _MTLIOCreateCompressionContext(path, type_, chunkSize)
	}


// Finishes compressing and saves the file that a compression context represents. [Full Topic]
//
// Added in macOS 13.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLIOFlushAndDestroyCompressionContext(_:)
func MTLIOFlushAndDestroyCompressionContext(context unsafe.Pointer) unsafe.Pointer {
	return _MTLIOFlushAndDestroyCompressionContext(context)
	}




