// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [MusicTrack] class.
var (
	MusicTrackClass     _MusicTrackClass
	MusicTrackClassOnce sync.Once
)

func getMusicTrackClass() _MusicTrackClass {
	MusicTrackClassOnce.Do(func() {
		MusicTrackClass = _MusicTrackClass{objc.GetClass("AVMusicTrack")}
	})
	return MusicTrackClass
}

type _MusicTrackClass struct {
	class objc.Class
}





// An interface definition for the [MusicTrack] class.
type IMusicTrack interface {
	objectivec.IObject
	

	// properties:
	DestinationAudioUnit() IAVAudioUnit
	SetDestinationAudioUnit(value IAVAudioUnit)
	DestinationMIDIEndpoint() objectivec.IObject
	SetDestinationMIDIEndpoint(value objectivec.IObject)
	LoopingEnabled() bool
	SetLoopingEnabled(value bool)
	Muted() bool
	SetMuted(value bool)
	Soloed() bool
	SetSoloed(value bool)
	LengthInBeats() MusicTimeStamp /* typedef */
	SetLengthInBeats(value MusicTimeStamp /* typedef */)
	LengthInSeconds() float64
	SetLengthInSeconds(value float64)
	LoopRange() objc.IObject /* cross-framework: AVBeatRange */
	SetLoopRange(value objc.IObject /* cross-framework: AVBeatRange */)
	NumberOfLoops() int
	SetNumberOfLoops(value int)
	OffsetTime() MusicTimeStamp /* typedef */
	SetOffsetTime(value MusicTimeStamp /* typedef */)
	TimeResolution() uint
	UsesAutomatedParameters() bool
	SetUsesAutomatedParameters(value bool)
	AVMusicTimeStampEndOfTrack() float64
	SetAVMusicTimeStampEndOfTrack(value float64)
	IsLoopingEnabled() bool
	SetIsLoopingEnabled(value bool)
	IsMuted() bool
	SetIsMuted(value bool)
	IsSoloed() bool
	SetIsSoloed(value bool)


	

	// methods:
	AddEventAtBeat(event IAVMusicEvent, beat MusicTimeStamp /* typedef */)
	ClearEventsInRange(range_ objc.IObject /* cross-framework: AVBeatRange */)
	CopyAndMergeEventsInRangeFromTrackMergeAtBeat(range_ objc.IObject /* cross-framework: AVBeatRange */, sourceTrack IAVMusicTrack, mergeStartBeat MusicTimeStamp /* typedef */)
	CopyEventsInRangeFromTrackInsertAtBeat(range_ objc.IObject /* cross-framework: AVBeatRange */, sourceTrack IAVMusicTrack, insertStartBeat MusicTimeStamp /* typedef */)
	CutEventsInRange(range_ objc.IObject /* cross-framework: AVBeatRange */)
	EnumerateEventsInRangeUsingBlock(range_ objc.IObject /* cross-framework: AVBeatRange */, block MusicEventEnumerationBlock /* not a class type */)
	MoveEventsInRangeByAmount(range_ objc.IObject /* cross-framework: AVBeatRange */, beatAmount MusicTimeStamp /* typedef */)


}





