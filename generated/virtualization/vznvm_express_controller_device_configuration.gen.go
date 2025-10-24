// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class VZNVMExpressControllerDeviceConfiguration */

/* debug [class_header]: Header for VZNVMExpressControllerDeviceConfiguration */
// The class instance for the [VZNVMExpressControllerDeviceConfiguration] class.
var (
	VZNVMExpressControllerDeviceConfigurationClass     _VZNVMExpressControllerDeviceConfigurationClass
	VZNVMExpressControllerDeviceConfigurationClassOnce sync.Once
)

func getVZNVMExpressControllerDeviceConfigurationClass() _VZNVMExpressControllerDeviceConfigurationClass {
	VZNVMExpressControllerDeviceConfigurationClassOnce.Do(func() {
		VZNVMExpressControllerDeviceConfigurationClass = _VZNVMExpressControllerDeviceConfigurationClass{objc.GetClass("VZNVMExpressControllerDeviceConfiguration")}
	})
	return VZNVMExpressControllerDeviceConfigurationClass
}

type _VZNVMExpressControllerDeviceConfigurationClass struct {
	class objc.Class
}

/* debug [class_header]: End header */

/* debug [class_interface]: Interface for VZNVMExpressControllerDeviceConfiguration */
// An interface definition for the [VZNVMExpressControllerDeviceConfiguration] class.
type IVZNVMExpressControllerDeviceConfiguration interface {
	IVZStorageDeviceConfiguration

	/* debug [class_interface_properties]: Properties for VZNVMExpressControllerDeviceConfiguration */
	// properties:
	/* debug [class_interface_properties]: End properties */

	/* debug [class_interface_methods]: Methods for VZNVMExpressControllerDeviceConfiguration */
	// methods:
	/* debug [class_interface_methods]: End methods */

}

/* debug [class_interface]: End interface */

/* debug [class_constructors]: Constructors for VZNVMExpressControllerDeviceConfiguration */
// Alloc allocates a new instance without initialization.
func (vc _VZNVMExpressControllerDeviceConfigurationClass) Alloc() VZNVMExpressControllerDeviceConfiguration {
	rv := objc.Send[VZNVMExpressControllerDeviceConfiguration](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (vc _VZNVMExpressControllerDeviceConfigurationClass) New() VZNVMExpressControllerDeviceConfiguration {
	rv := objc.Send[VZNVMExpressControllerDeviceConfiguration](objc.ID(vc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (v_ VZNVMExpressControllerDeviceConfiguration) Init() VZNVMExpressControllerDeviceConfiguration {
	rv := objc.Send[VZNVMExpressControllerDeviceConfiguration](v_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v_ VZNVMExpressControllerDeviceConfiguration) Autorelease() VZNVMExpressControllerDeviceConfiguration {
	rv := objc.Send[VZNVMExpressControllerDeviceConfiguration](v_.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZNVMExpressControllerDeviceConfiguration creates a new VZNVMExpressControllerDeviceConfiguration instance.
func NewVZNVMExpressControllerDeviceConfiguration() VZNVMExpressControllerDeviceConfiguration {
	return getVZNVMExpressControllerDeviceConfigurationClass().New()
}

/* debug [class_constructors]: End constructors */

/* debug [class_struct]: Struct for VZNVMExpressControllerDeviceConfiguration */
// The configuration object that represents an NVM Express Controller storage device.
//
// This device configuration creates a storage device that conforms to the . The device configuration is valid only if used with .

// The configuration object that represents an NVM Express Controller storage device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZNVMExpressControllerDeviceConfiguration
type VZNVMExpressControllerDeviceConfiguration struct {
	VZStorageDeviceConfiguration
}

// VZNVMExpressControllerDeviceConfigurationFrom constructs a [VZNVMExpressControllerDeviceConfiguration] from an unsafe.Pointer.
//
// The configuration object that represents an NVM Express Controller storage device.
func VZNVMExpressControllerDeviceConfigurationFrom(ptr unsafe.Pointer) VZNVMExpressControllerDeviceConfiguration {
	return VZNVMExpressControllerDeviceConfiguration{
		VZStorageDeviceConfiguration: VZStorageDeviceConfigurationFrom(ptr),
	}
}

/* debug [class_struct]: End struct */

/* debug [class_init_methods]: Init methods for VZNVMExpressControllerDeviceConfiguration */

// Creates a new NVM Express controller configuration with the storage device attachment you provide.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZNVMExpressControllerDeviceConfiguration/init(attachment:)
func NewVZNVMExpressControllerDeviceConfigurationWithAttachment(attachment IVZStorageDeviceAttachment) VZNVMExpressControllerDeviceConfiguration {
	instance := getVZNVMExpressControllerDeviceConfigurationClass().Alloc()
	rv := objc.Send[VZNVMExpressControllerDeviceConfiguration](instance.ID, objc.Sel("initWithAttachment:"), attachment)
	rv.Autorelease()
	return rv
} /* debug [class_init_methods/constructor]: NewVZNVMExpressControllerDeviceConfigurationWithAttachment */

/* debug [class_init_methods]: End init methods */

/* debug [class_methods]: Class methods for VZNVMExpressControllerDeviceConfiguration */
/* debug [class_methods]: End class methods */

/* debug [class_properties_class]: Class properties for VZNVMExpressControllerDeviceConfiguration */
/* debug [class_properties_class]: End class properties */

/* debug [instance_methods]: Instance methods for VZNVMExpressControllerDeviceConfiguration */
/* debug [instance_methods]: End instance methods */

/* debug [instance_properties]: Instance properties for VZNVMExpressControllerDeviceConfiguration */
/* debug [instance_properties]: End instance properties */

/* debug [class.gen.go]: End class VZNVMExpressControllerDeviceConfiguration */
