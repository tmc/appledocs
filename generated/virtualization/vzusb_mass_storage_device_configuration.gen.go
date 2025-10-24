// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class VZUSBMassStorageDeviceConfiguration */

/* debug [class_header]: Header for VZUSBMassStorageDeviceConfiguration */
// The class instance for the [VZUSBMassStorageDeviceConfiguration] class.
var (
	VZUSBMassStorageDeviceConfigurationClass     _VZUSBMassStorageDeviceConfigurationClass
	VZUSBMassStorageDeviceConfigurationClassOnce sync.Once
)

func getVZUSBMassStorageDeviceConfigurationClass() _VZUSBMassStorageDeviceConfigurationClass {
	VZUSBMassStorageDeviceConfigurationClassOnce.Do(func() {
		VZUSBMassStorageDeviceConfigurationClass = _VZUSBMassStorageDeviceConfigurationClass{objc.GetClass("VZUSBMassStorageDeviceConfiguration")}
	})
	return VZUSBMassStorageDeviceConfigurationClass
}

type _VZUSBMassStorageDeviceConfigurationClass struct {
	class objc.Class
}

/* debug [class_header]: End header */

/* debug [class_interface]: Interface for VZUSBMassStorageDeviceConfiguration */
// An interface definition for the [VZUSBMassStorageDeviceConfiguration] class.
type IVZUSBMassStorageDeviceConfiguration interface {
	IVZStorageDeviceConfiguration

	/* debug [class_interface_properties]: Properties for VZUSBMassStorageDeviceConfiguration */
	// properties:
	/* debug [class_interface_properties]: End properties */

	/* debug [class_interface_methods]: Methods for VZUSBMassStorageDeviceConfiguration */
	// methods:
	/* debug [class_interface_methods]: End methods */

}

/* debug [class_interface]: End interface */

/* debug [class_constructors]: Constructors for VZUSBMassStorageDeviceConfiguration */
// Alloc allocates a new instance without initialization.
func (vc _VZUSBMassStorageDeviceConfigurationClass) Alloc() VZUSBMassStorageDeviceConfiguration {
	rv := objc.Send[VZUSBMassStorageDeviceConfiguration](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (vc _VZUSBMassStorageDeviceConfigurationClass) New() VZUSBMassStorageDeviceConfiguration {
	rv := objc.Send[VZUSBMassStorageDeviceConfiguration](objc.ID(vc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (v_ VZUSBMassStorageDeviceConfiguration) Init() VZUSBMassStorageDeviceConfiguration {
	rv := objc.Send[VZUSBMassStorageDeviceConfiguration](v_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v_ VZUSBMassStorageDeviceConfiguration) Autorelease() VZUSBMassStorageDeviceConfiguration {
	rv := objc.Send[VZUSBMassStorageDeviceConfiguration](v_.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZUSBMassStorageDeviceConfiguration creates a new VZUSBMassStorageDeviceConfiguration instance.
func NewVZUSBMassStorageDeviceConfiguration() VZUSBMassStorageDeviceConfiguration {
	return getVZUSBMassStorageDeviceConfigurationClass().New()
}

/* debug [class_constructors]: End constructors */

/* debug [class_struct]: Struct for VZUSBMassStorageDeviceConfiguration */
// The configuration object that represents a USB Mass storage device.

// The configuration object that represents a USB Mass storage device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZUSBMassStorageDeviceConfiguration
type VZUSBMassStorageDeviceConfiguration struct {
	VZStorageDeviceConfiguration
}

// VZUSBMassStorageDeviceConfigurationFrom constructs a [VZUSBMassStorageDeviceConfiguration] from an unsafe.Pointer.
//
// The configuration object that represents a USB Mass storage device.
func VZUSBMassStorageDeviceConfigurationFrom(ptr unsafe.Pointer) VZUSBMassStorageDeviceConfiguration {
	return VZUSBMassStorageDeviceConfiguration{
		VZStorageDeviceConfiguration: VZStorageDeviceConfigurationFrom(ptr),
	}
}

/* debug [class_struct]: End struct */

/* debug [class_init_methods]: Init methods for VZUSBMassStorageDeviceConfiguration */

// Creates a new storage device configuration with the specified attachment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZUSBMassStorageDeviceConfiguration/init(attachment:)
func NewVZUSBMassStorageDeviceConfigurationWithAttachment(attachment IVZStorageDeviceAttachment) VZUSBMassStorageDeviceConfiguration {
	instance := getVZUSBMassStorageDeviceConfigurationClass().Alloc()
	rv := objc.Send[VZUSBMassStorageDeviceConfiguration](instance.ID, objc.Sel("initWithAttachment:"), attachment)
	rv.Autorelease()
	return rv
} /* debug [class_init_methods/constructor]: NewVZUSBMassStorageDeviceConfigurationWithAttachment */

/* debug [class_init_methods]: End init methods */

/* debug [class_methods]: Class methods for VZUSBMassStorageDeviceConfiguration */
/* debug [class_methods]: End class methods */

/* debug [class_properties_class]: Class properties for VZUSBMassStorageDeviceConfiguration */
/* debug [class_properties_class]: End class properties */

/* debug [instance_methods]: Instance methods for VZUSBMassStorageDeviceConfiguration */
/* debug [instance_methods]: End instance methods */

/* debug [instance_properties]: Instance properties for VZUSBMassStorageDeviceConfiguration */
/* debug [instance_properties]: End instance properties */

/* debug [class.gen.go]: End class VZUSBMassStorageDeviceConfiguration */
