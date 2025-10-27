//go:build darwin && ios

// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// iOS-only methods for SampleBufferRenderSynchronizer


// iOS-only properties

// The intended spatial audio experience applied to all AVSampleBufferAudioRenderers within this synchronizer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVSampleBufferRenderSynchronizer/intendedSpatialAudioExperience-2wthu
func (s_ SampleBufferRenderSynchronizer) IntendedSpatialAudioExperience() SpatialAudioExperience /* not a class type */ {
	rv := objc.Send[SpatialAudioExperience](s_.ID, objc.Sel("intendedSpatialAudioExperience"))
	return rv
}
func (s_ SampleBufferRenderSynchronizer) SetIntendedSpatialAudioExperience(value SpatialAudioExperience /* not a class type */) {
	s_.ID.Send(objc.RegisterName("setIntendedSpatialAudioExperience:"), value)
}





