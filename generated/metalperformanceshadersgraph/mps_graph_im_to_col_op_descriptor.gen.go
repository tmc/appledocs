// Code generated from Apple documentation for MetalPerformanceShadersGraph. DO NOT EDIT.

package metalperformanceshadersgraph

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [GraphImToColOpDescriptor] class.
var (
	GraphImToColOpDescriptorClass     _GraphImToColOpDescriptorClass
	GraphImToColOpDescriptorClassOnce sync.Once
)

func getGraphImToColOpDescriptorClass() _GraphImToColOpDescriptorClass {
	GraphImToColOpDescriptorClassOnce.Do(func() {
		GraphImToColOpDescriptorClass = _GraphImToColOpDescriptorClass{objc.GetClass("MPSGraphImToColOpDescriptor")}
	})
	return GraphImToColOpDescriptorClass
}

type _GraphImToColOpDescriptorClass struct {
	class objc.Class
}

// An interface definition for the [GraphImToColOpDescriptor] class.
type IGraphImToColOpDescriptor interface {
	IGraphObject
}

// The class that defines the parameters for an image to column or column to image operation.
//
// Use this descriptor with the following methods:
//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphImToColOpDescriptor
type GraphImToColOpDescriptor struct {
	GraphObject
}

// GraphImToColOpDescriptorFrom constructs a [GraphImToColOpDescriptor] from an unsafe.Pointer.
//
// The class that defines the parameters for an image to column or column to image operation.
func GraphImToColOpDescriptorFrom(ptr unsafe.Pointer) GraphImToColOpDescriptor {
	return GraphImToColOpDescriptor{
		GraphObject: GraphObjectFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (gc _GraphImToColOpDescriptorClass) Alloc() GraphImToColOpDescriptor {
	rv := objc.Send[GraphImToColOpDescriptor](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (gc _GraphImToColOpDescriptorClass) New() GraphImToColOpDescriptor {
	rv := objc.Send[GraphImToColOpDescriptor](objc.ID(gc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (g_ GraphImToColOpDescriptor) Init() GraphImToColOpDescriptor {
	rv := objc.Send[GraphImToColOpDescriptor](g_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g_ GraphImToColOpDescriptor) Autorelease() GraphImToColOpDescriptor {
	rv := objc.Send[GraphImToColOpDescriptor](g_.ID, objc.Sel("autorelease"))
	return rv
}

// NewGraphImToColOpDescriptor creates a new GraphImToColOpDescriptor instance.
func NewGraphImToColOpDescriptor() GraphImToColOpDescriptor {
	return getGraphImToColOpDescriptorClass().New()
}


// The property that defines the dilation in width dimension.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphImToColOpDescriptor/dilationRateInX
func (g_ GraphImToColOpDescriptor) DilationRateInX() uint {
	rv := objc.Send[uint](g_.ID, objc.Sel("dilationRateInX"))
	return rv
}


// SetDilationRateInX sets the value of the dilationRateInX property.
// The property that defines the dilation in width dimension.

//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphImToColOpDescriptor/dilationRateInX
func (g_ GraphImToColOpDescriptor) SetDilationRateInX(value uint) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setDilationRateInX:"), value)
}
// The property that defines the stride in height dimension.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphImToColOpDescriptor/strideInY
func (g_ GraphImToColOpDescriptor) StrideInY() uint {
	rv := objc.Send[uint](g_.ID, objc.Sel("strideInY"))
	return rv
}


// SetStrideInY sets the value of the strideInY property.
// The property that defines the stride in height dimension.

//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphImToColOpDescriptor/strideInY
func (g_ GraphImToColOpDescriptor) SetStrideInY(value uint) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setStrideInY:"), value)
}


