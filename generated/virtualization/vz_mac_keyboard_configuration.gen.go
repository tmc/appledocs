// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class VZMacKeyboardConfiguration */

/* debug [class_header]: Header for VZMacKeyboardConfiguration */
// The class instance for the [VZMacKeyboardConfiguration] class.
var (
	VZMacKeyboardConfigurationClass     _VZMacKeyboardConfigurationClass
	VZMacKeyboardConfigurationClassOnce sync.Once
)

func getVZMacKeyboardConfigurationClass() _VZMacKeyboardConfigurationClass {
	VZMacKeyboardConfigurationClassOnce.Do(func() {
		VZMacKeyboardConfigurationClass = _VZMacKeyboardConfigurationClass{objc.GetClass("VZMacKeyboardConfiguration")}
	})
	return VZMacKeyboardConfigurationClass
}

type _VZMacKeyboardConfigurationClass struct {
	class objc.Class
}

/* debug [class_header]: End header */

/* debug [class_interface]: Interface for VZMacKeyboardConfiguration */
// An interface definition for the [VZMacKeyboardConfiguration] class.
type IVZMacKeyboardConfiguration interface {
	IVZKeyboardConfiguration

	/* debug [class_interface_properties]: Properties for VZMacKeyboardConfiguration */
	// properties:
	/* debug [class_interface_properties]: End properties */

	/* debug [class_interface_methods]: Methods for VZMacKeyboardConfiguration */
	// methods:
	/* debug [class_interface_methods]: End methods */

}

/* debug [class_interface]: End interface */

/* debug [class_constructors]: Constructors for VZMacKeyboardConfiguration */
// Alloc allocates a new instance without initialization.
func (vc _VZMacKeyboardConfigurationClass) Alloc() VZMacKeyboardConfiguration {
	rv := objc.Send[VZMacKeyboardConfiguration](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (vc _VZMacKeyboardConfigurationClass) New() VZMacKeyboardConfiguration {
	rv := objc.Send[VZMacKeyboardConfiguration](objc.ID(vc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (v_ VZMacKeyboardConfiguration) Init() VZMacKeyboardConfiguration {
	rv := objc.Send[VZMacKeyboardConfiguration](v_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v_ VZMacKeyboardConfiguration) Autorelease() VZMacKeyboardConfiguration {
	rv := objc.Send[VZMacKeyboardConfiguration](v_.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZMacKeyboardConfiguration creates a new VZMacKeyboardConfiguration instance.
func NewVZMacKeyboardConfiguration() VZMacKeyboardConfiguration {
	return getVZMacKeyboardConfigurationClass().New()
}

/* debug [class_constructors]: End constructors */

/* debug [class_struct]: Struct for VZMacKeyboardConfiguration */
// A device that defines the configuration for a Mac keyboard.
//
// Use this configuration to attach a Mac keyboard configuration to a VM. A can use this device to send key events to the VM, including the Mac-specific key events, such as the Globe key.

// A device that defines the configuration for a Mac keyboard.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZMacKeyboardConfiguration
type VZMacKeyboardConfiguration struct {
	VZKeyboardConfiguration
}

// VZMacKeyboardConfigurationFrom constructs a [VZMacKeyboardConfiguration] from an unsafe.Pointer.
//
// A device that defines the configuration for a Mac keyboard.
func VZMacKeyboardConfigurationFrom(ptr unsafe.Pointer) VZMacKeyboardConfiguration {
	return VZMacKeyboardConfiguration{
		VZKeyboardConfiguration: VZKeyboardConfigurationFrom(ptr),
	}
}

/* debug [class_struct]: End struct */

/* debug [class_init_methods]: Init methods for VZMacKeyboardConfiguration */
/* debug [class_init_methods]: End init methods */

/* debug [class_methods]: Class methods for VZMacKeyboardConfiguration */
/* debug [class_methods]: End class methods */

/* debug [class_properties_class]: Class properties for VZMacKeyboardConfiguration */
/* debug [class_properties_class]: End class properties */

/* debug [instance_methods]: Instance methods for VZMacKeyboardConfiguration */
/* debug [instance_methods]: End instance methods */

/* debug [instance_properties]: Instance properties for VZMacKeyboardConfiguration */
/* debug [instance_properties]: End instance properties */

/* debug [class.gen.go]: End class VZMacKeyboardConfiguration */
