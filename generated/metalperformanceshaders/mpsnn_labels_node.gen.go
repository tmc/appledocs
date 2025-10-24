// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)





// The class instance for the [LabelsNode] class.
var (
	LabelsNodeClass     _LabelsNodeClass
	LabelsNodeClassOnce sync.Once
)

func getLabelsNodeClass() _LabelsNodeClass {
	LabelsNodeClassOnce.Do(func() {
		LabelsNodeClass = _LabelsNodeClass{objc.GetClass("MPSNNLabelsNode")}
	})
	return LabelsNodeClass
}

type _LabelsNodeClass struct {
	class objc.Class
}





// An interface definition for the [LabelsNode] class.
type ILabelsNode interface {
	IStateNode
	

	// properties:


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (lc _LabelsNodeClass) Alloc() LabelsNode {
	rv := objc.Send[LabelsNode](objc.ID(lc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (lc _LabelsNodeClass) New() LabelsNode {
	rv := objc.Send[LabelsNode](objc.ID(lc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (l_ LabelsNode) Init() LabelsNode {
	rv := objc.Send[LabelsNode](l_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (l_ LabelsNode) Autorelease() LabelsNode {
	rv := objc.Send[LabelsNode](l_.ID, objc.Sel("autorelease"))
	return rv
}

// NewLabelsNode creates a new LabelsNode instance.
func NewLabelsNode() LabelsNode {
	return getLabelsNodeClass().New()
}





// A placeholder node denoting the per-element weight buffer used by loss and gradient loss kernels.


// A placeholder node denoting the per-element weight buffer used by loss and gradient loss kernels.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNLabelsNode
type LabelsNode struct {
	StateNode
}

// LabelsNodeFrom constructs a [LabelsNode] from an unsafe.Pointer.
//
// A placeholder node denoting the per-element weight buffer used by loss and gradient loss kernels.
func LabelsNodeFrom(ptr unsafe.Pointer) LabelsNode {
	return LabelsNode{
		StateNode: StateNodeFrom(ptr),
	}
}































