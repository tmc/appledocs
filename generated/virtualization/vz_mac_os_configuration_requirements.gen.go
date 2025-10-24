// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class VZMacOSConfigurationRequirements */


/* debug [class_header]: Header for VZMacOSConfigurationRequirements */
// The class instance for the [VZMacOSConfigurationRequirements] class.
var (
	VZMacOSConfigurationRequirementsClass     _VZMacOSConfigurationRequirementsClass
	VZMacOSConfigurationRequirementsClassOnce sync.Once
)

func getVZMacOSConfigurationRequirementsClass() _VZMacOSConfigurationRequirementsClass {
	VZMacOSConfigurationRequirementsClassOnce.Do(func() {
		VZMacOSConfigurationRequirementsClass = _VZMacOSConfigurationRequirementsClass{objc.GetClass("VZMacOSConfigurationRequirements")}
	})
	return VZMacOSConfigurationRequirementsClass
}

type _VZMacOSConfigurationRequirementsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for VZMacOSConfigurationRequirements */
// An interface definition for the [VZMacOSConfigurationRequirements] class.
type IVZMacOSConfigurationRequirements interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for VZMacOSConfigurationRequirements */
	// properties:
	HardwareModel() IVZMacHardwareModel
	MinimumSupportedCPUCount() uint
	MinimumSupportedMemorySize() uint64
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for VZMacOSConfigurationRequirements */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for VZMacOSConfigurationRequirements */
// Alloc allocates a new instance without initialization.
func (vc _VZMacOSConfigurationRequirementsClass) Alloc() VZMacOSConfigurationRequirements {
	rv := objc.Send[VZMacOSConfigurationRequirements](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (vc _VZMacOSConfigurationRequirementsClass) New() VZMacOSConfigurationRequirements {
	rv := objc.Send[VZMacOSConfigurationRequirements](objc.ID(vc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (v_ VZMacOSConfigurationRequirements) Init() VZMacOSConfigurationRequirements {
	rv := objc.Send[VZMacOSConfigurationRequirements](v_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v_ VZMacOSConfigurationRequirements) Autorelease() VZMacOSConfigurationRequirements {
	rv := objc.Send[VZMacOSConfigurationRequirements](v_.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZMacOSConfigurationRequirements creates a new VZMacOSConfigurationRequirements instance.
func NewVZMacOSConfigurationRequirements() VZMacOSConfigurationRequirements {
	return getVZMacOSConfigurationRequirementsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for VZMacOSConfigurationRequirements */
// An object that describes the parameter constraints required by a specific configuration of macOS.


// An object that describes the parameter constraints required by a specific configuration of macOS.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZMacOSConfigurationRequirements
type VZMacOSConfigurationRequirements struct {
	objectivec.Object
}

// VZMacOSConfigurationRequirementsFrom constructs a [VZMacOSConfigurationRequirements] from an unsafe.Pointer.
//
// An object that describes the parameter constraints required by a specific configuration of macOS.
func VZMacOSConfigurationRequirementsFrom(ptr unsafe.Pointer) VZMacOSConfigurationRequirements {
	return VZMacOSConfigurationRequirements{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for VZMacOSConfigurationRequirements *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for VZMacOSConfigurationRequirements */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for VZMacOSConfigurationRequirements */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for VZMacOSConfigurationRequirements */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for VZMacOSConfigurationRequirements */

// The hardware model for this configuration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZMacOSConfigurationRequirements/hardwareModel
func (v_ VZMacOSConfigurationRequirements) HardwareModel() IVZMacHardwareModel {
	rv := objc.Send[VZMacHardwareModel](v_.ID, objc.Sel("hardwareModel"))
	return rv
}/* debug [instance_properties/getter]: hardwareModel */


// The minimum supported number of CPUs for this configuration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZMacOSConfigurationRequirements/minimumSupportedCPUCount
func (v_ VZMacOSConfigurationRequirements) MinimumSupportedCPUCount() uint {
	rv := objc.Send[uint](v_.ID, objc.Sel("minimumSupportedCPUCount"))
	return rv
}/* debug [instance_properties/getter]: minimumSupportedCPUCount */


// The minimum supported memory size for this configuration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZMacOSConfigurationRequirements/minimumSupportedMemorySize
func (v_ VZMacOSConfigurationRequirements) MinimumSupportedMemorySize() uint64 {
	rv := objc.Send[uint64](v_.ID, objc.Sel("minimumSupportedMemorySize"))
	return rv
}/* debug [instance_properties/getter]: minimumSupportedMemorySize */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class VZMacOSConfigurationRequirements */



