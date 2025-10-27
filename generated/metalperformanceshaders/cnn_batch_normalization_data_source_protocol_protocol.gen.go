// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/foundation"

	"github.com/tmc/appledocs/generated/objectivec"
)

// PCNNBatchNormalizationDataSource is the MPSCNNBatchNormalizationDataSource protocol interface.
//
// A protocol that defines methods that a batch normalization state uses to initialize scale factors, bias terms, and batch statistics.
//
// Availability:
//   - Mac Catalyst 13.0+
//   - iOS 11.3+
//   - iPadOS 11.3+
//   - macOS 10.13.4+
//   - tvOS 11.3+
//   - visionOS 1.0+
//
// See: doc://com.apple.metalperformanceshaders/documentation/MetalPerformanceShaders/MPSCNNBatchNormalizationDataSource
type PCNNBatchNormalizationDataSource interface {
	// Required methods
	Load()
	Beta()
	Mean()
	Variance()
	NumberOfFeatureChannels()
	Gamma()
	Purge()
	EncodeWithCoder(aCoder foundation.Coder)
	InitWithCoder(aDecoder foundation.Coder) objectivec.IObject
	UpdateGammaAndBetaWithCommandBufferBatchNormalizationState(commandBuffer unsafe.Pointer, batchNormalizationState ICNNBatchNormalizationState) ICNNNormalizationGammaAndBetaState
	Label()
	UpdateGammaAndBetaWithBatchNormalizationState(batchNormalizationState ICNNBatchNormalizationState) bool
	UpdateMeanAndVarianceWithBatchNormalizationState(batchNormalizationState ICNNBatchNormalizationState) bool
	UpdateMeanAndVarianceWithCommandBufferBatchNormalizationState(commandBuffer unsafe.Pointer, batchNormalizationState ICNNBatchNormalizationState) ICNNNormalizationMeanAndVarianceState
	CopyWithZoneDevice(zone Zone /* not a class type */, device unsafe.Pointer) objectivec.IObject
	// Optional methods
	Epsilon()
	HasEpsilon() bool
	Encode()
	HasEncode() bool
	UpdateGammaAndBeta()
	HasUpdateGammaAndBeta() bool
	UpdateMeanAndVariance()
	HasUpdateMeanAndVariance() bool
	Copy()
	HasCopy() bool
}
