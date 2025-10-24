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
	// properties:
	DestinationFeatureChannelOffset() int
	SetDestinationFeatureChannelOffset(value int)
	Offset() MPSOffset /* not a class type */
	SetOffset(value MPSOffset /* not a class type */)
	DestinationImageAllocator() ImageAllocator /* not a class type */
	SetDestinationImageAllocator(value ImageAllocator /* not a class type */)
	Format() ImageFeatureChannelFormat /* not a class type */
	SetFormat(value ImageFeatureChannelFormat /* not a class type */)
	IntermediateImageHandles() Handle /* not a class type */
	SetIntermediateImageHandles(value Handle /* not a class type */)
	OutputStateIsTemporary() bool
	SetOutputStateIsTemporary(value bool)
	ResultHandle() Handle /* not a class type */
	SetResultHandle(value Handle /* not a class type */)
	ResultImageIsNeeded() bool
	SetResultImageIsNeeded(value bool)
	ResultStateHandles() Handle /* not a class type */
	SetResultStateHandles(value Handle /* not a class type */)
	SourceImageHandles() Handle /* not a class type */
	SetSourceImageHandles(value Handle /* not a class type */)
	SourceStateHandles() Handle /* not a class type */
	SetSourceStateHandles(value Handle /* not a class type */)
	// methods:
}

// An optimized representation of a graph of neural network image and filter nodes.
//
// Once you have prepared a graph of , , and, if needed, objects, you may initialize a using the image node that you wish to appear as the result. The graph object will introspect the graph representation and determine which nodes are needed for inputs, and which nodes are produced as output state (if any). Nodes which are not needed to calculate the result image node are ignored. Some nodes may be internally concatenated with other nodes for better performance. During construction, the graph attached to the result node will be parsed and reduced to an optimized representation. This representation may be saved using the protocol for later recall. When decoding a using a , it will be created against the system default . If you would like to set the device, your should conform to the protocol.


// An optimized representation of a graph of neural network image and filter nodes.
//
// [Full Topic]
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



// The number of channels in the destination image to skip before writing output data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnkernel/destinationfeaturechanneloffset
func (g_ Graph) DestinationFeatureChannelOffset() int {
	rv := objc.Send[int](g_.ID, objc.Sel("destinationFeatureChannelOffset"))
	return rv
}


// The number of channels in the destination image to skip before writing output data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnkernel/destinationfeaturechanneloffset
func (g_ Graph) SetDestinationFeatureChannelOffset(value int) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setDestinationFeatureChannelOffset:"), value)
}


// The position of the destination image’s clip rectangle origin, relative to the source image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnkernel/offset
func (g_ Graph) Offset() MPSOffset /* not a class type */ {
	rv := objc.Send[Offset](g_.ID, objc.Sel("offset"))
	return rv
}


// The position of the destination image’s clip rectangle origin, relative to the source image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnkernel/offset
func (g_ Graph) SetOffset(value MPSOffset /* not a class type */) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setOffset:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnngraph/destinationimageallocator
func (g_ Graph) DestinationImageAllocator() ImageAllocator /* not a class type */ {
	rv := objc.Send[ImageAllocator](g_.ID, objc.Sel("destinationImageAllocator"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnngraph/destinationimageallocator
func (g_ Graph) SetDestinationImageAllocator(value ImageAllocator /* not a class type */) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setDestinationImageAllocator:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnngraph/format
func (g_ Graph) Format() ImageFeatureChannelFormat /* not a class type */ {
	rv := objc.Send[ImageFeatureChannelFormat](g_.ID, objc.Sel("format"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnngraph/format
func (g_ Graph) SetFormat(value ImageFeatureChannelFormat /* not a class type */) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setFormat:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnngraph/intermediateimagehandles
func (g_ Graph) IntermediateImageHandles() Handle /* not a class type */ {
	rv := objc.Send[Handle](g_.ID, objc.Sel("intermediateImageHandles"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnngraph/intermediateimagehandles
func (g_ Graph) SetIntermediateImageHandles(value Handle /* not a class type */) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setIntermediateImageHandles:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnngraph/outputstateistemporary
func (g_ Graph) OutputStateIsTemporary() bool {
	rv := objc.Send[bool](g_.ID, objc.Sel("outputStateIsTemporary"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnngraph/outputstateistemporary
func (g_ Graph) SetOutputStateIsTemporary(value bool) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setOutputStateIsTemporary:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnngraph/resulthandle
func (g_ Graph) ResultHandle() Handle /* not a class type */ {
	rv := objc.Send[Handle](g_.ID, objc.Sel("resultHandle"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnngraph/resulthandle
func (g_ Graph) SetResultHandle(value Handle /* not a class type */) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setResultHandle:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnngraph/resultimageisneeded
func (g_ Graph) ResultImageIsNeeded() bool {
	rv := objc.Send[bool](g_.ID, objc.Sel("resultImageIsNeeded"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnngraph/resultimageisneeded
func (g_ Graph) SetResultImageIsNeeded(value bool) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setResultImageIsNeeded:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnngraph/resultstatehandles
func (g_ Graph) ResultStateHandles() Handle /* not a class type */ {
	rv := objc.Send[Handle](g_.ID, objc.Sel("resultStateHandles"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnngraph/resultstatehandles
func (g_ Graph) SetResultStateHandles(value Handle /* not a class type */) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setResultStateHandles:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnngraph/sourceimagehandles
func (g_ Graph) SourceImageHandles() Handle /* not a class type */ {
	rv := objc.Send[Handle](g_.ID, objc.Sel("sourceImageHandles"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnngraph/sourceimagehandles
func (g_ Graph) SetSourceImageHandles(value Handle /* not a class type */) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setSourceImageHandles:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnngraph/sourcestatehandles
func (g_ Graph) SourceStateHandles() Handle /* not a class type */ {
	rv := objc.Send[Handle](g_.ID, objc.Sel("sourceStateHandles"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnngraph/sourcestatehandles
func (g_ Graph) SetSourceStateHandles(value Handle /* not a class type */) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setSourceStateHandles:"), value)
}



