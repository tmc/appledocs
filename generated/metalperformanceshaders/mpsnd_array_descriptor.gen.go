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

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarraydescriptor/numberofdimensions
func (n_ NDArrayDescriptor) NumberOfDimensions() int {
	rv := objc.Send[int](n_.ID, objc.Sel("numberOfDimensions"))
	return rv
}


// SetNumberOfDimensions sets the value of the numberOfDimensions property.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarraydescriptor/numberofdimensions
func (n_ NDArrayDescriptor) SetNumberOfDimensions(value int) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setNumberOfDimensions:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarraydescriptor/datatype
func (n_ NDArrayDescriptor) DataType() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](n_.ID, objc.Sel("dataType"))
	return rv
}


// SetDataType sets the value of the dataType property.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarraydescriptor/datatype
func (n_ NDArrayDescriptor) SetDataType(value unsafe.Pointer) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setDataType:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarraydescriptor/preferpackedrows
func (n_ NDArrayDescriptor) PreferPackedRows() bool {
	rv := objc.Send[bool](n_.ID, objc.Sel("preferPackedRows"))
	return rv
}


// SetPreferPackedRows sets the value of the preferPackedRows property.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarraydescriptor/preferpackedrows
func (n_ NDArrayDescriptor) SetPreferPackedRows(value bool) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setPreferPackedRows:"), value)
}



