// Code generated from Apple documentation for MetalPerformanceShadersGraph. DO NOT EDIT.

package metalperformanceshadersgraph

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [GraphDepthwiseConvolution2DOpDescriptor] class.
var (
	GraphDepthwiseConvolution2DOpDescriptorClass     _GraphDepthwiseConvolution2DOpDescriptorClass
	GraphDepthwiseConvolution2DOpDescriptorClassOnce sync.Once
)

func getGraphDepthwiseConvolution2DOpDescriptorClass() _GraphDepthwiseConvolution2DOpDescriptorClass {
	GraphDepthwiseConvolution2DOpDescriptorClassOnce.Do(func() {
		GraphDepthwiseConvolution2DOpDescriptorClass = _GraphDepthwiseConvolution2DOpDescriptorClass{objc.GetClass("MPSGraphDepthwiseConvolution2DOpDescriptor")}
	})
	return GraphDepthwiseConvolution2DOpDescriptorClass
}

type _GraphDepthwiseConvolution2DOpDescriptorClass struct {
	class objc.Class
}

// An interface definition for the [GraphDepthwiseConvolution2DOpDescriptor] class.
type IGraphDepthwiseConvolution2DOpDescriptor interface {
	IGraphObject
}

// A class that defines the parameters for a 2D-depthwise convolution operation.
//
// An defines constant parameters for 2D-depthwise convolutions. Use this class with , , and methods.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphDepthwiseConvolution2DOpDescriptor
type GraphDepthwiseConvolution2DOpDescriptor struct {
	GraphObject
}

// GraphDepthwiseConvolution2DOpDescriptorFrom constructs a [GraphDepthwiseConvolution2DOpDescriptor] from an unsafe.Pointer.
//
// A class that defines the parameters for a 2D-depthwise convolution operation.
func GraphDepthwiseConvolution2DOpDescriptorFrom(ptr unsafe.Pointer) GraphDepthwiseConvolution2DOpDescriptor {
	return GraphDepthwiseConvolution2DOpDescriptor{
		GraphObject: GraphObjectFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (gc _GraphDepthwiseConvolution2DOpDescriptorClass) Alloc() GraphDepthwiseConvolution2DOpDescriptor {
	rv := objc.Send[GraphDepthwiseConvolution2DOpDescriptor](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (gc _GraphDepthwiseConvolution2DOpDescriptorClass) New() GraphDepthwiseConvolution2DOpDescriptor {
	rv := objc.Send[GraphDepthwiseConvolution2DOpDescriptor](objc.ID(gc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (g_ GraphDepthwiseConvolution2DOpDescriptor) Init() GraphDepthwiseConvolution2DOpDescriptor {
	rv := objc.Send[GraphDepthwiseConvolution2DOpDescriptor](g_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g_ GraphDepthwiseConvolution2DOpDescriptor) Autorelease() GraphDepthwiseConvolution2DOpDescriptor {
	rv := objc.Send[GraphDepthwiseConvolution2DOpDescriptor](g_.ID, objc.Sel("autorelease"))
	return rv
}

// NewGraphDepthwiseConvolution2DOpDescriptor creates a new GraphDepthwiseConvolution2DOpDescriptor instance.
func NewGraphDepthwiseConvolution2DOpDescriptor() GraphDepthwiseConvolution2DOpDescriptor {
	return getGraphDepthwiseConvolution2DOpDescriptorClass().New()
}


// The dilation rate for the x dimension.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphDepthwiseConvolution2DOpDescriptor/dilationRateInX
func (g_ GraphDepthwiseConvolution2DOpDescriptor) DilationRateInX() uint {
	rv := objc.Send[uint](g_.ID, objc.Sel("dilationRateInX"))
	return rv
}


// SetDilationRateInX sets the value of the dilationRateInX property.
// The dilation rate for the x dimension.

//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphDepthwiseConvolution2DOpDescriptor/dilationRateInX
func (g_ GraphDepthwiseConvolution2DOpDescriptor) SetDilationRateInX(value uint) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setDilationRateInX:"), value)
}

// The explicit padding value for the x dimension the operation adds before the data.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphDepthwiseConvolution2DOpDescriptor/paddingLeft
func (g_ GraphDepthwiseConvolution2DOpDescriptor) PaddingLeft() uint {
	rv := objc.Send[uint](g_.ID, objc.Sel("paddingLeft"))
	return rv
}


// SetPaddingLeft sets the value of the paddingLeft property.
// The explicit padding value for the x dimension the operation adds before the data.

//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphDepthwiseConvolution2DOpDescriptor/paddingLeft
func (g_ GraphDepthwiseConvolution2DOpDescriptor) SetPaddingLeft(value uint) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setPaddingLeft:"), value)
}

// The explicit padding value for the x dimension operation adds after the data.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphDepthwiseConvolution2DOpDescriptor/paddingRight
func (g_ GraphDepthwiseConvolution2DOpDescriptor) PaddingRight() uint {
	rv := objc.Send[uint](g_.ID, objc.Sel("paddingRight"))
	return rv
}


// SetPaddingRight sets the value of the paddingRight property.
// The explicit padding value for the x dimension operation adds after the data.

//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphDepthwiseConvolution2DOpDescriptor/paddingRight
func (g_ GraphDepthwiseConvolution2DOpDescriptor) SetPaddingRight(value uint) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setPaddingRight:"), value)
}



