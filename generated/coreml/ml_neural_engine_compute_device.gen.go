// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [NeuralEngineComputeDevice] class.
var (
	NeuralEngineComputeDeviceClass     _NeuralEngineComputeDeviceClass
	NeuralEngineComputeDeviceClassOnce sync.Once
)

func getNeuralEngineComputeDeviceClass() _NeuralEngineComputeDeviceClass {
	NeuralEngineComputeDeviceClassOnce.Do(func() {
		NeuralEngineComputeDeviceClass = _NeuralEngineComputeDeviceClass{objc.GetClass("MLNeuralEngineComputeDevice")}
	})
	return NeuralEngineComputeDeviceClass
}

type _NeuralEngineComputeDeviceClass struct {
	class objc.Class
}

// An interface definition for the [NeuralEngineComputeDevice] class.
type INeuralEngineComputeDevice interface {
	objectivec.IObject
	// properties:
	TotalCoreCount() int
	// methods:
}

// An object that represents a Neural Engine compute device.


// An object that represents a Neural Engine compute device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLNeuralEngineComputeDevice
type NeuralEngineComputeDevice struct {
	objectivec.Object
}

// NeuralEngineComputeDeviceFrom constructs a [NeuralEngineComputeDevice] from an unsafe.Pointer.
//
// An object that represents a Neural Engine compute device.
func NeuralEngineComputeDeviceFrom(ptr unsafe.Pointer) NeuralEngineComputeDevice {
	return NeuralEngineComputeDevice{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (nc _NeuralEngineComputeDeviceClass) Alloc() NeuralEngineComputeDevice {
	rv := objc.Send[NeuralEngineComputeDevice](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (nc _NeuralEngineComputeDeviceClass) New() NeuralEngineComputeDevice {
	rv := objc.Send[NeuralEngineComputeDevice](objc.ID(nc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (n_ NeuralEngineComputeDevice) Init() NeuralEngineComputeDevice {
	rv := objc.Send[NeuralEngineComputeDevice](n_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n_ NeuralEngineComputeDevice) Autorelease() NeuralEngineComputeDevice {
	rv := objc.Send[NeuralEngineComputeDevice](n_.ID, objc.Sel("autorelease"))
	return rv
}

// NewNeuralEngineComputeDevice creates a new NeuralEngineComputeDevice instance.
func NewNeuralEngineComputeDevice() NeuralEngineComputeDevice {
	return getNeuralEngineComputeDeviceClass().New()
}



// The total number of cores in the Neural Engine.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLNeuralEngineComputeDevice/totalCoreCount
func (n_ NeuralEngineComputeDevice) TotalCoreCount() int {
	rv := objc.Send[int](n_.ID, objc.Sel("totalCoreCount"))
	return rv
}



