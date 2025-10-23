// Code generated from Apple documentation for AudioToolbox. DO NOT EDIT.

package audiotoolbox

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [AutomaticSpatialAudio] class.
var (
	AutomaticSpatialAudioClass     _AutomaticSpatialAudioClass
	AutomaticSpatialAudioClassOnce sync.Once
)

func getAutomaticSpatialAudioClass() _AutomaticSpatialAudioClass {
	AutomaticSpatialAudioClassOnce.Do(func() {
		AutomaticSpatialAudioClass = _AutomaticSpatialAudioClass{objc.GetClass("CAAutomaticSpatialAudio")}
	})
	return AutomaticSpatialAudioClass
}

type _AutomaticSpatialAudioClass struct {
	class objc.Class
}

// An interface definition for the [AutomaticSpatialAudio] class.
type IAutomaticSpatialAudio interface {
	ISpatialAudioExperience
}

// A spatial audio experience determined by the system.
//
// The Objective-C version of the Swift type.


// A spatial audio experience determined by the system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/CAAutomaticSpatialAudio
type AutomaticSpatialAudio struct {
	SpatialAudioExperience
}

// AutomaticSpatialAudioFrom constructs a [AutomaticSpatialAudio] from an unsafe.Pointer.
//
// A spatial audio experience determined by the system.
func AutomaticSpatialAudioFrom(ptr unsafe.Pointer) AutomaticSpatialAudio {
	return AutomaticSpatialAudio{
		SpatialAudioExperience: SpatialAudioExperienceFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (ac _AutomaticSpatialAudioClass) Alloc() AutomaticSpatialAudio {
	rv := objc.Send[AutomaticSpatialAudio](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ac _AutomaticSpatialAudioClass) New() AutomaticSpatialAudio {
	rv := objc.Send[AutomaticSpatialAudio](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AutomaticSpatialAudio) Init() AutomaticSpatialAudio {
	rv := objc.Send[AutomaticSpatialAudio](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AutomaticSpatialAudio) Autorelease() AutomaticSpatialAudio {
	rv := objc.Send[AutomaticSpatialAudio](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAutomaticSpatialAudio creates a new AutomaticSpatialAudio instance.
func NewAutomaticSpatialAudio() AutomaticSpatialAudio {
	return getAutomaticSpatialAudioClass().New()
}




