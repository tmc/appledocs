// Code generated from Apple documentation for MetalPerformanceShadersGraph. DO NOT EDIT.

package metalperformanceshadersgraph

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class MPSGraphImToColOpDescriptor */


/* debug [class_header]: Header for MPSGraphImToColOpDescriptor */
// The class instance for the [GraphImToColOpDescriptor] class.
var (
	GraphImToColOpDescriptorClass     _GraphImToColOpDescriptorClass
	GraphImToColOpDescriptorClassOnce sync.Once
)

func getGraphImToColOpDescriptorClass() _GraphImToColOpDescriptorClass {
	GraphImToColOpDescriptorClassOnce.Do(func() {
		GraphImToColOpDescriptorClass = _GraphImToColOpDescriptorClass{objc.GetClass("MPSGraphImToColOpDescriptor")}
	})
	return GraphImToColOpDescriptorClass
}

type _GraphImToColOpDescriptorClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for GraphImToColOpDescriptor */
// An interface definition for the [GraphImToColOpDescriptor] class.
type IGraphImToColOpDescriptor interface {
	IGraphObject
	
/* debug [class_interface_properties]: Properties for GraphImToColOpDescriptor */
	// properties:
	DataLayout() GraphTensorNamedDataLayout
	SetDataLayout(value GraphTensorNamedDataLayout)
	DilationRateInX() uint
	SetDilationRateInX(value uint)
	DilationRateInY() uint
	SetDilationRateInY(value uint)
	KernelHeight() uint
	SetKernelHeight(value uint)
	KernelWidth() uint
	SetKernelWidth(value uint)
	PaddingBottom() uint
	SetPaddingBottom(value uint)
	PaddingLeft() uint
	SetPaddingLeft(value uint)
	PaddingRight() uint
	SetPaddingRight(value uint)
	PaddingTop() uint
	SetPaddingTop(value uint)
	StrideInX() uint
	SetStrideInX(value uint)
	StrideInY() uint
	SetStrideInY(value uint)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for GraphImToColOpDescriptor */
	// methods:
	SetExplicitPaddingWithPaddingLeftPaddingRightPaddingTopPaddingBottom(paddingLeft uint, paddingRight uint, paddingTop uint, paddingBottom uint)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for GraphImToColOpDescriptor */
// Alloc allocates a new instance without initialization.
func (gc _GraphImToColOpDescriptorClass) Alloc() GraphImToColOpDescriptor {
	rv := objc.Send[GraphImToColOpDescriptor](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (gc _GraphImToColOpDescriptorClass) New() GraphImToColOpDescriptor {
	rv := objc.Send[GraphImToColOpDescriptor](objc.ID(gc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (g_ GraphImToColOpDescriptor) Init() GraphImToColOpDescriptor {
	rv := objc.Send[GraphImToColOpDescriptor](g_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g_ GraphImToColOpDescriptor) Autorelease() GraphImToColOpDescriptor {
	rv := objc.Send[GraphImToColOpDescriptor](g_.ID, objc.Sel("autorelease"))
	return rv
}

// NewGraphImToColOpDescriptor creates a new GraphImToColOpDescriptor instance.
func NewGraphImToColOpDescriptor() GraphImToColOpDescriptor {
	return getGraphImToColOpDescriptorClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for GraphImToColOpDescriptor */
// The class that defines the parameters for an image to column or column to image operation.
//
// Use this descriptor with the following methods:


// The class that defines the parameters for an image to column or column to image operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphImToColOpDescriptor
type GraphImToColOpDescriptor struct {
	GraphObject
}

// GraphImToColOpDescriptorFrom constructs a [GraphImToColOpDescriptor] from an unsafe.Pointer.
//
// The class that defines the parameters for an image to column or column to image operation.
func GraphImToColOpDescriptorFrom(ptr unsafe.Pointer) GraphImToColOpDescriptor {
	return GraphImToColOpDescriptor{
		GraphObject: GraphObjectFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for GraphImToColOpDescriptor */

// Creates column to image descriptor with given values for parameters.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphImToColOpDescriptor/init(kernelWidth:kernelHeight:strideInX:strideInY:dilationRateInX:dilationRateInY:dataLayout:)
func NewGraphImToColOpDescriptorWithKernelWidthKernelHeightStrideInXStrideInYDilationRateInXDilationRateInYDataLayout(kernelWidth uint, kernelHeight uint, strideInX uint, strideInY uint, dilationRateInX uint, dilationRateInY uint, dataLayout GraphTensorNamedDataLayout) GraphImToColOpDescriptor {
	rv := objc.Send[GraphImToColOpDescriptor](objc.ID(getGraphImToColOpDescriptorClass().class), objc.Sel("descriptorWithKernelWidth:kernelHeight:strideInX:strideInY:dilationRateInX:dilationRateInY:dataLayout:"), kernelWidth, kernelHeight, strideInX, strideInY, dilationRateInX, dilationRateInY, dataLayout)
	return rv
}/* debug [class_init_methods/constructor]: NewGraphImToColOpDescriptorWithKernelWidthKernelHeightStrideInXStrideInYDilationRateInXDilationRateInYDataLayout */


// Creates an image to column descriptor with given values for parameters.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphImToColOpDescriptor/init(kernelWidth:kernelHeight:strideInX:strideInY:dilationRateInX:dilationRateInY:paddingLeft:paddingRight:paddingTop:paddingBottom:dataLayout:)
func NewGraphImToColOpDescriptorWithKernelWidthKernelHeightStrideInXStrideInYDilationRateInXDilationRateInYPaddingLeftPaddingRightPaddingTopPaddingBottomDataLayout(kernelWidth uint, kernelHeight uint, strideInX uint, strideInY uint, dilationRateInX uint, dilationRateInY uint, paddingLeft uint, paddingRight uint, paddingTop uint, paddingBottom uint, dataLayout GraphTensorNamedDataLayout) GraphImToColOpDescriptor {
	rv := objc.Send[GraphImToColOpDescriptor](objc.ID(getGraphImToColOpDescriptorClass().class), objc.Sel("descriptorWithKernelWidth:kernelHeight:strideInX:strideInY:dilationRateInX:dilationRateInY:paddingLeft:paddingRight:paddingTop:paddingBottom:dataLayout:"), kernelWidth, kernelHeight, strideInX, strideInY, dilationRateInX, dilationRateInY, paddingLeft, paddingRight, paddingTop, paddingBottom, dataLayout)
	return rv
}/* debug [class_init_methods/constructor]: NewGraphImToColOpDescriptorWithKernelWidthKernelHeightStrideInXStrideInYDilationRateInXDilationRateInYPaddingLeftPaddingRightPaddingTopPaddingBottomDataLayout */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for GraphImToColOpDescriptor */

// Creates column to image descriptor with given values for parameters.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphImToColOpDescriptor/init(kernelWidth:kernelHeight:strideInX:strideInY:dilationRateInX:dilationRateInY:dataLayout:)
func (gc _GraphImToColOpDescriptorClass) DescriptorWithKernelWidthKernelHeightStrideInXStrideInYDilationRateInXDilationRateInYDataLayout(kernelWidth uint, kernelHeight uint, strideInX uint, strideInY uint, dilationRateInX uint, dilationRateInY uint, dataLayout GraphTensorNamedDataLayout) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(gc.class), objc.Sel("descriptorWithKernelWidth:kernelHeight:strideInX:strideInY:dilationRateInX:dilationRateInY:dataLayout:"), kernelWidth, kernelHeight, strideInX, strideInY, dilationRateInX, dilationRateInY, dataLayout)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=DescriptorWithKernelWidthKernelHeightStrideInXStrideInYDilationRateInXDilationRateInYDataLayout) */


// Creates an image to column descriptor with given values for parameters.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphImToColOpDescriptor/init(kernelWidth:kernelHeight:strideInX:strideInY:dilationRateInX:dilationRateInY:paddingLeft:paddingRight:paddingTop:paddingBottom:dataLayout:)
func (gc _GraphImToColOpDescriptorClass) DescriptorWithKernelWidthKernelHeightStrideInXStrideInYDilationRateInXDilationRateInYPaddingLeftPaddingRightPaddingTopPaddingBottomDataLayout(kernelWidth uint, kernelHeight uint, strideInX uint, strideInY uint, dilationRateInX uint, dilationRateInY uint, paddingLeft uint, paddingRight uint, paddingTop uint, paddingBottom uint, dataLayout GraphTensorNamedDataLayout) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(gc.class), objc.Sel("descriptorWithKernelWidth:kernelHeight:strideInX:strideInY:dilationRateInX:dilationRateInY:paddingLeft:paddingRight:paddingTop:paddingBottom:dataLayout:"), kernelWidth, kernelHeight, strideInX, strideInY, dilationRateInX, dilationRateInY, paddingLeft, paddingRight, paddingTop, paddingBottom, dataLayout)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=DescriptorWithKernelWidthKernelHeightStrideInXStrideInYDilationRateInXDilationRateInYPaddingLeftPaddingRightPaddingTopPaddingBottomDataLayout) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for GraphImToColOpDescriptor */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for GraphImToColOpDescriptor */

// Sets the descriptor’s padding to the given values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphImToColOpDescriptor/setExplicitPaddingWithPaddingLeft(_:paddingRight:paddingTop:paddingBottom:)
func (g_ GraphImToColOpDescriptor) SetExplicitPaddingWithPaddingLeftPaddingRightPaddingTopPaddingBottom(paddingLeft uint, paddingRight uint, paddingTop uint, paddingBottom uint) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setExplicitPaddingWithPaddingLeft:paddingRight:paddingTop:paddingBottom:"), paddingLeft, paddingRight, paddingTop, paddingBottom)
}/* debug [instance_methods/method]: SetExplicitPaddingWithPaddingLeftPaddingRightPaddingTopPaddingBottom */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for GraphImToColOpDescriptor */

// The property that defines the layout of source or output tensor. e.g. for layout
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphImToColOpDescriptor/dataLayout
func (g_ GraphImToColOpDescriptor) DataLayout() GraphTensorNamedDataLayout {
	rv := objc.Send[GraphTensorNamedDataLayout](g_.ID, objc.Sel("dataLayout"))
	return rv
}/* debug [instance_properties/getter]: dataLayout */


// The property that defines the layout of source or output tensor. e.g. for layout
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphImToColOpDescriptor/dataLayout
func (g_ GraphImToColOpDescriptor) SetDataLayout(value GraphTensorNamedDataLayout) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setDataLayout:"), value)
}/* debug [instance_properties/setter]: dataLayout */


// The property that defines the dilation in width dimension.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphImToColOpDescriptor/dilationRateInX
func (g_ GraphImToColOpDescriptor) DilationRateInX() uint {
	rv := objc.Send[uint](g_.ID, objc.Sel("dilationRateInX"))
	return rv
}/* debug [instance_properties/getter]: dilationRateInX */


// The property that defines the dilation in width dimension.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphImToColOpDescriptor/dilationRateInX
func (g_ GraphImToColOpDescriptor) SetDilationRateInX(value uint) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setDilationRateInX:"), value)
}/* debug [instance_properties/setter]: dilationRateInX */


// The property that defines the dilation in height dimension.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphImToColOpDescriptor/dilationRateInY
func (g_ GraphImToColOpDescriptor) DilationRateInY() uint {
	rv := objc.Send[uint](g_.ID, objc.Sel("dilationRateInY"))
	return rv
}/* debug [instance_properties/getter]: dilationRateInY */


// The property that defines the dilation in height dimension.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphImToColOpDescriptor/dilationRateInY
func (g_ GraphImToColOpDescriptor) SetDilationRateInY(value uint) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setDilationRateInY:"), value)
}/* debug [instance_properties/setter]: dilationRateInY */


// The property that defines the kernel size in height dimension.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphImToColOpDescriptor/kernelHeight
func (g_ GraphImToColOpDescriptor) KernelHeight() uint {
	rv := objc.Send[uint](g_.ID, objc.Sel("kernelHeight"))
	return rv
}/* debug [instance_properties/getter]: kernelHeight */


// The property that defines the kernel size in height dimension.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphImToColOpDescriptor/kernelHeight
func (g_ GraphImToColOpDescriptor) SetKernelHeight(value uint) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setKernelHeight:"), value)
}/* debug [instance_properties/setter]: kernelHeight */


// The property that defines the kernel size in width dimension.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphImToColOpDescriptor/kernelWidth
func (g_ GraphImToColOpDescriptor) KernelWidth() uint {
	rv := objc.Send[uint](g_.ID, objc.Sel("kernelWidth"))
	return rv
}/* debug [instance_properties/getter]: kernelWidth */


// The property that defines the kernel size in width dimension.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphImToColOpDescriptor/kernelWidth
func (g_ GraphImToColOpDescriptor) SetKernelWidth(value uint) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setKernelWidth:"), value)
}/* debug [instance_properties/setter]: kernelWidth */


