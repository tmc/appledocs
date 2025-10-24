// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSException */


/* debug [class_header]: Header for NSException */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for Exception */
// An interface definition for the [Exception] class.
type IException interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for Exception */
	// properties:
	CallStackReturnAddresses() INumber
	SetCallStackReturnAddresses(value INumber)
	CallStackSymbols() IString
	SetCallStackSymbols(value IString)
	Name() ExceptionName /* typedef */
	SetName(value ExceptionName /* typedef */)
	Reason() IString
	SetReason(value IString)
	UserInfo() objectivec.IObject
	SetUserInfo(value objectivec.IObject)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for Exception */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for Exception */
// Alloc allocates a new instance without initialization.
func (ec _ExceptionClass) Alloc() Exception {
	rv := objc.Send[Exception](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for Exception */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for Exception *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for Exception */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for Exception */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for Exception */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for Exception */

// The call return addresses related to a raised exception.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsexception/callstackreturnaddresses
func (e_ Exception) CallStackReturnAddresses() INumber {
	rv := objc.Send[Number](e_.ID, objc.Sel("callStackReturnAddresses"))
	return rv
}/* debug [instance_properties/getter]: callStackReturnAddresses */


// The call return addresses related to a raised exception.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsexception/callstackreturnaddresses
func (e_ Exception) SetCallStackReturnAddresses(value INumber) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setCallStackReturnAddresses:"), value)
}/* debug [instance_properties/setter]: callStackReturnAddresses */


// An array containing the current call stack symbols.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsexception/callstacksymbols
func (e_ Exception) CallStackSymbols() IString {
	rv := objc.Send[String](e_.ID, objc.Sel("callStackSymbols"))
	return rv
}/* debug [instance_properties/getter]: callStackSymbols */


// An array containing the current call stack symbols.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsexception/callstacksymbols
func (e_ Exception) SetCallStackSymbols(value IString) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setCallStackSymbols:"), value)
}/* debug [instance_properties/setter]: callStackSymbols */


// A string used to uniquely identify the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsexception/name-swift.property
func (e_ Exception) Name() ExceptionName /* typedef */ {
	rv := objc.Send[String](e_.ID, objc.Sel("name"))
	return rv
}/* debug [instance_properties/getter]: name */


// A string used to uniquely identify the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsexception/name-swift.property
func (e_ Exception) SetName(value ExceptionName /* typedef */) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setName:"), value)
}/* debug [instance_properties/setter]: name */


// A string containing a “human-readable” reason for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsexception/reason-swift.property
func (e_ Exception) Reason() IString {
	rv := objc.Send[String](e_.ID, objc.Sel("reason"))
	return rv
}/* debug [instance_properties/getter]: reason */


// A string containing a “human-readable” reason for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsexception/reason-swift.property
func (e_ Exception) SetReason(value IString) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setReason:"), value)
}/* debug [instance_properties/setter]: reason */


// A dictionary containing application-specific data pertaining to the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsexception/userinfo-swift.property
func (e_ Exception) UserInfo() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](e_.ID, objc.Sel("userInfo"))
	return rv
}/* debug [instance_properties/getter]: userInfo */


// A dictionary containing application-specific data pertaining to the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsexception/userinfo-swift.property
func (e_ Exception) SetUserInfo(value objectivec.IObject) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setUserInfo:"), value)
}/* debug [instance_properties/setter]: userInfo */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSException */



