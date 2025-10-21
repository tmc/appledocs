// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [VZVirtioTraditionalMemoryBalloonDevice] class.
var (
	VZVirtioTraditionalMemoryBalloonDeviceClass     _VZVirtioTraditionalMemoryBalloonDeviceClass
	VZVirtioTraditionalMemoryBalloonDeviceClassOnce sync.Once
)

func getVZVirtioTraditionalMemoryBalloonDeviceClass() _VZVirtioTraditionalMemoryBalloonDeviceClass {
	VZVirtioTraditionalMemoryBalloonDeviceClassOnce.Do(func() {
		VZVirtioTraditionalMemoryBalloonDeviceClass = _VZVirtioTraditionalMemoryBalloonDeviceClass{objc.GetClass("VZVirtioTraditionalMemoryBalloonDevice")}
	})
	return VZVirtioTraditionalMemoryBalloonDeviceClass
}

type _VZVirtioTraditionalMemoryBalloonDeviceClass struct {
	class objc.Class
}

// An interface definition for the [VZVirtioTraditionalMemoryBalloonDevice] class.
type IVZVirtioTraditionalMemoryBalloonDevice interface {
	IVZMemoryBalloonDevice
}

// The object you use to change the amount of memory allocated to the guest system.
//
// A object implements a Virtio-compliant balloon memory device, which lets you change the amount of physical memory assigned to the guest operating system. The virtual machine has no insight into the amount of memory its guest operating system uses. A memory balloon device lets you ask the guest operating system to relinquish memory voluntarily, which you might do if memory resources become scarce. You don’t create a object directly. Instead, create a object and assign it to the property of your virtual machine configuration. In response, the virtual machine creates this object and assigns it to its property. To use a memory balloon device, change the value in the property when your virtual machine is running. If the new value is smaller than the amount of currently assigned memory, the guest system may return a list of unused memory pages using the memory balloon device. If it does, the virtual machine releases those pages back to the host computer. If it doesn’t return any memory pages, the virtual machine leaves the guest’s memory size unchanged. If the new value is larger than the amount of currently assigned memory, the virtual machine reserves more pages for the guest operating system. For optimal performance, the guest operating system should compact its memory before returning any pages back to the memory balloon device. Compacting the memory reduces fragmentation, and allows it to return contiguous blocks of free pages in the memory balloon device.
//
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtioTraditionalMemoryBalloonDevice
type VZVirtioTraditionalMemoryBalloonDevice struct {
	VZMemoryBalloonDevice
}

// VZVirtioTraditionalMemoryBalloonDeviceFrom constructs a [VZVirtioTraditionalMemoryBalloonDevice] from an unsafe.Pointer.
//
// The object you use to change the amount of memory allocated to the guest system.
func VZVirtioTraditionalMemoryBalloonDeviceFrom(ptr unsafe.Pointer) VZVirtioTraditionalMemoryBalloonDevice {
	return VZVirtioTraditionalMemoryBalloonDevice{
		VZMemoryBalloonDevice: VZMemoryBalloonDeviceFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (vc _VZVirtioTraditionalMemoryBalloonDeviceClass) Alloc() VZVirtioTraditionalMemoryBalloonDevice {
	rv := objc.Send[VZVirtioTraditionalMemoryBalloonDevice](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (vc _VZVirtioTraditionalMemoryBalloonDeviceClass) New() VZVirtioTraditionalMemoryBalloonDevice {
	rv := objc.Send[VZVirtioTraditionalMemoryBalloonDevice](objc.ID(vc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (v_ VZVirtioTraditionalMemoryBalloonDevice) Init() VZVirtioTraditionalMemoryBalloonDevice {
	rv := objc.Send[VZVirtioTraditionalMemoryBalloonDevice](v_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v_ VZVirtioTraditionalMemoryBalloonDevice) Autorelease() VZVirtioTraditionalMemoryBalloonDevice {
	rv := objc.Send[VZVirtioTraditionalMemoryBalloonDevice](v_.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZVirtioTraditionalMemoryBalloonDevice creates a new VZVirtioTraditionalMemoryBalloonDevice instance.
func NewVZVirtioTraditionalMemoryBalloonDevice() VZVirtioTraditionalMemoryBalloonDevice {
	return getVZVirtioTraditionalMemoryBalloonDeviceClass().New()
}


// The target amount of memory, in bytes, to make available to the virtual machine.
//
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtioTraditionalMemoryBalloonDevice/targetVirtualMachineMemorySize
func (v_ VZVirtioTraditionalMemoryBalloonDevice) TargetVirtualMachineMemorySize() uint64 {
	rv := objc.Send[uint64](v_.ID, objc.Sel("targetVirtualMachineMemorySize"))
	return rv
}


// SetTargetVirtualMachineMemorySize sets the value of the targetVirtualMachineMemorySize property.
// The target amount of memory, in bytes, to make available to the virtual machine.

//
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtioTraditionalMemoryBalloonDevice/targetVirtualMachineMemorySize
func (v_ VZVirtioTraditionalMemoryBalloonDevice) SetTargetVirtualMachineMemorySize(value uint64) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setTargetVirtualMachineMemorySize:"), value)
}


