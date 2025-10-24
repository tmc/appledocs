// Code generated from Apple documentation for CryptoTokenKit. DO NOT EDIT.

package cryptotokenkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class TKSmartCardUserInteractionForPINOperation */


/* debug [class_header]: Header for TKSmartCardUserInteractionForPINOperation */
// The class instance for the [TKSmartCardUserInteractionForPINOperation] class.
var (
	TKSmartCardUserInteractionForPINOperationClass     _TKSmartCardUserInteractionForPINOperationClass
	TKSmartCardUserInteractionForPINOperationClassOnce sync.Once
)

func getTKSmartCardUserInteractionForPINOperationClass() _TKSmartCardUserInteractionForPINOperationClass {
	TKSmartCardUserInteractionForPINOperationClassOnce.Do(func() {
		TKSmartCardUserInteractionForPINOperationClass = _TKSmartCardUserInteractionForPINOperationClass{objc.GetClass("TKSmartCardUserInteractionForPINOperation")}
	})
	return TKSmartCardUserInteractionForPINOperationClass
}

type _TKSmartCardUserInteractionForPINOperationClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for TKSmartCardUserInteractionForPINOperation */
// An interface definition for the [TKSmartCardUserInteractionForPINOperation] class.
type ITKSmartCardUserInteractionForPINOperation interface {
	ITKSmartCardUserInteraction
	
/* debug [class_interface_properties]: Properties for TKSmartCardUserInteractionForPINOperation */
	// properties:
	Locale() foundation.Locale
	SetLocale(value foundation.Locale)
	PINCompletion() TKSmartCardPINCompletion
	SetPINCompletion(value TKSmartCardPINCompletion)
	PINMessageIndices() []foundation.Number
	SetPINMessageIndices(value []foundation.Number)
	ResultData() objc.IObject /* cross-framework: NSData */
	SetResultData(value objc.IObject /* cross-framework: NSData */)
	ResultSW() unsafe.Pointer
	SetResultSW(value unsafe.Pointer)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for TKSmartCardUserInteractionForPINOperation */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for TKSmartCardUserInteractionForPINOperation */
// Alloc allocates a new instance without initialization.
func (tc _TKSmartCardUserInteractionForPINOperationClass) Alloc() TKSmartCardUserInteractionForPINOperation {
	rv := objc.Send[TKSmartCardUserInteractionForPINOperation](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (tc _TKSmartCardUserInteractionForPINOperationClass) New() TKSmartCardUserInteractionForPINOperation {
	rv := objc.Send[TKSmartCardUserInteractionForPINOperation](objc.ID(tc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (t_ TKSmartCardUserInteractionForPINOperation) Init() TKSmartCardUserInteractionForPINOperation {
	rv := objc.Send[TKSmartCardUserInteractionForPINOperation](t_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t_ TKSmartCardUserInteractionForPINOperation) Autorelease() TKSmartCardUserInteractionForPINOperation {
	rv := objc.Send[TKSmartCardUserInteractionForPINOperation](t_.ID, objc.Sel("autorelease"))
	return rv
}

// NewTKSmartCardUserInteractionForPINOperation creates a new TKSmartCardUserInteractionForPINOperation instance.
func NewTKSmartCardUserInteractionForPINOperation() TKSmartCardUserInteractionForPINOperation {
	return getTKSmartCardUserInteractionForPINOperationClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for TKSmartCardUserInteractionForPINOperation */
// A representation of user interaction for secure PIN operations on a Smart Card reader.
//
// There are two types of user interactions: those for secure PIN change and those for secure PIN validation. These interactions are instances of the , or subclasses of , respectively. You interact with instances of one of the subclasses of when calling the and methods on an object. The result of a user interaction is available once the interaction has completed.


// A representation of user interaction for secure PIN operations on a Smart Card reader.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCardUserInteractionForPINOperation
type TKSmartCardUserInteractionForPINOperation struct {
	TKSmartCardUserInteraction
}

// TKSmartCardUserInteractionForPINOperationFrom constructs a [TKSmartCardUserInteractionForPINOperation] from an unsafe.Pointer.
//
// A representation of user interaction for secure PIN operations on a Smart Card reader.
func TKSmartCardUserInteractionForPINOperationFrom(ptr unsafe.Pointer) TKSmartCardUserInteractionForPINOperation {
	return TKSmartCardUserInteractionForPINOperation{
		TKSmartCardUserInteraction: TKSmartCardUserInteractionFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for TKSmartCardUserInteractionForPINOperation *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for TKSmartCardUserInteractionForPINOperation */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for TKSmartCardUserInteractionForPINOperation */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for TKSmartCardUserInteractionForPINOperation */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for TKSmartCardUserInteractionForPINOperation */

// The locale for the displayed messages. If , the user’s current locale is used. By default, this value is the current locale of the system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCardUserInteractionForPINOperation/locale
func (t_ TKSmartCardUserInteractionForPINOperation) Locale() foundation.Locale {
	rv := objc.Send[foundation.Locale](t_.ID, objc.Sel("locale"))
	return rv
}/* debug [instance_properties/getter]: locale */


// The locale for the displayed messages. If , the user’s current locale is used. By default, this value is the current locale of the system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCardUserInteractionForPINOperation/locale
func (t_ TKSmartCardUserInteractionForPINOperation) SetLocale(value foundation.Locale) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setLocale:"), value)
}/* debug [instance_properties/setter]: locale */


// The conditions under which PIN entry should be considered complete.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCardUserInteractionForPINOperation/pinCompletion
func (t_ TKSmartCardUserInteractionForPINOperation) PINCompletion() TKSmartCardPINCompletion {
	rv := objc.Send[TKSmartCardPINCompletion](t_.ID, objc.Sel("PINCompletion"))
	return rv
}/* debug [instance_properties/getter]: PINCompletion */


// The conditions under which PIN entry should be considered complete.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCardUserInteractionForPINOperation/pinCompletion
func (t_ TKSmartCardUserInteractionForPINOperation) SetPINCompletion(value TKSmartCardPINCompletion) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setPINCompletion:"), value)
}/* debug [instance_properties/setter]: PINCompletion */


// A list of message indices referring to a predefined message table, used to specify the type and number of messages displayed during the PIN operation. by default.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCardUserInteractionForPINOperation/pinMessageIndices
func (t_ TKSmartCardUserInteractionForPINOperation) PINMessageIndices() []foundation.Number {
	rv := objc.Send[[]foundation.Number](t_.ID, objc.Sel("PINMessageIndices"))
	return rv
}/* debug [instance_properties/getter]: PINMessageIndices */


// A list of message indices referring to a predefined message table, used to specify the type and number of messages displayed during the PIN operation. by default.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCardUserInteractionForPINOperation/pinMessageIndices
func (t_ TKSmartCardUserInteractionForPINOperation) SetPINMessageIndices(value []foundation.Number) {
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](t_.ID, objc.Sel("setPINMessageIndices:"), nsArray)
}/* debug [instance_properties/setter]: PINMessageIndices */


// The returned data without SW1-SW2 bytes, if any.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCardUserInteractionForPINOperation/resultData
func (t_ TKSmartCardUserInteractionForPINOperation) ResultData() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](t_.ID, objc.Sel("resultData"))
	return rv
}/* debug [instance_properties/getter]: resultData */


// The returned data without SW1-SW2 bytes, if any.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCardUserInteractionForPINOperation/resultData
func (t_ TKSmartCardUserInteractionForPINOperation) SetResultData(value objc.IObject /* cross-framework: NSData */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setResultData:"), value)
}/* debug [instance_properties/setter]: resultData */


// The SW1-SW2 status bytes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCardUserInteractionForPINOperation/resultSW
func (t_ TKSmartCardUserInteractionForPINOperation) ResultSW() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("resultSW"))
	return rv
}/* debug [instance_properties/getter]: resultSW */


// The SW1-SW2 status bytes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCardUserInteractionForPINOperation/resultSW
func (t_ TKSmartCardUserInteractionForPINOperation) SetResultSW(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setResultSW:"), value)
}/* debug [instance_properties/setter]: resultSW */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class TKSmartCardUserInteractionForPINOperation */



