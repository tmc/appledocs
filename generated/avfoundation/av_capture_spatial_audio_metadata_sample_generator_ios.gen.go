//go:build darwin && ios

// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// iOS-only methods for CaptureSpatialAudioMetadataSampleGenerator


// Analyzes the provided audio sample buffer for its contribution to the spatial audio timed metadata value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureSpatialAudioMetadataSampleGenerator/analyzeAudioSample(_:)
func (c_ CaptureSpatialAudioMetadataSampleGenerator) AnalyzeAudioSample(sbuf SampleBufferRef /* not a class type */) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("analyzeAudioSample:"), sbuf)
	return rv
}

// Creates a sample buffer containing a spatial audio timed metadata sample computed from all analyzed audio buffers, and resets the analyzer to its initial state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureSpatialAudioMetadataSampleGenerator/newTimedMetadataSampleBufferAndResetAnalyzer()
func (c_ CaptureSpatialAudioMetadataSampleGenerator) NewTimedMetadataSampleBufferAndResetAnalyzer() SampleBufferRef /* not a class type */ {
	rv := objc.Send[SampleBufferRef](c_.ID, objc.Sel("newTimedMetadataSampleBufferAndResetAnalyzer"))
	return rv
}

// Calling this method resets the analyzer to its initial state so that a new run of audio sample buffers can be analyzed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureSpatialAudioMetadataSampleGenerator/resetAnalyzer()
func (c_ CaptureSpatialAudioMetadataSampleGenerator) ResetAnalyzer() {
	objc.Send[objc.ID](c_.ID, objc.Sel("resetAnalyzer"))
}

// iOS-only properties

// Returns the format description of the sample buffer returned from the method.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureSpatialAudioMetadataSampleGenerator/timedMetadataSampleBufferFormatDescription
func (c_ CaptureSpatialAudioMetadataSampleGenerator) TimedMetadataSampleBufferFormatDescription() FormatDescriptionRef /* not a class type */ {
	rv := objc.Send[FormatDescriptionRef](c_.ID, objc.Sel("timedMetadataSampleBufferFormatDescription"))
	return rv
}





