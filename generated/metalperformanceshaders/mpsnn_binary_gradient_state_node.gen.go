// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)





// The class instance for the [BinaryGradientStateNode] class.
var (
	BinaryGradientStateNodeClass     _BinaryGradientStateNodeClass
	BinaryGradientStateNodeClassOnce sync.Once
)

func getBinaryGradientStateNodeClass() _BinaryGradientStateNodeClass {
	BinaryGradientStateNodeClassOnce.Do(func() {
		BinaryGradientStateNodeClass = _BinaryGradientStateNodeClass{objc.GetClass("MPSNNBinaryGradientStateNode")}
	})
	return BinaryGradientStateNodeClass
}

type _BinaryGradientStateNodeClass struct {
	class objc.Class
}





// An interface definition for the [BinaryGradientStateNode] class.
type IBinaryGradientStateNode interface {
	IStateNode
	

	// properties:
	Label() objc.IObject /* cross-framework: NSString */
	SetLabel(value objc.IObject /* cross-framework: NSString */)
	PaddingPolicy() Padding /* not a class type */
	SetPaddingPolicy(value Padding /* not a class type */)
	ResultImage() IMPSNNImageNode
	SetResultImage(value IMPSNNImageNode)
	ResultState() IMPSNNStateNode
	SetResultState(value IMPSNNStateNode)
	ResultStates() IMPSNNStateNode
	SetResultStates(value IMPSNNStateNode)


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (bc _BinaryGradientStateNodeClass) Alloc() BinaryGradientStateNode {
	rv := objc.Send[BinaryGradientStateNode](objc.ID(bc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (bc _BinaryGradientStateNodeClass) New() BinaryGradientStateNode {
	rv := objc.Send[BinaryGradientStateNode](objc.ID(bc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (b_ BinaryGradientStateNode) Init() BinaryGradientStateNode {
	rv := objc.Send[BinaryGradientStateNode](b_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (b_ BinaryGradientStateNode) Autorelease() BinaryGradientStateNode {
	rv := objc.Send[BinaryGradientStateNode](b_.ID, objc.Sel("autorelease"))
	return rv
}

// NewBinaryGradientStateNode creates a new BinaryGradientStateNode instance.
func NewBinaryGradientStateNode() BinaryGradientStateNode {
	return getBinaryGradientStateNodeClass().New()
}





// A representation of the state created to record the properties of a binary gradient kernel.


// A representation of the state created to record the properties of a binary gradient kernel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNBinaryGradientStateNode
type BinaryGradientStateNode struct {
	StateNode
}

// BinaryGradientStateNodeFrom constructs a [BinaryGradientStateNode] from an unsafe.Pointer.
//
// A representation of the state created to record the properties of a binary gradient kernel.
func BinaryGradientStateNodeFrom(ptr unsafe.Pointer) BinaryGradientStateNode {
	return BinaryGradientStateNode{
		StateNode: StateNodeFrom(ptr),
	}
}

























// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnfilternode/label
func (b_ BinaryGradientStateNode) Label() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](b_.ID, objc.Sel("label"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnfilternode/label
func (b_ BinaryGradientStateNode) SetLabel(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setLabel:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnfilternode/paddingpolicy
func (b_ BinaryGradientStateNode) PaddingPolicy() Padding /* not a class type */ {
	rv := objc.Send[Padding](b_.ID, objc.Sel("paddingPolicy"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnfilternode/paddingpolicy
func (b_ BinaryGradientStateNode) SetPaddingPolicy(value Padding /* not a class type */) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setPaddingPolicy:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnfilternode/resultimage
func (b_ BinaryGradientStateNode) ResultImage() IMPSNNImageNode {
	rv := objc.Send[ImageNode](b_.ID, objc.Sel("resultImage"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnfilternode/resultimage
func (b_ BinaryGradientStateNode) SetResultImage(value IMPSNNImageNode) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setResultImage:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnfilternode/resultstate
func (b_ BinaryGradientStateNode) ResultState() IMPSNNStateNode {
	rv := objc.Send[StateNode](b_.ID, objc.Sel("resultState"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnfilternode/resultstate
func (b_ BinaryGradientStateNode) SetResultState(value IMPSNNStateNode) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setResultState:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnfilternode/resultstates
func (b_ BinaryGradientStateNode) ResultStates() IMPSNNStateNode {
	rv := objc.Send[StateNode](b_.ID, objc.Sel("resultStates"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnfilternode/resultstates
func (b_ BinaryGradientStateNode) SetResultStates(value IMPSNNStateNode) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setResultStates:"), value)
}








