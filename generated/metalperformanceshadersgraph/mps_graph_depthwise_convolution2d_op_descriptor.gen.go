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
	DilationRateInX() uint
	SetDilationRateInX(value uint)
	PaddingLeft() uint
	SetPaddingLeft(value uint)
	PaddingRight() uint
	SetPaddingRight(value uint)
	DataLayout() GraphTensorNamedDataLayout
	SetDataLayout(value IGraphTensorNamedDataLayout)
	DilationRateInY() int
	SetDilationRateInY(value int)
	PaddingBottom() int
	SetPaddingBottom(value int)
	PaddingStyle() GraphPaddingStyle
	SetPaddingStyle(value GraphPaddingStyle)
	PaddingTop() int
	SetPaddingTop(value int)
	StrideInX() int
	SetStrideInX(value int)
	StrideInY() int
	SetStrideInY(value int)
	WeightsLayout() GraphTensorNamedDataLayout
	SetWeightsLayout(value IGraphTensorNamedDataLayout)
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

// The data layout of the input data in the forward pass.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphdepthwiseconvolution2dopdescriptor/datalayout
func (g_ GraphDepthwiseConvolution2DOpDescriptor) DataLayout() GraphTensorNamedDataLayout {
	rv := objc.Send[GraphTensorNamedDataLayout](g_.ID, objc.Sel("dataLayout"))
	return rv
}


// SetDataLayout sets the value of the dataLayout property.
// The data layout of the input data in the forward pass.

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphdepthwiseconvolution2dopdescriptor/datalayout
func (g_ GraphDepthwiseConvolution2DOpDescriptor) SetDataLayout(value IGraphTensorNamedDataLayout) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setDataLayout:"), value)
}

// The dilation rate for the y dimension.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphdepthwiseconvolution2dopdescriptor/dilationrateiny
func (g_ GraphDepthwiseConvolution2DOpDescriptor) DilationRateInY() int {
	rv := objc.Send[int](g_.ID, objc.Sel("dilationRateInY"))
	return rv
}


// SetDilationRateInY sets the value of the dilationRateInY property.
// The dilation rate for the y dimension.

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphdepthwiseconvolution2dopdescriptor/dilationrateiny
func (g_ GraphDepthwiseConvolution2DOpDescriptor) SetDilationRateInY(value int) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setDilationRateInY:"), value)
}

// The explicit padding value for the y dimension operation adds after the data.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphdepthwiseconvolution2dopdescriptor/paddingbottom
func (g_ GraphDepthwiseConvolution2DOpDescriptor) PaddingBottom() int {
	rv := objc.Send[int](g_.ID, objc.Sel("paddingBottom"))
	return rv
}


// SetPaddingBottom sets the value of the paddingBottom property.
// The explicit padding value for the y dimension operation adds after the data.

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphdepthwiseconvolution2dopdescriptor/paddingbottom
func (g_ GraphDepthwiseConvolution2DOpDescriptor) SetPaddingBottom(value int) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setPaddingBottom:"), value)
}

// The padding style for the operation.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphdepthwiseconvolution2dopdescriptor/paddingstyle
func (g_ GraphDepthwiseConvolution2DOpDescriptor) PaddingStyle() GraphPaddingStyle {
	rv := objc.Send[GraphPaddingStyle](g_.ID, objc.Sel("paddingStyle"))
	return rv
}


// SetPaddingStyle sets the value of the paddingStyle property.
// The padding style for the operation.

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphdepthwiseconvolution2dopdescriptor/paddingstyle
func (g_ GraphDepthwiseConvolution2DOpDescriptor) SetPaddingStyle(value GraphPaddingStyle) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setPaddingStyle:"), value)
}

// The explicit padding value for the y dimension operation adds before the data.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphdepthwiseconvolution2dopdescriptor/paddingtop
func (g_ GraphDepthwiseConvolution2DOpDescriptor) PaddingTop() int {
	rv := objc.Send[int](g_.ID, objc.Sel("paddingTop"))
	return rv
}


// SetPaddingTop sets the value of the paddingTop property.
// The explicit padding value for the y dimension operation adds before the data.

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphdepthwiseconvolution2dopdescriptor/paddingtop
func (g_ GraphDepthwiseConvolution2DOpDescriptor) SetPaddingTop(value int) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setPaddingTop:"), value)
}

// The stride for the x dimension.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphdepthwiseconvolution2dopdescriptor/strideinx
func (g_ GraphDepthwiseConvolution2DOpDescriptor) StrideInX() int {
	rv := objc.Send[int](g_.ID, objc.Sel("strideInX"))
	return rv
}


// SetStrideInX sets the value of the strideInX property.
// The stride for the x dimension.

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphdepthwiseconvolution2dopdescriptor/strideinx
func (g_ GraphDepthwiseConvolution2DOpDescriptor) SetStrideInX(value int) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setStrideInX:"), value)
}

// The stride for the y dimension.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphdepthwiseconvolution2dopdescriptor/strideiny
func (g_ GraphDepthwiseConvolution2DOpDescriptor) StrideInY() int {
	rv := objc.Send[int](g_.ID, objc.Sel("strideInY"))
	return rv
}


// SetStrideInY sets the value of the strideInY property.
// The stride for the y dimension.

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphdepthwiseconvolution2dopdescriptor/strideiny
func (g_ GraphDepthwiseConvolution2DOpDescriptor) SetStrideInY(value int) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setStrideInY:"), value)
}

// The data layout of the weights.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphdepthwiseconvolution2dopdescriptor/weightslayout
func (g_ GraphDepthwiseConvolution2DOpDescriptor) WeightsLayout() GraphTensorNamedDataLayout {
	rv := objc.Send[GraphTensorNamedDataLayout](g_.ID, objc.Sel("weightsLayout"))
	return rv
}


// SetWeightsLayout sets the value of the weightsLayout property.
// The data layout of the weights.

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphdepthwiseconvolution2dopdescriptor/weightslayout
func (g_ GraphDepthwiseConvolution2DOpDescriptor) SetWeightsLayout(value IGraphTensorNamedDataLayout) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setWeightsLayout:"), value)
}