// Alloc allocates a new instance without initialization.
func (mc _MusicTrackClass) Alloc() MusicTrack {
	rv := objc.Send[MusicTrack](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MusicTrackClass) New() MusicTrack {
	rv := objc.Send[MusicTrack](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MusicTrack) Init() MusicTrack {
	rv := objc.Send[MusicTrack](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MusicTrack) Autorelease() MusicTrack {
	rv := objc.Send[MusicTrack](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMusicTrack creates a new MusicTrack instance.
func NewMusicTrack() MusicTrack {
	return getMusicTrackClass().New()
}





// A collection of music events that you can offset, set to a muted state, modify independently from other track events, and send to a specified destination.


// A collection of music events that you can offset, set to a muted state, modify independently from other track events, and send to a specified destination.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVMusicTrack
type MusicTrack struct {
	objectivec.Object
}

// MusicTrackFrom constructs a [MusicTrack] from an unsafe.Pointer.
//
// A collection of music events that you can offset, set to a muted state, modify independently from other track events, and send to a specified destination.
func MusicTrackFrom(ptr unsafe.Pointer) MusicTrack {
	return MusicTrack{objectivec.Object{objc.ID(ptr)}}
}




















// Adds a music event to a track at the time you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVMusicTrack/addEvent(_:at:)
func (m_ MusicTrack) AddEventAtBeat(event IAVMusicEvent, beat MusicTimeStamp /* typedef */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("addEvent:atBeat:"), event, beat)
}


// Removes all events in the given beat range from the music track.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVMusicTrack/clearEvents(in:)
func (m_ MusicTrack) ClearEventsInRange(range_ objc.IObject /* cross-framework: AVBeatRange */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("clearEventsInRange:"), range_)
}


// Copies the events from the source track and merges them into the current music track.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVMusicTrack/copyAndMergeEvents(in:from:mergeAt:)
func (m_ MusicTrack) CopyAndMergeEventsInRangeFromTrackMergeAtBeat(range_ objc.IObject /* cross-framework: AVBeatRange */, sourceTrack IAVMusicTrack, mergeStartBeat MusicTimeStamp /* typedef */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("copyAndMergeEventsInRange:fromTrack:mergeAtBeat:"), range_, sourceTrack, mergeStartBeat)
}


// Copies the events from the source track and splices them into the current music track.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVMusicTrack/copyEvents(in:from:insertAt:)
func (m_ MusicTrack) CopyEventsInRangeFromTrackInsertAtBeat(range_ objc.IObject /* cross-framework: AVBeatRange */, sourceTrack IAVMusicTrack, insertStartBeat MusicTimeStamp /* typedef */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("copyEventsInRange:fromTrack:insertAtBeat:"), range_, sourceTrack, insertStartBeat)
}


// Splices all events in the beat range from the music track.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVMusicTrack/cutEvents(in:)
func (m_ MusicTrack) CutEventsInRange(range_ objc.IObject /* cross-framework: AVBeatRange */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("cutEventsInRange:"), range_)
}


// Iterates through the music events within the track.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVMusicTrack/enumerateEvents(in:using:)
func (m_ MusicTrack) EnumerateEventsInRangeUsingBlock(range_ objc.IObject /* cross-framework: AVBeatRange */, block MusicEventEnumerationBlock /* not a class type */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("enumerateEventsInRange:usingBlock:"), range_, block)
}


// Moves the beat location of all events in the given beat range by the amount you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVMusicTrack/moveEvents(in:by:)
func (m_ MusicTrack) MoveEventsInRangeByAmount(range_ objc.IObject /* cross-framework: AVBeatRange */, beatAmount MusicTimeStamp /* typedef */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("moveEventsInRange:byAmount:"), range_, beatAmount)
}







// The audio unit that receives the track’s events.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVMusicTrack/destinationAudioUnit
func (m_ MusicTrack) DestinationAudioUnit() IAVAudioUnit {
	rv := objc.Send[AudioUnit](m_.ID, objc.Sel("destinationAudioUnit"))
	return rv
}


// The audio unit that receives the track’s events.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVMusicTrack/destinationAudioUnit
func (m_ MusicTrack) SetDestinationAudioUnit(value IAVAudioUnit) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDestinationAudioUnit:"), value)
}


// The MIDI endpoint you specify as the track’s target.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVMusicTrack/destinationMIDIEndpoint
func (m_ MusicTrack) DestinationMIDIEndpoint() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](m_.ID, objc.Sel("destinationMIDIEndpoint"))
	return rv
}


// The MIDI endpoint you specify as the track’s target.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVMusicTrack/destinationMIDIEndpoint
func (m_ MusicTrack) SetDestinationMIDIEndpoint(value objectivec.IObject) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDestinationMIDIEndpoint:"), value)
}


// A Boolean value that indicates whether the track is in a looping state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVMusicTrack/isLoopingEnabled
func (m_ MusicTrack) LoopingEnabled() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("loopingEnabled"))
	return rv
}


// A Boolean value that indicates whether the track is in a looping state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVMusicTrack/isLoopingEnabled
func (m_ MusicTrack) SetLoopingEnabled(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setLoopingEnabled:"), value)
}


// A Boolean value that indicates whether the track is in a muted state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVMusicTrack/isMuted
func (m_ MusicTrack) Muted() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("muted"))
	return rv
}


// A Boolean value that indicates whether the track is in a muted state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVMusicTrack/isMuted
func (m_ MusicTrack) SetMuted(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMuted:"), value)
}


// A Boolean value that indicates whether the track is in a soloed state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVMusicTrack/isSoloed
func (m_ MusicTrack) Soloed() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("soloed"))
	return rv
}


// A Boolean value that indicates whether the track is in a soloed state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVMusicTrack/isSoloed
func (m_ MusicTrack) SetSoloed(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSoloed:"), value)
}


// The total duration of the track, in beats.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVMusicTrack/lengthInBeats
func (m_ MusicTrack) LengthInBeats() MusicTimeStamp /* typedef */ {
	rv := objc.Send[float64](m_.ID, objc.Sel("lengthInBeats"))
	return rv
}


