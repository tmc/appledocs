// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MIDIPlayer] class.
var (
	MIDIPlayerClass     _MIDIPlayerClass
	MIDIPlayerClassOnce sync.Once
)

func getMIDIPlayerClass() _MIDIPlayerClass {
	MIDIPlayerClassOnce.Do(func() {
		MIDIPlayerClass = _MIDIPlayerClass{objc.GetClass("AVMIDIPlayer")}
	})
	return MIDIPlayerClass
}

type _MIDIPlayerClass struct {
	class objc.Class
}

// An interface definition for the [MIDIPlayer] class.
type IMIDIPlayer interface {
	objectivec.IObject
	Play(completionHandler unsafe.Pointer)
	PrepareToPlay()
	Stop()
}

// An object that plays MIDI data through a system sound module.
//
// For more information about preparing your app to play audio, see .
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVMIDIPlayer
type MIDIPlayer struct {
	objectivec.Object
}

// MIDIPlayerFrom constructs a [MIDIPlayer] from an unsafe.Pointer.
//
// An object that plays MIDI data through a system sound module.
func MIDIPlayerFrom(ptr unsafe.Pointer) MIDIPlayer {
	return MIDIPlayer{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MIDIPlayerClass) Alloc() MIDIPlayer {
	rv := objc.Send[MIDIPlayer](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MIDIPlayerClass) New() MIDIPlayer {
	rv := objc.Send[MIDIPlayer](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MIDIPlayer) Init() MIDIPlayer {
	rv := objc.Send[MIDIPlayer](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MIDIPlayer) Autorelease() MIDIPlayer {
	rv := objc.Send[MIDIPlayer](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMIDIPlayer creates a new MIDIPlayer instance.
func NewMIDIPlayer() MIDIPlayer {
	return getMIDIPlayerClass().New()
}




// Creates a player to play a MIDI file with the specified soundbank.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVMIDIPlayer/init(contentsOf:soundBankURL:)
func NewMIDIPlayerWithContentsOfURLSoundBankURLError(inURL unsafe.Pointer, bankURL unsafe.Pointer, outError unsafe.Pointer) MIDIPlayer {
	instance := getMIDIPlayerClass().Alloc()
	rv := objc.Send[MIDIPlayer](instance.ID, objc.Sel("initWithContentsOfURL:soundBankURL:error:"), inURL, bankURL, outError)
	rv.Autorelease()
	return rv
}



// Creates a player to play MIDI data with the specified soundbank.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVMIDIPlayer/init(data:soundBankURL:)
func NewMIDIPlayerWithDataSoundBankURLError(data unsafe.Pointer, bankURL unsafe.Pointer, outError unsafe.Pointer) MIDIPlayer {
	instance := getMIDIPlayerClass().Alloc()
	rv := objc.Send[MIDIPlayer](instance.ID, objc.Sel("initWithData:soundBankURL:error:"), data, bankURL, outError)
	rv.Autorelease()
	return rv
}


// Plays the MIDI sequence.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVMIDIPlayer/play(_:)
func (m_ MIDIPlayer) Play(completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("play:"), completionHandler)
}

// Prepares the player to play the sequence by prerolling all events.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVMIDIPlayer/prepareToPlay()
func (m_ MIDIPlayer) PrepareToPlay() {
	objc.Send[objc.ID](m_.ID, objc.Sel("prepareToPlay"))
}

// Stops playing the sequence.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVMIDIPlayer/stop()
func (m_ MIDIPlayer) Stop() {
	objc.Send[objc.ID](m_.ID, objc.Sel("stop"))
}

// The current playback position, in seconds.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVMIDIPlayer/currentPosition
func (m_ MIDIPlayer) CurrentPosition() TimeInterval {
	rv := objc.Send[TimeInterval](m_.ID, objc.Sel("currentPosition"))
	return rv
}


// SetCurrentPosition sets the value of the currentPosition property.
// The current playback position, in seconds.

//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVMIDIPlayer/currentPosition
func (m_ MIDIPlayer) SetCurrentPosition(value TimeInterval) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCurrentPosition:"), value)
}

// The duration, in seconds, of the currently loaded file.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVMIDIPlayer/duration
func (m_ MIDIPlayer) Duration() TimeInterval {
	rv := objc.Send[TimeInterval](m_.ID, objc.Sel("duration"))
	return rv
}

// A Boolean value that indicates whether the sequence is playing.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVMIDIPlayer/isPlaying
func (m_ MIDIPlayer) Playing() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("playing"))
	return rv
}

// The playback rate of the player.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVMIDIPlayer/rate
func (m_ MIDIPlayer) Rate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("rate"))
	return rv
}


// SetRate sets the value of the rate property.
// The playback rate of the player.

//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVMIDIPlayer/rate
func (m_ MIDIPlayer) SetRate(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setRate:"), value)
}


