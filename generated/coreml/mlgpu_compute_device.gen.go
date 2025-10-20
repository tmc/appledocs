// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [GPUComputeDevice] class.
var (
	GPUComputeDeviceClass     _GPUComputeDeviceClass
	GPUComputeDeviceClassOnce sync.Once
)

func getGPUComputeDeviceClass() _GPUComputeDeviceClass {
	GPUComputeDeviceClassOnce.Do(func() {
		GPUComputeDeviceClass = _GPUComputeDeviceClass{objc.GetClass("MLGPUComputeDevice")}
	})
	return GPUComputeDeviceClass
}

type _GPUComputeDeviceClass struct {
	class objc.Class
}

// An interface definition for the [GPUComputeDevice] class.
type IGPUComputeDevice interface {
	objectivec.IObject
}

// An object that represents a GPU compute device.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLGPUComputeDevice
type GPUComputeDevice struct {
	objectivec.Object
}

// GPUComputeDeviceFrom constructs a [GPUComputeDevice] from an unsafe.Pointer.
//
// An object that represents a GPU compute device.
func GPUComputeDeviceFrom(ptr unsafe.Pointer) GPUComputeDevice {
	return GPUComputeDevice{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (gc _GPUComputeDeviceClass) Alloc() GPUComputeDevice {
	rv := objc.Send[GPUComputeDevice](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (gc _GPUComputeDeviceClass) New() GPUComputeDevice {
	rv := objc.Send[GPUComputeDevice](objc.ID(gc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (g_ GPUComputeDevice) Init() GPUComputeDevice {
	rv := objc.Send[GPUComputeDevice](g_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g_ GPUComputeDevice) Autorelease() GPUComputeDevice {
	rv := objc.Send[GPUComputeDevice](g_.ID, objc.Sel("autorelease"))
	return rv
}

// NewGPUComputeDevice creates a new GPUComputeDevice instance.
func NewGPUComputeDevice() GPUComputeDevice {
	return getGPUComputeDeviceClass().New()
}


// The device that represents the underlying metal device.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLGPUComputeDevice/metalDevice
func (g_ GPUComputeDevice) MetalDevice() objc.ID {
	rv := objc.Send[objc.ID](g_.ID, objc.Sel("metalDevice"))
	return rv
}



