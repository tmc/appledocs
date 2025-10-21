// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
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
}

// An immutable representation of inspectable details of an integrated timeline object.
//
// A snapshot doesn’t reflect the new timeline state as playback progresses. You can request a new snapshot instance from an that reflect the latest timeline state.
//
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

// Alloc allocates a new instance without initialization.
func (pc _PlayerItemIntegratedTimelineSnapshotClass) Alloc() PlayerItemIntegratedTimelineSnapshot {
	rv := objc.Send[PlayerItemIntegratedTimelineSnapshot](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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


// The currently playing segment.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemIntegratedTimelineSnapshot/currentSegment
func (p_ PlayerItemIntegratedTimelineSnapshot) CurrentSegment() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("currentSegment"))
	return rv
}

// The current time on the integrated timeline when the system created the snapshot.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemIntegratedTimelineSnapshot/currentTime
func (p_ PlayerItemIntegratedTimelineSnapshot) CurrentTime() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("currentTime"))
	return rv
}

// The segments for this snapshot.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemIntegratedTimelineSnapshot/segments
func (p_ PlayerItemIntegratedTimelineSnapshot) Segments() []unsafe.Pointer {
	rv := objc.Send[[]unsafe.Pointer](p_.ID, objc.Sel("segments"))
	return rv
}

// An immutable representation of the timeline state at time of request.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritemintegratedtimeline/currentsnapshot
func (p_ PlayerItemIntegratedTimelineSnapshot) CurrentSnapshot() AVPlayerItemIntegratedTimelineSnapshot {
	rv := objc.Send[AVPlayerItemIntegratedTimelineSnapshot](p_.ID, objc.Sel("currentSnapshot"))
	return rv
}


// SetCurrentSnapshot sets the value of the currentSnapshot property.
// An immutable representation of the timeline state at time of request.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritemintegratedtimeline/currentsnapshot
func (p_ PlayerItemIntegratedTimelineSnapshot) SetCurrentSnapshot(value IAVPlayerItemIntegratedTimelineSnapshot) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setCurrentSnapshot:"), value)
}

// The current date on the integrated timeline when the system created the snapshot.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritemintegratedtimelinesnapshot/currentdate
func (p_ PlayerItemIntegratedTimelineSnapshot) CurrentDate() foundation.Date {
	rv := objc.Send[foundation.Date](p_.ID, objc.Sel("currentDate"))
	return rv
}


// SetCurrentDate sets the value of the currentDate property.
// The current date on the integrated timeline when the system created the snapshot.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritemintegratedtimelinesnapshot/currentdate
func (p_ PlayerItemIntegratedTimelineSnapshot) SetCurrentDate(value foundation.IDate) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setCurrentDate:"), value)
}

// The total duration of the primary item and scheduled interstitial events.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritemintegratedtimelinesnapshot/duration
func (p_ PlayerItemIntegratedTimelineSnapshot) Duration() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("duration"))
	return rv
}


// SetDuration sets the value of the duration property.
// The total duration of the primary item and scheduled interstitial events.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritemintegratedtimelinesnapshot/duration
func (p_ PlayerItemIntegratedTimelineSnapshot) SetDuration(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setDuration:"), value)
}



