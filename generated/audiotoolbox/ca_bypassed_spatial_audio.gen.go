// Code generated from Apple documentation for AudioToolbox. DO NOT EDIT.

package audiotoolbox

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [BypassedSpatialAudio] class.
var (
	BypassedSpatialAudioClass     _BypassedSpatialAudioClass
	BypassedSpatialAudioClassOnce sync.Once
)

func getBypassedSpatialAudioClass() _BypassedSpatialAudioClass {
	BypassedSpatialAudioClassOnce.Do(func() {
		BypassedSpatialAudioClass = _BypassedSpatialAudioClass{objc.GetClass("CABypassedSpatialAudio")}
	})
	return BypassedSpatialAudioClass
}

type _BypassedSpatialAudioClass struct {
	class objc.Class
}

// An interface definition for the [BypassedSpatialAudio] class.
type IBypassedSpatialAudio interface {
	ISpatialAudioExperience
}

// An experience in which the system does not apply spatial processing to the audio stream.
//
// The Objective-C version of the Swift type.


// An experience in which the system does not apply spatial processing to the audio stream.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/CABypassedSpatialAudio

type BypassedSpatialAudio struct {
	SpatialAudioExperience
}

// BypassedSpatialAudioFrom constructs a [BypassedSpatialAudio] from an unsafe.Pointer.
//
// An experience in which the system does not apply spatial processing to the audio stream.
func BypassedSpatialAudioFrom(ptr unsafe.Pointer) BypassedSpatialAudio {
	return BypassedSpatialAudio{
		SpatialAudioExperience: SpatialAudioExperienceFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (bc _BypassedSpatialAudioClass) Alloc() BypassedSpatialAudio {
	rv := objc.Send[BypassedSpatialAudio](objc.ID(bc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (bc _BypassedSpatialAudioClass) New() BypassedSpatialAudio {
	rv := objc.Send[BypassedSpatialAudio](objc.ID(bc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (b_ BypassedSpatialAudio) Init() BypassedSpatialAudio {
	rv := objc.Send[BypassedSpatialAudio](b_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (b_ BypassedSpatialAudio) Autorelease() BypassedSpatialAudio {
	rv := objc.Send[BypassedSpatialAudio](b_.ID, objc.Sel("autorelease"))
	return rv
}

// NewBypassedSpatialAudio creates a new BypassedSpatialAudio instance.
func NewBypassedSpatialAudio() BypassedSpatialAudio {
	return getBypassedSpatialAudioClass().New()
}




