// Code generated from Apple documentation for MetalPerformanceShadersGraph. DO NOT EDIT.

package metalperformanceshadersgraph

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [Graph] class.
var (
	GraphClass     _GraphClass
	GraphClassOnce sync.Once
)

func getGraphClass() _GraphClass {
	GraphClassOnce.Do(func() {
		GraphClass = _GraphClass{objc.GetClass("MPSGraph")}
	})
	return GraphClass
}

type _GraphClass struct {
	class objc.Class
}

// An interface definition for the [Graph] class.
type IGraph interface {
	IGraphObject
	GRUWithSourceTensorRecurrentWeightInputWeightBiasDescriptorName(source IMPSGraphTensor, recurrentWeight IMPSGraphTensor, inputWeight IMPSGraphTensor, bias IMPSGraphTensor, descriptor IMPSGraphGRUDescriptor, name appkit.string) []GraphTensor
	GRUGradientsWithSourceTensorRecurrentWeightSourceGradientZStateOutputFwdInputWeightBiasDescriptorName(source IMPSGraphTensor, recurrentWeight IMPSGraphTensor, sourceGradient IMPSGraphTensor, zState IMPSGraphTensor, outputFwd IMPSGraphTensor, inputWeight IMPSGraphTensor, bias IMPSGraphTensor, descriptor IMPSGraphGRUDescriptor, name appkit.string) []GraphTensor
	HermiteanToRealFFTWithTensorAxesTensorDescriptorName(tensor IMPSGraphTensor, axesTensor IMPSGraphTensor, descriptor IMPSGraphFFTDescriptor, name appkit.string) GraphTensor
	LSTMWithSourceTensorRecurrentWeightInitStateInitCellDescriptorName(source IMPSGraphTensor, recurrentWeight IMPSGraphTensor, initState IMPSGraphTensor, initCell IMPSGraphTensor, descriptor IMPSGraphLSTMDescriptor, name appkit.string) []GraphTensor
	ColToImWithSourceTensorOutputShapeDescriptorName(source IMPSGraphTensor, outputShape unsafe.Pointer, descriptor IMPSGraphImToColOpDescriptor, name appkit.string) GraphTensor
	DepthwiseConvolution3DWithSourceTensorWeightsTensorDescriptorName(source IMPSGraphTensor, weights IMPSGraphTensor, descriptor IMPSGraphDepthwiseConvolution3DOpDescriptor, name appkit.string) GraphTensor
	DepthwiseConvolution3DDataGradientWithIncomingGradientTensorWeightsTensorOutputShapeDescriptorName(incomingGradient IMPSGraphTensor, weights IMPSGraphTensor, outputShape unsafe.Pointer, descriptor IMPSGraphDepthwiseConvolution3DOpDescriptor, name appkit.string) GraphTensor
	FastFourierTransformWithTensorAxesDescriptorName(tensor IMPSGraphTensor, axes []foundation.INumber, descriptor IMPSGraphFFTDescriptor, name appkit.string) GraphTensor
	IdentityWithTensorName(tensor IMPSGraphTensor, name appkit.string) GraphTensor
	ImToColWithSourceTensorDescriptorName(source IMPSGraphTensor, descriptor IMPSGraphImToColOpDescriptor, name appkit.string) GraphTensor
	MaxPooling2DWithSourceTensorDescriptorName(source IMPSGraphTensor, descriptor IMPSGraphPooling2DOpDescriptor, name appkit.string) GraphTensor
	MaxPooling2DGradientWithGradientTensorSourceTensorDescriptorName(gradient IMPSGraphTensor, source IMPSGraphTensor, descriptor IMPSGraphPooling2DOpDescriptor, name appkit.string) GraphTensor
	MaxPooling2DReturnIndicesWithSourceTensorDescriptorName(source IMPSGraphTensor, descriptor IMPSGraphPooling2DOpDescriptor, name appkit.string) []GraphTensor
	MaxPooling4DWithSourceTensorDescriptorName(source IMPSGraphTensor, descriptor IMPSGraphPooling4DOpDescriptor, name appkit.string) GraphTensor
	MaxPooling4DReturnIndicesWithSourceTensorDescriptorName(source IMPSGraphTensor, descriptor IMPSGraphPooling4DOpDescriptor, name appkit.string) []GraphTensor
	RealToHermiteanFFTWithTensorAxesTensorDescriptorName(tensor IMPSGraphTensor, axesTensor IMPSGraphTensor, descriptor IMPSGraphFFTDescriptor, name appkit.string) GraphTensor
	SingleGateRNNWithSourceTensorRecurrentWeightInitStateDescriptorName(source IMPSGraphTensor, recurrentWeight IMPSGraphTensor, initState IMPSGraphTensor, descriptor IMPSGraphSingleGateRNNDescriptor, name appkit.string) []GraphTensor
	SingleGateRNNWithSourceTensorRecurrentWeightInputWeightBiasInitStateMaskDescriptorName(source IMPSGraphTensor, recurrentWeight IMPSGraphTensor, inputWeight IMPSGraphTensor, bias IMPSGraphTensor, initState IMPSGraphTensor, mask IMPSGraphTensor, descriptor IMPSGraphSingleGateRNNDescriptor, name appkit.string) []GraphTensor
}

