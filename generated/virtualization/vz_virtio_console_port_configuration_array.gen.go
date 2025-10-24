// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class VZVirtioConsolePortConfigurationArray */


/* debug [class_header]: Header for VZVirtioConsolePortConfigurationArray */
// The class instance for the [VZVirtioConsolePortConfigurationArray] class.
var (
	VZVirtioConsolePortConfigurationArrayClass     _VZVirtioConsolePortConfigurationArrayClass
	VZVirtioConsolePortConfigurationArrayClassOnce sync.Once
)

func getVZVirtioConsolePortConfigurationArrayClass() _VZVirtioConsolePortConfigurationArrayClass {
	VZVirtioConsolePortConfigurationArrayClassOnce.Do(func() {
		VZVirtioConsolePortConfigurationArrayClass = _VZVirtioConsolePortConfigurationArrayClass{objc.GetClass("VZVirtioConsolePortConfigurationArray")}
	})
	return VZVirtioConsolePortConfigurationArrayClass
}

type _VZVirtioConsolePortConfigurationArrayClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for VZVirtioConsolePortConfigurationArray */
// An interface definition for the [VZVirtioConsolePortConfigurationArray] class.
type IVZVirtioConsolePortConfigurationArray interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for VZVirtioConsolePortConfigurationArray */
	// properties:
	MaximumPortCount() uint32 /* not a class type */
	SetMaximumPortCount(value uint32 /* not a class type */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for VZVirtioConsolePortConfigurationArray */
	// methods:
	SetObjectAtIndexedSubscript(configuration IVZVirtioConsolePortConfiguration, portIndex uint)
	ObjectAtIndexedSubscript(portIndex uint) IVZVirtioConsolePortConfiguration
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for VZVirtioConsolePortConfigurationArray */
// Alloc allocates a new instance without initialization.
func (vc _VZVirtioConsolePortConfigurationArrayClass) Alloc() VZVirtioConsolePortConfigurationArray {
	rv := objc.Send[VZVirtioConsolePortConfigurationArray](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (vc _VZVirtioConsolePortConfigurationArrayClass) New() VZVirtioConsolePortConfigurationArray {
	rv := objc.Send[VZVirtioConsolePortConfigurationArray](objc.ID(vc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (v_ VZVirtioConsolePortConfigurationArray) Init() VZVirtioConsolePortConfigurationArray {
	rv := objc.Send[VZVirtioConsolePortConfigurationArray](v_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v_ VZVirtioConsolePortConfigurationArray) Autorelease() VZVirtioConsolePortConfigurationArray {
	rv := objc.Send[VZVirtioConsolePortConfigurationArray](v_.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZVirtioConsolePortConfigurationArray creates a new VZVirtioConsolePortConfigurationArray instance.
func NewVZVirtioConsolePortConfigurationArray() VZVirtioConsolePortConfigurationArray {
	return getVZVirtioConsolePortConfigurationArrayClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for VZVirtioConsolePortConfigurationArray */
// A class that represents a collection of Virtio console port configurations.
//
// This array stores a collection of port configurations for a . The index in the array corresponds to the port index that the VM uses. You can set a value, but the value must be larger than the highest indexed port. If there’s no value set, the framework uses the value the highest indexed port.


// A class that represents a collection of Virtio console port configurations.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtioConsolePortConfigurationArray
type VZVirtioConsolePortConfigurationArray struct {
	objectivec.Object
}

// VZVirtioConsolePortConfigurationArrayFrom constructs a [VZVirtioConsolePortConfigurationArray] from an unsafe.Pointer.
//
// A class that represents a collection of Virtio console port configurations.
func VZVirtioConsolePortConfigurationArrayFrom(ptr unsafe.Pointer) VZVirtioConsolePortConfigurationArray {
	return VZVirtioConsolePortConfigurationArray{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for VZVirtioConsolePortConfigurationArray *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for VZVirtioConsolePortConfigurationArray */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for VZVirtioConsolePortConfigurationArray */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for VZVirtioConsolePortConfigurationArray */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtioConsolePortConfigurationArray/setObject:atIndexedSubscript:
func (v_ VZVirtioConsolePortConfigurationArray) SetObjectAtIndexedSubscript(configuration IVZVirtioConsolePortConfiguration, portIndex uint) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setObject:atIndexedSubscript:"), configuration, portIndex)
}/* debug [instance_methods/method]: SetObjectAtIndexedSubscript */


// Returns the Virtio console port configuration as the specified index.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtioConsolePortConfigurationArray/subscript(_:)
func (v_ VZVirtioConsolePortConfigurationArray) ObjectAtIndexedSubscript(portIndex uint) IVZVirtioConsolePortConfiguration {
	rv := objc.Send[VZVirtioConsolePortConfiguration](v_.ID, objc.Sel("objectAtIndexedSubscript:"), portIndex)
	return rv
}/* debug [instance_methods/method]: ObjectAtIndexedSubscript */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for VZVirtioConsolePortConfigurationArray */

// An unsigned integer that represents the maximum number of ports allocated by this device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtioConsolePortConfigurationArray/maximumPortCount
func (v_ VZVirtioConsolePortConfigurationArray) MaximumPortCount() uint32 /* not a class type */ {
	rv := objc.Send[uint32](v_.ID, objc.Sel("maximumPortCount"))
	return rv
}/* debug [instance_properties/getter]: maximumPortCount */


// An unsigned integer that represents the maximum number of ports allocated by this device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtioConsolePortConfigurationArray/maximumPortCount
func (v_ VZVirtioConsolePortConfigurationArray) SetMaximumPortCount(value uint32 /* not a class type */) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setMaximumPortCount:"), value)
}/* debug [instance_properties/setter]: maximumPortCount */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class VZVirtioConsolePortConfigurationArray */



