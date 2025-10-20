// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [Graph] class.
var (
	GraphClass     _GraphClass
	GraphClassOnce sync.Once
)

func getGraphClass() _GraphClass {
	GraphClassOnce.Do(func() {
		GraphClass = _GraphClass{objc.GetClass("MPSNNGraph")}
	})
	return GraphClass
}

type _GraphClass struct {
	class objc.Class
}

// An interface definition for the [Graph] class.
type IGraph interface {
	IKernel
}

// An optimized representation of a graph of neural network image and filter nodes.
//
// Once you have prepared a graph of , , and, if needed, objects, you may initialize a using the image node that you wish to appear as the result. The graph object will introspect the graph representation and determine which nodes are needed for inputs, and which nodes are produced as output state (if any). Nodes which are not needed to calculate the result image node are ignored. Some nodes may be internally concatenated with other nodes for better performance. During construction, the graph attached to the result node will be parsed and reduced to an optimized representation. This representation may be saved using the protocol for later recall. When decoding a using a , it will be created against the system default . If you would like to set the device, your should conform to the protocol.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNGraph
type Graph struct {
	Kernel
}

// GraphFrom constructs a [Graph] from an unsafe.Pointer.
//
// An optimized representation of a graph of neural network image and filter nodes.
func GraphFrom(ptr unsafe.Pointer) Graph {
	return Graph{
		Kernel: KernelFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (gc _GraphClass) Alloc() Graph {
	rv := objc.Send[Graph](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (gc _GraphClass) New() Graph {
	rv := objc.Send[Graph](objc.ID(gc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (g_ Graph) Init() Graph {
	rv := objc.Send[Graph](g_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g_ Graph) Autorelease() Graph {
	rv := objc.Send[Graph](g_.ID, objc.Sel("autorelease"))
	return rv
}

// NewGraph creates a new Graph instance.
func NewGraph() Graph {
	return getGraphClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNGraph/resultStateHandles
func (g_ Graph) ResultStateHandles() []objc.ID {
	rv := objc.Send[[]objc.ID](g_.ID, objc.Sel("resultStateHandles"))
	return rv
}



