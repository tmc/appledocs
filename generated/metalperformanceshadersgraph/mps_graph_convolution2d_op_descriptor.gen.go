// Code generated from Apple documentation for MetalPerformanceShadersGraph. DO NOT EDIT.

package metalperformanceshadersgraph

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class MPSGraphConvolution2DOpDescriptor */


/* debug [class_header]: Header for MPSGraphConvolution2DOpDescriptor */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for GraphConvolution2DOpDescriptor */
// An interface definition for the [GraphConvolution2DOpDescriptor] class.
type IGraphConvolution2DOpDescriptor interface {
	IGraphObject
	
/* debug [class_interface_properties]: Properties for GraphConvolution2DOpDescriptor */
	// properties:
	DataLayout() GraphTensorNamedDataLayout
	SetDataLayout(value GraphTensorNamedDataLayout)
	DilationRateInX() uint
	SetDilationRateInX(value uint)
	DilationRateInY() uint
	SetDilationRateInY(value uint)
	Groups() uint
	SetGroups(value uint)
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
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for GraphConvolution2DOpDescriptor */
	// methods:
	SetExplicitPaddingWithPaddingLeftPaddingRightPaddingTopPaddingBottom(paddingLeft uint, paddingRight uint, paddingTop uint, paddingBottom uint)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for GraphConvolution2DOpDescriptor */
// Alloc allocates a new instance without initialization.
func (gc _GraphConvolution2DOpDescriptorClass) Alloc() GraphConvolution2DOpDescriptor {
	rv := objc.Send[GraphConvolution2DOpDescriptor](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for GraphConvolution2DOpDescriptor */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for GraphConvolution2DOpDescriptor */

// Creates a convolution descriptor with given values for parameters.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphConvolution2DOpDescriptor/init(strideInX:strideInY:dilationRateInX:dilationRateInY:groups:paddingLeft:paddingRight:paddingTop:paddingBottom:paddingStyle:dataLayout:weightsLayout:)
func NewGraphConvolution2DOpDescriptorWithStrideInXStrideInYDilationRateInXDilationRateInYGroupsPaddingLeftPaddingRightPaddingTopPaddingBottomPaddingStyleDataLayoutWeightsLayout(strideInX uint, strideInY uint, dilationRateInX uint, dilationRateInY uint, groups uint, paddingLeft uint, paddingRight uint, paddingTop uint, paddingBottom uint, paddingStyle GraphPaddingStyle, dataLayout GraphTensorNamedDataLayout, weightsLayout GraphTensorNamedDataLayout) GraphConvolution2DOpDescriptor {
	rv := objc.Send[GraphConvolution2DOpDescriptor](objc.ID(getGraphConvolution2DOpDescriptorClass().class), objc.Sel("descriptorWithStrideInX:strideInY:dilationRateInX:dilationRateInY:groups:paddingLeft:paddingRight:paddingTop:paddingBottom:paddingStyle:dataLayout:weightsLayout:"), strideInX, strideInY, dilationRateInX, dilationRateInY, groups, paddingLeft, paddingRight, paddingTop, paddingBottom, paddingStyle, dataLayout, weightsLayout)
	return rv
}/* debug [class_init_methods/constructor]: NewGraphConvolution2DOpDescriptorWithStrideInXStrideInYDilationRateInXDilationRateInYGroupsPaddingLeftPaddingRightPaddingTopPaddingBottomPaddingStyleDataLayoutWeightsLayout */


// Creates a convolution descriptor with given values for parameters.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphConvolution2DOpDescriptor/init(strideInX:strideInY:dilationRateInX:dilationRateInY:groups:paddingStyle:dataLayout:weightsLayout:)
func NewGraphConvolution2DOpDescriptorWithStrideInXStrideInYDilationRateInXDilationRateInYGroupsPaddingStyleDataLayoutWeightsLayout(strideInX uint, strideInY uint, dilationRateInX uint, dilationRateInY uint, groups uint, paddingStyle GraphPaddingStyle, dataLayout GraphTensorNamedDataLayout, weightsLayout GraphTensorNamedDataLayout) GraphConvolution2DOpDescriptor {
	rv := objc.Send[GraphConvolution2DOpDescriptor](objc.ID(getGraphConvolution2DOpDescriptorClass().class), objc.Sel("descriptorWithStrideInX:strideInY:dilationRateInX:dilationRateInY:groups:paddingStyle:dataLayout:weightsLayout:"), strideInX, strideInY, dilationRateInX, dilationRateInY, groups, paddingStyle, dataLayout, weightsLayout)
	return rv
}/* debug [class_init_methods/constructor]: NewGraphConvolution2DOpDescriptorWithStrideInXStrideInYDilationRateInXDilationRateInYGroupsPaddingStyleDataLayoutWeightsLayout */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for GraphConvolution2DOpDescriptor */

// Creates a convolution descriptor with given values for parameters.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphConvolution2DOpDescriptor/init(strideInX:strideInY:dilationRateInX:dilationRateInY:groups:paddingLeft:paddingRight:paddingTop:paddingBottom:paddingStyle:dataLayout:weightsLayout:)
func (gc _GraphConvolution2DOpDescriptorClass) DescriptorWithStrideInXStrideInYDilationRateInXDilationRateInYGroupsPaddingLeftPaddingRightPaddingTopPaddingBottomPaddingStyleDataLayoutWeightsLayout(strideInX uint, strideInY uint, dilationRateInX uint, dilationRateInY uint, groups uint, paddingLeft uint, paddingRight uint, paddingTop uint, paddingBottom uint, paddingStyle GraphPaddingStyle, dataLayout GraphTensorNamedDataLayout, weightsLayout GraphTensorNamedDataLayout) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(gc.class), objc.Sel("descriptorWithStrideInX:strideInY:dilationRateInX:dilationRateInY:groups:paddingLeft:paddingRight:paddingTop:paddingBottom:paddingStyle:dataLayout:weightsLayout:"), strideInX, strideInY, dilationRateInX, dilationRateInY, groups, paddingLeft, paddingRight, paddingTop, paddingBottom, paddingStyle, dataLayout, weightsLayout)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=DescriptorWithStrideInXStrideInYDilationRateInXDilationRateInYGroupsPaddingLeftPaddingRightPaddingTopPaddingBottomPaddingStyleDataLayoutWeightsLayout) */


// Creates a convolution descriptor with given values for parameters.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphConvolution2DOpDescriptor/init(strideInX:strideInY:dilationRateInX:dilationRateInY:groups:paddingStyle:dataLayout:weightsLayout:)
func (gc _GraphConvolution2DOpDescriptorClass) DescriptorWithStrideInXStrideInYDilationRateInXDilationRateInYGroupsPaddingStyleDataLayoutWeightsLayout(strideInX uint, strideInY uint, dilationRateInX uint, dilationRateInY uint, groups uint, paddingStyle GraphPaddingStyle, dataLayout GraphTensorNamedDataLayout, weightsLayout GraphTensorNamedDataLayout) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(gc.class), objc.Sel("descriptorWithStrideInX:strideInY:dilationRateInX:dilationRateInY:groups:paddingStyle:dataLayout:weightsLayout:"), strideInX, strideInY, dilationRateInX, dilationRateInY, groups, paddingStyle, dataLayout, weightsLayout)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=DescriptorWithStrideInXStrideInYDilationRateInXDilationRateInYGroupsPaddingStyleDataLayoutWeightsLayout) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for GraphConvolution2DOpDescriptor */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for GraphConvolution2DOpDescriptor */