// The total duration of the track, in beats.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVMusicTrack/lengthInBeats
func (m_ MusicTrack) SetLengthInBeats(value MusicTimeStamp /* typedef */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setLengthInBeats:"), value)
}


// The total duration of the track, in seconds.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVMusicTrack/lengthInSeconds
func (m_ MusicTrack) LengthInSeconds() float64 {
	rv := objc.Send[float64](m_.ID, objc.Sel("lengthInSeconds"))
	return rv
}


// The total duration of the track, in seconds.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVMusicTrack/lengthInSeconds
func (m_ MusicTrack) SetLengthInSeconds(value float64) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setLengthInSeconds:"), value)
}


// The timestamp range for the loop, in beats.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVMusicTrack/loopRange
func (m_ MusicTrack) LoopRange() objc.IObject /* cross-framework: AVBeatRange */ {
	rv := objc.Send[objc.ID](m_.ID, objc.Sel("loopRange"))
	return rv
}


// The timestamp range for the loop, in beats.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVMusicTrack/loopRange
func (m_ MusicTrack) SetLoopRange(value objc.IObject /* cross-framework: AVBeatRange */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setLoopRange:"), value)
}


// The number of times the track’s loop repeats.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVMusicTrack/numberOfLoops
func (m_ MusicTrack) NumberOfLoops() int {
	rv := objc.Send[int](m_.ID, objc.Sel("numberOfLoops"))
	return rv
}


// The number of times the track’s loop repeats.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVMusicTrack/numberOfLoops
func (m_ MusicTrack) SetNumberOfLoops(value int) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNumberOfLoops:"), value)
}


// The offset of the track’s start time, in beats.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVMusicTrack/offsetTime
func (m_ MusicTrack) OffsetTime() MusicTimeStamp /* typedef */ {
	rv := objc.Send[float64](m_.ID, objc.Sel("offsetTime"))
	return rv
}


// The offset of the track’s start time, in beats.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVMusicTrack/offsetTime
func (m_ MusicTrack) SetOffsetTime(value MusicTimeStamp /* typedef */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOffsetTime:"), value)
}


// The time resolution value for the sequence, in ticks (pulses) per quarter note.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVMusicTrack/timeResolution
func (m_ MusicTrack) TimeResolution() uint {
	rv := objc.Send[uint](m_.ID, objc.Sel("timeResolution"))
	return rv
}


// A Boolean value that indicates whether the track is an automation track.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVMusicTrack/usesAutomatedParameters
func (m_ MusicTrack) UsesAutomatedParameters() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("usesAutomatedParameters"))
	return rv
}


// A Boolean value that indicates whether the track is an automation track.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVMusicTrack/usesAutomatedParameters
func (m_ MusicTrack) SetUsesAutomatedParameters(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setUsesAutomatedParameters:"), value)
}


// A timestamp you use to access all events in a music track through a beat range.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avmusictimestampendoftrack
func (m_ MusicTrack) AVMusicTimeStampEndOfTrack() float64 {
	rv := objc.Send[float64](m_.ID, objc.Sel("AVMusicTimeStampEndOfTrack"))
	return rv
}


// A timestamp you use to access all events in a music track through a beat range.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avmusictimestampendoftrack
func (m_ MusicTrack) SetAVMusicTimeStampEndOfTrack(value float64) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAVMusicTimeStampEndOfTrack:"), value)
}


// A Boolean value that indicates whether the track is in a looping state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avmusictrack/isloopingenabled
func (m_ MusicTrack) IsLoopingEnabled() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("isLoopingEnabled"))
	return rv
}


// A Boolean value that indicates whether the track is in a looping state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avmusictrack/isloopingenabled
func (m_ MusicTrack) SetIsLoopingEnabled(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setIsLoopingEnabled:"), value)
}


// A Boolean value that indicates whether the track is in a muted state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avmusictrack/ismuted
func (m_ MusicTrack) IsMuted() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("isMuted"))
	return rv
}


// A Boolean value that indicates whether the track is in a muted state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avmusictrack/ismuted
func (m_ MusicTrack) SetIsMuted(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setIsMuted:"), value)
}


// A Boolean value that indicates whether the track is in a soloed state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avmusictrack/issoloed
func (m_ MusicTrack) IsSoloed() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("isSoloed"))
	return rv
}


// A Boolean value that indicates whether the track is in a soloed state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avmusictrack/issoloed
func (m_ MusicTrack) SetIsSoloed(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setIsSoloed:"), value)
}








