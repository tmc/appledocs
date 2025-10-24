// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class VZStorageDevice */


/* debug [class_header]: Header for VZStorageDevice */
// The class instance for the [VZStorageDevice] class.
var (
	VZStorageDeviceClass     _VZStorageDeviceClass
	VZStorageDeviceClassOnce sync.Once
)

func getVZStorageDeviceClass() _VZStorageDeviceClass {
	VZStorageDeviceClassOnce.Do(func() {
		VZStorageDeviceClass = _VZStorageDeviceClass{objc.GetClass("VZStorageDevice")}
	})
	return VZStorageDeviceClass
}

type _VZStorageDeviceClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for VZStorageDevice */
// An interface definition for the [VZStorageDevice] class.
type IVZStorageDevice interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for VZStorageDevice */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for VZStorageDevice */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for VZStorageDevice */
// Alloc allocates a new instance without initialization.
func (vc _VZStorageDeviceClass) Alloc() VZStorageDevice {
	rv := objc.Send[VZStorageDevice](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (vc _VZStorageDeviceClass) New() VZStorageDevice {
	rv := objc.Send[VZStorageDevice](objc.ID(vc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (v_ VZStorageDevice) Init() VZStorageDevice {
	rv := objc.Send[VZStorageDevice](v_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v_ VZStorageDevice) Autorelease() VZStorageDevice {
	rv := objc.Send[VZStorageDevice](v_.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZStorageDevice creates a new VZStorageDevice instance.
func NewVZStorageDevice() VZStorageDevice {
	return getVZStorageDeviceClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for VZStorageDevice */
// A class that represents a storage device in a VM.
//
// Don’t create a directly. Use one of its subclasses, such as , instead.


// A class that represents a storage device in a VM.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZStorageDevice
type VZStorageDevice struct {
	objectivec.Object
}

// VZStorageDeviceFrom constructs a [VZStorageDevice] from an unsafe.Pointer.
//
// A class that represents a storage device in a VM.
func VZStorageDeviceFrom(ptr unsafe.Pointer) VZStorageDevice {
	return VZStorageDevice{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for VZStorageDevice *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for VZStorageDevice */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for VZStorageDevice */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for VZStorageDevice */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for VZStorageDevice */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class VZStorageDevice */