// Sets the left, right, top, and bottom padding values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphConvolution2DOpDescriptor/setExplicitPaddingWithPaddingLeft(_:paddingRight:paddingTop:paddingBottom:)
func (g_ GraphConvolution2DOpDescriptor) SetExplicitPaddingWithPaddingLeftPaddingRightPaddingTopPaddingBottom(paddingLeft uint, paddingRight uint, paddingTop uint, paddingBottom uint) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setExplicitPaddingWithPaddingLeft:paddingRight:paddingTop:paddingBottom:"), paddingLeft, paddingRight, paddingTop, paddingBottom)
}/* debug [instance_methods/method]: SetExplicitPaddingWithPaddingLeftPaddingRightPaddingTopPaddingBottom */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for GraphConvolution2DOpDescriptor */

// The named layout of data in the source tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphConvolution2DOpDescriptor/dataLayout
func (g_ GraphConvolution2DOpDescriptor) DataLayout() GraphTensorNamedDataLayout {
	rv := objc.Send[GraphTensorNamedDataLayout](g_.ID, objc.Sel("dataLayout"))
	return rv
}/* debug [instance_properties/getter]: dataLayout */


// The named layout of data in the source tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphConvolution2DOpDescriptor/dataLayout
func (g_ GraphConvolution2DOpDescriptor) SetDataLayout(value GraphTensorNamedDataLayout) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setDataLayout:"), value)
}/* debug [instance_properties/setter]: dataLayout */


