// Code generated from Apple documentation for CryptoTokenKit. DO NOT EDIT.

package cryptotokenkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class TKSmartCardSlot */


/* debug [class_header]: Header for TKSmartCardSlot */
// The class instance for the [TKSmartCardSlot] class.
var (
	TKSmartCardSlotClass     _TKSmartCardSlotClass
	TKSmartCardSlotClassOnce sync.Once
)

func getTKSmartCardSlotClass() _TKSmartCardSlotClass {
	TKSmartCardSlotClassOnce.Do(func() {
		TKSmartCardSlotClass = _TKSmartCardSlotClass{objc.GetClass("TKSmartCardSlot")}
	})
	return TKSmartCardSlotClass
}

type _TKSmartCardSlotClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for TKSmartCardSlot */
// An interface definition for the [TKSmartCardSlot] class.
type ITKSmartCardSlot interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for TKSmartCardSlot */
	// properties:
	ATR() ITKSmartCardATR
	MaxInputLength() int
	MaxOutputLength() int
	Name() objc.IObject /* cross-framework: NSString */
	State() TKSmartCardSlotState
	SlotNames() objc.IObject /* cross-framework: NSString */
	SetSlotNames(value objc.IObject /* cross-framework: NSString */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for TKSmartCardSlot */
	// methods:
	MakeSmartCard() ITKSmartCard
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for TKSmartCardSlot */
// Alloc allocates a new instance without initialization.
func (tc _TKSmartCardSlotClass) Alloc() TKSmartCardSlot {
	rv := objc.Send[TKSmartCardSlot](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (tc _TKSmartCardSlotClass) New() TKSmartCardSlot {
	rv := objc.Send[TKSmartCardSlot](objc.ID(tc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (t_ TKSmartCardSlot) Init() TKSmartCardSlot {
	rv := objc.Send[TKSmartCardSlot](t_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t_ TKSmartCardSlot) Autorelease() TKSmartCardSlot {
	rv := objc.Send[TKSmartCardSlot](t_.ID, objc.Sel("autorelease"))
	return rv
}

// NewTKSmartCardSlot creates a new TKSmartCardSlot instance.
func NewTKSmartCardSlot() TKSmartCardSlot {
	return getTKSmartCardSlotClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for TKSmartCardSlot */
// A single smart card reader slot in the system.
//
// Use the class to manage all the smart card reader slots available to the system. You can retrieve the names of available smart card reader slots for a system using the property of a manager object, and access instances of using the method.


// A single smart card reader slot in the system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCardSlot
type TKSmartCardSlot struct {
	objectivec.Object
}

// TKSmartCardSlotFrom constructs a [TKSmartCardSlot] from an unsafe.Pointer.
//
// A single smart card reader slot in the system.
func TKSmartCardSlotFrom(ptr unsafe.Pointer) TKSmartCardSlot {
	return TKSmartCardSlot{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for TKSmartCardSlot *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for TKSmartCardSlot */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for TKSmartCardSlot */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for TKSmartCardSlot */

// Creates a new object representing the currently inserted Smart Card.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCardSlot/makeSmartCard()
func (t_ TKSmartCardSlot) MakeSmartCard() ITKSmartCard {
	rv := objc.Send[TKSmartCard](t_.ID, objc.Sel("makeSmartCard"))
	return rv
}/* debug [instance_methods/method]: MakeSmartCard */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for TKSmartCardSlot */

// The ATR (Answer to Reset) of the inserted Smart Card, or if no Smart Card is inserted or the inserted Smart Card is mute.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCardSlot/atr
func (t_ TKSmartCardSlot) ATR() ITKSmartCardATR {
	rv := objc.Send[TKSmartCardATR](t_.ID, objc.Sel("ATR"))
	return rv
}/* debug [instance_properties/getter]: ATR */


// The maximum length of input APDU (Application Protocol Data Unit) that the Smart Card reader slot is able to transfer to the Smart Card.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCardSlot/maxInputLength
func (t_ TKSmartCardSlot) MaxInputLength() int {
	rv := objc.Send[int](t_.ID, objc.Sel("maxInputLength"))
	return rv
}/* debug [instance_properties/getter]: maxInputLength */


// The maximum length of output APDU (Application Protocol Data Unit) that the Smart Card reader slot is able to transfer from the Smart Card.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCardSlot/maxOutputLength
func (t_ TKSmartCardSlot) MaxOutputLength() int {
	rv := objc.Send[int](t_.ID, objc.Sel("maxOutputLength"))
	return rv
}/* debug [instance_properties/getter]: maxOutputLength */


// The name of the Smart Card reader slot.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCardSlot/name
func (t_ TKSmartCardSlot) Name() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](t_.ID, objc.Sel("name"))
	return rv
}/* debug [instance_properties/getter]: name */


// The current state of the Smart Card reader slot.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCardSlot/state-swift.property
func (t_ TKSmartCardSlot) State() TKSmartCardSlotState {
	rv := objc.Send[TKSmartCardSlotState](t_.ID, objc.Sel("state"))
	return rv
}/* debug [instance_properties/getter]: state */


// A list of identifiers for all the Smart Card reader slots available to the system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cryptotokenkit/tksmartcardslotmanager/slotnames
func (t_ TKSmartCardSlot) SlotNames() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](t_.ID, objc.Sel("slotNames"))
	return rv
}/* debug [instance_properties/getter]: slotNames */


// A list of identifiers for all the Smart Card reader slots available to the system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cryptotokenkit/tksmartcardslotmanager/slotnames
func (t_ TKSmartCardSlot) SetSlotNames(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setSlotNames:"), value)
}/* debug [instance_properties/setter]: slotNames */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class TKSmartCardSlot */



