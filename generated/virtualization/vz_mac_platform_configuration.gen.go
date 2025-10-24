// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class VZMacPlatformConfiguration */

/* debug [class_header]: Header for VZMacPlatformConfiguration */
// The class instance for the [VZMacPlatformConfiguration] class.
var (
	VZMacPlatformConfigurationClass     _VZMacPlatformConfigurationClass
	VZMacPlatformConfigurationClassOnce sync.Once
)

func getVZMacPlatformConfigurationClass() _VZMacPlatformConfigurationClass {
	VZMacPlatformConfigurationClassOnce.Do(func() {
		VZMacPlatformConfigurationClass = _VZMacPlatformConfigurationClass{objc.GetClass("VZMacPlatformConfiguration")}
	})
	return VZMacPlatformConfigurationClass
}

type _VZMacPlatformConfigurationClass struct {
	class objc.Class
}

/* debug [class_header]: End header */

/* debug [class_interface]: Interface for VZMacPlatformConfiguration */
// An interface definition for the [VZMacPlatformConfiguration] class.
type IVZMacPlatformConfiguration interface {
	IVZPlatformConfiguration

	/* debug [class_interface_properties]: Properties for VZMacPlatformConfiguration */
	// properties:
	AuxiliaryStorage() IVZMacAuxiliaryStorage
	SetAuxiliaryStorage(value IVZMacAuxiliaryStorage)
	HardwareModel() IVZMacHardwareModel
	SetHardwareModel(value IVZMacHardwareModel)
	MachineIdentifier() IVZMacMachineIdentifier
	SetMachineIdentifier(value IVZMacMachineIdentifier)
	MostFeaturefulSupportedConfiguration() IVZMacOSConfigurationRequirements
	SetMostFeaturefulSupportedConfiguration(value IVZMacOSConfigurationRequirements)
	/* debug [class_interface_properties]: End properties */

	/* debug [class_interface_methods]: Methods for VZMacPlatformConfiguration */
	// methods:
	/* debug [class_interface_methods]: End methods */

}

/* debug [class_interface]: End interface */

