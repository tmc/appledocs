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
	// properties:
	UserInfo() objc.ID
	CallStackReturnAddresses() INumber
	SetCallStackReturnAddresses(value INumber)
	CallStackSymbols() string /* primitive/slice/pointer */
	SetCallStackSymbols(value string /* primitive/slice/pointer */)
	Name() ExceptionName /* foo */
	SetName(value ExceptionName /* foo */)
	Reason() string /* primitive/slice/pointer */
	SetReason(value string /* primitive/slice/pointer */)
	// methods:
}

// An object that represents a special condition that interrupts the normal flow of program execution.
//
// Use to implement exception handling. An exception is a special condition that interrupts the normal flow of program execution. Each application can interrupt the program for different reasons. For example, one application might interpret saving a file in a directory that is write-protected as an exception. In this sense, the exception is equivalent to an error. Another application might interpret the user’s key-press (for example, Control-C) as an exception: an indication that a long-running process should abort.


// An object that represents a special condition that interrupts the normal flow of program execution.
//
// [Full Topic]
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



// A dictionary containing application-specific data pertaining to the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSException/userInfo-swift.property
func (e_ Exception) UserInfo() objc.ID {
	rv := objc.Send[objc.ID](e_.ID, objc.Sel("userInfo"))
	return rv
}


// The call return addresses related to a raised exception.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsexception/callstackreturnaddresses
func (e_ Exception) CallStackReturnAddresses() INumber {
	rv := objc.Send[Number](e_.ID, objc.Sel("callStackReturnAddresses"))
	return rv
}


// The call return addresses related to a raised exception.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsexception/callstackreturnaddresses
func (e_ Exception) SetCallStackReturnAddresses(value INumber) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setCallStackReturnAddresses:"), value)
}


// An array containing the current call stack symbols.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsexception/callstacksymbols
func (e_ Exception) CallStackSymbols() string /* primitive/slice/pointer */ {
	rv := objc.Send[string](e_.ID, objc.Sel("callStackSymbols"))
	return rv
}


// An array containing the current call stack symbols.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsexception/callstacksymbols
func (e_ Exception) SetCallStackSymbols(value string /* primitive/slice/pointer */) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setCallStackSymbols:"), objc.String(value))
}


// A string used to uniquely identify the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsexception/name-swift.property
func (e_ Exception) Name() ExceptionName /* foo */ {
	rv := objc.Send[ExceptionName](e_.ID, objc.Sel("name"))
	return rv
}


// A string used to uniquely identify the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsexception/name-swift.property
func (e_ Exception) SetName(value ExceptionName /* foo */) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setName:"), value)
}


// A string containing a “human-readable” reason for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsexception/reason-swift.property
func (e_ Exception) Reason() string /* primitive/slice/pointer */ {
	rv := objc.Send[string](e_.ID, objc.Sel("reason"))
	return rv
}


// A string containing a “human-readable” reason for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsexception/reason-swift.property
func (e_ Exception) SetReason(value string /* primitive/slice/pointer */) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setReason:"), objc.String(value))
}



