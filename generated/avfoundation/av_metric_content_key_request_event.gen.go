// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)





// The class instance for the [MetricContentKeyRequestEvent] class.
var (
	MetricContentKeyRequestEventClass     _MetricContentKeyRequestEventClass
	MetricContentKeyRequestEventClassOnce sync.Once
)

func getMetricContentKeyRequestEventClass() _MetricContentKeyRequestEventClass {
	MetricContentKeyRequestEventClassOnce.Do(func() {
		MetricContentKeyRequestEventClass = _MetricContentKeyRequestEventClass{objc.GetClass("AVMetricContentKeyRequestEvent")}
	})
	return MetricContentKeyRequestEventClass
}

type _MetricContentKeyRequestEventClass struct {
	class objc.Class
}





// An interface definition for the [MetricContentKeyRequestEvent] class.
type IMetricContentKeyRequestEvent interface {
	IMetricEvent
	

	// properties:
	ContentKeySpecifier() IAVContentKeySpecifier
	IsClientInitiated() bool
	MediaResourceRequestEvent() IAVMetricMediaResourceRequestEvent
	MediaType() MediaType /* typedef */


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (mc _MetricContentKeyRequestEventClass) Alloc() MetricContentKeyRequestEvent {
	rv := objc.Send[MetricContentKeyRequestEvent](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MetricContentKeyRequestEventClass) New() MetricContentKeyRequestEvent {
	rv := objc.Send[MetricContentKeyRequestEvent](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MetricContentKeyRequestEvent) Init() MetricContentKeyRequestEvent {
	rv := objc.Send[MetricContentKeyRequestEvent](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MetricContentKeyRequestEvent) Autorelease() MetricContentKeyRequestEvent {
	rv := objc.Send[MetricContentKeyRequestEvent](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMetricContentKeyRequestEvent creates a new MetricContentKeyRequestEvent instance.
func NewMetricContentKeyRequestEvent() MetricContentKeyRequestEvent {
	return getMetricContentKeyRequestEventClass().New()
}





// An event that represents a live streaming content key resource request.


// An event that represents a live streaming content key resource request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMetricContentKeyRequestEvent
type MetricContentKeyRequestEvent struct {
	MetricEvent
}

// MetricContentKeyRequestEventFrom constructs a [MetricContentKeyRequestEvent] from an unsafe.Pointer.
//
// An event that represents a live streaming content key resource request.
func MetricContentKeyRequestEventFrom(ptr unsafe.Pointer) MetricContentKeyRequestEvent {
	return MetricContentKeyRequestEvent{
		MetricEvent: MetricEventFrom(ptr),
	}
}

























// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMetricContentKeyRequestEvent/contentKeySpecifier
func (m_ MetricContentKeyRequestEvent) ContentKeySpecifier() IAVContentKeySpecifier {
	rv := objc.Send[ContentKeySpecifier](m_.ID, objc.Sel("contentKeySpecifier"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMetricContentKeyRequestEvent/isClientInitiated
func (m_ MetricContentKeyRequestEvent) IsClientInitiated() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("isClientInitiated"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMetricContentKeyRequestEvent/mediaResourceRequestEvent
func (m_ MetricContentKeyRequestEvent) MediaResourceRequestEvent() IAVMetricMediaResourceRequestEvent {
	rv := objc.Send[MetricMediaResourceRequestEvent](m_.ID, objc.Sel("mediaResourceRequestEvent"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMetricContentKeyRequestEvent/mediaType
func (m_ MetricContentKeyRequestEvent) MediaType() MediaType /* typedef */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("mediaType"))
	return rv
}








