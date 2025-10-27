// Code generated from Apple documentation for MetalPerformanceShadersGraph. DO NOT EDIT.

package metalperformanceshadersgraph

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)





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





// An interface definition for the [GraphImToColOpDescriptor] class.
type IGraphImToColOpDescriptor interface {
	IGraphObject
	

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


	

	// methods:
	SetExplicitPaddingWithPaddingLeftPaddingRightPaddingTopPaddingBottom(paddingLeft uint, paddingRight uint, paddingTop uint, paddingBottom uint)


}





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






// Creates column to image descriptor with given values for parameters.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphImToColOpDescriptor/init(kernelWidth:kernelHeight:strideInX:strideInY:dilationRateInX:dilationRateInY:dataLayout:)
func NewGraphImToColOpDescriptorWithKernelWidthKernelHeightStrideInXStrideInYDilationRateInXDilationRateInYDataLayout(kernelWidth uint, kernelHeight uint, strideInX uint, strideInY uint, dilationRateInX uint, dilationRateInY uint, dataLayout GraphTensorNamedDataLayout) GraphImToColOpDescriptor {
	rv := objc.Send[GraphImToColOpDescriptor](objc.ID(getGraphImToColOpDescriptorClass().class), objc.Sel("descriptorWithKernelWidth:kernelHeight:strideInX:strideInY:dilationRateInX:dilationRateInY:dataLayout:"), kernelWidth, kernelHeight, strideInX, strideInY, dilationRateInX, dilationRateInY, dataLayout)
	return rv
}


// Creates an image to column descriptor with given values for parameters.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphImToColOpDescriptor/init(kernelWidth:kernelHeight:strideInX:strideInY:dilationRateInX:dilationRateInY:paddingLeft:paddingRight:paddingTop:paddingBottom:dataLayout:)
func NewGraphImToColOpDescriptorWithKernelWidthKernelHeightStrideInXStrideInYDilationRateInXDilationRateInYPaddingLeftPaddingRightPaddingTopPaddingBottomDataLayout(kernelWidth uint, kernelHeight uint, strideInX uint, strideInY uint, dilationRateInX uint, dilationRateInY uint, paddingLeft uint, paddingRight uint, paddingTop uint, paddingBottom uint, dataLayout GraphTensorNamedDataLayout) GraphImToColOpDescriptor {
	rv := objc.Send[GraphImToColOpDescriptor](objc.ID(getGraphImToColOpDescriptorClass().class), objc.Sel("descriptorWithKernelWidth:kernelHeight:strideInX:strideInY:dilationRateInX:dilationRateInY:paddingLeft:paddingRight:paddingTop:paddingBottom:dataLayout:"), kernelWidth, kernelHeight, strideInX, strideInY, dilationRateInX, dilationRateInY, paddingLeft, paddingRight, paddingTop, paddingBottom, dataLayout)
	return rv
}







// Creates column to image descriptor with given values for parameters.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphImToColOpDescriptor/init(kernelWidth:kernelHeight:strideInX:strideInY:dilationRateInX:dilationRateInY:dataLayout:)
func (gc _GraphImToColOpDescriptorClass) DescriptorWithKernelWidthKernelHeightStrideInXStrideInYDilationRateInXDilationRateInYDataLayout(kernelWidth uint, kernelHeight uint, strideInX uint, strideInY uint, dilationRateInX uint, dilationRateInY uint, dataLayout GraphTensorNamedDataLayout) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(gc.class), objc.Sel("descriptorWithKernelWidth:kernelHeight:strideInX:strideInY:dilationRateInX:dilationRateInY:dataLayout:"), kernelWidth, kernelHeight, strideInX, strideInY, dilationRateInX, dilationRateInY, dataLayout)
	return rv
}


// Creates an image to column descriptor with given values for parameters.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphImToColOpDescriptor/init(kernelWidth:kernelHeight:strideInX:strideInY:dilationRateInX:dilationRateInY:paddingLeft:paddingRight:paddingTop:paddingBottom:dataLayout:)
func (gc _GraphImToColOpDescriptorClass) DescriptorWithKernelWidthKernelHeightStrideInXStrideInYDilationRateInXDilationRateInYPaddingLeftPaddingRightPaddingTopPaddingBottomDataLayout(kernelWidth uint, kernelHeight uint, strideInX uint, strideInY uint, dilationRateInX uint, dilationRateInY uint, paddingLeft uint, paddingRight uint, paddingTop uint, paddingBottom uint, dataLayout GraphTensorNamedDataLayout) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(gc.class), objc.Sel("descriptorWithKernelWidth:kernelHeight:strideInX:strideInY:dilationRateInX:dilationRateInY:paddingLeft:paddingRight:paddingTop:paddingBottom:dataLayout:"), kernelWidth, kernelHeight, strideInX, strideInY, dilationRateInX, dilationRateInY, paddingLeft, paddingRight, paddingTop, paddingBottom, dataLayout)
	return rv
}












// Sets the descriptor’s padding to the given values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphImToColOpDescriptor/setExplicitPaddingWithPaddingLeft(_:paddingRight:paddingTop:paddingBottom:)
func (g_ GraphImToColOpDescriptor) SetExplicitPaddingWithPaddingLeftPaddingRightPaddingTopPaddingBottom(paddingLeft uint, paddingRight uint, paddingTop uint, paddingBottom uint) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setExplicitPaddingWithPaddingLeft:paddingRight:paddingTop:paddingBottom:"), paddingLeft, paddingRight, paddingTop, paddingBottom)
}







