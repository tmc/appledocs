// Code generated from Apple documentation for ExceptionHandling. DO NOT EDIT.

package exceptionhandling

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSExceptionHandler */


/* debug [class_header]: Header for NSExceptionHandler */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ExceptionHandler */
// An interface definition for the [ExceptionHandler] class.
type IExceptionHandler interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for ExceptionHandler */
	// properties:
	UserInfo() unsafe.Pointer
	SetUserInfo(value unsafe.Pointer)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for ExceptionHandler */
	// methods:
	Delegate() objc.ID
	ExceptionHandlingMask() uint
	ExceptionHangingMask() uint
	SetDelegate(anObject objc.IObject)
	SetExceptionHandlingMask(aMask uint)
	SetExceptionHangingMask(aMask uint)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ExceptionHandler */
// Alloc allocates a new instance without initialization.
func (ec _ExceptionHandlerClass) Alloc() ExceptionHandler {
	rv := objc.Send[ExceptionHandler](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ExceptionHandler */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ExceptionHandler *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ExceptionHandler */

// Returns the singleton instance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ExceptionHandling/NSExceptionHandler/default()
func (ec _ExceptionHandlerClass) DefaultExceptionHandler() IExceptionHandler {
	rv := objc.Send[ExceptionHandler](objc.ID(ec.class), objc.Sel("defaultExceptionHandler"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=DefaultExceptionHandler) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ExceptionHandler */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ExceptionHandler */

// Returns the delegate of the object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ExceptionHandling/NSExceptionHandler/delegate()
func (e_ ExceptionHandler) Delegate() objc.ID {
	rv := objc.Send[objc.ID](e_.ID, objc.Sel("delegate"))
	return rv
}/* debug [instance_methods/method]: Delegate */


// Returns a bit mask representing the types of exceptions monitored by the receiver and its handling and logging behavior.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ExceptionHandling/NSExceptionHandler/exceptionHandlingMask()
func (e_ ExceptionHandler) ExceptionHandlingMask() uint {
	rv := objc.Send[uint](e_.ID, objc.Sel("exceptionHandlingMask"))
	return rv
}/* debug [instance_methods/method]: ExceptionHandlingMask */


// Returns a bit mask representing the types of exceptions that will halt execution for debugging.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ExceptionHandling/NSExceptionHandler/exceptionHangingMask()
func (e_ ExceptionHandler) ExceptionHangingMask() uint {
	rv := objc.Send[uint](e_.ID, objc.Sel("exceptionHangingMask"))
	return rv
}/* debug [instance_methods/method]: ExceptionHangingMask */


// Sets the delegate of the object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ExceptionHandling/NSExceptionHandler/setDelegate(_:)
func (e_ ExceptionHandler) SetDelegate(anObject objc.IObject) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setDelegate:"), anObject)
}/* debug [instance_methods/method]: SetDelegate */


// Sets the bit mask of constants specifying the types of exceptions monitored by the receiver and its handling and logging behavior.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ExceptionHandling/NSExceptionHandler/setExceptionHandlingMask(_:)
func (e_ ExceptionHandler) SetExceptionHandlingMask(aMask uint) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setExceptionHandlingMask:"), aMask)
}/* debug [instance_methods/method]: SetExceptionHandlingMask */


// Sets the bit mask of constants specifying the types of exceptions that will halt execution for debugging.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ExceptionHandling/NSExceptionHandler/setExceptionHangingMask(_:)
func (e_ ExceptionHandler) SetExceptionHangingMask(aMask uint) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setExceptionHangingMask:"), aMask)
}/* debug [instance_methods/method]: SetExceptionHangingMask */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ExceptionHandler */

// A dictionary containing application-specific data pertaining to the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSException/userInfo-swift.property
func (e_ ExceptionHandler) UserInfo() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](e_.ID, objc.Sel("userInfo"))
	return rv
}/* debug [instance_properties/getter]: userInfo */


// A dictionary containing application-specific data pertaining to the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSException/userInfo-swift.property
func (e_ ExceptionHandler) SetUserInfo(value unsafe.Pointer) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setUserInfo:"), value)
}/* debug [instance_properties/setter]: userInfo */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSExceptionHandler */



