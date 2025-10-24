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
	// properties:
	ActualDeviceType() CDeviceType /* not a class type */
	SetActualDeviceType(value CDeviceType /* not a class type */)
	GpuDevices() Device /* not a class type */
	SetGpuDevices(value Device /* not a class type */)
	Type() CDeviceType /* not a class type */
	SetType(value CDeviceType /* not a class type */)
	// methods:
}

// An object that represents the CPU or one or more GPUs the framework uses to execute a neural network.


// An object that represents the CPU or one or more GPUs the framework uses to execute a neural network.
//
// [Full Topic]
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



// The active device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlcdevice/actualdevicetype
func (c_ CDevice) ActualDeviceType() CDeviceType /* not a class type */ {
	rv := objc.Send[CDeviceType](c_.ID, objc.Sel("actualDeviceType"))
	return rv
}


// The active device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlcdevice/actualdevicetype
func (c_ CDevice) SetActualDeviceType(value CDeviceType /* not a class type */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setActualDeviceType:"), value)
}


// An array that contains the specific Metal devices you use to execute neural networks.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlcdevice/gpudevices
func (c_ CDevice) GpuDevices() Device /* not a class type */ {
	rv := objc.Send[Device](c_.ID, objc.Sel("gpuDevices"))
	return rv
}


// An array that contains the specific Metal devices you use to execute neural networks.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlcdevice/gpudevices
func (c_ CDevice) SetGpuDevices(value Device /* not a class type */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setGpuDevices:"), value)
}


// The type you specify when creating the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlcdevice/type
func (c_ CDevice) Type() CDeviceType /* not a class type */ {
	rv := objc.Send[CDeviceType](c_.ID, objc.Sel("type"))
	return rv
}


// The type you specify when creating the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlcdevice/type
func (c_ CDevice) SetType(value CDeviceType /* not a class type */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setType:"), value)
}



