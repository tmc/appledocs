// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [PlayerItemIntegratedTimeline] class.
var (
	PlayerItemIntegratedTimelineClass     _PlayerItemIntegratedTimelineClass
	PlayerItemIntegratedTimelineClassOnce sync.Once
)

func getPlayerItemIntegratedTimelineClass() _PlayerItemIntegratedTimelineClass {
	PlayerItemIntegratedTimelineClassOnce.Do(func() {
		PlayerItemIntegratedTimelineClass = _PlayerItemIntegratedTimelineClass{objc.GetClass("AVPlayerItemIntegratedTimeline")}
	})
	return PlayerItemIntegratedTimelineClass
}

type _PlayerItemIntegratedTimelineClass struct {
	class objc.Class
}

// An interface definition for the [PlayerItemIntegratedTimeline] class.
type IPlayerItemIntegratedTimeline interface {
	objectivec.IObject
	CurrentDate() foundation.Date
	SetCurrentDate(value foundation.Date)
	CurrentSnapshot() PlayerItemIntegratedTimelineSnapshot
	SetCurrentSnapshot(value PlayerItemIntegratedTimelineSnapshot)
	CurrentTime() unsafe.Pointer
	SetCurrentTime(value unsafe.Pointer)
}

// An object that models the timeline and playback sequence of a primary player item and scheduled interstitial events.
//
// The timeline models all regions to traverse during playback. A player may not present portions of the primary item when exiting an interstitial event with a positive resumption offset.


// An object that models the timeline and playback sequence of a primary player item and scheduled interstitial events.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemIntegratedTimeline
type PlayerItemIntegratedTimeline struct {
	objectivec.Object
}

// PlayerItemIntegratedTimelineFrom constructs a [PlayerItemIntegratedTimeline] from an unsafe.Pointer.
//
// An object that models the timeline and playback sequence of a primary player item and scheduled interstitial events.
func PlayerItemIntegratedTimelineFrom(ptr unsafe.Pointer) PlayerItemIntegratedTimeline {
	return PlayerItemIntegratedTimeline{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (pc _PlayerItemIntegratedTimelineClass) Alloc() PlayerItemIntegratedTimeline {
	rv := objc.Send[PlayerItemIntegratedTimeline](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _PlayerItemIntegratedTimelineClass) New() PlayerItemIntegratedTimeline {
	rv := objc.Send[PlayerItemIntegratedTimeline](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PlayerItemIntegratedTimeline) Init() PlayerItemIntegratedTimeline {
	rv := objc.Send[PlayerItemIntegratedTimeline](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PlayerItemIntegratedTimeline) Autorelease() PlayerItemIntegratedTimeline {
	rv := objc.Send[PlayerItemIntegratedTimeline](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPlayerItemIntegratedTimeline creates a new PlayerItemIntegratedTimeline instance.
func NewPlayerItemIntegratedTimeline() PlayerItemIntegratedTimeline {
	return getPlayerItemIntegratedTimelineClass().New()
}



// The current date of playback.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritemintegratedtimeline/currentdate
func (p_ PlayerItemIntegratedTimeline) CurrentDate() foundation.Date {
	rv := objc.Send[foundation.Date](p_.ID, objc.Sel("currentDate"))
	return rv
}


// The current date of playback.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritemintegratedtimeline/currentdate
func (p_ PlayerItemIntegratedTimeline) SetCurrentDate(value foundation.Date) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setCurrentDate:"), value)
}


// An immutable representation of the timeline state at time of request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritemintegratedtimeline/currentsnapshot
func (p_ PlayerItemIntegratedTimeline) CurrentSnapshot() PlayerItemIntegratedTimelineSnapshot {
	rv := objc.Send[PlayerItemIntegratedTimelineSnapshot](p_.ID, objc.Sel("currentSnapshot"))
	return rv
}


// An immutable representation of the timeline state at time of request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritemintegratedtimeline/currentsnapshot
func (p_ PlayerItemIntegratedTimeline) SetCurrentSnapshot(value PlayerItemIntegratedTimelineSnapshot) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setCurrentSnapshot:"), value)
}


// The current time on the integrated timeline.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritemintegratedtimeline/currenttime
func (p_ PlayerItemIntegratedTimeline) CurrentTime() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("currentTime"))
	return rv
}


// The current time on the integrated timeline.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritemintegratedtimeline/currenttime
func (p_ PlayerItemIntegratedTimeline) SetCurrentTime(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setCurrentTime:"), value)
}



