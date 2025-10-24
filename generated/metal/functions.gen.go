// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

/* debug [functions.gen.go]: Generating 8 functions for Metal */
import (
	"unsafe"

	"github.com/ebitengine/purego"
)


// Metal Functions (8 total)
//
// Type-safe package-level functions with graceful error handling.
// Missing symbols are silently ignored during init; functions will panic when called if unavailable.

var (
	_MTLCopyAllDevices func() []unsafe.Pointer
	_MTLCopyAllDevicesWithObserver func(unsafe.Pointer, DeviceNotificationHandler) []unsafe.Pointer
	_MTLCreateSystemDefaultDevice func() unsafe.Pointer
	_MTLIOCompressionContextAppendData func(IOCompressionContext, unsafe.Pointer, uintptr)
	_MTLIOCompressionContextDefaultChunkSize func() uintptr
	_MTLIOCreateCompressionContext func(unsafe.Pointer, IOCompressionMethod, uintptr) IOCompressionContext
	_MTLIOFlushAndDestroyCompressionContext func(IOCompressionContext) IOCompressionStatus
	_MTLRemoveDeviceObserver func(unsafe.Pointer)
)

func init() {
	lib, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
	tryRegister(&_MTLCopyAllDevices, lib, "MTLCopyAllDevices")
	tryRegister(&_MTLCopyAllDevicesWithObserver, lib, "MTLCopyAllDevicesWithObserver")
	tryRegister(&_MTLCreateSystemDefaultDevice, lib, "MTLCreateSystemDefaultDevice")
	tryRegister(&_MTLIOCompressionContextAppendData, lib, "MTLIOCompressionContextAppendData")
	tryRegister(&_MTLIOCompressionContextDefaultChunkSize, lib, "MTLIOCompressionContextDefaultChunkSize")
	tryRegister(&_MTLIOCreateCompressionContext, lib, "MTLIOCreateCompressionContext")
	tryRegister(&_MTLIOFlushAndDestroyCompressionContext, lib, "MTLIOFlushAndDestroyCompressionContext")
	tryRegister(&_MTLRemoveDeviceObserver, lib, "MTLRemoveDeviceObserver")
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



// Returns an array of all the Metal device instances in the system.
//
// Added in macOS 10.11.
// Returns an array of all the Metal device instances in the system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLCopyAllDevices()
func MTLCopyAllDevices() []unsafe.Pointer {
	return _MTLCopyAllDevices()
}/* debug [functions.gen.go/function]: MTLCopyAllDevices */

// Returns an array of all the Metal GPU devices in the system and registers a notification handler that Metal calls when the device list changes.
//
// Added in macOS 10.13.
// Returns an array of all the Metal GPU devices in the system and registers a notification handler that Metal calls when the device list changes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLCopyAllDevicesWithObserver
func MTLCopyAllDevicesWithObserver(observer unsafe.Pointer, handler DeviceNotificationHandler) []unsafe.Pointer {
	return _MTLCopyAllDevicesWithObserver(observer, handler)
}/* debug [functions.gen.go/function]: MTLCopyAllDevicesWithObserver */

// Returns the device instance Metal selects as the default.
//
// Added in macOS 10.11.
// Returns the device instance Metal selects as the default.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLCreateSystemDefaultDevice()
func MTLCreateSystemDefaultDevice() unsafe.Pointer {
	return _MTLCreateSystemDefaultDevice()
}/* debug [functions.gen.go/function]: MTLCreateSystemDefaultDevice */

// Adds data to a compression context.
//
// Added in macOS 13.0.
// Adds data to a compression context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLIOCompressionContextAppendData(_:_:_:)
func MTLIOCompressionContextAppendData(context IOCompressionContext, data unsafe.Pointer, size uintptr) {
	_MTLIOCompressionContextAppendData(context, data, size)
}/* debug [functions.gen.go/function]: MTLIOCompressionContextAppendData */

// Returns a compression chunk size you can use as a default for creating a compression context.
//
// Added in macOS .
// Returns a compression chunk size you can use as a default for creating a compression context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLIOCompressionContextDefaultChunkSize()
func MTLIOCompressionContextDefaultChunkSize() uintptr {
	return _MTLIOCompressionContextDefaultChunkSize()
}/* debug [functions.gen.go/function]: MTLIOCompressionContextDefaultChunkSize */

// Creates a compression context that you use to compress data into a single file.
//
// Added in macOS 13.0.
// Creates a compression context that you use to compress data into a single file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLIOCreateCompressionContext
func MTLIOCreateCompressionContext(path unsafe.Pointer, type_ IOCompressionMethod, chunkSize uintptr) IOCompressionContext {
	return _MTLIOCreateCompressionContext(path, type_, chunkSize)
}/* debug [functions.gen.go/function]: MTLIOCreateCompressionContext */

// Finishes compressing and saves the file that a compression context represents.
//
// Added in macOS 13.0.
// Finishes compressing and saves the file that a compression context represents.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLIOFlushAndDestroyCompressionContext(_:)
func MTLIOFlushAndDestroyCompressionContext(context IOCompressionContext) IOCompressionStatus {
	return _MTLIOFlushAndDestroyCompressionContext(context)
}/* debug [functions.gen.go/function]: MTLIOFlushAndDestroyCompressionContext */

// Removes a registered observer of device notifications.
//
// Added in macOS 10.13.
// Removes a registered observer of device notifications.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLRemoveDeviceObserver(_:)
func MTLRemoveDeviceObserver(observer unsafe.Pointer) {
	_MTLRemoveDeviceObserver(observer)
}/* debug [functions.gen.go/function]: MTLRemoveDeviceObserver */




