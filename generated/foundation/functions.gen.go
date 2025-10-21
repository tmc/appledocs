// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/ebitengine/purego"
	coregraphics "github.com/tmc/appledocs/generated/coregraphics"
)


// Foundation Functions (35 total)
//
// Type-safe package-level functions with graceful error handling.
// Missing symbols are silently ignored during init; functions will panic when called if unavailable.

var (
	_NSAllocateCollectable func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_NSAllocateMemoryPages func(unsafe.Pointer) unsafe.Pointer
	_NSClassFromString func(unsafe.Pointer) unsafe.Pointer
	_NSCopyMemoryPages func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_NSCopyObject func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_NSCountFrames func() unsafe.Pointer
	_NSDecimalCompact func(unsafe.Pointer) unsafe.Pointer
	_NSDecimalPower func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_NSExtraRefCount func(unsafe.Pointer) unsafe.Pointer
	_NSFileTypeForHFSTypeCode func(unsafe.Pointer) unsafe.Pointer
	_NSFrameAddress func(unsafe.Pointer) unsafe.Pointer
	_NSFullUserName func() unsafe.Pointer
	_NSGetUncaughtExceptionHandler func() unsafe.Pointer
	_NSHFSTypeCodeFromFileType func(unsafe.Pointer) unsafe.Pointer
	_NSHFSTypeOfFile func(unsafe.Pointer) unsafe.Pointer
	_NSHomeDirectory func() unsafe.Pointer
	_NSHomeDirectoryForUser func(unsafe.Pointer) unsafe.Pointer
	_NSIncrementExtraRefCount func(unsafe.Pointer) unsafe.Pointer
	_NSIntegralRectWithOptions func(coregraphics.CGRect, unsafe.Pointer) coregraphics.CGRect
	_NSIsFreedObject func(unsafe.Pointer) unsafe.Pointer
	_NSLog func(unsafe.Pointer) unsafe.Pointer
	_NSLogv func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_NSMouseInRect func(coregraphics.CGPoint, coregraphics.CGRect, unsafe.Pointer) unsafe.Pointer
	_NSOpenStepRootDirectory func() unsafe.Pointer
	_NSReallocateCollectable func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_NSRecordAllocationEvent func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_NSReturnAddress func(unsafe.Pointer) unsafe.Pointer
	_NSRoundUpToMultipleOfPageSize func(unsafe.Pointer) unsafe.Pointer
	_NSSearchPathForDirectoriesInDomains func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_NSSetUncaughtExceptionHandler func() unsafe.Pointer
	_NSSizeFromString func(unsafe.Pointer) coregraphics.CGSize
	_NSStringFromProtocol func(unsafe.Pointer) unsafe.Pointer
	_NSTemporaryDirectory func() unsafe.Pointer
	_NSUserName func() unsafe.Pointer
	_NXReadNSObjectFromCoder func(unsafe.Pointer) unsafe.Pointer
)

