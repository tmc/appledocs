// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [CPUComputeDevice] class.
var (
	CPUComputeDeviceClass     _CPUComputeDeviceClass
	CPUComputeDeviceClassOnce sync.Once
)

func getCPUComputeDeviceClass() _CPUComputeDeviceClass {
	CPUComputeDeviceClassOnce.Do(func() {
		CPUComputeDeviceClass = _CPUComputeDeviceClass{objc.GetClass("MLCPUComputeDevice")}
	})
	return CPUComputeDeviceClass
}

type _CPUComputeDeviceClass struct {
	class objc.Class
}

// An interface definition for the [CPUComputeDevice] class.
type ICPUComputeDevice interface {
	objectivec.IObject
	// properties:
	// methods:
}

// An object that represents a CPU compute device.


// An object that represents a CPU compute device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLCPUComputeDevice
type CPUComputeDevice struct {
	objectivec.Object
}

// CPUComputeDeviceFrom constructs a [CPUComputeDevice] from an unsafe.Pointer.
//
// An object that represents a CPU compute device.
func CPUComputeDeviceFrom(ptr unsafe.Pointer) CPUComputeDevice {
	return CPUComputeDevice{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (cc _CPUComputeDeviceClass) Alloc() CPUComputeDevice {
	rv := objc.Send[CPUComputeDevice](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CPUComputeDeviceClass) New() CPUComputeDevice {
	rv := objc.Send[CPUComputeDevice](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CPUComputeDevice) Init() CPUComputeDevice {
	rv := objc.Send[CPUComputeDevice](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CPUComputeDevice) Autorelease() CPUComputeDevice {
	rv := objc.Send[CPUComputeDevice](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCPUComputeDevice creates a new CPUComputeDevice instance.
func NewCPUComputeDevice() CPUComputeDevice {
	return getCPUComputeDeviceClass().New()
}




