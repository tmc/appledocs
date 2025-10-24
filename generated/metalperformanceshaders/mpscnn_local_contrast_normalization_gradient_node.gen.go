// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [CNNLocalContrastNormalizationGradientNode] class.
var (
	CNNLocalContrastNormalizationGradientNodeClass     _CNNLocalContrastNormalizationGradientNodeClass
	CNNLocalContrastNormalizationGradientNodeClassOnce sync.Once
)

func getCNNLocalContrastNormalizationGradientNodeClass() _CNNLocalContrastNormalizationGradientNodeClass {
	CNNLocalContrastNormalizationGradientNodeClassOnce.Do(func() {
		CNNLocalContrastNormalizationGradientNodeClass = _CNNLocalContrastNormalizationGradientNodeClass{objc.GetClass("MPSCNNLocalContrastNormalizationGradientNode")}
	})
	return CNNLocalContrastNormalizationGradientNodeClass
}

type _CNNLocalContrastNormalizationGradientNodeClass struct {
	class objc.Class
}





// An interface definition for the [CNNLocalContrastNormalizationGradientNode] class.
type ICNNLocalContrastNormalizationGradientNode interface {
	IGradientFilterNode
	

	// properties:
	KernelWidth() objectivec.IObject
	SetKernelWidth(value objectivec.IObject)
	Alpha() objectivec.IObject
	SetAlpha(value objectivec.IObject)
	Beta() objectivec.IObject
	SetBeta(value objectivec.IObject)
	Ps() objectivec.IObject
	SetPs(value objectivec.IObject)
	Delta() objectivec.IObject
	SetDelta(value objectivec.IObject)
	P0() objectivec.IObject
	SetP0(value objectivec.IObject)
	KernelHeight() objectivec.IObject
	SetKernelHeight(value objectivec.IObject)
	Pm() objectivec.IObject
	SetPm(value objectivec.IObject)


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (cc _CNNLocalContrastNormalizationGradientNodeClass) Alloc() CNNLocalContrastNormalizationGradientNode {
	rv := objc.Send[CNNLocalContrastNormalizationGradientNode](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CNNLocalContrastNormalizationGradientNodeClass) New() CNNLocalContrastNormalizationGradientNode {
	rv := objc.Send[CNNLocalContrastNormalizationGradientNode](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNNLocalContrastNormalizationGradientNode) Init() CNNLocalContrastNormalizationGradientNode {
	rv := objc.Send[CNNLocalContrastNormalizationGradientNode](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNNLocalContrastNormalizationGradientNode) Autorelease() CNNLocalContrastNormalizationGradientNode {
	rv := objc.Send[CNNLocalContrastNormalizationGradientNode](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNNLocalContrastNormalizationGradientNode creates a new CNNLocalContrastNormalizationGradientNode instance.
func NewCNNLocalContrastNormalizationGradientNode() CNNLocalContrastNormalizationGradientNode {
	return getCNNLocalContrastNormalizationGradientNodeClass().New()
}





// A representation of a gradient local-contrast normalization kernel.


// A representation of a gradient local-contrast normalization kernel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNLocalContrastNormalizationGradientNode
type CNNLocalContrastNormalizationGradientNode struct {
	GradientFilterNode
}

// CNNLocalContrastNormalizationGradientNodeFrom constructs a [CNNLocalContrastNormalizationGradientNode] from an unsafe.Pointer.
//
// A representation of a gradient local-contrast normalization kernel.
func CNNLocalContrastNormalizationGradientNodeFrom(ptr unsafe.Pointer) CNNLocalContrastNormalizationGradientNode {
	return CNNLocalContrastNormalizationGradientNode{
		GradientFilterNode: GradientFilterNodeFrom(ptr),
	}
}






// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnlocalcontrastnormalizationgradientnode/2948016-initwithsourcegradient
func NewCNNLocalContrastNormalizationGradientNodeWithSourceGradientSourceImageGradientStateKernelWidthKernelHeight(sourceGradient IImageNode, sourceImage IImageNode, gradientState IGradientStateNode, kernelWidth uint, kernelHeight uint) CNNLocalContrastNormalizationGradientNode {
	instance := getCNNLocalContrastNormalizationGradientNodeClass().Alloc()
	rv := objc.Send[CNNLocalContrastNormalizationGradientNode](instance.ID, objc.Sel("initWithSourceGradient:sourceImage:gradientState:kernelWidth:kernelHeight:"), sourceGradient, sourceImage, gradientState, kernelWidth, kernelHeight)
	rv.Autorelease()
	return rv
}







// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnlocalcontrastnormalizationgradientnode/2948047-nodewithsourcegradient
func (cc _CNNLocalContrastNormalizationGradientNodeClass) NodeWithSourceGradientSourceImageGradientStateKernelWidthKernelHeight(sourceGradient IImageNode, sourceImage IImageNode, gradientState IGradientStateNode, kernelWidth uint, kernelHeight uint) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(cc.class), objc.Sel("nodeWithSourceGradient:sourceImage:gradientState:kernelWidth:kernelHeight:"), sourceGradient, sourceImage, gradientState, kernelWidth, kernelHeight)
	return rv
}

















// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnlocalcontrastnormalizationgradientnode/2947965-kernelwidth
func (c_ CNNLocalContrastNormalizationGradientNode) KernelWidth() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("kernelWidth"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnlocalcontrastnormalizationgradientnode/2947965-kernelwidth
func (c_ CNNLocalContrastNormalizationGradientNode) SetKernelWidth(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setKernelWidth:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnlocalcontrastnormalizationgradientnode/2947973-alpha
func (c_ CNNLocalContrastNormalizationGradientNode) Alpha() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("alpha"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnlocalcontrastnormalizationgradientnode/2947973-alpha
func (c_ CNNLocalContrastNormalizationGradientNode) SetAlpha(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAlpha:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnlocalcontrastnormalizationgradientnode/2947977-beta
func (c_ CNNLocalContrastNormalizationGradientNode) Beta() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("beta"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnlocalcontrastnormalizationgradientnode/2947977-beta
func (c_ CNNLocalContrastNormalizationGradientNode) SetBeta(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setBeta:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnlocalcontrastnormalizationgradientnode/2948008-ps
func (c_ CNNLocalContrastNormalizationGradientNode) Ps() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("ps"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnlocalcontrastnormalizationgradientnode/2948008-ps
func (c_ CNNLocalContrastNormalizationGradientNode) SetPs(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPs:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnlocalcontrastnormalizationgradientnode/2948014-delta
func (c_ CNNLocalContrastNormalizationGradientNode) Delta() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("delta"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnlocalcontrastnormalizationgradientnode/2948014-delta
func (c_ CNNLocalContrastNormalizationGradientNode) SetDelta(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setDelta:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnlocalcontrastnormalizationgradientnode/2948017-p0
func (c_ CNNLocalContrastNormalizationGradientNode) P0() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("p0"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnlocalcontrastnormalizationgradientnode/2948017-p0
func (c_ CNNLocalContrastNormalizationGradientNode) SetP0(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setP0:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnlocalcontrastnormalizationgradientnode/2948019-kernelheight
func (c_ CNNLocalContrastNormalizationGradientNode) KernelHeight() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("kernelHeight"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnlocalcontrastnormalizationgradientnode/2948019-kernelheight
func (c_ CNNLocalContrastNormalizationGradientNode) SetKernelHeight(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setKernelHeight:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnlocalcontrastnormalizationgradientnode/2948053-pm
func (c_ CNNLocalContrastNormalizationGradientNode) Pm() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("pm"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnlocalcontrastnormalizationgradientnode/2948053-pm
func (c_ CNNLocalContrastNormalizationGradientNode) SetPm(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPm:"), value)
}