// The amount by which the weights tensor expands in the -direction.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphConvolution2DOpDescriptor/dilationRateInX
func (g_ GraphConvolution2DOpDescriptor) DilationRateInX() uint {
	rv := objc.Send[uint](g_.ID, objc.Sel("dilationRateInX"))
	return rv
}/* debug [instance_properties/getter]: dilationRateInX */


// The amount by which the weights tensor expands in the -direction.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphConvolution2DOpDescriptor/dilationRateInX
func (g_ GraphConvolution2DOpDescriptor) SetDilationRateInX(value uint) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setDilationRateInX:"), value)
}/* debug [instance_properties/setter]: dilationRateInX */


// The amount by which the weights tensor expands in the -direction.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphConvolution2DOpDescriptor/dilationRateInY
func (g_ GraphConvolution2DOpDescriptor) DilationRateInY() uint {
	rv := objc.Send[uint](g_.ID, objc.Sel("dilationRateInY"))
	return rv
}/* debug [instance_properties/getter]: dilationRateInY */


// The amount by which the weights tensor expands in the -direction.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphConvolution2DOpDescriptor/dilationRateInY
func (g_ GraphConvolution2DOpDescriptor) SetDilationRateInY(value uint) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setDilationRateInY:"), value)
}/* debug [instance_properties/setter]: dilationRateInY */


// The number of partitions of the input and output channels.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphConvolution2DOpDescriptor/groups
func (g_ GraphConvolution2DOpDescriptor) Groups() uint {
	rv := objc.Send[uint](g_.ID, objc.Sel("groups"))
	return rv
}/* debug [instance_properties/getter]: groups */


// The number of partitions of the input and output channels.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphConvolution2DOpDescriptor/groups
func (g_ GraphConvolution2DOpDescriptor) SetGroups(value uint) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setGroups:"), value)
}/* debug [instance_properties/setter]: groups */


// The number of zeros added at the bottom of the source tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphConvolution2DOpDescriptor/paddingBottom
func (g_ GraphConvolution2DOpDescriptor) PaddingBottom() uint {
	rv := objc.Send[uint](g_.ID, objc.Sel("paddingBottom"))
	return rv
}/* debug [instance_properties/getter]: paddingBottom */


// The number of zeros added at the bottom of the source tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphConvolution2DOpDescriptor/paddingBottom
func (g_ GraphConvolution2DOpDescriptor) SetPaddingBottom(value uint) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setPaddingBottom:"), value)
}/* debug [instance_properties/setter]: paddingBottom */


// The number of zeros added on the left side of the source tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphConvolution2DOpDescriptor/paddingLeft
func (g_ GraphConvolution2DOpDescriptor) PaddingLeft() uint {
	rv := objc.Send[uint](g_.ID, objc.Sel("paddingLeft"))
	return rv
}/* debug [instance_properties/getter]: paddingLeft */


// The number of zeros added on the left side of the source tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphConvolution2DOpDescriptor/paddingLeft
func (g_ GraphConvolution2DOpDescriptor) SetPaddingLeft(value uint) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setPaddingLeft:"), value)
}/* debug [instance_properties/setter]: paddingLeft */


