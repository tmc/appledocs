// Code generated from Apple documentation for MetalPerformanceShadersGraph. DO NOT EDIT.

package metalperformanceshadersgraph

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)





// The class instance for the [GraphConvolution3DOpDescriptor] class.
var (
	GraphConvolution3DOpDescriptorClass     _GraphConvolution3DOpDescriptorClass
	GraphConvolution3DOpDescriptorClassOnce sync.Once
)

func getGraphConvolution3DOpDescriptorClass() _GraphConvolution3DOpDescriptorClass {
	GraphConvolution3DOpDescriptorClassOnce.Do(func() {
		GraphConvolution3DOpDescriptorClass = _GraphConvolution3DOpDescriptorClass{objc.GetClass("MPSGraphConvolution3DOpDescriptor")}
	})
	return GraphConvolution3DOpDescriptorClass
}

type _GraphConvolution3DOpDescriptorClass struct {
	class objc.Class
}





// An interface definition for the [GraphConvolution3DOpDescriptor] class.
type IGraphConvolution3DOpDescriptor interface {
	IGraphObject
	

	// properties:
	DataLayout() GraphTensorNamedDataLayout
	SetDataLayout(value GraphTensorNamedDataLayout)
	DilationRateInX() uint
	SetDilationRateInX(value uint)
	DilationRateInY() uint
	SetDilationRateInY(value uint)
	DilationRateInZ() uint
	SetDilationRateInZ(value uint)
	Groups() uint
	SetGroups(value uint)
	PaddingBack() uint
	SetPaddingBack(value uint)
	PaddingBottom() uint
	SetPaddingBottom(value uint)
	PaddingFront() uint
	SetPaddingFront(value uint)
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
	StrideInZ() uint
	SetStrideInZ(value uint)
	WeightsLayout() GraphTensorNamedDataLayout
	SetWeightsLayout(value GraphTensorNamedDataLayout)


	

	// methods:
	SetExplicitPaddingWithPaddingLeftPaddingRightPaddingTopPaddingBottomPaddingFrontPaddingBack(paddingLeft uint, paddingRight uint, paddingTop uint, paddingBottom uint, paddingFront uint, paddingBack uint)


}





