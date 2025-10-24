// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
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
	

	// properties:
	CurrentDate() objc.IObject /* cross-framework: NSDate */
	CurrentSnapshot() IAVPlayerItemIntegratedTimelineSnapshot
	CurrentTime() objc.IObject /* cross-framework: Time */


	

	// methods:
	AddBoundaryTimeObserverForSegmentOffsetsIntoSegmentQueueUsingBlock(segment IAVPlayerItemSegment, offsetsIntoSegment objc.IObject /* cross-framework: NSArray */, queue objectivec.IObject, block bool) unsafe.Pointer
	AddPeriodicTimeObserverForIntervalQueueUsingBlock(interval objc.IObject /* cross-framework: Time */, queue objectivec.IObject, block unsafe.Pointer) unsafe.Pointer
	RemoveTimeObserver(observer unsafe.Pointer)
	SeekToDateCompletionHandler(date objc.IObject /* cross-framework: NSDate */, completionHandler unsafe.Pointer)
	SeekToTimeToleranceBeforeToleranceAfterCompletionHandler(time objc.IObject /* cross-framework: Time */, toleranceBefore objc.IObject /* cross-framework: Time */, toleranceAfter objc.IObject /* cross-framework: Time */, completionHandler unsafe.Pointer)


}





// Alloc allocates a new instance without initialization.
func (pc _PlayerItemIntegratedTimelineClass) Alloc() PlayerItemIntegratedTimeline {
	rv := objc.Send[PlayerItemIntegratedTimeline](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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




















// Requests invocation of a block when traversing an offset in a segment during playback.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemIntegratedTimeline/addBoundaryTimeObserverForSegment:offsetsIntoSegment:queue:usingBlock:
func (p_ PlayerItemIntegratedTimeline) AddBoundaryTimeObserverForSegmentOffsetsIntoSegmentQueueUsingBlock(segment IAVPlayerItemSegment, offsetsIntoSegment objc.IObject /* cross-framework: NSArray */, queue objectivec.IObject, block bool) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("addBoundaryTimeObserverForSegment:offsetsIntoSegment:queue:usingBlock:"), segment, offsetsIntoSegment, queue, block)
	return rv
}


// Requests invocation of a block during playback to report changing time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemIntegratedTimeline/addPeriodicTimeObserverForInterval:queue:usingBlock:
func (p_ PlayerItemIntegratedTimeline) AddPeriodicTimeObserverForIntervalQueueUsingBlock(interval objc.IObject /* cross-framework: Time */, queue objectivec.IObject, block unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("addPeriodicTimeObserverForInterval:queue:usingBlock:"), interval, queue, block)
	return rv
}


// Cancels a previously registered time observer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemIntegratedTimeline/removeTimeObserver:
func (p_ PlayerItemIntegratedTimeline) RemoveTimeObserver(observer unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("removeTimeObserver:"), observer)
}


// Seeks to a particular date in the integrated time domain.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemIntegratedTimeline/seek(to:completionHandler:)
func (p_ PlayerItemIntegratedTimeline) SeekToDateCompletionHandler(date objc.IObject /* cross-framework: NSDate */, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("seekToDate:completionHandler:"), date, completionHandler)
}


// Seeks to a particular time in the integrated time domain.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemIntegratedTimeline/seek(to:toleranceBefore:toleranceAfter:completionHandler:)
func (p_ PlayerItemIntegratedTimeline) SeekToTimeToleranceBeforeToleranceAfterCompletionHandler(time objc.IObject /* cross-framework: Time */, toleranceBefore objc.IObject /* cross-framework: Time */, toleranceAfter objc.IObject /* cross-framework: Time */, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("seekToTime:toleranceBefore:toleranceAfter:completionHandler:"), time, toleranceBefore, toleranceAfter, completionHandler)
}







// The current date of playback.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemIntegratedTimeline/currentDate
func (p_ PlayerItemIntegratedTimeline) CurrentDate() objc.IObject /* cross-framework: NSDate */ {
	rv := objc.Send[foundation.NSDate](p_.ID, objc.Sel("currentDate"))
	return rv
}


// An immutable representation of the timeline state at time of request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemIntegratedTimeline/currentSnapshot
func (p_ PlayerItemIntegratedTimeline) CurrentSnapshot() IAVPlayerItemIntegratedTimelineSnapshot {
	rv := objc.Send[PlayerItemIntegratedTimelineSnapshot](p_.ID, objc.Sel("currentSnapshot"))
	return rv
}


// The current time on the integrated timeline.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemIntegratedTimeline/currentTime
func (p_ PlayerItemIntegratedTimeline) CurrentTime() objc.IObject /* cross-framework: Time */ {
	rv := objc.Send[corevideo.Time](p_.ID, objc.Sel("currentTime"))
	return rv
}








