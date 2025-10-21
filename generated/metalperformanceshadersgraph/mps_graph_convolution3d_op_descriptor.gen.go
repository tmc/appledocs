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
}

// A class that describes the properties of a 3D-convolution operator.
//
// Use an instance of this class is to add a 3D-convolution operator with desired properties to the graph.
//
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

// Alloc allocates a new instance without initialization.
func (gc _GraphConvolution3DOpDescriptorClass) Alloc() GraphConvolution3DOpDescriptor {
	rv := objc.Send[GraphConvolution3DOpDescriptor](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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




// Creates a convolution descriptor with given values for parameters.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphConvolution3DOpDescriptor/init(strideInX:strideInY:strideInZ:dilationRateInX:dilationRateInY:dilationRateInZ:groups:paddingLeft:paddingRight:paddingTop:paddingBottom:paddingFront:paddingBack:paddingStyle:dataLayout:weightsLayout:)
func NewGraphConvolution3DOpDescriptorWithStrideInXStrideInYStrideInZDilationRateInXDilationRateInYDilationRateInZGroupsPaddingLeftPaddingRightPaddingTopPaddingBottomPaddingFrontPaddingBackPaddingStyleDataLayoutWeightsLayout(strideInX uint, strideInY uint, strideInZ uint, dilationRateInX uint, dilationRateInY uint, dilationRateInZ uint, groups uint, paddingLeft uint, paddingRight uint, paddingTop uint, paddingBottom uint, paddingFront uint, paddingBack uint, paddingStyle unsafe.Pointer, dataLayout unsafe.Pointer, weightsLayout unsafe.Pointer) GraphConvolution3DOpDescriptor {
	rv := objc.Send[GraphConvolution3DOpDescriptor](objc.ID(getGraphConvolution3DOpDescriptorClass().class), objc.Sel("descriptorWithStrideInX:strideInY:strideInZ:dilationRateInX:dilationRateInY:dilationRateInZ:groups:paddingLeft:paddingRight:paddingTop:paddingBottom:paddingFront:paddingBack:paddingStyle:dataLayout:weightsLayout:"), strideInX, strideInY, strideInZ, dilationRateInX, dilationRateInY, dilationRateInZ, groups, paddingLeft, paddingRight, paddingTop, paddingBottom, paddingFront, paddingBack, paddingStyle, dataLayout, weightsLayout)
	return rv
}


// Creates a convolution descriptor with given values for parameters.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphConvolution3DOpDescriptor/init(strideInX:strideInY:strideInZ:dilationRateInX:dilationRateInY:dilationRateInZ:groups:paddingLeft:paddingRight:paddingTop:paddingBottom:paddingFront:paddingBack:paddingStyle:dataLayout:weightsLayout:)
func (gc _GraphConvolution3DOpDescriptorClass) DescriptorWithStrideInXStrideInYStrideInZDilationRateInXDilationRateInYDilationRateInZGroupsPaddingLeftPaddingRightPaddingTopPaddingBottomPaddingFrontPaddingBackPaddingStyleDataLayoutWeightsLayout(strideInX uint, strideInY uint, strideInZ uint, dilationRateInX uint, dilationRateInY uint, dilationRateInZ uint, groups uint, paddingLeft uint, paddingRight uint, paddingTop uint, paddingBottom uint, paddingFront uint, paddingBack uint, paddingStyle unsafe.Pointer, dataLayout unsafe.Pointer, weightsLayout unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(gc.class), objc.Sel("descriptorWithStrideInX:strideInY:strideInZ:dilationRateInX:dilationRateInY:dilationRateInZ:groups:paddingLeft:paddingRight:paddingTop:paddingBottom:paddingFront:paddingBack:paddingStyle:dataLayout:weightsLayout:"), strideInX, strideInY, strideInZ, dilationRateInX, dilationRateInY, dilationRateInZ, groups, paddingLeft, paddingRight, paddingTop, paddingBottom, paddingFront, paddingBack, paddingStyle, dataLayout, weightsLayout)
	return rv
}

// The named layout of data in the source tensor.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphConvolution3DOpDescriptor/dataLayout
func (g_ GraphConvolution3DOpDescriptor) DataLayout() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("dataLayout"))
	return rv
}