// Alloc allocates a new instance without initialization.
func (gc _GraphConvolution3DOpDescriptorClass) Alloc() GraphConvolution3DOpDescriptor {
	rv := objc.Send[GraphConvolution3DOpDescriptor](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (gc _GraphConvolution3DOpDescriptorClass) New() GraphConvolution3DOpDescriptor {
	rv := objc.Send[GraphConvolution3DOpDescriptor](objc.ID(gc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (g_ GraphConvolution3DOpDescriptor) Init() GraphConvolution3DOpDescriptor {
	rv := objc.Send[GraphConvolution3DOpDescriptor](g_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g_ GraphConvolution3DOpDescriptor) Autorelease() GraphConvolution3DOpDescriptor {
	rv := objc.Send[GraphConvolution3DOpDescriptor](g_.ID, objc.Sel("autorelease"))
	return rv
}

// NewGraphConvolution3DOpDescriptor creates a new GraphConvolution3DOpDescriptor instance.
func NewGraphConvolution3DOpDescriptor() GraphConvolution3DOpDescriptor {
	return getGraphConvolution3DOpDescriptorClass().New()
}





// A class that describes the properties of a 3D-convolution operator.
//
// Use an instance of this class is to add a 3D-convolution operator with desired properties to the graph.


// A class that describes the properties of a 3D-convolution operator.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphConvolution3DOpDescriptor
type GraphConvolution3DOpDescriptor struct {
	GraphObject
}

// GraphConvolution3DOpDescriptorFrom constructs a [GraphConvolution3DOpDescriptor] from an unsafe.Pointer.
//
// A class that describes the properties of a 3D-convolution operator.
func GraphConvolution3DOpDescriptorFrom(ptr unsafe.Pointer) GraphConvolution3DOpDescriptor {
	return GraphConvolution3DOpDescriptor{
		GraphObject: GraphObjectFrom(ptr),
	}
}






// Creates a convolution descriptor with given values for parameters.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphConvolution3DOpDescriptor/init(strideInX:strideInY:strideInZ:dilationRateInX:dilationRateInY:dilationRateInZ:groups:paddingLeft:paddingRight:paddingTop:paddingBottom:paddingFront:paddingBack:paddingStyle:dataLayout:weightsLayout:)
func NewGraphConvolution3DOpDescriptorWithStrideInXStrideInYStrideInZDilationRateInXDilationRateInYDilationRateInZGroupsPaddingLeftPaddingRightPaddingTopPaddingBottomPaddingFrontPaddingBackPaddingStyleDataLayoutWeightsLayout(strideInX uint, strideInY uint, strideInZ uint, dilationRateInX uint, dilationRateInY uint, dilationRateInZ uint, groups uint, paddingLeft uint, paddingRight uint, paddingTop uint, paddingBottom uint, paddingFront uint, paddingBack uint, paddingStyle GraphPaddingStyle, dataLayout GraphTensorNamedDataLayout, weightsLayout GraphTensorNamedDataLayout) GraphConvolution3DOpDescriptor {
	rv := objc.Send[GraphConvolution3DOpDescriptor](objc.ID(getGraphConvolution3DOpDescriptorClass().class), objc.Sel("descriptorWithStrideInX:strideInY:strideInZ:dilationRateInX:dilationRateInY:dilationRateInZ:groups:paddingLeft:paddingRight:paddingTop:paddingBottom:paddingFront:paddingBack:paddingStyle:dataLayout:weightsLayout:"), strideInX, strideInY, strideInZ, dilationRateInX, dilationRateInY, dilationRateInZ, groups, paddingLeft, paddingRight, paddingTop, paddingBottom, paddingFront, paddingBack, paddingStyle, dataLayout, weightsLayout)
	return rv
}


// Creates a convolution descriptor with given values for parameters.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphConvolution3DOpDescriptor/init(strideInX:strideInY:strideInZ:dilationRateInX:dilationRateInY:dilationRateInZ:groups:paddingStyle:dataLayout:weightsLayout:)
func NewGraphConvolution3DOpDescriptorWithStrideInXStrideInYStrideInZDilationRateInXDilationRateInYDilationRateInZGroupsPaddingStyleDataLayoutWeightsLayout(strideInX uint, strideInY uint, strideInZ uint, dilationRateInX uint, dilationRateInY uint, dilationRateInZ uint, groups uint, paddingStyle GraphPaddingStyle, dataLayout GraphTensorNamedDataLayout, weightsLayout GraphTensorNamedDataLayout) GraphConvolution3DOpDescriptor {
	rv := objc.Send[GraphConvolution3DOpDescriptor](objc.ID(getGraphConvolution3DOpDescriptorClass().class), objc.Sel("descriptorWithStrideInX:strideInY:strideInZ:dilationRateInX:dilationRateInY:dilationRateInZ:groups:paddingStyle:dataLayout:weightsLayout:"), strideInX, strideInY, strideInZ, dilationRateInX, dilationRateInY, dilationRateInZ, groups, paddingStyle, dataLayout, weightsLayout)
	return rv
}







// Creates a convolution descriptor with given values for parameters.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphConvolution3DOpDescriptor/init(strideInX:strideInY:strideInZ:dilationRateInX:dilationRateInY:dilationRateInZ:groups:paddingLeft:paddingRight:paddingTop:paddingBottom:paddingFront:paddingBack:paddingStyle:dataLayout:weightsLayout:)
func (gc _GraphConvolution3DOpDescriptorClass) DescriptorWithStrideInXStrideInYStrideInZDilationRateInXDilationRateInYDilationRateInZGroupsPaddingLeftPaddingRightPaddingTopPaddingBottomPaddingFrontPaddingBackPaddingStyleDataLayoutWeightsLayout(strideInX uint, strideInY uint, strideInZ uint, dilationRateInX uint, dilationRateInY uint, dilationRateInZ uint, groups uint, paddingLeft uint, paddingRight uint, paddingTop uint, paddingBottom uint, paddingFront uint, paddingBack uint, paddingStyle GraphPaddingStyle, dataLayout GraphTensorNamedDataLayout, weightsLayout GraphTensorNamedDataLayout) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(gc.class), objc.Sel("descriptorWithStrideInX:strideInY:strideInZ:dilationRateInX:dilationRateInY:dilationRateInZ:groups:paddingLeft:paddingRight:paddingTop:paddingBottom:paddingFront:paddingBack:paddingStyle:dataLayout:weightsLayout:"), strideInX, strideInY, strideInZ, dilationRateInX, dilationRateInY, dilationRateInZ, groups, paddingLeft, paddingRight, paddingTop, paddingBottom, paddingFront, paddingBack, paddingStyle, dataLayout, weightsLayout)
	return rv
}


// Creates a convolution descriptor with given values for parameters.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphConvolution3DOpDescriptor/init(strideInX:strideInY:strideInZ:dilationRateInX:dilationRateInY:dilationRateInZ:groups:paddingStyle:dataLayout:weightsLayout:)
func (gc _GraphConvolution3DOpDescriptorClass) DescriptorWithStrideInXStrideInYStrideInZDilationRateInXDilationRateInYDilationRateInZGroupsPaddingStyleDataLayoutWeightsLayout(strideInX uint, strideInY uint, strideInZ uint, dilationRateInX uint, dilationRateInY uint, dilationRateInZ uint, groups uint, paddingStyle GraphPaddingStyle, dataLayout GraphTensorNamedDataLayout, weightsLayout GraphTensorNamedDataLayout) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(gc.class), objc.Sel("descriptorWithStrideInX:strideInY:strideInZ:dilationRateInX:dilationRateInY:dilationRateInZ:groups:paddingStyle:dataLayout:weightsLayout:"), strideInX, strideInY, strideInZ, dilationRateInX, dilationRateInY, dilationRateInZ, groups, paddingStyle, dataLayout, weightsLayout)
	return rv
}












// Sets the left, right, top, bottom, front, and back padding values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphConvolution3DOpDescriptor/setExplicitPaddingWithPaddingLeft(_:paddingRight:paddingTop:paddingBottom:paddingFront:paddingBack:)
func (g_ GraphConvolution3DOpDescriptor) SetExplicitPaddingWithPaddingLeftPaddingRightPaddingTopPaddingBottomPaddingFrontPaddingBack(paddingLeft uint, paddingRight uint, paddingTop uint, paddingBottom uint, paddingFront uint, paddingBack uint) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setExplicitPaddingWithPaddingLeft:paddingRight:paddingTop:paddingBottom:paddingFront:paddingBack:"), paddingLeft, paddingRight, paddingTop, paddingBottom, paddingFront, paddingBack)
}







// The named layout of data in the source tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphConvolution3DOpDescriptor/dataLayout
func (g_ GraphConvolution3DOpDescriptor) DataLayout() GraphTensorNamedDataLayout {
	rv := objc.Send[GraphTensorNamedDataLayout](g_.ID, objc.Sel("dataLayout"))
	return rv
}


// The named layout of data in the source tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphConvolution3DOpDescriptor/dataLayout
func (g_ GraphConvolution3DOpDescriptor) SetDataLayout(value GraphTensorNamedDataLayout) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setDataLayout:"), value)
}


