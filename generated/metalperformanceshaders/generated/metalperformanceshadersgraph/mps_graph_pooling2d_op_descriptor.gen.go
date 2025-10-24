// Code generated from Apple documentation for MetalPerformanceShadersGraph. DO NOT EDIT.

package metalperformanceshadersgraph

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class MPSGraphPooling2DOpDescriptor */


/* debug [class_header]: Header for MPSGraphPooling2DOpDescriptor */
// The class instance for the [GraphPooling2DOpDescriptor] class.
var (
	GraphPooling2DOpDescriptorClass     _GraphPooling2DOpDescriptorClass
	GraphPooling2DOpDescriptorClassOnce sync.Once
)

func getGraphPooling2DOpDescriptorClass() _GraphPooling2DOpDescriptorClass {
	GraphPooling2DOpDescriptorClassOnce.Do(func() {
		GraphPooling2DOpDescriptorClass = _GraphPooling2DOpDescriptorClass{objc.GetClass("MPSGraphPooling2DOpDescriptor")}
	})
	return GraphPooling2DOpDescriptorClass
}

type _GraphPooling2DOpDescriptorClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for GraphPooling2DOpDescriptor */
// An interface definition for the [GraphPooling2DOpDescriptor] class.
type IGraphPooling2DOpDescriptor interface {
	IGraphObject
	
/* debug [class_interface_properties]: Properties for GraphPooling2DOpDescriptor */
	// properties:
	CeilMode() bool
	SetCeilMode(value bool)
	DataLayout() GraphTensorNamedDataLayout
	SetDataLayout(value GraphTensorNamedDataLayout)
	DilationRateInX() uint
	SetDilationRateInX(value uint)
	DilationRateInY() uint
	SetDilationRateInY(value uint)
	IncludeZeroPadToAverage() bool
	SetIncludeZeroPadToAverage(value bool)
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
	PaddingStyle() GraphPaddingStyle
	SetPaddingStyle(value GraphPaddingStyle)
	PaddingTop() uint
	SetPaddingTop(value uint)
	ReturnIndicesDataType() DataType /* not a class type */
	SetReturnIndicesDataType(value DataType /* not a class type */)
	ReturnIndicesMode() GraphPoolingReturnIndicesMode
	SetReturnIndicesMode(value GraphPoolingReturnIndicesMode)
	StrideInX() uint
	SetStrideInX(value uint)
	StrideInY() uint
	SetStrideInY(value uint)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for GraphPooling2DOpDescriptor */
	// methods:
	SetExplicitPaddingWithPaddingLeftPaddingRightPaddingTopPaddingBottom(paddingLeft uint, paddingRight uint, paddingTop uint, paddingBottom uint)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for GraphPooling2DOpDescriptor */
// Alloc allocates a new instance without initialization.
func (gc _GraphPooling2DOpDescriptorClass) Alloc() GraphPooling2DOpDescriptor {
	rv := objc.Send[GraphPooling2DOpDescriptor](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (gc _GraphPooling2DOpDescriptorClass) New() GraphPooling2DOpDescriptor {
	rv := objc.Send[GraphPooling2DOpDescriptor](objc.ID(gc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (g_ GraphPooling2DOpDescriptor) Init() GraphPooling2DOpDescriptor {
	rv := objc.Send[GraphPooling2DOpDescriptor](g_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g_ GraphPooling2DOpDescriptor) Autorelease() GraphPooling2DOpDescriptor {
	rv := objc.Send[GraphPooling2DOpDescriptor](g_.ID, objc.Sel("autorelease"))
	return rv
}

// NewGraphPooling2DOpDescriptor creates a new GraphPooling2DOpDescriptor instance.
func NewGraphPooling2DOpDescriptor() GraphPooling2DOpDescriptor {
	return getGraphPooling2DOpDescriptorClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for GraphPooling2DOpDescriptor */
// The class that defines the parameters for a 2D pooling operation.
//
// Use this descriptor with the following methods:


// The class that defines the parameters for a 2D pooling operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphPooling2DOpDescriptor
type GraphPooling2DOpDescriptor struct {
	GraphObject
}

// GraphPooling2DOpDescriptorFrom constructs a [GraphPooling2DOpDescriptor] from an unsafe.Pointer.
//
// The class that defines the parameters for a 2D pooling operation.
func GraphPooling2DOpDescriptorFrom(ptr unsafe.Pointer) GraphPooling2DOpDescriptor {
	return GraphPooling2DOpDescriptor{
		GraphObject: GraphObjectFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for GraphPooling2DOpDescriptor */

// Creates a 2D pooling descriptor with given values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphPooling2DOpDescriptor/init(kernelWidth:kernelHeight:strideInX:strideInY:dilationRateInX:dilationRateInY:paddingLeft:paddingRight:paddingTop:paddingBottom:paddingStyle:dataLayout:)
func NewGraphPooling2DOpDescriptorWithKernelWidthKernelHeightStrideInXStrideInYDilationRateInXDilationRateInYPaddingLeftPaddingRightPaddingTopPaddingBottomPaddingStyleDataLayout(kernelWidth uint, kernelHeight uint, strideInX uint, strideInY uint, dilationRateInX uint, dilationRateInY uint, paddingLeft uint, paddingRight uint, paddingTop uint, paddingBottom uint, paddingStyle GraphPaddingStyle, dataLayout GraphTensorNamedDataLayout) GraphPooling2DOpDescriptor {
	rv := objc.Send[GraphPooling2DOpDescriptor](objc.ID(getGraphPooling2DOpDescriptorClass().class), objc.Sel("descriptorWithKernelWidth:kernelHeight:strideInX:strideInY:dilationRateInX:dilationRateInY:paddingLeft:paddingRight:paddingTop:paddingBottom:paddingStyle:dataLayout:"), kernelWidth, kernelHeight, strideInX, strideInY, dilationRateInX, dilationRateInY, paddingLeft, paddingRight, paddingTop, paddingBottom, paddingStyle, dataLayout)
	return rv
}/* debug [class_init_methods/constructor]: NewGraphPooling2DOpDescriptorWithKernelWidthKernelHeightStrideInXStrideInYDilationRateInXDilationRateInYPaddingLeftPaddingRightPaddingTopPaddingBottomPaddingStyleDataLayout */


// Creates a 2D pooling descriptor with given values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphPooling2DOpDescriptor/init(kernelWidth:kernelHeight:strideInX:strideInY:paddingStyle:dataLayout:)
func NewGraphPooling2DOpDescriptorWithKernelWidthKernelHeightStrideInXStrideInYPaddingStyleDataLayout(kernelWidth uint, kernelHeight uint, strideInX uint, strideInY uint, paddingStyle GraphPaddingStyle, dataLayout GraphTensorNamedDataLayout) GraphPooling2DOpDescriptor {
	rv := objc.Send[GraphPooling2DOpDescriptor](objc.ID(getGraphPooling2DOpDescriptorClass().class), objc.Sel("descriptorWithKernelWidth:kernelHeight:strideInX:strideInY:paddingStyle:dataLayout:"), kernelWidth, kernelHeight, strideInX, strideInY, paddingStyle, dataLayout)
	return rv
}/* debug [class_init_methods/constructor]: NewGraphPooling2DOpDescriptorWithKernelWidthKernelHeightStrideInXStrideInYPaddingStyleDataLayout */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for GraphPooling2DOpDescriptor */

// Creates a 2D pooling descriptor with given values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphPooling2DOpDescriptor/init(kernelWidth:kernelHeight:strideInX:strideInY:dilationRateInX:dilationRateInY:paddingLeft:paddingRight:paddingTop:paddingBottom:paddingStyle:dataLayout:)
func (gc _GraphPooling2DOpDescriptorClass) DescriptorWithKernelWidthKernelHeightStrideInXStrideInYDilationRateInXDilationRateInYPaddingLeftPaddingRightPaddingTopPaddingBottomPaddingStyleDataLayout(kernelWidth uint, kernelHeight uint, strideInX uint, strideInY uint, dilationRateInX uint, dilationRateInY uint, paddingLeft uint, paddingRight uint, paddingTop uint, paddingBottom uint, paddingStyle GraphPaddingStyle, dataLayout GraphTensorNamedDataLayout) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(gc.class), objc.Sel("descriptorWithKernelWidth:kernelHeight:strideInX:strideInY:dilationRateInX:dilationRateInY:paddingLeft:paddingRight:paddingTop:paddingBottom:paddingStyle:dataLayout:"), kernelWidth, kernelHeight, strideInX, strideInY, dilationRateInX, dilationRateInY, paddingLeft, paddingRight, paddingTop, paddingBottom, paddingStyle, dataLayout)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=DescriptorWithKernelWidthKernelHeightStrideInXStrideInYDilationRateInXDilationRateInYPaddingLeftPaddingRightPaddingTopPaddingBottomPaddingStyleDataLayout) */


// Creates a 2D pooling descriptor with given values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphPooling2DOpDescriptor/init(kernelWidth:kernelHeight:strideInX:strideInY:paddingStyle:dataLayout:)
func (gc _GraphPooling2DOpDescriptorClass) DescriptorWithKernelWidthKernelHeightStrideInXStrideInYPaddingStyleDataLayout(kernelWidth uint, kernelHeight uint, strideInX uint, strideInY uint, paddingStyle GraphPaddingStyle, dataLayout GraphTensorNamedDataLayout) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(gc.class), objc.Sel("descriptorWithKernelWidth:kernelHeight:strideInX:strideInY:paddingStyle:dataLayout:"), kernelWidth, kernelHeight, strideInX, strideInY, paddingStyle, dataLayout)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=DescriptorWithKernelWidthKernelHeightStrideInXStrideInYPaddingStyleDataLayout) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for GraphPooling2DOpDescriptor */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for GraphPooling2DOpDescriptor */

// Sets the explicit padding values and sets padding style to explicit.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphPooling2DOpDescriptor/setExplicitPaddingWithPaddingLeft(_:paddingRight:paddingTop:paddingBottom:)
func (g_ GraphPooling2DOpDescriptor) SetExplicitPaddingWithPaddingLeftPaddingRightPaddingTopPaddingBottom(paddingLeft uint, paddingRight uint, paddingTop uint, paddingBottom uint) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setExplicitPaddingWithPaddingLeft:paddingRight:paddingTop:paddingBottom:"), paddingLeft, paddingRight, paddingTop, paddingBottom)
}/* debug [instance_methods/method]: SetExplicitPaddingWithPaddingLeftPaddingRightPaddingTopPaddingBottom */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for GraphPooling2DOpDescriptor */

// Affects how the graph computes the output size.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphPooling2DOpDescriptor/ceilMode
func (g_ GraphPooling2DOpDescriptor) CeilMode() bool {
	rv := objc.Send[bool](g_.ID, objc.Sel("ceilMode"))
	return rv
}/* debug [instance_properties/getter]: ceilMode */


// Affects how the graph computes the output size.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphPooling2DOpDescriptor/ceilMode
func (g_ GraphPooling2DOpDescriptor) SetCeilMode(value bool) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setCeilMode:"), value)
}/* debug [instance_properties/setter]: ceilMode */


// Defines the data layout of the input data in the forward pass. See: .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphPooling2DOpDescriptor/dataLayout
func (g_ GraphPooling2DOpDescriptor) DataLayout() GraphTensorNamedDataLayout {
	rv := objc.Send[GraphTensorNamedDataLayout](g_.ID, objc.Sel("dataLayout"))
	return rv
}/* debug [instance_properties/getter]: dataLayout */


// Defines the data layout of the input data in the forward pass. See: .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphPooling2DOpDescriptor/dataLayout
func (g_ GraphPooling2DOpDescriptor) SetDataLayout(value GraphTensorNamedDataLayout) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setDataLayout:"), value)
}/* debug [instance_properties/setter]: dataLayout */


// Defines the dilation rate for the width dimension.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphPooling2DOpDescriptor/dilationRateInX
func (g_ GraphPooling2DOpDescriptor) DilationRateInX() uint {
	rv := objc.Send[uint](g_.ID, objc.Sel("dilationRateInX"))
	return rv
}/* debug [instance_properties/getter]: dilationRateInX */


// Defines the dilation rate for the width dimension.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphPooling2DOpDescriptor/dilationRateInX
func (g_ GraphPooling2DOpDescriptor) SetDilationRateInX(value uint) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setDilationRateInX:"), value)
}/* debug [instance_properties/setter]: dilationRateInX */


// Defines the dilation rate for the height dimension.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphPooling2DOpDescriptor/dilationRateInY
func (g_ GraphPooling2DOpDescriptor) DilationRateInY() uint {
	rv := objc.Send[uint](g_.ID, objc.Sel("dilationRateInY"))
	return rv
}/* debug [instance_properties/getter]: dilationRateInY */


// Defines the dilation rate for the height dimension.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphPooling2DOpDescriptor/dilationRateInY
func (g_ GraphPooling2DOpDescriptor) SetDilationRateInY(value uint) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setDilationRateInY:"), value)
}/* debug [instance_properties/setter]: dilationRateInY */


// Defines a mode for average pooling, where samples outside the input tensor count as zeroes in the average computation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphPooling2DOpDescriptor/includeZeroPadToAverage
func (g_ GraphPooling2DOpDescriptor) IncludeZeroPadToAverage() bool {
	rv := objc.Send[bool](g_.ID, objc.Sel("includeZeroPadToAverage"))
	return rv
}/* debug [instance_properties/getter]: includeZeroPadToAverage */


// Defines a mode for average pooling, where samples outside the input tensor count as zeroes in the average computation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphPooling2DOpDescriptor/includeZeroPadToAverage
func (g_ GraphPooling2DOpDescriptor) SetIncludeZeroPadToAverage(value bool) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setIncludeZeroPadToAverage:"), value)
}/* debug [instance_properties/setter]: includeZeroPadToAverage */