// The property that defines the layout of source or output tensor. e.g. for layout
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphImToColOpDescriptor/dataLayout
func (g_ GraphImToColOpDescriptor) DataLayout() GraphTensorNamedDataLayout {
	rv := objc.Send[GraphTensorNamedDataLayout](g_.ID, objc.Sel("dataLayout"))
	return rv
}


// The property that defines the layout of source or output tensor. e.g. for layout
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphImToColOpDescriptor/dataLayout
func (g_ GraphImToColOpDescriptor) SetDataLayout(value GraphTensorNamedDataLayout) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setDataLayout:"), value)
}


// The property that defines the dilation in width dimension.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphImToColOpDescriptor/dilationRateInX
func (g_ GraphImToColOpDescriptor) DilationRateInX() uint {
	rv := objc.Send[uint](g_.ID, objc.Sel("dilationRateInX"))
	return rv
}


// The property that defines the dilation in width dimension.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphImToColOpDescriptor/dilationRateInX
func (g_ GraphImToColOpDescriptor) SetDilationRateInX(value uint) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setDilationRateInX:"), value)
}


// The property that defines the dilation in height dimension.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphImToColOpDescriptor/dilationRateInY
func (g_ GraphImToColOpDescriptor) DilationRateInY() uint {
	rv := objc.Send[uint](g_.ID, objc.Sel("dilationRateInY"))
	return rv
}


// The property that defines the dilation in height dimension.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphImToColOpDescriptor/dilationRateInY
func (g_ GraphImToColOpDescriptor) SetDilationRateInY(value uint) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setDilationRateInY:"), value)
}


// The property that defines the kernel size in height dimension.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphImToColOpDescriptor/kernelHeight
func (g_ GraphImToColOpDescriptor) KernelHeight() uint {
	rv := objc.Send[uint](g_.ID, objc.Sel("kernelHeight"))
	return rv
}


// The property that defines the kernel size in height dimension.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphImToColOpDescriptor/kernelHeight
func (g_ GraphImToColOpDescriptor) SetKernelHeight(value uint) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setKernelHeight:"), value)
}


// The property that defines the kernel size in width dimension.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphImToColOpDescriptor/kernelWidth
func (g_ GraphImToColOpDescriptor) KernelWidth() uint {
	rv := objc.Send[uint](g_.ID, objc.Sel("kernelWidth"))
	return rv
}


// The property that defines the kernel size in width dimension.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphImToColOpDescriptor/kernelWidth
func (g_ GraphImToColOpDescriptor) SetKernelWidth(value uint) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setKernelWidth:"), value)
}


// The property that defines the padding in height dimension at the bottom.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphImToColOpDescriptor/paddingBottom
func (g_ GraphImToColOpDescriptor) PaddingBottom() uint {
	rv := objc.Send[uint](g_.ID, objc.Sel("paddingBottom"))
	return rv
}


// The property that defines the padding in height dimension at the bottom.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphImToColOpDescriptor/paddingBottom
func (g_ GraphImToColOpDescriptor) SetPaddingBottom(value uint) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setPaddingBottom:"), value)
}


// The property that defines the padding in width dimension on the left side.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphImToColOpDescriptor/paddingLeft
func (g_ GraphImToColOpDescriptor) PaddingLeft() uint {
	rv := objc.Send[uint](g_.ID, objc.Sel("paddingLeft"))
	return rv
}


// The property that defines the padding in width dimension on the left side.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphImToColOpDescriptor/paddingLeft
func (g_ GraphImToColOpDescriptor) SetPaddingLeft(value uint) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setPaddingLeft:"), value)
}


// The property that defines the padding in width dimension on the right side.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphImToColOpDescriptor/paddingRight
func (g_ GraphImToColOpDescriptor) PaddingRight() uint {
	rv := objc.Send[uint](g_.ID, objc.Sel("paddingRight"))
	return rv
}


// The property that defines the padding in width dimension on the right side.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphImToColOpDescriptor/paddingRight
func (g_ GraphImToColOpDescriptor) SetPaddingRight(value uint) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setPaddingRight:"), value)
}


// The property that defines the padding in height dimension at the top.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphImToColOpDescriptor/paddingTop
func (g_ GraphImToColOpDescriptor) PaddingTop() uint {
	rv := objc.Send[uint](g_.ID, objc.Sel("paddingTop"))
	return rv
}


// The property that defines the padding in height dimension at the top.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphImToColOpDescriptor/paddingTop
func (g_ GraphImToColOpDescriptor) SetPaddingTop(value uint) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setPaddingTop:"), value)
}


// The property that defines the stride in width dimension.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphImToColOpDescriptor/strideInX
func (g_ GraphImToColOpDescriptor) StrideInX() uint {
	rv := objc.Send[uint](g_.ID, objc.Sel("strideInX"))
	return rv
}


// The property that defines the stride in width dimension.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphImToColOpDescriptor/strideInX
func (g_ GraphImToColOpDescriptor) SetStrideInX(value uint) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setStrideInX:"), value)
}


// The property that defines the stride in height dimension.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphImToColOpDescriptor/strideInY
func (g_ GraphImToColOpDescriptor) StrideInY() uint {
	rv := objc.Send[uint](g_.ID, objc.Sel("strideInY"))
	return rv
}


// The property that defines the stride in height dimension.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphImToColOpDescriptor/strideInY
func (g_ GraphImToColOpDescriptor) SetStrideInY(value uint) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setStrideInY:"), value)
}







