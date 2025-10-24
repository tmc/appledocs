//go:build darwin && ios

// Code generated from Apple documentation for AudioToolbox. DO NOT EDIT.

package audiotoolbox

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// iOS-only methods for HeadTrackedSpatialAudio


// iOS-only properties

// The experience’s anchoring strategy.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/CAHeadTrackedSpatialAudio/anchoringStrategy
func (h_ HeadTrackedSpatialAudio) AnchoringStrategy() IAnchoringStrategy {
	rv := objc.Send[AnchoringStrategy](h_.ID, objc.Sel("anchoringStrategy"))
	return rv
}

// The experience’s sound stage size.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/CAHeadTrackedSpatialAudio/soundStageSize
func (h_ HeadTrackedSpatialAudio) SoundStageSize() SoundStageSize {
	rv := objc.Send[SoundStageSize](h_.ID, objc.Sel("soundStageSize"))
	return rv
}




