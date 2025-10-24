// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class VZMemoryBalloonDevice */


/* debug [class_header]: Header for VZMemoryBalloonDevice */
// The class instance for the [VZMemoryBalloonDevice] class.
var (
	VZMemoryBalloonDeviceClass     _VZMemoryBalloonDeviceClass
	VZMemoryBalloonDeviceClassOnce sync.Once
)

func getVZMemoryBalloonDeviceClass() _VZMemoryBalloonDeviceClass {
	VZMemoryBalloonDeviceClassOnce.Do(func() {
		VZMemoryBalloonDeviceClass = _VZMemoryBalloonDeviceClass{objc.GetClass("VZMemoryBalloonDevice")}
	})
	return VZMemoryBalloonDeviceClass
}

type _VZMemoryBalloonDeviceClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for VZMemoryBalloonDevice */
// An interface definition for the [VZMemoryBalloonDevice] class.
type IVZMemoryBalloonDevice interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for VZMemoryBalloonDevice */
	// properties:
	MemoryBalloonDevices() IVZMemoryBalloonDeviceConfiguration
	SetMemoryBalloonDevices(value IVZMemoryBalloonDeviceConfiguration)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for VZMemoryBalloonDevice */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for VZMemoryBalloonDevice */
// Alloc allocates a new instance without initialization.
func (vc _VZMemoryBalloonDeviceClass) Alloc() VZMemoryBalloonDevice {
	rv := objc.Send[VZMemoryBalloonDevice](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (vc _VZMemoryBalloonDeviceClass) New() VZMemoryBalloonDevice {
	rv := objc.Send[VZMemoryBalloonDevice](objc.ID(vc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (v_ VZMemoryBalloonDevice) Init() VZMemoryBalloonDevice {
	rv := objc.Send[VZMemoryBalloonDevice](v_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v_ VZMemoryBalloonDevice) Autorelease() VZMemoryBalloonDevice {
	rv := objc.Send[VZMemoryBalloonDevice](v_.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZMemoryBalloonDevice creates a new VZMemoryBalloonDevice instance.
func NewVZMemoryBalloonDevice() VZMemoryBalloonDevice {
	return getVZMemoryBalloonDeviceClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for VZMemoryBalloonDevice */
// The common behavior for memory devices.
//
// Don’t instantiate this class directly. To request a memory ballon device, add an appropriate configuration object to the property of the object that you use to configure the virtual machine. In response, the system instantiates the subclass of that matches your request. For example, if you supply a object in your configuration, the system creates a object.


// The common behavior for memory devices.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZMemoryBalloonDevice
type VZMemoryBalloonDevice struct {
	objectivec.Object
}

// VZMemoryBalloonDeviceFrom constructs a [VZMemoryBalloonDevice] from an unsafe.Pointer.
//
// The common behavior for memory devices.
func VZMemoryBalloonDeviceFrom(ptr unsafe.Pointer) VZMemoryBalloonDevice {
	return VZMemoryBalloonDevice{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for VZMemoryBalloonDevice *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for VZMemoryBalloonDevice */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for VZMemoryBalloonDevice */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for VZMemoryBalloonDevice */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for VZMemoryBalloonDevice */

// An array that you configure with a memory balloon device, used to update the memory in the VM.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/virtualization/vzvirtualmachineconfiguration/memoryballoondevices
func (v_ VZMemoryBalloonDevice) MemoryBalloonDevices() IVZMemoryBalloonDeviceConfiguration {
	rv := objc.Send[VZMemoryBalloonDeviceConfiguration](v_.ID, objc.Sel("memoryBalloonDevices"))
	return rv
}/* debug [instance_properties/getter]: memoryBalloonDevices */


// An array that you configure with a memory balloon device, used to update the memory in the VM.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/virtualization/vzvirtualmachineconfiguration/memoryballoondevices
func (v_ VZMemoryBalloonDevice) SetMemoryBalloonDevices(value IVZMemoryBalloonDeviceConfiguration) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setMemoryBalloonDevices:"), value)
}/* debug [instance_properties/setter]: memoryBalloonDevices */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class VZMemoryBalloonDevice */



