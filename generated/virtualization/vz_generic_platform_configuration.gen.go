// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class VZGenericPlatformConfiguration */


/* debug [class_header]: Header for VZGenericPlatformConfiguration */
// The class instance for the [VZGenericPlatformConfiguration] class.
var (
	VZGenericPlatformConfigurationClass     _VZGenericPlatformConfigurationClass
	VZGenericPlatformConfigurationClassOnce sync.Once
)

func getVZGenericPlatformConfigurationClass() _VZGenericPlatformConfigurationClass {
	VZGenericPlatformConfigurationClassOnce.Do(func() {
		VZGenericPlatformConfigurationClass = _VZGenericPlatformConfigurationClass{objc.GetClass("VZGenericPlatformConfiguration")}
	})
	return VZGenericPlatformConfigurationClass
}

type _VZGenericPlatformConfigurationClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for VZGenericPlatformConfiguration */
// An interface definition for the [VZGenericPlatformConfiguration] class.
type IVZGenericPlatformConfiguration interface {
	IVZPlatformConfiguration
	
/* debug [class_interface_properties]: Properties for VZGenericPlatformConfiguration */
	// properties:
	NestedVirtualizationEnabled() bool
	SetNestedVirtualizationEnabled(value bool)
	MachineIdentifier() IVZGenericMachineIdentifier
	SetMachineIdentifier(value IVZGenericMachineIdentifier)
	IsNestedVirtualizationEnabled() bool
	SetIsNestedVirtualizationEnabled(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for VZGenericPlatformConfiguration */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for VZGenericPlatformConfiguration */
// Alloc allocates a new instance without initialization.
func (vc _VZGenericPlatformConfigurationClass) Alloc() VZGenericPlatformConfiguration {
	rv := objc.Send[VZGenericPlatformConfiguration](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (vc _VZGenericPlatformConfigurationClass) New() VZGenericPlatformConfiguration {
	rv := objc.Send[VZGenericPlatformConfiguration](objc.ID(vc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (v_ VZGenericPlatformConfiguration) Init() VZGenericPlatformConfiguration {
	rv := objc.Send[VZGenericPlatformConfiguration](v_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v_ VZGenericPlatformConfiguration) Autorelease() VZGenericPlatformConfiguration {
	rv := objc.Send[VZGenericPlatformConfiguration](v_.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZGenericPlatformConfiguration creates a new VZGenericPlatformConfiguration instance.
func NewVZGenericPlatformConfiguration() VZGenericPlatformConfiguration {
	return getVZGenericPlatformConfigurationClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for VZGenericPlatformConfiguration */
// The platform configuration for a generic Intel or ARM virtual machine.


// The platform configuration for a generic Intel or ARM virtual machine.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZGenericPlatformConfiguration
type VZGenericPlatformConfiguration struct {
	VZPlatformConfiguration
}

// VZGenericPlatformConfigurationFrom constructs a [VZGenericPlatformConfiguration] from an unsafe.Pointer.
//
// The platform configuration for a generic Intel or ARM virtual machine.
func VZGenericPlatformConfigurationFrom(ptr unsafe.Pointer) VZGenericPlatformConfiguration {
	return VZGenericPlatformConfiguration{
		VZPlatformConfiguration: VZPlatformConfigurationFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for VZGenericPlatformConfiguration */
/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for VZGenericPlatformConfiguration */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for VZGenericPlatformConfiguration */

// A Boolean value that describes whether the platform configuration supports nested virtualization.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZGenericPlatformConfiguration/isNestedVirtualizationSupported
func (vc _VZGenericPlatformConfigurationClass) NestedVirtualizationSupported() bool {
	rv := objc.Send[bool](objc.ID(vc.class), objc.Sel("nestedVirtualizationSupported"))
	return rv
}/* debug [class_properties_class/property]: nestedVirtualizationSupported */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for VZGenericPlatformConfiguration */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for VZGenericPlatformConfiguration */

// A Boolean value that indicates whether nested virtualization is in an enabled state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZGenericPlatformConfiguration/isNestedVirtualizationEnabled
func (v_ VZGenericPlatformConfiguration) NestedVirtualizationEnabled() bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("nestedVirtualizationEnabled"))
	return rv
}/* debug [instance_properties/getter]: nestedVirtualizationEnabled */


// A Boolean value that indicates whether nested virtualization is in an enabled state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZGenericPlatformConfiguration/isNestedVirtualizationEnabled
func (v_ VZGenericPlatformConfiguration) SetNestedVirtualizationEnabled(value bool) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setNestedVirtualizationEnabled:"), value)
}/* debug [instance_properties/setter]: nestedVirtualizationEnabled */


// A Boolean value that describes whether the platform configuration supports nested virtualization.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZGenericPlatformConfiguration/isNestedVirtualizationSupported
func (v_ VZGenericPlatformConfiguration) NestedVirtualizationSupported() bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("nestedVirtualizationSupported"))
	return rv
}/* debug [instance_properties/getter]: nestedVirtualizationSupported */


// A value that represents a unique identifier for the virtual machine.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZGenericPlatformConfiguration/machineIdentifier
func (v_ VZGenericPlatformConfiguration) MachineIdentifier() IVZGenericMachineIdentifier {
	rv := objc.Send[VZGenericMachineIdentifier](v_.ID, objc.Sel("machineIdentifier"))
	return rv
}/* debug [instance_properties/getter]: machineIdentifier */


// A value that represents a unique identifier for the virtual machine.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZGenericPlatformConfiguration/machineIdentifier
func (v_ VZGenericPlatformConfiguration) SetMachineIdentifier(value IVZGenericMachineIdentifier) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setMachineIdentifier:"), value)
}/* debug [instance_properties/setter]: machineIdentifier */


// A Boolean value that indicates whether nested virtualization is in an enabled state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/virtualization/vzgenericplatformconfiguration/isnestedvirtualizationenabled
func (v_ VZGenericPlatformConfiguration) IsNestedVirtualizationEnabled() bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("isNestedVirtualizationEnabled"))
	return rv
}/* debug [instance_properties/getter]: isNestedVirtualizationEnabled */


// A Boolean value that indicates whether nested virtualization is in an enabled state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/virtualization/vzgenericplatformconfiguration/isnestedvirtualizationenabled
func (v_ VZGenericPlatformConfiguration) SetIsNestedVirtualizationEnabled(value bool) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setIsNestedVirtualizationEnabled:"), value)
}/* debug [instance_properties/setter]: isNestedVirtualizationEnabled */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class VZGenericPlatformConfiguration */


