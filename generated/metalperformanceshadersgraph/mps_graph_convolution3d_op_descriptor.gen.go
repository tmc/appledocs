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

