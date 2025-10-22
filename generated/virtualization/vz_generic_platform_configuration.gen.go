// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

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

// An interface definition for the [VZGenericPlatformConfiguration] class.
type IVZGenericPlatformConfiguration interface {
	IVZPlatformConfiguration
	NestedVirtualizationEnabled() bool
	SetNestedVirtualizationEnabled(value bool)
	MachineIdentifier() VZGenericMachineIdentifier
	SetMachineIdentifier(value IVZGenericMachineIdentifier)
	IsNestedVirtualizationEnabled() bool
	SetIsNestedVirtualizationEnabled(value bool)
}

// The platform configuration for a generic Intel or ARM virtual machine.
//
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

// Alloc allocates a new instance without initialization.
func (vc _VZGenericPlatformConfigurationClass) Alloc() VZGenericPlatformConfiguration {
	rv := objc.Send[VZGenericPlatformConfiguration](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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



// A Boolean value that describes whether the platform configuration supports nested virtualization.
//
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZGenericPlatformConfiguration/isNestedVirtualizationSupported
func (vc _VZGenericPlatformConfigurationClass) NestedVirtualizationSupported() bool {
	rv := objc.Send[bool](objc.ID(vc.class), objc.Sel("nestedVirtualizationSupported"))
	return rv
}
// A Boolean value that indicates whether nested virtualization is in an enabled state.
//
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZGenericPlatformConfiguration/isNestedVirtualizationEnabled
func (v_ VZGenericPlatformConfiguration) NestedVirtualizationEnabled() bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("nestedVirtualizationEnabled"))
	return rv
}


// SetNestedVirtualizationEnabled sets the value of the nestedVirtualizationEnabled property.
// A Boolean value that indicates whether nested virtualization is in an enabled state.

//
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZGenericPlatformConfiguration/isNestedVirtualizationEnabled
func (v_ VZGenericPlatformConfiguration) SetNestedVirtualizationEnabled(value bool) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setNestedVirtualizationEnabled:"), value)
}

// A Boolean value that describes whether the platform configuration supports nested virtualization.
//
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZGenericPlatformConfiguration/isNestedVirtualizationSupported
func (v_ VZGenericPlatformConfiguration) NestedVirtualizationSupported() bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("nestedVirtualizationSupported"))
	return rv
}

// A value that represents a unique identifier for the virtual machine.
//
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZGenericPlatformConfiguration/machineIdentifier
func (v_ VZGenericPlatformConfiguration) MachineIdentifier() VZGenericMachineIdentifier {
	rv := objc.Send[VZGenericMachineIdentifier](v_.ID, objc.Sel("machineIdentifier"))
	return rv
}


// SetMachineIdentifier sets the value of the machineIdentifier property.
// A value that represents a unique identifier for the virtual machine.

//
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZGenericPlatformConfiguration/machineIdentifier
func (v_ VZGenericPlatformConfiguration) SetMachineIdentifier(value IVZGenericMachineIdentifier) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setMachineIdentifier:"), value)
}

// A Boolean value that indicates whether nested virtualization is in an enabled state.
//
// [Full Topic]: https://developer.apple.com/documentation/virtualization/vzgenericplatformconfiguration/isnestedvirtualizationenabled
func (v_ VZGenericPlatformConfiguration) IsNestedVirtualizationEnabled() bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("isNestedVirtualizationEnabled"))
	return rv
}


// SetIsNestedVirtualizationEnabled sets the value of the isNestedVirtualizationEnabled property.
// A Boolean value that indicates whether nested virtualization is in an enabled state.

//
// [Full Topic]: https://developer.apple.com/documentation/virtualization/vzgenericplatformconfiguration/isnestedvirtualizationenabled
func (v_ VZGenericPlatformConfiguration) SetIsNestedVirtualizationEnabled(value bool) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setIsNestedVirtualizationEnabled:"), value)
}


