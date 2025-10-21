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


//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnfilternode/resultstates
func (s_ StateNode) ResultStates() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("resultStates"))
	return rv
}


// SetResultStates sets the value of the resultStates property.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnfilternode/resultstates
func (s_ StateNode) SetResultStates(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setResultStates:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnstatenode/synchronizeresource
func (s_ StateNode) SynchronizeResource() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("synchronizeResource"))
	return rv
}


// SetSynchronizeResource sets the value of the synchronizeResource property.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnstatenode/synchronizeresource
func (s_ StateNode) SetSynchronizeResource(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setSynchronizeResource:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnstatenode/handle
func (s_ StateNode) Handle() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("handle"))
	return rv
}


// SetHandle sets the value of the handle property.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnstatenode/handle
func (s_ StateNode) SetHandle(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setHandle:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnstatenode/exportfromgraph
func (s_ StateNode) ExportFromGraph() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("exportFromGraph"))
	return rv
}


// SetExportFromGraph sets the value of the exportFromGraph property.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnstatenode/exportfromgraph
func (s_ StateNode) SetExportFromGraph(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setExportFromGraph:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnfilternode/resultimage
func (s_ StateNode) ResultImage() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("resultImage"))
	return rv
}


// SetResultImage sets the value of the resultImage property.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnfilternode/resultimage
func (s_ StateNode) SetResultImage(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setResultImage:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnfilternode/resultstate
func (s_ StateNode) ResultState() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("resultState"))
	return rv
}


// SetResultState sets the value of the resultState property.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnfilternode/resultstate
func (s_ StateNode) SetResultState(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setResultState:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnfilternode/paddingpolicy
func (s_ StateNode) PaddingPolicy() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("paddingPolicy"))
	return rv
}


// SetPaddingPolicy sets the value of the paddingPolicy property.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnfilternode/paddingpolicy
func (s_ StateNode) SetPaddingPolicy(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setPaddingPolicy:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnfilternode/label
func (s_ StateNode) Label() string {
	rv := objc.Send[string](s_.ID, objc.Sel("label"))
	return rv
}


// SetLabel sets the value of the label property.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnfilternode/label
func (s_ StateNode) SetLabel(value string) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setLabel:"), objc.String(value))
}