// The optimized representation of a compute graph of operations and tensors.
//
// An MPSGraph is a symbolic representation of operations to be utilized to execute compute graphs on a device.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph
type Graph struct {
	GraphObject
}

// GraphFrom constructs a [Graph] from an unsafe.Pointer.
//
// The optimized representation of a compute graph of operations and tensors.
func GraphFrom(ptr unsafe.Pointer) Graph {
	return Graph{
		GraphObject: GraphObjectFrom(ptr),
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


// Creates a GRU operation and returns the value and optionally the training state tensor.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/GRU(_:recurrentWeight:inputWeight:bias:descriptor:name:)
func (g_ Graph) GRUWithSourceTensorRecurrentWeightInputWeightBiasDescriptorName(source IMPSGraphTensor, recurrentWeight IMPSGraphTensor, inputWeight IMPSGraphTensor, bias IMPSGraphTensor, descriptor IMPSGraphGRUDescriptor, name appkit.string) []GraphTensor {
	rv := objc.Send[[]GraphTensor](g_.ID, objc.Sel("GRUWithSourceTensor:recurrentWeight:inputWeight:bias:descriptor:name:"), source, recurrentWeight, inputWeight, bias, descriptor, name)
	return rv
}

// Creates a GRU gradient operation and returns the gradient tensor values.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/GRUGradients(_:recurrentWeight:sourceGradient:zState:outputFwd:inputWeight:bias:descriptor:name:)
func (g_ Graph) GRUGradientsWithSourceTensorRecurrentWeightSourceGradientZStateOutputFwdInputWeightBiasDescriptorName(source IMPSGraphTensor, recurrentWeight IMPSGraphTensor, sourceGradient IMPSGraphTensor, zState IMPSGraphTensor, outputFwd IMPSGraphTensor, inputWeight IMPSGraphTensor, bias IMPSGraphTensor, descriptor IMPSGraphGRUDescriptor, name appkit.string) []GraphTensor {
	rv := objc.Send[[]GraphTensor](g_.ID, objc.Sel("GRUGradientsWithSourceTensor:recurrentWeight:sourceGradient:zState:outputFwd:inputWeight:bias:descriptor:name:"), source, recurrentWeight, sourceGradient, zState, outputFwd, inputWeight, bias, descriptor, name)
	return rv
}

// Creates a Hermitean-to-real fast Fourier transform operation and returns the result tensor.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/HermiteanToRealFFT(_:axesTensor:descriptor:name:)
func (g_ Graph) HermiteanToRealFFTWithTensorAxesTensorDescriptorName(tensor IMPSGraphTensor, axesTensor IMPSGraphTensor, descriptor IMPSGraphFFTDescriptor, name appkit.string) GraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("HermiteanToRealFFTWithTensor:axesTensor:descriptor:name:"), tensor, axesTensor, descriptor, name)
	return rv
}

// Creates an LSTM operation and returns the value tensor and optionally the cell state tensor and the training state tensor.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/LSTM(_:recurrentWeight:initState:initCell:descriptor:name:)
func (g_ Graph) LSTMWithSourceTensorRecurrentWeightInitStateInitCellDescriptorName(source IMPSGraphTensor, recurrentWeight IMPSGraphTensor, initState IMPSGraphTensor, initCell IMPSGraphTensor, descriptor IMPSGraphLSTMDescriptor, name appkit.string) []GraphTensor {
	rv := objc.Send[[]GraphTensor](g_.ID, objc.Sel("LSTMWithSourceTensor:recurrentWeight:initState:initCell:descriptor:name:"), source, recurrentWeight, initState, initCell, descriptor, name)
	return rv
}

