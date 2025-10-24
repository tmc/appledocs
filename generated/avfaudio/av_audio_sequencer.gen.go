// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class AVAudioSequencer */


/* debug [class_header]: Header for AVAudioSequencer */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for AudioSequencer */
// An interface definition for the [AudioSequencer] class.
type IAudioSequencer interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for AudioSequencer */
	// properties:
	CurrentPositionInBeats() float64
	SetCurrentPositionInBeats(value float64)
	CurrentPositionInSeconds() float64
	SetCurrentPositionInSeconds(value float64)
	Playing() bool
	Rate() float32
	SetRate(value float32)
	TempoTrack() IAVMusicTrack
	Tracks() []MusicTrack
	UserInfo() foundation.IDictionary
	IsPlaying() bool
	SetIsPlaying(value bool)
	AVMusicTimeStampEndOfTrack() float64
	SetAVMusicTimeStampEndOfTrack(value float64)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for AudioSequencer */
	// methods:
	BeatsForHostTimeError(inHostTime uint64, outError objectivec.IObject) MusicTimeStamp /* typedef */
	BeatsForSeconds(seconds float64) MusicTimeStamp /* typedef */
	CreateAndAppendTrack() IMusicTrack
	DataWithSMPTEResolutionError(SMPTEResolution int, outError objectivec.IObject) foundation.Data
	HostTimeForBeatsError(inBeats MusicTimeStamp /* typedef */, outError objectivec.IObject) uint64
	LoadFromDataOptionsError(data objc.IObject /* cross-framework: NSData */, options MusicSequenceLoadOptions, outError objectivec.IObject) bool
	LoadFromURLOptionsError(fileURL objc.IObject /* cross-framework: NSURL */, options MusicSequenceLoadOptions, outError objectivec.IObject) bool
	PrepareToPlay()
	RemoveTrack(track IAVMusicTrack) bool
	ReverseEvents()
	SecondsForBeats(beats MusicTimeStamp /* typedef */) float64
	SetUserCallback(userCallback AudioSequencerUserCallback /* not a class type */)
	StartAndReturnError(outError objectivec.IObject) bool
	Stop()
	WriteToURLSMPTEResolutionReplaceExistingError(fileURL objc.IObject /* cross-framework: NSURL */, resolution int, replace bool, outError objectivec.IObject) bool
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for AudioSequencer */
// Alloc allocates a new instance without initialization.
func (ac _AudioSequencerClass) Alloc() AudioSequencer {
	rv := objc.Send[AudioSequencer](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for AudioSequencer */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for AudioSequencer */

// Creates an audio sequencer that the framework attaches to an audio engine instance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSequencer/init(audioEngine:)
func NewAudioSequencerWithAudioEngine(engine IAVAudioEngine) AudioSequencer {
	instance := getAudioSequencerClass().Alloc()
	rv := objc.Send[AudioSequencer](instance.ID, objc.Sel("initWithAudioEngine:"), engine)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewAudioSequencerWithAudioEngine */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for AudioSequencer */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for AudioSequencer */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for AudioSequencer */

// Gets the beat the system plays at the specified host time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSequencer/beats(forHostTime:error:)
func (a_ AudioSequencer) BeatsForHostTimeError(inHostTime uint64, outError objectivec.IObject) MusicTimeStamp /* typedef */ {
	rv := objc.Send[float64](a_.ID, objc.Sel("beatsForHostTime:error:"), inHostTime, outError)
	return rv
}/* debug [instance_methods/method]: BeatsForHostTimeError */


// Gets the beat position (timestamp) for the specified time in the track.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSequencer/beats(forSeconds:)
func (a_ AudioSequencer) BeatsForSeconds(seconds float64) MusicTimeStamp /* typedef */ {
	rv := objc.Send[float64](a_.ID, objc.Sel("beatsForSeconds:"), seconds)
	return rv
}/* debug [instance_methods/method]: BeatsForSeconds */


// Creates a new music track and appends it to the sequencer’s list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSequencer/createAndAppendTrack()
func (a_ AudioSequencer) CreateAndAppendTrack() IMusicTrack {
	rv := objc.Send[MusicTrack](a_.ID, objc.Sel("createAndAppendTrack"))
	return rv
}/* debug [instance_methods/method]: CreateAndAppendTrack */


// Gets a data object that contains the events from the sequence.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSequencer/data(withSMPTEResolution:error:)
func (a_ AudioSequencer) DataWithSMPTEResolutionError(SMPTEResolution int, outError objectivec.IObject) foundation.Data {
	rv := objc.Send[foundation.Data](a_.ID, objc.Sel("dataWithSMPTEResolution:error:"), SMPTEResolution, outError)
	return rv
}/* debug [instance_methods/method]: DataWithSMPTEResolutionError */


// Gets the host time the sequence plays at the specified position.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSequencer/hostTime(forBeats:error:)
func (a_ AudioSequencer) HostTimeForBeatsError(inBeats MusicTimeStamp /* typedef */, outError objectivec.IObject) uint64 {
	rv := objc.Send[uint64](a_.ID, objc.Sel("hostTimeForBeats:error:"), inBeats, outError)
	return rv
}/* debug [instance_methods/method]: HostTimeForBeatsError */


// Parses the data and adds its events to the sequence.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSequencer/load(from:options:)-8o58w
func (a_ AudioSequencer) LoadFromDataOptionsError(data objc.IObject /* cross-framework: NSData */, options MusicSequenceLoadOptions, outError objectivec.IObject) bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("loadFromData:options:error:"), data, options, outError)
	return rv
}/* debug [instance_methods/method]: LoadFromDataOptionsError */


// Loads the file the URL references and adds the events to the sequence.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSequencer/load(from:options:)-9kb6m
func (a_ AudioSequencer) LoadFromURLOptionsError(fileURL objc.IObject /* cross-framework: NSURL */, options MusicSequenceLoadOptions, outError objectivec.IObject) bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("loadFromURL:options:error:"), fileURL, options, outError)
	return rv
}/* debug [instance_methods/method]: LoadFromURLOptionsError */