// The amount by which weights tensor expands in the -direction.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphConvolution3DOpDescriptor/dilationRateInX
func (g_ GraphConvolution3DOpDescriptor) DilationRateInX() uint {
	rv := objc.Send[uint](g_.ID, objc.Sel("dilationRateInX"))
	return rv
}


// The amount by which weights tensor expands in the -direction.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphConvolution3DOpDescriptor/dilationRateInX
func (g_ GraphConvolution3DOpDescriptor) SetDilationRateInX(value uint) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setDilationRateInX:"), value)
}


// The amount by which weights tensor expands in the -direction.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphConvolution3DOpDescriptor/dilationRateInY
func (g_ GraphConvolution3DOpDescriptor) DilationRateInY() uint {
	rv := objc.Send[uint](g_.ID, objc.Sel("dilationRateInY"))
	return rv
}


// The amount by which weights tensor expands in the -direction.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphConvolution3DOpDescriptor/dilationRateInY
func (g_ GraphConvolution3DOpDescriptor) SetDilationRateInY(value uint) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setDilationRateInY:"), value)
}


// The amount by which weights tensor expands in the -direction.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphConvolution3DOpDescriptor/dilationRateInZ
func (g_ GraphConvolution3DOpDescriptor) DilationRateInZ() uint {
	rv := objc.Send[uint](g_.ID, objc.Sel("dilationRateInZ"))
	return rv
}


