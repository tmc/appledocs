// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [VZMacOSConfigurationRequirements] class.
type IVZMacOSConfigurationRequirements interface {
	objectivec.IObject
}

// An object that describes the parameter constraints required by a specific configuration of macOS.
//
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

// Alloc allocates a new instance without initialization.
func (vc _VZMacOSConfigurationRequirementsClass) Alloc() VZMacOSConfigurationRequirements {
	rv := objc.Send[VZMacOSConfigurationRequirements](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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


// The hardware model for this configuration.
//
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZMacOSConfigurationRequirements/hardwareModel
func (v_ VZMacOSConfigurationRequirements) HardwareModel() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID, objc.Sel("hardwareModel"))
	return rv
}

// The minimum supported number of CPUs for this configuration.
//
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZMacOSConfigurationRequirements/minimumSupportedCPUCount
func (v_ VZMacOSConfigurationRequirements) MinimumSupportedCPUCount() uint {
	rv := objc.Send[uint](v_.ID, objc.Sel("minimumSupportedCPUCount"))
	return rv
}

// The minimum supported memory size for this configuration.
//
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZMacOSConfigurationRequirements/minimumSupportedMemorySize
func (v_ VZMacOSConfigurationRequirements) MinimumSupportedMemorySize() uint64 {
	rv := objc.Send[uint64](v_.ID, objc.Sel("minimumSupportedMemorySize"))
	return rv
}



