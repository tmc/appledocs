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
	

	// properties:
	DataLayout() GraphTensorNamedDataLayout
	SetDataLayout(value GraphTensorNamedDataLayout)
	DilationRateInX() uint
	SetDilationRateInX(value uint)
	DilationRateInY() uint
	SetDilationRateInY(value uint)
	PaddingBottom() uint
	SetPaddingBottom(value uint)
	PaddingLeft() uint
	SetPaddingLeft(value uint)
	PaddingRight() uint
	SetPaddingRight(value uint)
	PaddingStyle() GraphPaddingStyle
	SetPaddingStyle(value GraphPaddingStyle)
	PaddingTop() uint
	SetPaddingTop(value uint)
	StrideInX() uint
	SetStrideInX(value uint)
	StrideInY() uint
	SetStrideInY(value uint)
	WeightsLayout() GraphTensorNamedDataLayout
	SetWeightsLayout(value GraphTensorNamedDataLayout)


	

	// methods:
	SetExplicitPaddingWithPaddingLeftPaddingRightPaddingTopPaddingBottom(paddingLeft uint, paddingRight uint, paddingTop uint, paddingBottom uint)


}





// Alloc allocates a new instance without initialization.
func (gc _GraphDepthwiseConvolution2DOpDescriptorClass) Alloc() GraphDepthwiseConvolution2DOpDescriptor {
	rv := objc.Send[GraphDepthwiseConvolution2DOpDescriptor](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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





// A class that defines the parameters for a 2D-depthwise convolution operation.
//
// An defines constant parameters for 2D-depthwise convolutions. Use this class with , , and methods.


// A class that defines the parameters for a 2D-depthwise convolution operation.
//
// [Full Topic]
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






// Creates a 2D-depthwise convolution descriptor with given properties and default values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphDepthwiseConvolution2DOpDescriptor/init(dataLayout:weightsLayout:)
func NewGraphDepthwiseConvolution2DOpDescriptorWithDataLayoutWeightsLayout(dataLayout GraphTensorNamedDataLayout, weightsLayout GraphTensorNamedDataLayout) GraphDepthwiseConvolution2DOpDescriptor {
	rv := objc.Send[GraphDepthwiseConvolution2DOpDescriptor](objc.ID(getGraphDepthwiseConvolution2DOpDescriptorClass().class), objc.Sel("descriptorWithDataLayout:weightsLayout:"), dataLayout, weightsLayout)
	return rv
}


// Creates a 2D-depthwise convolution descriptor with given values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphDepthwiseConvolution2DOpDescriptor/init(strideInX:strideInY:dilationRateInX:dilationRateInY:paddingLeft:paddingRight:paddingTop:paddingBottom:paddingStyle:dataLayout:weightsLayout:)
func NewGraphDepthwiseConvolution2DOpDescriptorWithStrideInXStrideInYDilationRateInXDilationRateInYPaddingLeftPaddingRightPaddingTopPaddingBottomPaddingStyleDataLayoutWeightsLayout(strideInX uint, strideInY uint, dilationRateInX uint, dilationRateInY uint, paddingLeft uint, paddingRight uint, paddingTop uint, paddingBottom uint, paddingStyle GraphPaddingStyle, dataLayout GraphTensorNamedDataLayout, weightsLayout GraphTensorNamedDataLayout) GraphDepthwiseConvolution2DOpDescriptor {
	rv := objc.Send[GraphDepthwiseConvolution2DOpDescriptor](objc.ID(getGraphDepthwiseConvolution2DOpDescriptorClass().class), objc.Sel("descriptorWithStrideInX:strideInY:dilationRateInX:dilationRateInY:paddingLeft:paddingRight:paddingTop:paddingBottom:paddingStyle:dataLayout:weightsLayout:"), strideInX, strideInY, dilationRateInX, dilationRateInY, paddingLeft, paddingRight, paddingTop, paddingBottom, paddingStyle, dataLayout, weightsLayout)
	return rv
}







// Creates a 2D-depthwise convolution descriptor with given properties and default values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphDepthwiseConvolution2DOpDescriptor/init(dataLayout:weightsLayout:)
func (gc _GraphDepthwiseConvolution2DOpDescriptorClass) DescriptorWithDataLayoutWeightsLayout(dataLayout GraphTensorNamedDataLayout, weightsLayout GraphTensorNamedDataLayout) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(gc.class), objc.Sel("descriptorWithDataLayout:weightsLayout:"), dataLayout, weightsLayout)
	return rv
}


