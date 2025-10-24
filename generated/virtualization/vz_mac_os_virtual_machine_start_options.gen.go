// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [VZMacOSVirtualMachineStartOptions] class.
var (
	VZMacOSVirtualMachineStartOptionsClass     _VZMacOSVirtualMachineStartOptionsClass
	VZMacOSVirtualMachineStartOptionsClassOnce sync.Once
)

func getVZMacOSVirtualMachineStartOptionsClass() _VZMacOSVirtualMachineStartOptionsClass {
	VZMacOSVirtualMachineStartOptionsClassOnce.Do(func() {
		VZMacOSVirtualMachineStartOptionsClass = _VZMacOSVirtualMachineStartOptionsClass{objc.GetClass("VZMacOSVirtualMachineStartOptions")}
	})
	return VZMacOSVirtualMachineStartOptionsClass
}

type _VZMacOSVirtualMachineStartOptionsClass struct {
	class objc.Class
}

// An interface definition for the [VZMacOSVirtualMachineStartOptions] class.
type IVZMacOSVirtualMachineStartOptions interface {
	IVZVirtualMachineStartOptions
	// properties:
	StartUpFromMacOSRecovery() bool
	SetStartUpFromMacOSRecovery(value bool)
	// methods:
}

// A class that describes start options for macOS VMs.


// A class that describes start options for macOS VMs.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZMacOSVirtualMachineStartOptions
type VZMacOSVirtualMachineStartOptions struct {
	VZVirtualMachineStartOptions
}

// VZMacOSVirtualMachineStartOptionsFrom constructs a [VZMacOSVirtualMachineStartOptions] from an unsafe.Pointer.
//
// A class that describes start options for macOS VMs.
func VZMacOSVirtualMachineStartOptionsFrom(ptr unsafe.Pointer) VZMacOSVirtualMachineStartOptions {
	return VZMacOSVirtualMachineStartOptions{
		VZVirtualMachineStartOptions: VZVirtualMachineStartOptionsFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (vc _VZMacOSVirtualMachineStartOptionsClass) Alloc() VZMacOSVirtualMachineStartOptions {
	rv := objc.Send[VZMacOSVirtualMachineStartOptions](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (vc _VZMacOSVirtualMachineStartOptionsClass) New() VZMacOSVirtualMachineStartOptions {
	rv := objc.Send[VZMacOSVirtualMachineStartOptions](objc.ID(vc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (v_ VZMacOSVirtualMachineStartOptions) Init() VZMacOSVirtualMachineStartOptions {
	rv := objc.Send[VZMacOSVirtualMachineStartOptions](v_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v_ VZMacOSVirtualMachineStartOptions) Autorelease() VZMacOSVirtualMachineStartOptions {
	rv := objc.Send[VZMacOSVirtualMachineStartOptions](v_.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZMacOSVirtualMachineStartOptions creates a new VZMacOSVirtualMachineStartOptions instance.
func NewVZMacOSVirtualMachineStartOptions() VZMacOSVirtualMachineStartOptions {
	return getVZMacOSVirtualMachineStartOptionsClass().New()
}



// A Boolean value that indicates whether the macOS guest should start in recovery mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZMacOSVirtualMachineStartOptions/startUpFromMacOSRecovery
func (v_ VZMacOSVirtualMachineStartOptions) StartUpFromMacOSRecovery() bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("startUpFromMacOSRecovery"))
	return rv
}


// A Boolean value that indicates whether the macOS guest should start in recovery mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZMacOSVirtualMachineStartOptions/startUpFromMacOSRecovery
func (v_ VZMacOSVirtualMachineStartOptions) SetStartUpFromMacOSRecovery(value bool) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setStartUpFromMacOSRecovery:"), value)
}