// SetDataLayout sets the value of the dataLayout property.
// The named layout of data in the source tensor.

//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphConvolution3DOpDescriptor/dataLayout
func (g_ GraphConvolution3DOpDescriptor) SetDataLayout(value unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setDataLayout:"), value)
}

// The amount by which weights tensor expands in the -direction.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphConvolution3DOpDescriptor/dilationRateInX
func (g_ GraphConvolution3DOpDescriptor) DilationRateInX() uint {
	rv := objc.Send[uint](g_.ID, objc.Sel("dilationRateInX"))
	return rv
}


// SetDilationRateInX sets the value of the dilationRateInX property.
// The amount by which weights tensor expands in the -direction.

//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphConvolution3DOpDescriptor/dilationRateInX
func (g_ GraphConvolution3DOpDescriptor) SetDilationRateInX(value uint) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setDilationRateInX:"), value)
}

// The number of zeros added at the back of the source tensor.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphConvolution3DOpDescriptor/paddingBack
func (g_ GraphConvolution3DOpDescriptor) PaddingBack() uint {
	rv := objc.Send[uint](g_.ID, objc.Sel("paddingBack"))
	return rv
}


// SetPaddingBack sets the value of the paddingBack property.
// The number of zeros added at the back of the source tensor.

//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphConvolution3DOpDescriptor/paddingBack
func (g_ GraphConvolution3DOpDescriptor) SetPaddingBack(value uint) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setPaddingBack:"), value)
}

// The number of zeros added at the front of the source tensor.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphConvolution3DOpDescriptor/paddingFront
func (g_ GraphConvolution3DOpDescriptor) PaddingFront() uint {
	rv := objc.Send[uint](g_.ID, objc.Sel("paddingFront"))
	return rv
}


// SetPaddingFront sets the value of the paddingFront property.
// The number of zeros added at the front of the source tensor.

//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphConvolution3DOpDescriptor/paddingFront
func (g_ GraphConvolution3DOpDescriptor) SetPaddingFront(value uint) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setPaddingFront:"), value)
}

// The amount by which weights tensor expands in the
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphconvolution3dopdescriptor/dilationrateiny
func (g_ GraphConvolution3DOpDescriptor) DilationRateInY() int {
	rv := objc.Send[int](g_.ID, objc.Sel("dilationRateInY"))
	return rv
}


// SetDilationRateInY sets the value of the dilationRateInY property.
// The amount by which weights tensor expands in the

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphconvolution3dopdescriptor/dilationrateiny
func (g_ GraphConvolution3DOpDescriptor) SetDilationRateInY(value int) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setDilationRateInY:"), value)
}

// The amount by which weights tensor expands in the
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphconvolution3dopdescriptor/dilationrateinz
func (g_ GraphConvolution3DOpDescriptor) DilationRateInZ() int {
	rv := objc.Send[int](g_.ID, objc.Sel("dilationRateInZ"))
	return rv
}


// SetDilationRateInZ sets the value of the dilationRateInZ property.
// The amount by which weights tensor expands in the

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphconvolution3dopdescriptor/dilationrateinz
func (g_ GraphConvolution3DOpDescriptor) SetDilationRateInZ(value int) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setDilationRateInZ:"), value)
}

// The number of partitions of the input and output channels.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphconvolution3dopdescriptor/groups
func (g_ GraphConvolution3DOpDescriptor) Groups() int {
	rv := objc.Send[int](g_.ID, objc.Sel("groups"))
	return rv
}


// SetGroups sets the value of the groups property.
// The number of partitions of the input and output channels.

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphconvolution3dopdescriptor/groups
func (g_ GraphConvolution3DOpDescriptor) SetGroups(value int) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setGroups:"), value)
}

// The number of zeros added at the bottom of the source tensor.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphconvolution3dopdescriptor/paddingbottom
func (g_ GraphConvolution3DOpDescriptor) PaddingBottom() int {
	rv := objc.Send[int](g_.ID, objc.Sel("paddingBottom"))
	return rv
}


// SetPaddingBottom sets the value of the paddingBottom property.
// The number of zeros added at the bottom of the source tensor.

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphconvolution3dopdescriptor/paddingbottom
func (g_ GraphConvolution3DOpDescriptor) SetPaddingBottom(value int) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setPaddingBottom:"), value)
}