// Creates a 2D-depthwise convolution descriptor with given values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphDepthwiseConvolution2DOpDescriptor/init(strideInX:strideInY:dilationRateInX:dilationRateInY:paddingLeft:paddingRight:paddingTop:paddingBottom:paddingStyle:dataLayout:weightsLayout:)
func (gc _GraphDepthwiseConvolution2DOpDescriptorClass) DescriptorWithStrideInXStrideInYDilationRateInXDilationRateInYPaddingLeftPaddingRightPaddingTopPaddingBottomPaddingStyleDataLayoutWeightsLayout(strideInX uint, strideInY uint, dilationRateInX uint, dilationRateInY uint, paddingLeft uint, paddingRight uint, paddingTop uint, paddingBottom uint, paddingStyle GraphPaddingStyle, dataLayout GraphTensorNamedDataLayout, weightsLayout GraphTensorNamedDataLayout) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(gc.class), objc.Sel("descriptorWithStrideInX:strideInY:dilationRateInX:dilationRateInY:paddingLeft:paddingRight:paddingTop:paddingBottom:paddingStyle:dataLayout:weightsLayout:"), strideInX, strideInY, dilationRateInX, dilationRateInY, paddingLeft, paddingRight, paddingTop, paddingBottom, paddingStyle, dataLayout, weightsLayout)
	return rv
}












// Sets the explicit padding values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphDepthwiseConvolution2DOpDescriptor/setExplicitPaddingWithPaddingLeft(_:paddingRight:paddingTop:paddingBottom:)
func (g_ GraphDepthwiseConvolution2DOpDescriptor) SetExplicitPaddingWithPaddingLeftPaddingRightPaddingTopPaddingBottom(paddingLeft uint, paddingRight uint, paddingTop uint, paddingBottom uint) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setExplicitPaddingWithPaddingLeft:paddingRight:paddingTop:paddingBottom:"), paddingLeft, paddingRight, paddingTop, paddingBottom)
}







// The data layout of the input data in the forward pass.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphDepthwiseConvolution2DOpDescriptor/dataLayout
func (g_ GraphDepthwiseConvolution2DOpDescriptor) DataLayout() GraphTensorNamedDataLayout {
	rv := objc.Send[GraphTensorNamedDataLayout](g_.ID, objc.Sel("dataLayout"))
	return rv
}


// The data layout of the input data in the forward pass.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphDepthwiseConvolution2DOpDescriptor/dataLayout
func (g_ GraphDepthwiseConvolution2DOpDescriptor) SetDataLayout(value GraphTensorNamedDataLayout) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setDataLayout:"), value)
}


// The dilation rate for the x dimension.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphDepthwiseConvolution2DOpDescriptor/dilationRateInX
func (g_ GraphDepthwiseConvolution2DOpDescriptor) DilationRateInX() uint {
	rv := objc.Send[uint](g_.ID, objc.Sel("dilationRateInX"))
	return rv
}


// The dilation rate for the x dimension.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphDepthwiseConvolution2DOpDescriptor/dilationRateInX
func (g_ GraphDepthwiseConvolution2DOpDescriptor) SetDilationRateInX(value uint) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setDilationRateInX:"), value)
}


// The dilation rate for the y dimension.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphDepthwiseConvolution2DOpDescriptor/dilationRateInY
func (g_ GraphDepthwiseConvolution2DOpDescriptor) DilationRateInY() uint {
	rv := objc.Send[uint](g_.ID, objc.Sel("dilationRateInY"))
	return rv
}


// The dilation rate for the y dimension.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphDepthwiseConvolution2DOpDescriptor/dilationRateInY
func (g_ GraphDepthwiseConvolution2DOpDescriptor) SetDilationRateInY(value uint) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setDilationRateInY:"), value)
}


// The explicit padding value for the y dimension operation adds after the data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphDepthwiseConvolution2DOpDescriptor/paddingBottom
func (g_ GraphDepthwiseConvolution2DOpDescriptor) PaddingBottom() uint {
	rv := objc.Send[uint](g_.ID, objc.Sel("paddingBottom"))
	return rv
}


