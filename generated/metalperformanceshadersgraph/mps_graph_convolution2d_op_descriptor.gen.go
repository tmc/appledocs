// Code generated from Apple documentation for MetalPerformanceShadersGraph. DO NOT EDIT.

package metalperformanceshadersgraph

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [GraphConvolution2DOpDescriptor] class.
var (
	GraphConvolution2DOpDescriptorClass     _GraphConvolution2DOpDescriptorClass
	GraphConvolution2DOpDescriptorClassOnce sync.Once
)

func getGraphConvolution2DOpDescriptorClass() _GraphConvolution2DOpDescriptorClass {
	GraphConvolution2DOpDescriptorClassOnce.Do(func() {
		GraphConvolution2DOpDescriptorClass = _GraphConvolution2DOpDescriptorClass{objc.GetClass("MPSGraphConvolution2DOpDescriptor")}
	})
	return GraphConvolution2DOpDescriptorClass
}

type _GraphConvolution2DOpDescriptorClass struct {
	class objc.Class
}

// An interface definition for the [GraphConvolution2DOpDescriptor] class.
type IGraphConvolution2DOpDescriptor interface {
	IGraphObject
	// properties:
	DataLayout() GraphTensorNamedDataLayout
	SetDataLayout(value GraphTensorNamedDataLayout)
	DilationRateInX() int
	SetDilationRateInX(value int)
	DilationRateInY() int
	SetDilationRateInY(value int)
	Groups() int
	SetGroups(value int)
	PaddingBottom() int
	SetPaddingBottom(value int)
	PaddingLeft() int
	SetPaddingLeft(value int)
	PaddingRight() int
	SetPaddingRight(value int)
	PaddingStyle() GraphPaddingStyle
	SetPaddingStyle(value GraphPaddingStyle)
	PaddingTop() int
	SetPaddingTop(value int)
	StrideInX() int
	SetStrideInX(value int)
	StrideInY() int
	SetStrideInY(value int)
	WeightsLayout() GraphTensorNamedDataLayout
	SetWeightsLayout(value GraphTensorNamedDataLayout)
	// methods:
}

// A class that describes the properties of a 2D-convolution operator.
//
// Use an instance of this class is to add a 2D-convolution operator with the desired properties to the graph.


// A class that describes the properties of a 2D-convolution operator.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphConvolution2DOpDescriptor
type GraphConvolution2DOpDescriptor struct {
	GraphObject
}

