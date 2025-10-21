// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [AudioUnitTimeEffect] class.
var (
	AudioUnitTimeEffectClass     _AudioUnitTimeEffectClass
	AudioUnitTimeEffectClassOnce sync.Once
)

func getAudioUnitTimeEffectClass() _AudioUnitTimeEffectClass {
	AudioUnitTimeEffectClassOnce.Do(func() {
		AudioUnitTimeEffectClass = _AudioUnitTimeEffectClass{objc.GetClass("AVAudioUnitTimeEffect")}
	})
	return AudioUnitTimeEffectClass
}

type _AudioUnitTimeEffectClass struct {
	class objc.Class
}

// An interface definition for the [AudioUnitTimeEffect] class.
type IAudioUnitTimeEffect interface {
	IAudioUnit
}

// An object that processes audio in nonreal time.
//
// A time effect audio unit represents an with a type ( . These effects don’t process audio in real time. The class is an example of a time effect unit.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitTimeEffect
type AudioUnitTimeEffect struct {
	AudioUnit
}

// AudioUnitTimeEffectFrom constructs a [AudioUnitTimeEffect] from an unsafe.Pointer.
//
// An object that processes audio in nonreal time.
func AudioUnitTimeEffectFrom(ptr unsafe.Pointer) AudioUnitTimeEffect {
	return AudioUnitTimeEffect{
		AudioUnit: AudioUnitFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (ac _AudioUnitTimeEffectClass) Alloc() AudioUnitTimeEffect {
	rv := objc.Send[AudioUnitTimeEffect](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ac _AudioUnitTimeEffectClass) New() AudioUnitTimeEffect {
	rv := objc.Send[AudioUnitTimeEffect](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AudioUnitTimeEffect) Init() AudioUnitTimeEffect {
	rv := objc.Send[AudioUnitTimeEffect](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AudioUnitTimeEffect) Autorelease() AudioUnitTimeEffect {
	rv := objc.Send[AudioUnitTimeEffect](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAudioUnitTimeEffect creates a new AudioUnitTimeEffect instance.
func NewAudioUnitTimeEffect() AudioUnitTimeEffect {
	return getAudioUnitTimeEffectClass().New()
}


// The bypass state of the audio unit.
//
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudiounittimeeffect/bypass
func (a_ AudioUnitTimeEffect) Bypass() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("bypass"))
	return rv
}


// SetBypass sets the value of the bypass property.
// The bypass state of the audio unit.

//
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudiounittimeeffect/bypass
func (a_ AudioUnitTimeEffect) SetBypass(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setBypass:"), value)
}



