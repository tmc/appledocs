//go:build darwin && ios

// Code generated from Apple documentation for AudioToolbox. DO NOT EDIT.

package audiotoolbox

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// iOS-only methods for FixedSpatialAudio


// iOS-only properties

// The experience’s sound stage size.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/CAFixedSpatialAudio/soundStageSize
func (f_ FixedSpatialAudio) SoundStageSize() SoundStageSize {
	rv := objc.Send[SoundStageSize](f_.ID, objc.Sel("soundStageSize"))
	return rv
}




