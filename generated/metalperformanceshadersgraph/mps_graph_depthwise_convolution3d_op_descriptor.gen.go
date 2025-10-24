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
	// properties:
	ChannelDimensionIndex() int
	SetChannelDimensionIndex(value int)
	DilationRates() objc.IObject /* cross-framework: NSNumber */
	SetDilationRates(value objc.IObject /* cross-framework: NSNumber */)
	PaddingStyle() GraphPaddingStyle
	SetPaddingStyle(value GraphPaddingStyle)
	PaddingValues() objc.IObject /* cross-framework: NSNumber */
	SetPaddingValues(value objc.IObject /* cross-framework: NSNumber */)
	Strides() objc.IObject /* cross-framework: NSNumber */
	SetStrides(value objc.IObject /* cross-framework: NSNumber */)
	// methods:
}

// The class that defines the parameters for a 3D-depthwise convolution operation.
//
// A defines constant parameters for 3D depthwise convolutions. Use this class with , and methods.


// The class that defines the parameters for a 3D-depthwise convolution operation.
//
// [Full Topic]
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



// The axis that contains the channels in the input and the weights, within
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphdepthwiseconvolution3dopdescriptor/channeldimensionindex
func (g_ GraphDepthwiseConvolution3DOpDescriptor) ChannelDimensionIndex() int {
	rv := objc.Send[int](g_.ID, objc.Sel("channelDimensionIndex"))
	return rv
}


// The axis that contains the channels in the input and the weights, within
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphdepthwiseconvolution3dopdescriptor/channeldimensionindex
func (g_ GraphDepthwiseConvolution3DOpDescriptor) SetChannelDimensionIndex(value int) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setChannelDimensionIndex:"), value)
}


// The dilation rates for spatial dimensions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphdepthwiseconvolution3dopdescriptor/dilationrates
func (g_ GraphDepthwiseConvolution3DOpDescriptor) DilationRates() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](g_.ID, objc.Sel("dilationRates"))
	return rv
}


// The dilation rates for spatial dimensions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphdepthwiseconvolution3dopdescriptor/dilationrates
func (g_ GraphDepthwiseConvolution3DOpDescriptor) SetDilationRates(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setDilationRates:"), value)
}


// The padding style for the operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphdepthwiseconvolution3dopdescriptor/paddingstyle
func (g_ GraphDepthwiseConvolution3DOpDescriptor) PaddingStyle() GraphPaddingStyle {
	rv := objc.Send[GraphPaddingStyle](g_.ID, objc.Sel("paddingStyle"))
	return rv
}


// The padding style for the operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphdepthwiseconvolution3dopdescriptor/paddingstyle
func (g_ GraphDepthwiseConvolution3DOpDescriptor) SetPaddingStyle(value GraphPaddingStyle) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setPaddingStyle:"), value)
}


// The padding values for spatial dimensions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphdepthwiseconvolution3dopdescriptor/paddingvalues
func (g_ GraphDepthwiseConvolution3DOpDescriptor) PaddingValues() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](g_.ID, objc.Sel("paddingValues"))
	return rv
}


// The padding values for spatial dimensions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphdepthwiseconvolution3dopdescriptor/paddingvalues
func (g_ GraphDepthwiseConvolution3DOpDescriptor) SetPaddingValues(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setPaddingValues:"), value)
}


// The strides for spatial dimensions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphdepthwiseconvolution3dopdescriptor/strides
func (g_ GraphDepthwiseConvolution3DOpDescriptor) Strides() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](g_.ID, objc.Sel("strides"))
	return rv
}


// The strides for spatial dimensions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphdepthwiseconvolution3dopdescriptor/strides
func (g_ GraphDepthwiseConvolution3DOpDescriptor) SetStrides(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setStrides:"), value)
}



