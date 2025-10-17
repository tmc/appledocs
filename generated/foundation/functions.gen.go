// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/ebitengine/purego"
)

// Foundation Functions (23 total)
//
// Type-safe package-level functions with graceful error handling.
// Missing symbols are silently ignored during init; functions will panic when called if unavailable.

var (
	_NSClassFromString func(unsafe.Pointer) unsafe.Pointer
	_NSCopyObject func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_NSCountFrames func() unsafe.Pointer
	_NSFileTypeForHFSTypeCode func(unsafe.Pointer) unsafe.Pointer
	_NSFrameAddress func(unsafe.Pointer) unsafe.Pointer
	_NSFullUserName func() unsafe.Pointer
	_NSGetUncaughtExceptionHandler func() unsafe.Pointer
	_NSHFSTypeCodeFromFileType func(unsafe.Pointer) unsafe.Pointer
	_NSHFSTypeOfFile func(unsafe.Pointer) unsafe.Pointer
	_NSHomeDirectory func() unsafe.Pointer
	_NSHomeDirectoryForUser func(unsafe.Pointer) unsafe.Pointer
	_NSIntegralRectWithOptions func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_NSIsFreedObject func(unsafe.Pointer) bool
	_NSLog func(unsafe.Pointer)
	_NSLogv func(unsafe.Pointer, unsafe.Pointer)
	_NSOpenStepRootDirectory func() unsafe.Pointer
	_NSRecordAllocationEvent func(int, unsafe.Pointer)
	_NSReturnAddress func(unsafe.Pointer) unsafe.Pointer
	_NSSearchPathForDirectoriesInDomains func(unsafe.Pointer, unsafe.Pointer, bool) unsafe.Pointer
	_NSSetUncaughtExceptionHandler func()
	_NSTemporaryDirectory func() unsafe.Pointer
	_NSUserName func() unsafe.Pointer
	_NXReadNSObjectFromCoder func(unsafe.Pointer) unsafe.Pointer
)

func init() {
	lib, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}

	tryRegister(&_NSClassFromString, lib, "NSClassFromString")
	tryRegister(&_NSCopyObject, lib, "NSCopyObject")
	tryRegister(&_NSCountFrames, lib, "NSCountFrames")
	tryRegister(&_NSFileTypeForHFSTypeCode, lib, "NSFileTypeForHFSTypeCode")
	tryRegister(&_NSFrameAddress, lib, "NSFrameAddress")
	tryRegister(&_NSFullUserName, lib, "NSFullUserName")
	tryRegister(&_NSGetUncaughtExceptionHandler, lib, "NSGetUncaughtExceptionHandler")
	tryRegister(&_NSHFSTypeCodeFromFileType, lib, "NSHFSTypeCodeFromFileType")
	tryRegister(&_NSHFSTypeOfFile, lib, "NSHFSTypeOfFile")
	tryRegister(&_NSHomeDirectory, lib, "NSHomeDirectory")
	tryRegister(&_NSHomeDirectoryForUser, lib, "NSHomeDirectoryForUser")
	tryRegister(&_NSIntegralRectWithOptions, lib, "NSIntegralRectWithOptions")
	tryRegister(&_NSIsFreedObject, lib, "NSIsFreedObject")
	tryRegister(&_NSLog, lib, "NSLog")
	tryRegister(&_NSLogv, lib, "NSLogv")
	tryRegister(&_NSOpenStepRootDirectory, lib, "NSOpenStepRootDirectory")
	tryRegister(&_NSRecordAllocationEvent, lib, "NSRecordAllocationEvent")
	tryRegister(&_NSReturnAddress, lib, "NSReturnAddress")
	tryRegister(&_NSSearchPathForDirectoriesInDomains, lib, "NSSearchPathForDirectoriesInDomains")
	tryRegister(&_NSSetUncaughtExceptionHandler, lib, "NSSetUncaughtExceptionHandler")
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


// Obtains a class by name. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSClassFromString(_:)
func NSClassFromString(aClassName unsafe.Pointer) unsafe.Pointer {
	return _NSClassFromString(aClassName)
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


// Adjusts the sides of a rectangle to integral values using the specified options. [Full Topic]
//
// Added in macOS 10.7.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSIntegralRectWithOptions(_:_:)
func NSIntegralRectWithOptions(aRect unsafe.Pointer, opts unsafe.Pointer) unsafe.Pointer {
	return _NSIntegralRectWithOptions(aRect, opts)
	}


// Returns a Boolean indicating whether the specified object has been freed. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSIsFreedObject
func NSIsFreedObject(anObject unsafe.Pointer) bool {
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


// Returns the root directory of the user’s system. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSOpenStepRootDirectory()
func NSOpenStepRootDirectory() unsafe.Pointer {
	return _NSOpenStepRootDirectory()
	}


// Notes an object or zone allocation event and various other statistics, such as the time and current thread. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSRecordAllocationEvent
func NSRecordAllocationEvent(eventType int, object unsafe.Pointer) {
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


// Creates a list of directory search paths. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSSearchPathForDirectoriesInDomains(_:_:_:)
func NSSearchPathForDirectoriesInDomains(directory unsafe.Pointer, domainMask unsafe.Pointer, expandTilde bool) unsafe.Pointer {
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




