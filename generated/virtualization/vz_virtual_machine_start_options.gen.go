// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [VZVirtualMachineStartOptions] class.
var (
	VZVirtualMachineStartOptionsClass     _VZVirtualMachineStartOptionsClass
	VZVirtualMachineStartOptionsClassOnce sync.Once
)

func getVZVirtualMachineStartOptionsClass() _VZVirtualMachineStartOptionsClass {
	VZVirtualMachineStartOptionsClassOnce.Do(func() {
		VZVirtualMachineStartOptionsClass = _VZVirtualMachineStartOptionsClass{objc.GetClass("VZVirtualMachineStartOptions")}
	})
	return VZVirtualMachineStartOptionsClass
}

type _VZVirtualMachineStartOptionsClass struct {
	class objc.Class
}

// An interface definition for the [VZVirtualMachineStartOptions] class.
type IVZVirtualMachineStartOptions interface {
	objectivec.IObject
}

// The abstract class for VM start options.
//
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtualMachineStartOptions
type VZVirtualMachineStartOptions struct {
	objectivec.Object
}

// VZVirtualMachineStartOptionsFrom constructs a [VZVirtualMachineStartOptions] from an unsafe.Pointer.
//
// The abstract class for VM start options.
func VZVirtualMachineStartOptionsFrom(ptr unsafe.Pointer) VZVirtualMachineStartOptions {
	return VZVirtualMachineStartOptions{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (vc _VZVirtualMachineStartOptionsClass) Alloc() VZVirtualMachineStartOptions {
	rv := objc.Send[VZVirtualMachineStartOptions](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (vc _VZVirtualMachineStartOptionsClass) New() VZVirtualMachineStartOptions {
	rv := objc.Send[VZVirtualMachineStartOptions](objc.ID(vc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (v_ VZVirtualMachineStartOptions) Init() VZVirtualMachineStartOptions {
	rv := objc.Send[VZVirtualMachineStartOptions](v_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v_ VZVirtualMachineStartOptions) Autorelease() VZVirtualMachineStartOptions {
	rv := objc.Send[VZVirtualMachineStartOptions](v_.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZVirtualMachineStartOptions creates a new VZVirtualMachineStartOptions instance.
func NewVZVirtualMachineStartOptions() VZVirtualMachineStartOptions {
	return getVZVirtualMachineStartOptionsClass().New()
}