// Defines the pooling window size for the height dimension.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphPooling2DOpDescriptor/kernelHeight
func (g_ GraphPooling2DOpDescriptor) KernelHeight() uint {
	rv := objc.Send[uint](g_.ID, objc.Sel("kernelHeight"))
	return rv
}/* debug [instance_properties/getter]: kernelHeight */


// Defines the pooling window size for the height dimension.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphPooling2DOpDescriptor/kernelHeight
func (g_ GraphPooling2DOpDescriptor) SetKernelHeight(value uint) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setKernelHeight:"), value)
}/* debug [instance_properties/setter]: kernelHeight */


// Defines the pooling window size for the width dimension.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphPooling2DOpDescriptor/kernelWidth
func (g_ GraphPooling2DOpDescriptor) KernelWidth() uint {
	rv := objc.Send[uint](g_.ID, objc.Sel("kernelWidth"))
	return rv
}/* debug [instance_properties/getter]: kernelWidth */


// Defines the pooling window size for the width dimension.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphPooling2DOpDescriptor/kernelWidth
func (g_ GraphPooling2DOpDescriptor) SetKernelWidth(value uint) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setKernelWidth:"), value)
}/* debug [instance_properties/setter]: kernelWidth */


// Defines the explicit padding value for the height dimension to add after the data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphPooling2DOpDescriptor/paddingBottom
func (g_ GraphPooling2DOpDescriptor) PaddingBottom() uint {
	rv := objc.Send[uint](g_.ID, objc.Sel("paddingBottom"))
	return rv
}/* debug [instance_properties/getter]: paddingBottom */


