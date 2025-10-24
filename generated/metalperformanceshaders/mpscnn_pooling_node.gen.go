// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [CNNPoolingNode] class.
var (
	CNNPoolingNodeClass     _CNNPoolingNodeClass
	CNNPoolingNodeClassOnce sync.Once
)

func getCNNPoolingNodeClass() _CNNPoolingNodeClass {
	CNNPoolingNodeClassOnce.Do(func() {
		CNNPoolingNodeClass = _CNNPoolingNodeClass{objc.GetClass("MPSCNNPoolingNode")}
	})
	return CNNPoolingNodeClass
}

type _CNNPoolingNodeClass struct {
	class objc.Class
}





// An interface definition for the [CNNPoolingNode] class.
type ICNNPoolingNode interface {
	IFilterNode
	

	// properties:
	KernelHeight() objectivec.IObject
	SetKernelHeight(value objectivec.IObject)
	KernelWidth() objectivec.IObject
	SetKernelWidth(value objectivec.IObject)
	StrideInPixelsX() objectivec.IObject
	SetStrideInPixelsX(value objectivec.IObject)
	StrideInPixelsY() objectivec.IObject
	SetStrideInPixelsY(value objectivec.IObject)


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (cc _CNNPoolingNodeClass) Alloc() CNNPoolingNode {
	rv := objc.Send[CNNPoolingNode](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CNNPoolingNodeClass) New() CNNPoolingNode {
	rv := objc.Send[CNNPoolingNode](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNNPoolingNode) Init() CNNPoolingNode {
	rv := objc.Send[CNNPoolingNode](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNNPoolingNode) Autorelease() CNNPoolingNode {
	rv := objc.Send[CNNPoolingNode](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNNPoolingNode creates a new CNNPoolingNode instance.
func NewCNNPoolingNode() CNNPoolingNode {
	return getCNNPoolingNodeClass().New()
}





// A representation of a MPS CNN pooling kernel.


// A representation of a MPS CNN pooling kernel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNPoolingNode
type CNNPoolingNode struct {
	FilterNode
}

// CNNPoolingNodeFrom constructs a [CNNPoolingNode] from an unsafe.Pointer.
//
// A representation of a MPS CNN pooling kernel.
func CNNPoolingNodeFrom(ptr unsafe.Pointer) CNNPoolingNode {
	return CNNPoolingNode{
		FilterNode: FilterNodeFrom(ptr),
	}
}






// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnpoolingnode/2866488-initwithsource
func NewCNNPoolingNodeWithSourceFilterSize(sourceNode IImageNode, size uint) CNNPoolingNode {
	instance := getCNNPoolingNodeClass().Alloc()
	rv := objc.Send[CNNPoolingNode](instance.ID, objc.Sel("initWithSource:filterSize:"), sourceNode, size)
	rv.Autorelease()
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnpoolingnode/2866444-initwithsource
func NewCNNPoolingNodeWithSourceFilterSizeStride(sourceNode IImageNode, size uint, stride uint) CNNPoolingNode {
	instance := getCNNPoolingNodeClass().Alloc()
	rv := objc.Send[CNNPoolingNode](instance.ID, objc.Sel("initWithSource:filterSize:stride:"), sourceNode, size, stride)
	rv.Autorelease()
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnpoolingnode/2866471-initwithsource
func NewCNNPoolingNodeWithSourceKernelWidthKernelHeightStrideInPixelsXStrideInPixelsY(sourceNode IImageNode, kernelWidth uint, kernelHeight uint, strideInPixelsX uint, strideInPixelsY uint) CNNPoolingNode {
	instance := getCNNPoolingNodeClass().Alloc()
	rv := objc.Send[CNNPoolingNode](instance.ID, objc.Sel("initWithSource:kernelWidth:kernelHeight:strideInPixelsX:strideInPixelsY:"), sourceNode, kernelWidth, kernelHeight, strideInPixelsX, strideInPixelsY)
	rv.Autorelease()
	return rv
}







// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnpoolingnode/2866508-nodewithsource
func (cc _CNNPoolingNodeClass) NodeWithSourceFilterSize(sourceNode IImageNode, size uint) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(cc.class), objc.Sel("nodeWithSource:filterSize:"), sourceNode, size)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnpoolingnode/2890831-nodewithsource
func (cc _CNNPoolingNodeClass) NodeWithSourceFilterSizeStride(sourceNode IImageNode, size uint, stride uint) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(cc.class), objc.Sel("nodeWithSource:filterSize:stride:"), sourceNode, size, stride)
	return rv
}

















// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnpoolingnode/2993001-kernelheight
func (c_ CNNPoolingNode) KernelHeight() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("kernelHeight"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnpoolingnode/2993001-kernelheight
func (c_ CNNPoolingNode) SetKernelHeight(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setKernelHeight:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnpoolingnode/2993002-kernelwidth
func (c_ CNNPoolingNode) KernelWidth() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("kernelWidth"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnpoolingnode/2993002-kernelwidth
func (c_ CNNPoolingNode) SetKernelWidth(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setKernelWidth:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnpoolingnode/2993003-strideinpixelsx
func (c_ CNNPoolingNode) StrideInPixelsX() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("strideInPixelsX"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnpoolingnode/2993003-strideinpixelsx
func (c_ CNNPoolingNode) SetStrideInPixelsX(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setStrideInPixelsX:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnpoolingnode/2993004-strideinpixelsy
func (c_ CNNPoolingNode) StrideInPixelsY() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("strideInPixelsY"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnpoolingnode/2993004-strideinpixelsy
func (c_ CNNPoolingNode) SetStrideInPixelsY(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setStrideInPixelsY:"), value)
}







