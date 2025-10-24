// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corevideo"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class AVPlayerItemIntegratedTimelineSnapshot */


/* debug [class_header]: Header for AVPlayerItemIntegratedTimelineSnapshot */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for PlayerItemIntegratedTimelineSnapshot */
// An interface definition for the [PlayerItemIntegratedTimelineSnapshot] class.
type IPlayerItemIntegratedTimelineSnapshot interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for PlayerItemIntegratedTimelineSnapshot */
	// properties:
	CurrentDate() objc.IObject /* cross-framework: NSDate */
	CurrentSegment() IAVPlayerItemSegment
	CurrentTime() objc.IObject /* cross-framework: Time */
	Duration() objc.IObject /* cross-framework: Time */
	Segments() []PlayerItemSegment
	CurrentSnapshot() IAVPlayerItemIntegratedTimelineSnapshot
	SetCurrentSnapshot(value IAVPlayerItemIntegratedTimelineSnapshot)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for PlayerItemIntegratedTimelineSnapshot */
	// methods:
	MapTimeToSegmentAtSegmentOffset(time objc.IObject /* cross-framework: Time */, timeSegmentOut IAVPlayerItemSegment, segmentOffsetOut objc.IObject /* cross-framework: Time */)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for PlayerItemIntegratedTimelineSnapshot */
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for PlayerItemIntegratedTimelineSnapshot */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for PlayerItemIntegratedTimelineSnapshot *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for PlayerItemIntegratedTimelineSnapshot */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for PlayerItemIntegratedTimelineSnapshot */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for PlayerItemIntegratedTimelineSnapshot */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemIntegratedTimelineSnapshot/mapTime:toSegment:atSegmentOffset:
func (p_ PlayerItemIntegratedTimelineSnapshot) MapTimeToSegmentAtSegmentOffset(time objc.IObject /* cross-framework: Time */, timeSegmentOut IAVPlayerItemSegment, segmentOffsetOut objc.IObject /* cross-framework: Time */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("mapTime:toSegment:atSegmentOffset:"), time, timeSegmentOut, segmentOffsetOut)
}/* debug [instance_methods/method]: MapTimeToSegmentAtSegmentOffset */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for PlayerItemIntegratedTimelineSnapshot */

// The current date on the integrated timeline when the system created the snapshot.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemIntegratedTimelineSnapshot/currentDate
func (p_ PlayerItemIntegratedTimelineSnapshot) CurrentDate() objc.IObject /* cross-framework: NSDate */ {
	rv := objc.Send[foundation.NSDate](p_.ID, objc.Sel("currentDate"))
	return rv
}/* debug [instance_properties/getter]: currentDate */


// The currently playing segment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemIntegratedTimelineSnapshot/currentSegment
func (p_ PlayerItemIntegratedTimelineSnapshot) CurrentSegment() IAVPlayerItemSegment {
	rv := objc.Send[PlayerItemSegment](p_.ID, objc.Sel("currentSegment"))
	return rv
}/* debug [instance_properties/getter]: currentSegment */


// The current time on the integrated timeline when the system created the snapshot.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemIntegratedTimelineSnapshot/currentTime
func (p_ PlayerItemIntegratedTimelineSnapshot) CurrentTime() objc.IObject /* cross-framework: Time */ {
	rv := objc.Send[corevideo.Time](p_.ID, objc.Sel("currentTime"))
	return rv
}/* debug [instance_properties/getter]: currentTime */


// The total duration of the primary item and scheduled interstitial events.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemIntegratedTimelineSnapshot/duration
func (p_ PlayerItemIntegratedTimelineSnapshot) Duration() objc.IObject /* cross-framework: Time */ {
	rv := objc.Send[corevideo.Time](p_.ID, objc.Sel("duration"))
	return rv
}/* debug [instance_properties/getter]: duration */


// The segments for this snapshot.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemIntegratedTimelineSnapshot/segments
func (p_ PlayerItemIntegratedTimelineSnapshot) Segments() []PlayerItemSegment {
	rv := objc.Send[[]PlayerItemSegment](p_.ID, objc.Sel("segments"))
	return rv
}/* debug [instance_properties/getter]: segments */


// An immutable representation of the timeline state at time of request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritemintegratedtimeline/currentsnapshot
func (p_ PlayerItemIntegratedTimelineSnapshot) CurrentSnapshot() IAVPlayerItemIntegratedTimelineSnapshot {
	rv := objc.Send[PlayerItemIntegratedTimelineSnapshot](p_.ID, objc.Sel("currentSnapshot"))
	return rv
}/* debug [instance_properties/getter]: currentSnapshot */


// An immutable representation of the timeline state at time of request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritemintegratedtimeline/currentsnapshot
func (p_ PlayerItemIntegratedTimelineSnapshot) SetCurrentSnapshot(value IAVPlayerItemIntegratedTimelineSnapshot) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setCurrentSnapshot:"), value)
}/* debug [instance_properties/setter]: currentSnapshot */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AVPlayerItemIntegratedTimelineSnapshot */