// Defines the explicit padding value for the height dimension to add after the data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphPooling2DOpDescriptor/paddingBottom
func (g_ GraphPooling2DOpDescriptor) SetPaddingBottom(value uint) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setPaddingBottom:"), value)
}/* debug [instance_properties/setter]: paddingBottom */


// Defines the explicit padding value for the width dimension to add before the data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphPooling2DOpDescriptor/paddingLeft
func (g_ GraphPooling2DOpDescriptor) PaddingLeft() uint {
	rv := objc.Send[uint](g_.ID, objc.Sel("paddingLeft"))
	return rv
}/* debug [instance_properties/getter]: paddingLeft */


// Defines the explicit padding value for the width dimension to add before the data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphPooling2DOpDescriptor/paddingLeft
func (g_ GraphPooling2DOpDescriptor) SetPaddingLeft(value uint) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setPaddingLeft:"), value)
}/* debug [instance_properties/setter]: paddingLeft */


// Defines the explicit padding value for the width dimension to add after the data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphPooling2DOpDescriptor/paddingRight
func (g_ GraphPooling2DOpDescriptor) PaddingRight() uint {
	rv := objc.Send[uint](g_.ID, objc.Sel("paddingRight"))
	return rv
}/* debug [instance_properties/getter]: paddingRight */