// Gets ready to play the sequence by prerolling all events.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSequencer/prepareToPlay()
func (a_ AudioSequencer) PrepareToPlay() {
	objc.Send[objc.ID](a_.ID, objc.Sel("prepareToPlay"))
}/* debug [instance_methods/method]: PrepareToPlay */


// Removes the music track from the sequencer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSequencer/removeTrack(_:)
func (a_ AudioSequencer) RemoveTrack(track IAVMusicTrack) bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("removeTrack:"), track)
	return rv
}/* debug [instance_methods/method]: RemoveTrack */


// Reverses the order of all events in all music tracks, including the tempo track.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSequencer/reverseEvents()
func (a_ AudioSequencer) ReverseEvents() {
	objc.Send[objc.ID](a_.ID, objc.Sel("reverseEvents"))
}/* debug [instance_methods/method]: ReverseEvents */


// Gets the time for the specified beat position (timestamp) in the track, in seconds.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSequencer/seconds(forBeats:)
func (a_ AudioSequencer) SecondsForBeats(beats MusicTimeStamp /* typedef */) float64 {
	rv := objc.Send[float64](a_.ID, objc.Sel("secondsForBeats:"), beats)
	return rv
}/* debug [instance_methods/method]: SecondsForBeats */


// Adds a callback that the sequencer calls each time it encounters a user event during playback.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSequencer/setUserCallback(_:)
func (a_ AudioSequencer) SetUserCallback(userCallback AudioSequencerUserCallback /* not a class type */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setUserCallback:"), userCallback)
}/* debug [instance_methods/method]: SetUserCallback */


// Starts the sequencer’s player.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSequencer/start()
func (a_ AudioSequencer) StartAndReturnError(outError objectivec.IObject) bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("startAndReturnError:"), outError)
	return rv
}/* debug [instance_methods/method]: StartAndReturnError */


// Stops the sequencer’s player.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSequencer/stop()
func (a_ AudioSequencer) Stop() {
	objc.Send[objc.ID](a_.ID, objc.Sel("stop"))
}/* debug [instance_methods/method]: Stop */


// Creates and writes a MIDI file from the events in the sequence.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSequencer/write(to:smpteResolution:replaceExisting:)
func (a_ AudioSequencer) WriteToURLSMPTEResolutionReplaceExistingError(fileURL objc.IObject /* cross-framework: NSURL */, resolution int, replace bool, outError objectivec.IObject) bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("writeToURL:SMPTEResolution:replaceExisting:error:"), fileURL, resolution, replace, outError)
	return rv
}/* debug [instance_methods/method]: WriteToURLSMPTEResolutionReplaceExistingError */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for AudioSequencer */

