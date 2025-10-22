// Code generated from Apple documentation for MetalPerformanceShadersGraph. DO NOT EDIT.

package metalperformanceshadersgraph

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

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

// An interface definition for the [GraphPooling2DOpDescriptor] class.
type IGraphPooling2DOpDescriptor interface {
	IGraphObject
	PaddingLeft() uint
	SetPaddingLeft(value uint)
	CeilMode() bool
	SetCeilMode(value bool)
	DataLayout() GraphTensorNamedDataLayout
	SetDataLayout(value IGraphTensorNamedDataLayout)
	DilationRateInX() int
	SetDilationRateInX(value int)
	DilationRateInY() int
	SetDilationRateInY(value int)
	IncludeZeroPadToAverage() bool
	SetIncludeZeroPadToAverage(value bool)
	KernelHeight() int
	SetKernelHeight(value int)
	KernelWidth() int
	SetKernelWidth(value int)
	PaddingBottom() int
	SetPaddingBottom(value int)
	PaddingRight() int
	SetPaddingRight(value int)
	PaddingStyle() GraphPaddingStyle
	SetPaddingStyle(value GraphPaddingStyle)
	PaddingTop() int
	SetPaddingTop(value int)
	ReturnIndicesDataType() unsafe.Pointer
	SetReturnIndicesDataType(value unsafe.Pointer)
	ReturnIndicesMode() GraphPoolingReturnIndicesMode
	SetReturnIndicesMode(value GraphPoolingReturnIndicesMode)
	StrideInX() int
	SetStrideInX(value int)
	StrideInY() int
	SetStrideInY(value int)
}

// The class that defines the parameters for a 2D pooling operation.
//
// Use this descriptor with the following methods:
//
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

// Alloc allocates a new instance without initialization.
func (gc _GraphPooling2DOpDescriptorClass) Alloc() GraphPooling2DOpDescriptor {
	rv := objc.Send[GraphPooling2DOpDescriptor](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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


// Defines the explicit padding value for the width dimension to add before the data.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphPooling2DOpDescriptor/paddingLeft
func (g_ GraphPooling2DOpDescriptor) PaddingLeft() uint {
	rv := objc.Send[uint](g_.ID, objc.Sel("paddingLeft"))
	return rv
}


// SetPaddingLeft sets the value of the paddingLeft property.
// Defines the explicit padding value for the width dimension to add before the data.

//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphPooling2DOpDescriptor/paddingLeft
func (g_ GraphPooling2DOpDescriptor) SetPaddingLeft(value uint) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setPaddingLeft:"), value)
}

// Affects how the graph computes the output size.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphpooling2dopdescriptor/ceilmode
func (g_ GraphPooling2DOpDescriptor) CeilMode() bool {
	rv := objc.Send[bool](g_.ID, objc.Sel("ceilMode"))
	return rv
}


// SetCeilMode sets the value of the ceilMode property.
// Affects how the graph computes the output size.

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphpooling2dopdescriptor/ceilmode
func (g_ GraphPooling2DOpDescriptor) SetCeilMode(value bool) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setCeilMode:"), value)
}

// Defines the data layout of the input data in the forward pass. See:
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphpooling2dopdescriptor/datalayout
func (g_ GraphPooling2DOpDescriptor) DataLayout() GraphTensorNamedDataLayout {
	rv := objc.Send[GraphTensorNamedDataLayout](g_.ID, objc.Sel("dataLayout"))
	return rv
}


// SetDataLayout sets the value of the dataLayout property.
// Defines the data layout of the input data in the forward pass. See:

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphpooling2dopdescriptor/datalayout
func (g_ GraphPooling2DOpDescriptor) SetDataLayout(value IGraphTensorNamedDataLayout) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setDataLayout:"), value)
}

// Defines the dilation rate for the width dimension.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphpooling2dopdescriptor/dilationrateinx
func (g_ GraphPooling2DOpDescriptor) DilationRateInX() int {
	rv := objc.Send[int](g_.ID, objc.Sel("dilationRateInX"))
	return rv
}


// SetDilationRateInX sets the value of the dilationRateInX property.
// Defines the dilation rate for the width dimension.

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphpooling2dopdescriptor/dilationrateinx
func (g_ GraphPooling2DOpDescriptor) SetDilationRateInX(value int) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setDilationRateInX:"), value)
}

// Defines the dilation rate for the height dimension.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphpooling2dopdescriptor/dilationrateiny
func (g_ GraphPooling2DOpDescriptor) DilationRateInY() int {
	rv := objc.Send[int](g_.ID, objc.Sel("dilationRateInY"))
	return rv
}


// SetDilationRateInY sets the value of the dilationRateInY property.
// Defines the dilation rate for the height dimension.

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphpooling2dopdescriptor/dilationrateiny
func (g_ GraphPooling2DOpDescriptor) SetDilationRateInY(value int) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setDilationRateInY:"), value)
}

// Defines a mode for average pooling, where samples outside the input tensor count as zeroes in the average computation.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphpooling2dopdescriptor/includezeropadtoaverage
func (g_ GraphPooling2DOpDescriptor) IncludeZeroPadToAverage() bool {
	rv := objc.Send[bool](g_.ID, objc.Sel("includeZeroPadToAverage"))
	return rv
}


// SetIncludeZeroPadToAverage sets the value of the includeZeroPadToAverage property.
// Defines a mode for average pooling, where samples outside the input tensor count as zeroes in the average computation.

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphpooling2dopdescriptor/includezeropadtoaverage
func (g_ GraphPooling2DOpDescriptor) SetIncludeZeroPadToAverage(value bool) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setIncludeZeroPadToAverage:"), value)
}

