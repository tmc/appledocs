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
}

// A class that describes the properties of a 2D-convolution operator.
//
// Use an instance of this class is to add a 2D-convolution operator with the desired properties to the graph.
//
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
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphConvolution2DOpDescriptor/init(strideInX:strideInY:dilationRateInX:dilationRateInY:groups:paddingLeft:paddingRight:paddingTop:paddingBottom:paddingStyle:dataLayout:weightsLayout:)
func NewGraphConvolution2DOpDescriptorWithStrideInXStrideInYDilationRateInXDilationRateInYGroupsPaddingLeftPaddingRightPaddingTopPaddingBottomPaddingStyleDataLayoutWeightsLayout(strideInX uint, strideInY uint, dilationRateInX uint, dilationRateInY uint, groups uint, paddingLeft uint, paddingRight uint, paddingTop uint, paddingBottom uint, paddingStyle unsafe.Pointer, dataLayout unsafe.Pointer, weightsLayout unsafe.Pointer) GraphConvolution2DOpDescriptor {
	rv := objc.Send[GraphConvolution2DOpDescriptor](objc.ID(getGraphConvolution2DOpDescriptorClass().class), objc.Sel("descriptorWithStrideInX:strideInY:dilationRateInX:dilationRateInY:groups:paddingLeft:paddingRight:paddingTop:paddingBottom:paddingStyle:dataLayout:weightsLayout:"), strideInX, strideInY, dilationRateInX, dilationRateInY, groups, paddingLeft, paddingRight, paddingTop, paddingBottom, paddingStyle, dataLayout, weightsLayout)
	return rv
}


// Creates a convolution descriptor with given values for parameters.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphConvolution2DOpDescriptor/init(strideInX:strideInY:dilationRateInX:dilationRateInY:groups:paddingLeft:paddingRight:paddingTop:paddingBottom:paddingStyle:dataLayout:weightsLayout:)
func (gc _GraphConvolution2DOpDescriptorClass) DescriptorWithStrideInXStrideInYDilationRateInXDilationRateInYGroupsPaddingLeftPaddingRightPaddingTopPaddingBottomPaddingStyleDataLayoutWeightsLayout(strideInX uint, strideInY uint, dilationRateInX uint, dilationRateInY uint, groups uint, paddingLeft uint, paddingRight uint, paddingTop uint, paddingBottom uint, paddingStyle unsafe.Pointer, dataLayout unsafe.Pointer, weightsLayout unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(gc.class), objc.Sel("descriptorWithStrideInX:strideInY:dilationRateInX:dilationRateInY:groups:paddingLeft:paddingRight:paddingTop:paddingBottom:paddingStyle:dataLayout:weightsLayout:"), strideInX, strideInY, dilationRateInX, dilationRateInY, groups, paddingLeft, paddingRight, paddingTop, paddingBottom, paddingStyle, dataLayout, weightsLayout)
	return rv
}

// The amount by which the weights tensor expands in the -direction.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphConvolution2DOpDescriptor/dilationRateInY
func (g_ GraphConvolution2DOpDescriptor) DilationRateInY() uint {
	rv := objc.Send[uint](g_.ID, objc.Sel("dilationRateInY"))
	return rv
}


// SetDilationRateInY sets the value of the dilationRateInY property.
// The amount by which the weights tensor expands in the -direction.

//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphConvolution2DOpDescriptor/dilationRateInY
func (g_ GraphConvolution2DOpDescriptor) SetDilationRateInY(value uint) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setDilationRateInY:"), value)
}