func init() {
	lib, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
	tryRegister(&_NSAllocateCollectable, lib, "NSAllocateCollectable")
	tryRegister(&_NSAllocateMemoryPages, lib, "NSAllocateMemoryPages")
	tryRegister(&_NSClassFromString, lib, "NSClassFromString")
	tryRegister(&_NSCopyMemoryPages, lib, "NSCopyMemoryPages")
	tryRegister(&_NSCopyObject, lib, "NSCopyObject")
	tryRegister(&_NSCountFrames, lib, "NSCountFrames")
	tryRegister(&_NSDecimalCompact, lib, "NSDecimalCompact")
	tryRegister(&_NSDecimalPower, lib, "NSDecimalPower")
	tryRegister(&_NSExtraRefCount, lib, "NSExtraRefCount")
	tryRegister(&_NSFileTypeForHFSTypeCode, lib, "NSFileTypeForHFSTypeCode")
	tryRegister(&_NSFrameAddress, lib, "NSFrameAddress")
	tryRegister(&_NSFullUserName, lib, "NSFullUserName")
	tryRegister(&_NSGetUncaughtExceptionHandler, lib, "NSGetUncaughtExceptionHandler")
	tryRegister(&_NSHFSTypeCodeFromFileType, lib, "NSHFSTypeCodeFromFileType")
	tryRegister(&_NSHFSTypeOfFile, lib, "NSHFSTypeOfFile")
	tryRegister(&_NSHomeDirectory, lib, "NSHomeDirectory")
	tryRegister(&_NSHomeDirectoryForUser, lib, "NSHomeDirectoryForUser")
	tryRegister(&_NSIncrementExtraRefCount, lib, "NSIncrementExtraRefCount")
	tryRegister(&_NSIntegralRectWithOptions, lib, "NSIntegralRectWithOptions")
	tryRegister(&_NSIsFreedObject, lib, "NSIsFreedObject")
	tryRegister(&_NSLog, lib, "NSLog")
	tryRegister(&_NSLogv, lib, "NSLogv")
	tryRegister(&_NSMouseInRect, lib, "NSMouseInRect")
	tryRegister(&_NSOpenStepRootDirectory, lib, "NSOpenStepRootDirectory")
	tryRegister(&_NSReallocateCollectable, lib, "NSReallocateCollectable")
	tryRegister(&_NSRecordAllocationEvent, lib, "NSRecordAllocationEvent")
	tryRegister(&_NSReturnAddress, lib, "NSReturnAddress")
	tryRegister(&_NSRoundUpToMultipleOfPageSize, lib, "NSRoundUpToMultipleOfPageSize")
	tryRegister(&_NSSearchPathForDirectoriesInDomains, lib, "NSSearchPathForDirectoriesInDomains")
	tryRegister(&_NSSetUncaughtExceptionHandler, lib, "NSSetUncaughtExceptionHandler")
	tryRegister(&_NSSizeFromString, lib, "NSSizeFromString")
	tryRegister(&_NSStringFromProtocol, lib, "NSStringFromProtocol")
	tryRegister(&_NSTemporaryDirectory, lib, "NSTemporaryDirectory")
	tryRegister(&_NSUserName, lib, "NSUserName")
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



// Allocates collectable memory. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAllocateCollectable
func NSAllocateCollectable(size unsafe.Pointer, options unsafe.Pointer) unsafe.Pointer {
	return _NSAllocateCollectable(size, options)
	}


// Allocates a new block of memory. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAllocateMemoryPages(_:)
func NSAllocateMemoryPages(bytes unsafe.Pointer) unsafe.Pointer {
	return _NSAllocateMemoryPages(bytes)
	}


// Obtains a class by name. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSClassFromString(_:)
func NSClassFromString(aClassName unsafe.Pointer) unsafe.Pointer {
	return _NSClassFromString(aClassName)
	}


// Copies a block of memory. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCopyMemoryPages(_:_:_:)
func NSCopyMemoryPages(source unsafe.Pointer, dest unsafe.Pointer, bytes unsafe.Pointer) {
	_NSCopyMemoryPages(source, dest, bytes)
	}


// Creates an exact copy of an object. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.8.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCopyObject
func NSCopyObject(object unsafe.Pointer, extraBytes unsafe.Pointer, zone unsafe.Pointer) unsafe.Pointer {
	return _NSCopyObject(object, extraBytes, zone)
	}


// Returns the number of call frames on the stack. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCountFrames
func NSCountFrames() unsafe.Pointer {
	return _NSCountFrames()
	}


// Compacts the decimal structure for efficiency. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDecimalCompact(_:)
func NSDecimalCompact(number unsafe.Pointer) {
	_NSDecimalCompact(number)
	}


// Raises the decimal value to the specified power. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDecimalPower(_:_:_:_:)
func NSDecimalPower(result unsafe.Pointer, number unsafe.Pointer, power unsafe.Pointer, roundingMode unsafe.Pointer) unsafe.Pointer {
	return _NSDecimalPower(result, number, power, roundingMode)
	}


// Returns the specified object’s reference count. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSExtraRefCount
func NSExtraRefCount(object unsafe.Pointer) unsafe.Pointer {
	return _NSExtraRefCount(object)
	}


// Returns a string encoding a file type code. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSFileTypeForHFSTypeCode(_:)
func NSFileTypeForHFSTypeCode(hfsFileTypeCode unsafe.Pointer) unsafe.Pointer {
	return _NSFileTypeForHFSTypeCode(hfsFileTypeCode)
	}


// Returns the value of the frame pointer of the specified frame. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSFrameAddress
func NSFrameAddress(frame unsafe.Pointer) unsafe.Pointer {
	return _NSFrameAddress(frame)
	}


// Returns a string containing the full name of the current user. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSFullUserName()
func NSFullUserName() unsafe.Pointer {
	return _NSFullUserName()
	}


// Returns the top-level error handler. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSGetUncaughtExceptionHandler()
func NSGetUncaughtExceptionHandler() unsafe.Pointer {
	return _NSGetUncaughtExceptionHandler()
	}


// Returns a file type code. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSHFSTypeCodeFromFileType(_:)
func NSHFSTypeCodeFromFileType(fileTypeString unsafe.Pointer) unsafe.Pointer {
	return _NSHFSTypeCodeFromFileType(fileTypeString)
	}


// Returns a string encoding a file type. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSHFSTypeOfFile(_:)
func NSHFSTypeOfFile(fullFilePath unsafe.Pointer) unsafe.Pointer {
	return _NSHFSTypeOfFile(fullFilePath)
	}


// Returns the path to either the user’s or application’s home directory, depending on the platform. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSHomeDirectory()
func NSHomeDirectory() unsafe.Pointer {
	return _NSHomeDirectory()
	}


// Returns the path to a given user’s home directory. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSHomeDirectoryForUser(_:)
func NSHomeDirectoryForUser(userName unsafe.Pointer) unsafe.Pointer {
	return _NSHomeDirectoryForUser(userName)
	}


// Increments the specified object’s reference count. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSIncrementExtraRefCount
func NSIncrementExtraRefCount(object unsafe.Pointer) {
	_NSIncrementExtraRefCount(object)
	}


// Adjusts the sides of a rectangle to integral values using the specified options. [Full Topic]
//
// Added in macOS 10.7.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSIntegralRectWithOptions(_:_:)
func NSIntegralRectWithOptions(aRect coregraphics.CGRect, opts unsafe.Pointer) coregraphics.CGRect {
	return _NSIntegralRectWithOptions(aRect, opts)
	}


// Returns a Boolean indicating whether the specified object has been freed. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSIsFreedObject
func NSIsFreedObject(anObject unsafe.Pointer) unsafe.Pointer {
	return _NSIsFreedObject(anObject)
	}


// Logs an error message to the Apple System Log facility. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSLog
func NSLog(format unsafe.Pointer) {
	_NSLog(format)
	}


// Logs an error message to the Apple System Log facility. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSLogv(_:_:)
func NSLogv(format unsafe.Pointer, args unsafe.Pointer) {
	_NSLogv(format, args)
	}


// Returns a Boolean value that indicates whether the point is in the specified rectangle. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMouseInRect(_:_:_:)
func NSMouseInRect(aPoint coregraphics.CGPoint, aRect coregraphics.CGRect, flipped unsafe.Pointer) unsafe.Pointer {
	return _NSMouseInRect(aPoint, aRect, flipped)
	}


// Returns the root directory of the user’s system. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSOpenStepRootDirectory()
func NSOpenStepRootDirectory() unsafe.Pointer {
	return _NSOpenStepRootDirectory()
	}


// Reallocates collectable memory. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSReallocateCollectable
func NSReallocateCollectable(ptr unsafe.Pointer, size unsafe.Pointer, options unsafe.Pointer) unsafe.Pointer {
	return _NSReallocateCollectable(ptr, size, options)
	}


// Notes an object or zone allocation event and various other statistics, such as the time and current thread. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSRecordAllocationEvent
func NSRecordAllocationEvent(eventType unsafe.Pointer, object unsafe.Pointer) {
	_NSRecordAllocationEvent(eventType, object)
	}


// Returns the value of the return address of the specified frame. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSReturnAddress
func NSReturnAddress(frame unsafe.Pointer) unsafe.Pointer {
	return _NSReturnAddress(frame)
	}


// Returns the specified number of bytes rounded up to a multiple of the page size. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSRoundUpToMultipleOfPageSize(_:)
func NSRoundUpToMultipleOfPageSize(bytes unsafe.Pointer) unsafe.Pointer {
	return _NSRoundUpToMultipleOfPageSize(bytes)
	}


// Creates a list of directory search paths. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSSearchPathForDirectoriesInDomains(_:_:_:)
func NSSearchPathForDirectoriesInDomains(directory unsafe.Pointer, domainMask unsafe.Pointer, expandTilde unsafe.Pointer) unsafe.Pointer {
	return _NSSearchPathForDirectoriesInDomains(directory, domainMask, expandTilde)
	}


// Changes the top-level error handler. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSSetUncaughtExceptionHandler(_:)
func NSSetUncaughtExceptionHandler() {
	_NSSetUncaughtExceptionHandler()
	}


// Returns an from a text-based representation. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSSizeFromString(_:)
func NSSizeFromString(aString unsafe.Pointer) coregraphics.CGSize {
	return _NSSizeFromString(aString)
	}


// Returns the name of a protocol as a string. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSStringFromProtocol(_:)
func NSStringFromProtocol(proto unsafe.Pointer) unsafe.Pointer {
	return _NSStringFromProtocol(proto)
	}


// Returns the path of the temporary directory for the current user. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSTemporaryDirectory()
func NSTemporaryDirectory() unsafe.Pointer {
	return _NSTemporaryDirectory()
	}


// Returns the logon name of the current user. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSUserName()
func NSUserName() unsafe.Pointer {
	return _NSUserName()
	}


// Returns the next object from the coder. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.5.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NXReadNSObjectFromCoder
func NXReadNSObjectFromCoder(decoder unsafe.Pointer) unsafe.Pointer {
	return _NXReadNSObjectFromCoder(decoder)
	}




