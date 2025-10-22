// Code generated from Apple documentation for AudioToolbox. DO NOT EDIT.

package audiotoolbox

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [HeadTrackedSpatialAudio] class.
var (
	HeadTrackedSpatialAudioClass     _HeadTrackedSpatialAudioClass
	HeadTrackedSpatialAudioClassOnce sync.Once
)

func getHeadTrackedSpatialAudioClass() _HeadTrackedSpatialAudioClass {
	HeadTrackedSpatialAudioClassOnce.Do(func() {
		HeadTrackedSpatialAudioClass = _HeadTrackedSpatialAudioClass{objc.GetClass("CAHeadTrackedSpatialAudio")}
	})
	return HeadTrackedSpatialAudioClass
}

type _HeadTrackedSpatialAudioClass struct {
	class objc.Class
}

// An interface definition for the [HeadTrackedSpatialAudio] class.
type IHeadTrackedSpatialAudio interface {
	ISpatialAudioExperience
	AnchoringStrategy() CAAnchoringStrategy
	SoundStageSize() SoundStageSize
}

// A spatial experience that takes user motion into account.
//
// The Objective-C version of the Swift type.


// A spatial experience that takes user motion into account.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/CAHeadTrackedSpatialAudio

type HeadTrackedSpatialAudio struct {
	SpatialAudioExperience
}

// HeadTrackedSpatialAudioFrom constructs a [HeadTrackedSpatialAudio] from an unsafe.Pointer.
//
// A spatial experience that takes user motion into account.
func HeadTrackedSpatialAudioFrom(ptr unsafe.Pointer) HeadTrackedSpatialAudio {
	return HeadTrackedSpatialAudio{
		SpatialAudioExperience: SpatialAudioExperienceFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (hc _HeadTrackedSpatialAudioClass) Alloc() HeadTrackedSpatialAudio {
	rv := objc.Send[HeadTrackedSpatialAudio](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (hc _HeadTrackedSpatialAudioClass) New() HeadTrackedSpatialAudio {
	rv := objc.Send[HeadTrackedSpatialAudio](objc.ID(hc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (h_ HeadTrackedSpatialAudio) Init() HeadTrackedSpatialAudio {
	rv := objc.Send[HeadTrackedSpatialAudio](h_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (h_ HeadTrackedSpatialAudio) Autorelease() HeadTrackedSpatialAudio {
	rv := objc.Send[HeadTrackedSpatialAudio](h_.ID, objc.Sel("autorelease"))
	return rv
}

// NewHeadTrackedSpatialAudio creates a new HeadTrackedSpatialAudio instance.
func NewHeadTrackedSpatialAudio() HeadTrackedSpatialAudio {
	return getHeadTrackedSpatialAudioClass().New()
}




// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/CAHeadTrackedSpatialAudio/initWithSoundStageSize:anchoringStrategy:

func NewHeadTrackedSpatialAudioWithSoundStageSizeAnchoringStrategy(soundStageSize ISoundStageSize, anchoringStrategy CAAnchoringStrategy) HeadTrackedSpatialAudio {
	instance := getHeadTrackedSpatialAudioClass().Alloc()
	rv := objc.Send[HeadTrackedSpatialAudio](instance.ID, objc.Sel("initWithSoundStageSize:anchoringStrategy:"), soundStageSize, anchoringStrategy)
	rv.Autorelease()
	return rv
}



// The experience’s anchoring strategy.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/CAHeadTrackedSpatialAudio/anchoringStrategy

func (h_ HeadTrackedSpatialAudio) AnchoringStrategy() CAAnchoringStrategy {
	rv := objc.Send[CAAnchoringStrategy](h_.ID, objc.Sel("anchoringStrategy"))
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


