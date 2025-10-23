// Code generated from Apple documentation for ExceptionHandling. DO NOT EDIT.

package exceptionhandling

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [ExceptionHandler] class.
var (
	ExceptionHandlerClass     _ExceptionHandlerClass
	ExceptionHandlerClassOnce sync.Once
)

func getExceptionHandlerClass() _ExceptionHandlerClass {
	ExceptionHandlerClassOnce.Do(func() {
		ExceptionHandlerClass = _ExceptionHandlerClass{objc.GetClass("NSExceptionHandler")}
	})
	return ExceptionHandlerClass
}

type _ExceptionHandlerClass struct {
	class objc.Class
}

// An interface definition for the [ExceptionHandler] class.
type IExceptionHandler interface {
	objectivec.IObject
	// properties:
	UserInfo() unsafe.Pointer
	SetUserInfo(value unsafe.Pointer)
	// methods:
	Delegate() objc.ID
	ExceptionHandlingMask() uint /* primitive/slice/pointer. */
	ExceptionHangingMask() uint /* primitive/slice/pointer. */
	SetDelegate(anObject objectivec.IObject)
	SetExceptionHandlingMask(aMask uint /* primitive/slice/pointer. */)
	SetExceptionHangingMask(aMask uint /* primitive/slice/pointer. */)
}

// The class provides facilities for monitoring and debugging exceptional conditions in Objective-C programs. It works by installing a special uncaught exception handler via the function. Consequently, to use the services of , you must not install your own custom uncaught exception handler.


// The class provides facilities for monitoring and debugging exceptional conditions in Objective-C programs. It works by installing a special uncaught exception handler via the function. Consequently, to use the services of , you must not install your own custom uncaught exception handler.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ExceptionHandling/NSExceptionHandler
type ExceptionHandler struct {
	objectivec.Object
}

// ExceptionHandlerFrom constructs a [ExceptionHandler] from an unsafe.Pointer.
//
// The class provides facilities for monitoring and debugging exceptional conditions in Objective-C programs. It works by installing a special uncaught exception handler via the function. Consequently, to use the services of , you must not install your own custom uncaught exception handler.
func ExceptionHandlerFrom(ptr unsafe.Pointer) ExceptionHandler {
	return ExceptionHandler{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ec _ExceptionHandlerClass) Alloc() ExceptionHandler {
	rv := objc.Send[ExceptionHandler](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ec _ExceptionHandlerClass) New() ExceptionHandler {
	rv := objc.Send[ExceptionHandler](objc.ID(ec.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (e_ ExceptionHandler) Init() ExceptionHandler {
	rv := objc.Send[ExceptionHandler](e_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e_ ExceptionHandler) Autorelease() ExceptionHandler {
	rv := objc.Send[ExceptionHandler](e_.ID, objc.Sel("autorelease"))
	return rv
}

// NewExceptionHandler creates a new ExceptionHandler instance.
func NewExceptionHandler() ExceptionHandler {
	return getExceptionHandlerClass().New()
}



// Returns the singleton instance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ExceptionHandling/NSExceptionHandler/default()
func (ec _ExceptionHandlerClass) DefaultExceptionHandler() IExceptionHandler {
	rv := objc.Send[ExceptionHandler](objc.ID(ec.class), objc.Sel("defaultExceptionHandler"))
	return rv
}


// Returns the delegate of the object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ExceptionHandling/NSExceptionHandler/delegate()
func (e_ ExceptionHandler) Delegate() objc.ID {
	rv := objc.Send[objc.ID](e_.ID, objc.Sel("delegate"))
	return rv
}


// Returns a bit mask representing the types of exceptions monitored by the receiver and its handling and logging behavior.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ExceptionHandling/NSExceptionHandler/exceptionHandlingMask()
func (e_ ExceptionHandler) ExceptionHandlingMask() uint /* primitive/slice/pointer. */ {
	rv := objc.Send[uint](e_.ID, objc.Sel("exceptionHandlingMask"))
	return rv
}


// Returns a bit mask representing the types of exceptions that will halt execution for debugging.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ExceptionHandling/NSExceptionHandler/exceptionHangingMask()
func (e_ ExceptionHandler) ExceptionHangingMask() uint /* primitive/slice/pointer. */ {
	rv := objc.Send[uint](e_.ID, objc.Sel("exceptionHangingMask"))
	return rv
}


// Sets the delegate of the object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ExceptionHandling/NSExceptionHandler/setDelegate(_:)
func (e_ ExceptionHandler) SetDelegate(anObject objectivec.IObject) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setDelegate:"), anObject)
}


// Sets the bit mask of constants specifying the types of exceptions monitored by the receiver and its handling and logging behavior.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ExceptionHandling/NSExceptionHandler/setExceptionHandlingMask(_:)
func (e_ ExceptionHandler) SetExceptionHandlingMask(aMask uint /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setExceptionHandlingMask:"), aMask)
}


// Sets the bit mask of constants specifying the types of exceptions that will halt execution for debugging.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ExceptionHandling/NSExceptionHandler/setExceptionHangingMask(_:)
func (e_ ExceptionHandler) SetExceptionHangingMask(aMask uint /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setExceptionHangingMask:"), aMask)
}


// A dictionary containing application-specific data pertaining to the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSException/userInfo-swift.property
func (e_ ExceptionHandler) UserInfo() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](e_.ID, objc.Sel("userInfo"))
	return rv
}


// A dictionary containing application-specific data pertaining to the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSException/userInfo-swift.property
func (e_ ExceptionHandler) SetUserInfo(value unsafe.Pointer) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setUserInfo:"), value)
}




