// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [VZMemoryBalloonDevice] class.
var (
	VZMemoryBalloonDeviceClass     _VZMemoryBalloonDeviceClass
	VZMemoryBalloonDeviceClassOnce sync.Once
)

func getVZMemoryBalloonDeviceClass() _VZMemoryBalloonDeviceClass {
	VZMemoryBalloonDeviceClassOnce.Do(func() {
		VZMemoryBalloonDeviceClass = _VZMemoryBalloonDeviceClass{objc.GetClass("VZMemoryBalloonDevice")}
	})
	return VZMemoryBalloonDeviceClass
}

type _VZMemoryBalloonDeviceClass struct {
	class objc.Class
}

// An interface definition for the [VZMemoryBalloonDevice] class.
type IVZMemoryBalloonDevice interface {
	objectivec.IObject
	MemoryBalloonDevices() VZMemoryBalloonDeviceConfiguration
	SetMemoryBalloonDevices(value IVZMemoryBalloonDeviceConfiguration)
}

// The common behavior for memory devices.
//
// Don’t instantiate this class directly. To request a memory ballon device, add an appropriate configuration object to the property of the object that you use to configure the virtual machine. In response, the system instantiates the subclass of that matches your request. For example, if you supply a object in your configuration, the system creates a object.
//
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZMemoryBalloonDevice
type VZMemoryBalloonDevice struct {
	objectivec.Object
}

// VZMemoryBalloonDeviceFrom constructs a [VZMemoryBalloonDevice] from an unsafe.Pointer.
//
// The common behavior for memory devices.
func VZMemoryBalloonDeviceFrom(ptr unsafe.Pointer) VZMemoryBalloonDevice {
	return VZMemoryBalloonDevice{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (vc _VZMemoryBalloonDeviceClass) Alloc() VZMemoryBalloonDevice {
	rv := objc.Send[VZMemoryBalloonDevice](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (vc _VZMemoryBalloonDeviceClass) New() VZMemoryBalloonDevice {
	rv := objc.Send[VZMemoryBalloonDevice](objc.ID(vc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (v_ VZMemoryBalloonDevice) Init() VZMemoryBalloonDevice {
	rv := objc.Send[VZMemoryBalloonDevice](v_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v_ VZMemoryBalloonDevice) Autorelease() VZMemoryBalloonDevice {
	rv := objc.Send[VZMemoryBalloonDevice](v_.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZMemoryBalloonDevice creates a new VZMemoryBalloonDevice instance.
func NewVZMemoryBalloonDevice() VZMemoryBalloonDevice {
	return getVZMemoryBalloonDeviceClass().New()
}


// An array that you configure with a memory balloon device, used to update the memory in the VM.
//
// [Full Topic]: https://developer.apple.com/documentation/virtualization/vzvirtualmachineconfiguration/memoryballoondevices
func (v_ VZMemoryBalloonDevice) MemoryBalloonDevices() VZMemoryBalloonDeviceConfiguration {
	rv := objc.Send[VZMemoryBalloonDeviceConfiguration](v_.ID, objc.Sel("memoryBalloonDevices"))
	return rv
}


// SetMemoryBalloonDevices sets the value of the memoryBalloonDevices property.
// An array that you configure with a memory balloon device, used to update the memory in the VM.

//
// [Full Topic]: https://developer.apple.com/documentation/virtualization/vzvirtualmachineconfiguration/memoryballoondevices
func (v_ VZMemoryBalloonDevice) SetMemoryBalloonDevices(value IVZMemoryBalloonDeviceConfiguration) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setMemoryBalloonDevices:"), value)
}



