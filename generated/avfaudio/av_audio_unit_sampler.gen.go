// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [AudioUnitSampler] class.
var (
	AudioUnitSamplerClass     _AudioUnitSamplerClass
	AudioUnitSamplerClassOnce sync.Once
)

func getAudioUnitSamplerClass() _AudioUnitSamplerClass {
	AudioUnitSamplerClassOnce.Do(func() {
		AudioUnitSamplerClass = _AudioUnitSamplerClass{objc.GetClass("AVAudioUnitSampler")}
	})
	return AudioUnitSamplerClass
}

type _AudioUnitSamplerClass struct {
	class objc.Class
}

// An interface definition for the [AudioUnitSampler] class.
type IAudioUnitSampler interface {
	IAudioUnitMIDIInstrument
}

// An object that you configure with one or more instrument samples, based on Apple’s Sampler audio unit.
//
// An is an for Apple’s Sampler audio unit. You configure the sampler by loading instruments from different types of files. These include an file, DLS, or SF2 sound bank; an EXS24 instrument; a single audio file; or an array of audio files. The output of a is a single stereo bus.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitSampler
type AudioUnitSampler struct {
	AudioUnitMIDIInstrument
}

// AudioUnitSamplerFrom constructs a [AudioUnitSampler] from an unsafe.Pointer.
//
// An object that you configure with one or more instrument samples, based on Apple’s Sampler audio unit.
func AudioUnitSamplerFrom(ptr unsafe.Pointer) AudioUnitSampler {
	return AudioUnitSampler{
		AudioUnitMIDIInstrument: AudioUnitMIDIInstrumentFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (ac _AudioUnitSamplerClass) Alloc() AudioUnitSampler {
	rv := objc.Send[AudioUnitSampler](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ac _AudioUnitSamplerClass) New() AudioUnitSampler {
	rv := objc.Send[AudioUnitSampler](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AudioUnitSampler) Init() AudioUnitSampler {
	rv := objc.Send[AudioUnitSampler](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AudioUnitSampler) Autorelease() AudioUnitSampler {
	rv := objc.Send[AudioUnitSampler](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAudioUnitSampler creates a new AudioUnitSampler instance.
func NewAudioUnitSampler() AudioUnitSampler {
	return getAudioUnitSamplerClass().New()
}




