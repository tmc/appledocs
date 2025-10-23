// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/ebitengine/purego"
)


// Foundation Functions (32 total)
//
// Type-safe package-level functions with graceful error handling.
// Missing symbols are silently ignored during init; functions will panic when called if unavailable.

var (
	_NSAllocateObject func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_NSClassFromString func(unsafe.Pointer) unsafe.Pointer
	_NSCopyObject func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_NSCountFrames func() unsafe.Pointer
	_NSDeallocateObject func(unsafe.Pointer) unsafe.Pointer
	_NSDecrementExtraRefCountWasZero func(unsafe.Pointer) bool
	_NSExtraRefCount func(unsafe.Pointer) unsafe.Pointer
	_NSFileTypeForHFSTypeCode func(unsafe.Pointer) unsafe.Pointer
	_NSFrameAddress func(unsafe.Pointer) unsafe.Pointer
	_NSFullUserName func() unsafe.Pointer
	_NSGetSizeAndAlignment func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_NSHFSTypeCodeFromFileType func(unsafe.Pointer) unsafe.Pointer
	_NSHFSTypeOfFile func(unsafe.Pointer) unsafe.Pointer
	_NSHomeDirectory func() unsafe.Pointer
	_NSHomeDirectoryForUser func(unsafe.Pointer) unsafe.Pointer
	_NSIncrementExtraRefCount func(unsafe.Pointer) unsafe.Pointer
	_NSIsFreedObject func(unsafe.Pointer) bool
	_NSLog func(unsafe.Pointer) unsafe.Pointer
	_NSLogv func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_NSOpenStepRootDirectory func() unsafe.Pointer
	_NSProtocolFromString func(unsafe.Pointer) unsafe.Pointer
	_NSRecordAllocationEvent func(int, unsafe.Pointer) unsafe.Pointer
	_NSReturnAddress func(unsafe.Pointer) unsafe.Pointer
	_NSSearchPathForDirectoriesInDomains func(unsafe.Pointer, unsafe.Pointer, bool) unsafe.Pointer
	_NSSelectorFromString func(unsafe.Pointer) unsafe.Pointer
	_NSShouldRetainWithZone func(unsafe.Pointer, unsafe.Pointer) bool
	_NSStringFromClass func(unsafe.Pointer) unsafe.Pointer
	_NSStringFromProtocol func(unsafe.Pointer) unsafe.Pointer
	_NSStringFromSelector func(unsafe.Pointer) unsafe.Pointer
	_NSTemporaryDirectory func() unsafe.Pointer
	_NSUserName func() unsafe.Pointer
	_NXReadNSObjectFromCoder func(unsafe.Pointer) unsafe.Pointer
)

