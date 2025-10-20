// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [NDArrayLUTQuantizationDescriptor] class.
var (
	NDArrayLUTQuantizationDescriptorClass     _NDArrayLUTQuantizationDescriptorClass
	NDArrayLUTQuantizationDescriptorClassOnce sync.Once
)

func getNDArrayLUTQuantizationDescriptorClass() _NDArrayLUTQuantizationDescriptorClass {
	NDArrayLUTQuantizationDescriptorClassOnce.Do(func() {
		NDArrayLUTQuantizationDescriptorClass = _NDArrayLUTQuantizationDescriptorClass{objc.GetClass("MPSNDArrayLUTQuantizationDescriptor")}
	})
	return NDArrayLUTQuantizationDescriptorClass
}

type _NDArrayLUTQuantizationDescriptorClass struct {
	class objc.Class
}

// An interface definition for the [NDArrayLUTQuantizationDescriptor] class.
type INDArrayLUTQuantizationDescriptor interface {
	INDArrayQuantizationDescriptor
}

//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArrayLUTQuantizationDescriptor
type NDArrayLUTQuantizationDescriptor struct {
	NDArrayQuantizationDescriptor
}

// NDArrayLUTQuantizationDescriptorFrom constructs a [NDArrayLUTQuantizationDescriptor] from an unsafe.Pointer.
func NDArrayLUTQuantizationDescriptorFrom(ptr unsafe.Pointer) NDArrayLUTQuantizationDescriptor {
	return NDArrayLUTQuantizationDescriptor{
		NDArrayQuantizationDescriptor: NDArrayQuantizationDescriptorFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (nc _NDArrayLUTQuantizationDescriptorClass) Alloc() NDArrayLUTQuantizationDescriptor {
	rv := objc.Send[NDArrayLUTQuantizationDescriptor](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (nc _NDArrayLUTQuantizationDescriptorClass) New() NDArrayLUTQuantizationDescriptor {
	rv := objc.Send[NDArrayLUTQuantizationDescriptor](objc.ID(nc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (n_ NDArrayLUTQuantizationDescriptor) Init() NDArrayLUTQuantizationDescriptor {
	rv := objc.Send[NDArrayLUTQuantizationDescriptor](n_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n_ NDArrayLUTQuantizationDescriptor) Autorelease() NDArrayLUTQuantizationDescriptor {
	rv := objc.Send[NDArrayLUTQuantizationDescriptor](n_.ID, objc.Sel("autorelease"))
	return rv
}

// NewNDArrayLUTQuantizationDescriptor creates a new NDArrayLUTQuantizationDescriptor instance.
func NewNDArrayLUTQuantizationDescriptor() NDArrayLUTQuantizationDescriptor {
	return getNDArrayLUTQuantizationDescriptorClass().New()
}

//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArrayLUTQuantizationDescriptor/init(dataType:vectorAxis:)
func NewNDArrayLUTQuantizationDescriptorWithDataTypeVectorAxis(quantizationDataType unsafe.Pointer, vectorAxis uint) NDArrayLUTQuantizationDescriptor {
	instance := getNDArrayLUTQuantizationDescriptorClass().Alloc()
	rv := objc.Send[NDArrayLUTQuantizationDescriptor](instance.ID, objc.Sel("initWithDataType:vectorAxis:"), quantizationDataType, vectorAxis)
	rv.Autorelease()
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArrayLUTQuantizationDescriptor/init(dataType:)
func NewNDArrayLUTQuantizationDescriptorWithDataType(quantizationDataType unsafe.Pointer) NDArrayLUTQuantizationDescriptor {
	instance := getNDArrayLUTQuantizationDescriptorClass().Alloc()
	rv := objc.Send[NDArrayLUTQuantizationDescriptor](instance.ID, objc.Sel("initWithDataType:"), quantizationDataType)
	rv.Autorelease()
	return rv
}
