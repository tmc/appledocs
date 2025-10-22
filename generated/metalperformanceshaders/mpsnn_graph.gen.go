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
	DestinationImageAllocator() objc.ID
	SetDestinationImageAllocator(value objc.ID)
	ResultStateHandles() []objc.ID
	DestinationFeatureChannelOffset() int
	SetDestinationFeatureChannelOffset(value int)
	Offset() unsafe.Pointer
	SetOffset(value unsafe.Pointer)
	Format() unsafe.Pointer
	SetFormat(value unsafe.Pointer)
	IntermediateImageHandles() unsafe.Pointer
	SetIntermediateImageHandles(value unsafe.Pointer)
	OutputStateIsTemporary() bool
	SetOutputStateIsTemporary(value bool)
	ResultHandle() unsafe.Pointer
	SetResultHandle(value unsafe.Pointer)
	ResultImageIsNeeded() bool
	SetResultImageIsNeeded(value bool)
	SourceImageHandles() unsafe.Pointer
	SetSourceImageHandles(value unsafe.Pointer)
	SourceStateHandles() unsafe.Pointer
	SetSourceStateHandles(value unsafe.Pointer)
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
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNGraph/destinationImageAllocator
func (g_ Graph) DestinationImageAllocator() objc.ID {
	rv := objc.Send[objc.ID](g_.ID, objc.Sel("destinationImageAllocator"))
	return rv
}


// SetDestinationImageAllocator sets the value of the destinationImageAllocator property.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNGraph/destinationImageAllocator
func (g_ Graph) SetDestinationImageAllocator(value objc.ID) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setDestinationImageAllocator:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNGraph/resultStateHandles
func (g_ Graph) ResultStateHandles() []objc.ID {
	rv := objc.Send[[]objc.ID](g_.ID, objc.Sel("resultStateHandles"))
	return rv
}

// The number of channels in the destination image to skip before writing output data.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnkernel/destinationfeaturechanneloffset
func (g_ Graph) DestinationFeatureChannelOffset() int {
	rv := objc.Send[int](g_.ID, objc.Sel("destinationFeatureChannelOffset"))
	return rv
}


// SetDestinationFeatureChannelOffset sets the value of the destinationFeatureChannelOffset property.
// The number of channels in the destination image to skip before writing output data.

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnkernel/destinationfeaturechanneloffset
func (g_ Graph) SetDestinationFeatureChannelOffset(value int) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setDestinationFeatureChannelOffset:"), value)
}

// The position of the destination image’s clip rectangle origin, relative to the source image.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnkernel/offset
func (g_ Graph) Offset() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("offset"))
	return rv
}


// SetOffset sets the value of the offset property.
// The position of the destination image’s clip rectangle origin, relative to the source image.

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnkernel/offset
func (g_ Graph) SetOffset(value unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setOffset:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnngraph/format
func (g_ Graph) Format() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("format"))
	return rv
}


// SetFormat sets the value of the format property.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnngraph/format
func (g_ Graph) SetFormat(value unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setFormat:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnngraph/intermediateimagehandles
func (g_ Graph) IntermediateImageHandles() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("intermediateImageHandles"))
	return rv
}


// SetIntermediateImageHandles sets the value of the intermediateImageHandles property.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnngraph/intermediateimagehandles
func (g_ Graph) SetIntermediateImageHandles(value unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setIntermediateImageHandles:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnngraph/outputstateistemporary
func (g_ Graph) OutputStateIsTemporary() bool {
	rv := objc.Send[bool](g_.ID, objc.Sel("outputStateIsTemporary"))
	return rv
}


// SetOutputStateIsTemporary sets the value of the outputStateIsTemporary property.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnngraph/outputstateistemporary
func (g_ Graph) SetOutputStateIsTemporary(value bool) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setOutputStateIsTemporary:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnngraph/resulthandle
func (g_ Graph) ResultHandle() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("resultHandle"))
	return rv
}


// SetResultHandle sets the value of the resultHandle property.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnngraph/resulthandle
func (g_ Graph) SetResultHandle(value unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setResultHandle:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnngraph/resultimageisneeded
func (g_ Graph) ResultImageIsNeeded() bool {
	rv := objc.Send[bool](g_.ID, objc.Sel("resultImageIsNeeded"))
	return rv
}


// SetResultImageIsNeeded sets the value of the resultImageIsNeeded property.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnngraph/resultimageisneeded
func (g_ Graph) SetResultImageIsNeeded(value bool) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setResultImageIsNeeded:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnngraph/sourceimagehandles
func (g_ Graph) SourceImageHandles() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("sourceImageHandles"))
	return rv
}


// SetSourceImageHandles sets the value of the sourceImageHandles property.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnngraph/sourceimagehandles
func (g_ Graph) SetSourceImageHandles(value unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setSourceImageHandles:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnngraph/sourcestatehandles
func (g_ Graph) SourceStateHandles() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("sourceStateHandles"))
	return rv
}


// SetSourceStateHandles sets the value of the sourceStateHandles property.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnngraph/sourcestatehandles
func (g_ Graph) SetSourceStateHandles(value unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setSourceStateHandles:"), value)
}



