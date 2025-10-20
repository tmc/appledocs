// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [AudioUnit] class.
var (
	AudioUnitClass     _AudioUnitClass
	AudioUnitClassOnce sync.Once
)

func getAudioUnitClass() _AudioUnitClass {
	AudioUnitClassOnce.Do(func() {
		AudioUnitClass = _AudioUnitClass{objc.GetClass("AVAudioUnit")}
	})
	return AudioUnitClass
}

type _AudioUnitClass struct {
	class objc.Class
}

// An interface definition for the [AudioUnit] class.
type IAudioUnit interface {
	IAudioNode
	LoadAudioUnitPresetAtURLError(url unsafe.Pointer, outError unsafe.Pointer) bool
}

// A subclass of the audio node class that, processes audio either in real time or nonreal time, depending on the type of the audio unit.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioUnit
type AudioUnit struct {
	AudioNode
}

// AudioUnitFrom constructs a [AudioUnit] from an unsafe.Pointer.
//
// A subclass of the audio node class that, processes audio either in real time or nonreal time, depending on the type of the audio unit.
func AudioUnitFrom(ptr unsafe.Pointer) AudioUnit {
	return AudioUnit{
		AudioNode: AudioNodeFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (ac _AudioUnitClass) Alloc() AudioUnit {
	rv := objc.Send[AudioUnit](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ac _AudioUnitClass) New() AudioUnit {
	rv := objc.Send[AudioUnit](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AudioUnit) Init() AudioUnit {
	rv := objc.Send[AudioUnit](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AudioUnit) Autorelease() AudioUnit {
	rv := objc.Send[AudioUnit](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAudioUnit creates a new AudioUnit instance.
func NewAudioUnit() AudioUnit {
	return getAudioUnitClass().New()
}


// Loads an audio unit using a specified preset.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioUnit/loadPreset(at:)
func (a_ AudioUnit) LoadAudioUnitPresetAtURLError(url unsafe.Pointer, outError unsafe.Pointer) bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("loadAudioUnitPresetAtURL:error:"), url, outError)
	return rv
}

// The underlying Core Audio audio unit.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioUnit/audioUnit
func (a_ AudioUnit) AudioUnit() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("audioUnit"))
	return rv
}



