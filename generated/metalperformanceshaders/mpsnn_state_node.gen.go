// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [StateNode] class.
var (
	StateNodeClass     _StateNodeClass
	StateNodeClassOnce sync.Once
)

func getStateNodeClass() _StateNodeClass {
	StateNodeClassOnce.Do(func() {
		StateNodeClass = _StateNodeClass{objc.GetClass("MPSNNStateNode")}
	})
	return StateNodeClass
}

type _StateNodeClass struct {
	class objc.Class
}

// An interface definition for the [StateNode] class.
type IStateNode interface {
	objectivec.IObject
}

// A placeholder node denoting the position in the graph of a state object.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNStateNode
type StateNode struct {
	objectivec.Object
}

// StateNodeFrom constructs a [StateNode] from an unsafe.Pointer.
//
// A placeholder node denoting the position in the graph of a state object.
func StateNodeFrom(ptr unsafe.Pointer) StateNode {
	return StateNode{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (sc _StateNodeClass) Alloc() StateNode {
	rv := objc.Send[StateNode](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (sc _StateNodeClass) New() StateNode {
	rv := objc.Send[StateNode](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ StateNode) Init() StateNode {
	rv := objc.Send[StateNode](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ StateNode) Autorelease() StateNode {
	rv := objc.Send[StateNode](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewStateNode creates a new StateNode instance.
func NewStateNode() StateNode {
	return getStateNodeClass().New()
}