// GraphConvolution2DOpDescriptorFrom constructs a [GraphConvolution2DOpDescriptor] from an unsafe.Pointer.
//
// A class that describes the properties of a 2D-convolution operator.
func GraphConvolution2DOpDescriptorFrom(ptr unsafe.Pointer) GraphConvolution2DOpDescriptor {
	return GraphConvolution2DOpDescriptor{
		GraphObject: GraphObjectFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (gc _GraphConvolution2DOpDescriptorClass) Alloc() GraphConvolution2DOpDescriptor {
	rv := objc.Send[GraphConvolution2DOpDescriptor](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (gc _GraphConvolution2DOpDescriptorClass) New() GraphConvolution2DOpDescriptor {
	rv := objc.Send[GraphConvolution2DOpDescriptor](objc.ID(gc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (g_ GraphConvolution2DOpDescriptor) Init() GraphConvolution2DOpDescriptor {
	rv := objc.Send[GraphConvolution2DOpDescriptor](g_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g_ GraphConvolution2DOpDescriptor) Autorelease() GraphConvolution2DOpDescriptor {
	rv := objc.Send[GraphConvolution2DOpDescriptor](g_.ID, objc.Sel("autorelease"))
	return rv
}

// NewGraphConvolution2DOpDescriptor creates a new GraphConvolution2DOpDescriptor instance.
func NewGraphConvolution2DOpDescriptor() GraphConvolution2DOpDescriptor {
	return getGraphConvolution2DOpDescriptorClass().New()
}



// Creates a convolution descriptor with given values for parameters.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphConvolution2DOpDescriptor/init(strideInX:strideInY:dilationRateInX:dilationRateInY:groups:paddingLeft:paddingRight:paddingTop:paddingBottom:paddingStyle:dataLayout:weightsLayout:)
func NewGraphConvolution2DOpDescriptorWithStrideInXStrideInYDilationRateInXDilationRateInYGroupsPaddingLeftPaddingRightPaddingTopPaddingBottomPaddingStyleDataLayoutWeightsLayout(strideInX uint, strideInY uint, dilationRateInX uint, dilationRateInY uint, groups uint, paddingLeft uint, paddingRight uint, paddingTop uint, paddingBottom uint, paddingStyle GraphPaddingStyle, dataLayout GraphTensorNamedDataLayout, weightsLayout GraphTensorNamedDataLayout) GraphConvolution2DOpDescriptor {
	rv := objc.Send[GraphConvolution2DOpDescriptor](objc.ID(getGraphConvolution2DOpDescriptorClass().class), objc.Sel("descriptorWithStrideInX:strideInY:dilationRateInX:dilationRateInY:groups:paddingLeft:paddingRight:paddingTop:paddingBottom:paddingStyle:dataLayout:weightsLayout:"), strideInX, strideInY, dilationRateInX, dilationRateInY, groups, paddingLeft, paddingRight, paddingTop, paddingBottom, paddingStyle, dataLayout, weightsLayout)
	return rv
}



// Creates a convolution descriptor with given values for parameters.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphConvolution2DOpDescriptor/init(strideInX:strideInY:dilationRateInX:dilationRateInY:groups:paddingLeft:paddingRight:paddingTop:paddingBottom:paddingStyle:dataLayout:weightsLayout:)
func (gc _GraphConvolution2DOpDescriptorClass) DescriptorWithStrideInXStrideInYDilationRateInXDilationRateInYGroupsPaddingLeftPaddingRightPaddingTopPaddingBottomPaddingStyleDataLayoutWeightsLayout(strideInX uint, strideInY uint, dilationRateInX uint, dilationRateInY uint, groups uint, paddingLeft uint, paddingRight uint, paddingTop uint, paddingBottom uint, paddingStyle GraphPaddingStyle, dataLayout GraphTensorNamedDataLayout, weightsLayout GraphTensorNamedDataLayout) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(gc.class), objc.Sel("descriptorWithStrideInX:strideInY:dilationRateInX:dilationRateInY:groups:paddingLeft:paddingRight:paddingTop:paddingBottom:paddingStyle:dataLayout:weightsLayout:"), strideInX, strideInY, dilationRateInX, dilationRateInY, groups, paddingLeft, paddingRight, paddingTop, paddingBottom, paddingStyle, dataLayout, weightsLayout)
	return rv
}


// The named layout of data in the source tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphconvolution2dopdescriptor/datalayout
func (g_ GraphConvolution2DOpDescriptor) DataLayout() GraphTensorNamedDataLayout {
	rv := objc.Send[GraphTensorNamedDataLayout](g_.ID, objc.Sel("dataLayout"))
	return rv
}


// The named layout of data in the source tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphconvolution2dopdescriptor/datalayout
func (g_ GraphConvolution2DOpDescriptor) SetDataLayout(value GraphTensorNamedDataLayout) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setDataLayout:"), value)
}


// The amount by which the weights tensor expands in the
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphconvolution2dopdescriptor/dilationrateinx
func (g_ GraphConvolution2DOpDescriptor) DilationRateInX() int {
	rv := objc.Send[int](g_.ID, objc.Sel("dilationRateInX"))
	return rv
}


// The amount by which the weights tensor expands in the
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphconvolution2dopdescriptor/dilationrateinx
func (g_ GraphConvolution2DOpDescriptor) SetDilationRateInX(value int) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setDilationRateInX:"), value)
}


// The amount by which the weights tensor expands in the
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphconvolution2dopdescriptor/dilationrateiny
func (g_ GraphConvolution2DOpDescriptor) DilationRateInY() int {
	rv := objc.Send[int](g_.ID, objc.Sel("dilationRateInY"))
	return rv
}


// The amount by which the weights tensor expands in the
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphconvolution2dopdescriptor/dilationrateiny
func (g_ GraphConvolution2DOpDescriptor) SetDilationRateInY(value int) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setDilationRateInY:"), value)
}


// The number of partitions of the input and output channels.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphconvolution2dopdescriptor/groups
func (g_ GraphConvolution2DOpDescriptor) Groups() int {
	rv := objc.Send[int](g_.ID, objc.Sel("groups"))
	return rv
}


// The number of partitions of the input and output channels.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphconvolution2dopdescriptor/groups
func (g_ GraphConvolution2DOpDescriptor) SetGroups(value int) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setGroups:"), value)
}


// The number of zeros added at the bottom of the source tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphconvolution2dopdescriptor/paddingbottom
func (g_ GraphConvolution2DOpDescriptor) PaddingBottom() int {
	rv := objc.Send[int](g_.ID, objc.Sel("paddingBottom"))
	return rv
}


// The number of zeros added at the bottom of the source tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphconvolution2dopdescriptor/paddingbottom
func (g_ GraphConvolution2DOpDescriptor) SetPaddingBottom(value int) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setPaddingBottom:"), value)
}


// The number of zeros added on the left side of the source tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphconvolution2dopdescriptor/paddingleft
func (g_ GraphConvolution2DOpDescriptor) PaddingLeft() int {
	rv := objc.Send[int](g_.ID, objc.Sel("paddingLeft"))
	return rv
}


// The number of zeros added on the left side of the source tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphconvolution2dopdescriptor/paddingleft
func (g_ GraphConvolution2DOpDescriptor) SetPaddingLeft(value int) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setPaddingLeft:"), value)
}


