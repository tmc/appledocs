// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [CNNPoolingGradientNode] class.
var (
	CNNPoolingGradientNodeClass     _CNNPoolingGradientNodeClass
	CNNPoolingGradientNodeClassOnce sync.Once
)

func getCNNPoolingGradientNodeClass() _CNNPoolingGradientNodeClass {
	CNNPoolingGradientNodeClassOnce.Do(func() {
		CNNPoolingGradientNodeClass = _CNNPoolingGradientNodeClass{objc.GetClass("MPSCNNPoolingGradientNode")}
	})
	return CNNPoolingGradientNodeClass
}

type _CNNPoolingGradientNodeClass struct {
	class objc.Class
}





// An interface definition for the [CNNPoolingGradientNode] class.
type ICNNPoolingGradientNode interface {
	IGradientFilterNode
	

	// properties:
	KernelHeight() objectivec.IObject
	SetKernelHeight(value objectivec.IObject)
	StrideInPixelsX() objectivec.IObject
	SetStrideInPixelsX(value objectivec.IObject)
	KernelWidth() objectivec.IObject
	SetKernelWidth(value objectivec.IObject)
	StrideInPixelsY() objectivec.IObject
	SetStrideInPixelsY(value objectivec.IObject)


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (cc _CNNPoolingGradientNodeClass) Alloc() CNNPoolingGradientNode {
	rv := objc.Send[CNNPoolingGradientNode](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CNNPoolingGradientNodeClass) New() CNNPoolingGradientNode {
	rv := objc.Send[CNNPoolingGradientNode](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNNPoolingGradientNode) Init() CNNPoolingGradientNode {
	rv := objc.Send[CNNPoolingGradientNode](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNNPoolingGradientNode) Autorelease() CNNPoolingGradientNode {
	rv := objc.Send[CNNPoolingGradientNode](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNNPoolingGradientNode creates a new CNNPoolingGradientNode instance.
func NewCNNPoolingGradientNode() CNNPoolingGradientNode {
	return getCNNPoolingGradientNodeClass().New()
}





// A representation of a gradient pooling kernel.


// A representation of a gradient pooling kernel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNPoolingGradientNode
type CNNPoolingGradientNode struct {
	GradientFilterNode
}

// CNNPoolingGradientNodeFrom constructs a [CNNPoolingGradientNode] from an unsafe.Pointer.
//
// A representation of a gradient pooling kernel.
func CNNPoolingGradientNodeFrom(ptr unsafe.Pointer) CNNPoolingGradientNode {
	return CNNPoolingGradientNode{
		GradientFilterNode: GradientFilterNodeFrom(ptr),
	}
}






// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnpoolinggradientnode/2948011-initwithsourcegradient
func NewCNNPoolingGradientNodeWithSourceGradientSourceImageGradientStateKernelWidthKernelHeightStrideInPixelsXStrideInPixelsYPaddingPolicy(sourceGradient IImageNode, sourceImage IImageNode, gradientState IGradientStateNode, kernelWidth uint, kernelHeight uint, strideInPixelsX uint, strideInPixelsY uint, paddingPolicy unsafe.Pointer) CNNPoolingGradientNode {
	instance := getCNNPoolingGradientNodeClass().Alloc()
	rv := objc.Send[CNNPoolingGradientNode](instance.ID, objc.Sel("initWithSourceGradient:sourceImage:gradientState:kernelWidth:kernelHeight:strideInPixelsX:strideInPixelsY:paddingPolicy:"), sourceGradient, sourceImage, gradientState, kernelWidth, kernelHeight, strideInPixelsX, strideInPixelsY, paddingPolicy)
	rv.Autorelease()
	return rv
}







// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnpoolinggradientnode/2948045-nodewithsourcegradient
func (cc _CNNPoolingGradientNodeClass) NodeWithSourceGradientSourceImageGradientStateKernelWidthKernelHeightStrideInPixelsXStrideInPixelsYPaddingPolicy(sourceGradient IImageNode, sourceImage IImageNode, gradientState IGradientStateNode, kernelWidth uint, kernelHeight uint, strideInPixelsX uint, strideInPixelsY uint, paddingPolicy unsafe.Pointer) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(cc.class), objc.Sel("nodeWithSourceGradient:sourceImage:gradientState:kernelWidth:kernelHeight:strideInPixelsX:strideInPixelsY:paddingPolicy:"), sourceGradient, sourceImage, gradientState, kernelWidth, kernelHeight, strideInPixelsX, strideInPixelsY, paddingPolicy)
	return rv
}

















// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnpoolinggradientnode/2947992-kernelheight
func (c_ CNNPoolingGradientNode) KernelHeight() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("kernelHeight"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnpoolinggradientnode/2947992-kernelheight
func (c_ CNNPoolingGradientNode) SetKernelHeight(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setKernelHeight:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnpoolinggradientnode/2948018-strideinpixelsx
func (c_ CNNPoolingGradientNode) StrideInPixelsX() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("strideInPixelsX"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnpoolinggradientnode/2948018-strideinpixelsx
func (c_ CNNPoolingGradientNode) SetStrideInPixelsX(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setStrideInPixelsX:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnpoolinggradientnode/2948034-kernelwidth
func (c_ CNNPoolingGradientNode) KernelWidth() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("kernelWidth"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnpoolinggradientnode/2948034-kernelwidth
func (c_ CNNPoolingGradientNode) SetKernelWidth(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setKernelWidth:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnpoolinggradientnode/2948048-strideinpixelsy
func (c_ CNNPoolingGradientNode) StrideInPixelsY() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("strideInPixelsY"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnpoolinggradientnode/2948048-strideinpixelsy
func (c_ CNNPoolingGradientNode) SetStrideInPixelsY(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setStrideInPixelsY:"), value)
}







