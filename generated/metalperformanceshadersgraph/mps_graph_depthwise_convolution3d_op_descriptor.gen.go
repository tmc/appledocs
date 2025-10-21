// Code generated from Apple documentation for MetalPerformanceShadersGraph. DO NOT EDIT.

package metalperformanceshadersgraph

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [GraphDepthwiseConvolution3DOpDescriptor] class.
var (
	GraphDepthwiseConvolution3DOpDescriptorClass     _GraphDepthwiseConvolution3DOpDescriptorClass
	GraphDepthwiseConvolution3DOpDescriptorClassOnce sync.Once
)

func getGraphDepthwiseConvolution3DOpDescriptorClass() _GraphDepthwiseConvolution3DOpDescriptorClass {
	GraphDepthwiseConvolution3DOpDescriptorClassOnce.Do(func() {
		GraphDepthwiseConvolution3DOpDescriptorClass = _GraphDepthwiseConvolution3DOpDescriptorClass{objc.GetClass("MPSGraphDepthwiseConvolution3DOpDescriptor")}
	})
	return GraphDepthwiseConvolution3DOpDescriptorClass
}

type _GraphDepthwiseConvolution3DOpDescriptorClass struct {
	class objc.Class
}

// An interface definition for the [GraphDepthwiseConvolution3DOpDescriptor] class.
type IGraphDepthwiseConvolution3DOpDescriptor interface {
	IGraphObject
}

// The class that defines the parameters for a 3D-depthwise convolution operation.
//
// A defines constant parameters for 3D depthwise convolutions. Use this class with , and methods.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphDepthwiseConvolution3DOpDescriptor
type GraphDepthwiseConvolution3DOpDescriptor struct {
	GraphObject
}

// GraphDepthwiseConvolution3DOpDescriptorFrom constructs a [GraphDepthwiseConvolution3DOpDescriptor] from an unsafe.Pointer.
//
// The class that defines the parameters for a 3D-depthwise convolution operation.
func GraphDepthwiseConvolution3DOpDescriptorFrom(ptr unsafe.Pointer) GraphDepthwiseConvolution3DOpDescriptor {
	return GraphDepthwiseConvolution3DOpDescriptor{
		GraphObject: GraphObjectFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (gc _GraphDepthwiseConvolution3DOpDescriptorClass) Alloc() GraphDepthwiseConvolution3DOpDescriptor {
	rv := objc.Send[GraphDepthwiseConvolution3DOpDescriptor](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (gc _GraphDepthwiseConvolution3DOpDescriptorClass) New() GraphDepthwiseConvolution3DOpDescriptor {
	rv := objc.Send[GraphDepthwiseConvolution3DOpDescriptor](objc.ID(gc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (g_ GraphDepthwiseConvolution3DOpDescriptor) Init() GraphDepthwiseConvolution3DOpDescriptor {
	rv := objc.Send[GraphDepthwiseConvolution3DOpDescriptor](g_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g_ GraphDepthwiseConvolution3DOpDescriptor) Autorelease() GraphDepthwiseConvolution3DOpDescriptor {
	rv := objc.Send[GraphDepthwiseConvolution3DOpDescriptor](g_.ID, objc.Sel("autorelease"))
	return rv
}

// NewGraphDepthwiseConvolution3DOpDescriptor creates a new GraphDepthwiseConvolution3DOpDescriptor instance.
func NewGraphDepthwiseConvolution3DOpDescriptor() GraphDepthwiseConvolution3DOpDescriptor {
	return getGraphDepthwiseConvolution3DOpDescriptorClass().New()
}




// Creates a 3D depthwise convolution descriptor with default values.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphDepthwiseConvolution3DOpDescriptor/init(paddingStyle:)
func NewGraphDepthwiseConvolution3DOpDescriptorWithPaddingStyle(paddingStyle unsafe.Pointer) GraphDepthwiseConvolution3DOpDescriptor {
	rv := objc.Send[GraphDepthwiseConvolution3DOpDescriptor](objc.ID(getGraphDepthwiseConvolution3DOpDescriptorClass().class), objc.Sel("descriptorWithPaddingStyle:"), paddingStyle)
	return rv
}


// Creates a 3D depthwise convolution descriptor with default values.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphDepthwiseConvolution3DOpDescriptor/init(paddingStyle:)
func (gc _GraphDepthwiseConvolution3DOpDescriptorClass) DescriptorWithPaddingStyle(paddingStyle unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(gc.class), objc.Sel("descriptorWithPaddingStyle:"), paddingStyle)
	return rv
}

// The axis that contains the channels in the input and the weights, within
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphdepthwiseconvolution3dopdescriptor/channeldimensionindex
func (g_ GraphDepthwiseConvolution3DOpDescriptor) ChannelDimensionIndex() int {
	rv := objc.Send[int](g_.ID, objc.Sel("channelDimensionIndex"))
	return rv
}


// SetChannelDimensionIndex sets the value of the channelDimensionIndex property.
// The axis that contains the channels in the input and the weights, within

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphdepthwiseconvolution3dopdescriptor/channeldimensionindex
func (g_ GraphDepthwiseConvolution3DOpDescriptor) SetChannelDimensionIndex(value int) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setChannelDimensionIndex:"), value)
}

// The dilation rates for spatial dimensions.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphdepthwiseconvolution3dopdescriptor/dilationrates
func (g_ GraphDepthwiseConvolution3DOpDescriptor) DilationRates() foundation.Number {
	rv := objc.Send[foundation.Number](g_.ID, objc.Sel("dilationRates"))
	return rv
}


// SetDilationRates sets the value of the dilationRates property.
// The dilation rates for spatial dimensions.

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphdepthwiseconvolution3dopdescriptor/dilationrates
func (g_ GraphDepthwiseConvolution3DOpDescriptor) SetDilationRates(value foundation.Number) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setDilationRates:"), value)
}

// The padding style for the operation.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphdepthwiseconvolution3dopdescriptor/paddingstyle
func (g_ GraphDepthwiseConvolution3DOpDescriptor) PaddingStyle() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("paddingStyle"))
	return rv
}


// SetPaddingStyle sets the value of the paddingStyle property.
// The padding style for the operation.

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphdepthwiseconvolution3dopdescriptor/paddingstyle
func (g_ GraphDepthwiseConvolution3DOpDescriptor) SetPaddingStyle(value unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setPaddingStyle:"), value)
}

// The padding values for spatial dimensions.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphdepthwiseconvolution3dopdescriptor/paddingvalues
func (g_ GraphDepthwiseConvolution3DOpDescriptor) PaddingValues() foundation.Number {
	rv := objc.Send[foundation.Number](g_.ID, objc.Sel("paddingValues"))
	return rv
}


// SetPaddingValues sets the value of the paddingValues property.
// The padding values for spatial dimensions.

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphdepthwiseconvolution3dopdescriptor/paddingvalues
func (g_ GraphDepthwiseConvolution3DOpDescriptor) SetPaddingValues(value foundation.Number) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setPaddingValues:"), value)
}

// The strides for spatial dimensions.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphdepthwiseconvolution3dopdescriptor/strides
func (g_ GraphDepthwiseConvolution3DOpDescriptor) Strides() foundation.Number {
	rv := objc.Send[foundation.Number](g_.ID, objc.Sel("strides"))
	return rv
}


// SetStrides sets the value of the strides property.
// The strides for spatial dimensions.

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphdepthwiseconvolution3dopdescriptor/strides
func (g_ GraphDepthwiseConvolution3DOpDescriptor) SetStrides(value foundation.Number) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setStrides:"), value)
}


