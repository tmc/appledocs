// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/ebitengine/purego"
	objc "github.com/ebitengine/purego/objc"
)


// Foundation Functions (8 total)
//
// Type-safe package-level functions with graceful error handling.
// Missing symbols are silently ignored during init; functions will panic when called if unavailable.

var (
	_NSCountFrames func() uint64
	_NSFrameAddress func(uint64) unsafe.Pointer
	_NSIsFreedObject func(objc.ID) bool
	_NSLog func(unsafe.Pointer)
	_NSLogv func(unsafe.Pointer, unsafe.Pointer)
	_NSRecordAllocationEvent func(int, objc.ID)
	_NSReturnAddress func(uint64) unsafe.Pointer
	_NXReadNSObjectFromCoder func(unsafe.Pointer) unsafe.Pointer
)

func init() {
	lib, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
	tryRegister(&_NSCountFrames, lib, "NSCountFrames")
	tryRegister(&_NSFrameAddress, lib, "NSFrameAddress")
	tryRegister(&_NSIsFreedObject, lib, "NSIsFreedObject")
	tryRegister(&_NSLog, lib, "NSLog")
	tryRegister(&_NSLogv, lib, "NSLogv")
	tryRegister(&_NSRecordAllocationEvent, lib, "NSRecordAllocationEvent")
	tryRegister(&_NSReturnAddress, lib, "NSReturnAddress")
	tryRegister(&_NXReadNSObjectFromCoder, lib, "NXReadNSObjectFromCoder")
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



// Returns the number of call frames on the stack.
//
// Added in macOS 10.0.
// Returns the number of call frames on the stack.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCountFrames
func NSCountFrames() uint64 {
	return _NSCountFrames()
}

// Returns the value of the frame pointer of the specified frame.
//
// Added in macOS 10.0.
// Returns the value of the frame pointer of the specified frame.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSFrameAddress
func NSFrameAddress(frame uint64) unsafe.Pointer {
	return _NSFrameAddress(frame)
}

// Returns a Boolean indicating whether the specified object has been freed.
//
// Added in macOS 10.0.
// Returns a Boolean indicating whether the specified object has been freed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSIsFreedObject
func NSIsFreedObject(anObject objc.ID) bool {
	return _NSIsFreedObject(anObject)
}

// Logs an error message to the Apple System Log facility.
//
// Added in macOS 10.0.
// Logs an error message to the Apple System Log facility.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSLog
func NSLog(format unsafe.Pointer) {
	_NSLog(format)
}

// Logs an error message to the Apple System Log facility.
//
// Added in macOS 10.0.
// Logs an error message to the Apple System Log facility.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSLogv(_:_:)
func NSLogv(format unsafe.Pointer, args unsafe.Pointer) {
	_NSLogv(format, args)
}

// Notes an object or zone allocation event and various other statistics, such as the time and current thread.
//
// Added in macOS 10.0.
// Notes an object or zone allocation event and various other statistics, such as the time and current thread.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSRecordAllocationEvent
func NSRecordAllocationEvent(eventType int, object objc.ID) {
	_NSRecordAllocationEvent(eventType, object)
}

// Returns the value of the return address of the specified frame.
//
// Added in macOS 10.0.
// Returns the value of the return address of the specified frame.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSReturnAddress
func NSReturnAddress(frame uint64) unsafe.Pointer {
	return _NSReturnAddress(frame)
}

// Returns the next object from the coder.
//
// Deprecated: This function was deprecated in macOS 10.5.
//
// Added in macOS 10.0.
// Returns the next object from the coder.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NXReadNSObjectFromCoder
func NXReadNSObjectFromCoder(decoder unsafe.Pointer) unsafe.Pointer {
	return _NXReadNSObjectFromCoder(decoder)
}



