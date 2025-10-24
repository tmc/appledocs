// Code generated from Apple documentation for CoreBluetooth. DO NOT EDIT.

package corebluetooth

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class CBATTRequest */


/* debug [class_header]: Header for CBATTRequest */
// The class instance for the [CBATTRequest] class.
var (
	CBATTRequestClass     _CBATTRequestClass
	CBATTRequestClassOnce sync.Once
)

func getCBATTRequestClass() _CBATTRequestClass {
	CBATTRequestClassOnce.Do(func() {
		CBATTRequestClass = _CBATTRequestClass{objc.GetClass("CBATTRequest")}
	})
	return CBATTRequestClass
}

type _CBATTRequestClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CBATTRequest */
// An interface definition for the [CBATTRequest] class.
type ICBATTRequest interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for CBATTRequest */
	// properties:
	Central() ICBCentral
	Characteristic() ICBCharacteristic
	Offset() uint
	Value() objc.IObject /* cross-framework: NSData */
	SetValue(value objc.IObject /* cross-framework: NSData */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CBATTRequest */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CBATTRequest */
// Alloc allocates a new instance without initialization.
func (cc _CBATTRequestClass) Alloc() CBATTRequest {
	rv := objc.Send[CBATTRequest](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CBATTRequestClass) New() CBATTRequest {
	rv := objc.Send[CBATTRequest](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CBATTRequest) Init() CBATTRequest {
	rv := objc.Send[CBATTRequest](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CBATTRequest) Autorelease() CBATTRequest {
	rv := objc.Send[CBATTRequest](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCBATTRequest creates a new CBATTRequest instance.
func NewCBATTRequest() CBATTRequest {
	return getCBATTRequestClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CBATTRequest */
// A request that uses the Attribute Protocol (ATT).
//
// The class represents Attribute Protocol (ATT) read and write requests from remote central devices (represented by objects). Remote centrals use these ATT requests to read and write characteristic values on local peripherals (represented by objects). Local peripherals, on the other hand, use the properties of objects to respond to the read and write requests appropriately, using the method of the class.


// A request that uses the Attribute Protocol (ATT).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreBluetooth/CBATTRequest
type CBATTRequest struct {
	objectivec.Object
}

// CBATTRequestFrom constructs a [CBATTRequest] from an unsafe.Pointer.
//
// A request that uses the Attribute Protocol (ATT).
func CBATTRequestFrom(ptr unsafe.Pointer) CBATTRequest {
	return CBATTRequest{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CBATTRequest *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CBATTRequest */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CBATTRequest */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CBATTRequest */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CBATTRequest */

// The remote central device that originated the request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreBluetooth/CBATTRequest/central
func (c_ CBATTRequest) Central() ICBCentral {
	rv := objc.Send[CBCentral](c_.ID, objc.Sel("central"))
	return rv
}/* debug [instance_properties/getter]: central */


// The characteristic to read or write the value of.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreBluetooth/CBATTRequest/characteristic
func (c_ CBATTRequest) Characteristic() ICBCharacteristic {
	rv := objc.Send[CBCharacteristic](c_.ID, objc.Sel("characteristic"))
	return rv
}/* debug [instance_properties/getter]: characteristic */


// The zero-based index of the first byte for the read or write request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreBluetooth/CBATTRequest/offset
func (c_ CBATTRequest) Offset() uint {
	rv := objc.Send[uint](c_.ID, objc.Sel("offset"))
	return rv
}/* debug [instance_properties/getter]: offset */


// The data that the central reads from or writes to the peripheral.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreBluetooth/CBATTRequest/value
func (c_ CBATTRequest) Value() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](c_.ID, objc.Sel("value"))
	return rv
}/* debug [instance_properties/getter]: value */


// The data that the central reads from or writes to the peripheral.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreBluetooth/CBATTRequest/value
func (c_ CBATTRequest) SetValue(value objc.IObject /* cross-framework: NSData */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setValue:"), value)
}/* debug [instance_properties/setter]: value */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CBATTRequest */