// The number of zeros added on the left side of the source tensor.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphconvolution3dopdescriptor/paddingleft
func (g_ GraphConvolution3DOpDescriptor) PaddingLeft() int {
	rv := objc.Send[int](g_.ID, objc.Sel("paddingLeft"))
	return rv
}


// SetPaddingLeft sets the value of the paddingLeft property.
// The number of zeros added on the left side of the source tensor.

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphconvolution3dopdescriptor/paddingleft
func (g_ GraphConvolution3DOpDescriptor) SetPaddingLeft(value int) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setPaddingLeft:"), value)
}

// The number of zeros added on the right side of the source tensor.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphconvolution3dopdescriptor/paddingright
func (g_ GraphConvolution3DOpDescriptor) PaddingRight() int {
	rv := objc.Send[int](g_.ID, objc.Sel("paddingRight"))
	return rv
}


// SetPaddingRight sets the value of the paddingRight property.
// The number of zeros added on the right side of the source tensor.

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphconvolution3dopdescriptor/paddingright
func (g_ GraphConvolution3DOpDescriptor) SetPaddingRight(value int) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setPaddingRight:"), value)
}

// The type of padding that is applied to the source tensor.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphconvolution3dopdescriptor/paddingstyle
func (g_ GraphConvolution3DOpDescriptor) PaddingStyle() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("paddingStyle"))
	return rv
}


// SetPaddingStyle sets the value of the paddingStyle property.
// The type of padding that is applied to the source tensor.

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphconvolution3dopdescriptor/paddingstyle
func (g_ GraphConvolution3DOpDescriptor) SetPaddingStyle(value unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setPaddingStyle:"), value)
}

// The number of zeros added at the top of the source tensor.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphconvolution3dopdescriptor/paddingtop
func (g_ GraphConvolution3DOpDescriptor) PaddingTop() int {
	rv := objc.Send[int](g_.ID, objc.Sel("paddingTop"))
	return rv
}


// SetPaddingTop sets the value of the paddingTop property.
// The number of zeros added at the top of the source tensor.

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphconvolution3dopdescriptor/paddingtop
func (g_ GraphConvolution3DOpDescriptor) SetPaddingTop(value int) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setPaddingTop:"), value)
}

// The scale that maps
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphconvolution3dopdescriptor/strideinx
func (g_ GraphConvolution3DOpDescriptor) StrideInX() int {
	rv := objc.Send[int](g_.ID, objc.Sel("strideInX"))
	return rv
}


// SetStrideInX sets the value of the strideInX property.
// The scale that maps

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphconvolution3dopdescriptor/strideinx
func (g_ GraphConvolution3DOpDescriptor) SetStrideInX(value int) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setStrideInX:"), value)
}

// The scale that maps
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphconvolution3dopdescriptor/strideiny
func (g_ GraphConvolution3DOpDescriptor) StrideInY() int {
	rv := objc.Send[int](g_.ID, objc.Sel("strideInY"))
	return rv
}


// SetStrideInY sets the value of the strideInY property.
// The scale that maps

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphconvolution3dopdescriptor/strideiny
func (g_ GraphConvolution3DOpDescriptor) SetStrideInY(value int) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setStrideInY:"), value)
}

// The scale that maps
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphconvolution3dopdescriptor/strideinz
func (g_ GraphConvolution3DOpDescriptor) StrideInZ() int {
	rv := objc.Send[int](g_.ID, objc.Sel("strideInZ"))
	return rv
}


// SetStrideInZ sets the value of the strideInZ property.
// The scale that maps

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphconvolution3dopdescriptor/strideinz
func (g_ GraphConvolution3DOpDescriptor) SetStrideInZ(value int) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setStrideInZ:"), value)
}

// The named layout of data in the weights tensor.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphconvolution3dopdescriptor/weightslayout
func (g_ GraphConvolution3DOpDescriptor) WeightsLayout() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("weightsLayout"))
	return rv
}


// SetWeightsLayout sets the value of the weightsLayout property.
// The named layout of data in the weights tensor.

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphconvolution3dopdescriptor/weightslayout
func (g_ GraphConvolution3DOpDescriptor) SetWeightsLayout(value unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setWeightsLayout:"), value)
}


