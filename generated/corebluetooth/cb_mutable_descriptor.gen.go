// Code generated from Apple documentation for CoreBluetooth. DO NOT EDIT.

package corebluetooth

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class CBMutableDescriptor */


/* debug [class_header]: Header for CBMutableDescriptor */
// The class instance for the [CBMutableDescriptor] class.
var (
	CBMutableDescriptorClass     _CBMutableDescriptorClass
	CBMutableDescriptorClassOnce sync.Once
)

func getCBMutableDescriptorClass() _CBMutableDescriptorClass {
	CBMutableDescriptorClassOnce.Do(func() {
		CBMutableDescriptorClass = _CBMutableDescriptorClass{objc.GetClass("CBMutableDescriptor")}
	})
	return CBMutableDescriptorClass
}

type _CBMutableDescriptorClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CBMutableDescriptor */
// An interface definition for the [CBMutableDescriptor] class.
type ICBMutableDescriptor interface {
	ICBDescriptor
	
/* debug [class_interface_properties]: Properties for CBMutableDescriptor */
	// properties:
	CBUUIDCharacteristicFormatString() objc.IObject /* cross-framework: NSString */
	CBUUIDCharacteristicUserDescriptionString() objc.IObject /* cross-framework: NSString */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CBMutableDescriptor */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CBMutableDescriptor */
// Alloc allocates a new instance without initialization.
func (cc _CBMutableDescriptorClass) Alloc() CBMutableDescriptor {
	rv := objc.Send[CBMutableDescriptor](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CBMutableDescriptorClass) New() CBMutableDescriptor {
	rv := objc.Send[CBMutableDescriptor](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CBMutableDescriptor) Init() CBMutableDescriptor {
	rv := objc.Send[CBMutableDescriptor](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CBMutableDescriptor) Autorelease() CBMutableDescriptor {
	rv := objc.Send[CBMutableDescriptor](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCBMutableDescriptor creates a new CBMutableDescriptor instance.
func NewCBMutableDescriptor() CBMutableDescriptor {
	return getCBMutableDescriptorClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CBMutableDescriptor */
// An object that provides additional information about a local peripheral’s characteristic.
//
// You use the class to create a local characteristic descriptor. After you create a descriptor and associate it with a local characteristic, you can publish it to the peripheral’s local database using the method of the class. This also publishes the characteristic and local service to which the descriptor belongs. After you publish a local descriptor, Core Bluetooth caches the descriptor and you can no longer make changes to it. details predefined descriptor types and their corresponding value types. That said, only two of these are currently supported when creating local, mutable descriptors: the characteristic user description descriptor and the characteristic format descriptor. declares these as the constants and , respectively. The system automatically creates the extended properties descriptor and the client configuration descriptor, depending on the properties of the characteristic to which the descriptor belongs.


// An object that provides additional information about a local peripheral’s characteristic.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreBluetooth/CBMutableDescriptor
type CBMutableDescriptor struct {
	CBDescriptor
}

// CBMutableDescriptorFrom constructs a [CBMutableDescriptor] from an unsafe.Pointer.
//
// An object that provides additional information about a local peripheral’s characteristic.
func CBMutableDescriptorFrom(ptr unsafe.Pointer) CBMutableDescriptor {
	return CBMutableDescriptor{
		CBDescriptor: CBDescriptorFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CBMutableDescriptor */

// Creates a mutable descriptor with a specified value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreBluetooth/CBMutableDescriptor/init(type:value:)
func NewCBMutableDescriptorWithTypeValue(UUID ICBUUID, value objc.IObject) CBMutableDescriptor {
	instance := getCBMutableDescriptorClass().Alloc()
	rv := objc.Send[CBMutableDescriptor](instance.ID, objc.Sel("initWithType:value:"), UUID, value)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCBMutableDescriptorWithTypeValue */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CBMutableDescriptor */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CBMutableDescriptor */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CBMutableDescriptor */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CBMutableDescriptor */

// The UUID for the Presentation Format descriptor, as a string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corebluetooth/cbuuidcharacteristicformatstring
func (c_ CBMutableDescriptor) CBUUIDCharacteristicFormatString() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CBUUIDCharacteristicFormatString"))
	return rv
}/* debug [instance_properties/getter]: CBUUIDCharacteristicFormatString */


// The UUID for the User Description descriptor, as a string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corebluetooth/cbuuidcharacteristicuserdescriptionstring
func (c_ CBMutableDescriptor) CBUUIDCharacteristicUserDescriptionString() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CBUUIDCharacteristicUserDescriptionString"))
	return rv
}/* debug [instance_properties/getter]: CBUUIDCharacteristicUserDescriptionString */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CBMutableDescriptor */


