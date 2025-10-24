// Code generated from Apple documentation for CoreBluetooth. DO NOT EDIT.

package corebluetooth

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class CBAttribute */


/* debug [class_header]: Header for CBAttribute */
// The class instance for the [CBAttribute] class.
var (
	CBAttributeClass     _CBAttributeClass
	CBAttributeClassOnce sync.Once
)

func getCBAttributeClass() _CBAttributeClass {
	CBAttributeClassOnce.Do(func() {
		CBAttributeClass = _CBAttributeClass{objc.GetClass("CBAttribute")}
	})
	return CBAttributeClass
}

type _CBAttributeClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CBAttribute */
// An interface definition for the [CBAttribute] class.
type ICBAttribute interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for CBAttribute */
	// properties:
	UUID() ICBUUID
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CBAttribute */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CBAttribute */
// Alloc allocates a new instance without initialization.
func (cc _CBAttributeClass) Alloc() CBAttribute {
	rv := objc.Send[CBAttribute](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CBAttributeClass) New() CBAttribute {
	rv := objc.Send[CBAttribute](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CBAttribute) Init() CBAttribute {
	rv := objc.Send[CBAttribute](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CBAttribute) Autorelease() CBAttribute {
	rv := objc.Send[CBAttribute](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCBAttribute creates a new CBAttribute instance.
func NewCBAttribute() CBAttribute {
	return getCBAttributeClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CBAttribute */
// A representation of common aspects of services offered by a peripheral.
//
// Concrete subclasses of (and their mutable counterparts) represent the services a peripheral offers, the characteristics of those services, and the descriptors attached to those characteristics. The concrete subclasses are:


// A representation of common aspects of services offered by a peripheral.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreBluetooth/CBAttribute
type CBAttribute struct {
	objectivec.Object
}

// CBAttributeFrom constructs a [CBAttribute] from an unsafe.Pointer.
//
// A representation of common aspects of services offered by a peripheral.
func CBAttributeFrom(ptr unsafe.Pointer) CBAttribute {
	return CBAttribute{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CBAttribute *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CBAttribute */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CBAttribute */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CBAttribute */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CBAttribute */

// The Bluetooth-specific UUID of the attribute.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreBluetooth/CBAttribute/uuid
func (c_ CBAttribute) UUID() ICBUUID {
	rv := objc.Send[CBUUID](c_.ID, objc.Sel("UUID"))
	return rv
}/* debug [instance_properties/getter]: UUID */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CBAttribute */



