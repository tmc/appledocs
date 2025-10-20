// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [NDArrayDescriptor] class.
var (
	NDArrayDescriptorClass     _NDArrayDescriptorClass
	NDArrayDescriptorClassOnce sync.Once
)

func getNDArrayDescriptorClass() _NDArrayDescriptorClass {
	NDArrayDescriptorClassOnce.Do(func() {
		NDArrayDescriptorClass = _NDArrayDescriptorClass{objc.GetClass("MPSNDArrayDescriptor")}
	})
	return NDArrayDescriptorClass
}

type _NDArrayDescriptorClass struct {
	class objc.Class
}

// An interface definition for the [NDArrayDescriptor] class.
type INDArrayDescriptor interface {
	objectivec.IObject
	DimensionOrder() unsafe.Pointer
	SliceRangeForDimension(dimensionIndex uint) unsafe.Pointer
}

//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArrayDescriptor
type NDArrayDescriptor struct {
	objectivec.Object
}

// NDArrayDescriptorFrom constructs a [NDArrayDescriptor] from an unsafe.Pointer.
func NDArrayDescriptorFrom(ptr unsafe.Pointer) NDArrayDescriptor {
	return NDArrayDescriptor{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (nc _NDArrayDescriptorClass) Alloc() NDArrayDescriptor {
	rv := objc.Send[NDArrayDescriptor](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (nc _NDArrayDescriptorClass) New() NDArrayDescriptor {
	rv := objc.Send[NDArrayDescriptor](objc.ID(nc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (n_ NDArrayDescriptor) Init() NDArrayDescriptor {
	rv := objc.Send[NDArrayDescriptor](n_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n_ NDArrayDescriptor) Autorelease() NDArrayDescriptor {
	rv := objc.Send[NDArrayDescriptor](n_.ID, objc.Sel("autorelease"))
	return rv
}

// NewNDArrayDescriptor creates a new NDArrayDescriptor instance.
func NewNDArrayDescriptor() NDArrayDescriptor {
	return getNDArrayDescriptorClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArrayDescriptor/dimensionOrder()
func (n_ NDArrayDescriptor) DimensionOrder() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](n_.ID, objc.Sel("dimensionOrder"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArrayDescriptor/sliceRange(forDimension:)
func (n_ NDArrayDescriptor) SliceRangeForDimension(dimensionIndex uint) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](n_.ID, objc.Sel("sliceRangeForDimension:"), dimensionIndex)
	return rv
}



