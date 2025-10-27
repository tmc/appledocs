// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)





// The class instance for the [MetricPlayerItemInitialLikelyToKeepUpEvent] class.
var (
	MetricPlayerItemInitialLikelyToKeepUpEventClass     _MetricPlayerItemInitialLikelyToKeepUpEventClass
	MetricPlayerItemInitialLikelyToKeepUpEventClassOnce sync.Once
)

func getMetricPlayerItemInitialLikelyToKeepUpEventClass() _MetricPlayerItemInitialLikelyToKeepUpEventClass {
	MetricPlayerItemInitialLikelyToKeepUpEventClassOnce.Do(func() {
		MetricPlayerItemInitialLikelyToKeepUpEventClass = _MetricPlayerItemInitialLikelyToKeepUpEventClass{objc.GetClass("AVMetricPlayerItemInitialLikelyToKeepUpEvent")}
	})
	return MetricPlayerItemInitialLikelyToKeepUpEventClass
}

type _MetricPlayerItemInitialLikelyToKeepUpEventClass struct {
	class objc.Class
}





// An interface definition for the [MetricPlayerItemInitialLikelyToKeepUpEvent] class.
type IMetricPlayerItemInitialLikelyToKeepUpEvent interface {
	IMetricPlayerItemLikelyToKeepUpEvent
	

	// properties:
	ContentKeyRequestEvents() []MetricContentKeyRequestEvent
	MediaSegmentRequestEvents() []MetricHLSMediaSegmentRequestEvent
	PlaylistRequestEvents() []MetricHLSPlaylistRequestEvent


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (mc _MetricPlayerItemInitialLikelyToKeepUpEventClass) Alloc() MetricPlayerItemInitialLikelyToKeepUpEvent {
	rv := objc.Send[MetricPlayerItemInitialLikelyToKeepUpEvent](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MetricPlayerItemInitialLikelyToKeepUpEventClass) New() MetricPlayerItemInitialLikelyToKeepUpEvent {
	rv := objc.Send[MetricPlayerItemInitialLikelyToKeepUpEvent](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MetricPlayerItemInitialLikelyToKeepUpEvent) Init() MetricPlayerItemInitialLikelyToKeepUpEvent {
	rv := objc.Send[MetricPlayerItemInitialLikelyToKeepUpEvent](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MetricPlayerItemInitialLikelyToKeepUpEvent) Autorelease() MetricPlayerItemInitialLikelyToKeepUpEvent {
	rv := objc.Send[MetricPlayerItemInitialLikelyToKeepUpEvent](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMetricPlayerItemInitialLikelyToKeepUpEvent creates a new MetricPlayerItemInitialLikelyToKeepUpEvent instance.
func NewMetricPlayerItemInitialLikelyToKeepUpEvent() MetricPlayerItemInitialLikelyToKeepUpEvent {
	return getMetricPlayerItemInitialLikelyToKeepUpEventClass().New()
}





// An event that represents the initial state for whether playback is likely to continue without stalling.


// An event that represents the initial state for whether playback is likely to continue without stalling.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMetricPlayerItemInitialLikelyToKeepUpEvent
type MetricPlayerItemInitialLikelyToKeepUpEvent struct {
	MetricPlayerItemLikelyToKeepUpEvent
}

// MetricPlayerItemInitialLikelyToKeepUpEventFrom constructs a [MetricPlayerItemInitialLikelyToKeepUpEvent] from an unsafe.Pointer.
//
// An event that represents the initial state for whether playback is likely to continue without stalling.
func MetricPlayerItemInitialLikelyToKeepUpEventFrom(ptr unsafe.Pointer) MetricPlayerItemInitialLikelyToKeepUpEvent {
	return MetricPlayerItemInitialLikelyToKeepUpEvent{
		MetricPlayerItemLikelyToKeepUpEvent: MetricPlayerItemLikelyToKeepUpEventFrom(ptr),
	}
}

























// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMetricPlayerItemInitialLikelyToKeepUpEvent/contentKeyRequestEvents
func (m_ MetricPlayerItemInitialLikelyToKeepUpEvent) ContentKeyRequestEvents() []MetricContentKeyRequestEvent {
	rv := objc.Send[[]MetricContentKeyRequestEvent](m_.ID, objc.Sel("contentKeyRequestEvents"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMetricPlayerItemInitialLikelyToKeepUpEvent/mediaSegmentRequestEvents
func (m_ MetricPlayerItemInitialLikelyToKeepUpEvent) MediaSegmentRequestEvents() []MetricHLSMediaSegmentRequestEvent {
	rv := objc.Send[[]MetricHLSMediaSegmentRequestEvent](m_.ID, objc.Sel("mediaSegmentRequestEvents"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMetricPlayerItemInitialLikelyToKeepUpEvent/playlistRequestEvents
func (m_ MetricPlayerItemInitialLikelyToKeepUpEvent) PlaylistRequestEvents() []MetricHLSPlaylistRequestEvent {
	rv := objc.Send[[]MetricHLSPlaylistRequestEvent](m_.ID, objc.Sel("playlistRequestEvents"))
	return rv
}








