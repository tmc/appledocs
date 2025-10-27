// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)





// The class instance for the [MetricHLSMediaSegmentRequestEvent] class.
var (
	MetricHLSMediaSegmentRequestEventClass     _MetricHLSMediaSegmentRequestEventClass
	MetricHLSMediaSegmentRequestEventClassOnce sync.Once
)

func getMetricHLSMediaSegmentRequestEventClass() _MetricHLSMediaSegmentRequestEventClass {
	MetricHLSMediaSegmentRequestEventClassOnce.Do(func() {
		MetricHLSMediaSegmentRequestEventClass = _MetricHLSMediaSegmentRequestEventClass{objc.GetClass("AVMetricHLSMediaSegmentRequestEvent")}
	})
	return MetricHLSMediaSegmentRequestEventClass
}

type _MetricHLSMediaSegmentRequestEventClass struct {
	class objc.Class
}





// An interface definition for the [MetricHLSMediaSegmentRequestEvent] class.
type IMetricHLSMediaSegmentRequestEvent interface {
	IMetricEvent
	

	// properties:
	ByteRange() foundation.Range
	IndexFileURL() foundation.foundation.INSURL
	IsMapSegment() bool
	MediaResourceRequestEvent() IAVMetricMediaResourceRequestEvent
	MediaType() MediaType
	SegmentDuration() float64
	Url() foundation.foundation.INSURL


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (mc _MetricHLSMediaSegmentRequestEventClass) Alloc() MetricHLSMediaSegmentRequestEvent {
	rv := objc.Send[MetricHLSMediaSegmentRequestEvent](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MetricHLSMediaSegmentRequestEventClass) New() MetricHLSMediaSegmentRequestEvent {
	rv := objc.Send[MetricHLSMediaSegmentRequestEvent](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MetricHLSMediaSegmentRequestEvent) Init() MetricHLSMediaSegmentRequestEvent {
	rv := objc.Send[MetricHLSMediaSegmentRequestEvent](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MetricHLSMediaSegmentRequestEvent) Autorelease() MetricHLSMediaSegmentRequestEvent {
	rv := objc.Send[MetricHLSMediaSegmentRequestEvent](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMetricHLSMediaSegmentRequestEvent creates a new MetricHLSMediaSegmentRequestEvent instance.
func NewMetricHLSMediaSegmentRequestEvent() MetricHLSMediaSegmentRequestEvent {
	return getMetricHLSMediaSegmentRequestEventClass().New()
}





// An event that represents a live streaming media segment resource request.


// An event that represents a live streaming media segment resource request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMetricHLSMediaSegmentRequestEvent
type MetricHLSMediaSegmentRequestEvent struct {
	MetricEvent
}

// MetricHLSMediaSegmentRequestEventFrom constructs a [MetricHLSMediaSegmentRequestEvent] from an unsafe.Pointer.
//
// An event that represents a live streaming media segment resource request.
func MetricHLSMediaSegmentRequestEventFrom(ptr unsafe.Pointer) MetricHLSMediaSegmentRequestEvent {
	return MetricHLSMediaSegmentRequestEvent{
		MetricEvent: MetricEventFrom(ptr),
	}
}

























// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMetricHLSMediaSegmentRequestEvent/byteRange
func (m_ MetricHLSMediaSegmentRequestEvent) ByteRange() foundation.Range {
	rv := objc.Send[foundation.Range](m_.ID, objc.Sel("byteRange"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMetricHLSMediaSegmentRequestEvent/indexFileURL
func (m_ MetricHLSMediaSegmentRequestEvent) IndexFileURL() foundation.foundation.INSURL {
	rv := objc.Send[foundation.NSURL](m_.ID, objc.Sel("indexFileURL"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMetricHLSMediaSegmentRequestEvent/isMapSegment
func (m_ MetricHLSMediaSegmentRequestEvent) IsMapSegment() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("isMapSegment"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMetricHLSMediaSegmentRequestEvent/mediaResourceRequestEvent
func (m_ MetricHLSMediaSegmentRequestEvent) MediaResourceRequestEvent() IAVMetricMediaResourceRequestEvent {
	rv := objc.Send[MetricMediaResourceRequestEvent](m_.ID, objc.Sel("mediaResourceRequestEvent"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMetricHLSMediaSegmentRequestEvent/mediaType
func (m_ MetricHLSMediaSegmentRequestEvent) MediaType() MediaType {
	rv := objc.Send[MediaType](m_.ID, objc.Sel("mediaType"))
	return rv
}


// Returns the duration of segment in seconds.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMetricHLSMediaSegmentRequestEvent/segmentDuration
func (m_ MetricHLSMediaSegmentRequestEvent) SegmentDuration() float64 {
	rv := objc.Send[float64](m_.ID, objc.Sel("segmentDuration"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMetricHLSMediaSegmentRequestEvent/url
func (m_ MetricHLSMediaSegmentRequestEvent) Url() foundation.foundation.INSURL {
	rv := objc.Send[foundation.NSURL](m_.ID, objc.Sel("url"))
	return rv
}








