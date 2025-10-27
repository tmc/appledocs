// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [PlayerItemSegment] class.
var (
	PlayerItemSegmentClass     _PlayerItemSegmentClass
	PlayerItemSegmentClassOnce sync.Once
)

func getPlayerItemSegmentClass() _PlayerItemSegmentClass {
	PlayerItemSegmentClassOnce.Do(func() {
		PlayerItemSegmentClass = _PlayerItemSegmentClass{objc.GetClass("AVPlayerItemSegment")}
	})
	return PlayerItemSegmentClass
}

type _PlayerItemSegmentClass struct {
	class objc.Class
}





// An interface definition for the [PlayerItemSegment] class.
type IPlayerItemSegment interface {
	objectivec.IObject
	

	// properties:
	InterstitialEvent() IAVPlayerInterstitialEvent
	LoadedTimeRanges() []foundation.Value
	SegmentType() PlayerItemSegmentType
	StartDate() foundation.foundation.INSDate
	TimeMapping() objectivec.IObject
	CurrentDate() foundation.Date
	SetCurrentDate(value foundation.Date)
	CurrentSegment() IAVPlayerItemSegment
	SetCurrentSegment(value IAVPlayerItemSegment)
	CurrentTime() objectivec.IObject
	SetCurrentTime(value objectivec.IObject)
	Duration() objectivec.IObject
	SetDuration(value objectivec.IObject)
	Segments() IAVPlayerItemSegment
	SetSegments(value IAVPlayerItemSegment)


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (pc _PlayerItemSegmentClass) Alloc() PlayerItemSegment {
	rv := objc.Send[PlayerItemSegment](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (pc _PlayerItemSegmentClass) New() PlayerItemSegment {
	rv := objc.Send[PlayerItemSegment](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PlayerItemSegment) Init() PlayerItemSegment {
	rv := objc.Send[PlayerItemSegment](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PlayerItemSegment) Autorelease() PlayerItemSegment {
	rv := objc.Send[PlayerItemSegment](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPlayerItemSegment creates a new PlayerItemSegment instance.
func NewPlayerItemSegment() PlayerItemSegment {
	return getPlayerItemSegmentClass().New()
}





// An immutable object that represents a segment of time on the integrated timeline.


// An immutable object that represents a segment of time on the integrated timeline.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemSegment
type PlayerItemSegment struct {
	objectivec.Object
}

// PlayerItemSegmentFrom constructs a [PlayerItemSegment] from an unsafe.Pointer.
//
// An immutable object that represents a segment of time on the integrated timeline.
func PlayerItemSegmentFrom(ptr unsafe.Pointer) PlayerItemSegment {
	return PlayerItemSegment{objectivec.Object{objc.ID(ptr)}}
}

























// The associated interstitial event for this segment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemSegment/interstitialEvent
func (p_ PlayerItemSegment) InterstitialEvent() IAVPlayerInterstitialEvent {
	rv := objc.Send[PlayerInterstitialEvent](p_.ID, objc.Sel("interstitialEvent"))
	return rv
}


// The time ranges for the segment that have media data is readily available.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemSegment/loadedTimeRanges-2p0fl
func (p_ PlayerItemSegment) LoadedTimeRanges() []foundation.Value {
	rv := objc.Send[[]foundation.Value](p_.ID, objc.Sel("loadedTimeRanges"))
	return rv
}


// The type content this segment represents.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemSegment/segmentType-swift.property
func (p_ PlayerItemSegment) SegmentType() PlayerItemSegmentType {
	rv := objc.Send[PlayerItemSegmentType](p_.ID, objc.Sel("segmentType"))
	return rv
}


// The date at which a segment starts.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemSegment/startDate
func (p_ PlayerItemSegment) StartDate() foundation.foundation.INSDate {
	rv := objc.Send[foundation.NSDate](p_.ID, objc.Sel("startDate"))
	return rv
}


// The time mapping for this segment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemSegment/timeMapping
func (p_ PlayerItemSegment) TimeMapping() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](p_.ID, objc.Sel("timeMapping"))
	return rv
}


// The current date on the integrated timeline when the system created the snapshot.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritemintegratedtimelinesnapshot/currentdate
func (p_ PlayerItemSegment) CurrentDate() foundation.Date {
	rv := objc.Send[foundation.Date](p_.ID, objc.Sel("currentDate"))
	return rv
}


// The current date on the integrated timeline when the system created the snapshot.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritemintegratedtimelinesnapshot/currentdate
func (p_ PlayerItemSegment) SetCurrentDate(value foundation.Date) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setCurrentDate:"), value)
}


// The currently playing segment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritemintegratedtimelinesnapshot/currentsegment
func (p_ PlayerItemSegment) CurrentSegment() IAVPlayerItemSegment {
	rv := objc.Send[PlayerItemSegment](p_.ID, objc.Sel("currentSegment"))
	return rv
}


// The currently playing segment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritemintegratedtimelinesnapshot/currentsegment
func (p_ PlayerItemSegment) SetCurrentSegment(value IAVPlayerItemSegment) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setCurrentSegment:"), value)
}


// The current time on the integrated timeline when the system created the snapshot.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritemintegratedtimelinesnapshot/currenttime
func (p_ PlayerItemSegment) CurrentTime() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](p_.ID, objc.Sel("currentTime"))
	return rv
}


// The current time on the integrated timeline when the system created the snapshot.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritemintegratedtimelinesnapshot/currenttime
func (p_ PlayerItemSegment) SetCurrentTime(value objectivec.IObject) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setCurrentTime:"), value)
}


// The total duration of the primary item and scheduled interstitial events.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritemintegratedtimelinesnapshot/duration
func (p_ PlayerItemSegment) Duration() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](p_.ID, objc.Sel("duration"))
	return rv
}


// The total duration of the primary item and scheduled interstitial events.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritemintegratedtimelinesnapshot/duration
func (p_ PlayerItemSegment) SetDuration(value objectivec.IObject) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setDuration:"), value)
}


// The segments for this snapshot.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritemintegratedtimelinesnapshot/segments
func (p_ PlayerItemSegment) Segments() IAVPlayerItemSegment {
	rv := objc.Send[PlayerItemSegment](p_.ID, objc.Sel("segments"))
	return rv
}


// The segments for this snapshot.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritemintegratedtimelinesnapshot/segments
func (p_ PlayerItemSegment) SetSegments(value IAVPlayerItemSegment) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setSegments:"), value)
}