// Defines the explicit padding value for the width dimension to add after the data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphPooling2DOpDescriptor/paddingRight
func (g_ GraphPooling2DOpDescriptor) SetPaddingRight(value uint) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setPaddingRight:"), value)
}/* debug [instance_properties/setter]: paddingRight */


// Defines what kind of padding graph applies to the operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphPooling2DOpDescriptor/paddingStyle
func (g_ GraphPooling2DOpDescriptor) PaddingStyle() GraphPaddingStyle {
	rv := objc.Send[GraphPaddingStyle](g_.ID, objc.Sel("paddingStyle"))
	return rv
}/* debug [instance_properties/getter]: paddingStyle */


// Defines what kind of padding graph applies to the operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphPooling2DOpDescriptor/paddingStyle
func (g_ GraphPooling2DOpDescriptor) SetPaddingStyle(value GraphPaddingStyle) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setPaddingStyle:"), value)
}/* debug [instance_properties/setter]: paddingStyle */


// Defines the explicit padding value for the height dimension to add before the data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphPooling2DOpDescriptor/paddingTop
func (g_ GraphPooling2DOpDescriptor) PaddingTop() uint {
	rv := objc.Send[uint](g_.ID, objc.Sel("paddingTop"))
	return rv
}/* debug [instance_properties/getter]: paddingTop */