/* debug [class_constructors]: Constructors for VZMacPlatformConfiguration */
// Alloc allocates a new instance without initialization.
func (vc _VZMacPlatformConfigurationClass) Alloc() VZMacPlatformConfiguration {
	rv := objc.Send[VZMacPlatformConfiguration](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (vc _VZMacPlatformConfigurationClass) New() VZMacPlatformConfiguration {
	rv := objc.Send[VZMacPlatformConfiguration](objc.ID(vc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (v_ VZMacPlatformConfiguration) Init() VZMacPlatformConfiguration {
	rv := objc.Send[VZMacPlatformConfiguration](v_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v_ VZMacPlatformConfiguration) Autorelease() VZMacPlatformConfiguration {
	rv := objc.Send[VZMacPlatformConfiguration](v_.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZMacPlatformConfiguration creates a new VZMacPlatformConfiguration instance.
func NewVZMacPlatformConfiguration() VZMacPlatformConfiguration {
	return getVZMacPlatformConfigurationClass().New()
}

/* debug [class_constructors]: End constructors */

/* debug [class_struct]: Struct for VZMacPlatformConfiguration */
// The platform configuration for booting macOS on Apple silicon.
//
// When creating a VM, the and depend on the restore image that you use to install macOS. To choose the hardware model, start from . to get a supported configuration, then use its . property to get the hardware model. Use the hardware model to set up and to initialize a new auxiliary storage with . When you save a VM to disk and load it again, you must restore the , and properties to their original values. If you create multiple VMs from the same configuration, each should have a unique and .

// The platform configuration for booting macOS on Apple silicon.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZMacPlatformConfiguration
type VZMacPlatformConfiguration struct {
	VZPlatformConfiguration
}

// VZMacPlatformConfigurationFrom constructs a [VZMacPlatformConfiguration] from an unsafe.Pointer.
//
// The platform configuration for booting macOS on Apple silicon.
func VZMacPlatformConfigurationFrom(ptr unsafe.Pointer) VZMacPlatformConfiguration {
	return VZMacPlatformConfiguration{
		VZPlatformConfiguration: VZPlatformConfigurationFrom(ptr),
	}
}

/* debug [class_struct]: End struct */

/* debug [class_init_methods]: Init methods for VZMacPlatformConfiguration */
/* debug [class_init_methods]: End init methods */

/* debug [class_methods]: Class methods for VZMacPlatformConfiguration */
/* debug [class_methods]: End class methods */

/* debug [class_properties_class]: Class properties for VZMacPlatformConfiguration */
/* debug [class_properties_class]: End class properties */

/* debug [instance_methods]: Instance methods for VZMacPlatformConfiguration */
/* debug [instance_methods]: End instance methods */

/* debug [instance_properties]: Instance properties for VZMacPlatformConfiguration */

// The Mac auxiliary storage.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZMacPlatformConfiguration/auxiliaryStorage
func (v_ VZMacPlatformConfiguration) AuxiliaryStorage() IVZMacAuxiliaryStorage {
	rv := objc.Send[VZMacAuxiliaryStorage](v_.ID, objc.Sel("auxiliaryStorage"))
	return rv
} /* debug [instance_properties/getter]: auxiliaryStorage */

// The Mac auxiliary storage.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZMacPlatformConfiguration/auxiliaryStorage
func (v_ VZMacPlatformConfiguration) SetAuxiliaryStorage(value IVZMacAuxiliaryStorage) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setAuxiliaryStorage:"), value)
} /* debug [instance_properties/setter]: auxiliaryStorage */

// The Mac hardware model.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZMacPlatformConfiguration/hardwareModel
func (v_ VZMacPlatformConfiguration) HardwareModel() IVZMacHardwareModel {
	rv := objc.Send[VZMacHardwareModel](v_.ID, objc.Sel("hardwareModel"))
	return rv
} /* debug [instance_properties/getter]: hardwareModel */

// The Mac hardware model.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZMacPlatformConfiguration/hardwareModel
func (v_ VZMacPlatformConfiguration) SetHardwareModel(value IVZMacHardwareModel) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setHardwareModel:"), value)
} /* debug [instance_properties/setter]: hardwareModel */

// The Mac machine identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZMacPlatformConfiguration/machineIdentifier
func (v_ VZMacPlatformConfiguration) MachineIdentifier() IVZMacMachineIdentifier {
	rv := objc.Send[VZMacMachineIdentifier](v_.ID, objc.Sel("machineIdentifier"))
	return rv
} /* debug [instance_properties/getter]: machineIdentifier */

// The Mac machine identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZMacPlatformConfiguration/machineIdentifier
func (v_ VZMacPlatformConfiguration) SetMachineIdentifier(value IVZMacMachineIdentifier) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setMachineIdentifier:"), value)
} /* debug [instance_properties/setter]: machineIdentifier */

// This object represents the most fully featured configuration that’s supported by both the current host and by this restore image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/virtualization/vzmacosrestoreimage/mostfeaturefulsupportedconfiguration
func (v_ VZMacPlatformConfiguration) MostFeaturefulSupportedConfiguration() IVZMacOSConfigurationRequirements {
	rv := objc.Send[VZMacOSConfigurationRequirements](v_.ID, objc.Sel("mostFeaturefulSupportedConfiguration"))
	return rv
} /* debug [instance_properties/getter]: mostFeaturefulSupportedConfiguration */

// This object represents the most fully featured configuration that’s supported by both the current host and by this restore image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/virtualization/vzmacosrestoreimage/mostfeaturefulsupportedconfiguration
func (v_ VZMacPlatformConfiguration) SetMostFeaturefulSupportedConfiguration(value IVZMacOSConfigurationRequirements) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setMostFeaturefulSupportedConfiguration:"), value)
} /* debug [instance_properties/setter]: mostFeaturefulSupportedConfiguration */

/* debug [instance_properties]: End instance properties */

/* debug [class.gen.go]: End class VZMacPlatformConfiguration */
