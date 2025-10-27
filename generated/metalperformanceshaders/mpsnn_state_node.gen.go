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
	

	// properties:
	Handle() Handle get set /* not a class type */
	SetHandle(value Handle get set /* not a class type */)
	SynchronizeResource() objectivec.IObject
	SetSynchronizeResource(value objectivec.IObject)
	ExportFromGraph() objectivec.IObject
	SetExportFromGraph(value objectivec.IObject)
	Label() foundation.foundation.INSString
	SetLabel(value foundation.foundation.INSString)
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
func (sc _StateNodeClass) Alloc() StateNode {
	rv := objc.Send[StateNode](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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





// A placeholder node denoting the position in the graph of a state object.


// A placeholder node denoting the position in the graph of a state object.
//
// [Full Topic]
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

























// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnstatenode/2866426-handle
func (s_ StateNode) Handle() Handle get set /* not a class type */ {
	rv := objc.Send[objc.ID](s_.ID, objc.Sel("handle"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnstatenode/2866426-handle
func (s_ StateNode) SetHandle(value Handle get set /* not a class type */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setHandle:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnstatenode/2942639-synchronizeresource
func (s_ StateNode) SynchronizeResource() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](s_.ID, objc.Sel("synchronizeResource"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnstatenode/2942639-synchronizeresource
func (s_ StateNode) SetSynchronizeResource(value objectivec.IObject) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setSynchronizeResource:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnstatenode/2942640-exportfromgraph
func (s_ StateNode) ExportFromGraph() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](s_.ID, objc.Sel("exportFromGraph"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnstatenode/2942640-exportfromgraph
func (s_ StateNode) SetExportFromGraph(value objectivec.IObject) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setExportFromGraph:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnfilternode/label
func (s_ StateNode) Label() foundation.foundation.INSString {
	rv := objc.Send[foundation.NSString](s_.ID, objc.Sel("label"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnfilternode/label
func (s_ StateNode) SetLabel(value foundation.foundation.INSString) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setLabel:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnfilternode/paddingpolicy
func (s_ StateNode) PaddingPolicy() Padding /* not a class type */ {
	rv := objc.Send[Padding](s_.ID, objc.Sel("paddingPolicy"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnfilternode/paddingpolicy
func (s_ StateNode) SetPaddingPolicy(value Padding /* not a class type */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setPaddingPolicy:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnfilternode/resultimage
func (s_ StateNode) ResultImage() IMPSNNImageNode {
	rv := objc.Send[ImageNode](s_.ID, objc.Sel("resultImage"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnfilternode/resultimage
func (s_ StateNode) SetResultImage(value IMPSNNImageNode) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setResultImage:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnfilternode/resultstate
func (s_ StateNode) ResultState() IMPSNNStateNode {
	rv := objc.Send[StateNode](s_.ID, objc.Sel("resultState"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnfilternode/resultstate
func (s_ StateNode) SetResultState(value IMPSNNStateNode) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setResultState:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnfilternode/resultstates
func (s_ StateNode) ResultStates() IMPSNNStateNode {
	rv := objc.Send[StateNode](s_.ID, objc.Sel("resultStates"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnfilternode/resultstates
func (s_ StateNode) SetResultStates(value IMPSNNStateNode) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setResultStates:"), value)
}








