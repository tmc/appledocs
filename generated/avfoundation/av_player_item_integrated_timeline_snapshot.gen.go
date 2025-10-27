// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [PlayerItemIntegratedTimelineSnapshot] class.
var (
	PlayerItemIntegratedTimelineSnapshotClass     _PlayerItemIntegratedTimelineSnapshotClass
	PlayerItemIntegratedTimelineSnapshotClassOnce sync.Once
)

func getPlayerItemIntegratedTimelineSnapshotClass() _PlayerItemIntegratedTimelineSnapshotClass {
	PlayerItemIntegratedTimelineSnapshotClassOnce.Do(func() {
		PlayerItemIntegratedTimelineSnapshotClass = _PlayerItemIntegratedTimelineSnapshotClass{objc.GetClass("AVPlayerItemIntegratedTimelineSnapshot")}
	})
	return PlayerItemIntegratedTimelineSnapshotClass
}

type _PlayerItemIntegratedTimelineSnapshotClass struct {
	class objc.Class
}





// An interface definition for the [PlayerItemIntegratedTimelineSnapshot] class.
type IPlayerItemIntegratedTimelineSnapshot interface {
	objectivec.IObject
	

	// properties:
	CurrentDate() foundation.foundation.INSDate
	CurrentSegment() IAVPlayerItemSegment
	CurrentTime() objectivec.IObject
	Duration() objectivec.IObject
	Segments() []PlayerItemSegment
	CurrentSnapshot() IAVPlayerItemIntegratedTimelineSnapshot
	SetCurrentSnapshot(value IAVPlayerItemIntegratedTimelineSnapshot)


	

	// methods:
	MapTimeToSegmentAtSegmentOffset(time objectivec.IObject, timeSegmentOut IAVPlayerItemSegment, segmentOffsetOut objectivec.IObject)


}





// Alloc allocates a new instance without initialization.
func (pc _PlayerItemIntegratedTimelineSnapshotClass) Alloc() PlayerItemIntegratedTimelineSnapshot {
	rv := objc.Send[PlayerItemIntegratedTimelineSnapshot](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (pc _PlayerItemIntegratedTimelineSnapshotClass) New() PlayerItemIntegratedTimelineSnapshot {
	rv := objc.Send[PlayerItemIntegratedTimelineSnapshot](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PlayerItemIntegratedTimelineSnapshot) Init() PlayerItemIntegratedTimelineSnapshot {
	rv := objc.Send[PlayerItemIntegratedTimelineSnapshot](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PlayerItemIntegratedTimelineSnapshot) Autorelease() PlayerItemIntegratedTimelineSnapshot {
	rv := objc.Send[PlayerItemIntegratedTimelineSnapshot](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPlayerItemIntegratedTimelineSnapshot creates a new PlayerItemIntegratedTimelineSnapshot instance.
func NewPlayerItemIntegratedTimelineSnapshot() PlayerItemIntegratedTimelineSnapshot {
	return getPlayerItemIntegratedTimelineSnapshotClass().New()
}





// An immutable representation of inspectable details of an integrated timeline object.
//
// A snapshot doesn’t reflect the new timeline state as playback progresses. You can request a new snapshot instance from an that reflect the latest timeline state.


// An immutable representation of inspectable details of an integrated timeline object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemIntegratedTimelineSnapshot
type PlayerItemIntegratedTimelineSnapshot struct {
	objectivec.Object
}

// PlayerItemIntegratedTimelineSnapshotFrom constructs a [PlayerItemIntegratedTimelineSnapshot] from an unsafe.Pointer.
//
// An immutable representation of inspectable details of an integrated timeline object.
func PlayerItemIntegratedTimelineSnapshotFrom(ptr unsafe.Pointer) PlayerItemIntegratedTimelineSnapshot {
	return PlayerItemIntegratedTimelineSnapshot{objectivec.Object{objc.ID(ptr)}}
}




















// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemIntegratedTimelineSnapshot/mapTime:toSegment:atSegmentOffset:
func (p_ PlayerItemIntegratedTimelineSnapshot) MapTimeToSegmentAtSegmentOffset(time objectivec.IObject, timeSegmentOut IAVPlayerItemSegment, segmentOffsetOut objectivec.IObject) {
	objc.Send[objc.ID](p_.ID, objc.Sel("mapTime:toSegment:atSegmentOffset:"), time, timeSegmentOut, segmentOffsetOut)
}







// The current date on the integrated timeline when the system created the snapshot.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemIntegratedTimelineSnapshot/currentDate
func (p_ PlayerItemIntegratedTimelineSnapshot) CurrentDate() foundation.foundation.INSDate {
	rv := objc.Send[foundation.NSDate](p_.ID, objc.Sel("currentDate"))
	return rv
}


// The currently playing segment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemIntegratedTimelineSnapshot/currentSegment
func (p_ PlayerItemIntegratedTimelineSnapshot) CurrentSegment() IAVPlayerItemSegment {
	rv := objc.Send[PlayerItemSegment](p_.ID, objc.Sel("currentSegment"))
	return rv
}


// The current time on the integrated timeline when the system created the snapshot.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemIntegratedTimelineSnapshot/currentTime
func (p_ PlayerItemIntegratedTimelineSnapshot) CurrentTime() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](p_.ID, objc.Sel("currentTime"))
	return rv
}


// The total duration of the primary item and scheduled interstitial events.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemIntegratedTimelineSnapshot/duration
func (p_ PlayerItemIntegratedTimelineSnapshot) Duration() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](p_.ID, objc.Sel("duration"))
	return rv
}


// The segments for this snapshot.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemIntegratedTimelineSnapshot/segments
func (p_ PlayerItemIntegratedTimelineSnapshot) Segments() []PlayerItemSegment {
	rv := objc.Send[[]PlayerItemSegment](p_.ID, objc.Sel("segments"))
	return rv
}


// An immutable representation of the timeline state at time of request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritemintegratedtimeline/currentsnapshot
func (p_ PlayerItemIntegratedTimelineSnapshot) CurrentSnapshot() IAVPlayerItemIntegratedTimelineSnapshot {
	rv := objc.Send[PlayerItemIntegratedTimelineSnapshot](p_.ID, objc.Sel("currentSnapshot"))
	return rv
}


// An immutable representation of the timeline state at time of request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritemintegratedtimeline/currentsnapshot
func (p_ PlayerItemIntegratedTimelineSnapshot) SetCurrentSnapshot(value IAVPlayerItemIntegratedTimelineSnapshot) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setCurrentSnapshot:"), value)
}








