// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
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
	HostTimeForBeatsError(inBeats IMusicTimeStamp, outError unsafe.Pointer) uint64
	SecondsForBeats(beats IMusicTimeStamp) foundation.TimeInterval
	CurrentPositionInBeats() unsafe.Pointer
	SetCurrentPositionInBeats(value unsafe.Pointer)
	CurrentPositionInSeconds() unsafe.Pointer
	SetCurrentPositionInSeconds(value unsafe.Pointer)
	IsPlaying() bool
	SetIsPlaying(value bool)
	Rate() float32
	SetRate(value float32)
	TempoTrack() unsafe.Pointer
	SetTempoTrack(value unsafe.Pointer)
	Tracks() unsafe.Pointer
	SetTracks(value unsafe.Pointer)
	UserInfo() string
	SetUserInfo(value string)
	AVMusicTimeStampEndOfTrack() float64
	SetAVMusicTimeStampEndOfTrack(value float64)
}

// An object that plays audio from a collection of MIDI events the system organizes into music tracks.


// An object that plays audio from a collection of MIDI events the system organizes into music tracks.
//
// [Full Topic]
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
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSequencer/hostTime(forBeats:error:)

func (a_ AudioSequencer) HostTimeForBeatsError(inBeats IMusicTimeStamp, outError unsafe.Pointer) uint64 {
	rv := objc.Send[uint64](a_.ID, objc.Sel("hostTimeForBeats:error:"), inBeats, outError)
	return rv
}



// Gets the time for the specified beat position (timestamp) in the track, in seconds.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSequencer/seconds(forBeats:)

func (a_ AudioSequencer) SecondsForBeats(beats IMusicTimeStamp) foundation.TimeInterval {
	rv := objc.Send[foundation.TimeInterval](a_.ID, objc.Sel("secondsForBeats:"), beats)
	return rv
}


// The current playback position, in beats.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudiosequencer/currentpositioninbeats

func (a_ AudioSequencer) CurrentPositionInBeats() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("currentPositionInBeats"))
	return rv
}


// The current playback position, in beats.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudiosequencer/currentpositioninbeats

func (a_ AudioSequencer) SetCurrentPositionInBeats(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setCurrentPositionInBeats:"), value)
}


// The current playback position, in seconds.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudiosequencer/currentpositioninseconds

func (a_ AudioSequencer) CurrentPositionInSeconds() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("currentPositionInSeconds"))
	return rv
}


// The current playback position, in seconds.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudiosequencer/currentpositioninseconds

func (a_ AudioSequencer) SetCurrentPositionInSeconds(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setCurrentPositionInSeconds:"), value)
}


// A Boolean value that indicates whether the sequencer’s player is in a playing state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudiosequencer/isplaying

func (a_ AudioSequencer) IsPlaying() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("isPlaying"))
	return rv
}


// A Boolean value that indicates whether the sequencer’s player is in a playing state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudiosequencer/isplaying

func (a_ AudioSequencer) SetIsPlaying(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setIsPlaying:"), value)
}


// The playback rate of the sequencer’s player.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudiosequencer/rate

func (a_ AudioSequencer) Rate() float32 {
	rv := objc.Send[float32](a_.ID, objc.Sel("rate"))
	return rv
}


// The playback rate of the sequencer’s player.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudiosequencer/rate

func (a_ AudioSequencer) SetRate(value float32) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setRate:"), value)
}


// The track that contains tempo information about the sequence.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudiosequencer/tempotrack

func (a_ AudioSequencer) TempoTrack() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("tempoTrack"))
	return rv
}


// The track that contains tempo information about the sequence.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudiosequencer/tempotrack

func (a_ AudioSequencer) SetTempoTrack(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setTempoTrack:"), value)
}


// An array that contains all the tracks in the sequence.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudiosequencer/tracks

func (a_ AudioSequencer) Tracks() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("tracks"))
	return rv
}


// An array that contains all the tracks in the sequence.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudiosequencer/tracks

func (a_ AudioSequencer) SetTracks(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setTracks:"), value)
}


// A dictionary that contains metadata from a sequence.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudiosequencer/userinfo

func (a_ AudioSequencer) UserInfo() string {
	rv := objc.Send[string](a_.ID, objc.Sel("userInfo"))
	return rv
}


// A dictionary that contains metadata from a sequence.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudiosequencer/userinfo

func (a_ AudioSequencer) SetUserInfo(value string) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setUserInfo:"), objc.String(value))
}


// A timestamp you use to access all events in a music track through a beat range.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avmusictimestampendoftrack

func (a_ AudioSequencer) AVMusicTimeStampEndOfTrack() float64 {
	rv := objc.Send[float64](a_.ID, objc.Sel("AVMusicTimeStampEndOfTrack"))
	return rv
}


// A timestamp you use to access all events in a music track through a beat range.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avmusictimestampendoftrack

func (a_ AudioSequencer) SetAVMusicTimeStampEndOfTrack(value float64) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setAVMusicTimeStampEndOfTrack:"), value)
}



