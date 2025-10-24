// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)





// The class instance for the [AudioUnitVarispeed] class.
var (
	AudioUnitVarispeedClass     _AudioUnitVarispeedClass
	AudioUnitVarispeedClassOnce sync.Once
)

func getAudioUnitVarispeedClass() _AudioUnitVarispeedClass {
	AudioUnitVarispeedClassOnce.Do(func() {
		AudioUnitVarispeedClass = _AudioUnitVarispeedClass{objc.GetClass("AVAudioUnitVarispeed")}
	})
	return AudioUnitVarispeedClass
}

type _AudioUnitVarispeedClass struct {
	class objc.Class
}





// An interface definition for the [AudioUnitVarispeed] class.
type IAudioUnitVarispeed interface {
	IAudioUnitTimeEffect
	

	// properties:
	Rate() float32
	SetRate(value float32)


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (ac _AudioUnitVarispeedClass) Alloc() AudioUnitVarispeed {
	rv := objc.Send[AudioUnitVarispeed](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ac _AudioUnitVarispeedClass) New() AudioUnitVarispeed {
	rv := objc.Send[AudioUnitVarispeed](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AudioUnitVarispeed) Init() AudioUnitVarispeed {
	rv := objc.Send[AudioUnitVarispeed](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AudioUnitVarispeed) Autorelease() AudioUnitVarispeed {
	rv := objc.Send[AudioUnitVarispeed](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAudioUnitVarispeed creates a new AudioUnitVarispeed instance.
func NewAudioUnitVarispeed() AudioUnitVarispeed {
	return getAudioUnitVarispeedClass().New()
}





// An object that allows control of the playback rate.


// An object that allows control of the playback rate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitVarispeed
type AudioUnitVarispeed struct {
	AudioUnitTimeEffect
}

// AudioUnitVarispeedFrom constructs a [AudioUnitVarispeed] from an unsafe.Pointer.
//
// An object that allows control of the playback rate.
func AudioUnitVarispeedFrom(ptr unsafe.Pointer) AudioUnitVarispeed {
	return AudioUnitVarispeed{
		AudioUnitTimeEffect: AudioUnitTimeEffectFrom(ptr),
	}
}

























// The audio playback rate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitVarispeed/rate
func (a_ AudioUnitVarispeed) Rate() float32 {
	rv := objc.Send[float32](a_.ID, objc.Sel("rate"))
	return rv
}


// The audio playback rate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitVarispeed/rate
func (a_ AudioUnitVarispeed) SetRate(value float32) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setRate:"), value)
}