// The current playback position, in beats.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSequencer/currentPositionInBeats
func (a_ AudioSequencer) CurrentPositionInBeats() float64 {
	rv := objc.Send[float64](a_.ID, objc.Sel("currentPositionInBeats"))
	return rv
}/* debug [instance_properties/getter]: currentPositionInBeats */


// The current playback position, in beats.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSequencer/currentPositionInBeats
func (a_ AudioSequencer) SetCurrentPositionInBeats(value float64) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setCurrentPositionInBeats:"), value)
}/* debug [instance_properties/setter]: currentPositionInBeats */


// The current playback position, in seconds.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSequencer/currentPositionInSeconds
func (a_ AudioSequencer) CurrentPositionInSeconds() float64 {
	rv := objc.Send[float64](a_.ID, objc.Sel("currentPositionInSeconds"))
	return rv
}/* debug [instance_properties/getter]: currentPositionInSeconds */


// The current playback position, in seconds.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSequencer/currentPositionInSeconds
func (a_ AudioSequencer) SetCurrentPositionInSeconds(value float64) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setCurrentPositionInSeconds:"), value)
}/* debug [instance_properties/setter]: currentPositionInSeconds */


// A Boolean value that indicates whether the sequencer’s player is in a playing state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSequencer/isPlaying
func (a_ AudioSequencer) Playing() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("playing"))
	return rv
}/* debug [instance_properties/getter]: playing */


// The playback rate of the sequencer’s player.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSequencer/rate
func (a_ AudioSequencer) Rate() float32 {
	rv := objc.Send[float32](a_.ID, objc.Sel("rate"))
	return rv
}/* debug [instance_properties/getter]: rate */


// The playback rate of the sequencer’s player.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSequencer/rate
func (a_ AudioSequencer) SetRate(value float32) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setRate:"), value)
}/* debug [instance_properties/setter]: rate */


// The track that contains tempo information about the sequence.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSequencer/tempoTrack
func (a_ AudioSequencer) TempoTrack() IAVMusicTrack {
	rv := objc.Send[MusicTrack](a_.ID, objc.Sel("tempoTrack"))
	return rv
}/* debug [instance_properties/getter]: tempoTrack */


// An array that contains all the tracks in the sequence.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSequencer/tracks
func (a_ AudioSequencer) Tracks() []MusicTrack {
	rv := objc.Send[[]MusicTrack](a_.ID, objc.Sel("tracks"))
	return rv
}/* debug [instance_properties/getter]: tracks */


// A dictionary that contains metadata from a sequence.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSequencer/userInfo
func (a_ AudioSequencer) UserInfo() foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](a_.ID, objc.Sel("userInfo"))
	return rv
}/* debug [instance_properties/getter]: userInfo */


// A Boolean value that indicates whether the sequencer’s player is in a playing state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudiosequencer/isplaying
func (a_ AudioSequencer) IsPlaying() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("isPlaying"))
	return rv
}/* debug [instance_properties/getter]: isPlaying */


// A Boolean value that indicates whether the sequencer’s player is in a playing state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudiosequencer/isplaying
func (a_ AudioSequencer) SetIsPlaying(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setIsPlaying:"), value)
}/* debug [instance_properties/setter]: isPlaying */


// A timestamp you use to access all events in a music track through a beat range.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avmusictimestampendoftrack
func (a_ AudioSequencer) AVMusicTimeStampEndOfTrack() float64 {
	rv := objc.Send[float64](a_.ID, objc.Sel("AVMusicTimeStampEndOfTrack"))
	return rv
}/* debug [instance_properties/getter]: AVMusicTimeStampEndOfTrack */


// A timestamp you use to access all events in a music track through a beat range.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avmusictimestampendoftrack
func (a_ AudioSequencer) SetAVMusicTimeStampEndOfTrack(value float64) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setAVMusicTimeStampEndOfTrack:"), value)
}/* debug [instance_properties/setter]: AVMusicTimeStampEndOfTrack */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AVAudioSequencer */


