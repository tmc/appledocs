// Code generated from Apple documentation for AudioToolbox. DO NOT EDIT.

package audiotoolbox

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [SpatialAudioExperience] class.
var (
	SpatialAudioExperienceClass     _SpatialAudioExperienceClass
	SpatialAudioExperienceClassOnce sync.Once
)

func getSpatialAudioExperienceClass() _SpatialAudioExperienceClass {
	SpatialAudioExperienceClassOnce.Do(func() {
		SpatialAudioExperienceClass = _SpatialAudioExperienceClass{objc.GetClass("CASpatialAudioExperience")}
	})
	return SpatialAudioExperienceClass
}

type _SpatialAudioExperienceClass struct {
	class objc.Class
}

// An interface definition for the [SpatialAudioExperience] class.
type ISpatialAudioExperience interface {
	objectivec.IObject
}

// Configure an audio stream for spatial computing.
//
// The Objective-C version of the Swift type.


// Configure an audio stream for spatial computing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/CASpatialAudioExperience

type SpatialAudioExperience struct {
	objectivec.Object
}

// SpatialAudioExperienceFrom constructs a [SpatialAudioExperience] from an unsafe.Pointer.
//
// Configure an audio stream for spatial computing.
func SpatialAudioExperienceFrom(ptr unsafe.Pointer) SpatialAudioExperience {
	return SpatialAudioExperience{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (sc _SpatialAudioExperienceClass) Alloc() SpatialAudioExperience {
	rv := objc.Send[SpatialAudioExperience](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (sc _SpatialAudioExperienceClass) New() SpatialAudioExperience {
	rv := objc.Send[SpatialAudioExperience](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ SpatialAudioExperience) Init() SpatialAudioExperience {
	rv := objc.Send[SpatialAudioExperience](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ SpatialAudioExperience) Autorelease() SpatialAudioExperience {
	rv := objc.Send[SpatialAudioExperience](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSpatialAudioExperience creates a new SpatialAudioExperience instance.
func NewSpatialAudioExperience() SpatialAudioExperience {
	return getSpatialAudioExperienceClass().New()
}





