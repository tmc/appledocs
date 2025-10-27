// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objectivec"
)

// PCNNConvolutionDataSource is the MPSCNNConvolutionDataSource protocol interface.
//
// The protocol that provides convolution filter weights and bias terms.
//
// Availability:
//   - Mac Catalyst +
//   - iOS +
//   - iPadOS +
//   - macOS +
//   - tvOS +
//   - visionOS +
//
// See: doc://com.apple.metalperformanceshaders/documentation/MetalPerformanceShaders/MPSCNNConvolutionDataSource
type PCNNConvolutionDataSource interface {
	// Required methods
	BiasTerms()
	Load()
	Descriptor()
	Purge()
	DataType()
	Weights()
	Label()
	UpdateWithCommandBufferGradientStateSourceState(commandBuffer unsafe.Pointer, gradientState ICNNConvolutionGradientState, sourceState ICNNConvolutionWeightsAndBiasesState) ICNNConvolutionWeightsAndBiasesState
	UpdateWithGradientStateSourceState(gradientState ICNNConvolutionGradientState, sourceState ICNNConvolutionWeightsAndBiasesState) bool
	CopyWithZoneDevice(zone Zone /* not a class type */, device unsafe.Pointer) objectivec.IObject
	// Optional methods
	RangesForUInt8Kernel()
	HasRangesForUInt8Kernel() bool
	LookupTableForUInt8Kernel()
	HasLookupTableForUInt8Kernel() bool
	Update()
	HasUpdate() bool
	WeightsQuantizationType()
	HasWeightsQuantizationType() bool
	Copy()
	HasCopy() bool
	WeightsLayout()
	HasWeightsLayout() bool
	KernelWeightsDataType()
	HasKernelWeightsDataType() bool
}
