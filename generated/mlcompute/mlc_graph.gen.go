// Code generated from Apple documentation for MLCompute. DO NOT EDIT.

package mlcompute

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
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
	BindAndWriteDataForInputsToDeviceBatchSizeSynchronous(inputsData unsafe.Pointer, inputTensors unsafe.Pointer, device unsafe.Pointer, batchSize uint, synchronous bool) bool
	BindAndWriteDataForInputsToDeviceSynchronous(inputsData unsafe.Pointer, inputTensors unsafe.Pointer, device unsafe.Pointer, synchronous bool) bool
	ConcatenateWithSourcesDimension(sources unsafe.Pointer, dimension uint) unsafe.Pointer
	GatherWithDimensionSourceIndices(dimension uint, source unsafe.Pointer, indices unsafe.Pointer) unsafe.Pointer
	NodeWithLayerSource(layer unsafe.Pointer, source unsafe.Pointer) unsafe.Pointer
	NodeWithLayerSources(layer unsafe.Pointer, sources unsafe.Pointer) unsafe.Pointer
	NodeWithLayerSourcesDisableUpdate(layer unsafe.Pointer, sources unsafe.Pointer, disableUpdate bool) unsafe.Pointer
	NodeWithLayerSourcesLossLabels(layer unsafe.Pointer, sources unsafe.Pointer, lossLabels unsafe.Pointer) unsafe.Pointer
	ReshapeWithShapeSource(shape unsafe.Pointer, source unsafe.Pointer) unsafe.Pointer
	ResultTensorsForLayer(layer unsafe.Pointer) []CTensor
	ScatterWithDimensionSourceIndicesCopyFromReductionType(dimension uint, source unsafe.Pointer, indices unsafe.Pointer, copyFrom unsafe.Pointer, reductionType unsafe.Pointer) unsafe.Pointer
	SelectWithSourcesCondition(sources unsafe.Pointer, condition unsafe.Pointer) unsafe.Pointer
	SourceTensorsForLayer(layer unsafe.Pointer) []CTensor
	SplitWithSourceSplitCountDimension(source unsafe.Pointer, splitCount uint, dimension uint) []CTensor
	SplitWithSourceSplitSectionLengthsDimension(source unsafe.Pointer, splitSectionLengths unsafe.Pointer, dimension uint) []CTensor
	TransposeWithDimensionsSource(dimensions unsafe.Pointer, source unsafe.Pointer) unsafe.Pointer
}

// A graph of layers you use to build a training or inference graph.
//
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


// Creates a new graph.
//
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCGraph/graph
func (cc _CGraphClass) Graph() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("graph"))
	return rv
}

// Associates the given data with the input tensors, and if the device is a GPU, also copies the data to the device memory.
//
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCGraph/bindAndWriteData(_:forInputs:to:batchSize:synchronous:)
func (c_ CGraph) BindAndWriteDataForInputsToDeviceBatchSizeSynchronous(inputsData unsafe.Pointer, inputTensors unsafe.Pointer, device unsafe.Pointer, batchSize uint, synchronous bool) bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("bindAndWriteData:forInputs:toDevice:batchSize:synchronous:"), inputsData, inputTensors, device, batchSize, synchronous)
	return rv
}

// Associates the given data with the input tensors, and if the device is a GPU, also copies the data to the device memory.
//
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCGraph/bindAndWriteData(_:forInputs:to:synchronous:)
func (c_ CGraph) BindAndWriteDataForInputsToDeviceSynchronous(inputsData unsafe.Pointer, inputTensors unsafe.Pointer, device unsafe.Pointer, synchronous bool) bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("bindAndWriteData:forInputs:toDevice:synchronous:"), inputsData, inputTensors, device, synchronous)
	return rv
}

// Adds a new concatenation layer to the graph using the source tensors and concatenation dimension you specify.
//
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCGraph/concatenate(sources:dimension:)
func (c_ CGraph) ConcatenateWithSourcesDimension(sources unsafe.Pointer, dimension uint) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("concatenateWithSources:dimension:"), sources, dimension)
	return rv
}

// Adds a gather layer to the graph using the source tensor, dimension along which to index, and the indices you specify.
//
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCGraph/gather(withDimension:source:indices:)
func (c_ CGraph) GatherWithDimensionSourceIndices(dimension uint, source unsafe.Pointer, indices unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("gatherWithDimension:source:indices:"), dimension, source, indices)
	return rv
}

// Adds the layer and source tensor that you specify to the graph.
//
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCGraph/node(with:source:)
func (c_ CGraph) NodeWithLayerSource(layer unsafe.Pointer, source unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("nodeWithLayer:source:"), layer, source)
	return rv
}

