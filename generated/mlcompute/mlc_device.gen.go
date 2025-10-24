// Code generated from Apple documentation for MLCompute. DO NOT EDIT.

package mlcompute

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MLCDevice */


/* debug [class_header]: Header for MLCDevice */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CDevice */
// An interface definition for the [CDevice] class.
type ICDevice interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for CDevice */
	// properties:
	ActualDeviceType() CDeviceType
	GpuDevices() []objc.ID
	Type() CDeviceType
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CDevice */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CDevice */
// Alloc allocates a new instance without initialization.
func (cc _CDeviceClass) Alloc() CDevice {
	rv := objc.Send[CDevice](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CDevice */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CDevice */

// Creates a device using the GPUs you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCDevice/init(gpuDevices:)
func NewCDeviceWithGPUDevices(gpus []objc.ID) CDevice {
	rv := objc.Send[CDevice](objc.ID(getCDeviceClass().class), objc.Sel("deviceWithGPUDevices:"), gpus)
	return rv
}/* debug [class_init_methods/constructor]: NewCDeviceWithGPUDevices */


// Creates a device of the type you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCDevice/init(type:)
func NewCDeviceWithType(type_ CDeviceType) CDevice {
	rv := objc.Send[CDevice](objc.ID(getCDeviceClass().class), objc.Sel("deviceWithType:"), type_)
	return rv
}/* debug [class_init_methods/constructor]: NewCDeviceWithType */


// Creates a device that you can configure to use multiple compute devices.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCDevice/init(type:selectsMultipleComputeDevices:)
func NewCDeviceWithTypeSelectsMultipleComputeDevices(type_ CDeviceType, selectsMultipleComputeDevices bool) CDevice {
	rv := objc.Send[CDevice](objc.ID(getCDeviceClass().class), objc.Sel("deviceWithType:selectsMultipleComputeDevices:"), type_, selectsMultipleComputeDevices)
	return rv
}/* debug [class_init_methods/constructor]: NewCDeviceWithTypeSelectsMultipleComputeDevices */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CDevice */

// Creates a device that uses the Apple Neural Engine, if one exists.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCDevice/ane()
func (cc _CDeviceClass) AneDevice() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("aneDevice"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=AneDevice) */


// Creates a device that uses the CPU.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCDevice/cpu()
func (cc _CDeviceClass) CpuDevice() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("cpuDevice"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=CpuDevice) */


// Creates a device that uses a GPU, if one exists.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCDevice/gpu()
func (cc _CDeviceClass) GpuDevice() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("gpuDevice"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=GpuDevice) */


// Creates a device using the GPUs you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCDevice/init(gpuDevices:)
func (cc _CDeviceClass) DeviceWithGPUDevices(gpus []objc.ID) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("deviceWithGPUDevices:"), gpus)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=DeviceWithGPUDevices) */


// Creates a device of the type you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCDevice/init(type:)
func (cc _CDeviceClass) DeviceWithType(type_ CDeviceType) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("deviceWithType:"), type_)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=DeviceWithType) */


// Creates a device that you can configure to use multiple compute devices.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCDevice/init(type:selectsMultipleComputeDevices:)
func (cc _CDeviceClass) DeviceWithTypeSelectsMultipleComputeDevices(type_ CDeviceType, selectsMultipleComputeDevices bool) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("deviceWithType:selectsMultipleComputeDevices:"), type_, selectsMultipleComputeDevices)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=DeviceWithTypeSelectsMultipleComputeDevices) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CDevice */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CDevice */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CDevice */

// The active device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCDevice/actualDeviceType
func (c_ CDevice) ActualDeviceType() CDeviceType {
	rv := objc.Send[CDeviceType](c_.ID, objc.Sel("actualDeviceType"))
	return rv
}/* debug [instance_properties/getter]: actualDeviceType */


// An array that contains the specific Metal devices you use to execute neural networks.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCDevice/gpuDevices
func (c_ CDevice) GpuDevices() []objc.ID {
	rv := objc.Send[[]objc.ID](c_.ID, objc.Sel("gpuDevices"))
	return rv
}/* debug [instance_properties/getter]: gpuDevices */


// The type you specify when creating the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCDevice/type
func (c_ CDevice) Type() CDeviceType {
	rv := objc.Send[CDeviceType](c_.ID, objc.Sel("type"))
	return rv
}/* debug [instance_properties/getter]: type */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MLCDevice */