// The explicit padding value for the y dimension operation adds after the data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphDepthwiseConvolution2DOpDescriptor/paddingBottom
func (g_ GraphDepthwiseConvolution2DOpDescriptor) SetPaddingBottom(value uint) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setPaddingBottom:"), value)
}


// The explicit padding value for the x dimension the operation adds before the data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphDepthwiseConvolution2DOpDescriptor/paddingLeft
func (g_ GraphDepthwiseConvolution2DOpDescriptor) PaddingLeft() uint {
	rv := objc.Send[uint](g_.ID, objc.Sel("paddingLeft"))
	return rv
}


// The explicit padding value for the x dimension the operation adds before the data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphDepthwiseConvolution2DOpDescriptor/paddingLeft
func (g_ GraphDepthwiseConvolution2DOpDescriptor) SetPaddingLeft(value uint) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setPaddingLeft:"), value)
}


// The explicit padding value for the x dimension operation adds after the data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphDepthwiseConvolution2DOpDescriptor/paddingRight
func (g_ GraphDepthwiseConvolution2DOpDescriptor) PaddingRight() uint {
	rv := objc.Send[uint](g_.ID, objc.Sel("paddingRight"))
	return rv
}


// The explicit padding value for the x dimension operation adds after the data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphDepthwiseConvolution2DOpDescriptor/paddingRight
func (g_ GraphDepthwiseConvolution2DOpDescriptor) SetPaddingRight(value uint) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setPaddingRight:"), value)
}


// The padding style for the operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphDepthwiseConvolution2DOpDescriptor/paddingStyle
func (g_ GraphDepthwiseConvolution2DOpDescriptor) PaddingStyle() GraphPaddingStyle {
	rv := objc.Send[GraphPaddingStyle](g_.ID, objc.Sel("paddingStyle"))
	return rv
}


// The padding style for the operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphDepthwiseConvolution2DOpDescriptor/paddingStyle
func (g_ GraphDepthwiseConvolution2DOpDescriptor) SetPaddingStyle(value GraphPaddingStyle) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setPaddingStyle:"), value)
}


// The explicit padding value for the y dimension operation adds before the data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphDepthwiseConvolution2DOpDescriptor/paddingTop
func (g_ GraphDepthwiseConvolution2DOpDescriptor) PaddingTop() uint {
	rv := objc.Send[uint](g_.ID, objc.Sel("paddingTop"))
	return rv
}


// The explicit padding value for the y dimension operation adds before the data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphDepthwiseConvolution2DOpDescriptor/paddingTop
func (g_ GraphDepthwiseConvolution2DOpDescriptor) SetPaddingTop(value uint) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setPaddingTop:"), value)
}


// The stride for the x dimension.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphDepthwiseConvolution2DOpDescriptor/strideInX
func (g_ GraphDepthwiseConvolution2DOpDescriptor) StrideInX() uint {
	rv := objc.Send[uint](g_.ID, objc.Sel("strideInX"))
	return rv
}


// The stride for the x dimension.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphDepthwiseConvolution2DOpDescriptor/strideInX
func (g_ GraphDepthwiseConvolution2DOpDescriptor) SetStrideInX(value uint) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setStrideInX:"), value)
}


// The stride for the y dimension.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphDepthwiseConvolution2DOpDescriptor/strideInY
func (g_ GraphDepthwiseConvolution2DOpDescriptor) StrideInY() uint {
	rv := objc.Send[uint](g_.ID, objc.Sel("strideInY"))
	return rv
}


// The stride for the y dimension.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphDepthwiseConvolution2DOpDescriptor/strideInY
func (g_ GraphDepthwiseConvolution2DOpDescriptor) SetStrideInY(value uint) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setStrideInY:"), value)
}


// The data layout of the weights.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphDepthwiseConvolution2DOpDescriptor/weightsLayout
func (g_ GraphDepthwiseConvolution2DOpDescriptor) WeightsLayout() GraphTensorNamedDataLayout {
	rv := objc.Send[GraphTensorNamedDataLayout](g_.ID, objc.Sel("weightsLayout"))
	return rv
}


// The data layout of the weights.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphDepthwiseConvolution2DOpDescriptor/weightsLayout
func (g_ GraphDepthwiseConvolution2DOpDescriptor) SetWeightsLayout(value GraphTensorNamedDataLayout) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setWeightsLayout:"), value)
}







