// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class VZDirectorySharingDevice */

/* debug [class_header]: Header for VZDirectorySharingDevice */
// The class instance for the [VZDirectorySharingDevice] class.
var (
	VZDirectorySharingDeviceClass     _VZDirectorySharingDeviceClass
	VZDirectorySharingDeviceClassOnce sync.Once
)

func getVZDirectorySharingDeviceClass() _VZDirectorySharingDeviceClass {
	VZDirectorySharingDeviceClassOnce.Do(func() {
		VZDirectorySharingDeviceClass = _VZDirectorySharingDeviceClass{objc.GetClass("VZDirectorySharingDevice")}
	})
	return VZDirectorySharingDeviceClass
}

type _VZDirectorySharingDeviceClass struct {
	class objc.Class
}

/* debug [class_header]: End header */

/* debug [class_interface]: Interface for VZDirectorySharingDevice */
// An interface definition for the [VZDirectorySharingDevice] class.
type IVZDirectorySharingDevice interface {
	objectivec.IObject

	/* debug [class_interface_properties]: Properties for VZDirectorySharingDevice */
	// properties:
	/* debug [class_interface_properties]: End properties */

	/* debug [class_interface_methods]: Methods for VZDirectorySharingDevice */
	// methods:
	/* debug [class_interface_methods]: End methods */

}

/* debug [class_interface]: End interface */

/* debug [class_constructors]: Constructors for VZDirectorySharingDevice */
// Alloc allocates a new instance without initialization.
func (vc _VZDirectorySharingDeviceClass) Alloc() VZDirectorySharingDevice {
	rv := objc.Send[VZDirectorySharingDevice](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (vc _VZDirectorySharingDeviceClass) New() VZDirectorySharingDevice {
	rv := objc.Send[VZDirectorySharingDevice](objc.ID(vc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (v_ VZDirectorySharingDevice) Init() VZDirectorySharingDevice {
	rv := objc.Send[VZDirectorySharingDevice](v_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v_ VZDirectorySharingDevice) Autorelease() VZDirectorySharingDevice {
	rv := objc.Send[VZDirectorySharingDevice](v_.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZDirectorySharingDevice creates a new VZDirectorySharingDevice instance.
func NewVZDirectorySharingDevice() VZDirectorySharingDevice {
	return getVZDirectorySharingDeviceClass().New()
}

/* debug [class_constructors]: End constructors */

/* debug [class_struct]: Struct for VZDirectorySharingDevice */
// The base class that represents a directory sharing device in a VM.
//
// Don’t instantiate directly; configure a directory sharing device first by using through a subclass of . When you create a from the configuration, the directory sharing devices are available through the property. The real type of corresponds to the type used by the configuration. For example, a leads to a device of type .

// The base class that represents a directory sharing device in a VM.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZDirectorySharingDevice
type VZDirectorySharingDevice struct {
	objectivec.Object
}

// VZDirectorySharingDeviceFrom constructs a [VZDirectorySharingDevice] from an unsafe.Pointer.
//
// The base class that represents a directory sharing device in a VM.
func VZDirectorySharingDeviceFrom(ptr unsafe.Pointer) VZDirectorySharingDevice {
	return VZDirectorySharingDevice{objectivec.Object{objc.ID(ptr)}}
}

/* debug [class_struct]: End struct */

/* debug [class_init_methods]: Init methods for VZDirectorySharingDevice */ /* debug [class_init_methods]: End init methods */

/* debug [class_methods]: Class methods for VZDirectorySharingDevice */
/* debug [class_methods]: End class methods */

/* debug [class_properties_class]: Class properties for VZDirectorySharingDevice */
/* debug [class_properties_class]: End class properties */

/* debug [instance_methods]: Instance methods for VZDirectorySharingDevice */
/* debug [instance_methods]: End instance methods */

/* debug [instance_properties]: Instance properties for VZDirectorySharingDevice */
/* debug [instance_properties]: End instance properties */

/* debug [class.gen.go]: End class VZDirectorySharingDevice */
