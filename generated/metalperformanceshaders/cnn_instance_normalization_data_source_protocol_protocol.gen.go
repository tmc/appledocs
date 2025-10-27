// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/foundation"

	"github.com/tmc/appledocs/generated/objectivec"
)

// PCNNInstanceNormalizationDataSource is the MPSCNNInstanceNormalizationDataSource protocol interface.
//
// A protocol that defines methods that an instance normalization uses to initialize scale factors and bias terms.
//
// Availability:
//   - Mac Catalyst 13.0+
//   - iOS 11.3+
//   - iPadOS 11.3+
//   - macOS 10.13.4+
//   - tvOS 11.3+
//   - visionOS 1.0+
//
// See: doc://com.apple.metalperformanceshaders/documentation/MetalPerformanceShaders/MPSCNNInstanceNormalizationDataSource
type PCNNInstanceNormalizationDataSource interface {
	// Required methods
	EncodeWithCoder(aCoder foundation.Coder)
	InitWithCoder(aDecoder foundation.Coder) objectivec.IObject
	Label()
	Beta()
	Gamma()
	UpdateGammaAndBetaWithCommandBufferInstanceNormalizationStateBatch(commandBuffer unsafe.Pointer, instanceNormalizationStateBatch CNNInstanceNormalizationGradientStateBatch /* not a class type */) ICNNNormalizationGammaAndBetaState
	UpdateGammaAndBetaWithInstanceNormalizationStateBatch(instanceNormalizationStateBatch CNNInstanceNormalizationGradientStateBatch /* not a class type */) bool
	CopyWithZoneDevice(zone Zone /* not a class type */, device unsafe.Pointer) objectivec.IObject
	// Optional methods
	Encode()
	HasEncode() bool
	Epsilon()
	HasEpsilon() bool
	UpdateGammaAndBeta()
	HasUpdateGammaAndBeta() bool
	Copy()
	HasCopy() bool
	Load()
	HasLoad() bool
	Purge()
	HasPurge() bool
}
