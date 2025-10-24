// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class VZVirtioTraditionalMemoryBalloonDevice */


/* debug [class_header]: Header for VZVirtioTraditionalMemoryBalloonDevice */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for VZVirtioTraditionalMemoryBalloonDevice */
// An interface definition for the [VZVirtioTraditionalMemoryBalloonDevice] class.
type IVZVirtioTraditionalMemoryBalloonDevice interface {
	IVZMemoryBalloonDevice
	
/* debug [class_interface_properties]: Properties for VZVirtioTraditionalMemoryBalloonDevice */
	// properties:
	TargetVirtualMachineMemorySize() uint64
	SetTargetVirtualMachineMemorySize(value uint64)
	MemoryBalloonDevices() IVZMemoryBalloonDevice
	SetMemoryBalloonDevices(value IVZMemoryBalloonDevice)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for VZVirtioTraditionalMemoryBalloonDevice */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for VZVirtioTraditionalMemoryBalloonDevice */
// Alloc allocates a new instance without initialization.
func (vc _VZVirtioTraditionalMemoryBalloonDeviceClass) Alloc() VZVirtioTraditionalMemoryBalloonDevice {
	rv := objc.Send[VZVirtioTraditionalMemoryBalloonDevice](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for VZVirtioTraditionalMemoryBalloonDevice */
// The object you use to change the amount of memory allocated to the guest system.
//
// A object implements a Virtio-compliant balloon memory device, which lets you change the amount of physical memory assigned to the guest operating system. The virtual machine has no insight into the amount of memory its guest operating system uses. A memory balloon device lets you ask the guest operating system to relinquish memory voluntarily, which you might do if memory resources become scarce. You don’t create a object directly. Instead, create a object and assign it to the property of your virtual machine configuration. In response, the virtual machine creates this object and assigns it to its property. To use a memory balloon device, change the value in the property when your virtual machine is running. If the new value is smaller than the amount of currently assigned memory, the guest system may return a list of unused memory pages using the memory balloon device. If it does, the virtual machine releases those pages back to the host computer. If it doesn’t return any memory pages, the virtual machine leaves the guest’s memory size unchanged. If the new value is larger than the amount of currently assigned memory, the virtual machine reserves more pages for the guest operating system. For optimal performance, the guest operating system should compact its memory before returning any pages back to the memory balloon device. Compacting the memory reduces fragmentation, and allows it to return contiguous blocks of free pages in the memory balloon device.


// The object you use to change the amount of memory allocated to the guest system.
//
// [Full Topic]
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for VZVirtioTraditionalMemoryBalloonDevice *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for VZVirtioTraditionalMemoryBalloonDevice */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for VZVirtioTraditionalMemoryBalloonDevice */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for VZVirtioTraditionalMemoryBalloonDevice */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for VZVirtioTraditionalMemoryBalloonDevice */

// The target amount of memory, in bytes, to make available to the virtual machine.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtioTraditionalMemoryBalloonDevice/targetVirtualMachineMemorySize
func (v_ VZVirtioTraditionalMemoryBalloonDevice) TargetVirtualMachineMemorySize() uint64 {
	rv := objc.Send[uint64](v_.ID, objc.Sel("targetVirtualMachineMemorySize"))
	return rv
}/* debug [instance_properties/getter]: targetVirtualMachineMemorySize */


// The target amount of memory, in bytes, to make available to the virtual machine.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtioTraditionalMemoryBalloonDevice/targetVirtualMachineMemorySize
func (v_ VZVirtioTraditionalMemoryBalloonDevice) SetTargetVirtualMachineMemorySize(value uint64) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setTargetVirtualMachineMemorySize:"), value)
}/* debug [instance_properties/setter]: targetVirtualMachineMemorySize */


// The array of devices that you use to adjust the amount of memory available to the guest system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/virtualization/vzvirtualmachine/memoryballoondevices
func (v_ VZVirtioTraditionalMemoryBalloonDevice) MemoryBalloonDevices() IVZMemoryBalloonDevice {
	rv := objc.Send[VZMemoryBalloonDevice](v_.ID, objc.Sel("memoryBalloonDevices"))
	return rv
}/* debug [instance_properties/getter]: memoryBalloonDevices */


// The array of devices that you use to adjust the amount of memory available to the guest system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/virtualization/vzvirtualmachine/memoryballoondevices
func (v_ VZVirtioTraditionalMemoryBalloonDevice) SetMemoryBalloonDevices(value IVZMemoryBalloonDevice) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setMemoryBalloonDevices:"), value)
}/* debug [instance_properties/setter]: memoryBalloonDevices */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class VZVirtioTraditionalMemoryBalloonDevice */