// The amount by which weights tensor expands in the -direction.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphConvolution3DOpDescriptor/dilationRateInZ
func (g_ GraphConvolution3DOpDescriptor) SetDilationRateInZ(value uint) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setDilationRateInZ:"), value)
}


// The number of partitions of the input and output channels.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphConvolution3DOpDescriptor/groups
func (g_ GraphConvolution3DOpDescriptor) Groups() uint {
	rv := objc.Send[uint](g_.ID, objc.Sel("groups"))
	return rv
}


// The number of partitions of the input and output channels.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphConvolution3DOpDescriptor/groups
func (g_ GraphConvolution3DOpDescriptor) SetGroups(value uint) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setGroups:"), value)
}


// The number of zeros added at the back of the source tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphConvolution3DOpDescriptor/paddingBack
func (g_ GraphConvolution3DOpDescriptor) PaddingBack() uint {
	rv := objc.Send[uint](g_.ID, objc.Sel("paddingBack"))
	return rv
}


// The number of zeros added at the back of the source tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphConvolution3DOpDescriptor/paddingBack
func (g_ GraphConvolution3DOpDescriptor) SetPaddingBack(value uint) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setPaddingBack:"), value)
}


// The number of zeros added at the bottom of the source tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphConvolution3DOpDescriptor/paddingBottom
func (g_ GraphConvolution3DOpDescriptor) PaddingBottom() uint {
	rv := objc.Send[uint](g_.ID, objc.Sel("paddingBottom"))
	return rv
}


// The number of zeros added at the bottom of the source tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphConvolution3DOpDescriptor/paddingBottom
func (g_ GraphConvolution3DOpDescriptor) SetPaddingBottom(value uint) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setPaddingBottom:"), value)
}


// The number of zeros added at the front of the source tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphConvolution3DOpDescriptor/paddingFront
func (g_ GraphConvolution3DOpDescriptor) PaddingFront() uint {
	rv := objc.Send[uint](g_.ID, objc.Sel("paddingFront"))
	return rv
}


// The number of zeros added at the front of the source tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphConvolution3DOpDescriptor/paddingFront
func (g_ GraphConvolution3DOpDescriptor) SetPaddingFront(value uint) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setPaddingFront:"), value)
}


// The number of zeros added on the left side of the source tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphConvolution3DOpDescriptor/paddingLeft
func (g_ GraphConvolution3DOpDescriptor) PaddingLeft() uint {
	rv := objc.Send[uint](g_.ID, objc.Sel("paddingLeft"))
	return rv
}


// The number of zeros added on the left side of the source tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphConvolution3DOpDescriptor/paddingLeft
func (g_ GraphConvolution3DOpDescriptor) SetPaddingLeft(value uint) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setPaddingLeft:"), value)
}