// The number of zeros added on the right side of the source tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphConvolution2DOpDescriptor/paddingRight
func (g_ GraphConvolution2DOpDescriptor) PaddingRight() uint {
	rv := objc.Send[uint](g_.ID, objc.Sel("paddingRight"))
	return rv
}/* debug [instance_properties/getter]: paddingRight */


// The number of zeros added on the right side of the source tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphConvolution2DOpDescriptor/paddingRight
func (g_ GraphConvolution2DOpDescriptor) SetPaddingRight(value uint) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setPaddingRight:"), value)
}/* debug [instance_properties/setter]: paddingRight */


// The type of padding applied to the source tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphConvolution2DOpDescriptor/paddingStyle
func (g_ GraphConvolution2DOpDescriptor) PaddingStyle() GraphPaddingStyle {
	rv := objc.Send[GraphPaddingStyle](g_.ID, objc.Sel("paddingStyle"))
	return rv
}/* debug [instance_properties/getter]: paddingStyle */


// The type of padding applied to the source tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphConvolution2DOpDescriptor/paddingStyle
func (g_ GraphConvolution2DOpDescriptor) SetPaddingStyle(value GraphPaddingStyle) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setPaddingStyle:"), value)
}/* debug [instance_properties/setter]: paddingStyle */


// The number of zeros added at the top of the source tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphConvolution2DOpDescriptor/paddingTop
func (g_ GraphConvolution2DOpDescriptor) PaddingTop() uint {
	rv := objc.Send[uint](g_.ID, objc.Sel("paddingTop"))
	return rv
}/* debug [instance_properties/getter]: paddingTop */


// The number of zeros added at the top of the source tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphConvolution2DOpDescriptor/paddingTop
func (g_ GraphConvolution2DOpDescriptor) SetPaddingTop(value uint) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setPaddingTop:"), value)
}/* debug [instance_properties/setter]: paddingTop */


// The scale that maps -coordinate of the destination to -coordinate of the source.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphConvolution2DOpDescriptor/strideInX
func (g_ GraphConvolution2DOpDescriptor) StrideInX() uint {
	rv := objc.Send[uint](g_.ID, objc.Sel("strideInX"))
	return rv
}/* debug [instance_properties/getter]: strideInX */


// The scale that maps -coordinate of the destination to -coordinate of the source.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphConvolution2DOpDescriptor/strideInX
func (g_ GraphConvolution2DOpDescriptor) SetStrideInX(value uint) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setStrideInX:"), value)
}/* debug [instance_properties/setter]: strideInX */


// The scale that maps -coordinate of the destination to -coordinate of the source.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphConvolution2DOpDescriptor/strideInY
func (g_ GraphConvolution2DOpDescriptor) StrideInY() uint {
	rv := objc.Send[uint](g_.ID, objc.Sel("strideInY"))
	return rv
}/* debug [instance_properties/getter]: strideInY */


// The scale that maps -coordinate of the destination to -coordinate of the source.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphConvolution2DOpDescriptor/strideInY
func (g_ GraphConvolution2DOpDescriptor) SetStrideInY(value uint) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setStrideInY:"), value)
}/* debug [instance_properties/setter]: strideInY */


// The named layout of data in the weights tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphConvolution2DOpDescriptor/weightsLayout
func (g_ GraphConvolution2DOpDescriptor) WeightsLayout() GraphTensorNamedDataLayout {
	rv := objc.Send[GraphTensorNamedDataLayout](g_.ID, objc.Sel("weightsLayout"))
	return rv
}/* debug [instance_properties/getter]: weightsLayout */


// The named layout of data in the weights tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphConvolution2DOpDescriptor/weightsLayout
func (g_ GraphConvolution2DOpDescriptor) SetWeightsLayout(value GraphTensorNamedDataLayout) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setWeightsLayout:"), value)
}/* debug [instance_properties/setter]: weightsLayout */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSGraphConvolution2DOpDescriptor */


