// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)





// The class instance for the [AudioUnitDistortion] class.
var (
	AudioUnitDistortionClass     _AudioUnitDistortionClass
	AudioUnitDistortionClassOnce sync.Once
)

func getAudioUnitDistortionClass() _AudioUnitDistortionClass {
	AudioUnitDistortionClassOnce.Do(func() {
		AudioUnitDistortionClass = _AudioUnitDistortionClass{objc.GetClass("AVAudioUnitDistortion")}
	})
	return AudioUnitDistortionClass
}

type _AudioUnitDistortionClass struct {
	class objc.Class
}





// An interface definition for the [AudioUnitDistortion] class.
type IAudioUnitDistortion interface {
	IAudioUnitEffect
	

	// properties:
	PreGain() float32
	SetPreGain(value float32)
	WetDryMix() float32
	SetWetDryMix(value float32)


	

	// methods:
	LoadFactoryPreset(preset AudioUnitDistortionPreset)


}





// Alloc allocates a new instance without initialization.
func (ac _AudioUnitDistortionClass) Alloc() AudioUnitDistortion {
	rv := objc.Send[AudioUnitDistortion](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ac _AudioUnitDistortionClass) New() AudioUnitDistortion {
	rv := objc.Send[AudioUnitDistortion](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AudioUnitDistortion) Init() AudioUnitDistortion {
	rv := objc.Send[AudioUnitDistortion](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AudioUnitDistortion) Autorelease() AudioUnitDistortion {
	rv := objc.Send[AudioUnitDistortion](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAudioUnitDistortion creates a new AudioUnitDistortion instance.
func NewAudioUnitDistortion() AudioUnitDistortion {
	return getAudioUnitDistortionClass().New()
}





// An object that implements a multistage distortion effect.


// An object that implements a multistage distortion effect.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitDistortion
type AudioUnitDistortion struct {
	AudioUnitEffect
}

// AudioUnitDistortionFrom constructs a [AudioUnitDistortion] from an unsafe.Pointer.
//
// An object that implements a multistage distortion effect.
func AudioUnitDistortionFrom(ptr unsafe.Pointer) AudioUnitDistortion {
	return AudioUnitDistortion{
		AudioUnitEffect: AudioUnitEffectFrom(ptr),
	}
}




















// Configures the audio distortion unit by loading a distortion preset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitDistortion/loadFactoryPreset(_:)
func (a_ AudioUnitDistortion) LoadFactoryPreset(preset AudioUnitDistortionPreset) {
	objc.Send[objc.ID](a_.ID, objc.Sel("loadFactoryPreset:"), preset)
}







// The gain that the audio unit applies to the signal before distortion, in decibels.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitDistortion/preGain
func (a_ AudioUnitDistortion) PreGain() float32 {
	rv := objc.Send[float32](a_.ID, objc.Sel("preGain"))
	return rv
}


// The gain that the audio unit applies to the signal before distortion, in decibels.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitDistortion/preGain
func (a_ AudioUnitDistortion) SetPreGain(value float32) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setPreGain:"), value)
}


// The blend of the distorted and dry signals.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitDistortion/wetDryMix
func (a_ AudioUnitDistortion) WetDryMix() float32 {
	rv := objc.Send[float32](a_.ID, objc.Sel("wetDryMix"))
	return rv
}


// The blend of the distorted and dry signals.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitDistortion/wetDryMix
func (a_ AudioUnitDistortion) SetWetDryMix(value float32) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setWetDryMix:"), value)
}








