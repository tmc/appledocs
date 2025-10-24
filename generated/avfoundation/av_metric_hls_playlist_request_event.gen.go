// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)





// The class instance for the [MetricHLSPlaylistRequestEvent] class.
var (
	MetricHLSPlaylistRequestEventClass     _MetricHLSPlaylistRequestEventClass
	MetricHLSPlaylistRequestEventClassOnce sync.Once
)

func getMetricHLSPlaylistRequestEventClass() _MetricHLSPlaylistRequestEventClass {
	MetricHLSPlaylistRequestEventClassOnce.Do(func() {
		MetricHLSPlaylistRequestEventClass = _MetricHLSPlaylistRequestEventClass{objc.GetClass("AVMetricHLSPlaylistRequestEvent")}
	})
	return MetricHLSPlaylistRequestEventClass
}

type _MetricHLSPlaylistRequestEventClass struct {
	class objc.Class
}





// An interface definition for the [MetricHLSPlaylistRequestEvent] class.
type IMetricHLSPlaylistRequestEvent interface {
	IMetricEvent
	

	// properties:
	IsMultivariantPlaylist() bool
	MediaResourceRequestEvent() IAVMetricMediaResourceRequestEvent
	MediaType() MediaType /* typedef */
	Url() objc.IObject /* cross-framework: NSURL */


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (mc _MetricHLSPlaylistRequestEventClass) Alloc() MetricHLSPlaylistRequestEvent {
	rv := objc.Send[MetricHLSPlaylistRequestEvent](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MetricHLSPlaylistRequestEventClass) New() MetricHLSPlaylistRequestEvent {
	rv := objc.Send[MetricHLSPlaylistRequestEvent](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MetricHLSPlaylistRequestEvent) Init() MetricHLSPlaylistRequestEvent {
	rv := objc.Send[MetricHLSPlaylistRequestEvent](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MetricHLSPlaylistRequestEvent) Autorelease() MetricHLSPlaylistRequestEvent {
	rv := objc.Send[MetricHLSPlaylistRequestEvent](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMetricHLSPlaylistRequestEvent creates a new MetricHLSPlaylistRequestEvent instance.
func NewMetricHLSPlaylistRequestEvent() MetricHLSPlaylistRequestEvent {
	return getMetricHLSPlaylistRequestEventClass().New()
}





// An event that represents a live streaming playlist resource request.


// An event that represents a live streaming playlist resource request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMetricHLSPlaylistRequestEvent
type MetricHLSPlaylistRequestEvent struct {
	MetricEvent
}

// MetricHLSPlaylistRequestEventFrom constructs a [MetricHLSPlaylistRequestEvent] from an unsafe.Pointer.
//
// An event that represents a live streaming playlist resource request.
func MetricHLSPlaylistRequestEventFrom(ptr unsafe.Pointer) MetricHLSPlaylistRequestEvent {
	return MetricHLSPlaylistRequestEvent{
		MetricEvent: MetricEventFrom(ptr),
	}
}

























// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMetricHLSPlaylistRequestEvent/isMultivariantPlaylist
func (m_ MetricHLSPlaylistRequestEvent) IsMultivariantPlaylist() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("isMultivariantPlaylist"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMetricHLSPlaylistRequestEvent/mediaResourceRequestEvent
func (m_ MetricHLSPlaylistRequestEvent) MediaResourceRequestEvent() IAVMetricMediaResourceRequestEvent {
	rv := objc.Send[MetricMediaResourceRequestEvent](m_.ID, objc.Sel("mediaResourceRequestEvent"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMetricHLSPlaylistRequestEvent/mediaType
func (m_ MetricHLSPlaylistRequestEvent) MediaType() MediaType /* typedef */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("mediaType"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMetricHLSPlaylistRequestEvent/url
func (m_ MetricHLSPlaylistRequestEvent) Url() objc.IObject /* cross-framework: NSURL */ {
	rv := objc.Send[foundation.NSURL](m_.ID, objc.Sel("url"))
	return rv
}