// Adds the layer and source tensors that you specify to the graph.
//
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCGraph/node(with:sources:)
func (c_ CGraph) NodeWithLayerSources(layer unsafe.Pointer, sources unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("nodeWithLayer:sources:"), layer, sources)
	return rv
}

// Adds the layer, source tensors, and option to disable optimizer updates that you specify to the graph.
//
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCGraph/node(with:sources:disableUpdate:)
func (c_ CGraph) NodeWithLayerSourcesDisableUpdate(layer unsafe.Pointer, sources unsafe.Pointer, disableUpdate bool) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("nodeWithLayer:sources:disableUpdate:"), layer, sources, disableUpdate)
	return rv
}

// Adds the layer, sources, and loss labels tensors that you specify to the graph.
//
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCGraph/node(with:sources:lossLabels:)
func (c_ CGraph) NodeWithLayerSourcesLossLabels(layer unsafe.Pointer, sources unsafe.Pointer, lossLabels unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("nodeWithLayer:sources:lossLabels:"), layer, sources, lossLabels)
	return rv
}

// Adds a new reshape layer to the graph using the shape and source tensor you specify.
//
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCGraph/reshapeWithShape:source:
func (c_ CGraph) ReshapeWithShapeSource(shape unsafe.Pointer, source unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("reshapeWithShape:source:"), shape, source)
	return rv
}

// Gets the result tensors for a layer in the training graph.
//
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCGraph/resultTensors(for:)
func (c_ CGraph) ResultTensorsForLayer(layer unsafe.Pointer) []CTensor {
	rv := objc.Send[[]CTensor](c_.ID, objc.Sel("resultTensorsForLayer:"), layer)
	return rv
}

// Adds a scatter layer to the graph.
//
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCGraph/scatter(withDimension:source:indices:copyFrom:reductionType:)
func (c_ CGraph) ScatterWithDimensionSourceIndicesCopyFromReductionType(dimension uint, source unsafe.Pointer, indices unsafe.Pointer, copyFrom unsafe.Pointer, reductionType unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("scatterWithDimension:source:indices:copyFrom:reductionType:"), dimension, source, indices, copyFrom, reductionType)
	return rv
}

// Adds a select layer to the graph using the condition mask and source tensors you specify.
//
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCGraph/selectWithSources:condition:
func (c_ CGraph) SelectWithSourcesCondition(sources unsafe.Pointer, condition unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("selectWithSources:condition:"), sources, condition)
	return rv
}

// Gets the source tensors for a layer in the training graph.
//
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCGraph/sourceTensors(for:)
func (c_ CGraph) SourceTensorsForLayer(layer unsafe.Pointer) []CTensor {
	rv := objc.Send[[]CTensor](c_.ID, objc.Sel("sourceTensorsForLayer:"), layer)
	return rv
}

// Adds a new split layer to the graph using the source tensor, number of splits, and dimension to split the source tensor that you specify.
//
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCGraph/split(source:splitCount:dimension:)
func (c_ CGraph) SplitWithSourceSplitCountDimension(source unsafe.Pointer, splitCount uint, dimension uint) []CTensor {
	rv := objc.Send[[]CTensor](c_.ID, objc.Sel("splitWithSource:splitCount:dimension:"), source, splitCount, dimension)
	return rv
}

// Adds a new split layer to the graph using the source tensor, lengths of each split section, and dimension to split the source tensor that you specify.
//
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCGraph/splitWithSource:splitSectionLengths:dimension:
func (c_ CGraph) SplitWithSourceSplitSectionLengthsDimension(source unsafe.Pointer, splitSectionLengths unsafe.Pointer, dimension uint) []CTensor {
	rv := objc.Send[[]CTensor](c_.ID, objc.Sel("splitWithSource:splitSectionLengths:dimension:"), source, splitSectionLengths, dimension)
	return rv
}

// Adds a new transpose layer to the graph using the dimensions and source tensor you specify.
//
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCGraph/transposeWithDimensions:source:
func (c_ CGraph) TransposeWithDimensionsSource(dimensions unsafe.Pointer, source unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("transposeWithDimensions:source:"), dimensions, source)
	return rv
}

// The device you’ll use for compiling and executing a graph.
//
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCGraph/device
func (c_ CGraph) Device() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("device"))
	return rv
}

// An array that contains the layers in the graph.
//
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCGraph/layers
func (c_ CGraph) Layers() []CLayer {
	rv := objc.Send[[]CLayer](c_.ID, objc.Sel("layers"))
	return rv
}

// A DOT representation of the graph.
//
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCGraph/summarizedDOTDescription
func (c_ CGraph) SummarizedDOTDescription() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("summarizedDOTDescription"))
	return rv
}