// The number of zeros added on the right side of the source tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphConvolution3DOpDescriptor/paddingRight
func (g_ GraphConvolution3DOpDescriptor) PaddingRight() uint {
	rv := objc.Send[uint](g_.ID, objc.Sel("paddingRight"))
	return rv
}


// The number of zeros added on the right side of the source tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphConvolution3DOpDescriptor/paddingRight
func (g_ GraphConvolution3DOpDescriptor) SetPaddingRight(value uint) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setPaddingRight:"), value)
}


// The type of padding that is applied to the source tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphConvolution3DOpDescriptor/paddingStyle
func (g_ GraphConvolution3DOpDescriptor) PaddingStyle() GraphPaddingStyle {
	rv := objc.Send[GraphPaddingStyle](g_.ID, objc.Sel("paddingStyle"))
	return rv
}


// The type of padding that is applied to the source tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphConvolution3DOpDescriptor/paddingStyle
func (g_ GraphConvolution3DOpDescriptor) SetPaddingStyle(value GraphPaddingStyle) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setPaddingStyle:"), value)
}


// The number of zeros added at the top of the source tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphConvolution3DOpDescriptor/paddingTop
func (g_ GraphConvolution3DOpDescriptor) PaddingTop() uint {
	rv := objc.Send[uint](g_.ID, objc.Sel("paddingTop"))
	return rv
}


// The number of zeros added at the top of the source tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphConvolution3DOpDescriptor/paddingTop
func (g_ GraphConvolution3DOpDescriptor) SetPaddingTop(value uint) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setPaddingTop:"), value)
}


// The scale that maps -coordinate of destination to -coordinate of source.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphConvolution3DOpDescriptor/strideInX
func (g_ GraphConvolution3DOpDescriptor) StrideInX() uint {
	rv := objc.Send[uint](g_.ID, objc.Sel("strideInX"))
	return rv
}


// The scale that maps -coordinate of destination to -coordinate of source.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphConvolution3DOpDescriptor/strideInX
func (g_ GraphConvolution3DOpDescriptor) SetStrideInX(value uint) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setStrideInX:"), value)
}


// The scale that maps -coordinate of destination to -coordinate of source.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphConvolution3DOpDescriptor/strideInY
func (g_ GraphConvolution3DOpDescriptor) StrideInY() uint {
	rv := objc.Send[uint](g_.ID, objc.Sel("strideInY"))
	return rv
}


// The scale that maps -coordinate of destination to -coordinate of source.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphConvolution3DOpDescriptor/strideInY
func (g_ GraphConvolution3DOpDescriptor) SetStrideInY(value uint) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setStrideInY:"), value)
}


// The scale that maps -coordinate of destination to -coordinate of source.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphConvolution3DOpDescriptor/strideInZ
func (g_ GraphConvolution3DOpDescriptor) StrideInZ() uint {
	rv := objc.Send[uint](g_.ID, objc.Sel("strideInZ"))
	return rv
}


// The scale that maps -coordinate of destination to -coordinate of source.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphConvolution3DOpDescriptor/strideInZ
func (g_ GraphConvolution3DOpDescriptor) SetStrideInZ(value uint) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setStrideInZ:"), value)
}


// The named layout of data in the weights tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphConvolution3DOpDescriptor/weightsLayout
func (g_ GraphConvolution3DOpDescriptor) WeightsLayout() GraphTensorNamedDataLayout {
	rv := objc.Send[GraphTensorNamedDataLayout](g_.ID, objc.Sel("weightsLayout"))
	return rv
}


// The named layout of data in the weights tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphConvolution3DOpDescriptor/weightsLayout
func (g_ GraphConvolution3DOpDescriptor) SetWeightsLayout(value GraphTensorNamedDataLayout) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setWeightsLayout:"), value)
}







