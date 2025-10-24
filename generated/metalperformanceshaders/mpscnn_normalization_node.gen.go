// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [CNNNormalizationNode] class.
var (
	CNNNormalizationNodeClass     _CNNNormalizationNodeClass
	CNNNormalizationNodeClassOnce sync.Once
)

func getCNNNormalizationNodeClass() _CNNNormalizationNodeClass {
	CNNNormalizationNodeClassOnce.Do(func() {
		CNNNormalizationNodeClass = _CNNNormalizationNodeClass{objc.GetClass("MPSCNNNormalizationNode")}
	})
	return CNNNormalizationNodeClass
}

type _CNNNormalizationNodeClass struct {
	class objc.Class
}





// An interface definition for the [CNNNormalizationNode] class.
type ICNNNormalizationNode interface {
	IFilterNode
	

	// properties:
	Alpha() objectivec.IObject
	SetAlpha(value objectivec.IObject)
	Delta() objectivec.IObject
	SetDelta(value objectivec.IObject)
	Beta() objectivec.IObject
	SetBeta(value objectivec.IObject)


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (cc _CNNNormalizationNodeClass) Alloc() CNNNormalizationNode {
	rv := objc.Send[CNNNormalizationNode](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CNNNormalizationNodeClass) New() CNNNormalizationNode {
	rv := objc.Send[CNNNormalizationNode](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNNNormalizationNode) Init() CNNNormalizationNode {
	rv := objc.Send[CNNNormalizationNode](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNNNormalizationNode) Autorelease() CNNNormalizationNode {
	rv := objc.Send[CNNNormalizationNode](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNNNormalizationNode creates a new CNNNormalizationNode instance.
func NewCNNNormalizationNode() CNNNormalizationNode {
	return getCNNNormalizationNodeClass().New()
}





// Virtual base class for CNN normalization nodes.


// Virtual base class for CNN normalization nodes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNormalizationNode
type CNNNormalizationNode struct {
	FilterNode
}

// CNNNormalizationNodeFrom constructs a [CNNNormalizationNode] from an unsafe.Pointer.
//
// Virtual base class for CNN normalization nodes.
func CNNNormalizationNodeFrom(ptr unsafe.Pointer) CNNNormalizationNode {
	return CNNNormalizationNode{
		FilterNode: FilterNodeFrom(ptr),
	}
}






// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnnormalizationnode/2866425-initwithsource
func NewCNNNormalizationNodeWithSource(sourceNode IImageNode) CNNNormalizationNode {
	instance := getCNNNormalizationNodeClass().Alloc()
	rv := objc.Send[CNNNormalizationNode](instance.ID, objc.Sel("initWithSource:"), sourceNode)
	rv.Autorelease()
	return rv
}







// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnnormalizationnode/2866460-nodewithsource
func (cc _CNNNormalizationNodeClass) NodeWithSource(sourceNode IImageNode) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(cc.class), objc.Sel("nodeWithSource:"), sourceNode)
	return rv
}

















// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnnormalizationnode/2866474-alpha
func (c_ CNNNormalizationNode) Alpha() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("alpha"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnnormalizationnode/2866474-alpha
func (c_ CNNNormalizationNode) SetAlpha(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAlpha:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnnormalizationnode/2866482-delta
func (c_ CNNNormalizationNode) Delta() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("delta"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnnormalizationnode/2866482-delta
func (c_ CNNNormalizationNode) SetDelta(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setDelta:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnnormalizationnode/2866497-beta
func (c_ CNNNormalizationNode) Beta() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("beta"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnnormalizationnode/2866497-beta
func (c_ CNNNormalizationNode) SetBeta(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setBeta:"), value)
}







