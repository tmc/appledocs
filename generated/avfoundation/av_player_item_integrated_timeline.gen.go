// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [AVPlayerItemIntegratedTimeline] class.
var (
	aVPlayerItemIntegratedTimelineClass     _AVPlayerItemIntegratedTimelineClass
	aVPlayerItemIntegratedTimelineClassOnce sync.Once
)

func getAVPlayerItemIntegratedTimelineClass() _AVPlayerItemIntegratedTimelineClass {
	aVPlayerItemIntegratedTimelineClassOnce.Do(func() {
		aVPlayerItemIntegratedTimelineClass = _AVPlayerItemIntegratedTimelineClass{objc.GetClass("AVPlayerItemIntegratedTimeline")}
	})
	return aVPlayerItemIntegratedTimelineClass
}

type _AVPlayerItemIntegratedTimelineClass struct {
	class objc.Class
}

// An interface definition for the [AVPlayerItemIntegratedTimeline] class.
type IAVPlayerItemIntegratedTimeline interface {
	objectivec.IObject
}

// An object that models the timeline and playback sequence of a primary player item and scheduled interstitial events.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemIntegratedTimeline
type AVPlayerItemIntegratedTimeline struct {
	objectivec.Object
}

// AVPlayerItemIntegratedTimelineFrom constructs a [AVPlayerItemIntegratedTimeline] from an unsafe.Pointer.
//
// An object that models the timeline and playback sequence of a primary player item and scheduled interstitial events.
func AVPlayerItemIntegratedTimelineFrom(ptr unsafe.Pointer) AVPlayerItemIntegratedTimeline {
	return AVPlayerItemIntegratedTimeline{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ac _AVPlayerItemIntegratedTimelineClass) Alloc() AVPlayerItemIntegratedTimeline {
	rv := objc.Send[AVPlayerItemIntegratedTimeline](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ac _AVPlayerItemIntegratedTimelineClass) New() AVPlayerItemIntegratedTimeline {
	rv := objc.Send[AVPlayerItemIntegratedTimeline](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AVPlayerItemIntegratedTimeline) Init() AVPlayerItemIntegratedTimeline {
	rv := objc.Send[AVPlayerItemIntegratedTimeline](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AVPlayerItemIntegratedTimeline) Autorelease() AVPlayerItemIntegratedTimeline {
	rv := objc.Send[AVPlayerItemIntegratedTimeline](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVPlayerItemIntegratedTimeline creates a new AVPlayerItemIntegratedTimeline instance.
func NewAVPlayerItemIntegratedTimeline() AVPlayerItemIntegratedTimeline {
	return getAVPlayerItemIntegratedTimelineClass().New()
}




