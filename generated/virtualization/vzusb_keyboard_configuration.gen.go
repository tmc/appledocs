// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class VZUSBKeyboardConfiguration */

/* debug [class_header]: Header for VZUSBKeyboardConfiguration */
// The class instance for the [VZUSBKeyboardConfiguration] class.
var (
	VZUSBKeyboardConfigurationClass     _VZUSBKeyboardConfigurationClass
	VZUSBKeyboardConfigurationClassOnce sync.Once
)

func getVZUSBKeyboardConfigurationClass() _VZUSBKeyboardConfigurationClass {
	VZUSBKeyboardConfigurationClassOnce.Do(func() {
		VZUSBKeyboardConfigurationClass = _VZUSBKeyboardConfigurationClass{objc.GetClass("VZUSBKeyboardConfiguration")}
	})
	return VZUSBKeyboardConfigurationClass
}

type _VZUSBKeyboardConfigurationClass struct {
	class objc.Class
}

/* debug [class_header]: End header */

/* debug [class_interface]: Interface for VZUSBKeyboardConfiguration */
// An interface definition for the [VZUSBKeyboardConfiguration] class.
type IVZUSBKeyboardConfiguration interface {
	IVZKeyboardConfiguration

	/* debug [class_interface_properties]: Properties for VZUSBKeyboardConfiguration */
	// properties:
	/* debug [class_interface_properties]: End properties */

	/* debug [class_interface_methods]: Methods for VZUSBKeyboardConfiguration */
	// methods:
	/* debug [class_interface_methods]: End methods */

}

/* debug [class_interface]: End interface */

/* debug [class_constructors]: Constructors for VZUSBKeyboardConfiguration */
// Alloc allocates a new instance without initialization.
func (vc _VZUSBKeyboardConfigurationClass) Alloc() VZUSBKeyboardConfiguration {
	rv := objc.Send[VZUSBKeyboardConfiguration](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (vc _VZUSBKeyboardConfigurationClass) New() VZUSBKeyboardConfiguration {
	rv := objc.Send[VZUSBKeyboardConfiguration](objc.ID(vc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (v_ VZUSBKeyboardConfiguration) Init() VZUSBKeyboardConfiguration {
	rv := objc.Send[VZUSBKeyboardConfiguration](v_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v_ VZUSBKeyboardConfiguration) Autorelease() VZUSBKeyboardConfiguration {
	rv := objc.Send[VZUSBKeyboardConfiguration](v_.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZUSBKeyboardConfiguration creates a new VZUSBKeyboardConfiguration instance.
func NewVZUSBKeyboardConfiguration() VZUSBKeyboardConfiguration {
	return getVZUSBKeyboardConfigurationClass().New()
}

/* debug [class_constructors]: End constructors */

/* debug [class_struct]: Struct for VZUSBKeyboardConfiguration */
// A device that defines the configuration for a USB keyboard.
//
// A can use this device to send key events to the VM.

// A device that defines the configuration for a USB keyboard.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZUSBKeyboardConfiguration
type VZUSBKeyboardConfiguration struct {
	VZKeyboardConfiguration
}

// VZUSBKeyboardConfigurationFrom constructs a [VZUSBKeyboardConfiguration] from an unsafe.Pointer.
//
// A device that defines the configuration for a USB keyboard.
func VZUSBKeyboardConfigurationFrom(ptr unsafe.Pointer) VZUSBKeyboardConfiguration {
	return VZUSBKeyboardConfiguration{
		VZKeyboardConfiguration: VZKeyboardConfigurationFrom(ptr),
	}
}

/* debug [class_struct]: End struct */

/* debug [class_init_methods]: Init methods for VZUSBKeyboardConfiguration */
/* debug [class_init_methods]: End init methods */

/* debug [class_methods]: Class methods for VZUSBKeyboardConfiguration */
/* debug [class_methods]: End class methods */

/* debug [class_properties_class]: Class properties for VZUSBKeyboardConfiguration */
/* debug [class_properties_class]: End class properties */

/* debug [instance_methods]: Instance methods for VZUSBKeyboardConfiguration */
/* debug [instance_methods]: End instance methods */

/* debug [instance_properties]: Instance properties for VZUSBKeyboardConfiguration */
/* debug [instance_properties]: End instance properties */

/* debug [class.gen.go]: End class VZUSBKeyboardConfiguration */
