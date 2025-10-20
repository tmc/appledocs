// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [AudioSequencer] class.
var (
	AudioSequencerClass     _AudioSequencerClass
	AudioSequencerClassOnce sync.Once
)

func getAudioSequencerClass() _AudioSequencerClass {
	AudioSequencerClassOnce.Do(func() {
		AudioSequencerClass = _AudioSequencerClass{objc.GetClass("AVAudioSequencer")}
	})
	return AudioSequencerClass
}

type _AudioSequencerClass struct {
	class objc.Class
}

// An interface definition for the [AudioSequencer] class.
type IAudioSequencer interface {
	objectivec.IObject
	HostTimeForBeatsError(inBeats unsafe.Pointer, outError unsafe.Pointer) unsafe.Pointer
	SecondsForBeats(beats unsafe.Pointer) TimeInterval
}

// An object that plays audio from a collection of MIDI events the system organizes into music tracks.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSequencer
type AudioSequencer struct {
	objectivec.Object
}

// AudioSequencerFrom constructs a [AudioSequencer] from an unsafe.Pointer.
//
// An object that plays audio from a collection of MIDI events the system organizes into music tracks.
func AudioSequencerFrom(ptr unsafe.Pointer) AudioSequencer {
	return AudioSequencer{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ac _AudioSequencerClass) Alloc() AudioSequencer {
	rv := objc.Send[AudioSequencer](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ac _AudioSequencerClass) New() AudioSequencer {
	rv := objc.Send[AudioSequencer](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AudioSequencer) Init() AudioSequencer {
	rv := objc.Send[AudioSequencer](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AudioSequencer) Autorelease() AudioSequencer {
	rv := objc.Send[AudioSequencer](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAudioSequencer creates a new AudioSequencer instance.
func NewAudioSequencer() AudioSequencer {
	return getAudioSequencerClass().New()
}


// Gets the host time the sequence plays at the specified position.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSequencer/hostTime(forBeats:error:)
func (a_ AudioSequencer) HostTimeForBeatsError(inBeats unsafe.Pointer, outError unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("hostTimeForBeats:error:"), inBeats, outError)
	return rv
}

// Gets the time for the specified beat position (timestamp) in the track, in seconds.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSequencer/seconds(forBeats:)
func (a_ AudioSequencer) SecondsForBeats(beats unsafe.Pointer) TimeInterval {
	rv := objc.Send[TimeInterval](a_.ID, objc.Sel("secondsForBeats:"), beats)
	return rv
}



