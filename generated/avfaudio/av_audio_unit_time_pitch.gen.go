// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)





// The class instance for the [AudioUnitTimePitch] class.
var (
	AudioUnitTimePitchClass     _AudioUnitTimePitchClass
	AudioUnitTimePitchClassOnce sync.Once
)

func getAudioUnitTimePitchClass() _AudioUnitTimePitchClass {
	AudioUnitTimePitchClassOnce.Do(func() {
		AudioUnitTimePitchClass = _AudioUnitTimePitchClass{objc.GetClass("AVAudioUnitTimePitch")}
	})
	return AudioUnitTimePitchClass
}

type _AudioUnitTimePitchClass struct {
	class objc.Class
}





// An interface definition for the [AudioUnitTimePitch] class.
type IAudioUnitTimePitch interface {
	IAudioUnitTimeEffect
	

	// properties:
	Overlap() float32
	SetOverlap(value float32)
	Pitch() float32
	SetPitch(value float32)
	Rate() float32
	SetRate(value float32)


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (ac _AudioUnitTimePitchClass) Alloc() AudioUnitTimePitch {
	rv := objc.Send[AudioUnitTimePitch](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ac _AudioUnitTimePitchClass) New() AudioUnitTimePitch {
	rv := objc.Send[AudioUnitTimePitch](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AudioUnitTimePitch) Init() AudioUnitTimePitch {
	rv := objc.Send[AudioUnitTimePitch](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AudioUnitTimePitch) Autorelease() AudioUnitTimePitch {
	rv := objc.Send[AudioUnitTimePitch](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAudioUnitTimePitch creates a new AudioUnitTimePitch instance.
func NewAudioUnitTimePitch() AudioUnitTimePitch {
	return getAudioUnitTimePitchClass().New()
}





// An object that provides a good-quality playback rate and pitch shifting independently of each other.


// An object that provides a good-quality playback rate and pitch shifting independently of each other.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitTimePitch
type AudioUnitTimePitch struct {
	AudioUnitTimeEffect
}

// AudioUnitTimePitchFrom constructs a [AudioUnitTimePitch] from an unsafe.Pointer.
//
// An object that provides a good-quality playback rate and pitch shifting independently of each other.
func AudioUnitTimePitchFrom(ptr unsafe.Pointer) AudioUnitTimePitch {
	return AudioUnitTimePitch{
		AudioUnitTimeEffect: AudioUnitTimeEffectFrom(ptr),
	}
}

























// The amount of overlap between segments of the input audio signal.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitTimePitch/overlap
func (a_ AudioUnitTimePitch) Overlap() float32 {
	rv := objc.Send[float32](a_.ID, objc.Sel("overlap"))
	return rv
}


// The amount of overlap between segments of the input audio signal.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitTimePitch/overlap
func (a_ AudioUnitTimePitch) SetOverlap(value float32) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setOverlap:"), value)
}


// The amount to use to pitch shift the input signal.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitTimePitch/pitch
func (a_ AudioUnitTimePitch) Pitch() float32 {
	rv := objc.Send[float32](a_.ID, objc.Sel("pitch"))
	return rv
}


// The amount to use to pitch shift the input signal.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitTimePitch/pitch
func (a_ AudioUnitTimePitch) SetPitch(value float32) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setPitch:"), value)
}


// The playback rate of the input signal.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitTimePitch/rate
func (a_ AudioUnitTimePitch) Rate() float32 {
	rv := objc.Send[float32](a_.ID, objc.Sel("rate"))
	return rv
}


// The playback rate of the input signal.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitTimePitch/rate
func (a_ AudioUnitTimePitch) SetRate(value float32) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setRate:"), value)
}








