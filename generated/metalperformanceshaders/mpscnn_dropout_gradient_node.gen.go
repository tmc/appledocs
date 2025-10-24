// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [CNNDropoutGradientNode] class.
var (
	CNNDropoutGradientNodeClass     _CNNDropoutGradientNodeClass
	CNNDropoutGradientNodeClassOnce sync.Once
)

func getCNNDropoutGradientNodeClass() _CNNDropoutGradientNodeClass {
	CNNDropoutGradientNodeClassOnce.Do(func() {
		CNNDropoutGradientNodeClass = _CNNDropoutGradientNodeClass{objc.GetClass("MPSCNNDropoutGradientNode")}
	})
	return CNNDropoutGradientNodeClass
}

type _CNNDropoutGradientNodeClass struct {
	class objc.Class
}





// An interface definition for the [CNNDropoutGradientNode] class.
type ICNNDropoutGradientNode interface {
	IGradientFilterNode
	

	// properties:
	MaskStrideInPixels() Size get /* not a class type */
	SetMaskStrideInPixels(value Size get /* not a class type */)
	KeepProbability() objectivec.IObject
	SetKeepProbability(value objectivec.IObject)
	Seed() objectivec.IObject
	SetSeed(value objectivec.IObject)


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (cc _CNNDropoutGradientNodeClass) Alloc() CNNDropoutGradientNode {
	rv := objc.Send[CNNDropoutGradientNode](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CNNDropoutGradientNodeClass) New() CNNDropoutGradientNode {
	rv := objc.Send[CNNDropoutGradientNode](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNNDropoutGradientNode) Init() CNNDropoutGradientNode {
	rv := objc.Send[CNNDropoutGradientNode](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNNDropoutGradientNode) Autorelease() CNNDropoutGradientNode {
	rv := objc.Send[CNNDropoutGradientNode](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNNDropoutGradientNode creates a new CNNDropoutGradientNode instance.
func NewCNNDropoutGradientNode() CNNDropoutGradientNode {
	return getCNNDropoutGradientNodeClass().New()
}





// A representation of a gradient dropout filter.


// A representation of a gradient dropout filter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNDropoutGradientNode
type CNNDropoutGradientNode struct {
	GradientFilterNode
}

// CNNDropoutGradientNodeFrom constructs a [CNNDropoutGradientNode] from an unsafe.Pointer.
//
// A representation of a gradient dropout filter.
func CNNDropoutGradientNodeFrom(ptr unsafe.Pointer) CNNDropoutGradientNode {
	return CNNDropoutGradientNode{
		GradientFilterNode: GradientFilterNodeFrom(ptr),
	}
}






// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnndropoutgradientnode/2948001-initwithsourcegradient
func NewCNNDropoutGradientNodeWithSourceGradientSourceImageGradientStateKeepProbabilitySeedMaskStrideInPixels(sourceGradient IImageNode, sourceImage IImageNode, gradientState IGradientStateNode, keepProbability float32, seed uint, maskStrideInPixels objc.IObject /* cross-framework: MTLSize */) CNNDropoutGradientNode {
	instance := getCNNDropoutGradientNodeClass().Alloc()
	rv := objc.Send[CNNDropoutGradientNode](instance.ID, objc.Sel("initWithSourceGradient:sourceImage:gradientState:keepProbability:seed:maskStrideInPixels:"), sourceGradient, sourceImage, gradientState, keepProbability, seed, maskStrideInPixels)
	rv.Autorelease()
	return rv
}







// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnndropoutgradientnode/2947997-nodewithsourcegradient
func (cc _CNNDropoutGradientNodeClass) NodeWithSourceGradientSourceImageGradientStateKeepProbabilitySeedMaskStrideInPixels(sourceGradient IImageNode, sourceImage IImageNode, gradientState IGradientStateNode, keepProbability float32, seed uint, maskStrideInPixels objc.IObject /* cross-framework: MTLSize */) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(cc.class), objc.Sel("nodeWithSourceGradient:sourceImage:gradientState:keepProbability:seed:maskStrideInPixels:"), sourceGradient, sourceImage, gradientState, keepProbability, seed, maskStrideInPixels)
	return rv
}

















// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnndropoutgradientnode/2947972-maskstrideinpixels
func (c_ CNNDropoutGradientNode) MaskStrideInPixels() Size get /* not a class type */ {
	rv := objc.Send[objc.ID](c_.ID, objc.Sel("maskStrideInPixels"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnndropoutgradientnode/2947972-maskstrideinpixels
func (c_ CNNDropoutGradientNode) SetMaskStrideInPixels(value Size get /* not a class type */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setMaskStrideInPixels:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnndropoutgradientnode/2947988-keepprobability
func (c_ CNNDropoutGradientNode) KeepProbability() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("keepProbability"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnndropoutgradientnode/2947988-keepprobability
func (c_ CNNDropoutGradientNode) SetKeepProbability(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setKeepProbability:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnndropoutgradientnode/2948003-seed
func (c_ CNNDropoutGradientNode) Seed() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("seed"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnndropoutgradientnode/2948003-seed
func (c_ CNNDropoutGradientNode) SetSeed(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSeed:"), value)
}