// The property that defines the padding in height dimension at the bottom.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphImToColOpDescriptor/paddingBottom
func (g_ GraphImToColOpDescriptor) PaddingBottom() uint {
	rv := objc.Send[uint](g_.ID, objc.Sel("paddingBottom"))
	return rv
}/* debug [instance_properties/getter]: paddingBottom */


// The property that defines the padding in height dimension at the bottom.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphImToColOpDescriptor/paddingBottom
func (g_ GraphImToColOpDescriptor) SetPaddingBottom(value uint) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setPaddingBottom:"), value)
}/* debug [instance_properties/setter]: paddingBottom */


// The property that defines the padding in width dimension on the left side.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphImToColOpDescriptor/paddingLeft
func (g_ GraphImToColOpDescriptor) PaddingLeft() uint {
	rv := objc.Send[uint](g_.ID, objc.Sel("paddingLeft"))
	return rv
}/* debug [instance_properties/getter]: paddingLeft */


// The property that defines the padding in width dimension on the left side.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphImToColOpDescriptor/paddingLeft
func (g_ GraphImToColOpDescriptor) SetPaddingLeft(value uint) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setPaddingLeft:"), value)
}/* debug [instance_properties/setter]: paddingLeft */


// The property that defines the padding in width dimension on the right side.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphImToColOpDescriptor/paddingRight
func (g_ GraphImToColOpDescriptor) PaddingRight() uint {
	rv := objc.Send[uint](g_.ID, objc.Sel("paddingRight"))
	return rv
}/* debug [instance_properties/getter]: paddingRight */


