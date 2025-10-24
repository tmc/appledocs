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
	// properties:
	AudioAdvanceUnmuted() objc.IObject /* cross-framework: NSNumber */
	SetAudioAdvanceUnmuted(value objc.IObject /* cross-framework: NSNumber */)
	CurrentState() objc.IObject /* cross-framework: NSNumber */
	SetCurrentState(value objc.IObject /* cross-framework: NSNumber */)
	Data() objc.IObject /* cross-framework: NSData */
	SetData(value objc.IObject /* cross-framework: NSData */)
	Duration() objc.IObject /* cross-framework: NSNumber */
	SetDuration(value objc.IObject /* cross-framework: NSNumber */)
	PlaybackSpeed() objc.IObject /* cross-framework: NSNumber */
	SetPlaybackSpeed(value objc.IObject /* cross-framework: NSNumber */)
	SampledPosition() IMTRMediaPlaybackClusterPlaybackPositionStruct
	SetSampledPosition(value IMTRMediaPlaybackClusterPlaybackPositionStruct)
	SeekRangeEnd() objc.IObject /* cross-framework: NSNumber */
	SetSeekRangeEnd(value objc.IObject /* cross-framework: NSNumber */)
	SeekRangeStart() objc.IObject /* cross-framework: NSNumber */
	SetSeekRangeStart(value objc.IObject /* cross-framework: NSNumber */)
	StartTime() objc.IObject /* cross-framework: NSNumber */
	SetStartTime(value objc.IObject /* cross-framework: NSNumber */)
	// methods:
}



// [Full Topic]
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



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMediaPlaybackClusterStateChangedEvent/audioAdvanceUnmuted
func (m_ MTRMediaPlaybackClusterStateChangedEvent) AudioAdvanceUnmuted() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("audioAdvanceUnmuted"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMediaPlaybackClusterStateChangedEvent/audioAdvanceUnmuted
func (m_ MTRMediaPlaybackClusterStateChangedEvent) SetAudioAdvanceUnmuted(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAudioAdvanceUnmuted:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMediaPlaybackClusterStateChangedEvent/currentState
func (m_ MTRMediaPlaybackClusterStateChangedEvent) CurrentState() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("currentState"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMediaPlaybackClusterStateChangedEvent/currentState
func (m_ MTRMediaPlaybackClusterStateChangedEvent) SetCurrentState(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCurrentState:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMediaPlaybackClusterStateChangedEvent/data
func (m_ MTRMediaPlaybackClusterStateChangedEvent) Data() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](m_.ID, objc.Sel("data"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMediaPlaybackClusterStateChangedEvent/data
func (m_ MTRMediaPlaybackClusterStateChangedEvent) SetData(value objc.IObject /* cross-framework: NSData */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setData:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMediaPlaybackClusterStateChangedEvent/duration
func (m_ MTRMediaPlaybackClusterStateChangedEvent) Duration() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("duration"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMediaPlaybackClusterStateChangedEvent/duration
func (m_ MTRMediaPlaybackClusterStateChangedEvent) SetDuration(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDuration:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMediaPlaybackClusterStateChangedEvent/playbackSpeed
func (m_ MTRMediaPlaybackClusterStateChangedEvent) PlaybackSpeed() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("playbackSpeed"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMediaPlaybackClusterStateChangedEvent/playbackSpeed
func (m_ MTRMediaPlaybackClusterStateChangedEvent) SetPlaybackSpeed(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPlaybackSpeed:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMediaPlaybackClusterStateChangedEvent/sampledPosition
func (m_ MTRMediaPlaybackClusterStateChangedEvent) SampledPosition() IMTRMediaPlaybackClusterPlaybackPositionStruct {
	rv := objc.Send[MTRMediaPlaybackClusterPlaybackPositionStruct](m_.ID, objc.Sel("sampledPosition"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMediaPlaybackClusterStateChangedEvent/sampledPosition
func (m_ MTRMediaPlaybackClusterStateChangedEvent) SetSampledPosition(value IMTRMediaPlaybackClusterPlaybackPositionStruct) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSampledPosition:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMediaPlaybackClusterStateChangedEvent/seekRangeEnd
func (m_ MTRMediaPlaybackClusterStateChangedEvent) SeekRangeEnd() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("seekRangeEnd"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMediaPlaybackClusterStateChangedEvent/seekRangeEnd
func (m_ MTRMediaPlaybackClusterStateChangedEvent) SetSeekRangeEnd(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSeekRangeEnd:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMediaPlaybackClusterStateChangedEvent/seekRangeStart
func (m_ MTRMediaPlaybackClusterStateChangedEvent) SeekRangeStart() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("seekRangeStart"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMediaPlaybackClusterStateChangedEvent/seekRangeStart
func (m_ MTRMediaPlaybackClusterStateChangedEvent) SetSeekRangeStart(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSeekRangeStart:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMediaPlaybackClusterStateChangedEvent/startTime
func (m_ MTRMediaPlaybackClusterStateChangedEvent) StartTime() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("startTime"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMediaPlaybackClusterStateChangedEvent/startTime
func (m_ MTRMediaPlaybackClusterStateChangedEvent) SetStartTime(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setStartTime:"), value)
}



