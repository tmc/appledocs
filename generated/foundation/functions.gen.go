// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/ebitengine/purego"
)


// Foundation Functions (43 total)
//
// Type-safe package-level functions with graceful error handling.
// Missing symbols are silently ignored during init; functions will panic when called if unavailable.

var (
	_NSCountFrames func() uint64
	_NSDecimalAdd func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_NSDecimalCompact func(unsafe.Pointer)
	_NSDecimalCompare func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_NSDecimalCopy func(unsafe.Pointer, unsafe.Pointer)
	_NSDecimalDivide func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_NSDecimalMultiply func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_NSDecimalMultiplyByPowerOf10 func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_NSDecimalNormalize func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_NSDecimalPower func(unsafe.Pointer, unsafe.Pointer, uint64, unsafe.Pointer) unsafe.Pointer
	_NSDecimalRound func(unsafe.Pointer, unsafe.Pointer, int64, unsafe.Pointer)
	_NSDecimalString func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_NSDecimalSubtract func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_NSEdgeInsetsEqual func(unsafe.Pointer, unsafe.Pointer) bool
	_NSFileTypeForHFSTypeCode func(unsafe.Pointer) unsafe.Pointer
	_NSFrameAddress func(uint64) unsafe.Pointer
	_NSFullUserName func() unsafe.Pointer
	_NSGetUncaughtExceptionHandler func() unsafe.Pointer
	_NSHFSTypeCodeFromFileType func(unsafe.Pointer) unsafe.Pointer
	_NSHFSTypeOfFile func(unsafe.Pointer) unsafe.Pointer
	_NSHomeDirectory func() unsafe.Pointer
	_NSHomeDirectoryForUser func(unsafe.Pointer) unsafe.Pointer
	_NSIsFreedObject func(unsafe.Pointer) bool
	_NSLog func(unsafe.Pointer)
	_NSLogv func(unsafe.Pointer, unsafe.Pointer)
	_NSOpenStepRootDirectory func() unsafe.Pointer
	_NSPointFromString func(unsafe.Pointer) Point
	_NSRangeFromString func(unsafe.Pointer) unsafe.Pointer
	_NSRecordAllocationEvent func(int, unsafe.Pointer)
	_NSRectFromString func(unsafe.Pointer) Rect
	_NSReturnAddress func(uint64) unsafe.Pointer
	_NSSearchPathForDirectoriesInDomains func(unsafe.Pointer, unsafe.Pointer, bool) []unsafe.Pointer
	_NSSetUncaughtExceptionHandler func()
	_NSSizeFromString func(unsafe.Pointer) Size
	_NSTemporaryDirectory func() unsafe.Pointer
	_NSUserName func() unsafe.Pointer
	_NSZoneCalloc func(unsafe.Pointer, uint64, uint64) unsafe.Pointer
	_NSZoneFree func(unsafe.Pointer, unsafe.Pointer)
	_NSZoneFromPointer func(unsafe.Pointer) unsafe.Pointer
	_NSZoneMalloc func(unsafe.Pointer, uint64) unsafe.Pointer
	_NSZoneName func(unsafe.Pointer) unsafe.Pointer
	_NSZoneRealloc func(unsafe.Pointer, unsafe.Pointer, uint64) unsafe.Pointer
	_NXReadNSObjectFromCoder func(unsafe.Pointer) unsafe.Pointer
)

