// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/foundation"

	"github.com/tmc/appledocs/generated/objectivec"
)

// PVideoCompositionValidationHandling is the AVVideoCompositionValidationHandling protocol interface.
//
// Methods you can implement to indicate whether validation of a video composition should continue after specific errors are found.
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 5.0+
//   - iPadOS 5.0+
//   - macOS 10.8+
//   - tvOS 9.0+
//   - visionOS 1.0+
//
// See: doc://com.apple.avfoundation/documentation/AVFoundation/AVVideoCompositionValidationHandling
type PVideoCompositionValidationHandling interface {
	// Optional methods
	VideoCompositionShouldContinueValidatingAfterFindingEmptyTimeRange(videoComposition IAVVideoComposition, timeRange objectivec.IObject) bool
	HasVideoCompositionShouldContinueValidatingAfterFindingEmptyTimeRange() bool
	VideoCompositionShouldContinueValidatingAfterFindingInvalidTimeRangeInInstruction(videoComposition IAVVideoComposition, videoCompositionInstruction unsafe.Pointer) bool
	HasVideoCompositionShouldContinueValidatingAfterFindingInvalidTimeRangeInInstruction() bool
	VideoCompositionShouldContinueValidatingAfterFindingInvalidTrackIDInInstructionLayerInstructionAsset(videoComposition IAVVideoComposition, videoCompositionInstruction unsafe.Pointer, layerInstruction IAVVideoCompositionLayerInstruction, asset IAVAsset) bool
	HasVideoCompositionShouldContinueValidatingAfterFindingInvalidTrackIDInInstructionLayerInstructionAsset() bool
	VideoCompositionShouldContinueValidatingAfterFindingInvalidValueForKey(videoComposition IAVVideoComposition, key foundation.foundation.INSString) bool
	HasVideoCompositionShouldContinueValidatingAfterFindingInvalidValueForKey() bool
}
