// Code generated from Apple documentation for MLCompute. DO NOT EDIT.

package mlcompute

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [CDevice] class.
var (
	CDeviceClass     _CDeviceClass
	CDeviceClassOnce sync.Once
)

func getCDeviceClass() _CDeviceClass {
	CDeviceClassOnce.Do(func() {
		CDeviceClass = _CDeviceClass{objc.GetClass("MLCDevice")}
	})
	return CDeviceClass
}

type _CDeviceClass struct {
	class objc.Class
}

// An interface definition for the [CDevice] class.
type ICDevice interface {
	objectivec.IObject
}

// An object that represents the CPU or one or more GPUs the framework uses to execute a neural network.
//
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCDevice
type CDevice struct {
	objectivec.Object
}

// CDeviceFrom constructs a [CDevice] from an unsafe.Pointer.
//
// An object that represents the CPU or one or more GPUs the framework uses to execute a neural network.
func CDeviceFrom(ptr unsafe.Pointer) CDevice {
	return CDevice{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (cc _CDeviceClass) Alloc() CDevice {
	rv := objc.Send[CDevice](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CDeviceClass) New() CDevice {
	rv := objc.Send[CDevice](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CDevice) Init() CDevice {
	rv := objc.Send[CDevice](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CDevice) Autorelease() CDevice {
	rv := objc.Send[CDevice](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCDevice creates a new CDevice instance.
func NewCDevice() CDevice {
	return getCDeviceClass().New()
}




// Creates a device using the GPUs you specify.
//
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCDevice/init(gpuDevices:)
func NewCDeviceWithGPUDevices(gpus []objc.ID) CDevice {
	rv := objc.Send[CDevice](objc.ID(getCDeviceClass().class), objc.Sel("deviceWithGPUDevices:"), gpus)
	return rv
}



// Creates a device of the type you specify.
//
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCDevice/init(type:)
func NewCDeviceWithType(type_ CDeviceType) CDevice {
	rv := objc.Send[CDevice](objc.ID(getCDeviceClass().class), objc.Sel("deviceWithType:"), type_)
	return rv
}



// Creates a device that you can configure to use multiple compute devices.
//
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCDevice/init(type:selectsMultipleComputeDevices:)
func NewCDeviceWithTypeSelectsMultipleComputeDevices(type_ CDeviceType, selectsMultipleComputeDevices bool) CDevice {
	rv := objc.Send[CDevice](objc.ID(getCDeviceClass().class), objc.Sel("deviceWithType:selectsMultipleComputeDevices:"), type_, selectsMultipleComputeDevices)
	return rv
}


// Creates a device that uses the Apple Neural Engine, if one exists.
//
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCDevice/ane()
func (cc _CDeviceClass) AneDevice() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("aneDevice"))
	return rv
}

// Creates a device that uses the CPU.
//
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCDevice/cpu()
func (cc _CDeviceClass) CpuDevice() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("cpuDevice"))
	return rv
}

// Creates a device that uses a GPU, if one exists.
//
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCDevice/gpu()
func (cc _CDeviceClass) GpuDevice() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("gpuDevice"))
	return rv
}

// Creates a device using the GPUs you specify.
//
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCDevice/init(gpuDevices:)
func (cc _CDeviceClass) DeviceWithGPUDevices(gpus []objc.ID) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("deviceWithGPUDevices:"), gpus)
	return rv
}

// Creates a device of the type you specify.
//
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCDevice/init(type:)
func (cc _CDeviceClass) DeviceWithType(type_ CDeviceType) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("deviceWithType:"), type_)
	return rv
}

// Creates a device that you can configure to use multiple compute devices.
//
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCDevice/init(type:selectsMultipleComputeDevices:)
func (cc _CDeviceClass) DeviceWithTypeSelectsMultipleComputeDevices(type_ CDeviceType, selectsMultipleComputeDevices bool) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("deviceWithType:selectsMultipleComputeDevices:"), type_, selectsMultipleComputeDevices)
	return rv
}

// The active device.
//
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCDevice/actualDeviceType
func (c_ CDevice) ActualDeviceType() CDeviceType {
	rv := objc.Send[CDeviceType](c_.ID, objc.Sel("actualDeviceType"))
	return rv
}

// An array that contains the specific Metal devices you use to execute neural networks.
//
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCDevice/gpuDevices
func (c_ CDevice) GpuDevices() []objc.ID {
	rv := objc.Send[[]objc.ID](c_.ID, objc.Sel("gpuDevices"))
	return rv
}

// The type you specify when creating the device.
//
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCDevice/type
func (c_ CDevice) Type() CDeviceType {
	rv := objc.Send[CDeviceType](c_.ID, objc.Sel("type"))
	return rv
}