func init() {
	lib, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
	tryRegister(&_NSCountFrames, lib, "NSCountFrames")
	tryRegister(&_NSDecimalAdd, lib, "NSDecimalAdd")
	tryRegister(&_NSDecimalCompact, lib, "NSDecimalCompact")
	tryRegister(&_NSDecimalCompare, lib, "NSDecimalCompare")
	tryRegister(&_NSDecimalCopy, lib, "NSDecimalCopy")
	tryRegister(&_NSDecimalDivide, lib, "NSDecimalDivide")
	tryRegister(&_NSDecimalMultiply, lib, "NSDecimalMultiply")
	tryRegister(&_NSDecimalMultiplyByPowerOf10, lib, "NSDecimalMultiplyByPowerOf10")
	tryRegister(&_NSDecimalNormalize, lib, "NSDecimalNormalize")
	tryRegister(&_NSDecimalPower, lib, "NSDecimalPower")
	tryRegister(&_NSDecimalRound, lib, "NSDecimalRound")
	tryRegister(&_NSDecimalString, lib, "NSDecimalString")
	tryRegister(&_NSDecimalSubtract, lib, "NSDecimalSubtract")
	tryRegister(&_NSEdgeInsetsEqual, lib, "NSEdgeInsetsEqual")
	tryRegister(&_NSFileTypeForHFSTypeCode, lib, "NSFileTypeForHFSTypeCode")
	tryRegister(&_NSFrameAddress, lib, "NSFrameAddress")
	tryRegister(&_NSFullUserName, lib, "NSFullUserName")
	tryRegister(&_NSGetUncaughtExceptionHandler, lib, "NSGetUncaughtExceptionHandler")
	tryRegister(&_NSHFSTypeCodeFromFileType, lib, "NSHFSTypeCodeFromFileType")
	tryRegister(&_NSHFSTypeOfFile, lib, "NSHFSTypeOfFile")
	tryRegister(&_NSHomeDirectory, lib, "NSHomeDirectory")
	tryRegister(&_NSHomeDirectoryForUser, lib, "NSHomeDirectoryForUser")
	tryRegister(&_NSIsFreedObject, lib, "NSIsFreedObject")
	tryRegister(&_NSLog, lib, "NSLog")
	tryRegister(&_NSLogv, lib, "NSLogv")
	tryRegister(&_NSOpenStepRootDirectory, lib, "NSOpenStepRootDirectory")
	tryRegister(&_NSPointFromString, lib, "NSPointFromString")
	tryRegister(&_NSRangeFromString, lib, "NSRangeFromString")
	tryRegister(&_NSRecordAllocationEvent, lib, "NSRecordAllocationEvent")
	tryRegister(&_NSRectFromString, lib, "NSRectFromString")
	tryRegister(&_NSReturnAddress, lib, "NSReturnAddress")
	tryRegister(&_NSSearchPathForDirectoriesInDomains, lib, "NSSearchPathForDirectoriesInDomains")
	tryRegister(&_NSSetUncaughtExceptionHandler, lib, "NSSetUncaughtExceptionHandler")
	tryRegister(&_NSSizeFromString, lib, "NSSizeFromString")
	tryRegister(&_NSTemporaryDirectory, lib, "NSTemporaryDirectory")
	tryRegister(&_NSUserName, lib, "NSUserName")
	tryRegister(&_NSZoneCalloc, lib, "NSZoneCalloc")
	tryRegister(&_NSZoneFree, lib, "NSZoneFree")
	tryRegister(&_NSZoneFromPointer, lib, "NSZoneFromPointer")
	tryRegister(&_NSZoneMalloc, lib, "NSZoneMalloc")
	tryRegister(&_NSZoneName, lib, "NSZoneName")
	tryRegister(&_NSZoneRealloc, lib, "NSZoneRealloc")
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

// Adds two decimal values.
//
// Added in macOS 10.0.
// Adds two decimal values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDecimalAdd(_:_:_:_:)
func NSDecimalAdd(result unsafe.Pointer, leftOperand unsafe.Pointer, rightOperand unsafe.Pointer, roundingMode unsafe.Pointer) unsafe.Pointer {
	return _NSDecimalAdd(result, leftOperand, rightOperand, roundingMode)
}

// Compacts the decimal structure for efficiency.
//
// Added in macOS 10.0.
// Compacts the decimal structure for efficiency.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDecimalCompact(_:)
func NSDecimalCompact(number unsafe.Pointer) {
	_NSDecimalCompact(number)
}

// Compares two decimal values.
//
// Added in macOS 10.0.
// Compares two decimal values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDecimalCompare(_:_:)
func NSDecimalCompare(leftOperand unsafe.Pointer, rightOperand unsafe.Pointer) unsafe.Pointer {
	return _NSDecimalCompare(leftOperand, rightOperand)
}

// Copies the value of a decimal number.
//
// Added in macOS 10.0.
// Copies the value of a decimal number.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDecimalCopy(_:_:)
func NSDecimalCopy(destination unsafe.Pointer, source unsafe.Pointer) {
	_NSDecimalCopy(destination, source)
}

// Divides one decimal value by another.
//
// Added in macOS 10.0.
// Divides one decimal value by another.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDecimalDivide(_:_:_:_:)
func NSDecimalDivide(result unsafe.Pointer, leftOperand unsafe.Pointer, rightOperand unsafe.Pointer, roundingMode unsafe.Pointer) unsafe.Pointer {
	return _NSDecimalDivide(result, leftOperand, rightOperand, roundingMode)
}

// Multiplies two decimal numbers together.
//
// Added in macOS 10.0.
// Multiplies two decimal numbers together.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDecimalMultiply(_:_:_:_:)
func NSDecimalMultiply(result unsafe.Pointer, leftOperand unsafe.Pointer, rightOperand unsafe.Pointer, roundingMode unsafe.Pointer) unsafe.Pointer {
	return _NSDecimalMultiply(result, leftOperand, rightOperand, roundingMode)
}

// Multiplies a decimal by the specified power of 10.
//
// Added in macOS 10.0.
// Multiplies a decimal by the specified power of 10.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDecimalMultiplyByPowerOf10(_:_:_:_:)
func NSDecimalMultiplyByPowerOf10(result unsafe.Pointer, number unsafe.Pointer, power unsafe.Pointer, roundingMode unsafe.Pointer) unsafe.Pointer {
	return _NSDecimalMultiplyByPowerOf10(result, number, power, roundingMode)
}

// Normalizes the internal format of two decimal numbers to simplify later operations.
//
// Added in macOS 10.0.
// Normalizes the internal format of two decimal numbers to simplify later operations.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDecimalNormalize(_:_:_:)
func NSDecimalNormalize(number1 unsafe.Pointer, number2 unsafe.Pointer, roundingMode unsafe.Pointer) unsafe.Pointer {
	return _NSDecimalNormalize(number1, number2, roundingMode)
}

// Raises the decimal value to the specified power.
//
// Added in macOS 10.0.
// Raises the decimal value to the specified power.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDecimalPower(_:_:_:_:)
func NSDecimalPower(result unsafe.Pointer, number unsafe.Pointer, power uint64, roundingMode unsafe.Pointer) unsafe.Pointer {
	return _NSDecimalPower(result, number, power, roundingMode)
}

// Rounds off the decimal value.
//
// Added in macOS 10.0.
// Rounds off the decimal value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDecimalRound(_:_:_:_:)
func NSDecimalRound(result unsafe.Pointer, number unsafe.Pointer, scale int64, roundingMode unsafe.Pointer) {
	_NSDecimalRound(result, number, scale, roundingMode)
}

// Returns a string representation of the decimal value appropriate for the specified locale.
//
// Added in macOS 10.0.
// Returns a string representation of the decimal value appropriate for the specified locale.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDecimalString(_:_:)
func NSDecimalString(dcm unsafe.Pointer, locale unsafe.Pointer) unsafe.Pointer {
	return _NSDecimalString(dcm, locale)
}

// Subtracts one decimal value from another.
//
// Added in macOS 10.0.
// Subtracts one decimal value from another.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDecimalSubtract(_:_:_:_:)
func NSDecimalSubtract(result unsafe.Pointer, leftOperand unsafe.Pointer, rightOperand unsafe.Pointer, roundingMode unsafe.Pointer) unsafe.Pointer {
	return _NSDecimalSubtract(result, leftOperand, rightOperand, roundingMode)
}

// Returns a Boolean value that indicates whether two edge insets structures are equal.
//
// Added in macOS 10.10.
// Returns a Boolean value that indicates whether two edge insets structures are equal.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSEdgeInsetsEqual(_:_:)
func NSEdgeInsetsEqual(aInsets unsafe.Pointer, bInsets unsafe.Pointer) bool {
	return _NSEdgeInsetsEqual(aInsets, bInsets)
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
func NSFrameAddress(frame uint64) unsafe.Pointer {
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

// Returns the top-level error handler.
//
// Added in macOS 10.0.
// Returns the top-level error handler.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSGetUncaughtExceptionHandler()
func NSGetUncaughtExceptionHandler() unsafe.Pointer {
	return _NSGetUncaughtExceptionHandler()
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

// Returns a point from a text-based representation.
//
// Added in macOS 10.0.
// Returns a point from a text-based representation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSPointFromString(_:)
func NSPointFromString(aString unsafe.Pointer) Point {
	return _NSPointFromString(aString)
}

// Returns a range from a textual representation.
//
// Added in macOS 10.0.
// Returns a range from a textual representation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSRangeFromString(_:)
func NSRangeFromString(aString unsafe.Pointer) unsafe.Pointer {
	return _NSRangeFromString(aString)
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

// Returns a rectangle from a text-based representation.
//
// Added in macOS 10.0.
// Returns a rectangle from a text-based representation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSRectFromString(_:)
func NSRectFromString(aString unsafe.Pointer) Rect {
	return _NSRectFromString(aString)
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

// Creates a list of directory search paths.
//
// Added in macOS 10.0.
// Creates a list of directory search paths.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSSearchPathForDirectoriesInDomains(_:_:_:)
func NSSearchPathForDirectoriesInDomains(directory unsafe.Pointer, domainMask unsafe.Pointer, expandTilde bool) []unsafe.Pointer {
	return _NSSearchPathForDirectoriesInDomains(directory, domainMask, expandTilde)
}

// Changes the top-level error handler.
//
// Added in macOS 10.0.
// Changes the top-level error handler.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSSetUncaughtExceptionHandler(_:)
func NSSetUncaughtExceptionHandler() {
	_NSSetUncaughtExceptionHandler()
}

// Returns an from a text-based representation.
//
// Added in macOS 10.0.
// Returns an from a text-based representation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSSizeFromString(_:)
func NSSizeFromString(aString unsafe.Pointer) Size {
	return _NSSizeFromString(aString)
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

// Allocates memory in a zone.
//
// Added in macOS 10.0.
// Allocates memory in a zone.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSZoneCalloc
func NSZoneCalloc(zone unsafe.Pointer, numElems uint64, byteSize uint64) unsafe.Pointer {
	return _NSZoneCalloc(zone, numElems, byteSize)
}

// Deallocates a block of memory in the specified zone.
//
// Added in macOS 10.0.
// Deallocates a block of memory in the specified zone.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSZoneFree
func NSZoneFree(zone unsafe.Pointer, ptr unsafe.Pointer) {
	_NSZoneFree(zone, ptr)
}

// Gets the zone for a given block of memory.
//
// Added in macOS 10.0.
// Gets the zone for a given block of memory.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSZoneFromPointer
func NSZoneFromPointer(ptr unsafe.Pointer) unsafe.Pointer {
	return _NSZoneFromPointer(ptr)
}

// Allocates memory in a zone.
//
// Added in macOS 10.0.
// Allocates memory in a zone.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSZoneMalloc
func NSZoneMalloc(zone unsafe.Pointer, size uint64) unsafe.Pointer {
	return _NSZoneMalloc(zone, size)
}

// Returns the name of the specified zone.
//
// Added in macOS 10.0.
// Returns the name of the specified zone.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSZoneName
func NSZoneName(zone unsafe.Pointer) unsafe.Pointer {
	return _NSZoneName(zone)
}

// Allocates memory in a zone.
//
// Added in macOS 10.0.
// Allocates memory in a zone.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSZoneRealloc
func NSZoneRealloc(zone unsafe.Pointer, ptr unsafe.Pointer, size uint64) unsafe.Pointer {
	return _NSZoneRealloc(zone, ptr, size)
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



