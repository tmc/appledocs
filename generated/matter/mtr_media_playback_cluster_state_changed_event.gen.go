// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRMediaPlaybackClusterStateChangedEvent] class.
var (
	MTRMediaPlaybackClusterStateChangedEventClass     _MTRMediaPlaybackClusterStateChangedEventClass
	MTRMediaPlaybackClusterStateChangedEventClassOnce sync.Once
)

func getMTRMediaPlaybackClusterStateChangedEventClass() _MTRMediaPlaybackClusterStateChangedEventClass {
	MTRMediaPlaybackClusterStateChangedEventClassOnce.Do(func() {
		MTRMediaPlaybackClusterStateChangedEventClass = _MTRMediaPlaybackClusterStateChangedEventClass{objc.GetClass("MTRMediaPlaybackClusterStateChangedEvent")}
	})
	return MTRMediaPlaybackClusterStateChangedEventClass
}

type _MTRMediaPlaybackClusterStateChangedEventClass struct {
	class objc.Class
}

// An interface definition for the [MTRMediaPlaybackClusterStateChangedEvent] class.
type IMTRMediaPlaybackClusterStateChangedEvent interface {
	objectivec.IObject
	AudioAdvanceUnmuted() foundation.Number
	SetAudioAdvanceUnmuted(value foundation.INumber)
	CurrentState() foundation.Number
	SetCurrentState(value foundation.INumber)
	Data() foundation.NSData
	SetData(value foundation.IData)
	Duration() foundation.Number
	SetDuration(value foundation.INumber)
	PlaybackSpeed() foundation.Number
	SetPlaybackSpeed(value foundation.INumber)
	SampledPosition() MTRMediaPlaybackClusterPlaybackPositionStruct
	SetSampledPosition(value IMTRMediaPlaybackClusterPlaybackPositionStruct)
	SeekRangeEnd() foundation.Number
	SetSeekRangeEnd(value foundation.INumber)
	SeekRangeStart() foundation.Number
	SetSeekRangeStart(value foundation.INumber)
	StartTime() foundation.Number
	SetStartTime(value foundation.INumber)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMediaPlaybackClusterStateChangedEvent
type MTRMediaPlaybackClusterStateChangedEvent struct {
	objectivec.Object
}

// MTRMediaPlaybackClusterStateChangedEventFrom constructs a [MTRMediaPlaybackClusterStateChangedEvent] from an unsafe.Pointer.
func MTRMediaPlaybackClusterStateChangedEventFrom(ptr unsafe.Pointer) MTRMediaPlaybackClusterStateChangedEvent {
	return MTRMediaPlaybackClusterStateChangedEvent{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRMediaPlaybackClusterStateChangedEventClass) Alloc() MTRMediaPlaybackClusterStateChangedEvent {
	rv := objc.Send[MTRMediaPlaybackClusterStateChangedEvent](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRMediaPlaybackClusterStateChangedEventClass) New() MTRMediaPlaybackClusterStateChangedEvent {
	rv := objc.Send[MTRMediaPlaybackClusterStateChangedEvent](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRMediaPlaybackClusterStateChangedEvent) Init() MTRMediaPlaybackClusterStateChangedEvent {
	rv := objc.Send[MTRMediaPlaybackClusterStateChangedEvent](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRMediaPlaybackClusterStateChangedEvent) Autorelease() MTRMediaPlaybackClusterStateChangedEvent {
	rv := objc.Send[MTRMediaPlaybackClusterStateChangedEvent](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRMediaPlaybackClusterStateChangedEvent creates a new MTRMediaPlaybackClusterStateChangedEvent instance.
func NewMTRMediaPlaybackClusterStateChangedEvent() MTRMediaPlaybackClusterStateChangedEvent {
	return getMTRMediaPlaybackClusterStateChangedEventClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMediaPlaybackClusterStateChangedEvent/audioAdvanceUnmuted
func (m_ MTRMediaPlaybackClusterStateChangedEvent) AudioAdvanceUnmuted() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("audioAdvanceUnmuted"))
	return rv
}


// SetAudioAdvanceUnmuted sets the value of the audioAdvanceUnmuted property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMediaPlaybackClusterStateChangedEvent/audioAdvanceUnmuted
func (m_ MTRMediaPlaybackClusterStateChangedEvent) SetAudioAdvanceUnmuted(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAudioAdvanceUnmuted:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMediaPlaybackClusterStateChangedEvent/currentState
func (m_ MTRMediaPlaybackClusterStateChangedEvent) CurrentState() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("currentState"))
	return rv
}


// SetCurrentState sets the value of the currentState property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMediaPlaybackClusterStateChangedEvent/currentState
func (m_ MTRMediaPlaybackClusterStateChangedEvent) SetCurrentState(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCurrentState:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMediaPlaybackClusterStateChangedEvent/data
func (m_ MTRMediaPlaybackClusterStateChangedEvent) Data() foundation.NSData {
	rv := objc.Send[foundation.NSData](m_.ID, objc.Sel("data"))
	return rv
}


// SetData sets the value of the data property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMediaPlaybackClusterStateChangedEvent/data
func (m_ MTRMediaPlaybackClusterStateChangedEvent) SetData(value foundation.IData) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setData:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMediaPlaybackClusterStateChangedEvent/duration
func (m_ MTRMediaPlaybackClusterStateChangedEvent) Duration() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("duration"))
	return rv
}


// SetDuration sets the value of the duration property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMediaPlaybackClusterStateChangedEvent/duration
func (m_ MTRMediaPlaybackClusterStateChangedEvent) SetDuration(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDuration:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMediaPlaybackClusterStateChangedEvent/playbackSpeed
func (m_ MTRMediaPlaybackClusterStateChangedEvent) PlaybackSpeed() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("playbackSpeed"))
	return rv
}


// SetPlaybackSpeed sets the value of the playbackSpeed property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMediaPlaybackClusterStateChangedEvent/playbackSpeed
func (m_ MTRMediaPlaybackClusterStateChangedEvent) SetPlaybackSpeed(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPlaybackSpeed:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMediaPlaybackClusterStateChangedEvent/sampledPosition
func (m_ MTRMediaPlaybackClusterStateChangedEvent) SampledPosition() MTRMediaPlaybackClusterPlaybackPositionStruct {
	rv := objc.Send[MTRMediaPlaybackClusterPlaybackPositionStruct](m_.ID, objc.Sel("sampledPosition"))
	return rv
}


// SetSampledPosition sets the value of the sampledPosition property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMediaPlaybackClusterStateChangedEvent/sampledPosition
func (m_ MTRMediaPlaybackClusterStateChangedEvent) SetSampledPosition(value IMTRMediaPlaybackClusterPlaybackPositionStruct) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSampledPosition:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMediaPlaybackClusterStateChangedEvent/seekRangeEnd
func (m_ MTRMediaPlaybackClusterStateChangedEvent) SeekRangeEnd() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("seekRangeEnd"))
	return rv
}


// SetSeekRangeEnd sets the value of the seekRangeEnd property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMediaPlaybackClusterStateChangedEvent/seekRangeEnd
func (m_ MTRMediaPlaybackClusterStateChangedEvent) SetSeekRangeEnd(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSeekRangeEnd:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMediaPlaybackClusterStateChangedEvent/seekRangeStart
func (m_ MTRMediaPlaybackClusterStateChangedEvent) SeekRangeStart() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("seekRangeStart"))
	return rv
}


// SetSeekRangeStart sets the value of the seekRangeStart property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMediaPlaybackClusterStateChangedEvent/seekRangeStart
func (m_ MTRMediaPlaybackClusterStateChangedEvent) SetSeekRangeStart(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSeekRangeStart:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMediaPlaybackClusterStateChangedEvent/startTime
func (m_ MTRMediaPlaybackClusterStateChangedEvent) StartTime() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("startTime"))
	return rv
}


// SetStartTime sets the value of the startTime property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMediaPlaybackClusterStateChangedEvent/startTime
func (m_ MTRMediaPlaybackClusterStateChangedEvent) SetStartTime(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setStartTime:"), value)
}



