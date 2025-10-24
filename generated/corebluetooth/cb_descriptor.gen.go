// Code generated from Apple documentation for CoreBluetooth. DO NOT EDIT.

package corebluetooth

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class CBDescriptor */


/* debug [class_header]: Header for CBDescriptor */
// The class instance for the [CBDescriptor] class.
var (
	CBDescriptorClass     _CBDescriptorClass
	CBDescriptorClassOnce sync.Once
)

func getCBDescriptorClass() _CBDescriptorClass {
	CBDescriptorClassOnce.Do(func() {
		CBDescriptorClass = _CBDescriptorClass{objc.GetClass("CBDescriptor")}
	})
	return CBDescriptorClass
}

type _CBDescriptorClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CBDescriptor */
// An interface definition for the [CBDescriptor] class.
type ICBDescriptor interface {
	ICBAttribute
	
/* debug [class_interface_properties]: Properties for CBDescriptor */
	// properties:
	Characteristic() ICBCharacteristic
	Value() objc.ID
	CBUUIDCharacteristicAggregateFormatString() objc.IObject /* cross-framework: NSString */
	CBUUIDCharacteristicExtendedPropertiesString() objc.IObject /* cross-framework: NSString */
	CBUUIDCharacteristicFormatString() objc.IObject /* cross-framework: NSString */
	CBUUIDCharacteristicUserDescriptionString() objc.IObject /* cross-framework: NSString */
	CBUUIDClientCharacteristicConfigurationString() objc.IObject /* cross-framework: NSString */
	CBUUIDServerCharacteristicConfigurationString() objc.IObject /* cross-framework: NSString */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CBDescriptor */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CBDescriptor */
// Alloc allocates a new instance without initialization.
func (cc _CBDescriptorClass) Alloc() CBDescriptor {
	rv := objc.Send[CBDescriptor](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CBDescriptorClass) New() CBDescriptor {
	rv := objc.Send[CBDescriptor](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CBDescriptor) Init() CBDescriptor {
	rv := objc.Send[CBDescriptor](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CBDescriptor) Autorelease() CBDescriptor {
	rv := objc.Send[CBDescriptor](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCBDescriptor creates a new CBDescriptor instance.
func NewCBDescriptor() CBDescriptor {
	return getCBDescriptorClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CBDescriptor */
// An object that provides further information about a remote peripheral’s characteristic.
//
// and its subclass represent a descriptor of a peripheral’s characteristic. In partcular, objects represent the descriptors of a remote peripheral’s characteristic. Descriptors provide further information about a characteristic’s value. For example, they may describe the value in human-readable form and describe how to format the value for presentation purposes. Characteristic descriptors also indicate whether a characteristic’s value indicates or notifies a client (a central) when the value of the characteristic changes. details six predefined descriptors and their corresponding value types. lists the predefined descriptors and the constants that represent them.


// An object that provides further information about a remote peripheral’s characteristic.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreBluetooth/CBDescriptor
type CBDescriptor struct {
	CBAttribute
}

// CBDescriptorFrom constructs a [CBDescriptor] from an unsafe.Pointer.
//
// An object that provides further information about a remote peripheral’s characteristic.
func CBDescriptorFrom(ptr unsafe.Pointer) CBDescriptor {
	return CBDescriptor{
		CBAttribute: CBAttributeFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CBDescriptor *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CBDescriptor */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CBDescriptor */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CBDescriptor */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CBDescriptor */

// The characteristic to which this descriptor belongs.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreBluetooth/CBDescriptor/characteristic
func (c_ CBDescriptor) Characteristic() ICBCharacteristic {
	rv := objc.Send[CBCharacteristic](c_.ID, objc.Sel("characteristic"))
	return rv
}/* debug [instance_properties/getter]: characteristic */


// The value of the descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreBluetooth/CBDescriptor/value
func (c_ CBDescriptor) Value() objc.ID {
	rv := objc.Send[objc.ID](c_.ID, objc.Sel("value"))
	return rv
}/* debug [instance_properties/getter]: value */


// The UUID for the Aggregate Format descriptor, as a string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corebluetooth/cbuuidcharacteristicaggregateformatstring
func (c_ CBDescriptor) CBUUIDCharacteristicAggregateFormatString() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CBUUIDCharacteristicAggregateFormatString"))
	return rv
}/* debug [instance_properties/getter]: CBUUIDCharacteristicAggregateFormatString */


// The UUID for the Extended Properties descriptor, as a string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corebluetooth/cbuuidcharacteristicextendedpropertiesstring
func (c_ CBDescriptor) CBUUIDCharacteristicExtendedPropertiesString() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CBUUIDCharacteristicExtendedPropertiesString"))
	return rv
}/* debug [instance_properties/getter]: CBUUIDCharacteristicExtendedPropertiesString */


// The UUID for the Presentation Format descriptor, as a string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corebluetooth/cbuuidcharacteristicformatstring
func (c_ CBDescriptor) CBUUIDCharacteristicFormatString() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CBUUIDCharacteristicFormatString"))
	return rv
}/* debug [instance_properties/getter]: CBUUIDCharacteristicFormatString */


// The UUID for the User Description descriptor, as a string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corebluetooth/cbuuidcharacteristicuserdescriptionstring
func (c_ CBDescriptor) CBUUIDCharacteristicUserDescriptionString() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CBUUIDCharacteristicUserDescriptionString"))
	return rv
}/* debug [instance_properties/getter]: CBUUIDCharacteristicUserDescriptionString */


// The UUID for the Client Configuration descriptor, as a string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corebluetooth/cbuuidclientcharacteristicconfigurationstring
func (c_ CBDescriptor) CBUUIDClientCharacteristicConfigurationString() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CBUUIDClientCharacteristicConfigurationString"))
	return rv
}/* debug [instance_properties/getter]: CBUUIDClientCharacteristicConfigurationString */


// The UUID for the Server Configuration descriptor, as a string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corebluetooth/cbuuidservercharacteristicconfigurationstring
func (c_ CBDescriptor) CBUUIDServerCharacteristicConfigurationString() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CBUUIDServerCharacteristicConfigurationString"))
	return rv
}/* debug [instance_properties/getter]: CBUUIDServerCharacteristicConfigurationString */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CBDescriptor */



