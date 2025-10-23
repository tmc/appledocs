// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [AudioUnitMIDIInstrument] class.
var (
	AudioUnitMIDIInstrumentClass     _AudioUnitMIDIInstrumentClass
	AudioUnitMIDIInstrumentClassOnce sync.Once
)

func getAudioUnitMIDIInstrumentClass() _AudioUnitMIDIInstrumentClass {
	AudioUnitMIDIInstrumentClassOnce.Do(func() {
		AudioUnitMIDIInstrumentClass = _AudioUnitMIDIInstrumentClass{objc.GetClass("AVAudioUnitMIDIInstrument")}
	})
	return AudioUnitMIDIInstrumentClass
}

type _AudioUnitMIDIInstrumentClass struct {
	class objc.Class
}

// An interface definition for the [AudioUnitMIDIInstrument] class.
type IAudioUnitMIDIInstrument interface {
	objectivec.IObject
}

// A parent class referenced by other AVFAudio classes.


// A parent class referenced by other AVFAudio classes. [Full Topic]
type AudioUnitMIDIInstrument struct {
	objectivec.Object
}

// AudioUnitMIDIInstrumentFrom constructs a [AudioUnitMIDIInstrument] from an unsafe.Pointer.
//
// A parent class referenced by other AVFAudio classes.
func AudioUnitMIDIInstrumentFrom(ptr unsafe.Pointer) AudioUnitMIDIInstrument {
	return AudioUnitMIDIInstrument{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ac _AudioUnitMIDIInstrumentClass) Alloc() AudioUnitMIDIInstrument {
	rv := objc.Send[AudioUnitMIDIInstrument](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ac _AudioUnitMIDIInstrumentClass) New() AudioUnitMIDIInstrument {
	rv := objc.Send[AudioUnitMIDIInstrument](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AudioUnitMIDIInstrument) Init() AudioUnitMIDIInstrument {
	rv := objc.Send[AudioUnitMIDIInstrument](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AudioUnitMIDIInstrument) Autorelease() AudioUnitMIDIInstrument {
	rv := objc.Send[AudioUnitMIDIInstrument](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAudioUnitMIDIInstrument creates a new AudioUnitMIDIInstrument instance.
func NewAudioUnitMIDIInstrument() AudioUnitMIDIInstrument {
	return getAudioUnitMIDIInstrumentClass().New()
}




