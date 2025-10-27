// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/foundation"

	"github.com/tmc/appledocs/generated/objectivec"
)

// PCNNGroupNormalizationDataSource is the MPSCNNGroupNormalizationDataSource protocol interface.
//
// Availability:
//   - Mac Catalyst 13.0+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - tvOS 13.0+
//   - visionOS 1.0+
//
// See: doc://com.apple.metalperformanceshaders/documentation/MetalPerformanceShaders/MPSCNNGroupNormalizationDataSource
type PCNNGroupNormalizationDataSource interface {
	// Required methods
	Beta()
	CopyWithZoneDevice(zone Zone /* not a class type */, device unsafe.Pointer) objectivec.IObject
	EncodeWithCoder(aCoder foundation.Coder)
	Gamma()
	InitWithCoder(aDecoder foundation.Coder) objectivec.IObject
	Label()
	UpdateGammaAndBetaWithCommandBufferGroupNormalizationStateBatch(commandBuffer unsafe.Pointer, groupNormalizationStateBatch CNNGroupNormalizationGradientStateBatch /* not a class type */) ICNNNormalizationGammaAndBetaState
	UpdateGammaAndBetaWithGroupNormalizationStateBatch(groupNormalizationStateBatch CNNGroupNormalizationGradientStateBatch /* not a class type */) bool
	// Optional methods
	Copy()
	HasCopy() bool
	Encode()
	HasEncode() bool
	Epsilon()
	HasEpsilon() bool
	UpdateGammaAndBeta()
	HasUpdateGammaAndBeta() bool
}
