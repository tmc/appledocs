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
	exceptionClass     _ExceptionClass
	exceptionClassOnce sync.Once
)

func getExceptionClass() _ExceptionClass {
	exceptionClassOnce.Do(func() {
		exceptionClass = _ExceptionClass{objc.GetClass("NSException")}
	})
	return exceptionClass
}

type _ExceptionClass struct {
	class objc.Class
}

// An interface definition for the [Exception] class.
type IException interface {
	objectivec.IObject
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


// A dictionary containing application-specific data pertaining to the receiver.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSException/userInfo-swift.property
func (e_ Exception) UserInfo() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](e_.ID, objc.Sel("userInfo"))
	return rv
}