func init() {
	lib, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
	tryRegister(&_NSAllocateObject, lib, "NSAllocateObject")
	tryRegister(&_NSClassFromString, lib, "NSClassFromString")
	tryRegister(&_NSCopyObject, lib, "NSCopyObject")
	tryRegister(&_NSCountFrames, lib, "NSCountFrames")
	tryRegister(&_NSDeallocateObject, lib, "NSDeallocateObject")
	tryRegister(&_NSDecrementExtraRefCountWasZero, lib, "NSDecrementExtraRefCountWasZero")
	tryRegister(&_NSExtraRefCount, lib, "NSExtraRefCount")
	tryRegister(&_NSFileTypeForHFSTypeCode, lib, "NSFileTypeForHFSTypeCode")
	tryRegister(&_NSFrameAddress, lib, "NSFrameAddress")
	tryRegister(&_NSFullUserName, lib, "NSFullUserName")
	tryRegister(&_NSGetSizeAndAlignment, lib, "NSGetSizeAndAlignment")
	tryRegister(&_NSHFSTypeCodeFromFileType, lib, "NSHFSTypeCodeFromFileType")
	tryRegister(&_NSHFSTypeOfFile, lib, "NSHFSTypeOfFile")
	tryRegister(&_NSHomeDirectory, lib, "NSHomeDirectory")
	tryRegister(&_NSHomeDirectoryForUser, lib, "NSHomeDirectoryForUser")
	tryRegister(&_NSIncrementExtraRefCount, lib, "NSIncrementExtraRefCount")
	tryRegister(&_NSIsFreedObject, lib, "NSIsFreedObject")
	tryRegister(&_NSLog, lib, "NSLog")
	tryRegister(&_NSLogv, lib, "NSLogv")
	tryRegister(&_NSOpenStepRootDirectory, lib, "NSOpenStepRootDirectory")
	tryRegister(&_NSProtocolFromString, lib, "NSProtocolFromString")
	tryRegister(&_NSRecordAllocationEvent, lib, "NSRecordAllocationEvent")
	tryRegister(&_NSReturnAddress, lib, "NSReturnAddress")
	tryRegister(&_NSSearchPathForDirectoriesInDomains, lib, "NSSearchPathForDirectoriesInDomains")
	tryRegister(&_NSSelectorFromString, lib, "NSSelectorFromString")
	tryRegister(&_NSShouldRetainWithZone, lib, "NSShouldRetainWithZone")
	tryRegister(&_NSStringFromClass, lib, "NSStringFromClass")
	tryRegister(&_NSStringFromProtocol, lib, "NSStringFromProtocol")
	tryRegister(&_NSStringFromSelector, lib, "NSStringFromSelector")
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



// Creates and returns a new instance of a given class.
//
// Added in macOS 10.0.

// Creates and returns a new instance of a given class.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAllocateObject
func NSAllocateObject(aClass unsafe.Pointer, extraBytes unsafe.Pointer, zone unsafe.Pointer) unsafe.Pointer {
	return _NSAllocateObject(aClass, extraBytes, zone)
	}


// Obtains a class by name.
//
// Added in macOS 10.0.

// Obtains a class by name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSClassFromString(_:)
func NSClassFromString(aClassName unsafe.Pointer) unsafe.Pointer {
	return _NSClassFromString(aClassName)
	}


// Creates an exact copy of an object.
//
// Deprecated: This function was deprecated in macOS 10.8.
//
// Added in macOS 10.0.

// Creates an exact copy of an object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCopyObject
func NSCopyObject(object unsafe.Pointer, extraBytes unsafe.Pointer, zone unsafe.Pointer) unsafe.Pointer {
	return _NSCopyObject(object, extraBytes, zone)
	}


// Returns the number of call frames on the stack.
//
// Added in macOS 10.0.

// Returns the number of call frames on the stack.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCountFrames
func NSCountFrames() unsafe.Pointer {
	return _NSCountFrames()
	}


// Destroys an existing object.
//
// Added in macOS 10.0.

// Destroys an existing object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDeallocateObject
func NSDeallocateObject(object unsafe.Pointer) {
	_NSDeallocateObject(object)
	}


// Decrements the specified object’s reference count.
//
// Added in macOS 10.0.

// Decrements the specified object’s reference count.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDecrementExtraRefCountWasZero
func NSDecrementExtraRefCountWasZero(object unsafe.Pointer) bool {
	return _NSDecrementExtraRefCountWasZero(object)
	}


// Returns the specified object’s reference count.
//
// Added in macOS 10.0.

// Returns the specified object’s reference count.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSExtraRefCount
func NSExtraRefCount(object unsafe.Pointer) unsafe.Pointer {
	return _NSExtraRefCount(object)
	}


// Returns a string encoding a file type code.
//
// Added in macOS 10.0.

// Returns a string encoding a file type code.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSFileTypeForHFSTypeCode(_:)
func NSFileTypeForHFSTypeCode(hfsFileTypeCode unsafe.Pointer) unsafe.Pointer {
	return _NSFileTypeForHFSTypeCode(hfsFileTypeCode)
	}


// Returns the value of the frame pointer of the specified frame.
//
// Added in macOS 10.0.

// Returns the value of the frame pointer of the specified frame.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSFrameAddress
func NSFrameAddress(frame unsafe.Pointer) unsafe.Pointer {
	return _NSFrameAddress(frame)
	}


// Returns a string containing the full name of the current user.
//
// Added in macOS 10.0.

// Returns a string containing the full name of the current user.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSFullUserName()
func NSFullUserName() unsafe.Pointer {
	return _NSFullUserName()
	}


// Obtains the actual size and the aligned size of an encoded type.
//
// Added in macOS 10.0.

// Obtains the actual size and the aligned size of an encoded type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSGetSizeAndAlignment(_:_:_:)
func NSGetSizeAndAlignment(typePtr unsafe.Pointer, sizep unsafe.Pointer, alignp unsafe.Pointer) unsafe.Pointer {
	return _NSGetSizeAndAlignment(typePtr, sizep, alignp)
	}


// Returns a file type code.
//
// Added in macOS 10.0.

// Returns a file type code.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSHFSTypeCodeFromFileType(_:)
func NSHFSTypeCodeFromFileType(fileTypeString unsafe.Pointer) unsafe.Pointer {
	return _NSHFSTypeCodeFromFileType(fileTypeString)
	}


// Returns a string encoding a file type.
//
// Added in macOS 10.0.

// Returns a string encoding a file type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSHFSTypeOfFile(_:)
func NSHFSTypeOfFile(fullFilePath unsafe.Pointer) unsafe.Pointer {
	return _NSHFSTypeOfFile(fullFilePath)
	}


// Returns the path to either the user’s or application’s home directory, depending on the platform.
//
// Added in macOS 10.0.

// Returns the path to either the user’s or application’s home directory, depending on the platform.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSHomeDirectory()
func NSHomeDirectory() unsafe.Pointer {
	return _NSHomeDirectory()
	}


// Returns the path to a given user’s home directory.
//
// Added in macOS 10.0.

// Returns the path to a given user’s home directory.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSHomeDirectoryForUser(_:)
func NSHomeDirectoryForUser(userName unsafe.Pointer) unsafe.Pointer {
	return _NSHomeDirectoryForUser(userName)
	}


// Increments the specified object’s reference count.
//
// Added in macOS 10.0.

// Increments the specified object’s reference count.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSIncrementExtraRefCount
func NSIncrementExtraRefCount(object unsafe.Pointer) {
	_NSIncrementExtraRefCount(object)
	}


// Returns a Boolean indicating whether the specified object has been freed.
//
// Added in macOS 10.0.

// Returns a Boolean indicating whether the specified object has been freed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSIsFreedObject
func NSIsFreedObject(anObject unsafe.Pointer) bool {
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


// Returns the root directory of the user’s system.
//
// Added in macOS 10.0.

// Returns the root directory of the user’s system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSOpenStepRootDirectory()
func NSOpenStepRootDirectory() unsafe.Pointer {
	return _NSOpenStepRootDirectory()
	}


// Returns a the protocol with a given name.
//
// Added in macOS 10.5.

// Returns a the protocol with a given name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSProtocolFromString(_:)
func NSProtocolFromString(namestr unsafe.Pointer) unsafe.Pointer {
	return _NSProtocolFromString(namestr)
	}


// Notes an object or zone allocation event and various other statistics, such as the time and current thread.
//
// Added in macOS 10.0.

// Notes an object or zone allocation event and various other statistics, such as the time and current thread.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSRecordAllocationEvent
func NSRecordAllocationEvent(eventType int, object unsafe.Pointer) {
	_NSRecordAllocationEvent(eventType, object)
	}


// Returns the value of the return address of the specified frame.
//
// Added in macOS 10.0.

// Returns the value of the return address of the specified frame.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSReturnAddress
func NSReturnAddress(frame unsafe.Pointer) unsafe.Pointer {
	return _NSReturnAddress(frame)
	}


// Creates a list of directory search paths.
//
// Added in macOS 10.0.

// Creates a list of directory search paths.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSSearchPathForDirectoriesInDomains(_:_:_:)
func NSSearchPathForDirectoriesInDomains(directory unsafe.Pointer, domainMask unsafe.Pointer, expandTilde bool) unsafe.Pointer {
	return _NSSearchPathForDirectoriesInDomains(directory, domainMask, expandTilde)
	}


// Returns the selector with a given name.
//
// Added in macOS 10.0.

// Returns the selector with a given name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSSelectorFromString(_:)
func NSSelectorFromString(aSelectorName unsafe.Pointer) unsafe.Pointer {
	return _NSSelectorFromString(aSelectorName)
	}


// Indicates whether an object should be retained.
//
// Added in macOS 10.0.

// Indicates whether an object should be retained.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSShouldRetainWithZone
func NSShouldRetainWithZone(anObject unsafe.Pointer, requestedZone unsafe.Pointer) bool {
	return _NSShouldRetainWithZone(anObject, requestedZone)
	}


// Returns the name of a class as a string.
//
// Added in macOS 10.0.

// Returns the name of a class as a string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSStringFromClass(_:)
func NSStringFromClass(aClass unsafe.Pointer) unsafe.Pointer {
	return _NSStringFromClass(aClass)
	}


// Returns the name of a protocol as a string.
//
// Added in macOS 10.5.

// Returns the name of a protocol as a string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSStringFromProtocol(_:)
func NSStringFromProtocol(proto unsafe.Pointer) unsafe.Pointer {
	return _NSStringFromProtocol(proto)
	}


// Returns a string representation of a given selector.
//
// Added in macOS 10.0.

// Returns a string representation of a given selector.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSStringFromSelector(_:)
func NSStringFromSelector(aSelector unsafe.Pointer) unsafe.Pointer {
	return _NSStringFromSelector(aSelector)
	}


// Returns the path of the temporary directory for the current user.
//
// Added in macOS 10.0.

// Returns the path of the temporary directory for the current user.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSTemporaryDirectory()
func NSTemporaryDirectory() unsafe.Pointer {
	return _NSTemporaryDirectory()
	}


// Returns the logon name of the current user.
//
// Added in macOS 10.0.

// Returns the logon name of the current user.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSUserName()
func NSUserName() unsafe.Pointer {
	return _NSUserName()
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




