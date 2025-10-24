// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [NDArrayQuantizedMatrixMultiplication] class.
var (
	NDArrayQuantizedMatrixMultiplicationClass     _NDArrayQuantizedMatrixMultiplicationClass
	NDArrayQuantizedMatrixMultiplicationClassOnce sync.Once
)

func getNDArrayQuantizedMatrixMultiplicationClass() _NDArrayQuantizedMatrixMultiplicationClass {
	NDArrayQuantizedMatrixMultiplicationClassOnce.Do(func() {
		NDArrayQuantizedMatrixMultiplicationClass = _NDArrayQuantizedMatrixMultiplicationClass{objc.GetClass("MPSNDArrayQuantizedMatrixMultiplication")}
	})
	return NDArrayQuantizedMatrixMultiplicationClass
}

type _NDArrayQuantizedMatrixMultiplicationClass struct {
	class objc.Class
}

// An interface definition for the [NDArrayQuantizedMatrixMultiplication] class.
type INDArrayQuantizedMatrixMultiplication interface {
	INDArrayMatrixMultiplication
	// properties:
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArrayQuantizedMatrixMultiplication
type NDArrayQuantizedMatrixMultiplication struct {
	NDArrayMatrixMultiplication
}

// NDArrayQuantizedMatrixMultiplicationFrom constructs a [NDArrayQuantizedMatrixMultiplication] from an unsafe.Pointer.
func NDArrayQuantizedMatrixMultiplicationFrom(ptr unsafe.Pointer) NDArrayQuantizedMatrixMultiplication {
	return NDArrayQuantizedMatrixMultiplication{
		NDArrayMatrixMultiplication: NDArrayMatrixMultiplicationFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (nc _NDArrayQuantizedMatrixMultiplicationClass) Alloc() NDArrayQuantizedMatrixMultiplication {
	rv := objc.Send[NDArrayQuantizedMatrixMultiplication](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (nc _NDArrayQuantizedMatrixMultiplicationClass) New() NDArrayQuantizedMatrixMultiplication {
	rv := objc.Send[NDArrayQuantizedMatrixMultiplication](objc.ID(nc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (n_ NDArrayQuantizedMatrixMultiplication) Init() NDArrayQuantizedMatrixMultiplication {
	rv := objc.Send[NDArrayQuantizedMatrixMultiplication](n_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n_ NDArrayQuantizedMatrixMultiplication) Autorelease() NDArrayQuantizedMatrixMultiplication {
	rv := objc.Send[NDArrayQuantizedMatrixMultiplication](n_.ID, objc.Sel("autorelease"))
	return rv
}

// NewNDArrayQuantizedMatrixMultiplication creates a new NDArrayQuantizedMatrixMultiplication instance.
func NewNDArrayQuantizedMatrixMultiplication() NDArrayQuantizedMatrixMultiplication {
	return getNDArrayQuantizedMatrixMultiplicationClass().New()
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArrayQuantizedMatrixMultiplication/init(device:leftQuantizationDescriptor:rightQuantizationDescriptor:)
func NewNDArrayQuantizedMatrixMultiplicationWithDeviceLeftQuantizationDescriptorRightQuantizationDescriptor(device objectivec.IObject, leftQuantizationDescriptor IMPSNDArrayQuantizationDescriptor, rightQuantizationDescriptor IMPSNDArrayQuantizationDescriptor) NDArrayQuantizedMatrixMultiplication {
	instance := getNDArrayQuantizedMatrixMultiplicationClass().Alloc()
	rv := objc.Send[NDArrayQuantizedMatrixMultiplication](instance.ID, objc.Sel("initWithDevice:leftQuantizationDescriptor:rightQuantizationDescriptor:"), device, leftQuantizationDescriptor, rightQuantizationDescriptor)
	rv.Autorelease()
	return rv
}



