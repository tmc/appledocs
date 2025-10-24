// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [FilterNode] class.
var (
	FilterNodeClass     _FilterNodeClass
	FilterNodeClassOnce sync.Once
)

func getFilterNodeClass() _FilterNodeClass {
	FilterNodeClassOnce.Do(func() {
		FilterNodeClass = _FilterNodeClass{objc.GetClass("MPSNNFilterNode")}
	})
	return FilterNodeClass
}

type _FilterNodeClass struct {
	class objc.Class
}

// An interface definition for the [FilterNode] class.
type IFilterNode interface {
	objectivec.IObject
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

// A placeholder node denoting a neural network filter stage.


// A placeholder node denoting a neural network filter stage.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNFilterNode
type FilterNode struct {
	objectivec.Object
}

// FilterNodeFrom constructs a [FilterNode] from an unsafe.Pointer.
//
// A placeholder node denoting a neural network filter stage.
func FilterNodeFrom(ptr unsafe.Pointer) FilterNode {
	return FilterNode{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (fc _FilterNodeClass) Alloc() FilterNode {
	rv := objc.Send[FilterNode](objc.ID(fc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (fc _FilterNodeClass) New() FilterNode {
	rv := objc.Send[FilterNode](objc.ID(fc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (f_ FilterNode) Init() FilterNode {
	rv := objc.Send[FilterNode](f_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f_ FilterNode) Autorelease() FilterNode {
	rv := objc.Send[FilterNode](f_.ID, objc.Sel("autorelease"))
	return rv
}

// NewFilterNode creates a new FilterNode instance.
func NewFilterNode() FilterNode {
	return getFilterNodeClass().New()
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnfilternode/label
func (f_ FilterNode) Label() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](f_.ID, objc.Sel("label"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnfilternode/label
func (f_ FilterNode) SetLabel(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setLabel:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnfilternode/paddingpolicy
func (f_ FilterNode) PaddingPolicy() Padding /* not a class type */ {
	rv := objc.Send[Padding](f_.ID, objc.Sel("paddingPolicy"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnfilternode/paddingpolicy
func (f_ FilterNode) SetPaddingPolicy(value Padding /* not a class type */) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setPaddingPolicy:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnfilternode/resultimage
func (f_ FilterNode) ResultImage() IMPSNNImageNode {
	rv := objc.Send[ImageNode](f_.ID, objc.Sel("resultImage"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnfilternode/resultimage
func (f_ FilterNode) SetResultImage(value IMPSNNImageNode) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setResultImage:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnfilternode/resultstate
func (f_ FilterNode) ResultState() IMPSNNStateNode {
	rv := objc.Send[StateNode](f_.ID, objc.Sel("resultState"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnfilternode/resultstate
func (f_ FilterNode) SetResultState(value IMPSNNStateNode) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setResultState:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnfilternode/resultstates
func (f_ FilterNode) ResultStates() IMPSNNStateNode {
	rv := objc.Send[StateNode](f_.ID, objc.Sel("resultStates"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnfilternode/resultstates
func (f_ FilterNode) SetResultStates(value IMPSNNStateNode) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setResultStates:"), value)
}



