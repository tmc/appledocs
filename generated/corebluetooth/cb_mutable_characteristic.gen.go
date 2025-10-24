// Code generated from Apple documentation for CoreBluetooth. DO NOT EDIT.

package corebluetooth

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class CBMutableCharacteristic */


/* debug [class_header]: Header for CBMutableCharacteristic */
// The class instance for the [CBMutableCharacteristic] class.
var (
	CBMutableCharacteristicClass     _CBMutableCharacteristicClass
	CBMutableCharacteristicClassOnce sync.Once
)

func getCBMutableCharacteristicClass() _CBMutableCharacteristicClass {
	CBMutableCharacteristicClassOnce.Do(func() {
		CBMutableCharacteristicClass = _CBMutableCharacteristicClass{objc.GetClass("CBMutableCharacteristic")}
	})
	return CBMutableCharacteristicClass
}

type _CBMutableCharacteristicClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CBMutableCharacteristic */
// An interface definition for the [CBMutableCharacteristic] class.
type ICBMutableCharacteristic interface {
	ICBCharacteristic
	
/* debug [class_interface_properties]: Properties for CBMutableCharacteristic */
	// properties:
	Descriptors() []CBDescriptor
	SetDescriptors(value []CBDescriptor)
	Permissions() CBAttributePermissions
	SetPermissions(value CBAttributePermissions)
	Properties() CBCharacteristicProperties
	SetProperties(value CBCharacteristicProperties)
	SubscribedCentrals() []CBCentral
	Value() objc.IObject /* cross-framework: NSData */
	SetValue(value objc.IObject /* cross-framework: NSData */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CBMutableCharacteristic */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CBMutableCharacteristic */
// Alloc allocates a new instance without initialization.
func (cc _CBMutableCharacteristicClass) Alloc() CBMutableCharacteristic {
	rv := objc.Send[CBMutableCharacteristic](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CBMutableCharacteristicClass) New() CBMutableCharacteristic {
	rv := objc.Send[CBMutableCharacteristic](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CBMutableCharacteristic) Init() CBMutableCharacteristic {
	rv := objc.Send[CBMutableCharacteristic](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CBMutableCharacteristic) Autorelease() CBMutableCharacteristic {
	rv := objc.Send[CBMutableCharacteristic](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCBMutableCharacteristic creates a new CBMutableCharacteristic instance.
func NewCBMutableCharacteristic() CBMutableCharacteristic {
	return getCBMutableCharacteristicClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CBMutableCharacteristic */
// A characteristic of a local peripheral’s service.
//
// objects represent the characteristics of a local peripheral’s service. This class adds write access to many of the properties in the class, which it inherits from. You use this class to create a characteristic and to set its properties and permissions as desired. After you create and add a characteristic to a local service, you can publish it (and the service) to the peripheral’s local database with the method of the class. After you publish a characteristic, Core Bluetooth caches the characteristic and you can’t make changes to it.


// A characteristic of a local peripheral’s service.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreBluetooth/CBMutableCharacteristic
type CBMutableCharacteristic struct {
	CBCharacteristic
}

// CBMutableCharacteristicFrom constructs a [CBMutableCharacteristic] from an unsafe.Pointer.
//
// A characteristic of a local peripheral’s service.
func CBMutableCharacteristicFrom(ptr unsafe.Pointer) CBMutableCharacteristic {
	return CBMutableCharacteristic{
		CBCharacteristic: CBCharacteristicFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CBMutableCharacteristic */

// Creates a mutable characteristic with specified permissions, properties, and value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreBluetooth/CBMutableCharacteristic/init(type:properties:value:permissions:)
func NewCBMutableCharacteristicWithTypePropertiesValuePermissions(UUID ICBUUID, properties CBCharacteristicProperties, value objc.IObject /* cross-framework: NSData */, permissions CBAttributePermissions) CBMutableCharacteristic {
	instance := getCBMutableCharacteristicClass().Alloc()
	rv := objc.Send[CBMutableCharacteristic](instance.ID, objc.Sel("initWithType:properties:value:permissions:"), UUID, properties, value, permissions)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCBMutableCharacteristicWithTypePropertiesValuePermissions */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CBMutableCharacteristic */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CBMutableCharacteristic */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CBMutableCharacteristic */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CBMutableCharacteristic */

// An array of the characteristic’s descriptors.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreBluetooth/CBMutableCharacteristic/descriptors
func (c_ CBMutableCharacteristic) Descriptors() []CBDescriptor {
	rv := objc.Send[[]CBDescriptor](c_.ID, objc.Sel("descriptors"))
	return rv
}/* debug [instance_properties/getter]: descriptors */


// An array of the characteristic’s descriptors.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreBluetooth/CBMutableCharacteristic/descriptors
func (c_ CBMutableCharacteristic) SetDescriptors(value []CBDescriptor) {
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](c_.ID, objc.Sel("setDescriptors:"), nsArray)
}/* debug [instance_properties/setter]: descriptors */


// The permissions of the characteristic value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreBluetooth/CBMutableCharacteristic/permissions
func (c_ CBMutableCharacteristic) Permissions() CBAttributePermissions {
	rv := objc.Send[CBAttributePermissions](c_.ID, objc.Sel("permissions"))
	return rv
}/* debug [instance_properties/getter]: permissions */


// The permissions of the characteristic value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreBluetooth/CBMutableCharacteristic/permissions
func (c_ CBMutableCharacteristic) SetPermissions(value CBAttributePermissions) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPermissions:"), value)
}/* debug [instance_properties/setter]: permissions */


// The properties of the characteristic.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreBluetooth/CBMutableCharacteristic/properties
func (c_ CBMutableCharacteristic) Properties() CBCharacteristicProperties {
	rv := objc.Send[CBCharacteristicProperties](c_.ID, objc.Sel("properties"))
	return rv
}/* debug [instance_properties/getter]: properties */


// The properties of the characteristic.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreBluetooth/CBMutableCharacteristic/properties
func (c_ CBMutableCharacteristic) SetProperties(value CBCharacteristicProperties) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setProperties:"), value)
}/* debug [instance_properties/setter]: properties */


// A list of centrals that are currently subscribed to the characteristic’s value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreBluetooth/CBMutableCharacteristic/subscribedCentrals
func (c_ CBMutableCharacteristic) SubscribedCentrals() []CBCentral {
	rv := objc.Send[[]CBCentral](c_.ID, objc.Sel("subscribedCentrals"))
	return rv
}/* debug [instance_properties/getter]: subscribedCentrals */


// The value of the characteristic.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreBluetooth/CBMutableCharacteristic/value
func (c_ CBMutableCharacteristic) Value() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](c_.ID, objc.Sel("value"))
	return rv
}/* debug [instance_properties/getter]: value */


// The value of the characteristic.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreBluetooth/CBMutableCharacteristic/value
func (c_ CBMutableCharacteristic) SetValue(value objc.IObject /* cross-framework: NSData */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setValue:"), value)
}/* debug [instance_properties/setter]: value */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CBMutableCharacteristic */


