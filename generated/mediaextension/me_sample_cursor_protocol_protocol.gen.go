// Code generated from Apple documentation for MediaExtension. DO NOT EDIT.

package mediaextension

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/avfoundation"

	"github.com/tmc/appledocs/generated/corevideo"
)

// PMESampleCursor is the MESampleCursor protocol interface.
//
// A protocol that defines the information to provide about samples within a track of a media asset, and enables stepping through samples in the track in decode or presentation order.
//
// Availability:
//   - macOS 14.0+
//
// See: doc://com.apple.mediaextension/documentation/MediaExtension/MESampleCursor
type PMESampleCursor interface {
	// Required methods
	StepByDecodeTimeCompletionHandler(deltaDecodeTime objc.IObject /* cross-framework: Time */, completionHandler unsafe.Pointer)/* debug [protocol_interface/required_method]: StepByDecodeTimeCompletionHandler */
	StepByPresentationTimeCompletionHandler(deltaPresentationTime objc.IObject /* cross-framework: Time */, completionHandler unsafe.Pointer)/* debug [protocol_interface/required_method]: StepByPresentationTimeCompletionHandler */
	StepInDecodeOrderByCountCompletionHandler(stepCount int64, completionHandler unsafe.Pointer)/* debug [protocol_interface/required_method]: StepInDecodeOrderByCountCompletionHandler */
	StepInPresentationOrderByCountCompletionHandler(stepCount int64, completionHandler unsafe.Pointer)/* debug [protocol_interface/required_method]: StepInPresentationOrderByCountCompletionHandler */
	// Optional methods
	ChunkDetailsReturningError(error_ unsafe.Pointer) MESampleCursorChunk
	HasChunkDetailsReturningError() bool
	EstimatedSampleLocationReturningError(error_ unsafe.Pointer) MEEstimatedSampleLocation
	HasEstimatedSampleLocationReturningError() bool
	LoadPostDecodeProcessingMetadataWithCompletionHandler(completionHandler unsafe.Pointer)
	HasLoadPostDecodeProcessingMetadataWithCompletionHandler() bool
	LoadSampleBufferContainingSamplesToEndCursorCompletionHandler(endSampleCursor unsafe.Pointer, completionHandler unsafe.Pointer)
	HasLoadSampleBufferContainingSamplesToEndCursorCompletionHandler() bool
	RefineSampleLocationRefinementDataRefinementDataLengthRefinedLocationError(estimatedSampleLocation objc.IObject /* cross-framework: SampleCursorStorageRange */, refinementData unsafe.Pointer, refinementDataLength uintptr /* not a class type */, refinedLocationOut avfoundation.AVSampleCursorStorageRange, error_ unsafe.Pointer) bool
	HasRefineSampleLocationRefinementDataRefinementDataLengthRefinedLocationError() bool
	SampleLocationReturningError(error_ unsafe.Pointer) MESampleLocation
	HasSampleLocationReturningError() bool
	SamplesWithEarlierDTSsMayHaveLaterPTSsThanCursor(cursor unsafe.Pointer) bool
	HasSamplesWithEarlierDTSsMayHaveLaterPTSsThanCursor() bool
	SamplesWithLaterDTSsMayHaveEarlierPTSsThanCursor(cursor unsafe.Pointer) bool
	HasSamplesWithLaterDTSsMayHaveEarlierPTSsThanCursor() bool
}
