// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [AudioUnitGenerator] class.
var (
	AudioUnitGeneratorClass     _AudioUnitGeneratorClass
	AudioUnitGeneratorClassOnce sync.Once
)

func getAudioUnitGeneratorClass() _AudioUnitGeneratorClass {
	AudioUnitGeneratorClassOnce.Do(func() {
		AudioUnitGeneratorClass = _AudioUnitGeneratorClass{objc.GetClass("AVAudioUnitGenerator")}
	})
	return AudioUnitGeneratorClass
}

type _AudioUnitGeneratorClass struct {
	class objc.Class
}

// An interface definition for the [AudioUnitGenerator] class.
type IAudioUnitGenerator interface {
	IAudioUnit
	Bypass() bool
	SetBypass(value bool)
}

// An object that generates audio output.
//
// A generator represents an of type or . A generator has no audio input, but produces audio output. An example is a tone generator.


// An object that generates audio output.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitGenerator
type AudioUnitGenerator struct {
	AudioUnit
}

// AudioUnitGeneratorFrom constructs a [AudioUnitGenerator] from an unsafe.Pointer.
//
// An object that generates audio output.
func AudioUnitGeneratorFrom(ptr unsafe.Pointer) AudioUnitGenerator {
	return AudioUnitGenerator{
		AudioUnit: AudioUnitFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (ac _AudioUnitGeneratorClass) Alloc() AudioUnitGenerator {
	rv := objc.Send[AudioUnitGenerator](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ac _AudioUnitGeneratorClass) New() AudioUnitGenerator {
	rv := objc.Send[AudioUnitGenerator](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AudioUnitGenerator) Init() AudioUnitGenerator {
	rv := objc.Send[AudioUnitGenerator](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AudioUnitGenerator) Autorelease() AudioUnitGenerator {
	rv := objc.Send[AudioUnitGenerator](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAudioUnitGenerator creates a new AudioUnitGenerator instance.
func NewAudioUnitGenerator() AudioUnitGenerator {
	return getAudioUnitGeneratorClass().New()
}



// The bypass state of the audio unit.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudiounitgenerator/bypass
func (a_ AudioUnitGenerator) Bypass() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("bypass"))
	return rv
}


// The bypass state of the audio unit.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudiounitgenerator/bypass
func (a_ AudioUnitGenerator) SetBypass(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setBypass:"), value)
}