// The property that defines the padding in width dimension on the right side.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphImToColOpDescriptor/paddingRight
func (g_ GraphImToColOpDescriptor) SetPaddingRight(value uint) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setPaddingRight:"), value)
}/* debug [instance_properties/setter]: paddingRight */


// The property that defines the padding in height dimension at the top.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphImToColOpDescriptor/paddingTop
func (g_ GraphImToColOpDescriptor) PaddingTop() uint {
	rv := objc.Send[uint](g_.ID, objc.Sel("paddingTop"))
	return rv
}/* debug [instance_properties/getter]: paddingTop */


// The property that defines the padding in height dimension at the top.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphImToColOpDescriptor/paddingTop
func (g_ GraphImToColOpDescriptor) SetPaddingTop(value uint) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setPaddingTop:"), value)
}/* debug [instance_properties/setter]: paddingTop */


// The property that defines the stride in width dimension.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphImToColOpDescriptor/strideInX
func (g_ GraphImToColOpDescriptor) StrideInX() uint {
	rv := objc.Send[uint](g_.ID, objc.Sel("strideInX"))
	return rv
}/* debug [instance_properties/getter]: strideInX */


// The property that defines the stride in width dimension.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphImToColOpDescriptor/strideInX
func (g_ GraphImToColOpDescriptor) SetStrideInX(value uint) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setStrideInX:"), value)
}/* debug [instance_properties/setter]: strideInX */


// The property that defines the stride in height dimension.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphImToColOpDescriptor/strideInY
func (g_ GraphImToColOpDescriptor) StrideInY() uint {
	rv := objc.Send[uint](g_.ID, objc.Sel("strideInY"))
	return rv
}/* debug [instance_properties/getter]: strideInY */


// The property that defines the stride in height dimension.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphImToColOpDescriptor/strideInY
func (g_ GraphImToColOpDescriptor) SetStrideInY(value uint) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setStrideInY:"), value)
}/* debug [instance_properties/setter]: strideInY */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSGraphImToColOpDescriptor */


