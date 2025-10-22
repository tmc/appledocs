// Code generated from Apple documentation for AudioToolbox. DO NOT EDIT.

package audiotoolbox

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [FixedSpatialAudio] class.
var (
	FixedSpatialAudioClass     _FixedSpatialAudioClass
	FixedSpatialAudioClassOnce sync.Once
)

func getFixedSpatialAudioClass() _FixedSpatialAudioClass {
	FixedSpatialAudioClassOnce.Do(func() {
		FixedSpatialAudioClass = _FixedSpatialAudioClass{objc.GetClass("CAFixedSpatialAudio")}
	})
	return FixedSpatialAudioClass
}

type _FixedSpatialAudioClass struct {
	class objc.Class
}

// An interface definition for the [FixedSpatialAudio] class.
type IFixedSpatialAudio interface {
	ISpatialAudioExperience
	SoundStageSize() SoundStageSize
}

// A spatial experience that does not take user motion into account.
//
// The Objective-C version of the Swift type.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/CAFixedSpatialAudio
type FixedSpatialAudio struct {
	SpatialAudioExperience
}

// FixedSpatialAudioFrom constructs a [FixedSpatialAudio] from an unsafe.Pointer.
//
// A spatial experience that does not take user motion into account.
func FixedSpatialAudioFrom(ptr unsafe.Pointer) FixedSpatialAudio {
	return FixedSpatialAudio{
		SpatialAudioExperience: SpatialAudioExperienceFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (fc _FixedSpatialAudioClass) Alloc() FixedSpatialAudio {
	rv := objc.Send[FixedSpatialAudio](objc.ID(fc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (fc _FixedSpatialAudioClass) New() FixedSpatialAudio {
	rv := objc.Send[FixedSpatialAudio](objc.ID(fc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (f_ FixedSpatialAudio) Init() FixedSpatialAudio {
	rv := objc.Send[FixedSpatialAudio](f_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f_ FixedSpatialAudio) Autorelease() FixedSpatialAudio {
	rv := objc.Send[FixedSpatialAudio](f_.ID, objc.Sel("autorelease"))
	return rv
}

// NewFixedSpatialAudio creates a new FixedSpatialAudio instance.
func NewFixedSpatialAudio() FixedSpatialAudio {
	return getFixedSpatialAudioClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/CAFixedSpatialAudio/initWithSoundStageSize:
func NewFixedSpatialAudioWithSoundStageSize(soundStageSize ISoundStageSize) FixedSpatialAudio {
	instance := getFixedSpatialAudioClass().Alloc()
	rv := objc.Send[FixedSpatialAudio](instance.ID, objc.Sel("initWithSoundStageSize:"), soundStageSize)
	rv.Autorelease()
	return rv
}


// The experience’s sound stage size.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/CAFixedSpatialAudio/soundStageSize
func (f_ FixedSpatialAudio) SoundStageSize() SoundStageSize {
	rv := objc.Send[SoundStageSize](f_.ID, objc.Sel("soundStageSize"))
	return rv
}