// Defines the explicit padding value for the height dimension to add before the data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphPooling2DOpDescriptor/paddingTop
func (g_ GraphPooling2DOpDescriptor) SetPaddingTop(value uint) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setPaddingTop:"), value)
}/* debug [instance_properties/setter]: paddingTop */


// Defines the data type for returned indices. Use this in conjunction with API. Currently MPSGraph supports the following datatypes: . Default value: .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphPooling2DOpDescriptor/returnIndicesDataType
func (g_ GraphPooling2DOpDescriptor) ReturnIndicesDataType() DataType /* not a class type */ {
	rv := objc.Send[DataType](g_.ID, objc.Sel("returnIndicesDataType"))
	return rv
}/* debug [instance_properties/getter]: returnIndicesDataType */


// Defines the data type for returned indices. Use this in conjunction with API. Currently MPSGraph supports the following datatypes: . Default value: .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphPooling2DOpDescriptor/returnIndicesDataType
func (g_ GraphPooling2DOpDescriptor) SetReturnIndicesDataType(value DataType /* not a class type */) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setReturnIndicesDataType:"), value)
}/* debug [instance_properties/setter]: returnIndicesDataType */


// Defines the mode for returned indices of maximum values within each pooling window. Use this in conjunction with API. If then only the first result MPSGraph returns from will be valid and using the second result will assert. Default value: .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphPooling2DOpDescriptor/returnIndicesMode
func (g_ GraphPooling2DOpDescriptor) ReturnIndicesMode() GraphPoolingReturnIndicesMode {
	rv := objc.Send[GraphPoolingReturnIndicesMode](g_.ID, objc.Sel("returnIndicesMode"))
	return rv
}/* debug [instance_properties/getter]: returnIndicesMode */


// Defines the mode for returned indices of maximum values within each pooling window. Use this in conjunction with API. If then only the first result MPSGraph returns from will be valid and using the second result will assert. Default value: .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphPooling2DOpDescriptor/returnIndicesMode
func (g_ GraphPooling2DOpDescriptor) SetReturnIndicesMode(value GraphPoolingReturnIndicesMode) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setReturnIndicesMode:"), value)
}/* debug [instance_properties/setter]: returnIndicesMode */


// Defines the stride for the width dimension.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphPooling2DOpDescriptor/strideInX
func (g_ GraphPooling2DOpDescriptor) StrideInX() uint {
	rv := objc.Send[uint](g_.ID, objc.Sel("strideInX"))
	return rv
}/* debug [instance_properties/getter]: strideInX */


// Defines the stride for the width dimension.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphPooling2DOpDescriptor/strideInX
func (g_ GraphPooling2DOpDescriptor) SetStrideInX(value uint) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setStrideInX:"), value)
}/* debug [instance_properties/setter]: strideInX */


// Defines the stride for the height dimension.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphPooling2DOpDescriptor/strideInY
func (g_ GraphPooling2DOpDescriptor) StrideInY() uint {
	rv := objc.Send[uint](g_.ID, objc.Sel("strideInY"))
	return rv
}/* debug [instance_properties/getter]: strideInY */


// Defines the stride for the height dimension.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphPooling2DOpDescriptor/strideInY
func (g_ GraphPooling2DOpDescriptor) SetStrideInY(value uint) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setStrideInY:"), value)
}/* debug [instance_properties/setter]: strideInY */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSGraphPooling2DOpDescriptor */


