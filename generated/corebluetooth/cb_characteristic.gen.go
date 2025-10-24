// Code generated from Apple documentation for CoreBluetooth. DO NOT EDIT.

package corebluetooth

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class CBCharacteristic */


/* debug [class_header]: Header for CBCharacteristic */
// The class instance for the [CBCharacteristic] class.
var (
	CBCharacteristicClass     _CBCharacteristicClass
	CBCharacteristicClassOnce sync.Once
)

func getCBCharacteristicClass() _CBCharacteristicClass {
	CBCharacteristicClassOnce.Do(func() {
		CBCharacteristicClass = _CBCharacteristicClass{objc.GetClass("CBCharacteristic")}
	})
	return CBCharacteristicClass
}

type _CBCharacteristicClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CBCharacteristic */
// An interface definition for the [CBCharacteristic] class.
type ICBCharacteristic interface {
	ICBAttribute
	
/* debug [class_interface_properties]: Properties for CBCharacteristic */
	// properties:
	Descriptors() []CBDescriptor
	IsBroadcasted() bool
	IsNotifying() bool
	Properties() CBCharacteristicProperties
	Service() ICBService
	Value() objc.IObject /* cross-framework: NSData */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CBCharacteristic */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CBCharacteristic */
// Alloc allocates a new instance without initialization.
func (cc _CBCharacteristicClass) Alloc() CBCharacteristic {
	rv := objc.Send[CBCharacteristic](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CBCharacteristicClass) New() CBCharacteristic {
	rv := objc.Send[CBCharacteristic](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CBCharacteristic) Init() CBCharacteristic {
	rv := objc.Send[CBCharacteristic](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CBCharacteristic) Autorelease() CBCharacteristic {
	rv := objc.Send[CBCharacteristic](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCBCharacteristic creates a new CBCharacteristic instance.
func NewCBCharacteristic() CBCharacteristic {
	return getCBCharacteristicClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CBCharacteristic */
// A characteristic of a remote peripheral’s service.
//
// and its subclass represent further information about a peripheral’s service. In particular, objects represent the characteristics of a remote peripheral’s service. A characteristic contains a single value and any number of descriptors describing that value. The properties of a characteristic determine how you can use a characteristic’s value, and how you access the descriptors.


// A characteristic of a remote peripheral’s service.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreBluetooth/CBCharacteristic
type CBCharacteristic struct {
	CBAttribute
}

// CBCharacteristicFrom constructs a [CBCharacteristic] from an unsafe.Pointer.
//
// A characteristic of a remote peripheral’s service.
func CBCharacteristicFrom(ptr unsafe.Pointer) CBCharacteristic {
	return CBCharacteristic{
		CBAttribute: CBAttributeFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CBCharacteristic *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CBCharacteristic */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CBCharacteristic */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CBCharacteristic */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CBCharacteristic */

// A list of the descriptors discovered in this characteristic.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreBluetooth/CBCharacteristic/descriptors
func (c_ CBCharacteristic) Descriptors() []CBDescriptor {
	rv := objc.Send[[]CBDescriptor](c_.ID, objc.Sel("descriptors"))
	return rv
}/* debug [instance_properties/getter]: descriptors */


// A Boolean value that indicates whether the characteristic the service broadcasts this characteristic.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreBluetooth/CBCharacteristic/isBroadcasted
func (c_ CBCharacteristic) IsBroadcasted() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isBroadcasted"))
	return rv
}/* debug [instance_properties/getter]: isBroadcasted */


// A Boolean value that indicates whether the characteristic is currently notifying a subscribed central of its value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreBluetooth/CBCharacteristic/isNotifying
func (c_ CBCharacteristic) IsNotifying() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isNotifying"))
	return rv
}/* debug [instance_properties/getter]: isNotifying */


// The properties of the characteristic.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreBluetooth/CBCharacteristic/properties
func (c_ CBCharacteristic) Properties() CBCharacteristicProperties {
	rv := objc.Send[CBCharacteristicProperties](c_.ID, objc.Sel("properties"))
	return rv
}/* debug [instance_properties/getter]: properties */


// The service to which this characteristic belongs.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreBluetooth/CBCharacteristic/service
func (c_ CBCharacteristic) Service() ICBService {
	rv := objc.Send[CBService](c_.ID, objc.Sel("service"))
	return rv
}/* debug [instance_properties/getter]: service */


// The value of the characteristic.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreBluetooth/CBCharacteristic/value
func (c_ CBCharacteristic) Value() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](c_.ID, objc.Sel("value"))
	return rv
}/* debug [instance_properties/getter]: value */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CBCharacteristic */



