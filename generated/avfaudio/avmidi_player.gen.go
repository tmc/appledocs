// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class AVMIDIPlayer */


/* debug [class_header]: Header for AVMIDIPlayer */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MIDIPlayer */
// An interface definition for the [MIDIPlayer] class.
type IMIDIPlayer interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MIDIPlayer */
	// properties:
	CurrentPosition() float64
	SetCurrentPosition(value float64)
	Duration() float64
	Playing() bool
	Rate() float32
	SetRate(value float32)
	IsPlaying() bool
	SetIsPlaying(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MIDIPlayer */
	// methods:
	Play(completionHandler MIDIPlayerCompletionHandler /* not a class type */)
	PrepareToPlay()
	Stop()
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MIDIPlayer */
// Alloc allocates a new instance without initialization.
func (mc _MIDIPlayerClass) Alloc() MIDIPlayer {
	rv := objc.Send[MIDIPlayer](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MIDIPlayer */
// An object that plays MIDI data through a system sound module.
//
// For more information about preparing your app to play audio, see .


// An object that plays MIDI data through a system sound module.
//
// [Full Topic]
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MIDIPlayer */

// Creates a player to play a MIDI file with the specified soundbank.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVMIDIPlayer/init(contentsOf:soundBankURL:)
func NewMIDIPlayerWithContentsOfURLSoundBankURLError(inURL objc.IObject /* cross-framework: NSURL */, bankURL objc.IObject /* cross-framework: NSURL */, outError objectivec.IObject) MIDIPlayer {
	instance := getMIDIPlayerClass().Alloc()
	rv := objc.Send[MIDIPlayer](instance.ID, objc.Sel("initWithContentsOfURL:soundBankURL:error:"), inURL, bankURL, outError)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMIDIPlayerWithContentsOfURLSoundBankURLError */


// Creates a player to play MIDI data with the specified soundbank.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVMIDIPlayer/init(data:soundBankURL:)
func NewMIDIPlayerWithDataSoundBankURLError(data objc.IObject /* cross-framework: NSData */, bankURL objc.IObject /* cross-framework: NSURL */, outError objectivec.IObject) MIDIPlayer {
	instance := getMIDIPlayerClass().Alloc()
	rv := objc.Send[MIDIPlayer](instance.ID, objc.Sel("initWithData:soundBankURL:error:"), data, bankURL, outError)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMIDIPlayerWithDataSoundBankURLError */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MIDIPlayer */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MIDIPlayer */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MIDIPlayer */

// Plays the MIDI sequence.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVMIDIPlayer/play(_:)
func (m_ MIDIPlayer) Play(completionHandler MIDIPlayerCompletionHandler /* not a class type */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("play:"), completionHandler)
}/* debug [instance_methods/method]: Play */


// Prepares the player to play the sequence by prerolling all events.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVMIDIPlayer/prepareToPlay()
func (m_ MIDIPlayer) PrepareToPlay() {
	objc.Send[objc.ID](m_.ID, objc.Sel("prepareToPlay"))
}/* debug [instance_methods/method]: PrepareToPlay */


// Stops playing the sequence.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVMIDIPlayer/stop()
func (m_ MIDIPlayer) Stop() {
	objc.Send[objc.ID](m_.ID, objc.Sel("stop"))
}/* debug [instance_methods/method]: Stop */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MIDIPlayer */

// The current playback position, in seconds.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVMIDIPlayer/currentPosition
func (m_ MIDIPlayer) CurrentPosition() float64 {
	rv := objc.Send[float64](m_.ID, objc.Sel("currentPosition"))
	return rv
}/* debug [instance_properties/getter]: currentPosition */


// The current playback position, in seconds.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVMIDIPlayer/currentPosition
func (m_ MIDIPlayer) SetCurrentPosition(value float64) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCurrentPosition:"), value)
}/* debug [instance_properties/setter]: currentPosition */


// The duration, in seconds, of the currently loaded file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVMIDIPlayer/duration
func (m_ MIDIPlayer) Duration() float64 {
	rv := objc.Send[float64](m_.ID, objc.Sel("duration"))
	return rv
}/* debug [instance_properties/getter]: duration */


// A Boolean value that indicates whether the sequence is playing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVMIDIPlayer/isPlaying
func (m_ MIDIPlayer) Playing() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("playing"))
	return rv
}/* debug [instance_properties/getter]: playing */


// The playback rate of the player.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVMIDIPlayer/rate
func (m_ MIDIPlayer) Rate() float32 {
	rv := objc.Send[float32](m_.ID, objc.Sel("rate"))
	return rv
}/* debug [instance_properties/getter]: rate */


// The playback rate of the player.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVMIDIPlayer/rate
func (m_ MIDIPlayer) SetRate(value float32) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setRate:"), value)
}/* debug [instance_properties/setter]: rate */


// A Boolean value that indicates whether the sequence is playing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avmidiplayer/isplaying
func (m_ MIDIPlayer) IsPlaying() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("isPlaying"))
	return rv
}/* debug [instance_properties/getter]: isPlaying */


// A Boolean value that indicates whether the sequence is playing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avmidiplayer/isplaying
func (m_ MIDIPlayer) SetIsPlaying(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setIsPlaying:"), value)
}/* debug [instance_properties/setter]: isPlaying */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AVMIDIPlayer */


