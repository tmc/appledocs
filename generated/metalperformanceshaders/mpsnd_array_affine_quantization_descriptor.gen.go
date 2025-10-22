// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [NDArrayAffineQuantizationDescriptor] class.
var (
	NDArrayAffineQuantizationDescriptorClass     _NDArrayAffineQuantizationDescriptorClass
	NDArrayAffineQuantizationDescriptorClassOnce sync.Once
)

func getNDArrayAffineQuantizationDescriptorClass() _NDArrayAffineQuantizationDescriptorClass {
	NDArrayAffineQuantizationDescriptorClassOnce.Do(func() {
		NDArrayAffineQuantizationDescriptorClass = _NDArrayAffineQuantizationDescriptorClass{objc.GetClass("MPSNDArrayAffineQuantizationDescriptor")}
	})
	return NDArrayAffineQuantizationDescriptorClass
}

type _NDArrayAffineQuantizationDescriptorClass struct {
	class objc.Class
}

// An interface definition for the [NDArrayAffineQuantizationDescriptor] class.
type INDArrayAffineQuantizationDescriptor interface {
	INDArrayQuantizationDescriptor
	HasMinValue() bool
	SetHasMinValue(value bool)
	HasZeroPoint() bool
	SetHasZeroPoint(value bool)
	ImplicitZeroPoint() bool
	SetImplicitZeroPoint(value bool)
}

//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArrayAffineQuantizationDescriptor
type NDArrayAffineQuantizationDescriptor struct {
	NDArrayQuantizationDescriptor
}

// NDArrayAffineQuantizationDescriptorFrom constructs a [NDArrayAffineQuantizationDescriptor] from an unsafe.Pointer.
func NDArrayAffineQuantizationDescriptorFrom(ptr unsafe.Pointer) NDArrayAffineQuantizationDescriptor {
	return NDArrayAffineQuantizationDescriptor{
		NDArrayQuantizationDescriptor: NDArrayQuantizationDescriptorFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (nc _NDArrayAffineQuantizationDescriptorClass) Alloc() NDArrayAffineQuantizationDescriptor {
	rv := objc.Send[NDArrayAffineQuantizationDescriptor](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (nc _NDArrayAffineQuantizationDescriptorClass) New() NDArrayAffineQuantizationDescriptor {
	rv := objc.Send[NDArrayAffineQuantizationDescriptor](objc.ID(nc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (n_ NDArrayAffineQuantizationDescriptor) Init() NDArrayAffineQuantizationDescriptor {
	rv := objc.Send[NDArrayAffineQuantizationDescriptor](n_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n_ NDArrayAffineQuantizationDescriptor) Autorelease() NDArrayAffineQuantizationDescriptor {
	rv := objc.Send[NDArrayAffineQuantizationDescriptor](n_.ID, objc.Sel("autorelease"))
	return rv
}

// NewNDArrayAffineQuantizationDescriptor creates a new NDArrayAffineQuantizationDescriptor instance.
func NewNDArrayAffineQuantizationDescriptor() NDArrayAffineQuantizationDescriptor {
	return getNDArrayAffineQuantizationDescriptorClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArrayAffineQuantizationDescriptor/init(dataType:hasZeroPoint:hasMinValue:)
func NewNDArrayAffineQuantizationDescriptorWithDataTypeHasZeroPointHasMinValue(quantizationDataType unsafe.Pointer, hasZeroPoint bool, hasMinValue bool) NDArrayAffineQuantizationDescriptor {
	instance := getNDArrayAffineQuantizationDescriptorClass().Alloc()
	rv := objc.Send[NDArrayAffineQuantizationDescriptor](instance.ID, objc.Sel("initWithDataType:hasZeroPoint:hasMinValue:"), quantizationDataType, hasZeroPoint, hasMinValue)
	rv.Autorelease()
	return rv
}


//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArrayAffineQuantizationDescriptor/hasMinValue
func (n_ NDArrayAffineQuantizationDescriptor) HasMinValue() bool {
	rv := objc.Send[bool](n_.ID, objc.Sel("hasMinValue"))
	return rv
}


// SetHasMinValue sets the value of the hasMinValue property.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArrayAffineQuantizationDescriptor/hasMinValue
func (n_ NDArrayAffineQuantizationDescriptor) SetHasMinValue(value bool) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setHasMinValue:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArrayAffineQuantizationDescriptor/hasZeroPoint
func (n_ NDArrayAffineQuantizationDescriptor) HasZeroPoint() bool {
	rv := objc.Send[bool](n_.ID, objc.Sel("hasZeroPoint"))
	return rv
}


// SetHasZeroPoint sets the value of the hasZeroPoint property.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArrayAffineQuantizationDescriptor/hasZeroPoint
func (n_ NDArrayAffineQuantizationDescriptor) SetHasZeroPoint(value bool) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setHasZeroPoint:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArrayAffineQuantizationDescriptor/implicitZeroPoint
func (n_ NDArrayAffineQuantizationDescriptor) ImplicitZeroPoint() bool {
	rv := objc.Send[bool](n_.ID, objc.Sel("implicitZeroPoint"))
	return rv
}


// SetImplicitZeroPoint sets the value of the implicitZeroPoint property.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArrayAffineQuantizationDescriptor/implicitZeroPoint
func (n_ NDArrayAffineQuantizationDescriptor) SetImplicitZeroPoint(value bool) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setImplicitZeroPoint:"), value)
}


