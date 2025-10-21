// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [Exception] class.
var (
	ExceptionClass     _ExceptionClass
	ExceptionClassOnce sync.Once
)

func getExceptionClass() _ExceptionClass {
	ExceptionClassOnce.Do(func() {
		ExceptionClass = _ExceptionClass{objc.GetClass("NSException")}
	})
	return ExceptionClass
}

type _ExceptionClass struct {
	class objc.Class
}

// An interface definition for the [Exception] class.
type IException interface {
	objectivec.IObject
	Raise()
}

// An object that represents a special condition that interrupts the normal flow of program execution.
//
// Use to implement exception handling. An exception is a special condition that interrupts the normal flow of program execution. Each application can interrupt the program for different reasons. For example, one application might interpret saving a file in a directory that is write-protected as an exception. In this sense, the exception is equivalent to an error. Another application might interpret the user’s key-press (for example, Control-C) as an exception: an indication that a long-running process should abort.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSException
type Exception struct {
	objectivec.Object
}

// ExceptionFrom constructs a [Exception] from an unsafe.Pointer.
//
// An object that represents a special condition that interrupts the normal flow of program execution.
func ExceptionFrom(ptr unsafe.Pointer) Exception {
	return Exception{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ec _ExceptionClass) Alloc() Exception {
	rv := objc.Send[Exception](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ec _ExceptionClass) New() Exception {
	rv := objc.Send[Exception](objc.ID(ec.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (e_ Exception) Init() Exception {
	rv := objc.Send[Exception](e_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e_ Exception) Autorelease() Exception {
	rv := objc.Send[Exception](e_.ID, objc.Sel("autorelease"))
	return rv
}

// NewException creates a new Exception instance.
func NewException() Exception {
	return getExceptionClass().New()
}

// Initializes and returns a newly allocated exception object.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSException/init(name:reason:userInfo:)
func NewExceptionWithNameReasonUserInfo(aName unsafe.Pointer, aReason string, aUserInfo objc.ID) Exception {
	instance := getExceptionClass().Alloc()
	rv := objc.Send[Exception](instance.ID, objc.Sel("initWithName:reason:userInfo:"), aName, objc.String(aReason), aUserInfo)
	rv.Autorelease()
	return rv
}

// Creates and returns an exception object .
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSException/exceptionWithName:reason:userInfo:
func (ec _ExceptionClass) ExceptionWithNameReasonUserInfo(name unsafe.Pointer, reason string, userInfo objc.ID) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(ec.class), objc.Sel("exceptionWithName:reason:userInfo:"), name, objc.String(reason), userInfo)
	return rv
}

// Creates and raises an exception with the specified name, reason, and arguments.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSException/raise(_:format:arguments:)
func (ec _ExceptionClass) RaiseFormatArguments(name unsafe.Pointer, format string, argList unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(ec.class), objc.Sel("raise:format:arguments:"), name, objc.String(format), argList)
}

// A convenience method that creates and raises an exception.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSException/raise:format:
func (ec _ExceptionClass) RaiseFormat(name unsafe.Pointer, format string) {
	objc.Send[objc.ID](objc.ID(ec.class), objc.Sel("raise:format:"), name, objc.String(format))
}

// Raises the receiver, causing program flow to jump to the local exception handler.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSException/raise()
func (e_ Exception) Raise() {
	objc.Send[objc.ID](e_.ID, objc.Sel("raise"))
}

// The call return addresses related to a raised exception.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSException/callStackReturnAddresses
func (e_ Exception) CallStackReturnAddresses() []Number {
	rv := objc.Send[[]Number](e_.ID, objc.Sel("callStackReturnAddresses"))
	return rv
}

// An array containing the current call stack symbols.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSException/callStackSymbols
func (e_ Exception) CallStackSymbols() []string {
	rv := objc.Send[[]string](e_.ID, objc.Sel("callStackSymbols"))
	return rv
}

// A string used to uniquely identify the receiver.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSException/name-swift.property
func (e_ Exception) Name() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](e_.ID, objc.Sel("name"))
	return rv
}

// A string containing a “human-readable” reason for the receiver.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSException/reason-swift.property
func (e_ Exception) Reason() string {
	rv := objc.Send[string](e_.ID, objc.Sel("reason"))
	return rv
}

// A dictionary containing application-specific data pertaining to the receiver.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSException/userInfo-swift.property
func (e_ Exception) UserInfo() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](e_.ID, objc.Sel("userInfo"))
	return rv
}
