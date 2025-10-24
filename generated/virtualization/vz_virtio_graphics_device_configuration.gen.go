// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class VZVirtioGraphicsDeviceConfiguration */

/* debug [class_header]: Header for VZVirtioGraphicsDeviceConfiguration */
// The class instance for the [VZVirtioGraphicsDeviceConfiguration] class.
var (
	VZVirtioGraphicsDeviceConfigurationClass     _VZVirtioGraphicsDeviceConfigurationClass
	VZVirtioGraphicsDeviceConfigurationClassOnce sync.Once
)

func getVZVirtioGraphicsDeviceConfigurationClass() _VZVirtioGraphicsDeviceConfigurationClass {
	VZVirtioGraphicsDeviceConfigurationClassOnce.Do(func() {
		VZVirtioGraphicsDeviceConfigurationClass = _VZVirtioGraphicsDeviceConfigurationClass{objc.GetClass("VZVirtioGraphicsDeviceConfiguration")}
	})
	return VZVirtioGraphicsDeviceConfigurationClass
}

type _VZVirtioGraphicsDeviceConfigurationClass struct {
	class objc.Class
}

/* debug [class_header]: End header */

/* debug [class_interface]: Interface for VZVirtioGraphicsDeviceConfiguration */
// An interface definition for the [VZVirtioGraphicsDeviceConfiguration] class.
type IVZVirtioGraphicsDeviceConfiguration interface {
	IVZGraphicsDeviceConfiguration

	/* debug [class_interface_properties]: Properties for VZVirtioGraphicsDeviceConfiguration */
	// properties:
	Scanouts() []VZVirtioGraphicsScanoutConfiguration
	SetScanouts(value []VZVirtioGraphicsScanoutConfiguration)
	/* debug [class_interface_properties]: End properties */

	/* debug [class_interface_methods]: Methods for VZVirtioGraphicsDeviceConfiguration */
	// methods:
	/* debug [class_interface_methods]: End methods */

}

/* debug [class_interface]: End interface */

/* debug [class_constructors]: Constructors for VZVirtioGraphicsDeviceConfiguration */
// Alloc allocates a new instance without initialization.
func (vc _VZVirtioGraphicsDeviceConfigurationClass) Alloc() VZVirtioGraphicsDeviceConfiguration {
	rv := objc.Send[VZVirtioGraphicsDeviceConfiguration](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (vc _VZVirtioGraphicsDeviceConfigurationClass) New() VZVirtioGraphicsDeviceConfiguration {
	rv := objc.Send[VZVirtioGraphicsDeviceConfiguration](objc.ID(vc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (v_ VZVirtioGraphicsDeviceConfiguration) Init() VZVirtioGraphicsDeviceConfiguration {
	rv := objc.Send[VZVirtioGraphicsDeviceConfiguration](v_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v_ VZVirtioGraphicsDeviceConfiguration) Autorelease() VZVirtioGraphicsDeviceConfiguration {
	rv := objc.Send[VZVirtioGraphicsDeviceConfiguration](v_.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZVirtioGraphicsDeviceConfiguration creates a new VZVirtioGraphicsDeviceConfiguration instance.
func NewVZVirtioGraphicsDeviceConfiguration() VZVirtioGraphicsDeviceConfiguration {
	return getVZVirtioGraphicsDeviceConfigurationClass().New()
}

/* debug [class_constructors]: End constructors */

/* debug [class_struct]: Struct for VZVirtioGraphicsDeviceConfiguration */
// Configuration that represents the configuration of a Virtio graphics device for a Linux VM.

// Configuration that represents the configuration of a Virtio graphics device for a Linux VM.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtioGraphicsDeviceConfiguration
type VZVirtioGraphicsDeviceConfiguration struct {
	VZGraphicsDeviceConfiguration
}

// VZVirtioGraphicsDeviceConfigurationFrom constructs a [VZVirtioGraphicsDeviceConfiguration] from an unsafe.Pointer.
//
// Configuration that represents the configuration of a Virtio graphics device for a Linux VM.
func VZVirtioGraphicsDeviceConfigurationFrom(ptr unsafe.Pointer) VZVirtioGraphicsDeviceConfiguration {
	return VZVirtioGraphicsDeviceConfiguration{
		VZGraphicsDeviceConfiguration: VZGraphicsDeviceConfigurationFrom(ptr),
	}
}

/* debug [class_struct]: End struct */

/* debug [class_init_methods]: Init methods for VZVirtioGraphicsDeviceConfiguration */
/* debug [class_init_methods]: End init methods */

/* debug [class_methods]: Class methods for VZVirtioGraphicsDeviceConfiguration */
/* debug [class_methods]: End class methods */

/* debug [class_properties_class]: Class properties for VZVirtioGraphicsDeviceConfiguration */
/* debug [class_properties_class]: End class properties */

/* debug [instance_methods]: Instance methods for VZVirtioGraphicsDeviceConfiguration */
/* debug [instance_methods]: End instance methods */

/* debug [instance_properties]: Instance properties for VZVirtioGraphicsDeviceConfiguration */

// The array of output devices.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtioGraphicsDeviceConfiguration/scanouts
func (v_ VZVirtioGraphicsDeviceConfiguration) Scanouts() []VZVirtioGraphicsScanoutConfiguration {
	rv := objc.Send[[]VZVirtioGraphicsScanoutConfiguration](v_.ID, objc.Sel("scanouts"))
	return rv
} /* debug [instance_properties/getter]: scanouts */

// The array of output devices.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtioGraphicsDeviceConfiguration/scanouts
func (v_ VZVirtioGraphicsDeviceConfiguration) SetScanouts(value []VZVirtioGraphicsScanoutConfiguration) {
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](v_.ID, objc.Sel("setScanouts:"), nsArray)
} /* debug [instance_properties/setter]: scanouts */

/* debug [instance_properties]: End instance properties */

/* debug [class.gen.go]: End class VZVirtioGraphicsDeviceConfiguration */