// Creates a column to image operation and returns the result tensor.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/colToIm(_:outputShape:descriptor:name:)
func (g_ Graph) ColToImWithSourceTensorOutputShapeDescriptorName(source IMPSGraphTensor, outputShape unsafe.Pointer, descriptor IMPSGraphImToColOpDescriptor, name appkit.string) GraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("colToImWithSourceTensor:outputShape:descriptor:name:"), source, outputShape, descriptor, name)
	return rv
}

// Creates a 3D depthwise convolution operation and returns the result tensor.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/depthwiseConvolution3D(_:weights:descriptor:name:)
func (g_ Graph) DepthwiseConvolution3DWithSourceTensorWeightsTensorDescriptorName(source IMPSGraphTensor, weights IMPSGraphTensor, descriptor IMPSGraphDepthwiseConvolution3DOpDescriptor, name appkit.string) GraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("depthwiseConvolution3DWithSourceTensor:weightsTensor:descriptor:name:"), source, weights, descriptor, name)
	return rv
}

// Creates a 3D depthwise convolution gradient for data operation and returns the result tensor.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/depthwiseConvolution3DDataGradient(_:weights:outputShape:descriptor:name:)
func (g_ Graph) DepthwiseConvolution3DDataGradientWithIncomingGradientTensorWeightsTensorOutputShapeDescriptorName(incomingGradient IMPSGraphTensor, weights IMPSGraphTensor, outputShape unsafe.Pointer, descriptor IMPSGraphDepthwiseConvolution3DOpDescriptor, name appkit.string) GraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("depthwiseConvolution3DDataGradientWithIncomingGradientTensor:weightsTensor:outputShape:descriptor:name:"), incomingGradient, weights, outputShape, descriptor, name)
	return rv
}

// Creates a fast Fourier transform operation and returns the result tensor.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/fastFourierTransform(_:axes:descriptor:name:)
func (g_ Graph) FastFourierTransformWithTensorAxesDescriptorName(tensor IMPSGraphTensor, axes []foundation.INumber, descriptor IMPSGraphFFTDescriptor, name appkit.string) GraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("fastFourierTransformWithTensor:axes:descriptor:name:"), tensor, axes, descriptor, name)
	return rv
}

// Copies the input tensor values into the output, behaving as an identity operation.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/identity(with:name:)
func (g_ Graph) IdentityWithTensorName(tensor IMPSGraphTensor, name appkit.string) GraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("identityWithTensor:name:"), tensor, name)
	return rv
}

// Creates an imToCol operation and returns the result tensor.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/imToCol(_:descriptor:name:)
func (g_ Graph) ImToColWithSourceTensorDescriptorName(source IMPSGraphTensor, descriptor IMPSGraphImToColOpDescriptor, name appkit.string) GraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("imToColWithSourceTensor:descriptor:name:"), source, descriptor, name)
	return rv
}

// Creates a 2D max-pooling operation and returns the result tensor.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/maxPooling2D(withSourceTensor:descriptor:name:)
func (g_ Graph) MaxPooling2DWithSourceTensorDescriptorName(source IMPSGraphTensor, descriptor IMPSGraphPooling2DOpDescriptor, name appkit.string) GraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("maxPooling2DWithSourceTensor:descriptor:name:"), source, descriptor, name)
	return rv
}

// Creates a max-pooling gradient operation and returns the result tensor.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/maxPooling2DGradient(withGradientTensor:sourceTensor:descriptor:name:)
func (g_ Graph) MaxPooling2DGradientWithGradientTensorSourceTensorDescriptorName(gradient IMPSGraphTensor, source IMPSGraphTensor, descriptor IMPSGraphPooling2DOpDescriptor, name appkit.string) GraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("maxPooling2DGradientWithGradientTensor:sourceTensor:descriptor:name:"), gradient, source, descriptor, name)
	return rv
}

// Creates a 2D max-pooling operation and returns the result tensor and the corresponding indices tensor.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/maxPooling2DReturnIndices(_:descriptor:name:)
func (g_ Graph) MaxPooling2DReturnIndicesWithSourceTensorDescriptorName(source IMPSGraphTensor, descriptor IMPSGraphPooling2DOpDescriptor, name appkit.string) []GraphTensor {
	rv := objc.Send[[]GraphTensor](g_.ID, objc.Sel("maxPooling2DReturnIndicesWithSourceTensor:descriptor:name:"), source, descriptor, name)
	return rv
}

