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
}

// The class that defines the parameters for an image to column or column to image operation.
//
// Use this descriptor with the following methods:
//
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

// Alloc allocates a new instance without initialization.
func (gc _GraphImToColOpDescriptorClass) Alloc() GraphImToColOpDescriptor {
	rv := objc.Send[GraphImToColOpDescriptor](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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


// The property that defines the dilation in width dimension.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphImToColOpDescriptor/dilationRateInX
func (g_ GraphImToColOpDescriptor) DilationRateInX() uint {
	rv := objc.Send[uint](g_.ID, objc.Sel("dilationRateInX"))
	return rv
}


// SetDilationRateInX sets the value of the dilationRateInX property.
// The property that defines the dilation in width dimension.

//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphImToColOpDescriptor/dilationRateInX
func (g_ GraphImToColOpDescriptor) SetDilationRateInX(value uint) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setDilationRateInX:"), value)
}

// The property that defines the stride in height dimension.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphImToColOpDescriptor/strideInY
func (g_ GraphImToColOpDescriptor) StrideInY() uint {
	rv := objc.Send[uint](g_.ID, objc.Sel("strideInY"))
	return rv
}


// SetStrideInY sets the value of the strideInY property.
// The property that defines the stride in height dimension.

//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphImToColOpDescriptor/strideInY
func (g_ GraphImToColOpDescriptor) SetStrideInY(value uint) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setStrideInY:"), value)
}

// The property that defines the layout of source or output tensor.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphimtocolopdescriptor/datalayout
func (g_ GraphImToColOpDescriptor) DataLayout() GraphTensorNamedDataLayout {
	rv := objc.Send[GraphTensorNamedDataLayout](g_.ID, objc.Sel("dataLayout"))
	return rv
}


// SetDataLayout sets the value of the dataLayout property.
// The property that defines the layout of source or output tensor.

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphimtocolopdescriptor/datalayout
func (g_ GraphImToColOpDescriptor) SetDataLayout(value IGraphTensorNamedDataLayout) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setDataLayout:"), value)
}

// The property that defines the dilation in height dimension.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphimtocolopdescriptor/dilationrateiny
func (g_ GraphImToColOpDescriptor) DilationRateInY() int {
	rv := objc.Send[int](g_.ID, objc.Sel("dilationRateInY"))
	return rv
}


// SetDilationRateInY sets the value of the dilationRateInY property.
// The property that defines the dilation in height dimension.

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphimtocolopdescriptor/dilationrateiny
func (g_ GraphImToColOpDescriptor) SetDilationRateInY(value int) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setDilationRateInY:"), value)
}

// The property that defines the kernel size in height dimension.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphimtocolopdescriptor/kernelheight
func (g_ GraphImToColOpDescriptor) KernelHeight() int {
	rv := objc.Send[int](g_.ID, objc.Sel("kernelHeight"))
	return rv
}


// SetKernelHeight sets the value of the kernelHeight property.
// The property that defines the kernel size in height dimension.

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphimtocolopdescriptor/kernelheight
func (g_ GraphImToColOpDescriptor) SetKernelHeight(value int) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setKernelHeight:"), value)
}

// The property that defines the kernel size in width dimension.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphimtocolopdescriptor/kernelwidth
func (g_ GraphImToColOpDescriptor) KernelWidth() int {
	rv := objc.Send[int](g_.ID, objc.Sel("kernelWidth"))
	return rv
}


// SetKernelWidth sets the value of the kernelWidth property.
// The property that defines the kernel size in width dimension.

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphimtocolopdescriptor/kernelwidth
func (g_ GraphImToColOpDescriptor) SetKernelWidth(value int) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setKernelWidth:"), value)
}

// The property that defines the padding in height dimension at the bottom.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphimtocolopdescriptor/paddingbottom
func (g_ GraphImToColOpDescriptor) PaddingBottom() int {
	rv := objc.Send[int](g_.ID, objc.Sel("paddingBottom"))
	return rv
}


// SetPaddingBottom sets the value of the paddingBottom property.
// The property that defines the padding in height dimension at the bottom.

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphimtocolopdescriptor/paddingbottom
func (g_ GraphImToColOpDescriptor) SetPaddingBottom(value int) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setPaddingBottom:"), value)
}

// The property that defines the padding in width dimension on the left side.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphimtocolopdescriptor/paddingleft
func (g_ GraphImToColOpDescriptor) PaddingLeft() int {
	rv := objc.Send[int](g_.ID, objc.Sel("paddingLeft"))
	return rv
}


// SetPaddingLeft sets the value of the paddingLeft property.
// The property that defines the padding in width dimension on the left side.

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphimtocolopdescriptor/paddingleft
func (g_ GraphImToColOpDescriptor) SetPaddingLeft(value int) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setPaddingLeft:"), value)
}

// The property that defines the padding in width dimension on the right side.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphimtocolopdescriptor/paddingright
func (g_ GraphImToColOpDescriptor) PaddingRight() int {
	rv := objc.Send[int](g_.ID, objc.Sel("paddingRight"))
	return rv
}


// SetPaddingRight sets the value of the paddingRight property.
// The property that defines the padding in width dimension on the right side.

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphimtocolopdescriptor/paddingright
func (g_ GraphImToColOpDescriptor) SetPaddingRight(value int) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setPaddingRight:"), value)
}

// The property that defines the padding in height dimension at the top.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphimtocolopdescriptor/paddingtop
func (g_ GraphImToColOpDescriptor) PaddingTop() int {
	rv := objc.Send[int](g_.ID, objc.Sel("paddingTop"))
	return rv
}


// SetPaddingTop sets the value of the paddingTop property.
// The property that defines the padding in height dimension at the top.

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphimtocolopdescriptor/paddingtop
func (g_ GraphImToColOpDescriptor) SetPaddingTop(value int) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setPaddingTop:"), value)
}

// The property that defines the stride in width dimension.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphimtocolopdescriptor/strideinx
func (g_ GraphImToColOpDescriptor) StrideInX() int {
	rv := objc.Send[int](g_.ID, objc.Sel("strideInX"))
	return rv
}


// SetStrideInX sets the value of the strideInX property.
// The property that defines the stride in width dimension.

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphimtocolopdescriptor/strideinx
func (g_ GraphImToColOpDescriptor) SetStrideInX(value int) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setStrideInX:"), value)
}