// Defines the pooling window size for the height dimension.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphpooling2dopdescriptor/kernelheight
func (g_ GraphPooling2DOpDescriptor) KernelHeight() int {
	rv := objc.Send[int](g_.ID, objc.Sel("kernelHeight"))
	return rv
}


// SetKernelHeight sets the value of the kernelHeight property.
// Defines the pooling window size for the height dimension.

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphpooling2dopdescriptor/kernelheight
func (g_ GraphPooling2DOpDescriptor) SetKernelHeight(value int) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setKernelHeight:"), value)
}

// Defines the pooling window size for the width dimension.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphpooling2dopdescriptor/kernelwidth
func (g_ GraphPooling2DOpDescriptor) KernelWidth() int {
	rv := objc.Send[int](g_.ID, objc.Sel("kernelWidth"))
	return rv
}


// SetKernelWidth sets the value of the kernelWidth property.
// Defines the pooling window size for the width dimension.

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphpooling2dopdescriptor/kernelwidth
func (g_ GraphPooling2DOpDescriptor) SetKernelWidth(value int) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setKernelWidth:"), value)
}

// Defines the explicit padding value for the height dimension to add after the data.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphpooling2dopdescriptor/paddingbottom
func (g_ GraphPooling2DOpDescriptor) PaddingBottom() int {
	rv := objc.Send[int](g_.ID, objc.Sel("paddingBottom"))
	return rv
}


// SetPaddingBottom sets the value of the paddingBottom property.
// Defines the explicit padding value for the height dimension to add after the data.

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphpooling2dopdescriptor/paddingbottom
func (g_ GraphPooling2DOpDescriptor) SetPaddingBottom(value int) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setPaddingBottom:"), value)
}

// Defines the explicit padding value for the width dimension to add after the data.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphpooling2dopdescriptor/paddingright
func (g_ GraphPooling2DOpDescriptor) PaddingRight() int {
	rv := objc.Send[int](g_.ID, objc.Sel("paddingRight"))
	return rv
}


// SetPaddingRight sets the value of the paddingRight property.
// Defines the explicit padding value for the width dimension to add after the data.

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphpooling2dopdescriptor/paddingright
func (g_ GraphPooling2DOpDescriptor) SetPaddingRight(value int) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setPaddingRight:"), value)
}

// Defines what kind of padding graph applies to the operation.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphpooling2dopdescriptor/paddingstyle
func (g_ GraphPooling2DOpDescriptor) PaddingStyle() GraphPaddingStyle {
	rv := objc.Send[GraphPaddingStyle](g_.ID, objc.Sel("paddingStyle"))
	return rv
}


// SetPaddingStyle sets the value of the paddingStyle property.
// Defines what kind of padding graph applies to the operation.

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphpooling2dopdescriptor/paddingstyle
func (g_ GraphPooling2DOpDescriptor) SetPaddingStyle(value GraphPaddingStyle) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setPaddingStyle:"), value)
}

// Defines the explicit padding value for the height dimension to add before the data.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphpooling2dopdescriptor/paddingtop
func (g_ GraphPooling2DOpDescriptor) PaddingTop() int {
	rv := objc.Send[int](g_.ID, objc.Sel("paddingTop"))
	return rv
}


// SetPaddingTop sets the value of the paddingTop property.
// Defines the explicit padding value for the height dimension to add before the data.

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphpooling2dopdescriptor/paddingtop
func (g_ GraphPooling2DOpDescriptor) SetPaddingTop(value int) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setPaddingTop:"), value)
}

// Defines the data type for returned indices.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphpooling2dopdescriptor/returnindicesdatatype
func (g_ GraphPooling2DOpDescriptor) ReturnIndicesDataType() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("returnIndicesDataType"))
	return rv
}


// SetReturnIndicesDataType sets the value of the returnIndicesDataType property.
// Defines the data type for returned indices.

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphpooling2dopdescriptor/returnindicesdatatype
func (g_ GraphPooling2DOpDescriptor) SetReturnIndicesDataType(value unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setReturnIndicesDataType:"), value)
}

// Defines the mode for returned indices of maximum values within each pooling window.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphpooling2dopdescriptor/returnindicesmode
func (g_ GraphPooling2DOpDescriptor) ReturnIndicesMode() GraphPoolingReturnIndicesMode {
	rv := objc.Send[GraphPoolingReturnIndicesMode](g_.ID, objc.Sel("returnIndicesMode"))
	return rv
}


// SetReturnIndicesMode sets the value of the returnIndicesMode property.
// Defines the mode for returned indices of maximum values within each pooling window.

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphpooling2dopdescriptor/returnindicesmode
func (g_ GraphPooling2DOpDescriptor) SetReturnIndicesMode(value GraphPoolingReturnIndicesMode) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setReturnIndicesMode:"), value)
}

// Defines the stride for the width dimension.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphpooling2dopdescriptor/strideinx
func (g_ GraphPooling2DOpDescriptor) StrideInX() int {
	rv := objc.Send[int](g_.ID, objc.Sel("strideInX"))
	return rv
}


// SetStrideInX sets the value of the strideInX property.
// Defines the stride for the width dimension.

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphpooling2dopdescriptor/strideinx
func (g_ GraphPooling2DOpDescriptor) SetStrideInX(value int) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setStrideInX:"), value)
}

// Defines the stride for the height dimension.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphpooling2dopdescriptor/strideiny
func (g_ GraphPooling2DOpDescriptor) StrideInY() int {
	rv := objc.Send[int](g_.ID, objc.Sel("strideInY"))
	return rv
}


// SetStrideInY sets the value of the strideInY property.
// Defines the stride for the height dimension.

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphpooling2dopdescriptor/strideiny
func (g_ GraphPooling2DOpDescriptor) SetStrideInY(value int) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setStrideInY:"), value)
}



