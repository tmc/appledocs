// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class VZXHCIControllerConfiguration */

/* debug [class_header]: Header for VZXHCIControllerConfiguration */
// The class instance for the [VZXHCIControllerConfiguration] class.
var (
	VZXHCIControllerConfigurationClass     _VZXHCIControllerConfigurationClass
	VZXHCIControllerConfigurationClassOnce sync.Once
)

func getVZXHCIControllerConfigurationClass() _VZXHCIControllerConfigurationClass {
	VZXHCIControllerConfigurationClassOnce.Do(func() {
		VZXHCIControllerConfigurationClass = _VZXHCIControllerConfigurationClass{objc.GetClass("VZXHCIControllerConfiguration")}
	})
	return VZXHCIControllerConfigurationClass
}

type _VZXHCIControllerConfigurationClass struct {
	class objc.Class
}

/* debug [class_header]: End header */

/* debug [class_interface]: Interface for VZXHCIControllerConfiguration */
// An interface definition for the [VZXHCIControllerConfiguration] class.
type IVZXHCIControllerConfiguration interface {
	IVZUSBControllerConfiguration

	/* debug [class_interface_properties]: Properties for VZXHCIControllerConfiguration */
	// properties:
	/* debug [class_interface_properties]: End properties */

	/* debug [class_interface_methods]: Methods for VZXHCIControllerConfiguration */
	// methods:
	/* debug [class_interface_methods]: End methods */

}

/* debug [class_interface]: End interface */

/* debug [class_constructors]: Constructors for VZXHCIControllerConfiguration */
// Alloc allocates a new instance without initialization.
func (vc _VZXHCIControllerConfigurationClass) Alloc() VZXHCIControllerConfiguration {
	rv := objc.Send[VZXHCIControllerConfiguration](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (vc _VZXHCIControllerConfigurationClass) New() VZXHCIControllerConfiguration {
	rv := objc.Send[VZXHCIControllerConfiguration](objc.ID(vc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (v_ VZXHCIControllerConfiguration) Init() VZXHCIControllerConfiguration {
	rv := objc.Send[VZXHCIControllerConfiguration](v_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v_ VZXHCIControllerConfiguration) Autorelease() VZXHCIControllerConfiguration {
	rv := objc.Send[VZXHCIControllerConfiguration](v_.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZXHCIControllerConfiguration creates a new VZXHCIControllerConfiguration instance.
func NewVZXHCIControllerConfiguration() VZXHCIControllerConfiguration {
	return getVZXHCIControllerConfigurationClass().New()
}

/* debug [class_constructors]: End constructors */

/* debug [class_struct]: Struct for VZXHCIControllerConfiguration */
// The configuration object for the USB Extensible Host Controller Interface (XHCI) controller.
//
// Use this configuration to create a USB XHCI controller device for the guest.

// The configuration object for the USB Extensible Host Controller Interface (XHCI) controller.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZXHCIControllerConfiguration
type VZXHCIControllerConfiguration struct {
	VZUSBControllerConfiguration
}

// VZXHCIControllerConfigurationFrom constructs a [VZXHCIControllerConfiguration] from an unsafe.Pointer.
//
// The configuration object for the USB Extensible Host Controller Interface (XHCI) controller.
func VZXHCIControllerConfigurationFrom(ptr unsafe.Pointer) VZXHCIControllerConfiguration {
	return VZXHCIControllerConfiguration{
		VZUSBControllerConfiguration: VZUSBControllerConfigurationFrom(ptr),
	}
}

/* debug [class_struct]: End struct */

/* debug [class_init_methods]: Init methods for VZXHCIControllerConfiguration */
/* debug [class_init_methods]: End init methods */

/* debug [class_methods]: Class methods for VZXHCIControllerConfiguration */
/* debug [class_methods]: End class methods */

/* debug [class_properties_class]: Class properties for VZXHCIControllerConfiguration */
/* debug [class_properties_class]: End class properties */

/* debug [instance_methods]: Instance methods for VZXHCIControllerConfiguration */
/* debug [instance_methods]: End instance methods */

/* debug [instance_properties]: Instance properties for VZXHCIControllerConfiguration */
/* debug [instance_properties]: End instance properties */

/* debug [class.gen.go]: End class VZXHCIControllerConfiguration */