// The number of zeros added on the right side of the source tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphconvolution2dopdescriptor/paddingright
func (g_ GraphConvolution2DOpDescriptor) PaddingRight() int {
	rv := objc.Send[int](g_.ID, objc.Sel("paddingRight"))
	return rv
}


// The number of zeros added on the right side of the source tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphconvolution2dopdescriptor/paddingright
func (g_ GraphConvolution2DOpDescriptor) SetPaddingRight(value int) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setPaddingRight:"), value)
}


// The type of padding applied to the source tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphconvolution2dopdescriptor/paddingstyle
func (g_ GraphConvolution2DOpDescriptor) PaddingStyle() GraphPaddingStyle {
	rv := objc.Send[GraphPaddingStyle](g_.ID, objc.Sel("paddingStyle"))
	return rv
}


// The type of padding applied to the source tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphconvolution2dopdescriptor/paddingstyle
func (g_ GraphConvolution2DOpDescriptor) SetPaddingStyle(value GraphPaddingStyle) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setPaddingStyle:"), value)
}


// The number of zeros added at the top of the source tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphconvolution2dopdescriptor/paddingtop
func (g_ GraphConvolution2DOpDescriptor) PaddingTop() int {
	rv := objc.Send[int](g_.ID, objc.Sel("paddingTop"))
	return rv
}


// The number of zeros added at the top of the source tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphconvolution2dopdescriptor/paddingtop
func (g_ GraphConvolution2DOpDescriptor) SetPaddingTop(value int) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setPaddingTop:"), value)
}


// The scale that maps
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphconvolution2dopdescriptor/strideinx
func (g_ GraphConvolution2DOpDescriptor) StrideInX() int {
	rv := objc.Send[int](g_.ID, objc.Sel("strideInX"))
	return rv
}


// The scale that maps
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphconvolution2dopdescriptor/strideinx
func (g_ GraphConvolution2DOpDescriptor) SetStrideInX(value int) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setStrideInX:"), value)
}


// The scale that maps
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphconvolution2dopdescriptor/strideiny
func (g_ GraphConvolution2DOpDescriptor) StrideInY() int {
	rv := objc.Send[int](g_.ID, objc.Sel("strideInY"))
	return rv
}


// The scale that maps
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphconvolution2dopdescriptor/strideiny
func (g_ GraphConvolution2DOpDescriptor) SetStrideInY(value int) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setStrideInY:"), value)
}


// The named layout of data in the weights tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphconvolution2dopdescriptor/weightslayout
func (g_ GraphConvolution2DOpDescriptor) WeightsLayout() GraphTensorNamedDataLayout {
	rv := objc.Send[GraphTensorNamedDataLayout](g_.ID, objc.Sel("weightsLayout"))
	return rv
}


// The named layout of data in the weights tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphconvolution2dopdescriptor/weightslayout
func (g_ GraphConvolution2DOpDescriptor) SetWeightsLayout(value GraphTensorNamedDataLayout) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setWeightsLayout:"), value)
}


