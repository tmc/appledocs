// Code generated from Apple documentation for MLCompute. DO NOT EDIT.

package mlcompute

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [CGraph] class.
var (
	CGraphClass     _CGraphClass
	CGraphClassOnce sync.Once
)

func getCGraphClass() _CGraphClass {
	CGraphClassOnce.Do(func() {
		CGraphClass = _CGraphClass{objc.GetClass("MLCGraph")}
	})
	return CGraphClass
}

type _CGraphClass struct {
	class objc.Class
}

// An interface definition for the [CGraph] class.
type ICGraph interface {
	objectivec.IObject
	// properties:
	Device() IMLCDevice
	SetDevice(value IMLCDevice)
	Layers() IMLCLayer
	SetLayers(value IMLCLayer)
	SummarizedDOTDescription() objc.IObject /* cross-framework: NSString */
	SetSummarizedDOTDescription(value objc.IObject /* cross-framework: NSString */)
	// methods:
}

// A graph of layers you use to build a training or inference graph.


// A graph of layers you use to build a training or inference graph.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCGraph
type CGraph struct {
	objectivec.Object
}

// CGraphFrom constructs a [CGraph] from an unsafe.Pointer.
//
// A graph of layers you use to build a training or inference graph.
func CGraphFrom(ptr unsafe.Pointer) CGraph {
	return CGraph{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (cc _CGraphClass) Alloc() CGraph {
	rv := objc.Send[CGraph](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CGraphClass) New() CGraph {
	rv := objc.Send[CGraph](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CGraph) Init() CGraph {
	rv := objc.Send[CGraph](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CGraph) Autorelease() CGraph {
	rv := objc.Send[CGraph](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCGraph creates a new CGraph instance.
func NewCGraph() CGraph {
	return getCGraphClass().New()
}



// The device you’ll use for compiling and executing a graph.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlcgraph/device
func (c_ CGraph) Device() IMLCDevice {
	rv := objc.Send[CDevice](c_.ID, objc.Sel("device"))
	return rv
}


// The device you’ll use for compiling and executing a graph.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlcgraph/device
func (c_ CGraph) SetDevice(value IMLCDevice) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setDevice:"), value)
}


// An array that contains the layers in the graph.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlcgraph/layers
func (c_ CGraph) Layers() IMLCLayer {
	rv := objc.Send[CLayer](c_.ID, objc.Sel("layers"))
	return rv
}


// An array that contains the layers in the graph.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlcgraph/layers
func (c_ CGraph) SetLayers(value IMLCLayer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setLayers:"), value)
}


// A DOT representation of the graph.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlcgraph/summarizeddotdescription
func (c_ CGraph) SummarizedDOTDescription() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("summarizedDOTDescription"))
	return rv
}


// A DOT representation of the graph.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlcgraph/summarizeddotdescription
func (c_ CGraph) SetSummarizedDOTDescription(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSummarizedDOTDescription:"), value)
}