// Creates a 4D max-pooling operation and returns the result tensor.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/maxPooling4D(_:descriptor:name:)
func (g_ Graph) MaxPooling4DWithSourceTensorDescriptorName(source IMPSGraphTensor, descriptor IMPSGraphPooling4DOpDescriptor, name appkit.string) GraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("maxPooling4DWithSourceTensor:descriptor:name:"), source, descriptor, name)
	return rv
}

// Creates a 4D max-pooling operation and returns the result tensor and the corresponding indices tensor.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/maxPooling4DReturnIndices(_:descriptor:name:)
func (g_ Graph) MaxPooling4DReturnIndicesWithSourceTensorDescriptorName(source IMPSGraphTensor, descriptor IMPSGraphPooling4DOpDescriptor, name appkit.string) []GraphTensor {
	rv := objc.Send[[]GraphTensor](g_.ID, objc.Sel("maxPooling4DReturnIndicesWithSourceTensor:descriptor:name:"), source, descriptor, name)
	return rv
}

// Creates a Real-to-Hermitean fast Fourier transform operation and returns the result tensor.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/realToHermiteanFFT(_:axesTensor:descriptor:name:)
func (g_ Graph) RealToHermiteanFFTWithTensorAxesTensorDescriptorName(tensor IMPSGraphTensor, axesTensor IMPSGraphTensor, descriptor IMPSGraphFFTDescriptor, name appkit.string) GraphTensor {
	rv := objc.Send[GraphTensor](g_.ID, objc.Sel("realToHermiteanFFTWithTensor:axesTensor:descriptor:name:"), tensor, axesTensor, descriptor, name)
	return rv
}

// Creates a single-gate RNN operation and returns the value and optionally the training state tensor.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/singleGateRNN(_:recurrentWeight:initState:descriptor:name:)
func (g_ Graph) SingleGateRNNWithSourceTensorRecurrentWeightInitStateDescriptorName(source IMPSGraphTensor, recurrentWeight IMPSGraphTensor, initState IMPSGraphTensor, descriptor IMPSGraphSingleGateRNNDescriptor, name appkit.string) []GraphTensor {
	rv := objc.Send[[]GraphTensor](g_.ID, objc.Sel("singleGateRNNWithSourceTensor:recurrentWeight:initState:descriptor:name:"), source, recurrentWeight, initState, descriptor, name)
	return rv
}

// Creates a single-gate RNN operation and returns the value and optionally the training state tensor.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraph/singleGateRNN(_:recurrentWeight:inputWeight:bias:initState:mask:descriptor:name:)
func (g_ Graph) SingleGateRNNWithSourceTensorRecurrentWeightInputWeightBiasInitStateMaskDescriptorName(source IMPSGraphTensor, recurrentWeight IMPSGraphTensor, inputWeight IMPSGraphTensor, bias IMPSGraphTensor, initState IMPSGraphTensor, mask IMPSGraphTensor, descriptor IMPSGraphSingleGateRNNDescriptor, name appkit.string) []GraphTensor {
	rv := objc.Send[[]GraphTensor](g_.ID, objc.Sel("singleGateRNNWithSourceTensor:recurrentWeight:inputWeight:bias:initState:mask:descriptor:name:"), source, recurrentWeight, inputWeight, bias, initState, mask, descriptor, name)
	return rv
}

// Options for the graph.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraph/options
func (g_ Graph) Options() GraphOptions {
	rv := objc.Send[GraphOptions](g_.ID, objc.Sel("options"))
	return rv
}


// SetOptions sets the value of the options property.
// Options for the graph.

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraph/options
func (g_ Graph) SetOptions(value GraphOptions) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setOptions:"), value)
}

// Array of all the placeholder tensors.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraph/placeholdertensors
func (g_ Graph) PlaceholderTensors() MPSGraphTensor {
	rv := objc.Send[MPSGraphTensor](g_.ID, objc.Sel("placeholderTensors"))
	return rv
}


// SetPlaceholderTensors sets the value of the placeholderTensors property.
// Array of all the placeholder tensors.

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraph/placeholdertensors
func (g_ Graph) SetPlaceholderTensors(value IMPSGraphTensor) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setPlaceholderTensors:"), value)
}



